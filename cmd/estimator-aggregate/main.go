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

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var (
			dir = filepath.Join("_work", "dataset")
		)
		flag.StringVar(&dir, "dir", dir, "directory to store data")
		flag.Parse()

		out := &estimate.Aggregated{
			Organizations: map[string]*estimate.AggregatedOrg{},
		}

		orgs, err := os.ReadDir(dir)
		if err != nil {
			return errors.Wrap(err, "read orgs")
		}
		for _, org := range orgs {
			orgOut := &estimate.AggregatedOrg{
				Name:      org.Name(),
				Repos:     map[string]*estimate.AggregatedRepo{},
				Languages: map[string]int{},
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

				repoOut := &estimate.AggregatedRepo{
					ID:        e.RepoID,
					Name:      e.Repo,
					SLOC:      e.SLOC,
					PR:        e.PullRequests,
					Stars:     e.Stars,
					Commits:   e.Commits,
					OrgID:     e.OrgID,
					Languages: map[string]int{},
				}
				for _, v := range e.Code {
					repoOut.Languages[v.Name] += v.Code
				}
				orgOut.ID = e.OrgID
				orgOut.SLOC += repoOut.SLOC
				orgOut.PR += repoOut.PR
				orgOut.Commits += repoOut.Commits
				orgOut.Stars += repoOut.Stars
				orgOut.Repos[repoOut.Name] = repoOut
				for k, v := range repoOut.Languages {
					orgOut.Languages[k] += v
				}
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
