package cncf

import (
	"os"
	"testing"

	yaml "github.com/go-faster/yamlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	data, err := Load()
	require.NoError(t, err)
	require.NotEmpty(t, data.Projects)
}

func TestHas(t *testing.T) {
	db, err := New()
	require.NoError(t, err)
	for _, s := range []string{
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
		"open-telemetry",
		"kubearmor",
	} {
		assert.Truef(t, db.Has(s), "missing %s", s)
	}
}

func TestIntersect(t *testing.T) {
	var toAdd []string

	db, err := New()
	require.NoError(t, err)
	require.NotEmpty(t, db.Organizations)

	for _, org := range []string{
		"ClickHouse",
		"grpc",
		"open-telemetry",
		"prometheus",
		"cilium",
		"kubernetes",
		"kubernetes-sigs",
		"istio",
		"etcd-io",
		"kata-containers",
		"siderolabs",
		"openebs",
		"m3db",
		"grafana",
		"VictoriaMetrics",
		"vectordotdev",
		"envoyproxy",
		"helm",
		"docker",
		"go-faster",
		"gotd",
		"ogen-go",
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
		"longhorn",
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
		"ydb-platform",
		"cockroachdb",
		"ytsaurus",
		"kubearmor",
		"netdata",
		"cncf",
	} {
		if db.Has(org) {
			continue
		}
		toAdd = append(toAdd, org)
	}

	e := yaml.NewEncoder(os.Stdout)
	e.SetIndent(2)

	type Include struct {
		Organizations []string `yaml:"organizations"`
	}
	require.NoError(t, e.Encode(Include{Organizations: toAdd}))
}
