package main

import (
	"context"
	"flag"
	"strings"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/estimate"
	"estimator/internal/gh"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			orgName  = "go-faster"
			repoName = "jx"
		)
		flag.Parse()
		if v := flag.Arg(0); v != "" {
			elems := strings.SplitN(v, "/", 2)
			if len(elems) != 2 {
				return errors.Errorf("invalid repo name: %q", v)
			}
			orgName, repoName = elems[0], elems[1]
		}

		c, err := estimate.New(gh.Client(ctx), "_work").Get(ctx, orgName, repoName)
		if err != nil {
			return errors.Wrap(err, "estimate")
		}
		c.Print()

		return nil
	})
}
