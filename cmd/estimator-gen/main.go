package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/estimate"
	"estimator/internal/lang"
)

//go:embed README.md.tmpl
var tpl string

//go:embed data.json
var data []byte

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

func (s Stat) Lang() string {
	if !lang.In(s.Language) {
		return "N/A"
	}
	return s.Language
}

const maxLen = 35

func (s Stat) Title() string {
	if s.Name == "CNCF" || s.Name == "K8s" {
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
	if s.Org == "" {
		return fmt.Sprintf("https://github.com/%s", s.Name)
	}
	return fmt.Sprintf("https://github.com/%s/%s/", s.Org, s.Name)
}

type Context struct {
	Orgs          []Stat
	Repos         []Stat
	CNCF          []Stat
	LanguagesList []string
}

func (c Context) Languages() string {
	sort.Strings(c.LanguagesList)
	return strings.Join(c.LanguagesList, ", ")
}

// credit to https://github.com/DeyV/gotools/blob/master/numbers.go
func RoundPrec(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	if math.IsNaN(number) || math.IsInf(number, 0) {
		number = 0
	}

	var ret string
	var negative bool

	if number < 0 {
		number *= -1
		negative = true
	}

	d, fract := math.Modf(number)

	if decimals <= 0 {
		fract = 0
	} else {
		pow := math.Pow(10, float64(decimals))
		fract = RoundPrec(fract*pow, 0)
	}

	if thousandsSep == "" {
		ret = strconv.FormatFloat(d, 'f', 0, 64)
	} else if d >= 1 {
		var x float64
		for d >= 1 {
			d, x = math.Modf(d / 1000)
			x = x * 1000
			ret = strconv.FormatFloat(x, 'f', 0, 64) + ret
			if d >= 1 {
				ret = thousandsSep + ret
			}
		}
	} else {
		ret = "0"
	}

	fracts := strconv.FormatFloat(fract, 'f', 0, 64)

	// "0" pad left
	for i := len(fracts); i < decimals; i++ {
		fracts = "0" + fracts
	}

	ret += decPoint + fracts

	if negative {
		ret = "-" + ret
	}
	return ret
}

func RoundInt(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

func FormatNumber(input float64) string {
	x := RoundInt(input)
	xFormatted := NumberFormat(float64(x), 2, ".", ",")
	return xFormatted
}

func NearestThousandFormat(x int) string {
	num := float64(x)
	if math.Abs(num) < 999.5 {
		xNum := FormatNumber(num)
		xNumStr := xNum[:len(xNum)-3]
		return xNumStr
	}

	xNum := FormatNumber(num)
	// first, remove the .00 then convert to slice
	xNumStr := xNum[:len(xNum)-3]
	xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
	xNumSlice := strings.Fields(xNumCleaned)
	count := len(xNumSlice) - 2
	unit := [4]string{"K", "M", "B", "T"}
	xPart := unit[count]

	afterDecimal := ""
	if xNumSlice[1][0] != 0 {
		afterDecimal = "." + string(xNumSlice[1][0])
	}
	final := xNumSlice[0] + afterDecimal + xPart
	return final
}

func main() {
	// List of CNCF projects: https://www.cncf.io/projects/
	isCNCF := make(map[string]bool)
	for _, org := range []string{
		"istio",
		"etcd-io",
		"kubernetes",
		"kubernetes-sigs",
		"envoyproxy",
		"argoproj",
		"containerd",
		"coredns",
		"fluent",
		"goharbor",
		"jaegertracing",
		"linkerd",
		"open-policy-agent",
		"rook",
		"spiffe",
		"cri-o",
		"containernetworking",
		"cortexproject",
		"projectcontour",
		"nats-io",
		"notaryproject",
		"OpenObservability",
		"operator-framework",
		"thanos-io",
		"grpc",
		"longhorn",
		"fluent",
		"fluxcd",
		"fluent",
		"fluxcd",
		"helm",
		"prometheus",
		"theupdateframework",
		"vitessio",
		"backstage",
		"buildpacks",
		"cert-manager",
		"chaos-mesh",
		"cert-manager",
		"cloud-custodian",
		"cloudevents",
		"crossplane",
		"cubeFS",
		"dapr",
		"emissary-ingress",
		"in-toto",
		"kedacore",
		"keptn",
		"knative",
		"kubeedge",
		"kubevirt",
		"kyverno",
		"litmuschaos",
		"volcano-sh",
		"containerssh",
		"AthenZ",
		"carina-io",
		"k3s-io",
		"karmada-io",
	} {
		isCNCF[org] = true
	}

	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		t := template.New("template")
		t.Funcs(template.FuncMap{
			"formatNumber": NearestThousandFormat,
		})
		template.Must(t.Parse(tpl))

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

		c := Context{
			LanguagesList: lang.All(),
		}
		for _, org := range ag.Organizations {
			for _, repo := range org.Repos {
				c.Repos = append(c.Repos, Stat{
					Org:       org.Name,
					Name:      repo.Name,
					SLOC:      repo.SLOC,
					PR:        repo.PR,
					Commits:   repo.Commits,
					Stars:     repo.Stars,
					Language:  estimate.Max(repo.Languages),
					Languages: repo.Languages,
				})
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
			if isCNCF[org.Name] {
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

		var filteredRepos []Stat
		for _, repo := range c.Repos {
			if repo.SLOC < 4500 {
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
