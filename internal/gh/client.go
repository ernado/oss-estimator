package gh

import (
	"context"
	"os"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func Client(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(ctx, ts)
	c := github.NewClient(httpClient)
	return c

}
