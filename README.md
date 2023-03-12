# oss-estimator

Set of tools to get info about GitHub project that is needed to estimate the cost of running it.

## Tools

| Tool                                                     | Description                                                                                               |
|----------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [estimator-dl](./cmd/estimator-dl/main.go)               | Download from [gharchive.org](https://gharchive.org) and save to zstd-compressed clickhouse native format |
| [estimator-sloc](./cmd/estimator-sloc/main.go)           | SLOC count using [scc](https://github.com/boyter/scc/) (should be in $PATH)                               |
| [estimator-list](./cmd/estimator-list/main.go)           | Concurrently fetch popular oss repos and stat for them                                                    |
| [estimator-aggregate](./cmd/estimator-aggregate/main.go) | Aggregate stats                                                                                           |
| [estimator-gen](./cmd/estimator-gen/main.go)             | Generate README.md                                                                                        |


## Data used
- [gharchive.org](https://gharchive.org), GitHub events archive from 2014 until now
- Direct clone of GitHub repos (~50GB of data)
- [cncf/gitdm](https://github.com/cncf/gitdm/tree/master), CNCF affiliation of developers

## Example

```console
$ go run ./cmd/estimator-sloc open-telemetry/opentelemetry-ebpf
+-------------------+
| open-telemetry/op |
| entelemetry-ebpf  |
+---------+---------+
| SLOC    | 4295    |
| Commits | 267     |
| PR      | 110     |
| HEAD    | df4e6e6 |
+---------+---------+
```

```console
$ du -hs _work/
31G	_work
```

## Stats

Languages: BASH, Bazel, C, C Header, C#, C++, Cython, Dart, Erlang, Go, Go Template, Haskell, Java, JavaScript, Kotlin, Lua, Objective C, Objective C++, PHP, Perl, Protocol Buffers, Python, Ruby, Rust, Scala, Swift, TypeScript, Zig

### Effort

Effort is **rough estimate** how much developer-years it would take to write this project from scratch in proprietary,
enterprise environment (i.e. in a big company).

Current model is very simple: we estimate that average developer writes 12k SLOC per year.

TODO:
- Take into account COCOMO (not suitable for FOSS?)
- Estimate on commits and PRs (can vary a lot, also )
    - Average developer makes 3 commits per year
    - Exclude robot commits and PRs, like dependabot
- Research papers on FOSS effort estimation if proposed model is applicable

### Aggregates

Some organizations are highlighted in bold, they represent additional aggregated statistics.

#### CNCF

Organizations that are CNCF projects (graduated, incubating, sandbox).

#### K8s

Repositories in kubernetes and kuberentes-sigs organizations.

### Organizations
| Organization | SLOC | Commits | PRs | Stars | Language | Effort |
|--------------|------|---------|-----|-------|----------|--------|
| [**CNCF**](https://www.cncf.io/)  | 22.4M | 1.4M | 824.5K | 1.1M | Go | 1.8K |
| [torvalds](https://github.com/torvalds)  | 19.1M | 1.1M | 761 | 147.5K | C | 1.5K |
| [**K8s**](https://kubernetes.io/)  | 5.4M | 498.5K | 261.5K | 341.3K | Go | 453 |
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1K | 190.5K | 243.1K | Go | 286 |
| [tensorflow](https://github.com/tensorflow)  | 3.3M | 144.4K | 22.7K | 171.8K | C++ | 279 |
| [ydb-platform](https://github.com/ydb-platform)  | 3.1M | 18.5K | 1.5K | 3.1K | C++ | 261 |
| [elastic](https://github.com/elastic)  | 3.1M | 94.4K | 96.6K | 87.8K | Java | 261 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 1.9M | 117.4K | 71.3K | 98.2K | Go | 166 |
| [open-telemetry](https://github.com/open-telemetry)  | 1.7M | 61.8K | 61.4K | 28.6K | Go | 148 |
| [grpc](https://github.com/grpc)  | 1.6M | 74.1K | 38.7K | 86.8K | C++ | 140 |
| [ClickHouse](https://github.com/ClickHouse)  | 1.6M | 119.1K | 34.8K | 33.8K | C++ | 140 |
| [grafana](https://github.com/grafana)  | 1.6M | 76.6K | 50.2K | 103.1K | Go | 140 |
| [envoyproxy](https://github.com/envoyproxy)  | 1.5M | 46.3K | 24.6K | 28.6K | C++ | 131 |
| [rust-lang](https://github.com/rust-lang)  | 1.5M | 219.1K | 61.1K | 78.6K | Rust | 127 |
| [python](https://github.com/python)  | 1.3M | 116.2K | 39.4K | 51.7K | Python | 114 |
| [docker](https://github.com/docker)  | 1.2M | 173.7K | 30.5K | 146.4K | Go | 106 |
| [golang](https://github.com/golang)  | 1.6M | 55.7K | 2.8K | 109.1K | Go | 88 |
| [ziglang](https://github.com/ziglang)  | 901.1K | 22.8K | 6.7K | 20.3K | Zig | 75 |
| [nats-io](https://github.com/nats-io)  | 802.4K | 36.2K | 12.9K | 31.4K | Go | 67 |
| [m3db](https://github.com/m3db)  | 736.7K | 10.4K | 5.7K | 4.6K | Go | 61 |
| [cilium](https://github.com/cilium)  | 681.5K | 30.9K | 21.8K | 23.6K | Go | 57 |
| [vitessio](https://github.com/vitessio)  | 671.3K | 36.7K | 11.6K | 15.9K | Go | 56 |
| [fluent](https://github.com/fluent)  | 639.2K | 38.4K | 11.9K | 24.2K | C | 53 |
| [kubevirt](https://github.com/kubevirt)  | 631.7K | 60.5K | 22.5K | 5.4K | Go | 53 |
| [cubeFS](https://github.com/cubeFS)  | 591.2K | 2.7K | 1.4K | 3.3K | Go | 49 |
| [istio](https://github.com/istio)  | 583.5K | 47.5K | 62.3K | 38.4K | Go | 49 |
| [openebs](https://github.com/openebs)  | 524.6K | 15.5K | 9.6K | 10.1K | Go | 44 |
| [argoproj](https://github.com/argoproj)  | 517.7K | 18.1K | 16.8K | 34.2K | Go | 43 |
| [knative](https://github.com/knative)  | 432.8K | 30.5K | 32.2K | 12.3K | Go | 36 |
| [apache](https://github.com/apache)  | 420.4K | 22.2K | 521 | 5.6K | C++ | 35 |
| [AthenZ](https://github.com/AthenZ)  | 405.5K | 3.4K | 2.4K | 762 | Java | 34 |
| [backstage](https://github.com/backstage)  | 402.2K | 40.4K | 12.9K | 21.7K | TypeScript | 34 |
| [tikv](https://github.com/tikv)  | 395.8K | 7.8K | 10.2K | 12.7K | Rust | 33 |
| [prometheus](https://github.com/prometheus)  | 377.3K | 30.8K | 19.1K | 89.5K | Go | 31 |
| [facebook](https://github.com/facebook)  | 357.5K | 15.5K | 13.2K | 203.5K | JavaScript | 30 |
| [containerd](https://github.com/containerd)  | 353.3K | 23.2K | 11.6K | 24.7K | Go | 29 |
| [pixie-io](https://github.com/pixie-io)  | 352.4K | 11.7K | 550 | 4.3K | C++ | 29 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.1K | 8.5K | 7.1K | 15.4K | Go | 29 |
| [dapr](https://github.com/dapr)  | 335.5K | 17.4K | 10.7K | 26.3K | Go | 28 |
| [LINBIT](https://github.com/LINBIT)  | 310.8K | 4.3K | 11 | 649 | Java | 26 |
| [vectordotdev](https://github.com/vectordotdev)  | 308.9K | 14.3K | 10.1K | 14.3K | Rust | 26 |
| [kyverno](https://github.com/kyverno)  | 280.2K | 8.7K | 5.1K | 4.1K | Go | 23 |
| [operator-framework](https://github.com/operator-framework)  | 265.4K | 18.9K | 14.3K | 13.3K | Go | 22 |
| [keptn](https://github.com/keptn)  | 256.7K | 11.7K | 8.3K | 1.8K | Go | 21 |
| [uber](https://github.com/uber)  | 253.7K | 1.5K | 248 | 5.9K | Go | 21 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.4K | 8.7K | 7.6K | 23.6K | Go | 21 |
| [goharbor](https://github.com/goharbor)  | 242.4K | 16.9K | 10.3K | 21.1K | Go | 20 |
| [linkerd](https://github.com/linkerd)  | 237.9K | 13.1K | 13.3K | 17.9K | Rust | 20 |
| [siderolabs](https://github.com/siderolabs)  | 232.7K | 7.1K | 9.6K | 4.7K | Go | 19 |
| [kata-containers](https://github.com/kata-containers)  | 223.8K | 22.7K | 10.2K | 7.0K | Go | 19 |
| [Netflix](https://github.com/Netflix)  | 208.1K | 4.1K | 2.2K | 551 | Java | 17 |
| [spiffe](https://github.com/spiffe)  | 206.9K | 8.6K | 3.9K | 2.9K | Go | 17 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 205.9K | 8.1K | 2.5K | 10.9K | Go | 17 |
| [litmuschaos](https://github.com/litmuschaos)  | 203.2K | 7.4K | 7.9K | 4.5K | TypeScript | 17 |
| [fluxcd](https://github.com/fluxcd)  | 200.2K | 27.6K | 9.4K | 19.8K | Go | 17 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4K | 4.4K | 4.5K | 4.8K | Python | 16 |
| [longhorn](https://github.com/longhorn)  | 193.6K | 9.9K | 6.5K | 5.1K | Go | 16 |
| [etcd-io](https://github.com/etcd-io)  | 185.2K | 26.3K | 11.2K | 52.2K | Go | 15 |
| [containers](https://github.com/containers)  | 174.4K | 18.2K | 10.1K | 17.8K | Go | 15 |
| [cert-manager](https://github.com/cert-manager)  | 145.5K | 13.1K | 4.8K | 10.8K | Go | 12 |
| [kubeedge](https://github.com/kubeedge)  | 141.3K | 7.3K | 3.9K | 6.7K | Go | 12 |
| [rook](https://github.com/rook)  | 132.1K | 24.7K | 7.1K | 11.1K | Go | 11 |
| [projectcontour](https://github.com/projectcontour)  | 130.7K | 4.9K | 3.8K | 4.1K | Go | 11 |
| [crossplane](https://github.com/crossplane)  | 127.7K | 11.2K | 3.5K | 7.7K | Go | 11 |
| [thanos-io](https://github.com/thanos-io)  | 123.5K | 4.7K | 3.9K | 12.4K | Go | 10 |
| [cortexproject](https://github.com/cortexproject)  | 120.3K | 4.8K | 3.7K | 5.1K | Go | 10 |
| [buildpacks](https://github.com/buildpacks)  | 116.6K | 16.7K | 3.5K | 2.7K | Go | 10 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 106.7K | 3.2K | 3.1K | 5.8K | Go | 9 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 105.6K | 17.9K | 3.2K | 4.3K | Python | 9 |
| [kedacore](https://github.com/kedacore)  | 105.4K | 3.8K | 4.2K | 6.7K | Go | 9 |
| [VKCOM](https://github.com/VKCOM)  | 101.4K | 9.8K | 3.8K | 877 | TypeScript | 8 |
| [karmada-io](https://github.com/karmada-io)  | 98.8K | 4.1K | 2.6K | 3.1K | Go | 8 |
| [in-toto](https://github.com/in-toto)  | 92.2K | 9.4K | 1.2K | 1.6K | Python | 8 |
| [cloudevents](https://github.com/cloudevents)  | 79.5K | 3.9K | 2.6K | 5.8K | Go | 7 |
| [coredns](https://github.com/coredns)  | 76.2K | 5.5K | 4.6K | 11.5K | Go | 6 |
| [gotd](https://github.com/gotd)  | 74.3K | 6.2K | 2.4K | 1.1K | Go | 6 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 72.2K | 1.4K | 1.6K | 1.1K | Go | 6 |
| [helm](https://github.com/helm)  | 71.5K | 28.9K | 26.8K | 48.7K | Go | 6 |
| [cri-o](https://github.com/cri-o)  | 71.2K | 9.1K | 5.7K | 4.5K | Go | 6 |
| [volcano-sh](https://github.com/volcano-sh)  | 68.5K | 5.4K | 2.3K | 3.1K | Go | 6 |
| [notaryproject](https://github.com/notaryproject)  | 64.4K | 3.6K | 1.9K | 3.3K | Go | 5 |
| [containerssh](https://github.com/containerssh)  | 62.1K | 1.7K | 2.2K | 2.2K | Go | 5 |
| [vuejs](https://github.com/vuejs)  | 60.2K | 3.5K | 2.3K | 202.6K | TypeScript | 5 |
| [go-faster](https://github.com/go-faster)  | 50.7K | 4.7K | 670 | 151 | Go | 4 |
| [theupdateframework](https://github.com/theupdateframework)  | 45.3K | 8.6K | 2.5K | 2.6K | Rust | 4 |
| [vitalif](https://github.com/vitalif)  | 40.8K | 1.2K | 12 | 70 | C++ | 3 |
| [k3s-io](https://github.com/k3s-io)  | 38.5K | 3.4K | 3.2K | 25.3K | Go | 3 |
| [ogen-go](https://github.com/ogen-go)  | 38.1K | 3.8K | 1.5K | 472 | Go | 3 |
| [containernetworking](https://github.com/containernetworking)  | 34.7K | 2.8K | 1.1K | 6.4K | Go | 3 |
| [falcosecurity](https://github.com/falcosecurity)  | 13.9K | 3.5K | 1.4K | 5.6K | C++ | 1 |
| [carina-io](https://github.com/carina-io)  | 12.8K | 749 | 118 | 588 | Go | 1 |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3K | 251 | 117 | 2.1K | Go | 0 |


### Repositories
| Repository | SLOC | Commits | PRs | Stars | Language | Effort |
|------------|------|---------|-----|-------|----------|--------|
| [torvalds/linux](https://github.com/torvalds/linux/)  | 19.1M | 1.1M | 761 | 147.5K | C | 1.5K |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3.3M | 144.4K | 22.7K | 171.8K | C++ | 279 |
| [ydb-platform/ydb](https://github.com/ydb-platform/ydb/)  | 2.8M | 9.3K | 35 | 2.7K | C++ | 237 |
| [elastic/elasticsearch](https://github.com/elastic/elasticsearch/)  | 2.4M | 67.8K | 62.4K | 62.9K | Java | 202 |
| [rust-lang/rust](https://github.com/rust-lang/rust/)  | 1.5M | 219.1K | 61.1K | 78.6K | Rust | 127 |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1.4M | 114.4K | 73.5K | 96.4K | Go | 124 |
| [python/cpython](https://github.com/python/cpython/)  | 1.3M | 116.2K | 39.4K | 51.7K | Python | 114 |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1.1M | 109.7K | 32.2K | 27.4K | C++ | 93 |
| [golang/go](https://github.com/golang/go/)  | 1.6M | 55.7K | 2.8K | 109.1K | Go | 88 |
| [ziglang/zig](https://github.com/ziglang/zig/)  | 901.1K | 22.8K | 6.7K | 20.3K | Zig | 75 |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889.8K | 41.3K | 33.5K | 54.2K | TypeScript | 74 |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 807.5K | 17.3K | 17.1K | 21.5K | C++ | 67 |
| [elastic/beats](https://github.com/elastic/beats/)  | 617.7K | 16.2K | 25.7K | 11.5K | Go | 51 |
| [vitessio/vitess](https://github.com/vitessio/vitess/)  | 610.9K | 32.3K | 9.6K | 15.7K | Go | 51 |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 514.8K | 52.6K | 21.6K | 37.2K | C++ | 43 |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 512.2K | 96 | 103 | 22 | C++ | 43 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 504.5K | 8.4K | 454 | 205 | C++ | 42 |
| [cubeFS/cubefs](https://github.com/cubeFS/cubefs/)  | 478.3K | 2.2K | 1.2K | 3.1K | Go | 40 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477.3K | 4.2K | 3.5K | 4.3K | Go | 40 |
| [tikv/tikv](https://github.com/tikv/tikv/)  | 395.8K | 7.8K | 10.2K | 12.7K | Rust | 33 |
| [backstage/backstage](https://github.com/backstage/backstage/)  | 393.7K | 38.2K | 12.2K | 21.7K | TypeScript | 33 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389.2K | 9.6K | 16.1K | 1.5K | Go | 32 |
| [facebook/react](https://github.com/facebook/react/)  | 357.5K | 15.5K | 13.2K | 203.5K | JavaScript | 30 |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 352.4K | 11.7K | 550 | 4.3K | C++ | 29 |
| [AthenZ/athenz](https://github.com/AthenZ/athenz/)  | 346.7K | 2.5K | 1.8K | 751 | Java | 29 |
| [istio/istio](https://github.com/istio/istio/)  | 339.8K | 19.8K | 26.4K | 32.5K | Go | 28 |
| [kubernetes/autoscaler](https://github.com/kubernetes/autoscaler/)  | 332.8K | 6.4K | 3.7K | 6.5K | Go | 28 |
| [cilium/cilium](https://github.com/cilium/cilium/)  | 324.8K | 22.2K | 17.3K | 14.5K | Go | 27 |
| [docker/docker-ce](https://github.com/docker/docker-ce/)  | 324.4K | 54.3K | 662 | 5.5K | Go | 27 |
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310.8K | 4.3K | 11 | 649 | Java | 26 |
| [apache/mesos](https://github.com/apache/mesos/)  | 305.6K | 18.1K | 450 | 5.2K | C++ | 25 |
| [docker/labs](https://github.com/docker/labs/)  | 304.4K | 718 | 398 | 11.1K | PHP | 25 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284.4K | 1.5K | 1.3K | 465 | C Header | 24 |
| [kubevirt/kubevirt](https://github.com/kubevirt/kubevirt/)  | 281.1K | 16.3K | 6.9K | 3.9K | Go | 23 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274.6K | 9.3K | 9.3K | 12.9K | Rust | 23 |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 270.5K | 51.9K | 24.6K | 3.5K | Go | 23 |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 264.4K | 1.6K | 36 | 24 | JavaScript | 22 |
| [kyverno/kyverno](https://github.com/kyverno/kyverno/)  | 259.5K | 5.5K | 3.9K | 3.5K | Go | 22 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241.7K | 4.7K | 5.1K | 18.4K | Go | 20 |
| [open-policy-agent/opa](https://github.com/open-policy-agent/opa/)  | 235.3K | 4.4K | 3.5K | 7.7K | Go | 20 |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235.1K | 5.7K | 6.8K | 10.4K | Java | 20 |
| [uber/peloton](https://github.com/uber/peloton/)  | 216.3K | 705 | 10 | 582 | Go | 18 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194.1K | 181 | 125 | 1.1K | C Header | 16 |
| [goharbor/harbor](https://github.com/goharbor/harbor/)  | 191.4K | 11.5K | 8.2K | 19.5K | Go | 16 |
| [nats-io/nats-server](https://github.com/nats-io/nats-server/)  | 190.6K | 6.8K | 2.5K | 12.3K | Go | 16 |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186.6K | 1 | 2 | 2 | JavaScript | 16 |
| [fluent/fluent-bit](https://github.com/fluent/fluent-bit/)  | 182.4K | 9.2K | 3.3K | 4.3K | C | 15 |
| [cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian/)  | 175.1K | 4.1K | 4.3K | 4.7K | Python | 15 |
| [containers/podman](https://github.com/containers/podman/)  | 174.4K | 18.2K | 10.1K | 17.8K | Go | 15 |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173.3K | 8.8K | 5.3K | 2.7K | Go | 14 |
| [argoproj/argo-workflows](https://github.com/argoproj/argo-workflows/)  | 169.7K | 4.1K | 4.8K | 12.5K | Go | 14 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.8K | 8.8K | 5.6K | 1.2K | Java | 14 |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167.6K | 19.3K | 10.3K | 14.7K | Go | 14 |
| [keptn/keptn](https://github.com/keptn/keptn/)  | 162.2K | 8.2K | 6.7K | 1.6K | Go | 14 |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157.3K | 1.6K | 1.2K | 319 | Java | 13 |
| [spiffe/spire](https://github.com/spiffe/spire/)  | 156.6K | 5.8K | 2.7K | 1.3K | Go | 13 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155.1K | 10.9K | 6.7K | 47.8K | Go | 13 |
| [containerd/containerd](https://github.com/containerd/containerd/)  | 150.2K | 11.9K | 5.8K | 13.3K | Go | 13 |
| [argoproj/argo-cd](https://github.com/argoproj/argo-cd/)  | 149.4K | 5.1K | 5.8K | 12.3K | Go | 12 |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144.5K | 5.9K | 1.8K | 8.1K | Go | 12 |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143.2K | 4.4K | 3.9K | 17.6K | Go | 12 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.8K | 929 | 1.7K | 224 | C Header | 12 |
| [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize/)  | 139.4K | 6.3K | 3.2K | 9.4K | Go | 12 |
| [kubernetes/apiserver](https://github.com/kubernetes/apiserver/)  | 131.3K | 6.1K | 25 | 494 | Go | 11 |
| [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers/)  | 130.8K | 10.1K | 3.7K | 3.8K | Go | 11 |
| [siderolabs/talos](https://github.com/siderolabs/talos/)  | 130.1K | 3.8K | 5.1K | 3.6K | Go | 11 |
| [etcd-io/etcd](https://github.com/etcd-io/etcd/)  | 125.6K | 19.3K | 9.1K | 42.8K | Go | 10 |
| [cortexproject/cortex](https://github.com/cortexproject/cortex/)  | 118.9K | 4.6K | 3.4K | 5.1K | Go | 10 |
| [apache/aurora](https://github.com/apache/aurora/)  | 114.7K | 4.9K | 71 | 628 | Java | 10 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 113.6K | 2.8K | 63 | 10 | TypeScript | 9 |
| [thanos-io/thanos](https://github.com/thanos-io/thanos/)  | 113.2K | 3.1K | 3.6K | 11.5K | Go | 9 |
| [projectcontour/contour](https://github.com/projectcontour/contour/)  | 111.5K | 4.1K | 3.1K | 3.3K | Go | 9 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 106.9K | 2.2K | 1.8K | 218 | Go | 9 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105.9K | 1.4K | 1.7K | 980 | Go | 9 |
| [emissary-ingress/emissary](https://github.com/emissary-ingress/emissary/)  | 105.2K | 17.9K | 3.2K | 4.2K | Python | 9 |
| [dapr/dapr](https://github.com/dapr/dapr/)  | 100.9K | 3.9K | 3.1K | 20.6K | Go | 8 |
| [cert-manager/cert-manager](https://github.com/cert-manager/cert-manager/)  | 100.4K | 7.5K | 2.7K | 10.6K | Go | 8 |
| [openebs/maya](https://github.com/openebs/maya/)  | 100.7K | 1.7K | 1.6K | 180 | Go | 8 |
| [knative/serving](https://github.com/knative/serving/)  | 96.8K | 8.3K | 9.4K | 4.8K | Go | 8 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.3K | 3.2K | 3.6K | 1.4K | Java | 8 |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 95.9K | 2.9K | 313 | 2.2K | Go | 8 |
| [rook/rook](https://github.com/rook/rook/)  | 95.8K | 9.6K | 7.1K | 10.8K | Go | 8 |
| [karmada-io/karmada](https://github.com/karmada-io/karmada/)  | 94.6K | 3.6K | 2.3K | 3.1K | Go | 8 |
| [elastic/logstash](https://github.com/elastic/logstash/)  | 92.8K | 10.3K | 8.5K | 13.3K | Ruby | 8 |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90.6K | 20.9K | 7.6K | 26.1K | Go | 8 |
| [cubeFS/cubefs-blobstore](https://github.com/cubeFS/cubefs-blobstore/)  | 87.3K | 82 | 3 | 14 | Go | 7 |
| [knative/eventing](https://github.com/knative/eventing/)  | 87.6K | 3.5K | 4.3K | 1.2K | Go | 7 |
| [operator-framework/operator-lifecy~](https://github.com/operator-framework/operator-lifecycle-manager/)  | 87.1K | 4.1K | 2.1K | 1.4K | Go | 7 |
| [fluent/onigmo](https://github.com/fluent/onigmo/)  | 86.5K | 1.1K | 5 | 2 | C Header | 7 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84.4K | 3.4K | 3.1K | 210 | Go | 7 |
| [dapr/components-contrib](https://github.com/dapr/components-contrib/)  | 83.7K | 2.8K | 1.4K | 426 | Go | 7 |
| [argoproj/argo-rollouts](https://github.com/argoproj/argo-rollouts/)  | 82.5K | 1.2K | 1.5K | 1.9K | Go | 7 |
| [linkerd/linkerd2](https://github.com/linkerd/linkerd2/)  | 82.1K | 5.4K | 6.3K | 9.3K | Go | 7 |
| [fluent/fluentd](https://github.com/fluent/fluentd/)  | 80.4K | 6.5K | 1.8K | 11.8K | Ruby | 7 |
| [kubevirt/containerized-data-import~](https://github.com/kubevirt/containerized-data-importer/)  | 79.1K | 2.4K | 2.1K | 311 | Go | 7 |
| [docker/cli](https://github.com/docker/cli/)  | 78.9K | 8.4K | 2.6K | 3.9K | Go | 7 |
| [litmuschaos/litmus](https://github.com/litmuschaos/litmus/)  | 78.7K | 2.5K | 2.5K | 3.5K | TypeScript | 7 |
| [fluent/fluentbit-website-v3](https://github.com/fluent/fluentbit-website-v3/)  | 78.1K | 386 | 27 | 4 | JavaScript | 7 |
| [kubernetes/ingress-gce](https://github.com/kubernetes/ingress-gce/)  | 78.2K | 4.5K | 1.5K | 1.1K | Go | 7 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77.2K | 3.6K | 2.1K | 244 | Go | 6 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76.4K | 630 | 1.1K | 3.2K | Go | 6 |
| [chaos-mesh/chaos-mesh](https://github.com/chaos-mesh/chaos-mesh/)  | 74.9K | 1.6K | 2.3K | 5.5K | Go | 6 |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73.7K | 6.5K | 1.2K | 628 | Go | 6 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73.7K | 1.7K | 634 | 1.1K | Java | 6 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73.6K | 3.8K | 2.6K | 543 | Go | 6 |
| [linkerd/linkerd](https://github.com/linkerd/linkerd/)  | 73.2K | 1.4K | 1.3K | 5.3K | Scala | 6 |
| [grafana/k6](https://github.com/grafana/k6/)  | 73.4K | 5.5K | 1.4K | 19.6K | Go | 6 |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72.3K | 4.8K | 5.2K | 2.4K | TypeScript | 6 |
| [dragonflyoss/Dragonfly2](https://github.com/dragonflyoss/Dragonfly2/)  | 72.2K | 1.4K | 1.6K | 1.1K | Go | 6 |
| [jaegertracing/jaeger](https://github.com/jaegertracing/jaeger/)  | 71.8K | 1.9K | 2.3K | 17.2K | Go | 6 |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71.8K | 633 | 232 | 1.6K | JavaScript | 6 |
| [longhorn/longhorn-manager](https://github.com/longhorn/longhorn-manager/)  | 68.8K | 2.9K | 1.7K | 132 | Go | 6 |
| [kedacore/keda](https://github.com/kedacore/keda/)  | 68.2K | 1.8K | 2.1K | 6.4K | Go | 6 |
| [kubeedge/kubeedge](https://github.com/kubeedge/kubeedge/)  | 66.6K | 4.6K | 2.8K | 5.6K | Go | 6 |
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65.6K | 3.8K | 204 | 7.4K | Go | 5 |
| [nats-io/nats.c](https://github.com/nats-io/nats.c/)  | 64.9K | 972 | 399 | 301 | C | 5 |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63.9K | 2.7K | 1.4K | 2.1K | Go | 5 |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63.8K | 859 | 941 | 3.5K | C# | 5 |
| [fluent/fluent-bit-website-old](https://github.com/fluent/fluent-bit-website-old/)  | 63.1K | 19 | 0 | 2 | JavaScript | 5 |
| [ydb-platform/ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk/)  | 62.9K | 2.8K | 465 | 80 | Go | 5 |
| [knative/pkg](https://github.com/knative/pkg/)  | 61.1K | 1.9K | 2.3K | 234 | Go | 5 |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61.5K | 236 | 214 | 120 | Go | 5 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.1K | 1.7K | 2.6K | 1.7K | TypeScript | 5 |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60.9K | 1.2K | 447 | 27 | Rust | 5 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60.5K | 3.1K | 1.8K | 6.1K | Go | 5 |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60.2K | 3.5K | 2.3K | 202.6K | TypeScript | 5 |
| [kubernetes/legacy-cloud-providers](https://github.com/kubernetes/legacy-cloud-providers/)  | 60.2K | 1.8K | 0 | 47 | Go | 5 |
| [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/)  | 60.1K | 2.5K | 30 | 661 | Go | 5 |
| [ydb-platform/ydb-nodejs-genproto](https://github.com/ydb-platform/ydb-nodejs-genproto/)  | 59.3K | 9 | 3 | 1 | JavaScript | 5 |
| [buildpacks/pack](https://github.com/buildpacks/pack/)  | 58.4K | 3.3K | 957 | 1.9K | Go | 5 |
| [envoyproxy/envoy-mobile](https://github.com/envoyproxy/envoy-mobile/)  | 57.6K | 1.7K | 2.2K | 550 | Java | 5 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9K | 2.2K | 2.7K | 2.2K | C# | 5 |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56.8K | 6.9K | 4.9K | 14.5K | Go | 5 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 56.9K | 1.8K | 1.6K | 286 | Go | 5 |
| [volcano-sh/volcano](https://github.com/volcano-sh/volcano/)  | 55.6K | 4.6K | 1.6K | 2.8K | Go | 5 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 54.1K | 1.8K | 1.1K | 457 | Rust | 5 |
| [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy/)  | 54.8K | 2.1K | 2.2K | 1.7K | Rust | 5 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53.4K | 4.4K | 37 | 11 | Go | 4 |
| [gotd/td](https://github.com/gotd/td/)  | 52.6K | 3.8K | 788 | 1.5K | Go | 4 |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.3K | 1.5K | 990 | 420 | TypeScript | 4 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51.2K | 3.1K | 8 | 206 | Go | 4 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.9K | 1.9K | 2.3K | 3.6K | Go | 4 |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50.8K | 2.4K | 997 | 232 | Go | 4 |
| [coredns/coredns](https://github.com/coredns/coredns/)  | 50.6K | 3.5K | 3.7K | 10.3K | Go | 4 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50.6K | 233 | 194 | 9 | Go | 4 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector/)  | 50.3K | 4.4K | 5.9K | 2.7K | Go | 4 |
| [nats-io/nats.java](https://github.com/nats-io/nats.java/)  | 49.1K | 1.4K | 510 | 461 | Java | 4 |
| [kubernetes-sigs/multi-tenancy](https://github.com/kubernetes-sigs/multi-tenancy/)  | 48.9K | 2.3K | 1.6K | 930 | Go | 4 |
| [envoyproxy/pytooling](https://github.com/envoyproxy/pytooling/)  | 48.9K | 615 | 619 | 6 | Python | 4 |
| [jaegertracing/jaeger-ui](https://github.com/jaegertracing/jaeger-ui/)  | 48.8K | 1.9K | 781 | 894 | JavaScript | 4 |
| [nats-io/nats-streaming-server](https://github.com/nats-io/nats-streaming-server/)  | 48.8K | 1.7K | 651 | 2.4K | Go | 4 |
| [grpc/grpc-swift](https://github.com/grpc/grpc-swift/)  | 48.3K | 1.5K | 981 | 1.7K | Swift | 4 |
| [cri-o/cri-o](https://github.com/cri-o/cri-o/)  | 47.4K | 7.6K | 5.5K | 4.4K | Go | 4 |
| [nats-io/nats.net](https://github.com/nats-io/nats.net/)  | 46.7K | 747 | 430 | 566 | C# | 4 |
| [operator-framework/operator-sdk](https://github.com/operator-framework/operator-sdk/)  | 46.7K | 3.1K | 3.8K | 6.3K | Go | 4 |
| [knative/client](https://github.com/knative/client/)  | 46.6K | 1.1K | 1.2K | 303 | Go | 4 |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46.5K | 4.4K | 5 | 1 | Go | 4 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46.4K | 241 | 358 | 20 | C | 4 |
| [vitessio/website](https://github.com/vitessio/website/)  | 46.4K | 2.9K | 1.9K | 36 | JavaScript | 4 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45.7K | 790 | 309 | 38 | JavaScript | 4 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.7K | 1.7K | 1.2K | 413 | Python | 4 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45.2K | 1.8K | 543 | 113 | Go | 4 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44.8K | 4.4K | 4.9K | 12.3K | Go | 4 |
| [notaryproject/notary](https://github.com/notaryproject/notary/)  | 43.7K | 3.1K | 992 | 2.9K | Go | 4 |
| [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper/)  | 43.7K | 1.2K | 1.5K | 2.9K | Go | 4 |
| [containerd/nerdctl](https://github.com/containerd/nerdctl/)  | 43.5K | 2.4K | 1.2K | 5.6K | Go | 4 |
| [istio/old_mixer_repo](https://github.com/istio/old_mixer_repo/)  | 42.5K | 741 | 1.9K | 68 | Go | 4 |
| [crossplane/crossplane](https://github.com/crossplane/crossplane/)  | 42.4K | 4.6K | 2.1K | 6.7K | Go | 4 |
| [longhorn/longhorn-tests](https://github.com/longhorn/longhorn-tests/)  | 42.3K | 1.5K | 1.2K | 13 | Python | 4 |
| [grpc/grpc-node](https://github.com/grpc/grpc-node/)  | 42.3K | 4.2K | 1.3K | 3.8K | TypeScript | 4 |
| [open-telemetry/opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python/)  | 42.2K | 1.4K | 1.7K | 1.2K | Python | 4 |
| [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)  | 42.1K | 2.2K | 1.3K | 1.8K | Go | 4 |
| [open-telemetry/opentelemetry-cpp-c~](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42.5K | 157 | 180 | 78 | Python | 4 |
| [helm/helm](https://github.com/helm/helm/)  | 41.7K | 6.8K | 4.8K | 23.8K | Go | 3 |
| [open-telemetry/opentelemetry-ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4K | 268 | 112 | 85 | C++ | 3 |
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40.8K | 1.2K | 12 | 70 | C++ | 3 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 40.8K | 52 | 3 | 5 | JavaScript | 3 |
| [keptn/tutorials](https://github.com/keptn/tutorials/)  | 40.8K | 622 | 193 | 10 | JavaScript | 3 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40.4K | 9.5K | 2.8K | 757 | TypeScript | 3 |
| [dapr/java-sdk](https://github.com/dapr/java-sdk/)  | 39.6K | 409 | 490 | 223 | Java | 3 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39.8K | 177 | 142 | 13 | Go | 3 |
| [nats-io/nats.go](https://github.com/nats-io/nats.go/)  | 38.8K | 1.7K | 716 | 4.3K | Go | 3 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38.1K | 2.9K | 1.8K | 6.2K | Go | 3 |
| [uber/kraken](https://github.com/uber/kraken/)  | 37.4K | 867 | 238 | 5.3K | Go | 3 |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37.3K | 799 | 283 | 321 | Go | 3 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36.6K | 3.2K | 700 | 454 | Go | 3 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36.3K | 35 | 7 | 1 | C | 3 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36.3K | 298 | 358 | 83 | Go | 3 |
| [argoproj/argo-events](https://github.com/argoproj/argo-events/)  | 35.9K | 1.3K | 1.4K | 1.8K | Go | 3 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35.4K | 662 | 313 | 278 | Go | 3 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34.2K | 1.2K | 462 | 2.3K | Go | 3 |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34.1K | 13 | 0 | 6 | JavaScript | 3 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33.6K | 1.4K | 420 | 78 | Go | 3 |
| [containerssh/libcontainerssh](https://github.com/containerssh/libcontainerssh/)  | 33.6K | 324 | 384 | 28 | Go | 3 |
| [open-telemetry/opentelemetry-cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.5K | 1.5K | 1.1K | 432 | C++ | 3 |
| [kubeedge/sedna](https://github.com/kubeedge/sedna/)  | 33.4K | 475 | 215 | 428 | Python | 3 |
| [kubernetes/perf-tests](https://github.com/kubernetes/perf-tests/)  | 33.3K | 3.1K | 1.8K | 764 | Go | 3 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.3K | 732 | 850 | 230 | C# | 3 |
| [in-toto/github-action](https://github.com/in-toto/github-action/)  | 33.2K | 3 | 0 | 4 | JavaScript | 3 |
| [buildpacks/lifecycle](https://github.com/buildpacks/lifecycle/)  | 32.9K | 1.6K | 665 | 158 | Go | 3 |
| [operator-framework/operator-regist~](https://github.com/operator-framework/operator-registry/)  | 32.6K | 1.1K | 780 | 181 | Go | 3 |
| [goharbor/harbor-operator](https://github.com/goharbor/harbor-operator/)  | 32.4K | 1.3K | 537 | 293 | Go | 3 |
| [keptn/lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit/)  | 32.1K | 497 | 666 | 56 | Go | 3 |
| [cloudevents/sdk-go](https://github.com/cloudevents/sdk-go/)  | 32.1K | 624 | 542 | 613 | Go | 3 |
| [nats-io/nsc](https://github.com/nats-io/nsc/)  | 31.9K | 604 | 358 | 61 | Go | 3 |
| [prometheus/alertmanager](https://github.com/prometheus/alertmanager/)  | 31.8K | 2.8K | 1.6K | 5.5K | Go | 3 |
| [ydb-platform/xorm](https://github.com/ydb-platform/xorm/)  | 31.8K | 1.7K | 9 | 0 | Go | 3 |
| [litmuschaos/test-tools](https://github.com/litmuschaos/test-tools/)  | 31.8K | 226 | 300 | 27 | C | 3 |
| [kubevirt/hyperconverged-cluster-op~](https://github.com/kubevirt/hyperconverged-cluster-operator/)  | 31.3K | 1.4K | 2.1K | 106 | Go | 3 |
| [knative/func](https://github.com/knative/func/)  | 31.2K | 1.1K | 1.1K | 156 | Go | 3 |
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 30.9K | 508 | 591 | 519 | PHP | 3 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30.9K | 1.2K | 308 | 232 | Go | 3 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30.9K | 18 | 4 | 0 | C Header | 3 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5K | 574 | 597 | 997 | Rust | 3 |
| [fluxcd/source-controller](https://github.com/fluxcd/source-controller/)  | 30.1K | 1.9K | 759 | 191 | Go | 3 |
| [docker/machine](https://github.com/docker/machine/)  | 29.9K | 3.4K | 1.9K | 6.5K | Go | 3 |
| [cilium/ebpf](https://github.com/cilium/ebpf/)  | 29.9K | 1.2K | 686 | 4.1K | Go | 2 |
| [fluxcd/flux](https://github.com/fluxcd/flux/)  | 29.8K | 5.2K | 1.9K | 6.9K | Go | 2 |
| [nats-io/nats.deno](https://github.com/nats-io/nats.deno/)  | 29.7K | 442 | 413 | 115 | TypeScript | 2 |
| [m3db/m3cluster](https://github.com/m3db/m3cluster/)  | 29.6K | 238 | 227 | 21 | Go | 2 |
| [kubevirt/project-infra](https://github.com/kubevirt/project-infra/)  | 29.4K | 2.4K | 2.5K | 21 | Go | 2 |
| [k3s-io/k3s](https://github.com/k3s-io/k3s/)  | 29.2K | 2.6K | 2.6K | 22.4K | Go | 2 |
| [kubevirt/web-ui-components](https://github.com/kubevirt/web-ui-components/)  | 28.5K | 416 | 500 | 9 | JavaScript | 2 |
| [litmuschaos/litmus-go](https://github.com/litmuschaos/litmus-go/)  | 28.5K | 425 | 558 | 51 | Go | 2 |
| [jaegertracing/jaeger-operator](https://github.com/jaegertracing/jaeger-operator/)  | 28.4K | 1.9K | 1.3K | 904 | Go | 2 |
| [open-telemetry/opentelemetry-log-c~](https://github.com/open-telemetry/opentelemetry-log-collection/)  | 28.5K | 267 | 373 | 92 | Go | 2 |
| [envoyproxy/nighthawk](https://github.com/envoyproxy/nighthawk/)  | 27.8K | 578 | 655 | 272 | C++ | 2 |
| [kubernetes-sigs/alibaba-cloud-csi-~](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/)  | 27.6K | 1.3K | 590 | 444 | Go | 2 |
| [fluxcd/flagger](https://github.com/fluxcd/flagger/)  | 27.5K | 2.6K | 753 | 4.1K | Go | 2 |
| [fluxcd/pkg](https://github.com/fluxcd/pkg/)  | 27.5K | 1.7K | 454 | 30 | Go | 2 |
| [kubernetes/release](https://github.com/kubernetes/release/)  | 27.3K | 5.1K | 2.3K | 436 | Go | 2 |
| [docker/docker-py](https://github.com/docker/docker-py/)  | 27.3K | 3.2K | 1.5K | 6.1K | Python | 2 |
| [open-telemetry/opentelemetry-go-co~](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.2K | 1.4K | 3.2K | 684 | Go | 2 |
| [cilium/hubble-ui](https://github.com/cilium/hubble-ui/)  | 26.7K | 385 | 426 | 237 | TypeScript | 2 |
| [ClickHouse/libhdfs3](https://github.com/ClickHouse/libhdfs3/)  | 26.7K | 64 | 34 | 23 | C++ | 2 |
| [ClickHouse/clickhouse-cpp](https://github.com/ClickHouse/clickhouse-cpp/)  | 26.6K | 673 | 172 | 215 | C++ | 2 |
| [kubevirt/vm-import-operator](https://github.com/kubevirt/vm-import-operator/)  | 26.6K | 900 | 415 | 13 | Go | 2 |
| [dapr/dotnet-sdk](https://github.com/dapr/dotnet-sdk/)  | 26.4K | 550 | 465 | 967 | C# | 2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26.3K | 562 | 644 | 80 | Java | 2 |
| [containernetworking/plugins](https://github.com/containernetworking/plugins/)  | 26.2K | 1.6K | 484 | 1.7K | Go | 2 |
| [open-telemetry/opentelemetry-swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9K | 618 | 252 | 126 | Swift | 2 |
| [longhorn/longhorn-ui](https://github.com/longhorn/longhorn-ui/)  | 25.6K | 817 | 595 | 92 | JavaScript | 2 |
| [envoyproxy/java-control-plane](https://github.com/envoyproxy/java-control-plane/)  | 25.4K | 235 | 206 | 246 | Protocol Buffers | 2 |
| [grafana/grafana-plugin-sdk-go](https://github.com/grafana/grafana-plugin-sdk-go/)  | 25.3K | 416 | 531 | 157 | Go | 2 |
| [kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)  | 25.3K | 642 | 349 | 733 | Go | 2 |
| [containerd/stargz-snapshotter](https://github.com/containerd/stargz-snapshotter/)  | 25.1K | 1.3K | 1.3K | 796 | Go | 2 |
| [kubernetes-sigs/kubefed](https://github.com/kubernetes-sigs/kubefed/)  | 25.4K | 2.7K | 978 | 2.4K | Go | 2 |
| [docker/compose-cli](https://github.com/docker/compose-cli/)  | 24.9K | 3.2K | 1.2K | 910 | Go | 2 |
| [grafana/terraform-provider-grafana](https://github.com/grafana/terraform-provider-grafana/)  | 24.6K | 697 | 523 | 312 | Go | 2 |
| [siderolabs/theila](https://github.com/siderolabs/theila/)  | 24.3K | 100 | 114 | 45 | TypeScript | 2 |
| [fluent/cmetrics](https://github.com/fluent/cmetrics/)  | 24.9K | 431 | 149 | 9 | C | 2 |
| [ydb-platform/ydb-embedded-ui](https://github.com/ydb-platform/ydb-embedded-ui/)  | 24.5K | 611 | 334 | 13 | TypeScript | 2 |
| [in-toto/toto-pip](https://github.com/in-toto/toto-pip/)  | 23.7K | 4.9K | 0 | 0 | Python | 2 |
| [kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/)  | 23.6K | 2.5K | 1.2K | 4.3K | Go | 2 |
| [open-telemetry/opentelemetry-opera~](https://github.com/open-telemetry/opentelemetry-operator/)  | 23.6K | 781 | 1.1K | 675 | Go | 2 |
| [argoproj/gitops-engine](https://github.com/argoproj/gitops-engine/)  | 23.4K | 254 | 454 | 1.4K | Go | 2 |
| [longhorn/longhorn-engine](https://github.com/longhorn/longhorn-engine/)  | 23.3K | 1.5K | 715 | 297 | Go | 2 |
| [nats-io/nats-mq](https://github.com/nats-io/nats-mq/)  | 23.2K | 88 | 13 | 18 | Go | 2 |
| [openebs/elves](https://github.com/openebs/elves/)  | 23.2K | 234 | 47 | 17 | Python | 2 |
| [openebs/node-disk-manager](https://github.com/openebs/node-disk-manager/)  | 23.2K | 451 | 533 | 170 | Go | 2 |
| [istio/old_pilot_repo](https://github.com/istio/old_pilot_repo/)  | 22.9K | 688 | 1.2K | 138 | Go | 2 |
| [nats-io/nats.rs](https://github.com/nats-io/nats.rs/)  | 22.7K | 1.3K | 687 | 698 | Rust | 2 |
| [containerd/cri](https://github.com/containerd/cri/)  | 22.6K | 2.8K | 1.1K | 897 | Go | 2 |
| [istio/proxy](https://github.com/istio/proxy/)  | 22.4K | 2.4K | 4.2K | 700 | C++ | 2 |
| [kubernetes-sigs/azuredisk-csi-driv~](https://github.com/kubernetes-sigs/azuredisk-csi-driver/)  | 21.7K | 2.7K | 1.3K | 114 | Go | 2 |
| [envoyproxy/gateway](https://github.com/envoyproxy/gateway/)  | 21.6K | 539 | 632 | 886 | Go | 2 |
| [cri-o/image](https://github.com/cri-o/image/)  | 21.6K | 1.8K | 2 | 2 | Go | 2 |
| [istio/operator](https://github.com/istio/operator/)  | 21.2K | 508 | 780 | 174 | Go | 2 |
| [kedacore/sample-dotnet-worker-serv~](https://github.com/kedacore/sample-dotnet-worker-servicebus-queue/)  | 21.2K | 38 | 30 | 68 | JavaScript | 2 |
| [ydb-platform/ydb-java-sdk](https://github.com/ydb-platform/ydb-java-sdk/)  | 20.9K | 598 | 32 | 12 | Java | 2 |
| [containerd/overlaybd](https://github.com/containerd/overlaybd/)  | 20.9K | 285 | 147 | 144 | C++ | 2 |
| [kubernetes-sigs/structured-merge-d~](https://github.com/kubernetes-sigs/structured-merge-diff/)  | 20.9K | 495 | 208 | 74 | Go | 2 |
| [crossplane/crossplane-runtime](https://github.com/crossplane/crossplane-runtime/)  | 20.8K | 1.2K | 260 | 98 | Go | 2 |
| [prometheus/procfs](https://github.com/prometheus/procfs/)  | 20.8K | 572 | 372 | 618 | Go | 2 |
| [prometheus/client_golang](https://github.com/prometheus/client_golang/)  | 20.7K | 1.4K | 741 | 4.4K | Go | 2 |
| [docker/buildx](https://github.com/docker/buildx/)  | 20.6K | 1.4K | 723 | 2.4K | Go | 2 |
| [kubernetes/cloud-provider-vsphere](https://github.com/kubernetes/cloud-provider-vsphere/)  | 20.4K | 1.8K | 472 | 189 | Go | 2 |
| [nats-io/nats-site](https://github.com/nats-io/nats-site/)  | 20.4K | 2.7K | 743 | 83 | JavaScript | 2 |
| [fluxcd/flux2](https://github.com/fluxcd/flux2/)  | 20.4K | 2.6K | 1.2K | 4.5K | Go | 2 |
| [kubernetes/cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack/)  | 20.2K | 2.6K | 1.3K | 507 | Go | 2 |
| [open-telemetry/opentelemetry-ruby-~](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 19.9K | 895 | 241 | 29 | Ruby | 2 |
| [operator-framework/operatorhub.io](https://github.com/operator-framework/operatorhub.io/)  | 19.8K | 620 | 330 | 40 | TypeScript | 2 |
| [jaegertracing/jaeger-client-go](https://github.com/jaegertracing/jaeger-client-go/)  | 19.3K | 433 | 408 | 1.3K | Go | 2 |
| [prometheus/node_exporter](https://github.com/prometheus/node_exporter/)  | 19.2K | 1.9K | 1.4K | 8.5K | Go | 2 |
| [envoyproxy/data-plane-api](https://github.com/envoyproxy/data-plane-api/)  | 19.2K | 2.3K | 540 | 510 | Protocol Buffers | 2 |
| [kubernetes-sigs/descheduler](https://github.com/kubernetes-sigs/descheduler/)  | 19.7K | 1.3K | 667 | 3.2K | Go | 2 |
| [cilium/cilium-olm](https://github.com/cilium/cilium-olm/)  | 19.2K | 341 | 55 | 11 | Go Template | 2 |
| [grpc/grpc-dart](https://github.com/grpc/grpc-dart/)  | 18.7K | 228 | 297 | 759 | Dart | 2 |
| [crossplane/oam-kubernetes-runtime](https://github.com/crossplane/oam-kubernetes-runtime/)  | 18.6K | 415 | 203 | 279 | Go | 2 |
| [cloud-custodian/cel-python](https://github.com/cloud-custodian/cel-python/)  | 18.5K | 115 | 24 | 60 | Python | 2 |
| [cilium/cilium-cli](https://github.com/cilium/cilium-cli/)  | 18.3K | 1.3K | 1.1K | 205 | Go | 2 |
| [knative/client-pkg](https://github.com/knative/client-pkg/)  | 18.1K | 80 | 86 | 3 | Go | 2 |
| [AthenZ/athenz-authorizer](https://github.com/AthenZ/athenz-authorizer/)  | 18.5K | 85 | 2 | 1 | Go | 2 |
| [kubernetes-sigs/gcp-filestore-csi-~](https://github.com/kubernetes-sigs/gcp-filestore-csi-driver/)  | 17.8K | 593 | 403 | 68 | Go | 1 |
| [keptn/go-utils](https://github.com/keptn/go-utils/)  | 17.6K | 895 | 565 | 7 | Go | 1 |
| [istio/tools](https://github.com/istio/tools/)  | 17.4K | 1.7K | 2.3K | 302 | Go | 1 |
| [openebs/dynamic-nfs-provisioner](https://github.com/openebs/dynamic-nfs-provisioner/)  | 17.3K | 112 | 109 | 103 | Go | 1 |
| [docker/compose](https://github.com/docker/compose/)  | 17.3K | 4.1K | 3.5K | 28.7K | Go | 1 |
| [kubevirt/ssp-operator](https://github.com/kubevirt/ssp-operator/)  | 17.2K | 629 | 501 | 12 | Go | 1 |
| [docker/compose-on-kubernetes](https://github.com/docker/compose-on-kubernetes/)  | 17.2K | 236 | 114 | 1.4K | Go | 1 |
| [ydb-platform/ydb-python-sdk](https://github.com/ydb-platform/ydb-python-sdk/)  | 17.7K | 594 | 132 | 54 | Python | 1 |
| [kubernetes-sigs/etcdadm](https://github.com/kubernetes-sigs/etcdadm/)  | 16.8K | 1.2K | 248 | 652 | Go | 1 |
| [kubernetes/api](https://github.com/kubernetes/api/)  | 16.8K | 6.7K | 10 | 514 | Go | 1 |
| [kubevirt/kubevirt-tekton-tasks](https://github.com/kubevirt/kubevirt-tekton-tasks/)  | 16.6K | 526 | 209 | 20 | Go | 1 |
| [ClickHouse/libcxx](https://github.com/ClickHouse/libcxx/)  | 16.5K | 38 | 14 | 0 | C++ | 1 |
| [open-telemetry/opentelemetry-ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16.4K | 763 | 919 | 375 | Ruby | 1 |
| [argoproj/argocon21](https://github.com/argoproj/argocon21/)  | 16.3K | 67 | 19 | 7 | JavaScript | 1 |
| [ClickHouse/clickhouse-odbc](https://github.com/ClickHouse/clickhouse-odbc/)  | 16.2K | 1.1K | 172 | 222 | C++ | 1 |
| [kubernetes-sigs/aws-ebs-csi-driver](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/)  | 16.2K | 1.6K | 882 | 721 | Go | 1 |
| [ClickHouse/boost](https://github.com/ClickHouse/boost/)  | 16.1K | 83 | 28 | 1 | C++ | 1 |
| [ClickHouse/antlr4-runtime](https://github.com/ClickHouse/antlr4-runtime/)  | 16.1K | 306 | 0 | 1 | C++ | 1 |
| [fluxcd/go-git-providers](https://github.com/fluxcd/go-git-providers/)  | 16.1K | 347 | 142 | 62 | Go | 1 |
| [litmuschaos/litmus-e2e](https://github.com/litmuschaos/litmus-e2e/)  | 15.9K | 231 | 360 | 15 | JavaScript | 1 |
| [kubernetes-sigs/boskos](https://github.com/kubernetes-sigs/boskos/)  | 15.9K | 923 | 129 | 93 | Go | 1 |
| [knative/test-infra](https://github.com/knative/test-infra/)  | 15.8K | 2.7K | 3.0K | 82 | Go | 1 |
| [knative/toolbox](https://github.com/knative/toolbox/)  | 15.8K | 2.7K | 3 | 0 | Go | 1 |
| [openebs/dynamic-localpv-provisioner](https://github.com/openebs/dynamic-localpv-provisioner/)  | 15.8K | 154 | 123 | 88 | Go | 1 |
| [nats-io/natscli](https://github.com/nats-io/natscli/)  | 15.8K | 1.6K | 543 | 280 | Go | 1 |
| [kubernetes-sigs/gcp-compute-persis~](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver/)  | 15.7K | 1.3K | 824 | 133 | Go | 1 |
| [prometheus/client_java](https://github.com/prometheus/client_java/)  | 15.6K | 578 | 490 | 1.9K | Java | 1 |
| [grafana/carbon-relay-ng](https://github.com/grafana/carbon-relay-ng/)  | 15.5K | 1.1K | 258 | 454 | Go | 1 |
| [cloudevents/sdk-java](https://github.com/cloudevents/sdk-java/)  | 15.4K | 755 | 370 | 315 | Java | 1 |
| [go-faster/yamlx](https://github.com/go-faster/yamlx/)  | 15.4K | 540 | 38 | 5 | Go | 1 |
| [nats-io/nats-jms-bridge](https://github.com/nats-io/nats-jms-bridge/)  | 15.2K | 398 | 90 | 11 | Java | 1 |
| [kubernetes-sigs/azurefile-csi-driv~](https://github.com/kubernetes-sigs/azurefile-csi-driver/)  | 15.2K | 2.4K | 886 | 113 | Go | 1 |
| [ClickHouse/boost-extra](https://github.com/ClickHouse/boost-extra/)  | 15.2K | 34 | 0 | 0 | C++ | 1 |
| [kubernetes-sigs/promo-tools](https://github.com/kubernetes-sigs/promo-tools/)  | 15.1K | 1.6K | 602 | 125 | Go | 1 |
| [etcd-io/raft](https://github.com/etcd-io/raft/)  | 15.2K | 1.1K | 19 | 139 | Go | 1 |
| [open-policy-agent/frameworks](https://github.com/open-policy-agent/frameworks/)  | 15.1K | 373 | 238 | 105 | Go | 1 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-openstack/)  | 15.1K | 1.4K | 977 | 220 | Go | 1 |
| [ClickHouse/ch-go](https://github.com/ClickHouse/ch-go/)  | 14.9K | 1.9K | 230 | 230 | Go | 1 |
| [kubernetes/cloud-provider-aws](https://github.com/kubernetes/cloud-provider-aws/)  | 14.9K | 975 | 401 | 294 | Go | 1 |
| [dapr/cli](https://github.com/dapr/cli/)  | 14.7K | 783 | 654 | 275 | Go | 1 |
| [theupdateframework/rust-tuf](https://github.com/theupdateframework/rust-tuf/)  | 14.6K | 756 | 221 | 155 | Rust | 1 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-ibmcloud/)  | 14.5K | 727 | 840 | 55 | Go | 1 |
| [openebs/jiva](https://github.com/openebs/jiva/)  | 14.5K | 804 | 337 | 133 | Go | 1 |
| [kubernetes-sigs/kind](https://github.com/kubernetes-sigs/kind/)  | 14.5K | 3.8K | 1.5K | 11.2K | Go | 1 |
| [kubevirt/cluster-network-addons-op~](https://github.com/kubevirt/cluster-network-addons-operator/)  | 14.4K | 1.1K | 1.3K | 57 | Go | 1 |
| [docker/libcompose](https://github.com/docker/libcompose/)  | 14.3K | 673 | 361 | 587 | Go | 1 |
| [docker/go-docker](https://github.com/docker/go-docker/)  | 14.3K | 7 | 4 | 179 | Go | 1 |
| [kubernetes/cloud-provider](https://github.com/kubernetes/cloud-provider/)  | 14.1K | 1.4K | 0 | 182 | Go | 1 |
| [nats-io/go-nats](https://github.com/nats-io/go-nats/)  | 14.1K | 937 | 0 | 20 | Go | 1 |
| [VictoriaMetrics/grafana-datasource](https://github.com/VictoriaMetrics/grafana-datasource/)  | 14.1K | 189 | 29 | 23 | TypeScript | 1 |
| [fluent/fluentd-kubernetes-daemonset](https://github.com/fluent/fluentd-kubernetes-daemonset/)  | 14.1K | 826 | 1.4K | 1.1K | Ruby | 1 |
| [cubeFS/shuttle](https://github.com/cubeFS/shuttle/)  | 13.9K | 59 | 8 | 127 | Java | 1 |
| [falcosecurity/falco](https://github.com/falcosecurity/falco/)  | 13.9K | 3.5K | 1.4K | 5.6K | C++ | 1 |
| [nats-io/nats.py](https://github.com/nats-io/nats.py/)  | 13.8K | 687 | 228 | 634 | Python | 1 |
| [kubernetes/component-base](https://github.com/kubernetes/component-base/)  | 13.7K | 1.2K | 0 | 84 | Go | 1 |
| [kubernetes-sigs/node-feature-disco~](https://github.com/kubernetes-sigs/node-feature-discovery/)  | 13.6K | 1.5K | 833 | 541 | Go | 1 |
| [open-policy-agent/setup-opa](https://github.com/open-policy-agent/setup-opa/)  | 13.5K | 31 | 17 | 33 | JavaScript | 1 |
| [docker/awesome-compose](https://github.com/docker/awesome-compose/)  | 13.5K | 267 | 255 | 21.8K | JavaScript | 1 |
| [prometheus/promlens](https://github.com/prometheus/promlens/)  | 13.5K | 78 | 81 | 307 | TypeScript | 1 |
| [grafana/har-to-k6](https://github.com/grafana/har-to-k6/)  | 13.4K | 627 | 84 | 77 | JavaScript | 1 |
| [istio/bots](https://github.com/istio/bots/)  | 13.4K | 638 | 652 | 10 | Go | 1 |
| [kyverno/policy-reporter](https://github.com/kyverno/policy-reporter/)  | 13.4K | 445 | 188 | 170 | Go | 1 |
| [knative/operator](https://github.com/knative/operator/)  | 13.2K | 855 | 1.1K | 135 | Go | 1 |
| [m3db/m3coordinator](https://github.com/m3db/m3coordinator/)  | 13.2K | 68 | 54 | 4 | Go | 1 |
| [theupdateframework/go-tuf](https://github.com/theupdateframework/go-tuf/)  | 13.2K | 450 | 306 | 531 | Go | 1 |
| [m3db/m3ninx](https://github.com/m3db/m3ninx/)  | 13.2K | 67 | 77 | 3 | Go | 1 |
| [envoyproxy/playground](https://github.com/envoyproxy/playground/)  | 13.8K | 192 | 168 | 7 | JavaScript | 1 |
| [AthenZ/athenz-client-sidecar](https://github.com/AthenZ/athenz-client-sidecar/)  | 13.6K | 72 | 4 | 2 | Go | 1 |
| [go-faster/hx](https://github.com/go-faster/hx/)  | 13.3K | 1.4K | 12 | 4 | Go | 1 |
| [prometheus/common](https://github.com/prometheus/common/)  | 12.9K | 497 | 362 | 224 | Go | 1 |
| [containerd/nydus-snapshotter](https://github.com/containerd/nydus-snapshotter/)  | 12.9K | 915 | 327 | 59 | Go | 1 |
| [open-telemetry/opentelemetry-php-c~](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12.7K | 322 | 112 | 22 | PHP | 1 |
| [fluent/ctraces](https://github.com/fluent/ctraces/)  | 12.7K | 150 | 35 | 0 | C | 1 |
| [carina-io/carina](https://github.com/carina-io/carina/)  | 12.7K | 712 | 115 | 588 | Go | 1 |
| [m3db/m3x](https://github.com/m3db/m3x/)  | 12.6K | 208 | 206 | 17 | Go | 1 |
| [etcd-io/bbolt](https://github.com/etcd-io/bbolt/)  | 12.5K | 1.1K | 220 | 6.2K | Go | 1 |
| [kubeedge/mappers-go](https://github.com/kubeedge/mappers-go/)  | 12.5K | 158 | 74 | 33 | Go | 1 |
| [argoproj/applicationset](https://github.com/argoproj/applicationset/)  | 12.5K | 195 | 255 | 566 | Go | 1 |
| [ydb-platform/ydb-rs-sdk](https://github.com/ydb-platform/ydb-rs-sdk/)  | 12.4K | 604 | 105 | 37 | Rust | 1 |
| [rook/nfs](https://github.com/rook/nfs/)  | 12.4K | 7.4K | 22 | 28 | Go | 1 |
| [jaegertracing/jaeger-client-java](https://github.com/jaegertracing/jaeger-client-java/)  | 12.3K | 528 | 504 | 485 | Java | 1 |
| [siderolabs/sidero](https://github.com/siderolabs/sidero/)  | 12.3K | 366 | 839 | 257 | Go | 1 |
| [docker/app](https://github.com/docker/app/)  | 12.3K | 1.5K | 650 | 1.5K | Go | 1 |
| [grafana/ksonnet](https://github.com/grafana/ksonnet/)  | 12.1K | 434 | 0 | 0 | Go | 1 |
| [dapr/python-sdk](https://github.com/dapr/python-sdk/)  | 12.1K | 327 | 351 | 173 | Python | 1 |
| [m3db/m3db-operator](https://github.com/m3db/m3db-operator/)  | 11.9K | 230 | 247 | 134 | Go | 1 |
| [openebs/openebsctl](https://github.com/openebs/openebsctl/)  | 11.9K | 112 | 105 | 27 | Go | 1 |
| [kata-containers/agent](https://github.com/kata-containers/agent/)  | 11.8K | 833 | 515 | 237 | Go | 1 |
| [openebs/website](https://github.com/openebs/website/)  | 11.8K | 1.1K | 342 | 11 | TypeScript | 1 |
| [rook/cassandra](https://github.com/rook/cassandra/)  | 11.8K | 7.4K | 13 | 5 | Go | 1 |
| [spiffe/go-spiffe](https://github.com/spiffe/go-spiffe/)  | 11.7K | 215 | 155 | 97 | Go | 1 |
| [nats-io/jwt](https://github.com/nats-io/jwt/)  | 11.7K | 288 | 165 | 63 | Go | 1 |
| [fluent/fluent-operator](https://github.com/fluent/fluent-operator/)  | 11.6K | 619 | 444 | 369 | Go | 1 |
| [grafana/jslib.k6.io](https://github.com/grafana/jslib.k6.io/)  | 11.5K | 189 | 71 | 32 | JavaScript | 1 |
| [open-telemetry/opentelemetry-erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11.5K | 1.1K | 355 | 253 | Erlang | 1 |
| [spiffe/java-spiffe](https://github.com/spiffe/java-spiffe/)  | 11.5K | 321 | 90 | 18 | Java | 1 |
| [prometheus/pushgateway](https://github.com/prometheus/pushgateway/)  | 11.5K | 669 | 267 | 2.5K | JavaScript | 1 |
| [jaegertracing/legacy-client-java](https://github.com/jaegertracing/legacy-client-java/)  | 11.4K | 292 | 14 | 5 | Java | 1 |
| [kubernetes-sigs/blob-csi-driver](https://github.com/kubernetes-sigs/blob-csi-driver/)  | 11.3K | 1.6K | 645 | 97 | Go | 1 |
| [openebs/zfs-localpv](https://github.com/openebs/zfs-localpv/)  | 11.3K | 275 | 320 | 261 | Go | 1 |
| [docker/getting-started](https://github.com/docker/getting-started/)  | 11.3K | 190 | 212 | 2.5K | JavaScript | 1 |
| [jaegertracing/jaeger-client-csharp](https://github.com/jaegertracing/jaeger-client-csharp/)  | 11.2K | 285 | 105 | 308 | C# | 1 |
| [rook/rook-client-python](https://github.com/rook/rook-client-python/)  | 11.2K | 57 | 3 | 1 | Python | 1 |
| [ClickHouse/clickhouse-presentations](https://github.com/ClickHouse/clickhouse-presentations/)  | 11.2K | 515 | 39 | 833 | JavaScript | 1 |
| [m3db/m3msg](https://github.com/m3db/m3msg/)  | 11.1K | 62 | 55 | 15 | Go | 1 |
| [docker/engine-api](https://github.com/docker/engine-api/)  | 11.1K | 9.1K | 327 | 267 | Go | 1 |
| [kubernetes-sigs/kubebuilder-declar~](https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/)  | 11.8K | 591 | 262 | 188 | Go | 1 |
| [cilium/proxy](https://github.com/cilium/proxy/)  | 10.9K | 567 | 117 | 82 | C++ | 1 |
| [kubernetes-sigs/controller-tools](https://github.com/kubernetes-sigs/controller-tools/)  | 10.9K | 773 | 455 | 549 | Go | 1 |
| [operator-framework/helm-operator-p~](https://github.com/operator-framework/helm-operator-plugins/)  | 10.8K | 194 | 141 | 34 | Go | 1 |
| [jaegertracing/jaeger-client-cpp](https://github.com/jaegertracing/jaeger-client-cpp/)  | 10.8K | 110 | 140 | 136 | C++ | 1 |
| [knative/networking](https://github.com/knative/networking/)  | 10.8K | 840 | 728 | 57 | Go | 1 |
| [envoyproxy/xds-relay](https://github.com/envoyproxy/xds-relay/)  | 10.8K | 127 | 162 | 128 | Go | 1 |
| [openebs/lvm-localpv](https://github.com/openebs/lvm-localpv/)  | 10.7K | 182 | 144 | 141 | Go | 1 |
| [open-telemetry/assign-reviewers-ac~](https://github.com/open-telemetry/assign-reviewers-action/)  | 10.7K | 5 | 6 | 6 | JavaScript | 1 |
| [fluxcd/kustomize-controller](https://github.com/fluxcd/kustomize-controller/)  | 10.6K | 1.3K | 559 | 204 | Go | 1 |
| [kubernetes/code-generator](https://github.com/kubernetes/code-generator/)  | 10.6K | 1.7K | 17 | 1.3K | Go | 1 |
| [grafana/jmeter-to-k6](https://github.com/grafana/jmeter-to-k6/)  | 10.6K | 347 | 24 | 57 | JavaScript | 1 |
| [kubernetes/utils](https://github.com/kubernetes/utils/)  | 10.6K | 896 | 235 | 256 | Go | 1 |
| [cloudevents/sdk-csharp](https://github.com/cloudevents/sdk-csharp/)  | 10.5K | 279 | 146 | 199 | C# | 1 |
| [kubernetes/node-problem-detector](https://github.com/kubernetes/node-problem-detector/)  | 10.5K | 753 | 434 | 2.2K | Go | 1 |
| [grpc/grpc-kotlin](https://github.com/grpc/grpc-kotlin/)  | 10.4K | 230 | 184 | 977 | Kotlin | 1 |
| [kubernetes/cli-runtime](https://github.com/kubernetes/cli-runtime/)  | 10.4K | 913 | 5 | 240 | Go | 1 |
| [in-toto/in-toto](https://github.com/in-toto/in-toto/)  | 10.4K | 1.8K | 407 | 681 | Python | 1 |
| [chaos-mesh/chaosd](https://github.com/chaos-mesh/chaosd/)  | 10.4K | 179 | 132 | 103 | Go | 1 |
| [grafana/tanka](https://github.com/grafana/tanka/)  | 10.3K | 449 | 469 | 1.8K | Go | 1 |
| [containerd/containerd.io](https://github.com/containerd/containerd.io/)  | 10.2K | 274 | 143 | 38 | JavaScript | 1 |
| [AthenZ/authorization-proxy](https://github.com/AthenZ/authorization-proxy/)  | 10.2K | 98 | 12 | 2 | Go | 1 |
| [istio/pkg](https://github.com/istio/pkg/)  | 10.1K | 859 | 763 | 47 | Go | 1 |
| [kubernetes-sigs/kube-batch](https://github.com/kubernetes-sigs/kube-batch/)  | 10.1K | 1.3K | 691 | 1.3K | Go | 1 |
| [istio/istio.io](https://github.com/istio/istio.io/)  | 10.1K | 8.3K | 10.9K | 681 | JavaScript | 1 |
| [istio/old_mixerclient_repo](https://github.com/istio/old_mixerclient_repo/)  | 10.4K | 228 | 409 | 15 | C++ | 1 |
| [kubernetes/kubeadm](https://github.com/kubernetes/kubeadm/)  | 9.9K | 1.2K | 520 | 3.3K | Go | 1 |
| [containerd/rust-extensions](https://github.com/containerd/rust-extensions/)  | 9.9K | 294 | 103 | 92 | Rust | 1 |
| [operator-framework/rukpak](https://github.com/operator-framework/rukpak/)  | 9.8K | 435 | 362 | 37 | Go | 1 |
| [etcd-io/jetcd](https://github.com/etcd-io/jetcd/)  | 9.8K | 1.1K | 742 | 970 | Java | 1 |
| [containerd/ttrpc-rust](https://github.com/containerd/ttrpc-rust/)  | 9.7K | 357 | 151 | 149 | Rust | 1 |
| [kubevirt/hostpath-provisioner-oper~](https://github.com/kubevirt/hostpath-provisioner-operator/)  | 9.7K | 315 | 285 | 25 | Go | 1 |
| [kubernetes-sigs/secrets-store-csi-~](https://github.com/kubernetes-sigs/secrets-store-csi-driver/)  | 9.6K | 1.1K | 747 | 939 | Go | 1 |
| [fluent/fluentd-ui](https://github.com/fluent/fluentd-ui/)  | 9.6K | 1.8K | 357 | 580 | Ruby | 1 |
| [istio/ztunnel](https://github.com/istio/ztunnel/)  | 9.6K | 277 | 302 | 90 | Rust | 1 |
| [go-faster/jx](https://github.com/go-faster/jx/)  | 9.6K | 1.4K | 65 | 80 | Go | 1 |
| [dapr/js-sdk](https://github.com/dapr/js-sdk/)  | 9.6K | 585 | 242 | 159 | TypeScript | 1 |
| [cilium/kube-apate](https://github.com/cilium/kube-apate/)  | 9.5K | 18 | 0 | 4 | Go | 1 |
| [grafana/cortex-tools](https://github.com/grafana/cortex-tools/)  | 9.5K | 270 | 179 | 133 | Go | 1 |
| [theupdateframework/python-tuf](https://github.com/theupdateframework/python-tuf/)  | 9.5K | 5.6K | 1.5K | 1.4K | Python | 1 |
| [nats-io/jsm.go](https://github.com/nats-io/jsm.go/)  | 9.5K | 795 | 391 | 102 | Go | 1 |
| [kubernetes-sigs/cloud-provider-hua~](https://github.com/kubernetes-sigs/cloud-provider-huaweicloud/)  | 9.4K | 415 | 171 | 25 | Go | 1 |
| [containerd/nri](https://github.com/containerd/nri/)  | 9.4K | 86 | 28 | 99 | Go | 1 |
| [prometheus/statsd_exporter](https://github.com/prometheus/statsd_exporter/)  | 9.4K | 847 | 310 | 823 | Go | 1 |
| [kubernetes-sigs/gateway-api](https://github.com/kubernetes-sigs/gateway-api/)  | 9.3K | 2.3K | 1.1K | 871 | Go | 1 |
| [grafana/worldmap-panel](https://github.com/grafana/worldmap-panel/)  | 9.3K | 232 | 74 | 294 | JavaScript | 1 |
| [openebs/cstor-csi](https://github.com/openebs/cstor-csi/)  | 9.3K | 157 | 172 | 23 | Go | 1 |
| [nats-io/nats.rb](https://github.com/nats-io/nats.rb/)  | 9.3K | 756 | 118 | 867 | Ruby | 1 |
| [kubernetes-sigs/cri-tools](https://github.com/kubernetes-sigs/cri-tools/)  | 9.2K | 1.5K | 794 | 1.2K | Go | 1 |
| [openebs/jiva-operator](https://github.com/openebs/jiva-operator/)  | 9.2K | 148 | 170 | 37 | Go | 1 |
| [nats-io/nats-rest-config-proxy](https://github.com/nats-io/nats-rest-config-proxy/)  | 9.2K | 213 | 49 | 21 | Go | 1 |
| [kubernetes/kompose](https://github.com/kubernetes/kompose/)  | 9.2K | 1.4K | 895 | 8.3K | Go | 1 |
| [dapr/php-sdk](https://github.com/dapr/php-sdk/)  | 9.1K | 255 | 80 | 59 | PHP | 1 |
| [ClickHouse/clickhouse-jdbc-bridge](https://github.com/ClickHouse/clickhouse-jdbc-bridge/)  | 9.1K | 127 | 85 | 137 | Java | 1 |
| [in-toto/in-toto-golang](https://github.com/in-toto/in-toto-golang/)  | 9.1K | 721 | 146 | 77 | Go | 1 |
| [crossplane/terrajet](https://github.com/crossplane/terrajet/)  | 9.1K | 463 | 136 | 292 | Go | 1 |
| [envoyproxy/go-control-plane](https://github.com/envoyproxy/go-control-plane/)  | 9.1K | 1.3K | 398 | 1.2K | Go | 1 |
| [grpc/test-infra](https://github.com/grpc/test-infra/)  | 9.1K | 484 | 346 | 64 | Go | 1 |
| [jaegertracing/jaeger-client-node](https://github.com/jaegertracing/jaeger-client-node/)  | 9.1K | 450 | 420 | 549 | JavaScript | 1 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-gcp/)  | 8.9K | 1.1K | 673 | 131 | Go | 1 |
| [volcano-sh/kubegene](https://github.com/volcano-sh/kubegene/)  | 8.9K | 101 | 50 | 184 | Go | 1 |
| [operator-framework/api](https://github.com/operator-framework/api/)  | 8.8K | 294 | 230 | 19 | Go | 1 |
| [longhorn/bot](https://github.com/longhorn/bot/)  | 8.8K | 25 | 20 | 0 | JavaScript | 1 |
| [kubevirt/kubevirt-velero-plugin](https://github.com/kubevirt/kubevirt-velero-plugin/)  | 8.8K | 123 | 126 | 18 | Go | 1 |
| [ClickHouse/libcxxabi](https://github.com/ClickHouse/libcxxabi/)  | 8.8K | 14 | 4 | 0 | C Header | 1 |
| [kubernetes-sigs/apiserver-network-~](https://github.com/kubernetes-sigs/apiserver-network-proxy/)  | 8.7K | 640 | 316 | 262 | Go | 1 |
| [kubernetes/pod-security-admission](https://github.com/kubernetes/pod-security-admission/)  | 8.7K | 465 | 0 | 77 | Go | 1 |
| [ClickHouse/graphouse](https://github.com/ClickHouse/graphouse/)  | 8.7K | 542 | 160 | 251 | Java | 1 |
| [ClickHouse/clickhouse-connect](https://github.com/ClickHouse/clickhouse-connect/)  | 8.7K | 229 | 84 | 86 | Python | 1 |
| [buildpacks/imgutil](https://github.com/buildpacks/imgutil/)  | 8.6K | 416 | 149 | 21 | Go | 1 |
| [knative/observability](https://github.com/knative/observability/)  | 8.6K | 19 | 28 | 25 | Go | 1 |
| [kubernetes-sigs/cli-experimental](https://github.com/kubernetes-sigs/cli-experimental/)  | 8.6K | 278 | 200 | 65 | Go | 1 |
| [spiffe/tornjak](https://github.com/spiffe/tornjak/)  | 8.5K | 206 | 75 | 55 | TypeScript | 1 |
| [containernetworking/cni](https://github.com/containernetworking/cni/)  | 8.5K | 1.9K | 603 | 4.7K | Go | 1 |
| [AthenZ/garm](https://github.com/AthenZ/garm/)  | 8.5K | 126 | 7 | 3 | Go | 1 |
| [vitessio/arewefastyet](https://github.com/vitessio/arewefastyet/)  | 8.5K | 1.3K | 222 | 33 | Go | 1 |
| [crossplane/release-test](https://github.com/crossplane/release-test/)  | 8.4K | 1.2K | 0 | 3 | Go | 1 |
| [kubernetes/gengo](https://github.com/kubernetes/gengo/)  | 8.4K | 455 | 182 | 485 | Go | 1 |
| [open-policy-agent/contrib](https://github.com/open-policy-agent/contrib/)  | 8.3K | 197 | 167 | 285 | Go | 1 |
| [kubernetes-sigs/cloud-provider-bai~](https://github.com/kubernetes-sigs/cloud-provider-baiducloud/)  | 8.3K | 213 | 71 | 38 | Go | 1 |
| [nats-io/nats-kafka](https://github.com/nats-io/nats-kafka/)  | 8.3K | 98 | 58 | 91 | Go | 1 |
| [projectcontour/contour-operator](https://github.com/projectcontour/contour-operator/)  | 8.3K | 353 | 400 | 44 | Go | 1 |
| [cert-manager/approver-policy](https://github.com/cert-manager/approver-policy/)  | 8.2K | 486 | 193 | 48 | Go | 1 |
| [litmuschaos/litmus-ui](https://github.com/litmuschaos/litmus-ui/)  | 8.2K | 171 | 109 | 5 | TypeScript | 1 |
| [knative/build](https://github.com/knative/build/)  | 8.2K | 372 | 452 | 575 | Go | 1 |
| [m3db/m3em](https://github.com/m3db/m3em/)  | 8.1K | 25 | 19 | 1 | Go | 1 |
| [spiffe/spiffe.io](https://github.com/spiffe/spiffe.io/)  | 7.9K | 791 | 216 | 20 | JavaScript | 1 |
| [operator-framework/audit](https://github.com/operator-framework/audit/)  | 7.9K | 247 | 107 | 8 | Go | 1 |
| [openebs/mayastor-extensions](https://github.com/openebs/mayastor-extensions/)  | 7.8K | 220 | 125 | 10 | Rust | 1 |
| [envoyproxy/ratelimit](https://github.com/envoyproxy/ratelimit/)  | 7.8K | 150 | 235 | 1.7K | Go | 1 |
| [open-telemetry/opentelemetry-go-bu~](https://github.com/open-telemetry/opentelemetry-go-build-tools/)  | 7.7K | 309 | 242 | 21 | Go | 1 |
| [openebs/device-localpv](https://github.com/openebs/device-localpv/)  | 7.7K | 55 | 49 | 15 | Go | 1 |
| [nats-io/nats-pure.rb](https://github.com/nats-io/nats-pure.rb/)  | 7.6K | 304 | 60 | 95 | Ruby | 1 |
| [kubernetes-sigs/krew](https://github.com/kubernetes-sigs/krew/)  | 7.6K | 478 | 479 | 5.3K | Go | 1 |
| [kedacore/http-add-on](https://github.com/kedacore/http-add-on/)  | 7.6K | 262 | 413 | 176 | Go | 1 |
| [prometheus/client_python](https://github.com/prometheus/client_python/)  | 7.6K | 512 | 389 | 3.2K | Python | 1 |
| [grafana/grafana-polystat-panel](https://github.com/grafana/grafana-polystat-panel/)  | 7.5K | 199 | 124 | 74 | TypeScript | 1 |
| [fluent/fluent-plugin-opensearch](https://github.com/fluent/fluent-plugin-opensearch/)  | 7.5K | 1.3K | 29 | 39 | Ruby | 1 |
| [kubernetes/kube-aggregator](https://github.com/kubernetes/kube-aggregator/)  | 7.4K | 2.8K | 11 | 223 | Go | 1 |
| [grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client/)  | 7.4K | 402 | 136 | 71 | Go | 1 |
| [notaryproject/notation-go](https://github.com/notaryproject/notation-go/)  | 7.3K | 124 | 198 | 18 | Go | 1 |
| [dapr/go-sdk](https://github.com/dapr/go-sdk/)  | 7.3K | 239 | 257 | 387 | Go | 1 |
| [kubernetes-sigs/aws-iam-authentica~](https://github.com/kubernetes-sigs/aws-iam-authenticator/)  | 7.3K | 515 | 313 | 1.9K | Go | 1 |
| [kubernetes-sigs/aws-efs-csi-driver](https://github.com/kubernetes-sigs/aws-efs-csi-driver/)  | 7.3K | 689 | 585 | 568 | Go | 1 |
| [openebs/upgrade](https://github.com/openebs/upgrade/)  | 7.2K | 122 | 135 | 10 | Go | 1 |
| [grafana/postman-to-k6](https://github.com/grafana/postman-to-k6/)  | 7.2K | 607 | 58 | 268 | JavaScript | 1 |
| [kubernetes-sigs/instrumentation-to~](https://github.com/kubernetes-sigs/instrumentation-tools/)  | 7.1K | 92 | 6 | 25 | Go | 1 |
| [kubevirt/client-python](https://github.com/kubevirt/client-python/)  | 7.1K | 358 | 21 | 24 | Python | 1 |
| [docker/kitematic](https://github.com/docker/kitematic/)  | 7.1K | 2.3K | 534 | 12.2K | JavaScript | 1 |
| [grpc/grpc-web](https://github.com/grpc/grpc-web/)  | 7.7K | 888 | 622 | 7.3K | JavaScript | 1 |
| [fluxcd/notification-controller](https://github.com/fluxcd/notification-controller/)  | 7.5K | 829 | 349 | 119 | Go | 1 |
| [kubernetes/dns](https://github.com/kubernetes/dns/)  | 7.3K | 710 | 318 | 803 | Go | 1 |
| [grafana/opcua-datasource](https://github.com/grafana/opcua-datasource/)  | 7.3K | 370 | 44 | 45 | TypeScript | 1 |
| [prometheus/mysqld_exporter](https://github.com/prometheus/mysqld_exporter/)  | 6.9K | 605 | 412 | 1.6K | Go | 1 |
| [notaryproject/notation](https://github.com/notaryproject/notation/)  | 6.9K | 192 | 343 | 144 | Go | 1 |
| [openebs/libcstor](https://github.com/openebs/libcstor/)  | 6.9K | 77 | 89 | 14 | C | 1 |
| [kubeedge/edgemesh](https://github.com/kubeedge/edgemesh/)  | 6.9K | 534 | 265 | 182 | Go | 1 |
| [open-telemetry/opentelemetry-demo](https://github.com/open-telemetry/opentelemetry-demo/)  | 6.8K | 422 | 551 | 691 | TypeScript | 1 |
| [cilium/hubble](https://github.com/cilium/hubble/)  | 6.8K | 796 | 757 | 2.5K | Go | 1 |
| [containerd/cgroups](https://github.com/containerd/cgroups/)  | 6.8K | 517 | 218 | 928 | Go | 1 |
| [thanos-io/objstore](https://github.com/thanos-io/objstore/)  | 6.8K | 209 | 34 | 43 | Go | 1 |
| [grafana/attic](https://github.com/grafana/attic/)  | 6.8K | 426 | 0 | 1 | JavaScript | 1 |
| [envoyproxy/envoy-build-tools](https://github.com/envoyproxy/envoy-build-tools/)  | 6.8K | 320 | 182 | 36 | Bazel | 1 |
| [kubernetes-sigs/reference-docs](https://github.com/kubernetes-sigs/reference-docs/)  | 6.8K | 492 | 233 | 65 | Go | 1 |
| [nats-io/nats-operator](https://github.com/nats-io/nats-operator/)  | 6.8K | 588 | 206 | 555 | Go | 1 |
| [kubevirt/terraform-provider-kubevi~](https://github.com/kubevirt/terraform-provider-kubevirt/)  | 6.7K | 101 | 26 | 16 | Go | 1 |
| [grafana/kairosdb-datasource](https://github.com/grafana/kairosdb-datasource/)  | 6.7K | 110 | 43 | 31 | TypeScript | 1 |
| [cilium/image-tools](https://github.com/cilium/image-tools/)  | 6.7K | 239 | 192 | 11 | Perl | 1 |
| [argoproj/notifications-engine](https://github.com/argoproj/notifications-engine/)  | 6.7K | 81 | 108 | 180 | Go | 1 |
| [in-toto/in-toto-rs](https://github.com/in-toto/in-toto-rs/)  | 6.6K | 161 | 37 | 28 | Rust | 1 |
| [dapr/kit](https://github.com/dapr/kit/)  | 6.6K | 30 | 35 | 10 | Go | 1 |
| [helm/chartmuseum](https://github.com/helm/chartmuseum/)  | 6.6K | 613 | 333 | 3.1K | Go | 1 |
| [kubevirt/kvm-info-nfd-plugin](https://github.com/kubevirt/kvm-info-nfd-plugin/)  | 6.6K | 62 | 13 | 1 | C Header | 1 |
| [ydb-platform/ydb-nodejs-sdk](https://github.com/ydb-platform/ydb-nodejs-sdk/)  | 6.6K | 367 | 185 | 46 | TypeScript | 1 |
| [etcd-io/dbtester](https://github.com/etcd-io/dbtester/)  | 6.5K | 1.2K | 288 | 250 | Go | 1 |
| [argoproj/argo-cd-ui](https://github.com/argoproj/argo-cd-ui/)  | 6.5K | 385 | 131 | 14 | TypeScript | 1 |
| [kubernetes-sigs/apiserver-builder-~](https://github.com/kubernetes-sigs/apiserver-builder-alpha/)  | 6.5K | 1.7K | 434 | 717 | Go | 1 |
| [ydb-platform/ydb-go-examples](https://github.com/ydb-platform/ydb-go-examples/)  | 6.5K | 176 | 16 | 12 | Go | 1 |
| [envoyproxy/envoy-perf](https://github.com/envoyproxy/envoy-perf/)  | 6.5K | 142 | 160 | 107 | Python | 1 |
| [crossplane/gitlab-controller](https://github.com/crossplane/gitlab-controller/)  | 6.4K | 45 | 18 | 12 | Go | 1 |
| [AthenZ/terraform-provider-athenz](https://github.com/AthenZ/terraform-provider-athenz/)  | 6.4K | 94 | 56 | 3 | Go | 1 |
| [gotd/botapi](https://github.com/gotd/botapi/)  | 6.4K | 653 | 313 | 18 | Go | 1 |
| [nats-io/nats.ts](https://github.com/nats-io/nats.ts/)  | 6.3K | 533 | 86 | 180 | TypeScript | 1 |
| [kubernetes/csi-translation-lib](https://github.com/kubernetes/csi-translation-lib/)  | 6.3K | 658 | 0 | 10 | Go | 1 |
| [open-telemetry/opamp-go](https://github.com/open-telemetry/opamp-go/)  | 6.3K | 92 | 110 | 60 | Go | 1 |
| [kubernetes-sigs/kubetest2](https://github.com/kubernetes-sigs/kubetest2/)  | 6.3K | 535 | 177 | 259 | Go | 1 |
| [jaegertracing/jaeger-client-python](https://github.com/jaegertracing/jaeger-client-python/)  | 6.2K | 264 | 236 | 413 | Python | 1 |
| [containerd/continuity](https://github.com/containerd/continuity/)  | 6.2K | 372 | 157 | 126 | Go | 1 |
| [kubernetes-sigs/sig-storage-local-~](https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner/)  | 6.1K | 623 | 213 | 866 | Go | 1 |
| [kubernetes-sigs/slack-infra](https://github.com/kubernetes-sigs/slack-infra/)  | 6.1K | 108 | 44 | 85 | Go | 1 |
| [nats-io/java-nats-examples](https://github.com/nats-io/java-nats-examples/)  | 6.5K | 166 | 7 | 20 | Java | 1 |
| [kubevirt/hostpath-provisioner](https://github.com/kubevirt/hostpath-provisioner/)  | 6.5K | 149 | 164 | 51 | Go | 1 |
| [cubeFS/cubefs-for-android](https://github.com/cubeFS/cubefs-for-android/)  | 6.4K | 4 | 0 | 17 | Go | 1 |
| [notaryproject/notation-core-go](https://github.com/notaryproject/notation-core-go/)  | 6.4K | 42 | 99 | 5 | Go | 1 |
| [cilium/docsearch-scraper-webhook](https://github.com/cilium/docsearch-scraper-webhook/)  | 5.9K | 38 | 32 | 3 | Python | 1 |
| [kubeedge/examples](https://github.com/kubeedge/examples/)  | 5.9K | 132 | 64 | 201 | Go | 1 |
| [open-policy-agent/conftest](https://github.com/open-policy-agent/conftest/)  | 5.9K | 868 | 505 | 2.5K | Go | 1 |
| [siderolabs/kres](https://github.com/siderolabs/kres/)  | 5.9K | 152 | 199 | 22 | Go | 1 |
| [nats-io/stan.java](https://github.com/nats-io/stan.java/)  | 5.9K | 361 | 64 | 94 | Java | 1 |
| [ClickHouse/bzip2](https://github.com/ClickHouse/bzip2/)  | 5.9K | 141 | 0 | 1 | C | 1 |
| [kubevirt/client-go](https://github.com/kubevirt/client-go/)  | 5.9K | 301 | 9 | 25 | Go | 1 |
| [prometheus/blackbox_exporter](https://github.com/prometheus/blackbox_exporter/)  | 5.9K | 455 | 436 | 3.4K | Go | 1 |
| [grafana/devtools](https://github.com/grafana/devtools/)  | 5.9K | 139 | 22 | 9 | Go | 1 |
| [litmuschaos/charthub.litmuschaos.io](https://github.com/litmuschaos/charthub.litmuschaos.io/)  | 5.9K | 225 | 317 | 14 | TypeScript | 0 |
| [theupdateframework/tuf-js](https://github.com/theupdateframework/tuf-js/)  | 5.9K | 301 | 226 | 59 | TypeScript | 0 |
| [docker/go](https://github.com/docker/go/)  | 5.8K | 24 | 8 | 17 | Go | 0 |
| [kata-containers/govmm](https://github.com/kata-containers/govmm/)  | 5.8K | 402 | 151 | 304 | Go | 0 |
| [chaos-mesh/k8s_dns_chaos](https://github.com/chaos-mesh/k8s_dns_chaos/)  | 5.8K | 35 | 17 | 10 | Go | 0 |
| [docker/volumes-backup-extension](https://github.com/docker/volumes-backup-extension/)  | 5.8K | 228 | 76 | 47 | Go | 0 |
| [istio/api](https://github.com/istio/api/)  | 5.8K | 1.6K | 2.4K | 404 | Protocol Buffers | 0 |
| [kubernetes/website](https://github.com/kubernetes/website/)  | 5.7K | 39.7K | 29.1K | 3.6K | JavaScript | 0 |
| [docker/metadata-action](https://github.com/docker/metadata-action/)  | 5.7K | 318 | 157 | 563 | TypeScript | 0 |
| [dapr/cpp-sdk](https://github.com/dapr/cpp-sdk/)  | 5.6K | 45 | 27 | 29 | C Header | 0 |
| [cloudevents/sdk-rust](https://github.com/cloudevents/sdk-rust/)  | 5.6K | 194 | 148 | 125 | Rust | 0 |
| [grafana/azure-monitor-datasource](https://github.com/grafana/azure-monitor-datasource/)  | 5.6K | 170 | 18 | 91 | TypeScript | 0 |
| [fluxcd/helm-operator](https://github.com/fluxcd/helm-operator/)  | 5.6K | 4.6K | 351 | 656 | Go | 0 |
| [prometheus/compliance](https://github.com/prometheus/compliance/)  | 5.6K | 147 | 79 | 105 | Go | 0 |
| [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/)  | 5.5K | 989 | 602 | 4.5K | Go | 0 |
| [kubernetes/component-helpers](https://github.com/kubernetes/component-helpers/)  | 5.5K | 329 | 0 | 11 | Go | 0 |
| [m3db/m3storage](https://github.com/m3db/m3storage/)  | 5.5K | 38 | 15 | 3 | Go | 0 |
| [fluxcd/helm-controller](https://github.com/fluxcd/helm-controller/)  | 5.4K | 897 | 412 | 293 | Go | 0 |
| [fluent/fluentd-docs](https://github.com/fluent/fluentd-docs/)  | 5.4K | 2.4K | 458 | 49 | JavaScript | 0 |
| [longhorn/backing-image-manager](https://github.com/longhorn/backing-image-manager/)  | 5.4K | 126 | 105 | 4 | Go | 0 |
| [spiffe/spire-controller-manager](https://github.com/spiffe/spire-controller-manager/)  | 5.4K | 84 | 92 | 29 | Go | 0 |
| [openebs/api](https://github.com/openebs/api/)  | 5.3K | 88 | 99 | 7 | Go | 0 |
| [kubernetes-sigs/cluster-addons](https://github.com/kubernetes-sigs/cluster-addons/)  | 5.3K | 261 | 89 | 147 | Go | 0 |
| [nats-io/nats.py2](https://github.com/nats-io/nats.py2/)  | 5.3K | 320 | 50 | 62 | Python | 0 |
| [prometheus/codemirror-promql](https://github.com/prometheus/codemirror-promql/)  | 5.3K | 516 | 151 | 35 | TypeScript | 0 |
| [VictoriaMetrics/metricsql](https://github.com/VictoriaMetrics/metricsql/)  | 5.2K | 102 | 6 | 136 | Go | 0 |
| [cert-manager/release](https://github.com/cert-manager/release/)  | 5.2K | 246 | 108 | 4 | Go | 0 |
| [litmuschaos/website-litmuschaos](https://github.com/litmuschaos/website-litmuschaos/)  | 5.2K | 133 | 165 | 9 | TypeScript | 0 |
| [kubernetes-sigs/aws-fsx-csi-driver](https://github.com/kubernetes-sigs/aws-fsx-csi-driver/)  | 5.2K | 409 | 204 | 103 | Go | 0 |
| [nats-io/nats-account-server](https://github.com/nats-io/nats-account-server/)  | 5.2K | 209 | 60 | 69 | Go | 0 |
| [grafana/kubernetes-app](https://github.com/grafana/kubernetes-app/)  | 5.2K | 199 | 23 | 397 | TypeScript | 0 |
| [litmuschaos/chaos-runner](https://github.com/litmuschaos/chaos-runner/)  | 5.2K | 136 | 181 | 19 | Go | 0 |
| [helm/helm-classic](https://github.com/helm/helm-classic/)  | 5.2K | 574 | 274 | 578 | Go | 0 |
| [goharbor/terraform-provider-harbor](https://github.com/goharbor/terraform-provider-harbor/)  | 5.1K | 520 | 185 | 87 | Go | 0 |
| [k3s-io/kine](https://github.com/k3s-io/kine/)  | 5.1K | 139 | 102 | 934 | Go | 0 |
| [etcd-io/zetcd](https://github.com/etcd-io/zetcd/)  | 5.1K | 135 | 58 | 1.6K | Go | 0 |
| [fluxcd/image-reflector-controller](https://github.com/fluxcd/image-reflector-controller/)  | 5.1K | 583 | 253 | 79 | Go | 0 |
| [fluent/chunkio](https://github.com/fluent/chunkio/)  | 5.9K | 371 | 88 | 16 | C | 0 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-digitalocean/)  | 5.8K | 399 | 383 | 92 | Go | 0 |
| [longhorn/backupstore](https://github.com/longhorn/backupstore/)  | 5.8K | 188 | 107 | 3 | Go | 0 |
| [coredns/policy](https://github.com/coredns/policy/)  | 5.1K | 77 | 30 | 46 | Go | 0 |
| [ydb-platform/terraform-provider-ydb](https://github.com/ydb-platform/terraform-provider-ydb/)  | 4.9K | 70 | 10 | 1 | Go | 0 |
| [m3db/m3ctl](https://github.com/m3db/m3ctl/)  | 4.9K | 82 | 64 | 14 | Go | 0 |
| [fluxcd/webui](https://github.com/fluxcd/webui/)  | 4.9K | 162 | 38 | 142 | TypeScript | 0 |
| [grafana/gel-app](https://github.com/grafana/gel-app/)  | 4.9K | 119 | 47 | 2 | Go | 0 |
| [kubeedge/ianvs](https://github.com/kubeedge/ianvs/)  | 4.8K | 155 | 53 | 76 | Python | 0 |
| [goharbor/harbor-cluster-operator](https://github.com/goharbor/harbor-cluster-operator/)  | 4.8K | 390 | 114 | 28 | Go | 0 |
| [fluent/cfl](https://github.com/fluent/cfl/)  | 4.8K | 88 | 25 | 0 | C Header | 0 |
| [dapr/setup-dapr](https://github.com/dapr/setup-dapr/)  | 4.8K | 14 | 114 | 7 | JavaScript | 0 |
| [kyverno/policy-reporter-kyverno-pl~](https://github.com/kyverno/policy-reporter-kyverno-plugin/)  | 4.7K | 56 | 20 | 8 | Go | 0 |
| [prometheus/snmp_exporter](https://github.com/prometheus/snmp_exporter/)  | 4.7K | 553 | 416 | 1.1K | Go | 0 |
| [nats-io/nats-replicator](https://github.com/nats-io/nats-replicator/)  | 4.7K | 44 | 7 | 18 | Go | 0 |
| [longhorn/longhorn-instance-manager](https://github.com/longhorn/longhorn-instance-manager/)  | 4.7K | 366 | 195 | 8 | Go | 0 |
| [helm/monocular](https://github.com/helm/monocular/)  | 4.7K | 399 | 362 | 1.4K | Go | 0 |
| [kata-containers/tests](https://github.com/kata-containers/tests/)  | 4.7K | 4.3K | 2.7K | 135 | Go | 0 |
| [gotd/contrib](https://github.com/gotd/contrib/)  | 4.7K | 568 | 317 | 10 | Go | 0 |
| [containerssh/auditlog](https://github.com/containerssh/auditlog/)  | 4.7K | 118 | 476 | 1 | Go | 0 |
| [knative/client-contrib](https://github.com/knative/client-contrib/)  | 4.6K | 36 | 45 | 11 | Go | 0 |
| [argoproj/argo-ui](https://github.com/argoproj/argo-ui/)  | 4.6K | 409 | 302 | 176 | TypeScript | 0 |
| [kubernetes-sigs/image-builder](https://github.com/kubernetes-sigs/image-builder/)  | 4.6K | 2.5K | 815 | 248 | Go | 0 |
| [grafana/grafana-sensu-app](https://github.com/grafana/grafana-sensu-app/)  | 4.6K | 16 | 45 | 8 | TypeScript | 0 |
| [chaos-mesh/toda](https://github.com/chaos-mesh/toda/)  | 4.6K | 233 | 19 | 41 | Rust | 0 |
| [containerd/runwasi](https://github.com/containerd/runwasi/)  | 4.5K | 186 | 62 | 541 | Rust | 0 |
| [cloudevents/sdk-python](https://github.com/cloudevents/sdk-python/)  | 4.5K | 83 | 123 | 188 | Python | 0 |
| [nats-io/stan.net](https://github.com/nats-io/stan.net/)  | 4.5K | 190 | 114 | 134 | C# | 0 |
| [istio/old_auth_repo](https://github.com/istio/old_auth_repo/)  | 4.5K | 166 | 230 | 73 | Go | 0 |
| [openebs/velero-plugin](https://github.com/openebs/velero-plugin/)  | 4.5K | 119 | 137 | 43 | Go | 0 |
| [cert-manager/trust-manager](https://github.com/cert-manager/trust-manager/)  | 4.5K | 263 | 82 | 101 | Go | 0 |


### CNCF Projects
| Project | SLOC | Commits | PRs | Stars | Language | Effort |
|---------|------|---------|-----|-------|----------|--------|
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1K | 190.5K | 243.1K | Go | 286 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 1.9M | 117.4K | 71.3K | 98.2K | Go | 166 |
| [open-telemetry](https://github.com/open-telemetry)  | 1.7M | 61.8K | 61.4K | 28.6K | Go | 148 |
| [grpc](https://github.com/grpc)  | 1.6M | 74.1K | 38.7K | 86.8K | C++ | 140 |
| [envoyproxy](https://github.com/envoyproxy)  | 1.5M | 46.3K | 24.6K | 28.6K | C++ | 131 |
| [nats-io](https://github.com/nats-io)  | 802.4K | 36.2K | 12.9K | 31.4K | Go | 67 |
| [vitessio](https://github.com/vitessio)  | 671.3K | 36.7K | 11.6K | 15.9K | Go | 56 |
| [fluent](https://github.com/fluent)  | 639.2K | 38.4K | 11.9K | 24.2K | C | 53 |
| [kubevirt](https://github.com/kubevirt)  | 631.7K | 60.5K | 22.5K | 5.4K | Go | 53 |
| [cubeFS](https://github.com/cubeFS)  | 591.2K | 2.7K | 1.4K | 3.3K | Go | 49 |
| [istio](https://github.com/istio)  | 583.5K | 47.5K | 62.3K | 38.4K | Go | 49 |
| [argoproj](https://github.com/argoproj)  | 517.7K | 18.1K | 16.8K | 34.2K | Go | 43 |
| [knative](https://github.com/knative)  | 432.8K | 30.5K | 32.2K | 12.3K | Go | 36 |
| [AthenZ](https://github.com/AthenZ)  | 405.5K | 3.4K | 2.4K | 762 | Java | 34 |
| [backstage](https://github.com/backstage)  | 402.2K | 40.4K | 12.9K | 21.7K | TypeScript | 34 |
| [prometheus](https://github.com/prometheus)  | 377.3K | 30.8K | 19.1K | 89.5K | Go | 31 |
| [containerd](https://github.com/containerd)  | 353.3K | 23.2K | 11.6K | 24.7K | Go | 29 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.1K | 8.5K | 7.1K | 15.4K | Go | 29 |
| [dapr](https://github.com/dapr)  | 335.5K | 17.4K | 10.7K | 26.3K | Go | 28 |
| [kyverno](https://github.com/kyverno)  | 280.2K | 8.7K | 5.1K | 4.1K | Go | 23 |
| [operator-framework](https://github.com/operator-framework)  | 265.4K | 18.9K | 14.3K | 13.3K | Go | 22 |
| [keptn](https://github.com/keptn)  | 256.7K | 11.7K | 8.3K | 1.8K | Go | 21 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.4K | 8.7K | 7.6K | 23.6K | Go | 21 |
| [goharbor](https://github.com/goharbor)  | 242.4K | 16.9K | 10.3K | 21.1K | Go | 20 |
| [linkerd](https://github.com/linkerd)  | 237.9K | 13.1K | 13.3K | 17.9K | Rust | 20 |
| [spiffe](https://github.com/spiffe)  | 206.9K | 8.6K | 3.9K | 2.9K | Go | 17 |
| [litmuschaos](https://github.com/litmuschaos)  | 203.2K | 7.4K | 7.9K | 4.5K | TypeScript | 17 |
| [fluxcd](https://github.com/fluxcd)  | 200.2K | 27.6K | 9.4K | 19.8K | Go | 17 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4K | 4.4K | 4.5K | 4.8K | Python | 16 |
| [longhorn](https://github.com/longhorn)  | 193.6K | 9.9K | 6.5K | 5.1K | Go | 16 |
| [etcd-io](https://github.com/etcd-io)  | 185.2K | 26.3K | 11.2K | 52.2K | Go | 15 |
| [cert-manager](https://github.com/cert-manager)  | 145.5K | 13.1K | 4.8K | 10.8K | Go | 12 |
| [kubeedge](https://github.com/kubeedge)  | 141.3K | 7.3K | 3.9K | 6.7K | Go | 12 |
| [rook](https://github.com/rook)  | 132.1K | 24.7K | 7.1K | 11.1K | Go | 11 |
| [projectcontour](https://github.com/projectcontour)  | 130.7K | 4.9K | 3.8K | 4.1K | Go | 11 |
| [crossplane](https://github.com/crossplane)  | 127.7K | 11.2K | 3.5K | 7.7K | Go | 11 |
| [thanos-io](https://github.com/thanos-io)  | 123.5K | 4.7K | 3.9K | 12.4K | Go | 10 |
| [cortexproject](https://github.com/cortexproject)  | 120.3K | 4.8K | 3.7K | 5.1K | Go | 10 |
| [buildpacks](https://github.com/buildpacks)  | 116.6K | 16.7K | 3.5K | 2.7K | Go | 10 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 106.7K | 3.2K | 3.1K | 5.8K | Go | 9 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 105.6K | 17.9K | 3.2K | 4.3K | Python | 9 |
| [kedacore](https://github.com/kedacore)  | 105.4K | 3.8K | 4.2K | 6.7K | Go | 9 |
| [karmada-io](https://github.com/karmada-io)  | 98.8K | 4.1K | 2.6K | 3.1K | Go | 8 |
| [in-toto](https://github.com/in-toto)  | 92.2K | 9.4K | 1.2K | 1.6K | Python | 8 |
| [cloudevents](https://github.com/cloudevents)  | 79.5K | 3.9K | 2.6K | 5.8K | Go | 7 |
| [coredns](https://github.com/coredns)  | 76.2K | 5.5K | 4.6K | 11.5K | Go | 6 |
| [helm](https://github.com/helm)  | 71.5K | 28.9K | 26.8K | 48.7K | Go | 6 |
| [cri-o](https://github.com/cri-o)  | 71.2K | 9.1K | 5.7K | 4.5K | Go | 6 |
| [volcano-sh](https://github.com/volcano-sh)  | 68.5K | 5.4K | 2.3K | 3.1K | Go | 6 |
| [notaryproject](https://github.com/notaryproject)  | 64.4K | 3.6K | 1.9K | 3.3K | Go | 5 |
| [containerssh](https://github.com/containerssh)  | 62.1K | 1.7K | 2.2K | 2.2K | Go | 5 |
| [theupdateframework](https://github.com/theupdateframework)  | 45.3K | 8.6K | 2.5K | 2.6K | Rust | 4 |
| [k3s-io](https://github.com/k3s-io)  | 38.5K | 3.4K | 3.2K | 25.3K | Go | 3 |
| [containernetworking](https://github.com/containernetworking)  | 34.7K | 2.8K | 1.1K | 6.4K | Go | 3 |
| [carina-io](https://github.com/carina-io)  | 12.8K | 749 | 118 | 588 | Go | 1 |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3K | 251 | 117 | 2.1K | Go | 0 |

