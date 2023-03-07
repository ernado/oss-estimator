package entry

import (
	"bufio"
	"context"
	"io"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/klauspost/compress/zstd"

	"estimator/internal/speed"
)

type Reader struct {
	metric *speed.Metric
	e      Event
	j      *jx.Decoder
	z      *zstd.Decoder

	buf []byte
}

func (r *Reader) Close() {
	r.e.Reset()
	r.j = nil
	r.z.Close()
	r.buf = nil
}

type Decode struct {
	Reader  io.Reader
	OnEvent func(ctx context.Context, e *Event) error
	OnError func(ctx context.Context, data []byte, err error) error
}

func (r *Reader) Decode(ctx context.Context, d Decode) error {
	if err := r.z.Reset(d.Reader); err != nil {
		return errors.Wrap(err, "zstd reset")
	}

	s := bufio.NewScanner(r.z)
	s.Buffer(r.buf, len(r.buf))

	for s.Scan() {
		if err := ctx.Err(); err != nil {
			return err
		}
		buf := s.Bytes()
		r.j.ResetBytes(buf)
		r.metric.Add(uint64(len(buf)))
		r.e.Reset()
		if err := r.e.Decode(r.j); err != nil {
			if f := d.OnError; f != nil {
				if err := f(ctx, buf, err); err != nil {
					return errors.Wrap(err, "error callback")
				}
			}
			continue
		}
		if f := d.OnEvent; f != nil {
			if err := f(ctx, &r.e); err != nil {
				return errors.Wrap(err, "event callback")
			}
		}
	}

	return s.Err()
}

func NewReader(m *speed.Metric) *Reader {
	z, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
	if err != nil {
		panic(err)
	}
	return &Reader{
		metric: m,
		buf:    make([]byte, 1024*1024*150),
		z:      z,
		j:      jx.GetDecoder(),
	}
}
