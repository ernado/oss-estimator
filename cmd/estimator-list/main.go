package main

import (
	"context"
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

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			c           = gh.Client(ctx)
			e           = estimate.New(c, "_work")
			concurrency = 8
			jobs        = make(chan Job, concurrency)
		)

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			defer close(jobs)
			for _, org := range []string{
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
					select {
					case <-ctx.Done():
						return ctx.Err()
					case jobs <- Job{Org: org, Repo: repo.GetName()}:
					}
				}
			}
			for _, v := range [][2]string{
				{"ogen-go", "ogen"},
				{"gotd", "td"},
				{"kubernetes", "kubernetes"},
				{"cilium", "cilium"},
				{"cilium", "tetragon"},
				{"cilium", "hubble"},
				{"cilium", "hubble-ui"},
				{"cilium", "ebpf"},
				{"cilium", "cilium-cli"},
				{"cilium", "pwru"},
				{"cilium", "hubble-otel"},
				{"kata-containers", "kata-containers"},
				{"grafana", "grafana"},
				{"grafana", "loki"},
				{"grafana", "mimir"},
				{"grafana", "oncall"},
				{"grafana", "agent"},
				{"grafana", "k6"},
				{"VKCOM", "statshouse"},
				{"VKCOM", "VKUI"},
				{"VictoriaMetrics", "VictoriaMetrics"},
				{"VictoriaMetrics", "metrics"},
				{"VictoriaMetrics", "grafana-datasource"},
				{"VictoriaMetrics", "operator"},
				{"pixie-io", "pixie"},
				{"siderolabs", "talos"},
				{"falcosecurity", "falco"},
				{"ClickHouse", "ClickHouse"},
				{"vectordotdev", "vector"},
				{"prometheus", "node_exporter"},
				{"prometheus", "common"},
				{"prometheus", "exporter-toolkit"},
				{"prometheus", "prometheus"},
				{"prometheus", "pushgateway"},
				{"envoyproxy", "envoy"},
				{"istio", "istio"},
				{"istio", "ztunnel"},
				{"istio", "proxy"},
				{"etcd-io", "etcd"},
				{"apache", "mesos"},
				{"apache", "aurora"},
				{"uber", "peloton"},
				{"Netflix", "titus-executor"},
				{"Netflix", "titus-control-plane"},
				{"helm", "helm"},
				{"helm", "chartmuseum"},
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
