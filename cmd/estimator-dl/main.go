package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/go-faster/jx"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/entry"
)

func run(ctx context.Context, lg *zap.Logger) error {
	return nil
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		u := archive.GetURL(time.Now().AddDate(0, 0, -2))

		g, ctx := errgroup.WithContext(ctx)
		cmd := exec.CommandContext(ctx, "bash", "-c",
			fmt.Sprintf("wget -O - %s | zstd -d", u),
		)
		cmd.Stderr = os.Stderr
		out, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		g.Go(func() error {
			return cmd.Run()
		})
		g.Go(func() error {
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
				return err
			}
			if _, err := io.Copy(io.Discard, out); err != nil {
				return err
			}
			return nil
		})

		return g.Wait()
	})
}
