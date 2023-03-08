package main

import (
	"context"
	"flag"
	"path"
	"strings"
	"sync"

	"github.com/go-faster/errors"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/estimate"
	"estimator/internal/gh"
)

type Job struct {
	Org  string
	Repo string
}

type key [2]string

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			dir         = path.Join("_work", "dataset")
			concurrency = 8
		)
		flag.StringVar(&dir, "dir", dir, "directory to store data")
		flag.IntVar(&concurrency, "j", concurrency, "number of concurrent jobs")
		flag.Parse()

		var (
			c    = gh.Client(ctx)
			e    = estimate.New(c, dir)
			jobs = make(chan Job, concurrency)
		)

		skip := map[key]struct{}{
			{"ClickHouse", "clickhouse.github.io"}: {},
			{"ClickHouse", "llvm"}:                 {},
		}

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			defer close(jobs)
			for _, org := range []string{
				"ClickHouse",
				"grpc",
				"open-telemetry",
				"prometheus",
				"cilium",
				"kubernetes",
				"kubernetes-sigs",
				"istio",
				"etcd-io",
				"kata-containers",
				"siderolabs",
				"openebs",
				"m3db",
				"grafana",
				"VictoriaMetrics",
				"vectordotdev",
				"envoyproxy",
				"helm",
				"docker",
				"go-faster",
				"gotd",
				"ogen-go",
			} {
				repos, _, err := c.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
					ListOptions: github.ListOptions{
						PerPage: 500,
					},
				})
				if err != nil {
					return errors.Wrap(err, "list repos")
				}
				for _, repo := range repos {
					if repo.GetSize() == 0 || repo.GetFork() {
						continue
					}
					if _, ok := skip[key{org, repo.GetName()}]; ok {
						continue
					}
					if strings.Contains(repo.GetName(), "github.io") {
						continue
					}
					select {
					case <-ctx.Done():
						return ctx.Err()
					case jobs <- Job{Org: org, Repo: repo.GetName()}:
					}
				}
			}
			for _, v := range []key{
				{"VKCOM", "statshouse"},
				{"VKCOM", "VKUI"},
				{"pixie-io", "pixie"},
				{"falcosecurity", "falco"},
				{"apache", "mesos"},
				{"apache", "aurora"},
				{"uber", "peloton"},
				{"Netflix", "titus-executor"},
				{"Netflix", "titus-control-plane"},
				{"vitalif", "vitastor"},
				{"LINBIT", "linstor-server"},
				{"uber", "kraken"},
				{"containers", "podman"},
			} {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case jobs <- Job{Org: v[0], Repo: v[1]}:
				}
			}
			return nil
		})
		var mux sync.Mutex
		for i := 0; i < concurrency; i++ {
			g.Go(func() error {
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case j, ok := <-jobs:
						if !ok {
							return nil
						}
						stat, err := e.Get(ctx, j.Org, j.Repo)
						if err != nil {
							return errors.Wrap(err, "get repo")
						}
						mux.Lock()
						stat.Print()
						mux.Unlock()
					}
				}
			})
		}

		return g.Wait()
	})
}
