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

Languages: BASH, Bazel, C, C Header, C#, C++, Cython, Dart, Erlang, Go, Go Template, Haskell, Java, JavaScript, Kotlin, Lua, Objective C, Objective C++, PHP, Perl, Protocol Buffers, Python, Ruby, Rust, Scala, Swift, TypeScript, Zig.

### Effort

Effort is **rough estimate** how much developer-years it would take to write some project from scratch in proprietary,
enterprise environment (i.e. in a big company).

Current model is very simple: we estimate that average developer writes 12k SLOC per year.

TODO:
- Take into account COCOMO (not suitable for FOSS?)
- Estimate on commits and PRs (can vary a lot)
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
| [**CNCF**](https://www.cncf.io/)  | 22.4M | 1.4M | 824.5k | 1.2M | Go | 1.9k |
| [torvalds](https://github.com/torvalds)  | 19.1M | 1.2M | 761.0 | 148.2k | C | 1.6k |
| [**K8s**](https://kubernetes.io/)  | 5.4M | 498.5k | 261.5k | 341.4k | Go | 452.7 |
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1k | 190.5k | 243.1k | Go | 286.2 |
| [tensorflow](https://github.com/tensorflow)  | 3.4M | 145.1k | 22.8k | 172.2k | C++ | 279.2 |
| [ydb-platform](https://github.com/ydb-platform)  | 3.1M | 18.6k | 1.6k | 3.1k | C++ | 261.5 |
| [elastic](https://github.com/elastic)  | 3.1M | 94.5k | 96.8k | 88.0k | Java | 261.2 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.0M | 117.4k | 71.0k | 98.3k | Go | 166.5 |
| [cockroachdb](https://github.com/cockroachdb)  | 2.0M | 100.8k | 59.3k | 33.7k | Go | 163.1 |
| [open-telemetry](https://github.com/open-telemetry)  | 1.8M | 61.9k | 61.0k | 28.6k | Go | 148.4 |
| [grpc](https://github.com/grpc)  | 1.7M | 74.1k | 38.8k | 86.8k | C++ | 140.3 |
| [ClickHouse](https://github.com/ClickHouse)  | 1.7M | 119.2k | 34.8k | 33.9k | C++ | 140.0 |
| [grafana](https://github.com/grafana)  | 1.7M | 76.6k | 50.2k | 103.2k | Go | 139.9 |
| [envoyproxy](https://github.com/envoyproxy)  | 1.6M | 46.3k | 24.1k | 28.1k | C++ | 131.1 |
| [rust-lang](https://github.com/rust-lang)  | 1.5M | 220.1k | 61.4k | 79.1k | Rust | 127.1 |
| [python](https://github.com/python)  | 1.4M | 116.3k | 39.6k | 51.3k | Python | 114.4 |
| [docker](https://github.com/docker)  | 1.3M | 173.8k | 30.6k | 146.4k | Go | 105.5 |
| [golang](https://github.com/golang)  | 1.1M | 55.8k | 2.9k | 109.4k | Go | 88.4 |
| [ziglang](https://github.com/ziglang)  | 901.1k | 23.1k | 6.8k | 20.5k | Zig | 75.1 |
| [nats-io](https://github.com/nats-io)  | 802.5k | 36.2k | 12.9k | 31.0k | Go | 66.9 |
| [m3db](https://github.com/m3db)  | 736.8k | 10.4k | 5.1k | 4.6k | Go | 61.4 |
| [cilium](https://github.com/cilium)  | 681.6k | 31.0k | 21.9k | 23.6k | Go | 56.8 |
| [vitessio](https://github.com/vitessio)  | 671.3k | 36.7k | 11.0k | 15.9k | Go | 56.0 |
| [fluent](https://github.com/fluent)  | 639.2k | 38.0k | 12.0k | 24.2k | C | 53.3 |
| [kubevirt](https://github.com/kubevirt)  | 631.8k | 60.5k | 22.6k | 5.5k | Go | 52.6 |
| [cubeFS](https://github.com/cubeFS)  | 591.2k | 2.8k | 1.4k | 3.4k | Go | 49.3 |
| [istio](https://github.com/istio)  | 583.6k | 47.1k | 62.4k | 38.5k | Go | 48.6 |
| [openebs](https://github.com/openebs)  | 524.7k | 15.5k | 9.6k | 10.1k | Go | 43.7 |
| [argoproj](https://github.com/argoproj)  | 517.8k | 18.1k | 16.9k | 34.2k | Go | 43.1 |
| [knative](https://github.com/knative)  | 432.8k | 30.6k | 32.3k | 12.4k | Go | 36.1 |
| [apache](https://github.com/apache)  | 420.4k | 22.3k | 521.0 | 5.7k | C++ | 35.0 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762.0 | Java | 33.8 |
| [backstage](https://github.com/backstage)  | 402.7k | 40.4k | 12.9k | 21.8k | TypeScript | 33.6 |
| [tikv](https://github.com/tikv)  | 395.8k | 7.1k | 10.3k | 12.8k | Rust | 33.0 |
| [prometheus](https://github.com/prometheus)  | 377.3k | 30.9k | 19.1k | 89.5k | Go | 31.4 |
| [facebook](https://github.com/facebook)  | 357.5k | 15.6k | 13.3k | 204.1k | JavaScript | 29.8 |
| [containerd](https://github.com/containerd)  | 353.4k | 23.2k | 11.7k | 24.7k | Go | 29.4 |
| [pixie-io](https://github.com/pixie-io)  | 352.5k | 11.8k | 591.0 | 4.4k | C++ | 29.4 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.0k | 8.5k | 7.2k | 15.4k | Go | 28.6 |
| [dapr](https://github.com/dapr)  | 335.5k | 17.4k | 10.7k | 26.0k | Go | 28.0 |
| [LINBIT](https://github.com/LINBIT)  | 310.9k | 4.4k | 11.0 | 653.0 | Java | 25.9 |
| [vectordotdev](https://github.com/vectordotdev)  | 309.0k | 14.4k | 10.2k | 14.4k | Rust | 25.8 |
| [kyverno](https://github.com/kyverno)  | 280.0k | 8.8k | 5.1k | 4.0k | Go | 23.3 |
| [operator-framework](https://github.com/operator-framework)  | 265.0k | 18.9k | 14.0k | 13.3k | Go | 22.1 |
| [keptn](https://github.com/keptn)  | 256.8k | 11.8k | 8.4k | 1.9k | Go | 21.4 |
| [uber](https://github.com/uber)  | 253.8k | 1.6k | 248.0 | 6.0k | Go | 21.1 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.5k | 8.7k | 7.6k | 23.7k | Go | 20.9 |
| [goharbor](https://github.com/goharbor)  | 242.5k | 16.9k | 10.4k | 21.1k | Go | 20.2 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 13.0k | 13.4k | 18.0k | Rust | 19.8 |
| [siderolabs](https://github.com/siderolabs)  | 232.8k | 7.1k | 9.1k | 4.8k | Go | 19.4 |
| [kata-containers](https://github.com/kata-containers)  | 223.8k | 22.8k | 10.3k | 7.0k | Go | 18.6 |
| [Netflix](https://github.com/Netflix)  | 208.1k | 4.2k | 2.3k | 550.0 | Java | 17.4 |
| [spiffe](https://github.com/spiffe)  | 206.9k | 8.6k | 4.0k | 2.9k | Go | 17.2 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 206.0k | 8.1k | 2.6k | 11.0k | Go | 17.2 |
| [litmuschaos](https://github.com/litmuschaos)  | 203.2k | 7.5k | 7.1k | 4.1k | TypeScript | 16.9 |
| [fluxcd](https://github.com/fluxcd)  | 200.2k | 27.7k | 9.5k | 19.9k | Go | 16.7 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4k | 4.5k | 4.6k | 4.8k | Python | 16.4 |
| [longhorn](https://github.com/longhorn)  | 193.7k | 10.0k | 6.6k | 5.1k | Go | 16.1 |
| [etcd-io](https://github.com/etcd-io)  | 185.3k | 26.4k | 11.3k | 52.3k | Go | 15.4 |
| [containers](https://github.com/containers)  | 173.6k | 18.3k | 10.1k | 17.2k | Go | 14.5 |
| [cert-manager](https://github.com/cert-manager)  | 145.5k | 13.1k | 4.8k | 10.9k | Go | 12.1 |
| [kubeedge](https://github.com/kubeedge)  | 141.3k | 7.3k | 3.9k | 6.8k | Go | 11.8 |
| [rook](https://github.com/rook)  | 132.1k | 24.8k | 7.1k | 11.2k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 4.9k | 3.8k | 4.1k | Go | 10.9 |
| [crossplane](https://github.com/crossplane)  | 127.7k | 11.2k | 3.5k | 7.7k | Go | 10.6 |
| [thanos-io](https://github.com/thanos-io)  | 123.5k | 4.1k | 4.0k | 12.0k | Go | 10.3 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 4.9k | 3.8k | 5.1k | Go | 10.0 |
| [buildpacks](https://github.com/buildpacks)  | 116.7k | 16.1k | 3.5k | 2.8k | Go | 9.7 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 106.8k | 3.3k | 3.2k | 5.9k | Go | 8.9 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 105.7k | 18.0k | 3.3k | 4.0k | Python | 8.8 |
| [kedacore](https://github.com/kedacore)  | 105.0k | 3.9k | 4.2k | 6.8k | Go | 8.8 |
| [VKCOM](https://github.com/VKCOM)  | 101.5k | 9.9k | 3.2k | 883.0 | TypeScript | 8.4 |
| [karmada-io](https://github.com/karmada-io)  | 98.9k | 4.1k | 2.7k | 3.2k | Go | 8.2 |
| [in-toto](https://github.com/in-toto)  | 92.3k | 9.5k | 1.0k | 1.1k | Python | 7.7 |
| [cloudevents](https://github.com/cloudevents)  | 79.5k | 4.0k | 2.6k | 5.8k | Go | 6.6 |
| [coredns](https://github.com/coredns)  | 76.2k | 5.5k | 4.7k | 11.5k | Go | 6.3 |
| [gotd](https://github.com/gotd)  | 74.3k | 6.2k | 2.4k | 1.2k | Go | 6.2 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 72.3k | 1.5k | 1.7k | 1.2k | Go | 6.0 |
| [helm](https://github.com/helm)  | 71.5k | 28.9k | 26.9k | 48.8k | Go | 6.0 |
| [cri-o](https://github.com/cri-o)  | 71.3k | 9.2k | 5.8k | 4.5k | Go | 5.9 |
| [volcano-sh](https://github.com/volcano-sh)  | 68.5k | 5.5k | 2.0k | 3.2k | Go | 5.7 |
| [notaryproject](https://github.com/notaryproject)  | 64.4k | 3.7k | 1.9k | 3.3k | Go | 5.4 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 1.8k | 2.2k | 2.2k | Go | 5.2 |
| [vuejs](https://github.com/vuejs)  | 60.3k | 3.5k | 2.4k | 202.8k | TypeScript | 5.0 |
| [go-faster](https://github.com/go-faster)  | 50.8k | 4.7k | 670.0 | 151.0 | Go | 4.2 |
| [theupdateframework](https://github.com/theupdateframework)  | 45.4k | 8.6k | 2.6k | 2.6k | Rust | 3.8 |
| [vitalif](https://github.com/vitalif)  | 40.9k | 1.2k | 12.0 | 72.0 | C++ | 3.4 |
| [k3s-io](https://github.com/k3s-io)  | 38.5k | 3.5k | 3.3k | 25.3k | Go | 3.2 |
| [ogen-go](https://github.com/ogen-go)  | 38.1k | 3.9k | 1.1k | 472.0 | Go | 3.2 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 2.8k | 1.2k | 6.5k | Go | 2.9 |
| [falcosecurity](https://github.com/falcosecurity)  | 14.0k | 3.6k | 1.4k | 5.7k | C++ | 1.2 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 749.0 | 118.0 | 588.0 | Go | 1.1 |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |


### Repositories
| Repository | SLOC | Commits | PRs | Stars | Language | Effort |
|------------|------|---------|-----|-------|----------|--------|
| [torvalds/linux](https://github.com/torvalds/linux/)  | 19.1M | 1.2M | 761.0 | 148.2k | C | 1.6k |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3.4M | 145.1k | 22.8k | 172.2k | C++ | 279.2 |
| [ydb-platform/ydb](https://github.com/ydb-platform/ydb/)  | 2.8M | 9.4k | 35.0 | 2.8k | C++ | 237.4 |
| [elastic/elasticsearch](https://github.com/elastic/elasticsearch/)  | 2.4M | 67.9k | 62.5k | 63.1k | Java | 202.0 |
| [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/)  | 1.7M | 78.5k | 48.5k | 26.8k | Go | 139.7 |
| [rust-lang/rust](https://github.com/rust-lang/rust/)  | 1.5M | 220.1k | 61.4k | 79.1k | Rust | 127.1 |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1.5M | 114.5k | 73.6k | 96.5k | Go | 124.0 |
| [python/cpython](https://github.com/python/cpython/)  | 1.4M | 116.3k | 39.6k | 51.3k | Python | 114.4 |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1.1M | 109.8k | 32.2k | 27.4k | C++ | 92.7 |
| [golang/go](https://github.com/golang/go/)  | 1.1M | 55.8k | 2.9k | 109.4k | Go | 88.4 |
| [ziglang/zig](https://github.com/ziglang/zig/)  | 901.1k | 23.1k | 6.8k | 20.5k | Zig | 75.1 |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889.8k | 41.4k | 33.6k | 54.3k | TypeScript | 74.2 |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 807.6k | 17.3k | 17.1k | 21.6k | C++ | 67.3 |
| [elastic/beats](https://github.com/elastic/beats/)  | 617.8k | 16.3k | 25.8k | 11.6k | Go | 51.5 |
| [vitessio/vitess](https://github.com/vitessio/vitess/)  | 611.0k | 32.3k | 9.6k | 15.7k | Go | 50.9 |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 514.8k | 52.7k | 21.6k | 37.2k | C++ | 42.9 |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 512.3k | 96.0 | 103.0 | 22.0 | C++ | 42.7 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 504.5k | 8.5k | 454.0 | 205.0 | C++ | 42.0 |
| [cubeFS/cubefs](https://github.com/cubeFS/cubefs/)  | 478.0k | 2.3k | 1.3k | 3.1k | Go | 39.8 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477.3k | 4.2k | 3.6k | 4.4k | Go | 39.8 |
| [tikv/tikv](https://github.com/tikv/tikv/)  | 395.8k | 7.1k | 10.3k | 12.8k | Rust | 33.0 |
| [backstage/backstage](https://github.com/backstage/backstage/)  | 393.7k | 38.2k | 12.3k | 21.1k | TypeScript | 32.8 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389.2k | 9.6k | 16.2k | 1.5k | Go | 32.4 |
| [facebook/react](https://github.com/facebook/react/)  | 357.5k | 15.6k | 13.3k | 204.1k | JavaScript | 29.8 |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 352.5k | 11.8k | 591.0 | 4.4k | C++ | 29.4 |
| [AthenZ/athenz](https://github.com/AthenZ/athenz/)  | 346.1k | 2.6k | 1.9k | 751.0 | Java | 28.8 |
| [istio/istio](https://github.com/istio/istio/)  | 339.1k | 19.9k | 26.5k | 32.5k | Go | 28.3 |
| [kubernetes/autoscaler](https://github.com/kubernetes/autoscaler/)  | 332.8k | 6.5k | 3.7k | 6.6k | Go | 27.7 |
| [cilium/cilium](https://github.com/cilium/cilium/)  | 324.8k | 22.3k | 17.0k | 14.6k | Go | 27.1 |
| [docker/docker-ce](https://github.com/docker/docker-ce/)  | 324.4k | 54.3k | 662.0 | 5.6k | Go | 27.0 |
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310.9k | 4.4k | 11.0 | 653.0 | Java | 25.9 |
| [apache/mesos](https://github.com/apache/mesos/)  | 305.6k | 18.2k | 450.0 | 5.0k | C++ | 25.5 |
| [docker/labs](https://github.com/docker/labs/)  | 304.4k | 718.0 | 398.0 | 11.1k | PHP | 25.4 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284.4k | 1.6k | 1.3k | 465.0 | C Header | 23.7 |
| [kubevirt/kubevirt](https://github.com/kubevirt/kubevirt/)  | 281.1k | 16.3k | 7.0k | 3.9k | Go | 23.4 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274.6k | 9.4k | 9.4k | 12.9k | Rust | 22.9 |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 270.6k | 51.9k | 24.6k | 3.6k | Go | 22.6 |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 264.4k | 1.6k | 36.0 | 24.0 | JavaScript | 22.0 |
| [kyverno/kyverno](https://github.com/kyverno/kyverno/)  | 259.5k | 5.5k | 3.9k | 3.6k | Go | 21.6 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241.1k | 4.7k | 5.2k | 18.5k | Go | 20.1 |
| [open-policy-agent/opa](https://github.com/open-policy-agent/opa/)  | 235.4k | 4.5k | 3.5k | 7.8k | Go | 19.6 |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235.2k | 5.7k | 6.9k | 10.4k | Java | 19.6 |
| [uber/peloton](https://github.com/uber/peloton/)  | 216.4k | 705.0 | 10.0 | 582.0 | Go | 18.0 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194.2k | 181.0 | 125.0 | 1.2k | C Header | 16.2 |
| [goharbor/harbor](https://github.com/goharbor/harbor/)  | 191.5k | 11.6k | 8.0k | 19.5k | Go | 16.0 |
| [nats-io/nats-server](https://github.com/nats-io/nats-server/)  | 190.7k | 6.8k | 2.6k | 12.4k | Go | 15.9 |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186.6k | 1.0 | 2.0 | 2.0 | JavaScript | 15.6 |
| [fluent/fluent-bit](https://github.com/fluent/fluent-bit/)  | 182.4k | 9.3k | 3.3k | 4.4k | C | 15.2 |
| [cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian/)  | 175.1k | 4.1k | 4.4k | 4.7k | Python | 14.6 |
| [containers/podman](https://github.com/containers/podman/)  | 173.6k | 18.3k | 10.1k | 17.2k | Go | 14.5 |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173.3k | 8.9k | 5.3k | 2.8k | Go | 14.4 |
| [argoproj/argo-workflows](https://github.com/argoproj/argo-workflows/)  | 169.8k | 4.1k | 4.8k | 12.6k | Go | 14.2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.9k | 8.8k | 5.1k | 1.2k | Java | 14.0 |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167.6k | 19.4k | 10.4k | 14.8k | Go | 14.0 |
| [keptn/keptn](https://github.com/keptn/keptn/)  | 162.3k | 8.3k | 6.1k | 1.7k | Go | 13.5 |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157.3k | 1.7k | 1.3k | 319.0 | Java | 13.1 |
| [spiffe/spire](https://github.com/spiffe/spire/)  | 156.7k | 5.8k | 2.8k | 1.4k | Go | 13.1 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155.1k | 10.9k | 6.8k | 47.1k | Go | 12.9 |
| [containerd/containerd](https://github.com/containerd/containerd/)  | 150.2k | 12.0k | 5.9k | 13.3k | Go | 12.5 |
| [argoproj/argo-cd](https://github.com/argoproj/argo-cd/)  | 149.4k | 5.1k | 5.8k | 12.3k | Go | 12.4 |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144.5k | 5.9k | 1.8k | 8.1k | Go | 12.0 |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143.2k | 4.5k | 3.9k | 17.7k | Go | 11.9 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.1k | 929.0 | 1.8k | 224.0 | C Header | 11.8 |
| [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize/)  | 139.5k | 6.3k | 3.0k | 9.4k | Go | 11.6 |
| [kubernetes/apiserver](https://github.com/kubernetes/apiserver/)  | 131.3k | 6.2k | 25.0 | 494.0 | Go | 10.9 |
| [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers/)  | 130.8k | 10.1k | 3.0k | 3.1k | Go | 10.9 |
| [siderolabs/talos](https://github.com/siderolabs/talos/)  | 130.2k | 3.9k | 5.2k | 3.7k | Go | 10.8 |
| [etcd-io/etcd](https://github.com/etcd-io/etcd/)  | 125.7k | 19.3k | 9.2k | 42.8k | Go | 10.5 |
| [cortexproject/cortex](https://github.com/cortexproject/cortex/)  | 119.0k | 4.6k | 3.5k | 5.0k | Go | 9.9 |
| [apache/aurora](https://github.com/apache/aurora/)  | 114.8k | 4.1k | 71.0 | 629.0 | Java | 9.6 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 113.7k | 2.8k | 63.0 | 10.0 | TypeScript | 9.5 |
| [thanos-io/thanos](https://github.com/thanos-io/thanos/)  | 113.2k | 3.1k | 3.6k | 11.5k | Go | 9.4 |
| [projectcontour/contour](https://github.com/projectcontour/contour/)  | 111.5k | 4.0k | 3.2k | 3.3k | Go | 9.3 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 107.0k | 2.3k | 1.8k | 218.0 | Go | 8.9 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105.9k | 1.5k | 1.8k | 980.0 | Go | 8.8 |
| [emissary-ingress/emissary](https://github.com/emissary-ingress/emissary/)  | 105.3k | 17.9k | 3.3k | 4.0k | Python | 8.8 |
| [cockroachdb/pebble](https://github.com/cockroachdb/pebble/)  | 101.1k | 2.8k | 1.7k | 3.5k | Go | 8.4 |
| [dapr/dapr](https://github.com/dapr/dapr/)  | 101.0k | 4.0k | 3.1k | 20.7k | Go | 8.4 |
| [cert-manager/cert-manager](https://github.com/cert-manager/cert-manager/)  | 100.4k | 7.5k | 2.8k | 10.1k | Go | 8.4 |
| [openebs/maya](https://github.com/openebs/maya/)  | 100.1k | 1.8k | 1.7k | 180.0 | Go | 8.3 |
| [knative/serving](https://github.com/knative/serving/)  | 96.9k | 8.0k | 9.4k | 4.9k | Go | 8.1 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 3.3k | 3.6k | 1.5k | Java | 8.0 |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 96.0k | 3.0k | 313.0 | 2.2k | Go | 8.0 |
| [rook/rook](https://github.com/rook/rook/)  | 95.8k | 9.6k | 7.0k | 10.8k | Go | 8.0 |
| [karmada-io/karmada](https://github.com/karmada-io/karmada/)  | 94.6k | 3.6k | 2.4k | 3.1k | Go | 7.9 |
| [elastic/logstash](https://github.com/elastic/logstash/)  | 92.8k | 10.4k | 8.5k | 13.3k | Ruby | 7.7 |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90.1k | 21.0k | 7.6k | 26.0k | Go | 7.5 |
| [cubeFS/cubefs-blobstore](https://github.com/cubeFS/cubefs-blobstore/)  | 87.4k | 82.0 | 3.0 | 14.0 | Go | 7.3 |
| [knative/eventing](https://github.com/knative/eventing/)  | 87.1k | 3.5k | 4.3k | 1.3k | Go | 7.3 |
| [operator-framework/operator-lifecy~](https://github.com/operator-framework/operator-lifecycle-manager/)  | 87.0k | 4.2k | 2.2k | 1.5k | Go | 7.2 |
| [fluent/onigmo](https://github.com/fluent/onigmo/)  | 86.6k | 1.2k | 5.0 | 2.0 | C Header | 7.2 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84.5k | 3.4k | 3.2k | 210.0 | Go | 7.0 |
| [dapr/components-contrib](https://github.com/dapr/components-contrib/)  | 83.8k | 2.9k | 1.5k | 426.0 | Go | 7.0 |
| [argoproj/argo-rollouts](https://github.com/argoproj/argo-rollouts/)  | 82.6k | 1.2k | 1.5k | 1.9k | Go | 6.9 |
| [linkerd/linkerd2](https://github.com/linkerd/linkerd2/)  | 82.0k | 5.5k | 6.4k | 9.4k | Go | 6.8 |
| [fluent/fluentd](https://github.com/fluent/fluentd/)  | 80.5k | 6.5k | 1.9k | 11.8k | Ruby | 6.7 |
| [kubevirt/containerized-data-import~](https://github.com/kubevirt/containerized-data-importer/)  | 79.1k | 2.5k | 2.1k | 311.0 | Go | 6.6 |
| [docker/cli](https://github.com/docker/cli/)  | 78.9k | 8.5k | 2.7k | 3.9k | Go | 6.6 |
| [litmuschaos/litmus](https://github.com/litmuschaos/litmus/)  | 78.7k | 2.5k | 2.6k | 3.5k | TypeScript | 6.6 |
| [fluent/fluentbit-website-v3](https://github.com/fluent/fluentbit-website-v3/)  | 78.2k | 386.0 | 27.0 | 4.0 | JavaScript | 6.5 |
| [kubernetes/ingress-gce](https://github.com/kubernetes/ingress-gce/)  | 78.0k | 4.6k | 1.6k | 1.2k | Go | 6.5 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77.2k | 3.6k | 2.2k | 244.0 | Go | 6.4 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76.5k | 630.0 | 1.2k | 3.3k | Go | 6.4 |
| [chaos-mesh/chaos-mesh](https://github.com/chaos-mesh/chaos-mesh/)  | 74.9k | 1.7k | 2.4k | 5.5k | Go | 6.2 |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73.8k | 6.5k | 1.2k | 628.0 | Go | 6.2 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73.8k | 1.7k | 634.0 | 1.2k | Java | 6.2 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73.7k | 3.8k | 2.7k | 543.0 | Go | 6.1 |
| [linkerd/linkerd](https://github.com/linkerd/linkerd/)  | 73.3k | 1.4k | 1.4k | 5.4k | Scala | 6.1 |
| [grafana/k6](https://github.com/grafana/k6/)  | 73.0k | 5.1k | 1.5k | 19.7k | Go | 6.1 |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72.4k | 4.8k | 5.3k | 2.4k | TypeScript | 6.0 |
| [dragonflyoss/Dragonfly2](https://github.com/dragonflyoss/Dragonfly2/)  | 72.3k | 1.5k | 1.7k | 1.2k | Go | 6.0 |
| [jaegertracing/jaeger](https://github.com/jaegertracing/jaeger/)  | 71.9k | 2.0k | 2.3k | 17.3k | Go | 6.0 |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71.9k | 633.0 | 232.0 | 1.1k | JavaScript | 6.0 |
| [longhorn/longhorn-manager](https://github.com/longhorn/longhorn-manager/)  | 68.9k | 3.0k | 1.8k | 132.0 | Go | 5.7 |
| [kedacore/keda](https://github.com/kedacore/keda/)  | 68.3k | 1.9k | 2.1k | 6.0k | Go | 5.7 |
| [kubeedge/kubeedge](https://github.com/kubeedge/kubeedge/)  | 66.7k | 4.7k | 2.9k | 5.7k | Go | 5.6 |
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65.6k | 3.8k | 204.0 | 7.4k | Go | 5.5 |
| [nats-io/nats.c](https://github.com/nats-io/nats.c/)  | 64.9k | 972.0 | 399.0 | 301.0 | C | 5.4 |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63.9k | 2.8k | 1.5k | 2.1k | Go | 5.3 |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63.9k | 859.0 | 941.0 | 3.5k | C# | 5.3 |
| [fluent/fluent-bit-website-old](https://github.com/fluent/fluent-bit-website-old/)  | 63.2k | 19.0 | 0.0 | 2.0 | JavaScript | 5.3 |
| [ydb-platform/ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk/)  | 62.9k | 2.9k | 465.0 | 80.0 | Go | 5.2 |
| [knative/pkg](https://github.com/knative/pkg/)  | 61.2k | 2.0k | 2.3k | 234.0 | Go | 5.1 |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61.1k | 257.0 | 238.0 | 124.0 | Go | 5.1 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 1.7k | 2.0k | 1.8k | TypeScript | 5.1 |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60.9k | 1.2k | 447.0 | 27.0 | Rust | 5.1 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60.5k | 3.0k | 1.9k | 6.1k | Go | 5.0 |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60.3k | 3.5k | 2.4k | 202.8k | TypeScript | 5.0 |
| [kubernetes/legacy-cloud-providers](https://github.com/kubernetes/legacy-cloud-providers/)  | 60.2k | 1.8k | 0.0 | 47.0 | Go | 5.0 |
| [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/)  | 60.1k | 2.5k | 30.0 | 661.0 | Go | 5.0 |
| [ydb-platform/ydb-nodejs-genproto](https://github.com/ydb-platform/ydb-nodejs-genproto/)  | 59.3k | 9.0 | 3.0 | 1.0 | JavaScript | 5.0 |
| [buildpacks/pack](https://github.com/buildpacks/pack/)  | 58.4k | 3.4k | 957.0 | 1.9k | Go | 4.9 |
| [envoyproxy/envoy-mobile](https://github.com/envoyproxy/envoy-mobile/)  | 57.6k | 1.8k | 2.2k | 550.0 | Java | 4.8 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9k | 2.3k | 2.8k | 2.3k | C# | 4.7 |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56.9k | 6.9k | 4.9k | 14.5k | Go | 4.7 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 56.1k | 1.8k | 1.1k | 286.0 | Go | 4.7 |
| [volcano-sh/volcano](https://github.com/volcano-sh/volcano/)  | 55.7k | 4.6k | 1.6k | 2.9k | Go | 4.6 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 54.2k | 1.9k | 1.1k | 457.0 | Rust | 4.5 |
| [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy/)  | 54.1k | 2.2k | 2.3k | 1.7k | Rust | 4.5 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53.4k | 4.4k | 37.0 | 11.0 | Go | 4.5 |
| [gotd/td](https://github.com/gotd/td/)  | 52.6k | 3.1k | 788.0 | 1.1k | Go | 4.4 |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 1.6k | 990.0 | 420.0 | TypeScript | 4.3 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51.2k | 3.2k | 8.0 | 206.0 | Go | 4.3 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 1.9k | 2.4k | 3.7k | Go | 4.3 |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50.8k | 2.5k | 997.0 | 231.0 | Go | 4.2 |
| [coredns/coredns](https://github.com/coredns/coredns/)  | 50.7k | 3.6k | 3.8k | 10.4k | Go | 4.2 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50.7k | 233.0 | 194.0 | 9.0 | Go | 4.2 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector/)  | 50.3k | 4.4k | 5.1k | 2.8k | Go | 4.2 |
| [nats-io/nats.java](https://github.com/nats-io/nats.java/)  | 49.2k | 1.5k | 510.0 | 461.0 | Java | 4.1 |
| [kubernetes-sigs/multi-tenancy](https://github.com/kubernetes-sigs/multi-tenancy/)  | 48.9k | 2.3k | 1.1k | 930.0 | Go | 4.1 |
| [envoyproxy/pytooling](https://github.com/envoyproxy/pytooling/)  | 48.9k | 615.0 | 619.0 | 6.0 | Python | 4.1 |
| [jaegertracing/jaeger-ui](https://github.com/jaegertracing/jaeger-ui/)  | 48.9k | 1.1k | 781.0 | 894.0 | JavaScript | 4.1 |
| [nats-io/nats-streaming-server](https://github.com/nats-io/nats-streaming-server/)  | 48.9k | 1.7k | 651.0 | 2.4k | Go | 4.1 |
| [grpc/grpc-swift](https://github.com/grpc/grpc-swift/)  | 48.3k | 1.6k | 981.0 | 1.7k | Swift | 4.0 |
| [cri-o/cri-o](https://github.com/cri-o/cri-o/)  | 47.4k | 7.6k | 5.5k | 4.4k | Go | 4.0 |
| [nats-io/nats.net](https://github.com/nats-io/nats.net/)  | 46.8k | 747.0 | 430.0 | 566.0 | C# | 3.9 |
| [operator-framework/operator-sdk](https://github.com/operator-framework/operator-sdk/)  | 46.7k | 3.2k | 3.9k | 6.3k | Go | 3.9 |
| [knative/client](https://github.com/knative/client/)  | 46.6k | 1.1k | 1.3k | 303.0 | Go | 3.9 |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46.6k | 4.5k | 5.0 | 1.0 | Go | 3.9 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46.4k | 241.0 | 358.0 | 20.0 | C | 3.9 |
| [vitessio/website](https://github.com/vitessio/website/)  | 46.0k | 2.9k | 1.1k | 36.0 | JavaScript | 3.8 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45.8k | 790.0 | 309.0 | 38.0 | JavaScript | 3.8 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.8k | 1.8k | 1.0k | 413.0 | Python | 3.8 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45.3k | 1.1k | 543.0 | 113.0 | Go | 3.8 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44.9k | 4.4k | 4.9k | 12.3k | Go | 3.7 |
| [cockroachdb/gostdlib](https://github.com/cockroachdb/gostdlib/)  | 44.7k | 10.0 | 5.0 | 2.0 | Go | 3.7 |
| [notaryproject/notary](https://github.com/notaryproject/notary/)  | 43.8k | 3.0k | 992.0 | 3.0k | Go | 3.6 |
| [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper/)  | 43.7k | 1.2k | 1.5k | 3.0k | Go | 3.6 |
| [containerd/nerdctl](https://github.com/containerd/nerdctl/)  | 43.6k | 2.4k | 1.3k | 5.7k | Go | 3.6 |
| [istio/old_mixer_repo](https://github.com/istio/old_mixer_repo/)  | 42.5k | 741.0 | 1.1k | 68.0 | Go | 3.5 |
| [crossplane/crossplane](https://github.com/crossplane/crossplane/)  | 42.5k | 4.6k | 2.2k | 6.7k | Go | 3.5 |
| [longhorn/longhorn-tests](https://github.com/longhorn/longhorn-tests/)  | 42.4k | 1.5k | 1.3k | 13.0 | Python | 3.5 |
| [grpc/grpc-node](https://github.com/grpc/grpc-node/)  | 42.3k | 4.2k | 1.4k | 3.9k | TypeScript | 3.5 |
| [open-telemetry/opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python/)  | 42.2k | 1.4k | 1.7k | 1.2k | Python | 3.5 |
| [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)  | 42.1k | 2.2k | 1.3k | 1.9k | Go | 3.5 |
| [open-telemetry/opentelemetry-cpp-c~](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42.1k | 157.0 | 180.0 | 78.0 | Python | 3.5 |
| [helm/helm](https://github.com/helm/helm/)  | 41.7k | 6.8k | 4.8k | 23.9k | Go | 3.5 |
| [open-telemetry/opentelemetry-ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4k | 268.0 | 112.0 | 85.0 | C++ | 3.5 |
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40.9k | 1.2k | 12.0 | 72.0 | C++ | 3.4 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 40.8k | 52.0 | 3.0 | 5.0 | JavaScript | 3.4 |
| [keptn/tutorials](https://github.com/keptn/tutorials/)  | 40.8k | 622.0 | 193.0 | 10.0 | JavaScript | 3.4 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40.4k | 9.6k | 3.0k | 759.0 | TypeScript | 3.4 |
| [dapr/java-sdk](https://github.com/dapr/java-sdk/)  | 39.7k | 409.0 | 490.0 | 223.0 | Java | 3.3 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39.1k | 177.0 | 142.0 | 13.0 | Go | 3.3 |
| [nats-io/nats.go](https://github.com/nats-io/nats.go/)  | 38.8k | 1.8k | 716.0 | 4.4k | Go | 3.2 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38.0k | 3.0k | 1.8k | 6.2k | Go | 3.2 |
| [uber/kraken](https://github.com/uber/kraken/)  | 37.4k | 867.0 | 238.0 | 5.4k | Go | 3.1 |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37.3k | 799.0 | 283.0 | 321.0 | Go | 3.1 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36.7k | 3.3k | 700.0 | 454.0 | Go | 3.0 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36.4k | 35.0 | 7.0 | 1.0 | C | 3.0 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36.3k | 298.0 | 358.0 | 83.0 | Go | 3.0 |
| [argoproj/argo-events](https://github.com/argoproj/argo-events/)  | 35.9k | 1.3k | 1.5k | 1.9k | Go | 3.0 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35.5k | 662.0 | 313.0 | 278.0 | Go | 3.0 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34.2k | 1.3k | 462.0 | 2.3k | Go | 2.9 |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34.1k | 13.0 | 0.0 | 6.0 | JavaScript | 2.8 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33.7k | 1.5k | 420.0 | 78.0 | Go | 2.8 |
| [containerssh/libcontainerssh](https://github.com/containerssh/libcontainerssh/)  | 33.7k | 324.0 | 384.0 | 28.0 | Go | 2.8 |
| [open-telemetry/opentelemetry-cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.6k | 1.0k | 1.2k | 432.0 | C++ | 2.8 |
| [kubeedge/sedna](https://github.com/kubeedge/sedna/)  | 33.4k | 475.0 | 215.0 | 428.0 | Python | 2.8 |
| [kubernetes/perf-tests](https://github.com/kubernetes/perf-tests/)  | 33.4k | 3.1k | 1.9k | 764.0 | Go | 2.8 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.4k | 732.0 | 850.0 | 230.0 | C# | 2.8 |
| [in-toto/github-action](https://github.com/in-toto/github-action/)  | 33.2k | 3.0 | 0.0 | 4.0 | JavaScript | 2.8 |
| [buildpacks/lifecycle](https://github.com/buildpacks/lifecycle/)  | 32.9k | 1.6k | 665.0 | 158.0 | Go | 2.7 |
| [operator-framework/operator-regist~](https://github.com/operator-framework/operator-registry/)  | 32.6k | 1.1k | 780.0 | 181.0 | Go | 2.7 |
| [goharbor/harbor-operator](https://github.com/goharbor/harbor-operator/)  | 32.5k | 1.0k | 537.0 | 293.0 | Go | 2.7 |
| [keptn/lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit/)  | 32.2k | 497.0 | 666.0 | 56.0 | Go | 2.7 |
| [cloudevents/sdk-go](https://github.com/cloudevents/sdk-go/)  | 32.2k | 624.0 | 542.0 | 613.0 | Go | 2.7 |
| [nats-io/nsc](https://github.com/nats-io/nsc/)  | 31.9k | 604.0 | 358.0 | 61.0 | Go | 2.7 |
| [prometheus/alertmanager](https://github.com/prometheus/alertmanager/)  | 31.9k | 2.8k | 1.7k | 5.5k | Go | 2.7 |
| [ydb-platform/xorm](https://github.com/ydb-platform/xorm/)  | 31.9k | 1.7k | 9.0 | 0.0 | Go | 2.7 |
| [litmuschaos/test-tools](https://github.com/litmuschaos/test-tools/)  | 31.8k | 226.0 | 300.0 | 27.0 | C | 2.6 |
| [kubevirt/hyperconverged-cluster-op~](https://github.com/kubevirt/hyperconverged-cluster-operator/)  | 31.4k | 1.4k | 2.2k | 106.0 | Go | 2.6 |
| [knative/func](https://github.com/knative/func/)  | 31.2k | 1.2k | 1.1k | 156.0 | Go | 2.6 |
| [cockroachdb/docs](https://github.com/cockroachdb/docs/)  | 31.0k | 14.0k | 6.3k | 169.0 | Java | 2.6 |
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 508.0 | 591.0 | 519.0 | PHP | 2.6 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30.9k | 1.3k | 308.0 | 232.0 | Go | 2.6 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30.9k | 18.0 | 4.0 | 0.0 | C Header | 2.6 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 574.0 | 597.0 | 997.0 | Rust | 2.5 |
| [fluxcd/source-controller](https://github.com/fluxcd/source-controller/)  | 30.1k | 1.9k | 759.0 | 191.0 | Go | 2.5 |
| [docker/machine](https://github.com/docker/machine/)  | 29.9k | 3.5k | 2.0k | 6.6k | Go | 2.5 |
| [cilium/ebpf](https://github.com/cilium/ebpf/)  | 29.9k | 1.3k | 686.0 | 4.2k | Go | 2.5 |
| [fluxcd/flux](https://github.com/fluxcd/flux/)  | 29.9k | 5.3k | 1.9k | 6.9k | Go | 2.5 |
| [nats-io/nats.deno](https://github.com/nats-io/nats.deno/)  | 29.7k | 442.0 | 413.0 | 115.0 | TypeScript | 2.5 |
| [m3db/m3cluster](https://github.com/m3db/m3cluster/)  | 29.6k | 238.0 | 227.0 | 21.0 | Go | 2.5 |
| [kubevirt/project-infra](https://github.com/kubevirt/project-infra/)  | 29.5k | 2.4k | 2.5k | 21.0 | Go | 2.5 |
| [k3s-io/k3s](https://github.com/k3s-io/k3s/)  | 29.3k | 2.6k | 2.6k | 22.5k | Go | 2.4 |
| [kubevirt/web-ui-components](https://github.com/kubevirt/web-ui-components/)  | 28.5k | 416.0 | 500.0 | 9.0 | JavaScript | 2.4 |
| [litmuschaos/litmus-go](https://github.com/litmuschaos/litmus-go/)  | 28.5k | 425.0 | 558.0 | 51.0 | Go | 2.4 |
| [jaegertracing/jaeger-operator](https://github.com/jaegertracing/jaeger-operator/)  | 28.5k | 1.1k | 1.3k | 904.0 | Go | 2.4 |
| [open-telemetry/opentelemetry-log-c~](https://github.com/open-telemetry/opentelemetry-log-collection/)  | 28.1k | 267.0 | 373.0 | 92.0 | Go | 2.3 |
| [envoyproxy/nighthawk](https://github.com/envoyproxy/nighthawk/)  | 27.9k | 578.0 | 655.0 | 272.0 | C++ | 2.3 |
| [kubernetes-sigs/alibaba-cloud-csi-~](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/)  | 27.6k | 1.3k | 590.0 | 444.0 | Go | 2.3 |
| [fluxcd/flagger](https://github.com/fluxcd/flagger/)  | 27.6k | 2.7k | 753.0 | 4.1k | Go | 2.3 |
| [fluxcd/pkg](https://github.com/fluxcd/pkg/)  | 27.5k | 1.1k | 454.0 | 30.0 | Go | 2.3 |
| [kubernetes/release](https://github.com/kubernetes/release/)  | 27.3k | 5.2k | 2.4k | 436.0 | Go | 2.3 |
| [docker/docker-py](https://github.com/docker/docker-py/)  | 27.3k | 3.3k | 1.5k | 6.1k | Python | 2.3 |
| [open-telemetry/opentelemetry-go-co~](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.3k | 1.0k | 3.2k | 684.0 | Go | 2.3 |
| [cilium/hubble-ui](https://github.com/cilium/hubble-ui/)  | 26.8k | 385.0 | 426.0 | 237.0 | TypeScript | 2.2 |
| [ClickHouse/libhdfs3](https://github.com/ClickHouse/libhdfs3/)  | 26.7k | 64.0 | 34.0 | 23.0 | C++ | 2.2 |
| [ClickHouse/clickhouse-cpp](https://github.com/ClickHouse/clickhouse-cpp/)  | 26.7k | 673.0 | 172.0 | 215.0 | C++ | 2.2 |
| [kubevirt/vm-import-operator](https://github.com/kubevirt/vm-import-operator/)  | 26.6k | 900.0 | 415.0 | 13.0 | Go | 2.2 |
| [dapr/dotnet-sdk](https://github.com/dapr/dotnet-sdk/)  | 26.5k | 550.0 | 465.0 | 967.0 | C# | 2.2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26.3k | 562.0 | 644.0 | 80.0 | Java | 2.2 |
| [containernetworking/plugins](https://github.com/containernetworking/plugins/)  | 26.2k | 1.6k | 484.0 | 1.7k | Go | 2.2 |
| [open-telemetry/opentelemetry-swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9k | 618.0 | 252.0 | 126.0 | Swift | 2.2 |
| [cockroachdb/cockroach-operator](https://github.com/cockroachdb/cockroach-operator/)  | 25.8k | 589.0 | 461.0 | 216.0 | Go | 2.1 |
| [longhorn/longhorn-ui](https://github.com/longhorn/longhorn-ui/)  | 25.6k | 817.0 | 595.0 | 92.0 | JavaScript | 2.1 |
| [envoyproxy/java-control-plane](https://github.com/envoyproxy/java-control-plane/)  | 25.5k | 235.0 | 206.0 | 246.0 | Protocol Buffers | 2.1 |
| [grafana/grafana-plugin-sdk-go](https://github.com/grafana/grafana-plugin-sdk-go/)  | 25.4k | 416.0 | 531.0 | 157.0 | Go | 2.1 |
| [kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)  | 25.3k | 642.0 | 349.0 | 733.0 | Go | 2.1 |
| [containerd/stargz-snapshotter](https://github.com/containerd/stargz-snapshotter/)  | 25.1k | 1.4k | 1.0k | 796.0 | Go | 2.1 |
| [kubernetes-sigs/kubefed](https://github.com/kubernetes-sigs/kubefed/)  | 25.0k | 2.7k | 978.0 | 2.5k | Go | 2.1 |
| [docker/compose-cli](https://github.com/docker/compose-cli/)  | 25.0k | 3.0k | 1.2k | 910.0 | Go | 2.1 |
| [grafana/terraform-provider-grafana](https://github.com/grafana/terraform-provider-grafana/)  | 24.6k | 697.0 | 523.0 | 312.0 | Go | 2.0 |
| [siderolabs/theila](https://github.com/siderolabs/theila/)  | 24.3k | 100.0 | 114.0 | 45.0 | TypeScript | 2.0 |
| [fluent/cmetrics](https://github.com/fluent/cmetrics/)  | 24.1k | 431.0 | 149.0 | 9.0 | C | 2.0 |
| [ydb-platform/ydb-embedded-ui](https://github.com/ydb-platform/ydb-embedded-ui/)  | 24.1k | 611.0 | 334.0 | 13.0 | TypeScript | 2.0 |
| [in-toto/toto-pip](https://github.com/in-toto/toto-pip/)  | 23.7k | 4.9k | 0.0 | 0.0 | Python | 2.0 |
| [kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/)  | 23.7k | 2.6k | 1.2k | 4.4k | Go | 2.0 |
| [open-telemetry/opentelemetry-opera~](https://github.com/open-telemetry/opentelemetry-operator/)  | 23.6k | 781.0 | 1.0k | 675.0 | Go | 2.0 |
| [argoproj/gitops-engine](https://github.com/argoproj/gitops-engine/)  | 23.5k | 254.0 | 454.0 | 1.5k | Go | 2.0 |
| [longhorn/longhorn-engine](https://github.com/longhorn/longhorn-engine/)  | 23.3k | 1.5k | 715.0 | 297.0 | Go | 1.9 |
| [nats-io/nats-mq](https://github.com/nats-io/nats-mq/)  | 23.3k | 88.0 | 13.0 | 18.0 | Go | 1.9 |
| [openebs/elves](https://github.com/openebs/elves/)  | 23.2k | 234.0 | 47.0 | 17.0 | Python | 1.9 |
| [openebs/node-disk-manager](https://github.com/openebs/node-disk-manager/)  | 23.2k | 451.0 | 533.0 | 170.0 | Go | 1.9 |
| [istio/old_pilot_repo](https://github.com/istio/old_pilot_repo/)  | 23.0k | 688.0 | 1.2k | 138.0 | Go | 1.9 |
| [nats-io/nats.rs](https://github.com/nats-io/nats.rs/)  | 22.8k | 1.4k | 687.0 | 698.0 | Rust | 1.9 |
| [containerd/cri](https://github.com/containerd/cri/)  | 22.6k | 2.1k | 1.2k | 897.0 | Go | 1.9 |
| [istio/proxy](https://github.com/istio/proxy/)  | 22.4k | 2.5k | 4.2k | 700.0 | C++ | 1.9 |
| [kubernetes-sigs/azuredisk-csi-driv~](https://github.com/kubernetes-sigs/azuredisk-csi-driver/)  | 21.7k | 2.7k | 1.3k | 114.0 | Go | 1.8 |
| [envoyproxy/gateway](https://github.com/envoyproxy/gateway/)  | 21.6k | 539.0 | 632.0 | 886.0 | Go | 1.8 |
| [cri-o/image](https://github.com/cri-o/image/)  | 21.6k | 1.1k | 2.0 | 2.0 | Go | 1.8 |
| [istio/operator](https://github.com/istio/operator/)  | 21.3k | 508.0 | 780.0 | 174.0 | Go | 1.8 |
| [kedacore/sample-dotnet-worker-serv~](https://github.com/kedacore/sample-dotnet-worker-servicebus-queue/)  | 21.3k | 38.0 | 30.0 | 68.0 | JavaScript | 1.8 |
| [ydb-platform/ydb-java-sdk](https://github.com/ydb-platform/ydb-java-sdk/)  | 21.0k | 598.0 | 32.0 | 12.0 | Java | 1.8 |
| [containerd/overlaybd](https://github.com/containerd/overlaybd/)  | 20.9k | 285.0 | 147.0 | 144.0 | C++ | 1.7 |
| [kubernetes-sigs/structured-merge-d~](https://github.com/kubernetes-sigs/structured-merge-diff/)  | 20.9k | 495.0 | 208.0 | 74.0 | Go | 1.7 |
| [crossplane/crossplane-runtime](https://github.com/crossplane/crossplane-runtime/)  | 20.9k | 1.2k | 260.0 | 98.0 | Go | 1.7 |
| [prometheus/procfs](https://github.com/prometheus/procfs/)  | 20.8k | 572.0 | 372.0 | 618.0 | Go | 1.7 |
| [prometheus/client_golang](https://github.com/prometheus/client_golang/)  | 20.7k | 1.4k | 741.0 | 4.4k | Go | 1.7 |
| [docker/buildx](https://github.com/docker/buildx/)  | 20.6k | 1.4k | 723.0 | 2.5k | Go | 1.7 |
| [kubernetes/cloud-provider-vsphere](https://github.com/kubernetes/cloud-provider-vsphere/)  | 20.5k | 1.1k | 472.0 | 189.0 | Go | 1.7 |
| [nats-io/nats-site](https://github.com/nats-io/nats-site/)  | 20.5k | 2.1k | 743.0 | 83.0 | JavaScript | 1.7 |
| [fluxcd/flux2](https://github.com/fluxcd/flux2/)  | 20.4k | 2.7k | 1.2k | 4.5k | Go | 1.7 |
| [kubernetes/cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack/)  | 20.3k | 2.7k | 1.4k | 507.0 | Go | 1.7 |
| [open-telemetry/opentelemetry-ruby-~](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 20.0k | 895.0 | 241.0 | 29.0 | Ruby | 1.7 |
| [operator-framework/operatorhub.io](https://github.com/operator-framework/operatorhub.io/)  | 19.9k | 620.0 | 330.0 | 40.0 | TypeScript | 1.7 |
| [jaegertracing/jaeger-client-go](https://github.com/jaegertracing/jaeger-client-go/)  | 19.4k | 433.0 | 408.0 | 1.3k | Go | 1.6 |
| [prometheus/node_exporter](https://github.com/prometheus/node_exporter/)  | 19.3k | 2.0k | 1.4k | 8.6k | Go | 1.6 |
| [envoyproxy/data-plane-api](https://github.com/envoyproxy/data-plane-api/)  | 19.2k | 2.4k | 540.0 | 510.0 | Protocol Buffers | 1.6 |
| [kubernetes-sigs/descheduler](https://github.com/kubernetes-sigs/descheduler/)  | 19.1k | 1.3k | 667.0 | 3.2k | Go | 1.6 |
| [cilium/cilium-olm](https://github.com/cilium/cilium-olm/)  | 19.0k | 341.0 | 55.0 | 11.0 | Go Template | 1.6 |
| [grpc/grpc-dart](https://github.com/grpc/grpc-dart/)  | 18.7k | 228.0 | 297.0 | 759.0 | Dart | 1.6 |
| [crossplane/oam-kubernetes-runtime](https://github.com/crossplane/oam-kubernetes-runtime/)  | 18.7k | 415.0 | 203.0 | 279.0 | Go | 1.6 |
| [cloud-custodian/cel-python](https://github.com/cloud-custodian/cel-python/)  | 18.6k | 115.0 | 24.0 | 60.0 | Python | 1.6 |
| [cilium/cilium-cli](https://github.com/cilium/cilium-cli/)  | 18.4k | 1.4k | 1.1k | 205.0 | Go | 1.5 |
| [knative/client-pkg](https://github.com/knative/client-pkg/)  | 18.1k | 80.0 | 86.0 | 3.0 | Go | 1.5 |
| [AthenZ/athenz-authorizer](https://github.com/AthenZ/athenz-authorizer/)  | 18.1k | 85.0 | 2.0 | 1.0 | Go | 1.5 |
| [kubernetes-sigs/gcp-filestore-csi-~](https://github.com/kubernetes-sigs/gcp-filestore-csi-driver/)  | 17.9k | 593.0 | 403.0 | 68.0 | Go | 1.5 |
| [keptn/go-utils](https://github.com/keptn/go-utils/)  | 17.6k | 895.0 | 565.0 | 7.0 | Go | 1.5 |
| [istio/tools](https://github.com/istio/tools/)  | 17.5k | 1.8k | 2.4k | 302.0 | Go | 1.5 |
| [openebs/dynamic-nfs-provisioner](https://github.com/openebs/dynamic-nfs-provisioner/)  | 17.3k | 112.0 | 109.0 | 103.0 | Go | 1.4 |
| [docker/compose](https://github.com/docker/compose/)  | 17.3k | 4.0k | 3.5k | 28.8k | Go | 1.4 |
| [kubevirt/ssp-operator](https://github.com/kubevirt/ssp-operator/)  | 17.3k | 629.0 | 501.0 | 12.0 | Go | 1.4 |
| [docker/compose-on-kubernetes](https://github.com/docker/compose-on-kubernetes/)  | 17.2k | 236.0 | 114.0 | 1.4k | Go | 1.4 |
| [ydb-platform/ydb-python-sdk](https://github.com/ydb-platform/ydb-python-sdk/)  | 17.1k | 594.0 | 132.0 | 54.0 | Python | 1.4 |
| [kubernetes-sigs/etcdadm](https://github.com/kubernetes-sigs/etcdadm/)  | 16.9k | 1.2k | 248.0 | 652.0 | Go | 1.4 |
| [kubernetes/api](https://github.com/kubernetes/api/)  | 16.9k | 6.7k | 10.0 | 514.0 | Go | 1.4 |
| [kubevirt/kubevirt-tekton-tasks](https://github.com/kubevirt/kubevirt-tekton-tasks/)  | 16.7k | 526.0 | 209.0 | 20.0 | Go | 1.4 |
| [ClickHouse/libcxx](https://github.com/ClickHouse/libcxx/)  | 16.5k | 38.0 | 14.0 | 0.0 | C++ | 1.4 |
| [open-telemetry/opentelemetry-ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16.4k | 763.0 | 919.0 | 375.0 | Ruby | 1.4 |
| [argoproj/argocon21](https://github.com/argoproj/argocon21/)  | 16.4k | 67.0 | 19.0 | 7.0 | JavaScript | 1.4 |
| [ClickHouse/clickhouse-odbc](https://github.com/ClickHouse/clickhouse-odbc/)  | 16.3k | 1.1k | 172.0 | 222.0 | C++ | 1.4 |
| [kubernetes-sigs/aws-ebs-csi-driver](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/)  | 16.2k | 1.7k | 882.0 | 721.0 | Go | 1.4 |
| [ClickHouse/boost](https://github.com/ClickHouse/boost/)  | 16.2k | 83.0 | 28.0 | 1.0 | C++ | 1.4 |
| [ClickHouse/antlr4-runtime](https://github.com/ClickHouse/antlr4-runtime/)  | 16.2k | 306.0 | 0.0 | 1.0 | C++ | 1.4 |
| [fluxcd/go-git-providers](https://github.com/fluxcd/go-git-providers/)  | 16.1k | 347.0 | 142.0 | 62.0 | Go | 1.4 |
| [litmuschaos/litmus-e2e](https://github.com/litmuschaos/litmus-e2e/)  | 16.0k | 231.0 | 360.0 | 15.0 | JavaScript | 1.3 |
| [kubernetes-sigs/boskos](https://github.com/kubernetes-sigs/boskos/)  | 16.0k | 923.0 | 129.0 | 93.0 | Go | 1.3 |
| [knative/test-infra](https://github.com/knative/test-infra/)  | 15.9k | 2.7k | 3.0k | 82.0 | Go | 1.3 |
| [knative/toolbox](https://github.com/knative/toolbox/)  | 15.9k | 2.7k | 3.0 | 0.0 | Go | 1.3 |
| [openebs/dynamic-localpv-provisioner](https://github.com/openebs/dynamic-localpv-provisioner/)  | 15.9k | 154.0 | 123.0 | 88.0 | Go | 1.3 |
| [nats-io/natscli](https://github.com/nats-io/natscli/)  | 15.9k | 1.7k | 543.0 | 280.0 | Go | 1.3 |
| [kubernetes-sigs/gcp-compute-persis~](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver/)  | 15.7k | 1.4k | 824.0 | 133.0 | Go | 1.3 |
| [prometheus/client_java](https://github.com/prometheus/client_java/)  | 15.6k | 578.0 | 490.0 | 1.9k | Java | 1.3 |
| [grafana/carbon-relay-ng](https://github.com/grafana/carbon-relay-ng/)  | 15.6k | 1.1k | 258.0 | 454.0 | Go | 1.3 |
| [cloudevents/sdk-java](https://github.com/cloudevents/sdk-java/)  | 15.5k | 755.0 | 370.0 | 315.0 | Java | 1.3 |
| [go-faster/yamlx](https://github.com/go-faster/yamlx/)  | 15.4k | 540.0 | 38.0 | 5.0 | Go | 1.3 |
| [nats-io/nats-jms-bridge](https://github.com/nats-io/nats-jms-bridge/)  | 15.2k | 398.0 | 90.0 | 11.0 | Java | 1.3 |
| [kubernetes-sigs/azurefile-csi-driv~](https://github.com/kubernetes-sigs/azurefile-csi-driver/)  | 15.2k | 2.4k | 886.0 | 113.0 | Go | 1.3 |
| [ClickHouse/boost-extra](https://github.com/ClickHouse/boost-extra/)  | 15.2k | 34.0 | 0.0 | 0.0 | C++ | 1.3 |
| [kubernetes-sigs/promo-tools](https://github.com/kubernetes-sigs/promo-tools/)  | 15.1k | 1.6k | 602.0 | 125.0 | Go | 1.3 |
| [etcd-io/raft](https://github.com/etcd-io/raft/)  | 15.0k | 1.2k | 19.0 | 139.0 | Go | 1.2 |
| [open-policy-agent/frameworks](https://github.com/open-policy-agent/frameworks/)  | 15.0k | 373.0 | 238.0 | 105.0 | Go | 1.2 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-openstack/)  | 15.0k | 1.4k | 977.0 | 220.0 | Go | 1.2 |
| [ClickHouse/ch-go](https://github.com/ClickHouse/ch-go/)  | 15.0k | 1.1k | 230.0 | 230.0 | Go | 1.2 |
| [kubernetes/cloud-provider-aws](https://github.com/kubernetes/cloud-provider-aws/)  | 14.9k | 975.0 | 401.0 | 294.0 | Go | 1.2 |
| [dapr/cli](https://github.com/dapr/cli/)  | 14.7k | 783.0 | 654.0 | 275.0 | Go | 1.2 |
| [theupdateframework/rust-tuf](https://github.com/theupdateframework/rust-tuf/)  | 14.6k | 756.0 | 221.0 | 155.0 | Rust | 1.2 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-ibmcloud/)  | 14.6k | 727.0 | 840.0 | 55.0 | Go | 1.2 |
| [openebs/jiva](https://github.com/openebs/jiva/)  | 14.6k | 804.0 | 337.0 | 133.0 | Go | 1.2 |
| [kubernetes-sigs/kind](https://github.com/kubernetes-sigs/kind/)  | 14.6k | 3.8k | 1.6k | 11.2k | Go | 1.2 |
| [kubevirt/cluster-network-addons-op~](https://github.com/kubevirt/cluster-network-addons-operator/)  | 14.4k | 1.2k | 1.4k | 57.0 | Go | 1.2 |
| [docker/libcompose](https://github.com/docker/libcompose/)  | 14.4k | 673.0 | 361.0 | 587.0 | Go | 1.2 |
| [docker/go-docker](https://github.com/docker/go-docker/)  | 14.3k | 7.0 | 4.0 | 179.0 | Go | 1.2 |
| [kubernetes/cloud-provider](https://github.com/kubernetes/cloud-provider/)  | 14.2k | 1.0k | 0.0 | 182.0 | Go | 1.2 |
| [nats-io/go-nats](https://github.com/nats-io/go-nats/)  | 14.2k | 937.0 | 0.0 | 20.0 | Go | 1.2 |
| [VictoriaMetrics/grafana-datasource](https://github.com/VictoriaMetrics/grafana-datasource/)  | 14.1k | 189.0 | 29.0 | 23.0 | TypeScript | 1.2 |
| [fluent/fluentd-kubernetes-daemonset](https://github.com/fluent/fluentd-kubernetes-daemonset/)  | 14.1k | 826.0 | 1.0k | 1.2k | Ruby | 1.2 |
| [cubeFS/shuttle](https://github.com/cubeFS/shuttle/)  | 14.0k | 59.0 | 8.0 | 127.0 | Java | 1.2 |
| [falcosecurity/falco](https://github.com/falcosecurity/falco/)  | 14.0k | 3.6k | 1.4k | 5.7k | C++ | 1.2 |
| [nats-io/nats.py](https://github.com/nats-io/nats.py/)  | 13.8k | 687.0 | 228.0 | 634.0 | Python | 1.1 |
| [kubernetes/component-base](https://github.com/kubernetes/component-base/)  | 13.7k | 1.3k | 0.0 | 84.0 | Go | 1.1 |
| [kubernetes-sigs/node-feature-disco~](https://github.com/kubernetes-sigs/node-feature-discovery/)  | 13.6k | 1.6k | 833.0 | 541.0 | Go | 1.1 |
| [open-policy-agent/setup-opa](https://github.com/open-policy-agent/setup-opa/)  | 13.6k | 31.0 | 17.0 | 33.0 | JavaScript | 1.1 |
| [docker/awesome-compose](https://github.com/docker/awesome-compose/)  | 13.5k | 267.0 | 255.0 | 21.9k | JavaScript | 1.1 |
| [prometheus/promlens](https://github.com/prometheus/promlens/)  | 13.5k | 78.0 | 81.0 | 307.0 | TypeScript | 1.1 |
| [grafana/har-to-k6](https://github.com/grafana/har-to-k6/)  | 13.5k | 627.0 | 84.0 | 77.0 | JavaScript | 1.1 |
| [istio/bots](https://github.com/istio/bots/)  | 13.5k | 638.0 | 652.0 | 10.0 | Go | 1.1 |
| [kyverno/policy-reporter](https://github.com/kyverno/policy-reporter/)  | 13.5k | 445.0 | 188.0 | 170.0 | Go | 1.1 |
| [knative/operator](https://github.com/knative/operator/)  | 13.3k | 855.0 | 1.1k | 135.0 | Go | 1.1 |
| [m3db/m3coordinator](https://github.com/m3db/m3coordinator/)  | 13.3k | 68.0 | 54.0 | 4.0 | Go | 1.1 |
| [theupdateframework/go-tuf](https://github.com/theupdateframework/go-tuf/)  | 13.2k | 450.0 | 306.0 | 531.0 | Go | 1.1 |
| [m3db/m3ninx](https://github.com/m3db/m3ninx/)  | 13.2k | 67.0 | 77.0 | 3.0 | Go | 1.1 |
| [envoyproxy/playground](https://github.com/envoyproxy/playground/)  | 13.1k | 192.0 | 168.0 | 7.0 | JavaScript | 1.1 |
| [AthenZ/athenz-client-sidecar](https://github.com/AthenZ/athenz-client-sidecar/)  | 13.1k | 72.0 | 4.0 | 2.0 | Go | 1.1 |
| [go-faster/hx](https://github.com/go-faster/hx/)  | 13.0k | 1.5k | 12.0 | 4.0 | Go | 1.1 |
| [prometheus/common](https://github.com/prometheus/common/)  | 13.0k | 497.0 | 362.0 | 224.0 | Go | 1.1 |
| [containerd/nydus-snapshotter](https://github.com/containerd/nydus-snapshotter/)  | 12.9k | 915.0 | 327.0 | 59.0 | Go | 1.1 |
| [open-telemetry/opentelemetry-php-c~](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12.8k | 322.0 | 112.0 | 22.0 | PHP | 1.1 |
| [fluent/ctraces](https://github.com/fluent/ctraces/)  | 12.7k | 150.0 | 35.0 | 0.0 | C | 1.1 |
| [carina-io/carina](https://github.com/carina-io/carina/)  | 12.7k | 712.0 | 115.0 | 588.0 | Go | 1.1 |
| [m3db/m3x](https://github.com/m3db/m3x/)  | 12.7k | 208.0 | 206.0 | 17.0 | Go | 1.1 |
| [etcd-io/bbolt](https://github.com/etcd-io/bbolt/)  | 12.6k | 1.2k | 220.0 | 6.2k | Go | 1.1 |
| [kubeedge/mappers-go](https://github.com/kubeedge/mappers-go/)  | 12.6k | 158.0 | 74.0 | 33.0 | Go | 1.1 |
| [argoproj/applicationset](https://github.com/argoproj/applicationset/)  | 12.6k | 195.0 | 255.0 | 566.0 | Go | 1.1 |
| [cockroachdb/sequelize-cockroachdb](https://github.com/cockroachdb/sequelize-cockroachdb/)  | 12.5k | 289.0 | 114.0 | 49.0 | JavaScript | 1.0 |
| [ydb-platform/ydb-rs-sdk](https://github.com/ydb-platform/ydb-rs-sdk/)  | 12.4k | 604.0 | 105.0 | 37.0 | Rust | 1.0 |
| [rook/nfs](https://github.com/rook/nfs/)  | 12.4k | 7.4k | 22.0 | 28.0 | Go | 1.0 |
| [jaegertracing/jaeger-client-java](https://github.com/jaegertracing/jaeger-client-java/)  | 12.4k | 528.0 | 504.0 | 485.0 | Java | 1.0 |
| [siderolabs/sidero](https://github.com/siderolabs/sidero/)  | 12.4k | 366.0 | 839.0 | 257.0 | Go | 1.0 |
| [docker/app](https://github.com/docker/app/)  | 12.3k | 1.5k | 650.0 | 1.6k | Go | 1.0 |
| [grafana/ksonnet](https://github.com/grafana/ksonnet/)  | 12.2k | 434.0 | 0.0 | 0.0 | Go | 1.0 |
| [dapr/python-sdk](https://github.com/dapr/python-sdk/)  | 12.1k | 327.0 | 351.0 | 173.0 | Python | 1.0 |
| [m3db/m3db-operator](https://github.com/m3db/m3db-operator/)  | 12.0k | 230.0 | 247.0 | 134.0 | Go | 1.0 |
| [openebs/openebsctl](https://github.com/openebs/openebsctl/)  | 11.9k | 112.0 | 105.0 | 27.0 | Go | 990.0m |
| [kata-containers/agent](https://github.com/kata-containers/agent/)  | 11.9k | 833.0 | 515.0 | 237.0 | Go | 990.0m |
| [openebs/website](https://github.com/openebs/website/)  | 11.9k | 1.2k | 342.0 | 11.0 | TypeScript | 990.0m |
| [rook/cassandra](https://github.com/rook/cassandra/)  | 11.8k | 7.4k | 13.0 | 5.0 | Go | 980.0m |
| [spiffe/go-spiffe](https://github.com/spiffe/go-spiffe/)  | 11.8k | 215.0 | 155.0 | 97.0 | Go | 980.0m |
| [nats-io/jwt](https://github.com/nats-io/jwt/)  | 11.7k | 288.0 | 165.0 | 63.0 | Go | 980.0m |
| [fluent/fluent-operator](https://github.com/fluent/fluent-operator/)  | 11.7k | 619.0 | 444.0 | 369.0 | Go | 970.0m |
| [grafana/jslib.k6.io](https://github.com/grafana/jslib.k6.io/)  | 11.6k | 189.0 | 71.0 | 32.0 | JavaScript | 960.0m |
| [open-telemetry/opentelemetry-erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11.5k | 1.1k | 355.0 | 253.0 | Erlang | 960.0m |
| [spiffe/java-spiffe](https://github.com/spiffe/java-spiffe/)  | 11.5k | 321.0 | 90.0 | 18.0 | Java | 960.0m |
| [prometheus/pushgateway](https://github.com/prometheus/pushgateway/)  | 11.5k | 669.0 | 267.0 | 2.5k | JavaScript | 960.0m |
| [jaegertracing/legacy-client-java](https://github.com/jaegertracing/legacy-client-java/)  | 11.5k | 292.0 | 14.0 | 5.0 | Java | 960.0m |
| [kubernetes-sigs/blob-csi-driver](https://github.com/kubernetes-sigs/blob-csi-driver/)  | 11.4k | 1.7k | 645.0 | 97.0 | Go | 950.0m |
| [openebs/zfs-localpv](https://github.com/openebs/zfs-localpv/)  | 11.4k | 275.0 | 320.0 | 261.0 | Go | 950.0m |
| [docker/getting-started](https://github.com/docker/getting-started/)  | 11.4k | 190.0 | 212.0 | 2.5k | JavaScript | 950.0m |
| [jaegertracing/jaeger-client-csharp](https://github.com/jaegertracing/jaeger-client-csharp/)  | 11.3k | 285.0 | 105.0 | 308.0 | C# | 940.0m |
| [rook/rook-client-python](https://github.com/rook/rook-client-python/)  | 11.2k | 57.0 | 3.0 | 1.0 | Python | 940.0m |
| [ClickHouse/clickhouse-presentations](https://github.com/ClickHouse/clickhouse-presentations/)  | 11.2k | 515.0 | 39.0 | 833.0 | JavaScript | 930.0m |
| [m3db/m3msg](https://github.com/m3db/m3msg/)  | 11.2k | 62.0 | 55.0 | 15.0 | Go | 930.0m |
| [docker/engine-api](https://github.com/docker/engine-api/)  | 11.2k | 9.2k | 327.0 | 267.0 | Go | 930.0m |
| [kubernetes-sigs/kubebuilder-declar~](https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/)  | 11.1k | 591.0 | 262.0 | 188.0 | Go | 920.0m |
| [cilium/proxy](https://github.com/cilium/proxy/)  | 11.0k | 567.0 | 117.0 | 82.0 | C++ | 920.0m |
| [kubernetes-sigs/controller-tools](https://github.com/kubernetes-sigs/controller-tools/)  | 10.9k | 773.0 | 455.0 | 549.0 | Go | 910.0m |
| [operator-framework/helm-operator-p~](https://github.com/operator-framework/helm-operator-plugins/)  | 10.9k | 194.0 | 141.0 | 34.0 | Go | 910.0m |
| [jaegertracing/jaeger-client-cpp](https://github.com/jaegertracing/jaeger-client-cpp/)  | 10.9k | 110.0 | 140.0 | 136.0 | C++ | 910.0m |
| [knative/networking](https://github.com/knative/networking/)  | 10.9k | 840.0 | 728.0 | 57.0 | Go | 900.0m |
| [envoyproxy/xds-relay](https://github.com/envoyproxy/xds-relay/)  | 10.8k | 127.0 | 162.0 | 128.0 | Go | 900.0m |
| [openebs/lvm-localpv](https://github.com/openebs/lvm-localpv/)  | 10.7k | 182.0 | 144.0 | 141.0 | Go | 890.0m |
| [open-telemetry/assign-reviewers-ac~](https://github.com/open-telemetry/assign-reviewers-action/)  | 10.7k | 5.0 | 6.0 | 6.0 | JavaScript | 890.0m |
| [fluxcd/kustomize-controller](https://github.com/fluxcd/kustomize-controller/)  | 10.7k | 1.3k | 559.0 | 204.0 | Go | 890.0m |
| [kubernetes/code-generator](https://github.com/kubernetes/code-generator/)  | 10.6k | 1.1k | 17.0 | 1.4k | Go | 890.0m |
| [grafana/jmeter-to-k6](https://github.com/grafana/jmeter-to-k6/)  | 10.6k | 347.0 | 24.0 | 57.0 | JavaScript | 890.0m |
| [kubernetes/utils](https://github.com/kubernetes/utils/)  | 10.6k | 896.0 | 235.0 | 256.0 | Go | 880.0m |
| [cloudevents/sdk-csharp](https://github.com/cloudevents/sdk-csharp/)  | 10.6k | 279.0 | 146.0 | 199.0 | C# | 880.0m |
| [kubernetes/node-problem-detector](https://github.com/kubernetes/node-problem-detector/)  | 10.5k | 753.0 | 434.0 | 2.3k | Go | 880.0m |
| [grpc/grpc-kotlin](https://github.com/grpc/grpc-kotlin/)  | 10.5k | 230.0 | 184.0 | 977.0 | Kotlin | 870.0m |
| [kubernetes/cli-runtime](https://github.com/kubernetes/cli-runtime/)  | 10.5k | 913.0 | 5.0 | 240.0 | Go | 870.0m |
| [in-toto/in-toto](https://github.com/in-toto/in-toto/)  | 10.4k | 1.8k | 407.0 | 681.0 | Python | 870.0m |
| [chaos-mesh/chaosd](https://github.com/chaos-mesh/chaosd/)  | 10.4k | 179.0 | 132.0 | 103.0 | Go | 870.0m |
| [grafana/tanka](https://github.com/grafana/tanka/)  | 10.4k | 449.0 | 469.0 | 1.9k | Go | 860.0m |
| [containerd/containerd.io](https://github.com/containerd/containerd.io/)  | 10.3k | 274.0 | 143.0 | 38.0 | JavaScript | 860.0m |
| [AthenZ/authorization-proxy](https://github.com/AthenZ/authorization-proxy/)  | 10.3k | 98.0 | 12.0 | 2.0 | Go | 850.0m |
| [istio/pkg](https://github.com/istio/pkg/)  | 10.2k | 859.0 | 763.0 | 47.0 | Go | 850.0m |
| [kubernetes-sigs/kube-batch](https://github.com/kubernetes-sigs/kube-batch/)  | 10.2k | 1.3k | 691.0 | 1.0k | Go | 850.0m |
| [istio/istio.io](https://github.com/istio/istio.io/)  | 10.2k | 8.4k | 11.0k | 681.0 | JavaScript | 850.0m |
| [istio/old_mixerclient_repo](https://github.com/istio/old_mixerclient_repo/)  | 10.0k | 228.0 | 409.0 | 15.0 | C++ | 830.0m |
| [kubernetes/kubeadm](https://github.com/kubernetes/kubeadm/)  | 10.0k | 1.0k | 520.0 | 3.4k | Go | 830.0m |
| [containerd/rust-extensions](https://github.com/containerd/rust-extensions/)  | 10.0k | 294.0 | 103.0 | 92.0 | Rust | 830.0m |
| [operator-framework/rukpak](https://github.com/operator-framework/rukpak/)  | 9.8k | 435.0 | 362.0 | 37.0 | Go | 820.0m |
| [etcd-io/jetcd](https://github.com/etcd-io/jetcd/)  | 9.8k | 1.2k | 742.0 | 970.0 | Java | 820.0m |
| [containerd/ttrpc-rust](https://github.com/containerd/ttrpc-rust/)  | 9.8k | 357.0 | 151.0 | 149.0 | Rust | 820.0m |
| [kubevirt/hostpath-provisioner-oper~](https://github.com/kubevirt/hostpath-provisioner-operator/)  | 9.8k | 315.0 | 285.0 | 25.0 | Go | 810.0m |
| [kubernetes-sigs/secrets-store-csi-~](https://github.com/kubernetes-sigs/secrets-store-csi-driver/)  | 9.7k | 1.2k | 747.0 | 939.0 | Go | 810.0m |
| [fluent/fluentd-ui](https://github.com/fluent/fluentd-ui/)  | 9.7k | 1.8k | 357.0 | 580.0 | Ruby | 810.0m |
| [istio/ztunnel](https://github.com/istio/ztunnel/)  | 9.7k | 277.0 | 302.0 | 90.0 | Rust | 810.0m |
| [go-faster/jx](https://github.com/go-faster/jx/)  | 9.7k | 1.4k | 65.0 | 80.0 | Go | 810.0m |
| [dapr/js-sdk](https://github.com/dapr/js-sdk/)  | 9.6k | 585.0 | 242.0 | 159.0 | TypeScript | 800.0m |
| [cilium/kube-apate](https://github.com/cilium/kube-apate/)  | 9.6k | 18.0 | 0.0 | 4.0 | Go | 800.0m |
| [grafana/cortex-tools](https://github.com/grafana/cortex-tools/)  | 9.6k | 270.0 | 179.0 | 133.0 | Go | 800.0m |
| [theupdateframework/python-tuf](https://github.com/theupdateframework/python-tuf/)  | 9.6k | 5.6k | 1.5k | 1.5k | Python | 800.0m |
| [nats-io/jsm.go](https://github.com/nats-io/jsm.go/)  | 9.5k | 795.0 | 391.0 | 102.0 | Go | 790.0m |
| [kubernetes-sigs/cloud-provider-hua~](https://github.com/kubernetes-sigs/cloud-provider-huaweicloud/)  | 9.4k | 415.0 | 171.0 | 25.0 | Go | 790.0m |
| [containerd/nri](https://github.com/containerd/nri/)  | 9.4k | 86.0 | 28.0 | 99.0 | Go | 780.0m |
| [prometheus/statsd_exporter](https://github.com/prometheus/statsd_exporter/)  | 9.4k | 847.0 | 310.0 | 823.0 | Go | 780.0m |
| [kubernetes-sigs/gateway-api](https://github.com/kubernetes-sigs/gateway-api/)  | 9.4k | 2.3k | 1.1k | 871.0 | Go | 780.0m |
| [grafana/worldmap-panel](https://github.com/grafana/worldmap-panel/)  | 9.4k | 232.0 | 74.0 | 294.0 | JavaScript | 780.0m |
| [openebs/cstor-csi](https://github.com/openebs/cstor-csi/)  | 9.3k | 157.0 | 172.0 | 23.0 | Go | 780.0m |
| [nats-io/nats.rb](https://github.com/nats-io/nats.rb/)  | 9.3k | 756.0 | 118.0 | 867.0 | Ruby | 780.0m |
| [kubernetes-sigs/cri-tools](https://github.com/kubernetes-sigs/cri-tools/)  | 9.2k | 1.5k | 794.0 | 1.2k | Go | 770.0m |
| [openebs/jiva-operator](https://github.com/openebs/jiva-operator/)  | 9.2k | 148.0 | 170.0 | 37.0 | Go | 770.0m |
| [nats-io/nats-rest-config-proxy](https://github.com/nats-io/nats-rest-config-proxy/)  | 9.2k | 213.0 | 49.0 | 21.0 | Go | 770.0m |
| [kubernetes/kompose](https://github.com/kubernetes/kompose/)  | 9.2k | 1.5k | 895.0 | 8.3k | Go | 770.0m |
| [dapr/php-sdk](https://github.com/dapr/php-sdk/)  | 9.2k | 255.0 | 80.0 | 59.0 | PHP | 760.0m |
| [ClickHouse/clickhouse-jdbc-bridge](https://github.com/ClickHouse/clickhouse-jdbc-bridge/)  | 9.2k | 127.0 | 85.0 | 137.0 | Java | 760.0m |
| [in-toto/in-toto-golang](https://github.com/in-toto/in-toto-golang/)  | 9.1k | 721.0 | 146.0 | 77.0 | Go | 760.0m |
| [crossplane/terrajet](https://github.com/crossplane/terrajet/)  | 9.1k | 463.0 | 136.0 | 292.0 | Go | 760.0m |
| [envoyproxy/go-control-plane](https://github.com/envoyproxy/go-control-plane/)  | 9.1k | 1.4k | 398.0 | 1.3k | Go | 760.0m |
| [grpc/test-infra](https://github.com/grpc/test-infra/)  | 9.1k | 484.0 | 346.0 | 64.0 | Go | 760.0m |
| [jaegertracing/jaeger-client-node](https://github.com/jaegertracing/jaeger-client-node/)  | 9.1k | 450.0 | 420.0 | 549.0 | JavaScript | 760.0m |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-gcp/)  | 9.0k | 1.1k | 673.0 | 131.0 | Go | 750.0m |
| [volcano-sh/kubegene](https://github.com/volcano-sh/kubegene/)  | 9.0k | 101.0 | 50.0 | 184.0 | Go | 750.0m |
| [operator-framework/api](https://github.com/operator-framework/api/)  | 8.9k | 294.0 | 230.0 | 19.0 | Go | 740.0m |
| [longhorn/bot](https://github.com/longhorn/bot/)  | 8.9k | 25.0 | 20.0 | 0.0 | JavaScript | 740.0m |
| [kubevirt/kubevirt-velero-plugin](https://github.com/kubevirt/kubevirt-velero-plugin/)  | 8.9k | 123.0 | 126.0 | 18.0 | Go | 740.0m |
| [ClickHouse/libcxxabi](https://github.com/ClickHouse/libcxxabi/)  | 8.8k | 14.0 | 4.0 | 0.0 | C Header | 730.0m |
| [kubernetes-sigs/apiserver-network-~](https://github.com/kubernetes-sigs/apiserver-network-proxy/)  | 8.8k | 640.0 | 316.0 | 262.0 | Go | 730.0m |
| [kubernetes/pod-security-admission](https://github.com/kubernetes/pod-security-admission/)  | 8.8k | 465.0 | 0.0 | 77.0 | Go | 730.0m |
| [ClickHouse/graphouse](https://github.com/ClickHouse/graphouse/)  | 8.7k | 542.0 | 160.0 | 251.0 | Java | 730.0m |
| [ClickHouse/clickhouse-connect](https://github.com/ClickHouse/clickhouse-connect/)  | 8.7k | 229.0 | 84.0 | 86.0 | Python | 730.0m |
| [buildpacks/imgutil](https://github.com/buildpacks/imgutil/)  | 8.7k | 416.0 | 149.0 | 21.0 | Go | 720.0m |
| [knative/observability](https://github.com/knative/observability/)  | 8.6k | 19.0 | 28.0 | 25.0 | Go | 720.0m |
| [cockroachdb/errors](https://github.com/cockroachdb/errors/)  | 8.6k | 198.0 | 63.0 | 1.5k | Go | 720.0m |
| [kubernetes-sigs/cli-experimental](https://github.com/kubernetes-sigs/cli-experimental/)  | 8.6k | 278.0 | 200.0 | 65.0 | Go | 720.0m |
| [spiffe/tornjak](https://github.com/spiffe/tornjak/)  | 8.6k | 206.0 | 75.0 | 55.0 | TypeScript | 720.0m |
| [containernetworking/cni](https://github.com/containernetworking/cni/)  | 8.6k | 1.1k | 603.0 | 4.7k | Go | 710.0m |
| [AthenZ/garm](https://github.com/AthenZ/garm/)  | 8.5k | 126.0 | 7.0 | 3.0 | Go | 710.0m |
| [vitessio/arewefastyet](https://github.com/vitessio/arewefastyet/)  | 8.5k | 1.3k | 222.0 | 33.0 | Go | 710.0m |
| [crossplane/release-test](https://github.com/crossplane/release-test/)  | 8.5k | 1.2k | 0.0 | 3.0 | Go | 710.0m |
| [kubernetes/gengo](https://github.com/kubernetes/gengo/)  | 8.4k | 455.0 | 182.0 | 485.0 | Go | 700.0m |
| [open-policy-agent/contrib](https://github.com/open-policy-agent/contrib/)  | 8.4k | 197.0 | 167.0 | 285.0 | Go | 700.0m |
| [kubernetes-sigs/cloud-provider-bai~](https://github.com/kubernetes-sigs/cloud-provider-baiducloud/)  | 8.4k | 213.0 | 71.0 | 38.0 | Go | 700.0m |
| [nats-io/nats-kafka](https://github.com/nats-io/nats-kafka/)  | 8.3k | 98.0 | 58.0 | 91.0 | Go | 690.0m |
| [projectcontour/contour-operator](https://github.com/projectcontour/contour-operator/)  | 8.3k | 353.0 | 400.0 | 44.0 | Go | 690.0m |
| [cert-manager/approver-policy](https://github.com/cert-manager/approver-policy/)  | 8.3k | 486.0 | 193.0 | 48.0 | Go | 690.0m |
| [litmuschaos/litmus-ui](https://github.com/litmuschaos/litmus-ui/)  | 8.3k | 171.0 | 109.0 | 5.0 | TypeScript | 690.0m |
| [knative/build](https://github.com/knative/build/)  | 8.3k | 372.0 | 452.0 | 575.0 | Go | 690.0m |
| [m3db/m3em](https://github.com/m3db/m3em/)  | 8.1k | 25.0 | 19.0 | 1.0 | Go | 680.0m |
| [spiffe/spiffe.io](https://github.com/spiffe/spiffe.io/)  | 8.0k | 791.0 | 216.0 | 20.0 | JavaScript | 670.0m |
| [operator-framework/audit](https://github.com/operator-framework/audit/)  | 7.9k | 247.0 | 107.0 | 8.0 | Go | 660.0m |
| [openebs/mayastor-extensions](https://github.com/openebs/mayastor-extensions/)  | 7.9k | 220.0 | 125.0 | 10.0 | Rust | 660.0m |
| [envoyproxy/ratelimit](https://github.com/envoyproxy/ratelimit/)  | 7.9k | 150.0 | 235.0 | 1.8k | Go | 660.0m |
| [cockroachdb/roachprod](https://github.com/cockroachdb/roachprod/)  | 7.8k | 385.0 | 195.0 | 14.0 | Go | 650.0m |
| [open-telemetry/opentelemetry-go-bu~](https://github.com/open-telemetry/opentelemetry-go-build-tools/)  | 7.8k | 309.0 | 242.0 | 21.0 | Go | 650.0m |
| [cockroachdb/apd](https://github.com/cockroachdb/apd/)  | 7.8k | 324.0 | 83.0 | 516.0 | Go | 650.0m |
| [openebs/device-localpv](https://github.com/openebs/device-localpv/)  | 7.7k | 55.0 | 49.0 | 15.0 | Go | 650.0m |
| [nats-io/nats-pure.rb](https://github.com/nats-io/nats-pure.rb/)  | 7.7k | 304.0 | 60.0 | 95.0 | Ruby | 640.0m |
| [kubernetes-sigs/krew](https://github.com/kubernetes-sigs/krew/)  | 7.6k | 478.0 | 479.0 | 5.4k | Go | 640.0m |
| [kedacore/http-add-on](https://github.com/kedacore/http-add-on/)  | 7.6k | 262.0 | 413.0 | 176.0 | Go | 640.0m |
| [prometheus/client_python](https://github.com/prometheus/client_python/)  | 7.6k | 512.0 | 389.0 | 3.3k | Python | 630.0m |
| [grafana/grafana-polystat-panel](https://github.com/grafana/grafana-polystat-panel/)  | 7.6k | 199.0 | 124.0 | 74.0 | TypeScript | 630.0m |
| [fluent/fluent-plugin-opensearch](https://github.com/fluent/fluent-plugin-opensearch/)  | 7.6k | 1.3k | 29.0 | 39.0 | Ruby | 630.0m |
| [kubernetes/kube-aggregator](https://github.com/kubernetes/kube-aggregator/)  | 7.5k | 2.1k | 11.0 | 223.0 | Go | 620.0m |
| [grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client/)  | 7.5k | 402.0 | 136.0 | 71.0 | Go | 620.0m |
| [notaryproject/notation-go](https://github.com/notaryproject/notation-go/)  | 7.4k | 124.0 | 198.0 | 18.0 | Go | 610.0m |
| [dapr/go-sdk](https://github.com/dapr/go-sdk/)  | 7.4k | 239.0 | 257.0 | 387.0 | Go | 610.0m |
| [kubernetes-sigs/aws-iam-authentica~](https://github.com/kubernetes-sigs/aws-iam-authenticator/)  | 7.4k | 515.0 | 313.0 | 2.0k | Go | 610.0m |
| [kubernetes-sigs/aws-efs-csi-driver](https://github.com/kubernetes-sigs/aws-efs-csi-driver/)  | 7.3k | 689.0 | 585.0 | 568.0 | Go | 610.0m |
| [openebs/upgrade](https://github.com/openebs/upgrade/)  | 7.2k | 122.0 | 135.0 | 10.0 | Go | 600.0m |
| [grafana/postman-to-k6](https://github.com/grafana/postman-to-k6/)  | 7.2k | 607.0 | 58.0 | 268.0 | JavaScript | 600.0m |
| [kubernetes-sigs/instrumentation-to~](https://github.com/kubernetes-sigs/instrumentation-tools/)  | 7.2k | 92.0 | 6.0 | 25.0 | Go | 600.0m |
| [kubevirt/client-python](https://github.com/kubevirt/client-python/)  | 7.2k | 358.0 | 21.0 | 24.0 | Python | 600.0m |
| [docker/kitematic](https://github.com/docker/kitematic/)  | 7.1k | 2.3k | 534.0 | 12.3k | JavaScript | 590.0m |
| [grpc/grpc-web](https://github.com/grpc/grpc-web/)  | 7.1k | 888.0 | 622.0 | 7.4k | JavaScript | 590.0m |
| [fluxcd/notification-controller](https://github.com/fluxcd/notification-controller/)  | 7.1k | 829.0 | 349.0 | 119.0 | Go | 590.0m |
| [kubernetes/dns](https://github.com/kubernetes/dns/)  | 7.0k | 710.0 | 318.0 | 803.0 | Go | 590.0m |
| [grafana/opcua-datasource](https://github.com/grafana/opcua-datasource/)  | 7.0k | 370.0 | 44.0 | 45.0 | TypeScript | 590.0m |
| [prometheus/mysqld_exporter](https://github.com/prometheus/mysqld_exporter/)  | 7.0k | 605.0 | 412.0 | 1.7k | Go | 580.0m |
| [notaryproject/notation](https://github.com/notaryproject/notation/)  | 7.0k | 192.0 | 343.0 | 144.0 | Go | 580.0m |
| [openebs/libcstor](https://github.com/openebs/libcstor/)  | 7.0k | 77.0 | 89.0 | 14.0 | C | 580.0m |
| [kubeedge/edgemesh](https://github.com/kubeedge/edgemesh/)  | 6.9k | 534.0 | 265.0 | 182.0 | Go | 580.0m |
| [open-telemetry/opentelemetry-demo](https://github.com/open-telemetry/opentelemetry-demo/)  | 6.9k | 422.0 | 551.0 | 691.0 | TypeScript | 570.0m |
| [cilium/hubble](https://github.com/cilium/hubble/)  | 6.8k | 796.0 | 757.0 | 2.5k | Go | 570.0m |
| [containerd/cgroups](https://github.com/containerd/cgroups/)  | 6.8k | 517.0 | 218.0 | 928.0 | Go | 570.0m |
| [thanos-io/objstore](https://github.com/thanos-io/objstore/)  | 6.8k | 209.0 | 34.0 | 43.0 | Go | 570.0m |
| [grafana/attic](https://github.com/grafana/attic/)  | 6.8k | 426.0 | 0.0 | 1.0 | JavaScript | 570.0m |
| [envoyproxy/envoy-build-tools](https://github.com/envoyproxy/envoy-build-tools/)  | 6.8k | 320.0 | 182.0 | 36.0 | Bazel | 570.0m |
| [kubernetes-sigs/reference-docs](https://github.com/kubernetes-sigs/reference-docs/)  | 6.8k | 492.0 | 233.0 | 65.0 | Go | 570.0m |
| [nats-io/nats-operator](https://github.com/nats-io/nats-operator/)  | 6.8k | 588.0 | 206.0 | 555.0 | Go | 570.0m |
| [kubevirt/terraform-provider-kubevi~](https://github.com/kubevirt/terraform-provider-kubevirt/)  | 6.8k | 101.0 | 26.0 | 16.0 | Go | 570.0m |
| [grafana/kairosdb-datasource](https://github.com/grafana/kairosdb-datasource/)  | 6.8k | 110.0 | 43.0 | 31.0 | TypeScript | 560.0m |
| [cilium/image-tools](https://github.com/cilium/image-tools/)  | 6.7k | 239.0 | 192.0 | 11.0 | Perl | 560.0m |
| [argoproj/notifications-engine](https://github.com/argoproj/notifications-engine/)  | 6.7k | 81.0 | 108.0 | 180.0 | Go | 560.0m |
| [in-toto/in-toto-rs](https://github.com/in-toto/in-toto-rs/)  | 6.7k | 161.0 | 37.0 | 28.0 | Rust | 560.0m |
| [dapr/kit](https://github.com/dapr/kit/)  | 6.7k | 30.0 | 35.0 | 10.0 | Go | 550.0m |
| [helm/chartmuseum](https://github.com/helm/chartmuseum/)  | 6.7k | 613.0 | 333.0 | 3.1k | Go | 550.0m |
| [kubevirt/kvm-info-nfd-plugin](https://github.com/kubevirt/kvm-info-nfd-plugin/)  | 6.6k | 62.0 | 13.0 | 1.0 | C Header | 550.0m |
| [ydb-platform/ydb-nodejs-sdk](https://github.com/ydb-platform/ydb-nodejs-sdk/)  | 6.6k | 367.0 | 185.0 | 46.0 | TypeScript | 550.0m |
| [etcd-io/dbtester](https://github.com/etcd-io/dbtester/)  | 6.6k | 1.2k | 288.0 | 250.0 | Go | 550.0m |
| [argoproj/argo-cd-ui](https://github.com/argoproj/argo-cd-ui/)  | 6.6k | 385.0 | 131.0 | 14.0 | TypeScript | 550.0m |
| [kubernetes-sigs/apiserver-builder-~](https://github.com/kubernetes-sigs/apiserver-builder-alpha/)  | 6.6k | 1.1k | 434.0 | 717.0 | Go | 550.0m |
| [ydb-platform/ydb-go-examples](https://github.com/ydb-platform/ydb-go-examples/)  | 6.5k | 176.0 | 16.0 | 12.0 | Go | 540.0m |
| [envoyproxy/envoy-perf](https://github.com/envoyproxy/envoy-perf/)  | 6.5k | 142.0 | 160.0 | 107.0 | Python | 540.0m |
| [crossplane/gitlab-controller](https://github.com/crossplane/gitlab-controller/)  | 6.5k | 45.0 | 18.0 | 12.0 | Go | 540.0m |
| [AthenZ/terraform-provider-athenz](https://github.com/AthenZ/terraform-provider-athenz/)  | 6.4k | 94.0 | 56.0 | 3.0 | Go | 530.0m |
| [gotd/botapi](https://github.com/gotd/botapi/)  | 6.4k | 653.0 | 313.0 | 18.0 | Go | 530.0m |
| [nats-io/nats.ts](https://github.com/nats-io/nats.ts/)  | 6.4k | 533.0 | 86.0 | 180.0 | TypeScript | 530.0m |
| [cockroachdb/activerecord-cockroach~](https://github.com/cockroachdb/activerecord-cockroachdb-adapter/)  | 6.4k | 338.0 | 172.0 | 92.0 | Ruby | 530.0m |
| [kubernetes/csi-translation-lib](https://github.com/kubernetes/csi-translation-lib/)  | 6.4k | 658.0 | 0.0 | 10.0 | Go | 530.0m |
| [open-telemetry/opamp-go](https://github.com/open-telemetry/opamp-go/)  | 6.3k | 92.0 | 110.0 | 60.0 | Go | 530.0m |
| [kubernetes-sigs/kubetest2](https://github.com/kubernetes-sigs/kubetest2/)  | 6.3k | 535.0 | 177.0 | 259.0 | Go | 530.0m |
| [jaegertracing/jaeger-client-python](https://github.com/jaegertracing/jaeger-client-python/)  | 6.3k | 264.0 | 236.0 | 413.0 | Python | 520.0m |
| [containerd/continuity](https://github.com/containerd/continuity/)  | 6.2k | 372.0 | 157.0 | 126.0 | Go | 520.0m |
| [kubernetes-sigs/sig-storage-local-~](https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner/)  | 6.1k | 623.0 | 213.0 | 866.0 | Go | 510.0m |
| [kubernetes-sigs/slack-infra](https://github.com/kubernetes-sigs/slack-infra/)  | 6.1k | 108.0 | 44.0 | 85.0 | Go | 510.0m |
| [nats-io/java-nats-examples](https://github.com/nats-io/java-nats-examples/)  | 6.1k | 166.0 | 7.0 | 20.0 | Java | 500.0m |
| [kubevirt/hostpath-provisioner](https://github.com/kubevirt/hostpath-provisioner/)  | 6.0k | 149.0 | 164.0 | 51.0 | Go | 500.0m |
| [cubeFS/cubefs-for-android](https://github.com/cubeFS/cubefs-for-android/)  | 6.0k | 4.0 | 0.0 | 17.0 | Go | 500.0m |
| [notaryproject/notation-core-go](https://github.com/notaryproject/notation-core-go/)  | 6.0k | 42.0 | 99.0 | 5.0 | Go | 500.0m |
| [cilium/docsearch-scraper-webhook](https://github.com/cilium/docsearch-scraper-webhook/)  | 6.0k | 38.0 | 32.0 | 3.0 | Python | 500.0m |
| [kubeedge/examples](https://github.com/kubeedge/examples/)  | 6.0k | 132.0 | 64.0 | 201.0 | Go | 500.0m |
| [open-policy-agent/conftest](https://github.com/open-policy-agent/conftest/)  | 6.0k | 868.0 | 505.0 | 2.5k | Go | 500.0m |
| [siderolabs/kres](https://github.com/siderolabs/kres/)  | 6.0k | 152.0 | 199.0 | 22.0 | Go | 500.0m |
| [nats-io/stan.java](https://github.com/nats-io/stan.java/)  | 6.0k | 361.0 | 64.0 | 94.0 | Java | 500.0m |
| [ClickHouse/bzip2](https://github.com/ClickHouse/bzip2/)  | 6.0k | 141.0 | 0.0 | 1.0 | C | 500.0m |
| [kubevirt/client-go](https://github.com/kubevirt/client-go/)  | 6.0k | 301.0 | 9.0 | 25.0 | Go | 500.0m |
| [prometheus/blackbox_exporter](https://github.com/prometheus/blackbox_exporter/)  | 5.9k | 455.0 | 436.0 | 3.5k | Go | 500.0m |
| [grafana/devtools](https://github.com/grafana/devtools/)  | 5.9k | 139.0 | 22.0 | 9.0 | Go | 500.0m |
| [litmuschaos/charthub.litmuschaos.io](https://github.com/litmuschaos/charthub.litmuschaos.io/)  | 5.9k | 225.0 | 317.0 | 14.0 | TypeScript | 490.0m |
| [theupdateframework/tuf-js](https://github.com/theupdateframework/tuf-js/)  | 5.9k | 301.0 | 226.0 | 59.0 | TypeScript | 490.0m |
| [docker/go](https://github.com/docker/go/)  | 5.9k | 24.0 | 8.0 | 17.0 | Go | 490.0m |
| [kata-containers/govmm](https://github.com/kata-containers/govmm/)  | 5.9k | 402.0 | 151.0 | 304.0 | Go | 490.0m |
| [chaos-mesh/k8s_dns_chaos](https://github.com/chaos-mesh/k8s_dns_chaos/)  | 5.9k | 35.0 | 17.0 | 10.0 | Go | 490.0m |
| [docker/volumes-backup-extension](https://github.com/docker/volumes-backup-extension/)  | 5.9k | 228.0 | 76.0 | 47.0 | Go | 490.0m |
| [istio/api](https://github.com/istio/api/)  | 5.8k | 1.7k | 2.5k | 404.0 | Protocol Buffers | 480.0m |
| [kubernetes/website](https://github.com/kubernetes/website/)  | 5.8k | 39.1k | 29.1k | 3.7k | JavaScript | 480.0m |
| [docker/metadata-action](https://github.com/docker/metadata-action/)  | 5.7k | 318.0 | 157.0 | 563.0 | TypeScript | 480.0m |
| [dapr/cpp-sdk](https://github.com/dapr/cpp-sdk/)  | 5.7k | 45.0 | 27.0 | 29.0 | C Header | 470.0m |
| [cloudevents/sdk-rust](https://github.com/cloudevents/sdk-rust/)  | 5.7k | 194.0 | 148.0 | 125.0 | Rust | 470.0m |
| [grafana/azure-monitor-datasource](https://github.com/grafana/azure-monitor-datasource/)  | 5.6k | 170.0 | 18.0 | 91.0 | TypeScript | 470.0m |
| [fluxcd/helm-operator](https://github.com/fluxcd/helm-operator/)  | 5.6k | 4.6k | 351.0 | 656.0 | Go | 470.0m |
| [prometheus/compliance](https://github.com/prometheus/compliance/)  | 5.6k | 147.0 | 79.0 | 105.0 | Go | 470.0m |
| [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/)  | 5.6k | 989.0 | 602.0 | 4.5k | Go | 460.0m |
| [kubernetes/component-helpers](https://github.com/kubernetes/component-helpers/)  | 5.5k | 329.0 | 0.0 | 11.0 | Go | 460.0m |
| [m3db/m3storage](https://github.com/m3db/m3storage/)  | 5.5k | 38.0 | 15.0 | 3.0 | Go | 460.0m |
| [fluxcd/helm-controller](https://github.com/fluxcd/helm-controller/)  | 5.5k | 897.0 | 412.0 | 293.0 | Go | 460.0m |
| [fluent/fluentd-docs](https://github.com/fluent/fluentd-docs/)  | 5.5k | 2.4k | 458.0 | 49.0 | JavaScript | 450.0m |
| [longhorn/backing-image-manager](https://github.com/longhorn/backing-image-manager/)  | 5.5k | 126.0 | 105.0 | 4.0 | Go | 450.0m |
| [spiffe/spire-controller-manager](https://github.com/spiffe/spire-controller-manager/)  | 5.4k | 84.0 | 92.0 | 29.0 | Go | 450.0m |
| [openebs/api](https://github.com/openebs/api/)  | 5.4k | 88.0 | 99.0 | 7.0 | Go | 450.0m |
| [kubernetes-sigs/cluster-addons](https://github.com/kubernetes-sigs/cluster-addons/)  | 5.4k | 261.0 | 89.0 | 147.0 | Go | 450.0m |
| [nats-io/nats.py2](https://github.com/nats-io/nats.py2/)  | 5.4k | 320.0 | 50.0 | 62.0 | Python | 450.0m |
| [prometheus/codemirror-promql](https://github.com/prometheus/codemirror-promql/)  | 5.3k | 516.0 | 151.0 | 35.0 | TypeScript | 440.0m |
| [VictoriaMetrics/metricsql](https://github.com/VictoriaMetrics/metricsql/)  | 5.3k | 102.0 | 6.0 | 136.0 | Go | 440.0m |
| [cert-manager/release](https://github.com/cert-manager/release/)  | 5.3k | 246.0 | 108.0 | 4.0 | Go | 440.0m |
| [litmuschaos/website-litmuschaos](https://github.com/litmuschaos/website-litmuschaos/)  | 5.3k | 133.0 | 165.0 | 9.0 | TypeScript | 440.0m |
| [kubernetes-sigs/aws-fsx-csi-driver](https://github.com/kubernetes-sigs/aws-fsx-csi-driver/)  | 5.2k | 409.0 | 204.0 | 103.0 | Go | 440.0m |
| [nats-io/nats-account-server](https://github.com/nats-io/nats-account-server/)  | 5.2k | 209.0 | 60.0 | 69.0 | Go | 440.0m |
| [grafana/kubernetes-app](https://github.com/grafana/kubernetes-app/)  | 5.2k | 199.0 | 23.0 | 397.0 | TypeScript | 440.0m |
| [litmuschaos/chaos-runner](https://github.com/litmuschaos/chaos-runner/)  | 5.2k | 136.0 | 181.0 | 19.0 | Go | 440.0m |
| [helm/helm-classic](https://github.com/helm/helm-classic/)  | 5.2k | 574.0 | 274.0 | 578.0 | Go | 430.0m |
| [goharbor/terraform-provider-harbor](https://github.com/goharbor/terraform-provider-harbor/)  | 5.2k | 520.0 | 185.0 | 87.0 | Go | 430.0m |
| [k3s-io/kine](https://github.com/k3s-io/kine/)  | 5.2k | 139.0 | 102.0 | 934.0 | Go | 430.0m |
| [etcd-io/zetcd](https://github.com/etcd-io/zetcd/)  | 5.2k | 135.0 | 58.0 | 1.1k | Go | 430.0m |
| [fluxcd/image-reflector-controller](https://github.com/fluxcd/image-reflector-controller/)  | 5.1k | 583.0 | 253.0 | 79.0 | Go | 430.0m |
| [fluent/chunkio](https://github.com/fluent/chunkio/)  | 5.1k | 371.0 | 88.0 | 16.0 | C | 420.0m |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-digitalocean/)  | 5.1k | 399.0 | 383.0 | 92.0 | Go | 420.0m |
| [longhorn/backupstore](https://github.com/longhorn/backupstore/)  | 5.1k | 188.0 | 107.0 | 3.0 | Go | 420.0m |
| [coredns/policy](https://github.com/coredns/policy/)  | 5.0k | 77.0 | 30.0 | 46.0 | Go | 420.0m |
| [ydb-platform/terraform-provider-ydb](https://github.com/ydb-platform/terraform-provider-ydb/)  | 5.0k | 70.0 | 10.0 | 1.0 | Go | 420.0m |
| [m3db/m3ctl](https://github.com/m3db/m3ctl/)  | 5.0k | 82.0 | 64.0 | 14.0 | Go | 420.0m |
| [fluxcd/webui](https://github.com/fluxcd/webui/)  | 4.9k | 162.0 | 38.0 | 142.0 | TypeScript | 410.0m |
| [grafana/gel-app](https://github.com/grafana/gel-app/)  | 4.9k | 119.0 | 47.0 | 2.0 | Go | 410.0m |
| [kubeedge/ianvs](https://github.com/kubeedge/ianvs/)  | 4.9k | 155.0 | 53.0 | 76.0 | Python | 400.0m |
| [goharbor/harbor-cluster-operator](https://github.com/goharbor/harbor-cluster-operator/)  | 4.8k | 390.0 | 114.0 | 28.0 | Go | 400.0m |
| [fluent/cfl](https://github.com/fluent/cfl/)  | 4.8k | 88.0 | 25.0 | 0.0 | C Header | 400.0m |
| [dapr/setup-dapr](https://github.com/dapr/setup-dapr/)  | 4.8k | 14.0 | 114.0 | 7.0 | JavaScript | 400.0m |
| [kyverno/policy-reporter-kyverno-pl~](https://github.com/kyverno/policy-reporter-kyverno-plugin/)  | 4.8k | 56.0 | 20.0 | 8.0 | Go | 400.0m |
| [prometheus/snmp_exporter](https://github.com/prometheus/snmp_exporter/)  | 4.8k | 553.0 | 416.0 | 1.2k | Go | 400.0m |
| [nats-io/nats-replicator](https://github.com/nats-io/nats-replicator/)  | 4.8k | 44.0 | 7.0 | 18.0 | Go | 400.0m |
| [longhorn/longhorn-instance-manager](https://github.com/longhorn/longhorn-instance-manager/)  | 4.8k | 366.0 | 195.0 | 8.0 | Go | 400.0m |
| [helm/monocular](https://github.com/helm/monocular/)  | 4.7k | 399.0 | 362.0 | 1.4k | Go | 400.0m |
| [kata-containers/tests](https://github.com/kata-containers/tests/)  | 4.7k | 4.3k | 2.8k | 135.0 | Go | 390.0m |
| [gotd/contrib](https://github.com/gotd/contrib/)  | 4.7k | 568.0 | 317.0 | 10.0 | Go | 390.0m |
| [containerssh/auditlog](https://github.com/containerssh/auditlog/)  | 4.7k | 118.0 | 476.0 | 1.0 | Go | 390.0m |
| [knative/client-contrib](https://github.com/knative/client-contrib/)  | 4.7k | 36.0 | 45.0 | 11.0 | Go | 390.0m |
| [argoproj/argo-ui](https://github.com/argoproj/argo-ui/)  | 4.6k | 409.0 | 302.0 | 176.0 | TypeScript | 390.0m |
| [kubernetes-sigs/image-builder](https://github.com/kubernetes-sigs/image-builder/)  | 4.6k | 2.6k | 815.0 | 248.0 | Go | 390.0m |
| [grafana/grafana-sensu-app](https://github.com/grafana/grafana-sensu-app/)  | 4.6k | 16.0 | 45.0 | 8.0 | TypeScript | 380.0m |
| [chaos-mesh/toda](https://github.com/chaos-mesh/toda/)  | 4.6k | 233.0 | 19.0 | 41.0 | Rust | 380.0m |
| [containerd/runwasi](https://github.com/containerd/runwasi/)  | 4.6k | 186.0 | 62.0 | 541.0 | Rust | 380.0m |
| [cloudevents/sdk-python](https://github.com/cloudevents/sdk-python/)  | 4.6k | 83.0 | 123.0 | 188.0 | Python | 380.0m |
| [nats-io/stan.net](https://github.com/nats-io/stan.net/)  | 4.6k | 190.0 | 114.0 | 134.0 | C# | 380.0m |
| [istio/old_auth_repo](https://github.com/istio/old_auth_repo/)  | 4.6k | 166.0 | 230.0 | 73.0 | Go | 380.0m |
| [openebs/velero-plugin](https://github.com/openebs/velero-plugin/)  | 4.5k | 119.0 | 137.0 | 43.0 | Go | 380.0m |
| [cert-manager/trust-manager](https://github.com/cert-manager/trust-manager/)  | 4.5k | 263.0 | 82.0 | 101.0 | Go | 380.0m |


### CNCF Projects
| Project | SLOC | Commits | PRs | Stars | Language | Effort |
|---------|------|---------|-----|-------|----------|--------|
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1k | 190.5k | 243.1k | Go | 286.2 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.0M | 117.4k | 71.0k | 98.3k | Go | 166.5 |
| [open-telemetry](https://github.com/open-telemetry)  | 1.8M | 61.9k | 61.0k | 28.6k | Go | 148.4 |
| [grpc](https://github.com/grpc)  | 1.7M | 74.1k | 38.8k | 86.8k | C++ | 140.3 |
| [envoyproxy](https://github.com/envoyproxy)  | 1.6M | 46.3k | 24.1k | 28.1k | C++ | 131.1 |
| [nats-io](https://github.com/nats-io)  | 802.5k | 36.2k | 12.9k | 31.0k | Go | 66.9 |
| [vitessio](https://github.com/vitessio)  | 671.3k | 36.7k | 11.0k | 15.9k | Go | 56.0 |
| [fluent](https://github.com/fluent)  | 639.2k | 38.0k | 12.0k | 24.2k | C | 53.3 |
| [kubevirt](https://github.com/kubevirt)  | 631.8k | 60.5k | 22.6k | 5.5k | Go | 52.6 |
| [cubeFS](https://github.com/cubeFS)  | 591.2k | 2.8k | 1.4k | 3.4k | Go | 49.3 |
| [istio](https://github.com/istio)  | 583.6k | 47.1k | 62.4k | 38.5k | Go | 48.6 |
| [argoproj](https://github.com/argoproj)  | 517.8k | 18.1k | 16.9k | 34.2k | Go | 43.1 |
| [knative](https://github.com/knative)  | 432.8k | 30.6k | 32.3k | 12.4k | Go | 36.1 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762.0 | Java | 33.8 |
| [backstage](https://github.com/backstage)  | 402.7k | 40.4k | 12.9k | 21.8k | TypeScript | 33.6 |
| [prometheus](https://github.com/prometheus)  | 377.3k | 30.9k | 19.1k | 89.5k | Go | 31.4 |
| [containerd](https://github.com/containerd)  | 353.4k | 23.2k | 11.7k | 24.7k | Go | 29.4 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.0k | 8.5k | 7.2k | 15.4k | Go | 28.6 |
| [dapr](https://github.com/dapr)  | 335.5k | 17.4k | 10.7k | 26.0k | Go | 28.0 |
| [kyverno](https://github.com/kyverno)  | 280.0k | 8.8k | 5.1k | 4.0k | Go | 23.3 |
| [operator-framework](https://github.com/operator-framework)  | 265.0k | 18.9k | 14.0k | 13.3k | Go | 22.1 |
| [keptn](https://github.com/keptn)  | 256.8k | 11.8k | 8.4k | 1.9k | Go | 21.4 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.5k | 8.7k | 7.6k | 23.7k | Go | 20.9 |
| [goharbor](https://github.com/goharbor)  | 242.5k | 16.9k | 10.4k | 21.1k | Go | 20.2 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 13.0k | 13.4k | 18.0k | Rust | 19.8 |
| [spiffe](https://github.com/spiffe)  | 206.9k | 8.6k | 4.0k | 2.9k | Go | 17.2 |
| [litmuschaos](https://github.com/litmuschaos)  | 203.2k | 7.5k | 7.1k | 4.1k | TypeScript | 16.9 |
| [fluxcd](https://github.com/fluxcd)  | 200.2k | 27.7k | 9.5k | 19.9k | Go | 16.7 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4k | 4.5k | 4.6k | 4.8k | Python | 16.4 |
| [longhorn](https://github.com/longhorn)  | 193.7k | 10.0k | 6.6k | 5.1k | Go | 16.1 |
| [etcd-io](https://github.com/etcd-io)  | 185.3k | 26.4k | 11.3k | 52.3k | Go | 15.4 |
| [cert-manager](https://github.com/cert-manager)  | 145.5k | 13.1k | 4.8k | 10.9k | Go | 12.1 |
| [kubeedge](https://github.com/kubeedge)  | 141.3k | 7.3k | 3.9k | 6.8k | Go | 11.8 |
| [rook](https://github.com/rook)  | 132.1k | 24.8k | 7.1k | 11.2k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 4.9k | 3.8k | 4.1k | Go | 10.9 |
| [crossplane](https://github.com/crossplane)  | 127.7k | 11.2k | 3.5k | 7.7k | Go | 10.6 |
| [thanos-io](https://github.com/thanos-io)  | 123.5k | 4.1k | 4.0k | 12.0k | Go | 10.3 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 4.9k | 3.8k | 5.1k | Go | 10.0 |
| [buildpacks](https://github.com/buildpacks)  | 116.7k | 16.1k | 3.5k | 2.8k | Go | 9.7 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 106.8k | 3.3k | 3.2k | 5.9k | Go | 8.9 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 105.7k | 18.0k | 3.3k | 4.0k | Python | 8.8 |
| [kedacore](https://github.com/kedacore)  | 105.0k | 3.9k | 4.2k | 6.8k | Go | 8.8 |
| [karmada-io](https://github.com/karmada-io)  | 98.9k | 4.1k | 2.7k | 3.2k | Go | 8.2 |
| [in-toto](https://github.com/in-toto)  | 92.3k | 9.5k | 1.0k | 1.1k | Python | 7.7 |
| [cloudevents](https://github.com/cloudevents)  | 79.5k | 4.0k | 2.6k | 5.8k | Go | 6.6 |
| [coredns](https://github.com/coredns)  | 76.2k | 5.5k | 4.7k | 11.5k | Go | 6.3 |
| [helm](https://github.com/helm)  | 71.5k | 28.9k | 26.9k | 48.8k | Go | 6.0 |
| [cri-o](https://github.com/cri-o)  | 71.3k | 9.2k | 5.8k | 4.5k | Go | 5.9 |
| [volcano-sh](https://github.com/volcano-sh)  | 68.5k | 5.5k | 2.0k | 3.2k | Go | 5.7 |
| [notaryproject](https://github.com/notaryproject)  | 64.4k | 3.7k | 1.9k | 3.3k | Go | 5.4 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 1.8k | 2.2k | 2.2k | Go | 5.2 |
| [theupdateframework](https://github.com/theupdateframework)  | 45.4k | 8.6k | 2.6k | 2.6k | Rust | 3.8 |
| [k3s-io](https://github.com/k3s-io)  | 38.5k | 3.5k | 3.3k | 25.3k | Go | 3.2 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 2.8k | 1.2k | 6.5k | Go | 2.9 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 749.0 | 118.0 | 588.0 | Go | 1.1 |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |


### OTEL Instrumentations
| Project | SLOC | Commits | PRs | Stars | Effort |
|---------|------|---------|-----|-------|--------|
| [**OTEL**](https://opentelemetry.io/docs/instrumentation/)  | 1.1M | 32.5k | 30.8k | 16.7k | 89.23 |
| [java-instrumentation](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.9k | 8.8k | 5.1k | 1.2k | 13.99 |
| [dotnet-instrumentation](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.1k | 929.0 | 1.8k | 224.0 | 11.76 |
| [java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 3.3k | 3.6k | 1.5k | 8.03 |
| [js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 1.7k | 2.0k | 1.8k | 5.08 |
| [dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9k | 2.3k | 2.8k | 2.3k | 4.74 |
| [js-contrib](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 1.6k | 990.0 | 420.0 | 4.28 |
| [go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 1.9k | 2.4k | 3.7k | 4.26 |
| [python-contrib](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.8k | 1.8k | 1.0k | 413.0 | 3.82 |
| [python](https://github.com/open-telemetry/opentelemetry-python/)  | 42.2k | 1.4k | 1.7k | 1.2k | 3.52 |
| [cpp-contrib](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42.1k | 157.0 | 180.0 | 78.0 | 3.5 |
| [ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4k | 268.0 | 112.0 | 85.0 | 3.45 |
| [cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.6k | 1.0k | 1.2k | 432.0 | 2.8 |
| [dotnet-contrib](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.4k | 732.0 | 850.0 | 230.0 | 2.78 |
| [php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 508.0 | 591.0 | 519.0 | 2.58 |
| [rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 574.0 | 597.0 | 997.0 | 2.54 |
| [go-contrib](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.3k | 1.0k | 3.2k | 684.0 | 2.27 |
| [java-contrib](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26.3k | 562.0 | 644.0 | 80.0 | 2.2 |
| [swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9k | 618.0 | 252.0 | 126.0 | 2.16 |
| [ruby-contrib](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 20.0k | 895.0 | 241.0 | 29.0 | 1.66 |
| [ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16.4k | 763.0 | 919.0 | 375.0 | 1.37 |
| [php-contrib](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12.8k | 322.0 | 112.0 | 22.0 | 1.06 |
| [erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11.5k | 1.1k | 355.0 | 253.0 | 0.96 |
| [go-instrumentation](https://github.com/open-telemetry/opentelemetry-go-instrumentation/)  | 2.9k | 23.0 | 23.0 | 54.0 | 0.24 |
| [erlang-contrib](https://github.com/open-telemetry/opentelemetry-erlang-contrib/)  | 1.5k | 96.0 | 117.0 | 73.0 | 0.12 |
| [php-instrumentation](https://github.com/open-telemetry/opentelemetry-php-instrumentation/)  | 556.0 | 30.0 | 30.0 | 24.0 | 0.05 |

