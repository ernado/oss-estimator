package speed

import (
	"context"
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"go.uber.org/atomic"
)

type Metric struct {
	last  time.Time
	read  atomic.Uint64
	total atomic.Uint64
}

func (s *Metric) Report(ctx context.Context, name string) func() error {
	return func() error {
		fmt.Println("starting reporting")
		t := time.NewTicker(time.Millisecond * 300)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-t.C:
				fmt.Println(name, s.ConsumeSpeed())
			}
		}
	}
}

func (s *Metric) Write(p []byte) (n int, err error) {
	n = len(p)
	s.Add(uint64(n))
	return n, nil
}

func (s *Metric) Add(n uint64) {
	s.read.Add(n)
	s.total.Add(n)
}

func (s *Metric) Consume() uint64 {
	n := s.read.Load()
	s.Add(-n)
	return n
}

func (s *Metric) Rate() float64 {
	now := time.Now()
	delta := now.Sub(s.last)
	s.last = now
	n := s.Consume()

	opsPerSec := float64(n) / delta.Seconds()

	return opsPerSec
}

func (s *Metric) ConsumeSpeed() string {
	now := time.Now()
	delta := now.Sub(s.last)
	s.last = now
	n := s.Consume()

	bytesPerSec := float64(n) / delta.Seconds()

	return fmt.Sprintf("%s / sec", humanize.Bytes(uint64(bytesPerSec)))
}

func NewMetric() *Metric {
	return &Metric{last: time.Now()}
}
