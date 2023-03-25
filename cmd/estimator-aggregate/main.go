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
	"estimator/internal/cncf"
	"estimator/internal/estimate"
)

func write(path string, v interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "create")
	}
	defer func() { _ = f.Close() }()
	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	if err := e.Encode(v); err != nil {
		return errors.Wrap(err, "encode")
	}
	if err := f.Sync(); err != nil {
		return errors.Wrap(err, "sync")
	}
	if err := f.Close(); err != nil {
		return errors.Wrap(err, "sync")
	}
	return nil
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg struct {
			Dir    string
			OutDir string
		}
		flag.StringVar(&arg.Dir, "dir", filepath.Join("_work", "dataset"), "directory to read data from")
		flag.StringVar(&arg.OutDir, "o", "_data", "directory to read data from")
		flag.Parse()

		out := &estimate.Aggregated{
			Organizations: map[string]*estimate.AggregatedOrg{},
		}
		orgs, err := os.ReadDir(arg.Dir)
		if err != nil {
			return errors.Wrap(err, "read orgs")
		}

		db, err := cncf.New()
		if err != nil {
			return errors.Wrap(err, "open cncf db")
		}

		var cncfRepos []int64
		for _, org := range orgs {
			orgOut := &estimate.AggregatedOrg{
				Name:      org.Name(),
				Repos:     map[string]*estimate.AggregatedRepo{},
				Languages: map[string]int{},
			}
			if !org.IsDir() {
				continue
			}
			projects, err := os.ReadDir(filepath.Join(arg.Dir, org.Name()))
			if err != nil {
				return errors.Wrap(err, "read projects")
			}
			for _, project := range projects {
				if !project.IsDir() {
					continue
				}
				data, err := os.ReadFile(filepath.Join(arg.Dir, org.Name(), project.Name(), "cache.json"))
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
				if db.Has(e.Org) {
					cncfRepos = append(cncfRepos, e.RepoID)
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

		if err := write(filepath.Join(arg.OutDir, "aggregated.json"), out); err != nil {
			return errors.Wrap(err, "write")
		}
		if err := write(filepath.Join(arg.OutDir, "cncf-repos.json"), cncfRepos); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	})
}
