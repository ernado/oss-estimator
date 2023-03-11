package archive

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/klauspost/compress/zstd"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/entry"
)

func Format(date time.Time) string {
	return fmt.Sprintf("%s-%d", date.Format("2006-01-02"), date.Hour())
}

// GetURL returns GitHub archive URL for the given date.
func GetURL(date time.Time) string {
	// You can't describe 24-hour without leading zero.
	// Correct format is yyyy-MM-dd-H.
	return fmt.Sprintf("https://data.gharchive.org/%s.json.gz", Format(date))
}

type Downloader struct {
	uc  *UserCache
	lg  *zap.Logger
	dir string
}

func NewDownloader(lg *zap.Logger, uc *UserCache, dir string) *Downloader {
	return &Downloader{
		uc:  uc,
		lg:  lg,
		dir: dir,
	}
}

const EventDDL = `'WatchEvent'=1, 'PushEvent'=2, 'IssuesEvent'=3, 'PullRequestEvent'=4`

func (d *Downloader) Download(ctx context.Context, t time.Time) (err error) {
	targetDIR := filepath.Join(d.dir, "cache", t.Format("2006-01"))
	if err := os.MkdirAll(targetDIR, 0755); err != nil {
		return errors.Wrap(err, "mkdir")
	}
	outPathTarget := filepath.Join(targetDIR, fmt.Sprintf("%s.native.zst", t.Format("2006-01-02-15")))
	outPathTmp := outPathTarget + ".tmp"
	defer func() {
		if err != nil {
			_ = os.Remove(outPathTmp)
			_ = os.Remove(outPathTarget)
		}
	}()

	if stat, err := os.Stat(outPathTarget); err == nil && stat.Size() > 0 {
		// Already processed.
		d.lg.Info("Hit", zap.String("k", Format(t)))
		return nil
	}

	outFile, err := os.Create(outPathTmp)
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer func() { _ = outFile.Close() }()

	outWrite, err := zstd.NewWriter(outFile,
		zstd.WithEncoderConcurrency(1),
		zstd.WithEncoderLevel(zstd.SpeedBetterCompression),
	)
	if err != nil {
		return errors.Wrap(err, "zstd")
	}

	g, ctx := errgroup.WithContext(ctx)

	d.lg.Info("Download", zap.String("k", Format(t)))
	wget := exec.CommandContext(ctx, "wget",
		"-nv", "-O", "-", GetURL(t),
	)
	wget.Stderr = os.Stderr
	dl, err := wget.StdoutPipe()
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "zstd", "-d")
	cmd.Stdin = dl
	cmd.Stderr = os.Stderr
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	done := make(chan struct{})

	g.Go(func() error {
		if err := cmd.Start(); err != nil {
			return errors.Wrap(err, "zstd")
		}
		select {
		case <-ctx.Done():
			if err := cmd.Wait(); err != nil {
				return errors.Wrap(err, "zstd")
			}
		case <-done:
		}
		return nil
	})
	g.Go(func() error {
		if err := wget.Start(); err != nil {
			return errors.Wrap(err, "wget")
		}
		select {
		case <-ctx.Done():
			if err := wget.Wait(); err != nil {
				return errors.Wrap(err, "wget")
			}
		case <-done:
		}
		return nil
	})
	g.Go(func() error {
		defer close(done)
		s := bufio.NewScanner(out)
		var (
			j  jx.Decoder
			e  jx.Encoder
			ev entry.Event

			colEv      proto.ColEnum8
			colRepoID  proto.ColInt64
			colActorID proto.ColInt64
			colTime    proto.ColDateTime

			outBuf proto.Buffer
		)

		const maxTokSize = 1024 * 1024 * 150
		buf := make([]byte, 0, maxTokSize)
		s.Buffer(buf, len(buf))
		input := []proto.InputColumn{
			{Name: "event", Data: proto.Wrap(&colEv, EventDDL)},
			{Name: "repo", Data: &colRepoID},
			{Name: "actor", Data: &colActorID},
			{Name: "time", Data: &colTime},
		}
		write := func() error {
			b := proto.Block{Rows: colEv.Rows(), Columns: 4}
			if err := b.EncodeRawBlock(&outBuf, 54451, input); err != nil {
				return errors.Wrap(err, "encode")
			}
			if _, err := outWrite.Write(outBuf.Buf); err != nil {
				return errors.Wrap(err, "write")
			}
			outBuf.Reset()
			proto.Reset(
				&colEv,
				&colRepoID,
				&colActorID,
				&colTime,
			)
			return nil
		}

		var (
			noActorID int
			good      int
			total     int
			decodeErr int
		)

		for s.Scan() {
			if ctx.Err() != nil {
				return ctx.Err()
			}

			e.Reset()
			ev.Reset()

			if colEv.Rows() > 10_000 {
				if err := write(); err != nil {
					return errors.Wrap(err, "write")
				}
			}

			data := s.Bytes()
			j.ResetBytes(data)

			if err := ev.Decode(&j); err != nil {
				d.lg.Error("decode", zap.Error(err))
				decodeErr++
				continue
			}
			total++
			if !ev.Interesting() {
				continue
			}
			if ev.ActorID == 0 {
				actorID, err := d.uc.Get(ev.Actor)
				if err != nil {
					return errors.Wrap(err, "get")
				}
				ev.ActorID = actorID
			}
			if ev.ActorID == 0 {
				noActorID++
				continue
			}
			if err := d.uc.Add(ev.ActorID, ev.Actor); err != nil {
				return errors.Wrap(err, "add")
			}
			{
				good++
				colEv.Append(proto.Enum8(ev.Type))
				colRepoID.Append(ev.RepoID)
				colActorID.Append(ev.ActorID)
				colTime.Append(ev.Time)
			}
		}
		if err := s.Err(); err != nil {
			return errors.Wrap(err, "scan")
		}
		// Flush remaining data.
		if err := write(); err != nil {
			return errors.Wrap(err, "write")
		}
		if err := outWrite.Flush(); err != nil {
			return errors.Wrap(err, "flush")
		}
		if err := outWrite.Close(); err != nil {
			return errors.Wrap(err, "close")
		}
		if err := outFile.Sync(); err != nil {
			return errors.Wrap(err, "sync")
		}
		if err := outFile.Close(); err != nil {
			return errors.Wrap(err, "close")
		}
		if err := os.Rename(outPathTmp, outPathTarget); err != nil {
			return errors.Wrap(err, "rename")
		}
		d.lg.Info("Processed",
			zap.Int("no_actor_id", noActorID),
			zap.Int("good", good),
			zap.Int("total", total),
			zap.Int("err", decodeErr),
			zap.String("k", Format(t)),
		)

		return nil
	})

	return g.Wait()
}
