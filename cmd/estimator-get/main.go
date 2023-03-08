package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-faster/errors"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"estimator/internal/app"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		httpClient := oauth2.NewClient(ctx, ts)
		c := github.NewClient(httpClient)

		// Last page number for per-page: 1 will be total entities number.
		list := github.ListOptions{
			PerPage: 1,
		}

		var (
			commits      int
			pullRequests int
		)
		{
			_, res, err := c.Repositories.ListCommits(ctx, "kubernetes", "kubernetes", &github.CommitsListOptions{
				ListOptions: list,
			})
			if err != nil {
				return errors.Wrap(err, "list commits")
			}
			commits = res.LastPage
		}
		{
			_, res, err := c.PullRequests.List(ctx, "kubernetes", "kubernetes", &github.PullRequestListOptions{
				State:       "all",
				ListOptions: list,
			})
			if err != nil {
				return errors.Wrap(err, "list pull requests")
			}
			pullRequests = res.LastPage
		}

		fmt.Println("commits", commits, "pr", pullRequests)

		return nil
	})
}
