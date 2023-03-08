package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/estimate"
	"estimator/internal/gh"
)

// scc entry per language.
type statEntry struct {
	// Language name.
	Name string `json:"Name"`
	// SLOC count.
	Code int `json:"Code"`
}

type cacheEntry struct {
	Code         []statEntry `json:"Code"`
	Commits      int         `json:"Commits,omitempty"`
	PullRequests int         `json:"PullRequests,omitempty"`
	SLOC         int         `json:"SLOC,omitempty"`
}

func (c cacheEntry) Print() {
	fmt.Println("Total:")
	fmt.Println("", "SLOC", c.SLOC)
	fmt.Println("", "commits", c.Commits)
	fmt.Println("", "pull requests", c.PullRequests)
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			orgName  = "go-faster"
			repoName = "jx"
		)
		flag.StringVar(&orgName, "org", orgName, "GitHub organization name")
		flag.StringVar(&repoName, "repo", repoName, "GitHub repository name")
		flag.Parse()

		c, err := estimate.New(gh.Client(ctx), "_work").Get(ctx, orgName, repoName)
		if err != nil {
			return errors.Wrap(err, "estimate")
		}
		c.Print()

		return nil
	})
}
