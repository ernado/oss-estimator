package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/estimate"
)

type AggregatedRepo struct {
	Name    string `json:"Name"`
	SLOC    int    `json:"SLOC"`
	PR      int    `json:"PR"`
	Commits int    `json:"Commits"`
}

type AggregatedOrg struct {
	Name    string                     `json:"Name"`
	SLOC    int                        `json:"SLOC"`
	PR      int                        `json:"PR"`
	Commits int                        `json:"Commits"`
	Repos   map[string]*AggregatedRepo `json:"Repos,omitempty"`
}

type Aggregated struct {
	Organizations map[string]*AggregatedOrg `json:"Organizations,omitempty"`
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			dir = "_work"
		)
		flag.StringVar(&dir, "dir", dir, "directory to store data")
		flag.Parse()

		out := &Aggregated{
			Organizations: map[string]*AggregatedOrg{},
		}

		orgs, err := os.ReadDir(dir)
		if err != nil {
			return errors.Wrap(err, "read orgs")
		}
		for _, org := range orgs {
			orgOut := &AggregatedOrg{
				Name:  org.Name(),
				Repos: map[string]*AggregatedRepo{},
			}
			if !org.IsDir() {
				continue
			}
			projects, err := os.ReadDir(filepath.Join(dir, org.Name()))
			if err != nil {
				return errors.Wrap(err, "read projects")
			}
			for _, project := range projects {
				if !project.IsDir() {
					continue
				}
				data, err := os.ReadFile(filepath.Join(dir, org.Name(), project.Name(), "cache.json"))
				if os.IsNotExist(err) {
					continue
				}
				if err != nil {
					return errors.Wrap(err, "read cache")
				}
				var e estimate.Entry
				if err := json.Unmarshal(data, &e); err != nil {
					return errors.Wrap(err, "unmarshal cache")
				}
				if orgOut.Name != e.Org {
					return errors.Errorf("org mismatch: %s != %s", orgOut.Name, e.Org)
				}
				repoOut := &AggregatedRepo{
					Name:    e.Repo,
					SLOC:    e.SLOC,
					PR:      e.PullRequests,
					Commits: e.Commits,
				}
				orgOut.SLOC += repoOut.SLOC
				orgOut.PR += repoOut.PR
				orgOut.Commits += repoOut.Commits
				orgOut.Repos[repoOut.Name] = repoOut
			}
			out.Organizations[orgOut.Name] = orgOut
		}

		e := json.NewEncoder(os.Stdout)
		e.SetIndent("", "  ")
		if err := e.Encode(out); err != nil {
			return errors.Wrap(err, "encode")
		}

		return nil
	})
}
