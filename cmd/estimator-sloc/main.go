package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"estimator/internal/app"
	"estimator/internal/lang"
)

// scc entry per language.
type statEntry struct {
	// Language name.
	Name string `json:"Name"`
	// SLOC count.
	Code int `json:"Code"`
}

type cacheEntry struct {
	Code         []statEntry `json:"Code"`
	Commits      int         `json:"Commits,omitempty"`
	PullRequests int         `json:"PullRequests,omitempty"`
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

		p := filepath.Join("_work", "git", orgName, repoName)
		root := osfs.New(p)
		storageRoot := osfs.New(filepath.Join(p, ".git"))
		storage := filesystem.NewStorage(storageRoot, cache.NewObjectLRUDefault())

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		httpClient := oauth2.NewClient(ctx, ts)
		c := github.NewClient(httpClient)

		// Try to open first, so we don't need to call GitHub API.
		// Fast path.
		gitRepo, err := git.Open(storage, root)
		if err != nil {
			// Slow path, cloned repo doesn't exist.
			//
			// Fetching default branch and cloning.

			repo, _, err := c.Repositories.Get(ctx, orgName, repoName)
			if err != nil {
				return errors.Wrap(err, "get repository")
			}

			u, err := url.Parse(repo.GetCloneURL())
			if err != nil {
				return errors.Wrap(err, "parse clone URL")
			}
			u.User = url.UserPassword("git", os.Getenv("GITHUB_TOKEN"))

			// Fix partial clone.
			if err := os.RemoveAll(p); err != nil {
				lg.Warn("RemoveAll failed", zap.Error(err))
			}

			// git is significantly faster than go-git on big repos for cloning.
			cmd := exec.CommandContext(ctx, "git", "clone", "--depth=1", u.String(), p)
			out, outErr := new(bytes.Buffer), new(bytes.Buffer)
			cmd.Stdout = out
			cmd.Stderr = outErr

			if err := cmd.Run(); err != nil {
				if outErr.Len() > 0 {
					return errors.Wrapf(err, "run git: %s", outErr)
				}
				return errors.Wrap(err, "run git")
			}

			gitRepo, err = git.Open(storage, root)
			if err != nil {
				return errors.Wrap(err, "open git repo after clone")
			}
		}

		head, err := gitRepo.Head()
		if err != nil {
			return errors.Wrap(err, "head")
		}

		fmt.Println("git head:", head)

		// Initialize arguments for scc.
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

		fmt.Println("languages:")
		var total int
		for _, s := range stats {
			fmt.Println("", strings.ToLower(s.Name), s.Code)
			if lang.In(s.Name) {
				total += s.Code
			}
		}

		var (
			commits      int
			pullRequests int
		)
		// Last page number for per-page: 1 will be total entities number.
		list := github.ListOptions{
			PerPage: 1,
		}
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

		fmt.Println("Languages that are counted:", lang.All())

		fmt.Println("Total:")
		fmt.Println("", "SLOC", total)
		fmt.Println("", "commits", commits)
		fmt.Println("", "pull requests", pullRequests)

		return nil
	})
}
