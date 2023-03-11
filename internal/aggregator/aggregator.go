package aggregator

func IsCNCF(orgName string) bool {
	_, ok := _cncf[orgName]
	return ok
}

var _cncf = func() map[string]struct{} {
	out := make(map[string]struct{})
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
	} {
		out[s] = struct{}{}
	}

	return out
}()
