package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-faster/errors"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/gh"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		c := gh.Client(ctx)

		// Get all grpc libs.
		repos, _, err := c.Repositories.ListByOrg(ctx, "grpc", &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		})
		if err != nil {
			return errors.Wrap(err, "list repos")
		}
		var grpcLibs []string
		for _, repo := range repos {
			if strings.HasPrefix(repo.GetName(), "grpc-") {
				grpcLibs = append(grpcLibs, repo.GetName())
			}
		}

		fmt.Println(grpcLibs)

		return nil
	})
}
