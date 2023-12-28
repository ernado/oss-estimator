package cncf

import (
	_ "embed"
	"strings"
	"time"

	"github.com/go-faster/errors"
	yaml "github.com/go-faster/yamlx"
)

//go:embed projects.yaml
var projects []byte

type Project struct {
	Order     int       `yaml:"order"`
	Name      string    `yaml:"name"`
	Status    string    `yaml:"status"`
	StartDate time.Time `yaml:"start_date"`
	JoinDate  time.Time `yaml:"join_date"`
	Repo      string    `yaml:"main_repo"`
}

type Projects struct {
	Projects map[string]Project `yaml:"projects"`
}

func Load() (*Projects, error) {
	var v Projects
	if err := yaml.Unmarshal(projects, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

type Database struct {
	Organizations []string

	isCNCF map[string]struct{}
}

func (db *Database) Has(org string) bool {
	_, ok := db.isCNCF[strings.ToLower(org)]
	return ok
}

func New() (*Database, error) {
	var organizations []string
	data, err := Load()
	if err != nil {
		return nil, err
	}
	remap := map[string]string{
		"alibaba/sealer":     "sealerio/sealer",
		"SpectralOps/teller": "tellerops/teller",
		"G-Research/armada":  "armadaproject/armada",
	}
	isCNCF := make(map[string]struct{}, len(data.Projects))
	for _, p := range data.Projects {
		if p.Repo == "" || p.Status == "" || p.Status == "-" {
			continue
		}
		repo := p.Repo
		if v, ok := remap[repo]; ok {
			repo = v
		}
		idx := strings.Index(repo, "/")
		if idx == -1 {
			return nil, errors.Errorf("invalid repo for %q: %q", p.Name, repo)
		}
		org := strings.ToLower(repo[:idx])
		switch org {
		case "apache", "torvalds", "awslabs", "azure":
			continue
		}
		organizations = append(organizations, org)
		isCNCF[org] = struct{}{}
	}
	for _, org := range []string{
		"kubernetes-sigs",
		"kubeflow",
	} {
		organizations = append(organizations, org)
		isCNCF[org] = struct{}{}
	}
	return &Database{
		Organizations: organizations,
		isCNCF:        isCNCF,
	}, nil
}
