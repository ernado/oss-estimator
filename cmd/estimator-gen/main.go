package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/dustin/go-humanize"
	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	cncfdb "estimator/internal/cncf"
	"estimator/internal/estimate"
	"estimator/internal/lang"
)

//go:embed README.md.tmpl
var tpl string

type Stat struct {
	Name      string
	SLOC      int
	PR        int
	Commits   int
	Stars     int
	Org       string
	Language  string
	Languages map[string]int
}

func (s Stat) Effort() float64 {
	if s.SLOC == 0 {
		return 0
	}
	const linesPerMonth = 12_000
	v := float64(s.SLOC) / linesPerMonth
	const roundTo = 100
	return math.Round(v*roundTo) / roundTo
}

func (s Stat) Lang() string {
	if !lang.In(s.Language) {
		return "N/A"
	}
	return s.Language
}

const maxLen = 35

func (s Stat) Short() string {
	switch s.Name {
	case "CNCF", "K8s", "OTEL", "go-faster":
		// Markdown bold for aggregate.
		return "**" + s.Name + "**"
	}
	v := s.Name
	v = strings.TrimPrefix(v, "opentelemetry-")
	if len(v) > maxLen {
		return v[:maxLen-1] + "~"
	}
	return v
}

func (s Stat) Title() string {
	switch s.Name {
	case "CNCF", "K8s", "OTEL":
		// Markdown bold for aggregate.
		return "**" + s.Name + "**"
	}
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
	if s.Name == "CNCF" {
		return "https://www.cncf.io/"
	}
	if s.Name == "K8s" {
		return "https://kubernetes.io/"
	}
	if s.Name == "OTEL" {
		return "https://opentelemetry.io/docs/instrumentation/"
	}
	if s.Org == "" {
		return fmt.Sprintf("https://github.com/%s", s.Name)
	}
	return fmt.Sprintf("https://github.com/%s/%s/", s.Org, s.Name)
}

type Context struct {
	Orgs          []Stat
	Repos         []Stat
	CNCF          []Stat
	OTEL          []Stat
	Faster        []Stat
	LanguagesList []string
}

func (c Context) Languages() string {
	sort.Strings(c.LanguagesList)
	return strings.Join(c.LanguagesList, ", ")
}

func formatInt(x int) string {
	v := formatFloat(float64(x))
	return strings.TrimSuffix(v, ".0")
}

func formatFloat(num float64) string {
	v, u := humanize.ComputeSI(num)
	return fmt.Sprintf("%.1f%s", v, u)
}

func main() {
	// List of CNCF projects: https://www.cncf.io/projects/
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		t := template.New("template")
		t.Funcs(template.FuncMap{
			"formatInt":   formatInt,
			"formatFloat": formatFloat,
		})
		template.Must(t.Parse(tpl))

		db, err := cncfdb.New()
		if err != nil {
			return errors.Wrap(err, "open cncf db")
		}

		data, err := os.ReadFile(filepath.Join("_data", "aggregated.json"))
		if err != nil {
			return errors.Wrap(err, "read affiliations.json")
		}
		var ag estimate.Aggregated
		if err := json.Unmarshal(data, &ag); err != nil {
			return errors.Wrap(err, "unmarshal data")
		}

		// Aggregates.
		k8s := Stat{
			Name: "K8s",
		}
		cncf := Stat{
			Name: "CNCF",
		}
		otel := Stat{
			Name: "OTEL",
		}
		faster := Stat{
			Name: "go-faster",
		}
		c := Context{
			LanguagesList: lang.All(),
		}
		// Estimate go-faster projects.
		for _, org := range ag.Organizations {
			switch org.Name {
			case "ClickHouse", "ogen-go", "go-faster", "gotd", "pion":
			default:
				continue
			}
			sum := func(repo *estimate.AggregatedRepo) {
				if repo == nil {
					return
				}
				faster.PR += repo.PR
				faster.Commits += repo.Commits
				faster.Stars += repo.Stars
				faster.SLOC += repo.SLOC
				faster.Languages = estimate.Merge(faster.Languages, repo.Languages)
				v := Stat{
					Org:       org.Name,
					Name:      repo.Name,
					SLOC:      repo.SLOC,
					PR:        repo.PR,
					Commits:   repo.Commits,
					Stars:     repo.Stars,
					Language:  estimate.Max(repo.Languages),
					Languages: repo.Languages,
				}
				c.Faster = append(c.Faster, v)
			}
			switch org.Name {
			case "ClickHouse":
				sum(org.Repos["ch-go"])
				sum(org.Repos["ch-bench"])
			case "pion":
				sum(org.Repos["stun"])
			default:
				for _, repo := range org.Repos {
					if repo.SLOC == 0 {
						continue
					}
					switch repo.Name {
					case ".github", "porto", "yt-k8s-operator", "yt", "ytsaurus-ui", "minikube", "yamlx", "portoshim":
						continue
					}
					sum(repo)
				}
			}
		}
		for _, org := range ag.Organizations {
		Repo:
			for _, repo := range org.Repos {
				v := Stat{
					Org:       org.Name,
					Name:      repo.Name,
					SLOC:      repo.SLOC,
					PR:        repo.PR,
					Commits:   repo.Commits,
					Stars:     repo.Stars,
					Language:  estimate.Max(repo.Languages),
					Languages: repo.Languages,
				}
				c.Repos = append(c.Repos, v)
				if org.Name != "open-telemetry" {
					continue
				}
				if !strings.HasPrefix(v.Name, "opentelemetry-") {
					continue
				}
				for _, s := range []string{
					"vanityurls",
					"collection",
					"api",
					"collector",
					"docs",
					"sandbox",
					"profiling",
					"proto",
					"demo",
					"operator",
					"tools",
					"helm",
					"lambda",
					"specification",
					"sqlcommenter",
				} {
					if strings.Contains(v.Name, s) {
						continue Repo
					}
				}
				otel.PR += repo.PR
				otel.Commits += repo.Commits
				otel.Stars += repo.Stars
				otel.SLOC += repo.SLOC
				c.OTEL = append(c.OTEL, v)
			}
			v := Stat{
				Name:      org.Name,
				SLOC:      org.SLOC,
				PR:        org.PR,
				Commits:   org.Commits,
				Stars:     org.Stars,
				Language:  estimate.Max(org.Languages),
				Languages: org.Languages,
			}
			if db.Has(org.Name) {
				cncf.PR += org.PR
				cncf.Commits += org.Commits
				cncf.Stars += org.Stars
				cncf.SLOC += org.SLOC
				cncf.Languages = estimate.Merge(cncf.Languages, org.Languages)
				c.CNCF = append(c.CNCF, v)
			}
			c.Orgs = append(c.Orgs, v)
			switch org.Name {
			case "kubernetes", "kubernetes-sigs":
				k8s.PR += org.PR
				k8s.Commits += org.Commits
				k8s.Stars += org.Stars
				k8s.SLOC += org.SLOC
				k8s.Languages = estimate.Merge(k8s.Languages, org.Languages)
			}
		}

		// Add aggregates.
		k8s.Language = estimate.Max(k8s.Languages)
		cncf.Language = estimate.Max(cncf.Languages)
		c.Orgs = append(c.Orgs, k8s, cncf)
		c.OTEL = append(c.OTEL, otel)
		c.Faster = append(c.Faster, faster)

		comparator := func(s []Stat) func(i int, j int) bool {
			return func(i int, j int) bool {
				a, b := s[i], s[j]
				if a.SLOC == b.SLOC {
					return a.Name < b.Name
				}
				return s[i].SLOC > s[j].SLOC
			}
		}

		sort.SliceStable(c.Orgs, comparator(c.Orgs))
		sort.SliceStable(c.Repos, comparator(c.Repos))
		sort.SliceStable(c.CNCF, comparator(c.CNCF))
		sort.SliceStable(c.OTEL, comparator(c.OTEL))
		sort.SliceStable(c.Faster, comparator(c.Faster))

		var filteredRepos []Stat
		for _, repo := range c.Repos {
			if repo.SLOC < 25_000 {
				continue
			}
			filteredRepos = append(filteredRepos, repo)
		}
		c.Repos = filteredRepos

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
