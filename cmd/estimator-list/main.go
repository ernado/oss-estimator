package main

import (
	"context"
	_ "embed"
	"flag"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-faster/errors"
	yaml "github.com/go-faster/yamlx"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/cncf"
	"estimator/internal/estimate"
	"estimator/internal/gh"
)

type Job struct {
	Org  string
	Repo string
}

type Include struct {
	Orgs  []string `yaml:"orgs"`
	Repos []string `yaml:"repos"`
}

type Config struct {
	Exclude []string `yaml:"exclude"`
	Include Include  `yaml:"include"`
}

//go:embed config.yaml
var configRaw []byte

type key [2]string

func toKey(repo string) key {
	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		panic("invalid repo: " + repo)
	}
	return key{parts[0], parts[1]}
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			dir         = path.Join("_work", "dataset")
			concurrency = 8
			force       bool
			pull        bool
		)
		flag.BoolVar(&force, "f", force, "force update (re-fetch orgs)")
		flag.BoolVar(&pull, "pull", pull, "pull repositories for updates")
		flag.StringVar(&dir, "dir", dir, "directory to store data")
		flag.IntVar(&concurrency, "j", concurrency, "number of concurrent jobs")
		flag.Parse()

		var cfg Config
		if err := yaml.Unmarshal(configRaw, &cfg); err != nil {
			return errors.Wrap(err, "unmarshal config")
		}

		var (
			c    = gh.Client()
			e    = estimate.New(c, dir).WithPull(pull).WithForce(force)
			jobs = make(chan Job, concurrency)
		)

		if err := gh.Check(); err != nil {
			return errors.Wrap(err, "check tokens")
		}

		cncfDB, err := cncf.New()
		if err != nil {
			return errors.Wrap(err, "load cncf")
		}

		var orgs []string
		for _, org := range cfg.Include.Orgs {
			if cncfDB.Has(org) {
				lg.Warn("Already in CNCF", zap.String("org", org))
				continue
			}
			orgs = append(orgs, org)
		}
		orgs = append(orgs, cncfDB.Organizations...)
		skip := make(map[key]struct{}, len(cfg.Exclude))
		for _, repo := range cfg.Exclude {
			skip[toKey(repo)] = struct{}{}
		}
		var keys []key
		for _, repo := range cfg.Include.Repos {
			keys = append(keys, toKey(repo))
		}

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			defer close(jobs)
			for _, org := range orgs {
				if stat, err := os.Stat(path.Join(dir, org)); err == nil && stat.IsDir() && !force {
					lg.Debug("Skipping org (already exists)",
						zap.String("org", org),
					)
					continue
				}
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
			for _, v := range keys {
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

						bo := backoff.NewExponentialBackOff()
						bo.MaxElapsedTime = time.Minute
						bo.MaxInterval = time.Second * 5
						stat, err := backoff.RetryNotifyWithData[*estimate.Entry](func() (*estimate.Entry, error) {
							if err := ctx.Err(); err != nil {
								return nil, backoff.Permanent(err)
							}
							return e.Get(ctx, j.Org, j.Repo)
						}, bo, func(err error, d time.Duration) {
							lg.Warn("Retry", zap.String("repo", j.Repo), zap.Duration("backoff", d), zap.Error(err))
						})
						if err != nil {
							lg.Error("Failed", zap.String("repo", j.Repo), zap.Error(err))
							continue
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
