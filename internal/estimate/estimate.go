package estimate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/google/go-github/v50/github"

	"estimator/internal/lang"
)

// StatEntry is scc entry per language.
type StatEntry struct {
	// Language name.
	Name string `json:"Name"`
	// SLOC count.
	Code int `json:"Code"`
}

// Entry is a cache entry.
type Entry struct {
	Org          string      `json:"Org"`
	Repo         string      `json:"Repo"`
	Code         []StatEntry `json:"Code"`
	Commits      int         `json:"Commits,omitempty"`
	PullRequests int         `json:"PullRequests,omitempty"`
	SLOC         int         `json:"SLOC,omitempty"`
	Head         string      `json:"HEAD,omitempty"`
	RepoID       int64       `json:"RepoID,omitempty"`
	OrgID        int64       `json:"OrgID,omitempty"`
	Stars        int         `json:"Stars,omitempty"`
}

type AggregatedRepo struct {
	ID        int64          `json:"ID"`
	OrgID     int64          `json:"OrgID"`
	Name      string         `json:"Name"`
	SLOC      int            `json:"SLOC"`
	PR        int            `json:"PR"`
	Commits   int            `json:"Commits"`
	Stars     int            `json:"Stars"`
	Languages map[string]int `json:"Languages"`
}

type AggregatedOrg struct {
	ID        int64                      `json:"ID"`
	Name      string                     `json:"Name"`
	SLOC      int                        `json:"SLOC"`
	PR        int                        `json:"PR"`
	Commits   int                        `json:"Commits"`
	Stars     int                        `json:"Stars"`
	Languages map[string]int             `json:"Languages"`
	Repos     map[string]*AggregatedRepo `json:"Repos,omitempty"`
}

func Merge(dst, src map[string]int) map[string]int {
	if dst == nil {
		dst = make(map[string]int)
	}
	for k, v := range src {
		dst[k] += v
	}
	return dst
}

func Max(s map[string]int) string {
	type kv struct {
		Name string
		Code int
	}
	var elements []kv
	for k := range s {
		if !lang.In(k) {
			continue
		}
		elements = append(elements, kv{Name: k, Code: s[k]})
	}
	if len(elements) == 0 {
		return ""
	}
	slices.SortStableFunc(elements, func(a, b kv) int {
		if a.Code == b.Code {
			return strings.Compare(a.Name, b.Name)
		}
		return b.Code - a.Code
	})
	return elements[0].Name
}

type Aggregated struct {
	Organizations map[string]*AggregatedOrg `json:"Organizations,omitempty"`
}

func (e Entry) Print() {
	var b strings.Builder

	fullName := fmt.Sprintf("%s/%s", e.Org, e.Repo)
	b.WriteString(fmt.Sprintf("%-50s", fullName))

	b.WriteString(fmt.Sprintf(" SLOC=%-15d", e.SLOC))
	b.WriteString(fmt.Sprintf(" Commits=%-8d", e.Commits))
	b.WriteString(fmt.Sprintf(" PR=%-8d", e.PullRequests))
	b.WriteString(fmt.Sprintf(" Stars=%-8d", e.Stars))

	fmt.Println(b.String())
}

type Client struct {
	gh    *github.Client
	dir   string
	pull  bool
	force bool
}

func (c *Client) WithPull(v bool) *Client {
	c.pull = v
	return c
}

func (c *Client) WithForce(v bool) *Client {
	c.force = v
	return c
}

func New(ghClient *github.Client, dir string) *Client {
	c := &Client{
		gh:  ghClient,
		dir: dir,
	}

	return c
}

func (c *Client) Get(ctx context.Context, orgName, repoName string) (*Entry, error) {
	p := filepath.Join(c.dir, orgName, repoName)
	cacheEntryPath := filepath.Join(p, "cache.json")
	if data, err := os.ReadFile(cacheEntryPath); err == nil && !c.pull && !c.force {
		var ce Entry
		if err := json.Unmarshal(data, &ce); err != nil {
			return nil, errors.Wrap(err, "unmarshal cache entry")
		}
		return &ce, nil
	}

	gitRoot := filepath.Join(p, "git")
	root := osfs.New(gitRoot)
	storageRoot := osfs.New(filepath.Join(gitRoot, ".git"))
	storage := filesystem.NewStorage(storageRoot, cache.NewObjectLRUDefault())

	// Try to open first, so we don't need to call GitHub API.
	// Fast path.
	gitRepo, err := git.Open(storage, root)

	if err == nil {
		_, err = gitRepo.Head()
	}

	var (
		repoID int64
		orgID  int64
		stars  int
	)

	if err != nil {
		// Slow path, cloned repo doesn't exist.
		//
		// Fetching default branch and cloning.
		repo, _, err := c.gh.Repositories.Get(ctx, orgName, repoName)
		if err != nil {
			return nil, errors.Wrap(err, "get repository")
		}

		repoID = repo.GetID()
		stars = repo.GetStargazersCount()
		orgID = repo.GetOwner().GetID()
		u, err := url.Parse(repo.GetCloneURL())
		if err != nil {
			return nil, errors.Wrap(err, "parse clone URL")
		}
		// Fix partial clone.
		_ = os.RemoveAll(gitRoot)

		// git is significantly faster than go-git on big repos for cloning.
	fetch:
		cmd := exec.CommandContext(ctx, "git", "clone", "--depth=1", u.String(), gitRoot)
		out, outErr := new(bytes.Buffer), new(bytes.Buffer)
		cmd.Stdout = out
		cmd.Stderr = outErr

		if err := cmd.Run(); err != nil {
			if strings.Contains(outErr.String(), "fatal: Resource not accessible by personal access token") {
				// Try again without auth.
				u.User = nil
				goto fetch
			}
			if outErr.Len() > 0 {
				return nil, errors.Wrapf(err, "run git: %s", outErr)
			}
			return nil, errors.Wrap(err, "run git")
		}

		gitRepo, err = git.Open(storage, root)
		if err != nil {
			return nil, errors.Wrap(err, "open git repo after clone")
		}
	} else if c.pull {
		cmd := exec.CommandContext(ctx, "git", "pull")
		out, outErr := new(bytes.Buffer), new(bytes.Buffer)
		cmd.Stdout = out
		cmd.Stderr = outErr
		cmd.Dir = gitRoot

		gitRepo, err = git.Open(storage, root)
		if err != nil {
			return nil, errors.Wrap(err, "open git repo after clone")
		}
	}

	// Initialize arguments for scc.
	exclude := []string{
		"yaml", "yml", "md", "json",
	}
	args := []string{
		"--no-complexity",
		"--no-cocomo",
		"--no-min-gen",
		"--sort", "code",
		"-x", strings.Join(exclude, ","),
		"--format", "json",
	}

	ignoreDirs := []string{
		"vendor",
		"include",
		"third_party",
		"3rdparty",
		"contrib",
		"c-deps",
		"deps",
		"Godeps",
		"_workspace",

		".git",
		".yarn",
		".vendor",
		".github",
	}
	switch repoName {
	case "statshouse":
		ignoreDirs = append(ignoreDirs, "sqlite")
	case "fluent-bit":
		ignoreDirs = append(ignoreDirs, "lib")
	}

	// Ignore common vendor directories.
	for _, v := range ignoreDirs {
		args = append(args, "--exclude-dir", v)
	}

	// We can't use scc as library because of global state.
	out, outErr := new(bytes.Buffer), new(bytes.Buffer)
	cmd := exec.CommandContext(ctx, "scc", args...)
	cmd.Dir = gitRoot
	cmd.Stdout = out
	cmd.Stderr = outErr
	if err := cmd.Run(); err != nil {
		if outErr.Len() > 0 {
			return nil, errors.Wrapf(err, "run scc: %s", outErr)
		}
		return nil, errors.Wrap(err, "run scc")
	}

	d := json.NewDecoder(out)
	var stats []StatEntry
	if err := d.Decode(&stats); err != nil {
		return nil, errors.Wrap(err, "decode scc output")
	}
	var total int
	for _, s := range stats {
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
		_, res, err := c.gh.Repositories.ListCommits(ctx, orgName, repoName, &github.CommitsListOptions{
			ListOptions: list,
		})
		if err != nil {
			return nil, errors.Wrap(err, "list commits")
		}
		commits = res.LastPage
	}
	{
		_, res, err := c.gh.PullRequests.List(ctx, orgName, repoName, &github.PullRequestListOptions{
			State:       "all",
			ListOptions: list,
		})
		if err != nil {
			return nil, errors.Wrap(err, "list pull requests")
		}
		pullRequests = res.LastPage
	}

	head, err := gitRepo.Head()
	if err != nil {
		return nil, errors.Wrap(err, "get head")
	}

	if repoID == 0 {
		repo, _, err := c.gh.Repositories.Get(ctx, orgName, repoName)
		if err != nil {
			return nil, errors.Wrap(err, "get repository")
		}
		repoID, stars, orgID = repo.GetID(), repo.GetStargazersCount(), repo.GetOwner().GetID()
	}

	ce := Entry{
		OrgID:        orgID,
		PullRequests: pullRequests,
		Commits:      commits,
		Code:         stats,
		SLOC:         total,
		Head:         head.Hash().String(),
		Org:          orgName,
		Repo:         repoName,
		RepoID:       repoID,
		Stars:        stars,
	}
	cOut := new(bytes.Buffer)
	e := json.NewEncoder(cOut)
	e.SetIndent("", "  ")
	if err := e.Encode(ce); err != nil {
		return nil, errors.Wrap(err, "encode cache entry")
	}
	if err := os.WriteFile(cacheEntryPath, cOut.Bytes(), 0o644); err != nil {
		return nil, errors.Wrap(err, "write cache entry")
	}

	return &ce, nil

}
