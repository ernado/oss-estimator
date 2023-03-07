package archive

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-faster/errors"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"
	ht "github.com/ogen-go/ogen/http"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"estimator/internal/entry"
)

// Client to gh archive.
type Client struct {
	api ht.Client
	dir string
	lg  *zap.Logger
}

func New(api ht.Client, dir string, lg *zap.Logger) *Client {
	return &Client{
		api: api,
		dir: dir,
		lg:  lg,
	}
}

type Options struct {
	Progress *Progress
	Key      string
}

type Result struct {
	Path string

	SizeOutput  int64
	SizeInput   int64
	SizeContent int64

	SHA256Content string
	SHA256Input   string
	SHA256Output  string
}

type Hash struct {
	io.Writer

	SHA256 hash.Hash
}

func hexHashField(name string, h hash.Hash) zap.Field {
	return zap.String(name, hex.EncodeToString(h.Sum(nil)))
}

func hexHash(h hash.Hash) string {
	return hex.EncodeToString(h.Sum(nil))
}

func (h Hash) Fields() []zap.Field {
	return []zap.Field{
		hexHashField("sha256", h.SHA256),
	}
}

func NewHash() Hash {
	h := Hash{
		SHA256: sha256.New(),
	}
	h.Writer = io.MultiWriter(
		h.SHA256,
	)
	return h
}

func GetURL(date time.Time) string {
	// You can't describe 24-hour without leading zero.
	// Correct format is yyyy-MM-dd-H.
	return fmt.Sprintf("https://data.gharchive.org/%s-%d.json.gz",
		date.Format("2006-01-02"), date.Hour(),
	)
}

type Progress struct {
	// Cancel is called when scope of download is done.
	Cancel context.CancelFunc

	// total is total input size in bytes.
	total atomic.Int64
	// ready is downloaded input size in bytes.
	ready atomic.Int64
	last  atomic.Int64
}

// Total input size in bytes.
func (p *Progress) Total() int64 {
	if p == nil {
		return 0
	}
	return p.total.Load()
}

// Consume last progress in bytes.
//
// If zero, no progress were made.
func (p *Progress) Consume() int64 {
	if p == nil {
		return 0
	}

	// No need for CAS here, not being atomic is OK,
	// we miss a couple of writes in worst case.
	last := p.last.Load()
	p.last.Store(0)

	return last
}

func (p *Progress) ReadyBytes() int64 {
	return p.ready.Load()
}

// Ready returns [0..1] value of readiness.
func (p *Progress) Ready() float64 {
	if p == nil {
		return 0
	}
	var (
		done  = p.ready.Load()
		total = p.total.Load()
	)
	if total == 0 {
		return 0
	}
	return float64(done) / float64(total)
}

func (p *Progress) setTotal(n int64) {
	if p == nil {
		return
	}
	p.total.Store(n)
}

type writerFunc func(b []byte) (n int, err error)

func (f writerFunc) Write(b []byte) (n int, err error) { return f(b) }

func (p *Progress) tee() writerFunc {
	if p == nil {
		return func(b []byte) (n int, err error) {
			return len(b), nil
		}
	}
	return func(b []byte) (n int, err error) {
		p.ready.Add(int64(len(b)))
		return len(b), nil
	}
}

func (p *Progress) Tee() io.Writer {
	return p.tee()
}

func (p *Progress) Write(b []byte) (n int, err error) {
	p.last.Add(int64(len(b)))
	return len(b), nil
}

func (p *Progress) done() {
	if p == nil || p.Cancel == nil {
		return
	}
	p.Cancel()
}

var ErrNotFound = errors.New("not found")

func (c *Client) Inventory(ctx context.Context, key string) (*Result, error) {
	lg := c.lg.With(
		zap.String("key", key),
	)

	date, err := time.Parse(entry.Layout, key)
	if err != nil {
		return nil, errors.Wrap(err, "key parse")
	}

	outName := fmt.Sprintf("%s.json.zst", date.Format(entry.Layout))
	outPath := filepath.Join(c.dir, outName)
	f, err := os.Open(outPath)
	if os.IsNotExist(err) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "open")
	}
	defer func() {
		_ = f.Close()
	}()

	link := GetURL(date)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, link, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "req")
	}
	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("%s: bad code %d", link, res.StatusCode)
	}
	lg.Info("HEAD",
		zap.String("url", link),
		zap.Duration("duration", time.Since(start).Round(time.Millisecond)),
	)

	stat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "stat")
	}

	r, err := zstd.NewReader(f)
	if err != nil {
		return nil, errors.Wrap(err, "zstd")
	}
	defer r.Close()

	sizeContent, err := io.Copy(io.Discard, r)
	if err != nil {
		return nil, errors.Wrap(err, "copy")
	}

	return &Result{
		SizeContent: sizeContent,
		SizeInput:   res.ContentLength,
		SizeOutput:  stat.Size(),
	}, nil
}

func (c *Client) Download(ctx context.Context, opt Options) (*Result, error) {
	defer opt.Progress.done()

	lg := c.lg.With(
		zap.String("key", opt.Key),
	)

	if err := os.MkdirAll(c.dir, 0755); err != nil {
		return nil, errors.Wrap(err, "mkdir")
	}

	date, err := time.Parse(entry.Layout, opt.Key)
	if err != nil {
		return nil, errors.Wrap(err, "key parse")
	}

	link := GetURL(date)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "req")
	}
	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("%s: bad code %d", link, res.StatusCode)
	}

	opt.Progress.setTotal(res.ContentLength)
	lg.Info("Found",
		zap.String("method", req.Method),
		zap.String("url", link),
		zap.Duration("duration", time.Since(start).Round(time.Millisecond)),
	)

	var (
		hInput   = NewHash() // hash for input
		hContent = NewHash() // hash for json
		hOutput  = NewHash() // hash for output
	)

	// Reading body as gzip.
	bodyBuf := bufio.NewReader(res.Body)
	reader, err := gzip.NewReader(
		io.TeeReader(bodyBuf,
			// Use TeeReader to track input hash and progress.
			io.MultiWriter(hInput, opt.Progress.Tee()),
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "gzip")
	}

	outName := fmt.Sprintf("%s.json.zst", date.Format(entry.Layout))
	outPath := filepath.Join(c.dir, outName)
	out, err := os.Create(outPath)
	if err != nil {
		return nil, errors.Wrap(err, "create file")
	}
	defer func() {
		_ = out.Close()
	}()
	defer func() {
		if err == nil {
			// Cleaning up only on failure.
			return
		}

		lg.Warn("Cleaning up file")
		if err := os.Remove(outPath); err != nil {
			lg.Error("Failed to cleanup", zap.Error(err))
		}
	}()

	outWriter, err := zstd.NewWriter(
		io.MultiWriter(out, hOutput, opt.Progress),
	)
	if err != nil {
		return nil, errors.Wrap(err, "zstd writer init")
	}

	// Up to 1 MiB for copy.
	buf := make([]byte, 1024*1024)
	content, err := io.CopyBuffer(io.MultiWriter(outWriter, hContent), reader, buf)
	if err != nil {
		return nil, errors.Wrap(err, "copy")
	}

	if err := outWriter.Close(); err != nil {
		return nil, errors.Wrap(err, "zstd")
	}

	// Calculating re-compression ratio for fun and profit.
	stat, err := out.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "stat")
	}
	ratio := float64(res.ContentLength) / float64(stat.Size())
	absRatio := float64(content) / float64(stat.Size())

	if err := out.Close(); err != nil {
		return nil, errors.Wrap(err, "close file")
	}

	lg.Info("Processed",
		zap.String("path", outPath),
		zap.String("date", date.Format(entry.Layout)),
		zap.Int64("bytes_output", stat.Size()),
		zap.Int64("bytes_content", content),
		zap.Int64("bytes_input", res.ContentLength),
		zap.String("relative_ratio", fmt.Sprintf("%.0f%%", ratio*100)),
		zap.String("absolute_ratio", fmt.Sprintf("%.0f%%", absRatio*100)),
		zap.Duration("duration", time.Since(start).Round(time.Millisecond)),
	)

	lg.Info("Input", hInput.Fields()...)
	lg.Info("Content", hContent.Fields()...)
	lg.Info("Output", hOutput.Fields()...)

	return &Result{
		Path: outPath,

		SizeContent: content,
		SizeInput:   res.ContentLength,
		SizeOutput:  stat.Size(),

		SHA256Content: hexHash(hContent.SHA256),
		SHA256Input:   hexHash(hInput.SHA256),
		SHA256Output:  hexHash(hOutput.SHA256),
	}, nil
}
