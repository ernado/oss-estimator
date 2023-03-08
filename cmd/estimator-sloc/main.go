package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-faster/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"estimator/internal/app"
)

// scc entry per language.
type statEntry struct {
	// Language name.
	Name string `json:"Name"`
	// SLOC count.
	Code int `json:"Code"`
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

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		httpClient := oauth2.NewClient(ctx, ts)
		c := github.NewClient(httpClient)

		p := filepath.Join("_work", "git", orgName, repoName)
		root := osfs.New(p)
		storageRoot := osfs.New(filepath.Join(p, ".git"))
		storage := filesystem.NewStorage(storageRoot, cache.NewObjectLRUDefault())

		// Try to open first, so we don't need to call GitHub API.
		// Fast path.
		gitRepo, err := git.Open(storage, root)
		if err != nil {
			// Slow path.
			repo, _, err := c.Repositories.Get(ctx, orgName, repoName)
			if err != nil {
				return errors.Wrap(err, "get repository")
			}

			ref := plumbing.NewBranchReferenceName(repo.GetDefaultBranch())
			gitRepo, err = git.CloneContext(ctx, storage, root, &git.CloneOptions{
				Depth: 1,
				URL:   repo.GetCloneURL(),
				Auth: &http.BasicAuth{
					Username: "read",
					Password: os.Getenv("GITHUB_TOKEN"),
				},
				ReferenceName: ref,
				SingleBranch:  true,
			})
			if err != nil {
				return errors.Wrap(err, "clone")
			}
		}

		head, err := gitRepo.Head()
		if err != nil {
			return errors.Wrap(err, "head")
		}
		fmt.Println("head", head)

		args := []string{
			"--no-complexity",
			"--no-cocomo",
			"--no-min-gen",
			"--sort", "code",
			"-x", "yaml,yml,md",
			"--format", "json",
		}

		// Ignore common vendor directories.
		for _, v := range []string{
			"vendor",
			"include",
			"third_party",
			"3rdparty",
		} {
			args = append(args, "--exclude-dir", v)
		}

		// We can't use scc as library because of global state.
		out, outErr := new(bytes.Buffer), new(bytes.Buffer)
		cmd := exec.CommandContext(ctx, "scc", args...)
		cmd.Dir = p
		cmd.Stdout = out
		cmd.Stderr = outErr
		if err := cmd.Run(); err != nil {
			if outErr.Len() > 0 {
				return errors.Wrapf(err, "run scc: %s", outErr)
			}
			return errors.Wrap(err, "run scc")
		}

		d := json.NewDecoder(out)
		var stats []statEntry
		if err := d.Decode(&stats); err != nil {
			return errors.Wrap(err, "decode scc output")
		}

		for _, s := range stats {
			fmt.Println(s.Name, s.Code)
		}

		return nil
	})
}
