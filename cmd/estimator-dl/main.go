package main

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/entry"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		u := archive.GetURL(time.Now().AddDate(0, 0, -2))

		g, ctx := errgroup.WithContext(ctx)

		wget := exec.CommandContext(ctx, "wget",
			"-nv", "-O", "-", u,
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
				d jx.Decoder
				e entry.Event
			)
			buf := make([]byte, 0, 1024*1024)
			s.Buffer(buf, len(buf))
			for s.Scan() {
				e.Reset()

				data := s.Bytes()
				d.ResetBytes(data)

				if err := e.Decode(&d); err != nil {
					lg.Error("decode", zap.Error(err))
				}
			}
			if err := s.Err(); err != nil {
				return errors.Wrap(err, "scan")
			}
			return nil
		})

		return g.Wait()
	})
}
