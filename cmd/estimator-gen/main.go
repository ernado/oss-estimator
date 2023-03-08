package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/template"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/estimate"
)

//go:embed README.md.tmpl
var tpl string

//go:embed data.json
var data []byte

type Stat struct {
	Name    string
	SLOC    int
	PR      int
	Commits int
	Org     string
}

const maxLen = 35

func (s Stat) Title() string {
	v := s.Name
	if s.Org != "" {
		v = s.Org + "/" + s.Name
	}
	if len(v) > maxLen {
		return v[:maxLen-1] + "~"
	}
	return v
}

func (s Stat) URL() string {
	if s.Org == "" {
		return fmt.Sprintf("https://github.com/%s", s.Name)
	}
	return fmt.Sprintf("https://github.com/%s/%s/", s.Org, s.Name)
}

type Context struct {
	Orgs  []Stat
	Repos []Stat
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		t := template.Must(template.New("template").Parse(tpl))
		var v estimate.Aggregated
		if err := json.Unmarshal(data, &v); err != nil {
			return errors.Wrap(err, "unmarshal data")
		}
		var c Context
		for _, org := range v.Organizations {
			c.Orgs = append(c.Orgs, Stat{
				Name:    org.Name,
				SLOC:    org.SLOC,
				PR:      org.PR,
				Commits: org.Commits,
			})
			for _, repo := range org.Repos {
				c.Repos = append(c.Repos, Stat{
					Org:     org.Name,
					Name:    repo.Name,
					SLOC:    repo.SLOC,
					PR:      repo.PR,
					Commits: repo.Commits,
				})
			}
		}

		sort.Slice(c.Orgs, func(i, j int) bool {
			return c.Orgs[i].SLOC > c.Orgs[j].SLOC
		})
		sort.Slice(c.Repos, func(i, j int) bool {
			return c.Repos[i].SLOC > c.Repos[j].SLOC
		})

		f, err := os.Create("README.md")
		if err != nil {
			return errors.Wrap(err, "create file")
		}
		defer func() {
			_ = f.Close()
		}()
		return t.Execute(f, c)
	})
}
