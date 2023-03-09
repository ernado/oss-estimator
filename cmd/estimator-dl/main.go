package main

import (
	"context"
	"flag"
	"path/filepath"
	"time"

	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/entry"
)

func main() {
	// clickhouse local --structure "event Enum8('WatchEvent'=1, 'PushEvent'=2, 'IssuesEvent'=3, 'PullRequestEvent'=4), repo Int64, actor Int64, time DateTime" --input-format Native --interactive --file events.ch.native.zst
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg = struct {
			Jobs int
			Dir  string
		}{
			Jobs: 8,
			Dir:  filepath.Join("_work"),
		}
		flag.IntVar(&arg.Jobs, "j", arg.Jobs, "concurrent jobs")
		flag.StringVar(&arg.Dir, "dir", arg.Dir, "directory to store data")
		flag.Parse()

		if arg.Jobs <= 0 {
			arg.Jobs = 1
		}

		uc, err := archive.NewUserCache(filepath.Join(arg.Dir, "users.bbolt"), 250_000)
		if err != nil {
			return errors.Wrap(err, "cache")
		}
		defer func() {
			lg.Info("Closing")
			_ = uc.Close()
		}()
		var (
			d     = archive.NewDownloader(lg, uc, arg.Dir)
			start = entry.Start()
			end   = time.Now().AddDate(0, 0, -2)
			jobs  = make(chan time.Time, arg.Jobs)
		)
		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			defer close(jobs)
			for t := start; t.Before(end); t = t.Add(entry.Delta) {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case jobs <- t:
				}
			}
			return nil
		})
		for i := 0; i < arg.Jobs; i++ {
			g.Go(func() error {
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case t, ok := <-jobs:
						if !ok {
							return nil
						}
						if err := d.Download(ctx, t); err != nil {
							return errors.Wrap(err, "download")
						}
					}
				}
			})
		}
		if err := g.Wait(); err != nil {
			return errors.Wrap(err, "wait")
		}
		if err := uc.Close(); err != nil {
			return errors.Wrap(err, "close")
		}

		return nil
	})
}
