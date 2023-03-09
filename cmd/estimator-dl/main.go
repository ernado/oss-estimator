package main

import (
	"context"
	"flag"
	"path/filepath"
	"time"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/archive"
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

		uc, err := archive.NewUserCache(arg.Dir, 250_000)
		if err != nil {
			return errors.Wrap(err, "cache")
		}
		defer func() {
			lg.Info("Closing")
			_ = uc.Close()
		}()

		d := archive.NewDownloader(lg, uc, arg.Dir)
		if err := d.Download(ctx, time.Now().AddDate(0, 0, -2)); err != nil {
			return errors.Wrap(err, "download")
		}
		if err := uc.Close(); err != nil {
			return errors.Wrap(err, "close")
		}

		return nil
	})
}
