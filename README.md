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
- Direct clone of GitHub repos (~140GB of data)
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
137G	_work
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
| [**CNCF**](https://www.cncf.io/)  | 53.4M | 3.5M | 2.1M | 2.3M | Go | 4.5k |
| [torvalds](https://github.com/torvalds)  | 19.1M | 1.4M | 1.3k | 204.2k | C | 1.6k |
| [apache](https://github.com/apache)  | 15.7M | 653.6k | 71.8k | 103.6k | Java | 1.3k |
| [ydb-platform](https://github.com/ydb-platform)  | 7.3M | 80.4k | 32.2k | 5.8k | C++ | 610.8 |
| [**K8s**](https://kubernetes.io/)  | 5.5M | 679.9k | 374.3k | 438.6k | Go | 462.3 |
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 489.9k | 248.2k | 301.2k | Go | 286.8 |
| [ytsaurus](https://github.com/ytsaurus)  | 3.3M | 94.1k | 2.6k | 2.2k | C++ | 272.2 |
| [tensorflow](https://github.com/tensorflow)  | 3.3M | 185.4k | 58.8k | 191.9k | C++ | 271.2 |
| [elastic](https://github.com/elastic)  | 3.1M | 121.1k | 141.1k | 102.1k | Java | 260.9 |
| [envoyproxy](https://github.com/envoyproxy)  | 2.9M | 100.9k | 46.5k | 37.1k | JavaScript | 238.6 |
| [open-telemetry](https://github.com/open-telemetry)  | 2.3M | 128.4k | 138.2k | 60.4k | Go | 191.3 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.1M | 190.0k | 126.2k | 137.5k | Go | 175.5 |
| [godotengine](https://github.com/godotengine)  | 2.1M | 93.3k | 59.1k | 119.2k | C++ | 171.0 |
| [cockroachdb](https://github.com/cockroachdb)  | 2.0M | 138.5k | 94.2k | 41.7k | Go | 163.1 |
| [grpc](https://github.com/grpc)  | 1.7M | 87.0k | 52.3k | 104.5k | C++ | 145.0 |
| [grafana](https://github.com/grafana)  | 1.7M | 119.6k | 103.9k | 142.7k | Go | 142.0 |
| [ClickHouse](https://github.com/ClickHouse)  | 1.6M | 208.5k | 66.7k | 52.8k | C++ | 137.5 |
| [konveyor](https://github.com/konveyor)  | 1.6M | 14.3k | 14.6k | 1.2k | JavaScript | 136.4 |
| [inclavare-containers](https://github.com/inclavare-containers)  | 1.6M | 25.6k | 1.3k | 748 | C++ | 131.7 |
| [rust-lang](https://github.com/rust-lang)  | 1.5M | 307.2k | 87.1k | 106.9k | Rust | 126.9 |
| [openkruise](https://github.com/openkruise)  | 1.4M | 2.2k | 2.5k | 6.5k | JavaScript | 113.0 |
| [docker](https://github.com/docker)  | 1.3M | 197.9k | 45.2k | 195.0k | Go | 110.7 |
| [python](https://github.com/python)  | 1.3M | 128.6k | 64.9k | 69.2k | Python | 107.1 |
| [netdata](https://github.com/netdata)  | 1.2M | 35.9k | 19.5k | 77.3k | JavaScript | 103.5 |
| [prestodb](https://github.com/prestodb)  | 1.2M | 26.0k | 15.3k | 15.7k | Java | 97.0 |
| [kubewarden](https://github.com/kubewarden)  | 1.1M | 32.9k | 15.4k | 804 | JavaScript | 92.4 |
| [golang](https://github.com/golang)  | 1.1M | 64.0k | 4.7k | 130.1k | Go | 88.3 |
| [nodejs](https://github.com/nodejs)  | 955.2k | 45.2k | 38.1k | 113.6k | JavaScript | 79.6 |
| [kubeflow](https://github.com/kubeflow)  | 936.3k | 26.8k | 32.2k | 33.7k | Go | 78.0 |
| [tarantool](https://github.com/tarantool)  | 908.3k | 31.0k | 7.6k | 5.5k | Lua | 75.7 |
| [ziglang](https://github.com/ziglang)  | 900.7k | 35.2k | 12.0k | 41.5k | Zig | 75.1 |
| [brigadecore](https://github.com/brigadecore)  | 872.9k | 6.3k | 2.8k | 3.0k | JavaScript | 72.7 |
| [tikv](https://github.com/tikv)  | 862.5k | 20.1k | 26.7k | 30.1k | Rust | 71.9 |
| [nats-io](https://github.com/nats-io)  | 858.5k | 50.2k | 20.8k | 44.9k | Go | 71.5 |
| [keycloak](https://github.com/keycloak)  | 849.2k | 44.6k | 31.8k | 38.0k | Java | 70.8 |
| [cdk8s-team](https://github.com/cdk8s-team)  | 805.9k | 36.8k | 38.1k | 5.2k | Go | 67.2 |
| [m3db](https://github.com/m3db)  | 736.7k | 10.5k | 5.2k | 5.1k | Go | 61.4 |
| [cilium](https://github.com/cilium)  | 733.4k | 61.9k | 46.4k | 43.6k | Go | 61.1 |
| [falcosecurity](https://github.com/falcosecurity)  | 705.2k | 33.3k | 14.9k | 11.0k | C Header | 58.8 |
| [kubevirt](https://github.com/kubevirt)  | 683.0k | 122.3k | 38.7k | 8.8k | Go | 56.9 |
| [vitessio](https://github.com/vitessio)  | 678.5k | 42.1k | 16.7k | 20.6k | Go | 56.5 |
| [moby](https://github.com/moby)  | 660.9k | 88.6k | 37.5k | 92.8k | Go | 55.1 |
| [fluent](https://github.com/fluent)  | 660.2k | 47.7k | 17.8k | 29.7k | C | 55.0 |
| [cubefs](https://github.com/cubefs)  | 657.8k | 7.0k | 3.2k | 6.0k | Go | 54.8 |
| [istio](https://github.com/istio)  | 583.2k | 61.8k | 84.2k | 44.7k | Go | 48.6 |
| [openebs](https://github.com/openebs)  | 559.4k | 21.7k | 12.9k | 12.8k | Go | 46.6 |
| [argoproj](https://github.com/argoproj)  | 508.3k | 29.6k | 31.1k | 51.3k | Go | 42.4 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 503.3k | 18.2k | 15.2k | 22.1k | Go | 41.9 |
| [opengemini](https://github.com/opengemini)  | 500.0k | 1.2k | 1.3k | 1.3k | Go | 41.7 |
| [denoland](https://github.com/denoland)  | 497.6k | 14.4k | 16.2k | 104.6k | JavaScript | 41.5 |
| [cncf](https://github.com/cncf)  | 496.7k | 78.5k | 23.6k | 36.4k | JavaScript | 41.4 |
| [metrico](https://github.com/metrico)  | 495.5k | 23.6k | 1.6k | 2.6k | Rust | 41.3 |
| [googlecontainertools](https://github.com/googlecontainertools)  | 493.2k | 21.6k | 34.6k | 66.7k | Go | 41.1 |
| [pravega](https://github.com/pravega)  | 471.3k | 6.7k | 6.0k | 2.8k | Java | 39.3 |
| [backstage](https://github.com/backstage)  | 470.3k | 81.1k | 30.2k | 33.2k | TypeScript | 39.2 |
| [kubevela](https://github.com/kubevela)  | 448.4k | 9.2k | 8.1k | 8.5k | Go | 37.4 |
| [athenz](https://github.com/athenz)  | 433.0k | 5.0k | 3.6k | 1.0k | Java | 36.1 |
| [knative](https://github.com/knative)  | 432.7k | 40.5k | 40.8k | 15.1k | Go | 36.1 |
| [go-faster](https://github.com/go-faster)  | 430.0k | 37.4k | 3.3k | 525 | Go | 35.8 |
| [awslabs](https://github.com/awslabs)  | 427.5k | 11.1k | 4.5k | 45.6k | Java | 35.6 |
| [opencurve](https://github.com/opencurve)  | 418.8k | 6.4k | 2.7k | 2.6k | C++ | 34.9 |
| [werf](https://github.com/werf)  | 410.4k | 20.1k | 8.1k | 6.8k | JavaScript | 34.2 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762 | Java | 33.8 |
| [koordinator-sh](https://github.com/koordinator-sh)  | 399.1k | 2.0k | 2.4k | 1.6k | Go | 33.2 |
| [thanos-io](https://github.com/thanos-io)  | 383.9k | 21.5k | 6.3k | 14.7k | Go | 32.0 |
| [prometheus](https://github.com/prometheus)  | 381.1k | 44.4k | 30.1k | 120.9k | Go | 31.8 |
| [kubeovn](https://github.com/kubeovn)  | 377.0k | 24.1k | 4.8k | 2.2k | Go | 31.4 |
| [pixie-io](https://github.com/pixie-io)  | 376.8k | 13.7k | 2.5k | 6.5k | C++ | 31.4 |
| [vectordotdev](https://github.com/vectordotdev)  | 374.2k | 28.7k | 16.0k | 22.2k | Rust | 31.2 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 371.7k | 40.4k | 8.6k | 20.9k | Go | 31.0 |
| [dapr](https://github.com/dapr)  | 360.5k | 28.0k | 17.7k | 32.5k | Go | 30.0 |
| [facebook](https://github.com/facebook)  | 357.5k | 21.0k | 18.3k | 239.5k | JavaScript | 29.8 |
| [containerd](https://github.com/containerd)  | 355.2k | 35.7k | 19.7k | 37.7k | Go | 29.6 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 347.2k | 19.2k | 15.2k | 20.0k | Go | 28.9 |
| [longhorn](https://github.com/longhorn)  | 345.1k | 19.0k | 17.2k | 8.2k | C | 28.8 |
| [chaosblade-io](https://github.com/chaosblade-io)  | 340.5k | 1.7k | 1.4k | 7.7k | JavaScript | 28.4 |
| [django](https://github.com/django)  | 339.9k | 33.9k | 19.9k | 85.3k | Python | 28.3 |
| [operator-framework](https://github.com/operator-framework)  | 338.1k | 26.6k | 21.7k | 16.3k | Go | 28.2 |
| [radius-project](https://github.com/radius-project)  | 315.1k | 10.1k | 6.8k | 1.8k | C# | 26.3 |
| [LINBIT](https://github.com/LINBIT)  | 310.8k | 5.3k | 18 | 1.1k | Java | 25.9 |
| [kyverno](https://github.com/kyverno)  | 305.0k | 18.6k | 18.1k | 8.5k | Go | 25.4 |
| [open-cluster-management-io](https://github.com/open-cluster-management-io)  | 303.3k | 10.2k | 6.8k | 2.6k | Go | 25.3 |
| [parallaxsecond](https://github.com/parallaxsecond)  | 292.2k | 4.7k | 2.0k | 813 | Rust | 24.4 |
| [confidential-containers](https://github.com/confidential-containers)  | 283.0k | 29.3k | 5.4k | 1.2k | Go | 23.6 |
| [gravity-ui](https://github.com/gravity-ui)  | 265.9k | 11.3k | 9.4k | 2.0k | TypeScript | 22.2 |
| [inspektor-gadget](https://github.com/inspektor-gadget)  | 263.3k | 16.2k | 4.4k | 2.6k | C Header | 21.9 |
| [antrea-io](https://github.com/antrea-io)  | 262.3k | 6.7k | 8.4k | 1.8k | Go | 21.9 |
| [keptn](https://github.com/keptn)  | 256.1k | 14.0k | 11.0k | 2.3k | Go | 21.3 |
| [siderolabs](https://github.com/siderolabs)  | 253.4k | 12.0k | 15.2k | 11.6k | Go | 21.1 |
| [strimzi](https://github.com/strimzi)  | 253.2k | 9.6k | 8.9k | 6.4k | Java | 21.1 |
| [uber](https://github.com/uber)  | 251.7k | 1.6k | 342 | 7.2k | Go | 21.0 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.0k | 12.6k | 13.0k | 29.4k | Go | 20.8 |
| [kusionstack](https://github.com/kusionstack)  | 247.3k | 3.4k | 3.1k | 3.0k | Go | 20.6 |
| [goharbor](https://github.com/goharbor)  | 245.9k | 19.2k | 13.6k | 29.0k | Go | 20.5 |
| [kumahq](https://github.com/kumahq)  | 241.2k | 16.2k | 18.6k | 4.2k | Go | 20.1 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 18.7k | 20.5k | 20.2k | Rust | 19.8 |
| [kata-containers](https://github.com/kata-containers)  | 225.2k | 41.7k | 13.7k | 10.7k | Go | 18.8 |
| [buildpacks](https://github.com/buildpacks)  | 225.1k | 33.7k | 5.4k | 3.9k | Go | 18.8 |
| [spiffe](https://github.com/spiffe)  | 220.8k | 21.8k | 9.0k | 4.6k | Go | 18.4 |
| [litmuschaos](https://github.com/litmuschaos)  | 211.6k | 8.8k | 9.2k | 5.6k | TypeScript | 17.6 |
| [Netflix](https://github.com/Netflix)  | 208.1k | 4.2k | 2.3k | 549 | Java | 17.3 |
| [matplotlib](https://github.com/matplotlib)  | 206.9k | 53.3k | 19.6k | 21.8k | Python | 17.2 |
| [alibaba](https://github.com/alibaba)  | 205.2k | 4.1k | 859 | 27.2k | Java | 17.1 |
| [kcl-lang](https://github.com/kcl-lang)  | 200.3k | 9.0k | 4.3k | 2.6k | Rust | 16.7 |
| [fluxcd](https://github.com/fluxcd)  | 200.1k | 37.0k | 16.1k | 25.4k | Go | 16.7 |
| [megaease](https://github.com/megaease)  | 199.7k | 4.5k | 1.9k | 11.5k | Go | 16.6 |
| [pion](https://github.com/pion)  | 197.0k | 10.4k | 8.3k | 26.0k | Go | 16.4 |
| [wasmcloud](https://github.com/wasmcloud)  | 192.4k | 19.0k | 9.5k | 3.5k | Rust | 16.0 |
| [flatcar](https://github.com/flatcar)  | 192.3k | 54.1k | 6.8k | 1.9k | Go | 16.0 |
| [tremor-rs](https://github.com/tremor-rs)  | 187.6k | 8.1k | 4.3k | 1.0k | Rust | 15.6 |
| [etcd-io](https://github.com/etcd-io)  | 186.5k | 34.2k | 17.3k | 64.2k | Go | 15.6 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 181.9k | 5.5k | 6.0k | 6.1k | Python | 15.2 |
| [carvel-dev](https://github.com/carvel-dev)  | 180.0k | 10.4k | 5.3k | 5.0k | Go | 15.0 |
| [containers](https://github.com/containers)  | 173.5k | 26.0k | 14.9k | 29.1k | Go | 14.5 |
| [cert-manager](https://github.com/cert-manager)  | 165.2k | 27.9k | 12.4k | 14.9k | Go | 13.8 |
| [loxilb-io](https://github.com/loxilb-io)  | 162.6k | 9.2k | 1.2k | 2.0k | C | 13.6 |
| [kubeslice](https://github.com/kubeslice)  | 157.9k | 7.7k | 1.7k | 510 | Go | 13.2 |
| [meshery](https://github.com/meshery)  | 155.9k | 69.4k | 18.3k | 10.6k | JavaScript | 13.0 |
| [kuadrant](https://github.com/kuadrant)  | 154.5k | 10.9k | 4.5k | 493 | Go | 12.9 |
| [crossplane](https://github.com/crossplane)  | 154.3k | 19.7k | 7.5k | 12.7k | Go | 12.8 |
| [networkservicemesh](https://github.com/networkservicemesh)  | 152.1k | 55.2k | 43.5k | 704 | Go | 12.7 |
| [in-toto](https://github.com/in-toto)  | 151.9k | 13.6k | 4.0k | 2.5k | Python | 12.7 |
| [perses](https://github.com/perses)  | 150.4k | 5.3k | 4.3k | 1.8k | TypeScript | 12.5 |
| [connectrpc](https://github.com/connectrpc)  | 149.0k | 5.1k | 8.2k | 7.2k | Go | 12.4 |
| [kubeedge](https://github.com/kubeedge)  | 148.3k | 9.8k | 5.8k | 8.8k | Go | 12.4 |
| [kube-vip](https://github.com/kube-vip)  | 145.0k | 2.2k | 1.0k | 2.8k | C Header | 12.1 |
| [drakkan](https://github.com/drakkan)  | 142.0k | 2.2k | 286 | 11.1k | Go | 11.8 |
| [devfile](https://github.com/devfile)  | 138.2k | 6.7k | 4.3k | 607 | Go | 11.5 |
| [getporter](https://github.com/getporter)  | 136.8k | 8.5k | 3.9k | 1.5k | Go | 11.4 |
| [ovn-org](https://github.com/ovn-org)  | 136.7k | 24.6k | 990 | 890 | C | 11.4 |
| [armadaproject](https://github.com/armadaproject)  | 136.5k | 8.0k | 4.0k | 579 | Go | 11.4 |
| [fluid-cloudnative](https://github.com/fluid-cloudnative)  | 136.3k | 3.7k | 4.0k | 1.8k | Go | 11.4 |
| [kubescape](https://github.com/kubescape)  | 133.7k | 21.9k | 5.6k | 11.5k | Go | 11.1 |
| [wasmedge](https://github.com/wasmedge)  | 133.5k | 6.0k | 3.2k | 10.7k | Rust | 11.1 |
| [rook](https://github.com/rook)  | 131.8k | 28.4k | 10.4k | 13.5k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 7.3k | 5.7k | 4.6k | Go | 10.9 |
| [sermant-io](https://github.com/sermant-io)  | 130.4k | 2.5k | 1.4k | 1.4k | Java | 10.9 |
| [kcp-dev](https://github.com/kcp-dev)  | 128.5k | 8.9k | 3.4k | 2.8k | Go | 10.7 |
| [keylime](https://github.com/keylime)  | 125.0k | 3.5k | 2.5k | 661 | C | 10.4 |
| [nocalhost](https://github.com/nocalhost)  | 122.6k | 6.2k | 2.1k | 2.0k | Go | 10.2 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 6.5k | 5.5k | 5.9k | Go | 10.0 |
| [artifacthub](https://github.com/artifacthub)  | 116.7k | 2.1k | 3.5k | 1.9k | TypeScript | 9.7 |
| [openelb](https://github.com/openelb)  | 116.3k | 695 | 397 | 1.8k | JavaScript | 9.7 |
| [pipe-cd](https://github.com/pipe-cd)  | 115.1k | 7.3k | 5.3k | 1.4k | Go | 9.6 |
| [serverless-devs](https://github.com/serverless-devs)  | 114.0k | 5.3k | 772 | 1.9k | JavaScript | 9.5 |
| [kptdev](https://github.com/kptdev)  | 113.6k | 4.3k | 4.6k | 2.0k | Go | 9.5 |
| [paralus](https://github.com/paralus)  | 112.6k | 2.1k | 1.3k | 1.2k | Go | 9.4 |
| [notaryproject](https://github.com/notaryproject)  | 107.7k | 6.8k | 5.6k | 4.4k | Go | 9.0 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 107.1k | 3.7k | 3.7k | 7.8k | Go | 8.9 |
| [kedacore](https://github.com/kedacore)  | 105.3k | 5.9k | 7.2k | 10.8k | Go | 8.8 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 103.5k | 19.6k | 4.1k | 4.5k | Python | 8.6 |
| [VKCOM](https://github.com/VKCOM)  | 101.5k | 15.1k | 8.5k | 1.4k | TypeScript | 8.4 |
| [karmada-io](https://github.com/karmada-io)  | 100.5k | 8.9k | 5.9k | 5.3k | Go | 8.4 |
| [kubeclipper](https://github.com/kubeclipper)  | 99.5k | 686 | 1.0k | 510 | Go | 8.3 |
| [openyurtio](https://github.com/openyurtio)  | 98.6k | 2.9k | 2.6k | 2.1k | Go | 8.2 |
| [bfenetworks](https://github.com/bfenetworks)  | 98.2k | 1.5k | 875 | 6.4k | Go | 8.2 |
| [v6d-io](https://github.com/v6d-io)  | 92.2k | 1.5k | 1.4k | 928 | C++ | 7.7 |
| [spidernet-io](https://github.com/spidernet-io)  | 92.1k | 7.3k | 6.7k | 912 | Go | 7.7 |
| [serverlessworkflow](https://github.com/serverlessworkflow)  | 90.4k | 4.5k | 2.1k | 1.5k | C# | 7.5 |
| [kairos-io](https://github.com/kairos-io)  | 88.2k | 14.2k | 9.7k | 1.7k | Go | 7.3 |
| [fnproject](https://github.com/fnproject)  | 87.4k | 8.5k | 2.5k | 6.4k | Go | 7.3 |
| [opencost](https://github.com/opencost)  | 82.3k | 6.7k | 3.5k | 6.3k | Go | 6.9 |
| [openservicemesh](https://github.com/openservicemesh)  | 81.9k | 5.2k | 4.1k | 2.6k | Go | 6.8 |
| [hwameistor](https://github.com/hwameistor)  | 80.4k | 3.4k | 1.5k | 627 | Go | 6.7 |
| [easegress-io](https://github.com/easegress-io)  | 79.9k | 2.2k | 1.3k | 5.9k | Go | 6.7 |
| [firecracker-microvm](https://github.com/firecracker-microvm)  | 79.7k | 7.4k | 3.9k | 30.7k | Rust | 6.6 |
| [cloudevents](https://github.com/cloudevents)  | 79.0k | 4.8k | 3.3k | 8.3k | Go | 6.6 |
| [opencontainers](https://github.com/opencontainers)  | 78.9k | 14.0k | 6.4k | 18.2k | Go | 6.6 |
| [oras-project](https://github.com/oras-project)  | 76.9k | 5.7k | 3.3k | 2.5k | Go | 6.4 |
| [coredns](https://github.com/coredns)  | 76.2k | 6.6k | 6.1k | 14.9k | Go | 6.3 |
| [project-zot](https://github.com/project-zot)  | 75.8k | 1.9k | 3.1k | 1.5k | Go | 6.3 |
| [devspace-sh](https://github.com/devspace-sh)  | 74.9k | 6.1k | 1.9k | 4.8k | Go | 6.2 |
| [volcano-sh](https://github.com/volcano-sh)  | 74.6k | 7.3k | 3.5k | 5.5k | Go | 6.2 |
| [gotd](https://github.com/gotd)  | 74.5k | 9.5k | 4.3k | 2.1k | Go | 6.2 |
| [spectralops](https://github.com/spectralops)  | 73.0k | 359 | 63 | 1.2k | JavaScript | 6.1 |
| [tinkerbell](https://github.com/tinkerbell)  | 71.4k | 18.0k | 5.1k | 2.7k | Go | 6.0 |
| [helm](https://github.com/helm)  | 71.3k | 32.2k | 29.5k | 56.4k | Go | 5.9 |
| [cri-o](https://github.com/cri-o)  | 71.1k | 12.8k | 8.5k | 5.7k | Go | 5.9 |
| [ten-nancy](https://github.com/ten-nancy)  | 71.0k | 326 | 7 | 44 | C++ | 5.9 |
| [oscal-compass](https://github.com/oscal-compass)  | 71.0k | 2.4k | 1.6k | 366 | Python | 5.9 |
| [metal3-io](https://github.com/metal3-io)  | 70.9k | 17.4k | 10.5k | 1.6k | Go | 5.9 |
| [meilisearch](https://github.com/meilisearch)  | 70.7k | 12.6k | 2.6k | 53.5k | Rust | 5.9 |
| [submariner-io](https://github.com/submariner-io)  | 69.8k | 13.9k | 16.6k | 3.0k | Go | 5.8 |
| [runatlantis](https://github.com/runatlantis)  | 68.5k | 4.9k | 20.6k | 8.7k | Go | 5.7 |
| [open-feature](https://github.com/open-feature)  | 67.6k | 15.4k | 14.2k | 3.5k | Go | 5.6 |
| [microcks](https://github.com/microcks)  | 66.0k | 6.9k | 2.8k | 2.1k | Java | 5.5 |
| [openfga](https://github.com/openfga)  | 65.8k | 9.3k | 7.2k | 5.0k | Go | 5.5 |
| [kanisterio](https://github.com/kanisterio)  | 63.6k | 3.3k | 3.3k | 847 | Go | 5.3 |
| [ent](https://github.com/ent)  | 63.2k | 2.3k | 2.5k | 16.6k | Go | 5.3 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 2.0k | 2.6k | 3.0k | Go | 5.2 |
| [telepresenceio](https://github.com/telepresenceio)  | 61.1k | 10.8k | 2.4k | 7.0k | Go | 5.1 |
| [vuejs](https://github.com/vuejs)  | 60.3k | 3.6k | 2.6k | 209.5k | TypeScript | 5.0 |
| [openfaas](https://github.com/openfaas)  | 59.0k | 9.0k | 4.1k | 31.9k | Go | 4.9 |
| [projectriff](https://github.com/projectriff)  | 57.7k | 6.6k | 4.2k | 887 | Go | 4.8 |
| [external-secrets](https://github.com/external-secrets)  | 57.3k | 4.5k | 3.9k | 8.6k | Go | 4.8 |
| [cartography-cncf](https://github.com/cartography-cncf)  | 57.2k | 1.2k | 1.4k | 3.6k | Python | 4.8 |
| [ariga](https://github.com/ariga)  | 55.7k | 2.7k | 2.8k | 7.4k | Go | 4.6 |
| [project-hami](https://github.com/project-hami)  | 53.2k | 1.9k | 1.1k | 2.6k | Go | 4.4 |
| [rkt](https://github.com/rkt)  | 52.4k | 5.6k | 2.5k | 8.8k | Go | 4.4 |
| [trickstercache](https://github.com/trickstercache)  | 52.1k | 1.3k | 640 | 2.1k | Go | 4.3 |
| [cloud-bulldozer](https://github.com/cloud-bulldozer)  | 51.9k | 7.9k | 3.6k | 914 | Python | 4.3 |
| [headlamp-k8s](https://github.com/headlamp-k8s)  | 51.4k | 6.0k | 2.2k | 2.7k | TypeScript | 4.3 |
| [kaito-project](https://github.com/kaito-project)  | 51.0k | 6.1k | 2.1k | 1.3k | Go | 4.2 |
| [kubestellar](https://github.com/kubestellar)  | 50.6k | 9.6k | 4.1k | 823 | Go | 4.2 |
| [slimtoolkit](https://github.com/slimtoolkit)  | 49.7k | 1.3k | 451 | 22.5k | Go | 4.1 |
| [shipwright-io](https://github.com/shipwright-io)  | 48.8k | 4.3k | 2.4k | 835 | Go | 4.1 |
| [theupdateframework](https://github.com/theupdateframework)  | 48.6k | 11.8k | 4.8k | 3.1k | Python | 4.0 |
| [superedge](https://github.com/superedge)  | 48.2k | 524 | 461 | 1.1k | Go | 4.0 |
| [virtual-kubelet](https://github.com/virtual-kubelet)  | 48.1k | 3.0k | 1.7k | 5.2k | Go | 4.0 |
| [kudobuilder](https://github.com/kudobuilder)  | 47.9k | 2.3k | 2.2k | 2.2k | Go | 4.0 |
| [xline-kv](https://github.com/xline-kv)  | 47.6k | 1.6k | 1.0k | 713 | Rust | 4.0 |
| [distribution](https://github.com/distribution)  | 45.7k | 6.1k | 2.4k | 10.3k | Go | 3.8 |
| [k3s-io](https://github.com/k3s-io)  | 44.5k | 5.9k | 6.5k | 37.0k | Go | 3.7 |
| [openfunction](https://github.com/openfunction)  | 44.4k | 1.7k | 1.2k | 1.7k | Go | 3.7 |
| [skooner-k8s](https://github.com/skooner-k8s)  | 44.0k | 543 | 259 | 1.4k | JavaScript | 3.7 |
| [kubedl-io](https://github.com/kubedl-io)  | 44.0k | 1.4k | 409 | 618 | Go | 3.7 |
| [tauri-apps](https://github.com/tauri-apps)  | 43.4k | 5.7k | 6.0k | 97.1k | Rust | 3.6 |
| [krkn-chaos](https://github.com/krkn-chaos)  | 43.2k | 2.0k | 1.6k | 507 | Python | 3.6 |
| [kubearmor](https://github.com/kubearmor)  | 41.6k | 7.4k | 3.3k | 2.2k | Go | 3.5 |
| [metallb](https://github.com/metallb)  | 40.9k | 3.5k | 2.4k | 8.0k | Go | 3.4 |
| [vitalif](https://github.com/vitalif)  | 40.8k | 2.3k | 22 | 193 | C++ | 3.4 |
| [capsule-rs](https://github.com/capsule-rs)  | 40.4k | 252 | 140 | 356 | Rust | 3.4 |
| [vscode-kubernetes-tools](https://github.com/vscode-kubernetes-tools)  | 39.4k | 956 | 983 | 752 | TypeScript | 3.3 |
| [devstream-io](https://github.com/devstream-io)  | 39.3k | 663 | 1.1k | 887 | Go | 3.3 |
| [ogen-go](https://github.com/ogen-go)  | 39.2k | 6.0k | 2.1k | 1.9k | Go | 3.3 |
| [aeraki-mesh](https://github.com/aeraki-mesh)  | 39.1k | 1.1k | 503 | 1.0k | Go | 3.2 |
| [youki-dev](https://github.com/youki-dev)  | 37.9k | 6.1k | 3.1k | 7.2k | Rust | 3.2 |
| [jquery](https://github.com/jquery)  | 37.2k | 6.8k | 3.1k | 59.6k | JavaScript | 3.1 |
| [kube-logging](https://github.com/kube-logging)  | 35.9k | 4.7k | 2.3k | 1.8k | Go | 3.0 |
| [kuasar-io](https://github.com/kuasar-io)  | 35.7k | 379 | 186 | 1.3k | Rust | 3.0 |
| [bpfman](https://github.com/bpfman)  | 35.5k | 3.4k | 1.6k | 711 | Rust | 3.0 |
| [score-spec](https://github.com/score-spec)  | 35.5k | 2.6k | 1.3k | 8.9k | Go | 3.0 |
| [sustainable-computing-io](https://github.com/sustainable-computing-io)  | 35.1k | 5.6k | 3.5k | 1.6k | Go | 2.9 |
| [kmesh-net](https://github.com/kmesh-net)  | 35.1k | 3.7k | 1.3k | 708 | Go | 2.9 |
| [ratify-project](https://github.com/ratify-project)  | 35.0k | 1.5k | 1.7k | 255 | Go | 2.9 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 3.4k | 1.6k | 8.4k | Go | 2.9 |
| [clusternet](https://github.com/clusternet)  | 33.4k | 1.2k | 801 | 1.4k | Go | 2.8 |
| [sealerio](https://github.com/sealerio)  | 32.6k | 1.2k | 1.4k | 2.1k | Go | 2.7 |
| [kube-rs](https://github.com/kube-rs)  | 31.9k | 4.8k | 2.0k | 4.0k | Rust | 2.7 |
| [bank-vaults](https://github.com/bank-vaults)  | 31.7k | 13.6k | 7.9k | 2.5k | Go | 2.6 |
| [clusterpedia-io](https://github.com/clusterpedia-io)  | 31.6k | 1.6k | 871 | 886 | Go | 2.6 |
| [dexidp](https://github.com/dexidp)  | 29.3k | 3.9k | 3.1k | 10.3k | Go | 2.4 |
| [piraeusdatastore](https://github.com/piraeusdatastore)  | 28.9k | 2.4k | 1.2k | 1.2k | Go | 2.4 |
| [opentracing](https://github.com/opentracing)  | 27.2k | 2.7k | 1.5k | 10.2k | Java | 2.3 |
| [krustlet](https://github.com/krustlet)  | 23.6k | 2.4k | 535 | 3.6k | Rust | 2.0 |
| [alacritty](https://github.com/alacritty)  | 23.1k | 2.5k | 2.5k | 60.5k | Rust | 1.9 |
| [project-stacker](https://github.com/project-stacker)  | 22.4k | 1.5k | 721 | 327 | Go | 1.9 |
| [curiefense](https://github.com/curiefense)  | 22.3k | 2.9k | 832 | 617 | Rust | 1.9 |
| [redhat-chaos](https://github.com/redhat-chaos)  | 22.3k | 911 | 784 | 317 | Python | 1.9 |
| [projectcapsule](https://github.com/projectcapsule)  | 21.9k | 2.3k | 2.0k | 2.0k | Go | 1.8 |
| [fabedge](https://github.com/fabedge)  | 21.6k | 1.1k | 529 | 535 | Go | 1.8 |
| [project-akri](https://github.com/project-akri)  | 20.5k | 1.0k | 604 | 1.2k | Rust | 1.7 |
| [k8gb-io](https://github.com/k8gb-io)  | 19.4k | 2.0k | 1.8k | 1.1k | Go | 1.6 |
| [foniod](https://github.com/foniod)  | 18.1k | 1.7k | 495 | 2.1k | Rust | 1.5 |
| [lima-vm](https://github.com/lima-vm)  | 17.4k | 5.9k | 2.9k | 19.1k | Go | 1.4 |
| [schemahero](https://github.com/schemahero)  | 17.0k | 1.7k | 1.1k | 1.1k | Go | 1.4 |
| [hexa-org](https://github.com/hexa-org)  | 16.5k | 956 | 181 | 125 | Go | 1.4 |
| [getsops](https://github.com/getsops)  | 15.3k | 2.4k | 1.1k | 19.7k | Go | 1.3 |
| [k8up-io](https://github.com/k8up-io)  | 14.1k | 2.4k | 864 | 874 | Go | 1.2 |
| [kubean-io](https://github.com/kubean-io)  | 13.2k | 2.8k | 1.3k | 552 | Go | 1.1 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 824 | 150 | 712 | Go | 1.1 |
| [pallets](https://github.com/pallets)  | 12.8k | 5.5k | 2.7k | 70.5k | Python | 1.1 |
| [kuberhealthy](https://github.com/kuberhealthy)  | 11.8k | 3.3k | 1.0k | 2.2k | Go | 990.0m |
| [servicemeshinterface](https://github.com/servicemeshinterface)  | 11.6k | 974 | 375 | 1.2k | Go | 970.0m |
| [k8sgpt-ai](https://github.com/k8sgpt-ai)  | 11.0k | 2.6k | 2.0k | 7.5k | Go | 920.0m |
| [ko-build](https://github.com/ko-build)  | 9.5k | 1.8k | 1.4k | 8.2k | Go | 790.0m |
| [eraser-dev](https://github.com/eraser-dev)  | 9.4k | 724 | 935 | 568 | Go | 780.0m |
| [devspace-cloud](https://github.com/devspace-cloud)  | 8.7k | 348 | 83 | 113 | Go | 720.0m |
| [kube-burner](https://github.com/kube-burner)  | 8.0k | 1.7k | 1.0k | 675 | Go | 660.0m |
| [tellerops](https://github.com/tellerops)  | 7.7k | 246 | 167 | 3.2k | Go | 640.0m |
| [cni-genie](https://github.com/cni-genie)  | 5.4k | 490 | 136 | 538 | Go | 450.0m |
| [project-copacetic](https://github.com/project-copacetic)  | 5.4k | 1.0k | 1.2k | 1.5k | Go | 450.0m |
| [merbridge](https://github.com/merbridge)  | 4.8k | 296 | 386 | 806 | Go | 400.0m |
| [opcr-io](https://github.com/opcr-io)  | 3.6k | 734 | 257 | 279 | Go | 300.0m |
| [krator-rs](https://github.com/krator-rs)  | 2.9k | 498 | 34 | 167 | Rust | 250.0m |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251 | 117 | 2.0k | Go | 190.0m |
| [openobservability](https://github.com/openobservability)  | 2.3k | 251 | 117 | 2.0k | Go | 190.0m |
| [service-mesh-performance](https://github.com/service-mesh-performance)  | 2.2k | 906 | 261 | 305 | JavaScript | 180.0m |
| [kubereboot](https://github.com/kubereboot)  | 2.0k | 1.9k | 1.1k | 2.4k | Go | 170.0m |
| [g-research](https://github.com/g-research)  | 870 | 9 | 0 | 4 | Go | 70.0m |
| [gitops-working-group](https://github.com/gitops-working-group)  | 0 | 43 | 57 | 552 | N/A | 0.0 |


### Repositories
| Repository | SLOC | Commits | PRs | Stars | Language | Effort |
|------------|------|---------|-----|-------|----------|--------|
| [torvalds/linux](https://github.com/torvalds/linux/)  | 19.1M | 1.4M | 1.3k | 204.2k | C | 1.6k |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3.3M | 185.4k | 58.8k | 191.9k | C++ | 271.2 |
| [ydb-platform/nbs](https://github.com/ydb-platform/nbs/)  | 3.2M | 8.0k | 3.8k | 80 | C++ | 266.6 |
| [ydb-platform/ydb](https://github.com/ydb-platform/ydb/)  | 2.8M | 38.3k | 20.4k | 4.6k | C++ | 237.0 |
| [ytsaurus/ytsaurus](https://github.com/ytsaurus/ytsaurus/)  | 2.4M | 85.6k | 911 | 2.1k | C++ | 202.7 |
| [elastic/elasticsearch](https://github.com/elastic/elasticsearch/)  | 2.4M | 90.6k | 94.4k | 74.9k | Java | 201.8 |
| [apache/hadoop](https://github.com/apache/hadoop/)  | 2.0M | 28.0k | 8.0k | 15.3k | Java | 165.5 |
| [godotengine/godot](https://github.com/godotengine/godot/)  | 2.0M | 77.5k | 51.4k | 101.8k | C++ | 164.8 |
| [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/)  | 1.7M | 110.0k | 77.6k | 31.3k | Go | 139.7 |
| [rust-lang/rust](https://github.com/rust-lang/rust/)  | 1.5M | 307.2k | 87.1k | 106.9k | Rust | 126.9 |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1.5M | 132.9k | 86.0k | 117.8k | Go | 124.0 |
| [apache/camel](https://github.com/apache/camel/)  | 1.4M | 76.9k | 19.4k | 6.0k | Java | 117.8 |
| [python/cpython](https://github.com/python/cpython/)  | 1.3M | 128.6k | 64.9k | 69.2k | Python | 107.1 |
| [openkruise/game.openkruise.io](https://github.com/openkruise/game.openkruise.io/)  | 1.2M | 4 | 0 | 0 | JavaScript | 98.3 |
| [apache/hive](https://github.com/apache/hive/)  | 1.1M | 17.9k | 6.1k | 5.8k | Java | 93.2 |
| [apache/harmony-classlib](https://github.com/apache/harmony-classlib/)  | 1.1M | 7.5k | 3 | 6 | Java | 92.2 |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1.1M | 192.4k | 61.7k | 43.2k | C++ | 88.9 |
| [golang/go](https://github.com/golang/go/)  | 1.1M | 64.0k | 4.7k | 130.1k | Go | 88.3 |
| [kubewarden/kubecon-24-eu-kubewarden](https://github.com/kubewarden/kubecon-24-eu-kubewarden/)  | 1.0M | 50 | 45 | 2 | JavaScript | 86.7 |
| [prestodb/presto](https://github.com/prestodb/presto/)  | 1.0M | 20.8k | 13.9k | 14.6k | Java | 86.0 |
| [nodejs/node](https://github.com/nodejs/node/)  | 955.2k | 45.2k | 38.1k | 113.6k | JavaScript | 79.6 |
| [envoyproxy/mobile-website](https://github.com/envoyproxy/mobile-website/)  | 916.7k | 5.2k | 19 | 1 | JavaScript | 76.4 |
| [ziglang/zig](https://github.com/ziglang/zig/)  | 900.7k | 35.2k | 12.0k | 41.5k | Zig | 75.1 |
| [konveyor/github-actions](https://github.com/konveyor/github-actions/)  | 894.4k | 17 | 14 | 0 | JavaScript | 74.5 |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889.7k | 62.7k | 67.6k | 70.2k | TypeScript | 74.1 |
| [inclavare-containers/tng-envoy](https://github.com/inclavare-containers/tng-envoy/)  | 866.5k | 21.5k | 0 | 1 | C++ | 72.2 |
| [brigadecore/go-libgit2](https://github.com/brigadecore/go-libgit2/)  | 778.0k | 3 | 0 | 0 | JavaScript | 64.8 |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 752.9k | 25.3k | 28.2k | 26.6k | C++ | 62.7 |
| [keycloak/keycloak](https://github.com/keycloak/keycloak/)  | 675.1k | 29.3k | 21.4k | 30.0k | Java | 56.3 |
| [tarantool/tarantool](https://github.com/tarantool/tarantool/)  | 648.6k | 19.4k | 3.9k | 3.6k | Lua | 54.0 |
| [apache/cassandra](https://github.com/apache/cassandra/)  | 631.4k | 31.4k | 4.4k | 9.4k | Java | 52.6 |
| [elastic/beats](https://github.com/elastic/beats/)  | 616.0k | 19.3k | 35.5k | 12.5k | Go | 51.3 |
| [ytsaurus/ytsaurus-cpp-sdk](https://github.com/ytsaurus/ytsaurus-cpp-sdk/)  | 615.1k | 5.1k | 3 | 3 | C++ | 51.3 |
| [vitessio/vitess](https://github.com/vitessio/vitess/)  | 610.7k | 34.8k | 14.3k | 20.3k | Go | 50.9 |
| [apache/lucenenet](https://github.com/apache/lucenenet/)  | 606.9k | 7.0k | 810 | 2.3k | C# | 50.6 |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 509.1k | 57.5k | 28.7k | 43.8k | C++ | 42.4 |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 506.7k | 159 | 262 | 50 | C++ | 42.2 |
| [falcosecurity/libs](https://github.com/falcosecurity/libs/)  | 504.4k | 10.5k | 2.4k | 286 | C Header | 42.0 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 502.6k | 8.5k | 454 | 205 | C++ | 41.9 |
| [denoland/deno](https://github.com/denoland/deno/)  | 497.6k | 14.4k | 16.2k | 104.6k | JavaScript | 41.5 |
| [cubefs/cubefs](https://github.com/cubefs/cubefs/)  | 480.1k | 5.8k | 2.8k | 5.3k | Go | 40.0 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477.3k | 4.3k | 3.7k | 4.9k | Go | 39.8 |
| [apache/openjpa](https://github.com/apache/openjpa/)  | 461.7k | 5.4k | 126 | 144 | Java | 38.5 |
| [apache/activemq](https://github.com/apache/activemq/)  | 447.4k | 12.1k | 1.5k | 2.4k | Java | 37.3 |
| [apache/poi](https://github.com/apache/poi/)  | 424.0k | 13.5k | 895 | 2.1k | Java | 35.3 |
| [opengemini/openGemini](https://github.com/opengemini/openGemini/)  | 410.5k | 450 | 569 | 1.1k | Go | 34.2 |
| [konveyor/tackle-diva](https://github.com/konveyor/tackle-diva/)  | 400.2k | 113 | 94 | 39 | Java | 33.4 |
| [tikv/tikv](https://github.com/tikv/tikv/)  | 395.8k | 8.3k | 13.4k | 16.2k | Rust | 33.0 |
| [netdata/charts](https://github.com/netdata/charts/)  | 395.7k | 836 | 195 | 17 | JavaScript | 33.0 |
| [backstage/backstage](https://github.com/backstage/backstage/)  | 393.6k | 67.7k | 22.9k | 31.5k | TypeScript | 32.8 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389.2k | 19.2k | 34.0k | 4.0k | Go | 32.4 |
| [netdata/netdata](https://github.com/netdata/netdata/)  | 378.8k | 21.0k | 12.4k | 76.3k | C | 31.6 |
| [envoyproxy/archive](https://github.com/envoyproxy/archive/)  | 377.7k | 635 | 37 | 2 | JavaScript | 31.5 |
| [open-telemetry/opentelemetry-ebpf-~](https://github.com/open-telemetry/opentelemetry-ebpf-instrumentation/)  | 367.3k | 592 | 654 | 174 | C Header | 30.6 |
| [cdk8s-team/cdk8s-examples](https://github.com/cdk8s-team/cdk8s-examples/)  | 360.6k | 592 | 605 | 52 | Java | 30.1 |
| [facebook/react](https://github.com/facebook/react/)  | 357.5k | 21.0k | 18.3k | 239.5k | JavaScript | 29.8 |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 351.2k | 12.6k | 1.6k | 6.2k | C++ | 29.3 |
| [athenz/athenz](https://github.com/athenz/athenz/)  | 348.7k | 3.5k | 2.8k | 964 | Java | 29.1 |
| [apache/ofbiz](https://github.com/apache/ofbiz/)  | 347.2k | 24.4k | 27 | 800 | Java | 28.9 |
| [AthenZ/athenz](https://github.com/AthenZ/athenz/)  | 346.1k | 2.6k | 1.9k | 751 | Java | 28.8 |
| [django/django](https://github.com/django/django/)  | 339.9k | 33.9k | 19.9k | 85.3k | Python | 28.3 |
| [istio/istio](https://github.com/istio/istio/)  | 338.8k | 25.5k | 35.5k | 37.5k | Go | 28.2 |
| [opencurve/curve](https://github.com/opencurve/curve/)  | 338.4k | 3.1k | 2.1k | 2.4k | C++ | 28.2 |
| [kubernetes/autoscaler](https://github.com/kubernetes/autoscaler/)  | 332.8k | 9.6k | 6.0k | 8.6k | Go | 27.7 |
| [docker/docker-ce](https://github.com/docker/docker-ce/)  | 324.4k | 54.3k | 662 | 5.6k | Go | 27.0 |
| [cilium/cilium](https://github.com/cilium/cilium/)  | 323.8k | 39.4k | 30.5k | 22.6k | Go | 27.0 |
| [apache/harmony-drlvm](https://github.com/apache/harmony-drlvm/)  | 316.2k | 2.2k | 0 | 14 | C++ | 26.4 |
| [apache/jackrabbit](https://github.com/apache/jackrabbit/)  | 315.5k | 9.4k | 279 | 354 | Java | 26.3 |
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310.8k | 5.3k | 18 | 1.1k | Java | 25.9 |
| [docker/labs](https://github.com/docker/labs/)  | 304.4k | 718 | 398 | 11.1k | PHP | 25.4 |
| [apache/mesos](https://github.com/apache/mesos/)  | 302.9k | 18.4k | 582 | 5.3k | C++ | 25.2 |
| [kubeovn/kubevirt-dpdk](https://github.com/kubeovn/kubevirt-dpdk/)  | 299.8k | 17.1k | 0 | 8 | Go | 25.0 |
| [kubeflow/pipelines](https://github.com/kubeflow/pipelines/)  | 287.5k | 6.5k | 8.1k | 4.0k | Python | 24.0 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284.4k | 3.2k | 2.7k | 786 | C Header | 23.7 |
| [pravega/pravega](https://github.com/pravega/pravega/)  | 283.8k | 3.3k | 3.4k | 2.0k | Java | 23.6 |
| [kubevirt/kubevirt](https://github.com/kubevirt/kubevirt/)  | 281.1k | 25.7k | 12.1k | 6.4k | Go | 23.4 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274.4k | 13.0k | 13.9k | 20.4k | Rust | 22.9 |
| [ydb-platform/ydb-slo-action](https://github.com/ydb-platform/ydb-slo-action/)  | 271.9k | 42 | 22 | 2 | JavaScript | 22.7 |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 269.0k | 64.6k | 31.0k | 4.0k | Go | 22.4 |
| [cdk8s-team/cdk8s](https://github.com/cdk8s-team/cdk8s/)  | 267.9k | 2.3k | 2.2k | 4.7k | TypeScript | 22.3 |
| [metrico/influxdb_iox](https://github.com/metrico/influxdb_iox/)  | 262.3k | 14.0k | 0 | 13 | Rust | 21.9 |
| [apache/cayenne](https://github.com/apache/cayenne/)  | 261.6k | 7.4k | 626 | 336 | Java | 21.8 |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 261.3k | 1.6k | 36 | 24 | JavaScript | 21.8 |
| [ydb-platform/ydb-cpp-sdk](https://github.com/ydb-platform/ydb-cpp-sdk/)  | 259.7k | 3.7k | 239 | 11 | C++ | 21.6 |
| [kyverno/kyverno](https://github.com/kyverno/kyverno/)  | 259.5k | 9.0k | 9.6k | 7.0k | Go | 21.6 |
| [inclavare-containers/deterministic~](https://github.com/inclavare-containers/deterministic-builds/)  | 259.3k | 28 | 4 | 8 | C Header | 21.6 |
| [apache/wicket](https://github.com/apache/wicket/)  | 257.5k | 22.1k | 1.2k | 769 | Java | 21.5 |
| [apache/httpd](https://github.com/apache/httpd/)  | 254.8k | 34.2k | 508 | 3.8k | C | 21.2 |
| [apache/pig](https://github.com/apache/pig/)  | 253.8k | 3.8k | 40 | 686 | Java | 21.1 |
| [inclavare-containers/kong_coco](https://github.com/inclavare-containers/kong_coco/)  | 253.3k | 9 | 14 | 0 | Lua | 21.1 |
| [inspektor-gadget/inspektor-gadget](https://github.com/inspektor-gadget/inspektor-gadget/)  | 251.3k | 8.5k | 3.8k | 2.6k | C Header | 20.9 |
| [koordinator-sh/koordinator](https://github.com/koordinator-sh/koordinator/)  | 250.9k | 1.6k | 1.9k | 1.6k | Go | 20.9 |
| [moby/moby](https://github.com/moby/moby/)  | 244.3k | 54.3k | 27.2k | 70.9k | Go | 20.4 |
| [ydb-platform/yql-parsers](https://github.com/ydb-platform/yql-parsers/)  | 242.2k | 23 | 7 | 0 | TypeScript | 20.2 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241.7k | 10.4k | 13.8k | 26.7k | Go | 20.1 |
| [apache/cocoon](https://github.com/apache/cocoon/)  | 241.6k | 13.2k | 64 | 28 | Java | 20.1 |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235.2k | 6.9k | 8.6k | 11.9k | Java | 19.6 |
| [open-policy-agent/opa](https://github.com/open-policy-agent/opa/)  | 234.5k | 5.9k | 5.1k | 10.7k | Go | 19.5 |
| [netdata/ebpf-co-re](https://github.com/netdata/ebpf-co-re/)  | 231.9k | 87 | 93 | 14 | C Header | 19.3 |
| [thanos-io/thanos-prometheus](https://github.com/thanos-io/thanos-prometheus/)  | 229.2k | 15.3k | 178 | 0 | Go | 19.1 |
| [kumahq/kuma](https://github.com/kumahq/kuma/)  | 221.4k | 8.9k | 11.7k | 3.9k | Go | 18.4 |
| [apache/xmlgraphics-fop](https://github.com/apache/xmlgraphics-fop/)  | 219.7k | 8.6k | 61 | 206 | Java | 18.3 |
| [chaosblade-io/chaosblade-box](https://github.com/chaosblade-io/chaosblade-box/)  | 219.6k | 50 | 50 | 226 | JavaScript | 18.3 |
| [apache/directory-studio](https://github.com/apache/directory-studio/)  | 216.4k | 5.8k | 90 | 140 | Java | 18.0 |
| [uber/peloton](https://github.com/uber/peloton/)  | 214.5k | 705 | 10 | 649 | Go | 17.9 |
| [matplotlib/matplotlib](https://github.com/matplotlib/matplotlib/)  | 206.9k | 53.3k | 19.6k | 21.8k | Python | 17.2 |
| [antrea-io/antrea](https://github.com/antrea-io/antrea/)  | 198.7k | 4.2k | 5.7k | 1.7k | Go | 16.6 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194.2k | 614 | 493 | 3.5k | C Header | 16.2 |
| [apache/xmlgraphics-batik](https://github.com/apache/xmlgraphics-batik/)  | 193.4k | 3.6k | 64 | 235 | Java | 16.1 |
| [nats-io/nats-server](https://github.com/nats-io/nats-server/)  | 190.7k | 11.5k | 4.8k | 18.3k | Go | 15.9 |
| [goharbor/harbor](https://github.com/goharbor/harbor/)  | 190.6k | 12.5k | 9.9k | 26.6k | Go | 15.9 |
| [apache/tapestry4](https://github.com/apache/tapestry4/)  | 190.4k | 3.8k | 13 | 7 | Java | 15.9 |
| [alibaba/fastjson](https://github.com/alibaba/fastjson/)  | 188.1k | 4.0k | 848 | 25.2k | Java | 15.7 |
| [apache/avro](https://github.com/apache/avro/)  | 187.8k | 4.8k | 3.5k | 3.2k | Java | 15.7 |
| [googlecontainertools/skaffold](https://github.com/googlecontainertools/skaffold/)  | 186.7k | 8.9k | 5.3k | 14.6k | Go | 15.6 |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186.6k | 1 | 2 | 2 | JavaScript | 15.6 |
| [fluent/fluent-bit](https://github.com/fluent/fluent-bit/)  | 182.6k | 12.3k | 5.3k | 7.3k | C | 15.2 |
| [apache/directory-server](https://github.com/apache/directory-server/)  | 181.4k | 10.3k | 285 | 167 | Java | 15.1 |
| [cncf/cncf.io](https://github.com/cncf/cncf.io/)  | 180.9k | 2.4k | 159 | 96 | PHP | 15.1 |
| [strimzi/strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator/)  | 175.0k | 6.9k | 6.6k | 5.4k | Java | 14.6 |
| [containers/podman](https://github.com/containers/podman/)  | 173.5k | 26.0k | 14.9k | 29.1k | Go | 14.5 |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173.3k | 13.9k | 9.0k | 4.0k | Go | 14.4 |
| [apache/xalan-j](https://github.com/apache/xalan-j/)  | 170.9k | 4.6k | 6 | 23 | Java | 14.2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.8k | 12.6k | 9.7k | 2.3k | Java | 14.0 |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167.5k | 22.6k | 12.2k | 16.4k | Go | 14.0 |
| [apache/tika](https://github.com/apache/tika/)  | 167.0k | 9.7k | 2.4k | 3.3k | Java | 13.9 |
| [cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian/)  | 162.6k | 4.9k | 5.5k | 5.8k | Python | 13.6 |
| [keptn/keptn](https://github.com/keptn/keptn/)  | 162.2k | 8.4k | 6.3k | 1.8k | Go | 13.5 |
| [apache/zookeeper](https://github.com/apache/zookeeper/)  | 162.1k | 2.7k | 2.3k | 12.6k | Java | 13.5 |
| [argoproj/argo-workflows](https://github.com/argoproj/argo-workflows/)  | 160.3k | 5.9k | 7.2k | 16.1k | Go | 13.4 |
| [apache/couchdb](https://github.com/apache/couchdb/)  | 158.8k | 14.1k | 3.6k | 6.7k | Erlang | 13.2 |
| [apache/shindig](https://github.com/apache/shindig/)  | 158.2k | 4.8k | 7 | 66 | Java | 13.2 |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157.3k | 1.7k | 1.3k | 318 | Java | 13.1 |
| [spiffe/spire](https://github.com/spiffe/spire/)  | 156.7k | 7.4k | 4.7k | 2.1k | Go | 13.1 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155.0k | 16.2k | 10.4k | 60.7k | Go | 12.9 |
| [werf/actions](https://github.com/werf/actions/)  | 153.1k | 132 | 73 | 86 | JavaScript | 12.8 |
| [longhorn/libqcow](https://github.com/longhorn/libqcow/)  | 152.6k | 1 | 0 | 0 | C | 12.7 |
| [go-faster/ytsaurus-ui](https://github.com/go-faster/ytsaurus-ui/)  | 151.4k | 636 | 0 | 0 | TypeScript | 12.6 |
| [apache/etch](https://github.com/apache/etch/)  | 150.7k | 549 | 0 | 17 | Java | 12.6 |
| [argoproj/argo-cd](https://github.com/argoproj/argo-cd/)  | 149.4k | 9.4k | 12.9k | 20.8k | Go | 12.4 |
| [apache/stdcxx](https://github.com/apache/stdcxx/)  | 148.0k | 2.2k | 0 | 65 | C++ | 12.3 |
| [containerd/containerd](https://github.com/containerd/containerd/)  | 147.3k | 15.7k | 8.9k | 19.4k | Go | 12.3 |
| [kubevela/samples](https://github.com/kubevela/samples/)  | 146.8k | 228 | 129 | 138 | JavaScript | 12.2 |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144.5k | 11.4k | 4.9k | 15.1k | Go | 12.1 |
| [kubevela/kubevela](https://github.com/kubevela/kubevela/)  | 144.3k | 4.0k | 4.6k | 7.5k | Go | 12.0 |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143.2k | 5.6k | 5.7k | 22.4k | Go | 11.9 |
| [drakkan/sftpgo](https://github.com/drakkan/sftpgo/)  | 142.0k | 2.2k | 286 | 11.1k | Go | 11.8 |
| [apache/maven-plugins](https://github.com/apache/maven-plugins/)  | 141.6k | 15.1k | 146 | 238 | Java | 11.8 |
| [kubeflow/kfp-tekton](https://github.com/kubeflow/kfp-tekton/)  | 141.3k | 815 | 1.1k | 181 | TypeScript | 11.8 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.1k | 2.4k | 3.7k | 436 | C Header | 11.8 |
| [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize/)  | 139.5k | 7.0k | 3.5k | 11.7k | Go | 11.6 |
| [radius-project/radius](https://github.com/radius-project/radius/)  | 136.1k | 3.4k | 3.8k | 1.6k | Go | 11.3 |
| [tikv/pd](https://github.com/tikv/pd/)  | 135.9k | 4.7k | 6.9k | 1.1k | Go | 11.3 |
| [fluid-cloudnative/fluid](https://github.com/fluid-cloudnative/fluid/)  | 134.5k | 3.3k | 3.9k | 1.8k | Go | 11.2 |
| [kube-vip/kube-gateway](https://github.com/kube-vip/kube-gateway/)  | 134.1k | 10 | 4 | 5 | C Header | 11.2 |
| [confidential-containers/kata-conta~](https://github.com/confidential-containers/kata-containers-CCv0/)  | 133.2k | 10.9k | 2 | 5 | Go | 11.1 |
| [radius-project/bicep](https://github.com/radius-project/bicep/)  | 132.9k | 4.3k | 777 | 6 | C# | 11.1 |
| [apache/apr-iconv](https://github.com/apache/apr-iconv/)  | 132.0k | 210 | 0 | 22 | C | 11.0 |
| [kubernetes/apiserver](https://github.com/kubernetes/apiserver/)  | 131.3k | 8.4k | 32 | 692 | Go | 10.9 |
| [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers/)  | 130.8k | 16.9k | 6.1k | 6.8k | Go | 10.9 |
| [ytsaurus/ytsaurus-ui](https://github.com/ytsaurus/ytsaurus-ui/)  | 130.7k | 2.1k | 1.1k | 42 | TypeScript | 10.9 |
| [siderolabs/talos](https://github.com/siderolabs/talos/)  | 130.2k | 5.8k | 7.5k | 8.8k | Go | 10.8 |
| [sermant-io/Sermant](https://github.com/sermant-io/Sermant/)  | 126.1k | 1.9k | 1.1k | 1.4k | Java | 10.5 |
| [etcd-io/etcd](https://github.com/etcd-io/etcd/)  | 124.4k | 23.9k | 13.2k | 50.5k | Go | 10.4 |
| [cncf/demo](https://github.com/cncf/demo/)  | 122.9k | 670 | 161 | 78 | JavaScript | 10.2 |
| [armadaproject/armada](https://github.com/armadaproject/armada/)  | 119.1k | 3.3k | 3.7k | 557 | Go | 9.9 |
| [cortexproject/cortex](https://github.com/cortexproject/cortex/)  | 119.0k | 5.9k | 4.9k | 5.7k | Go | 9.9 |
| [artifacthub/hub](https://github.com/artifacthub/hub/)  | 116.7k | 2.0k | 3.5k | 1.9k | TypeScript | 9.7 |
| [parallaxsecond/rust-tss-esapi](https://github.com/parallaxsecond/rust-tss-esapi/)  | 116.6k | 896 | 425 | 111 | Rust | 9.7 |
| [thanos-io/thanos](https://github.com/thanos-io/thanos/)  | 113.2k | 4.5k | 5.1k | 13.7k | Go | 9.4 |
| [apache/aurora](https://github.com/apache/aurora/)  | 113.0k | 4.1k | 71 | 635 | Java | 9.4 |
| [projectcontour/contour](https://github.com/projectcontour/contour/)  | 111.5k | 5.1k | 4.9k | 3.8k | Go | 9.3 |
| [netdata/go.d.plugin](https://github.com/netdata/go.d.plugin/)  | 110.3k | 1.8k | 1.3k | 166 | Go | 9.2 |
| [openelb/website](https://github.com/openelb/website/)  | 107.6k | 183 | 97 | 3 | JavaScript | 9.0 |
| [moby/buildkit](https://github.com/moby/buildkit/)  | 107.2k | 8.3k | 3.8k | 9.4k | Go | 8.9 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 107.0k | 3.1k | 3.1k | 332 | Go | 8.9 |
| [apache/synapse](https://github.com/apache/synapse/)  | 106.1k | 4.4k | 85 | 75 | Java | 8.8 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105.9k | 3.2k | 4.5k | 1.6k | Go | 8.8 |
| [googlecontainertools/kpt-config-sy~](https://github.com/googlecontainertools/kpt-config-sync/)  | 105.0k | 983 | 1.2k | 222 | Go | 8.8 |
| [kubeflow/kfp-tekton-backend](https://github.com/kubeflow/kfp-tekton-backend/)  | 104.3k | 14 | 71 | 8 | TypeScript | 8.7 |
| [koordinator-sh/koordetector](https://github.com/koordinator-sh/koordetector/)  | 104.2k | 12 | 31 | 11 | C Header | 8.7 |
| [VictoriaMetrics/VictoriaLogs](https://github.com/VictoriaMetrics/VictoriaLogs/)  | 103.6k | 11.5k | 134 | 472 | Go | 8.6 |
| [pipe-cd/pipecd](https://github.com/pipe-cd/pipecd/)  | 103.3k | 4.9k | 4.9k | 1.2k | Go | 8.6 |
| [kubernetes-sigs/headlamp](https://github.com/kubernetes-sigs/headlamp/)  | 102.2k | 7.0k | 2.7k | 4.5k | TypeScript | 8.5 |
| [apache/ode](https://github.com/apache/ode/)  | 101.8k | 3.8k | 8 | 45 | Java | 8.5 |
| [cockroachdb/pebble](https://github.com/cockroachdb/pebble/)  | 101.1k | 5.2k | 3.9k | 5.5k | Go | 8.4 |
| [perses/perses](https://github.com/perses/perses/)  | 101.0k | 2.9k | 2.9k | 1.6k | TypeScript | 8.4 |
| [dapr/dapr](https://github.com/dapr/dapr/)  | 101.0k | 6.4k | 4.9k | 25.2k | Go | 8.4 |
| [cert-manager/cert-manager](https://github.com/cert-manager/cert-manager/)  | 100.4k | 10.0k | 4.2k | 13.2k | Go | 8.4 |
| [openebs/maya](https://github.com/openebs/maya/)  | 100.1k | 1.8k | 1.7k | 180 | Go | 8.3 |
| [apache/apr](https://github.com/apache/apr/)  | 99.3k | 9.6k | 62 | 510 | C | 8.3 |
| [apache/continuum](https://github.com/apache/continuum/)  | 99.2k | 4.9k | 33 | 7 | Java | 8.3 |
| [openkruise/kruise](https://github.com/openkruise/kruise/)  | 98.1k | 1.1k | 1.4k | 5.0k | Go | 8.2 |
| [emissary-ingress/emissary](https://github.com/emissary-ingress/emissary/)  | 97.8k | 18.3k | 4.0k | 4.5k | Python | 8.2 |
| [keylime/tpm4720-keylime](https://github.com/keylime/tpm4720-keylime/)  | 97.6k | 10 | 0 | 6 | C | 8.1 |
| [moby/containerd](https://github.com/moby/containerd/)  | 97.5k | 8.4k | 2 | 22 | Go | 8.1 |
| [knative/serving](https://github.com/knative/serving/)  | 96.9k | 9.5k | 11.1k | 5.9k | Go | 8.1 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 4.8k | 5.4k | 2.3k | Java | 8.0 |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 95.9k | 3.9k | 324 | 3.1k | Go | 8.0 |
| [apache/spamassassin](https://github.com/apache/spamassassin/)  | 95.9k | 30.8k | 7 | 313 | Perl | 8.0 |
| [rook/rook](https://github.com/rook/rook/)  | 95.6k | 12.9k | 10.1k | 13.1k | Go | 8.0 |
| [karmada-io/karmada](https://github.com/karmada-io/karmada/)  | 94.6k | 7.0k | 4.8k | 5.1k | Go | 7.9 |
| [elastic/logstash](https://github.com/elastic/logstash/)  | 92.3k | 11.2k | 11.2k | 14.7k | Ruby | 7.7 |
| [kusionstack/karpor](https://github.com/kusionstack/karpor/)  | 92.3k | 434 | 527 | 1.5k | Go | 7.7 |
| [v6d-io/v6d](https://github.com/v6d-io/v6d/)  | 92.2k | 1.5k | 1.4k | 926 | C++ | 7.7 |
| [tremor-rs/tremor-runtime](https://github.com/tremor-rs/tremor-runtime/)  | 91.8k | 4.5k | 2.3k | 913 | Rust | 7.7 |
| [ovn-org/ovn](https://github.com/ovn-org/ovn/)  | 91.7k | 22.2k | 118 | 637 | C | 7.6 |
| [apache/commons-lang](https://github.com/apache/commons-lang/)  | 91.1k | 8.9k | 1.5k | 2.9k | Java | 7.6 |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90.1k | 26.3k | 11.4k | 31.0k | Go | 7.5 |
| [kubeclipper/kubeclipper](https://github.com/kubeclipper/kubeclipper/)  | 90.0k | 427 | 684 | 395 | Go | 7.5 |
| [apache/maven](https://github.com/apache/maven/)  | 89.9k | 15.8k | 2.8k | 4.8k | Java | 7.5 |
| [dragonflyoss/dragonfly](https://github.com/dragonflyoss/dragonfly/)  | 88.7k | 2.9k | 3.5k | 2.8k | Go | 7.4 |
| [meshery/meshery](https://github.com/meshery/meshery/)  | 87.6k | 46.9k | 11.0k | 8.6k | Go | 7.3 |
| [cubefs/cubefs-blobstore](https://github.com/cubefs/cubefs-blobstore/)  | 87.4k | 82 | 3 | 17 | Go | 7.3 |
| [wasmedge/WasmEdge](https://github.com/wasmedge/WasmEdge/)  | 87.2k | 4.1k | 2.6k | 10.2k | C++ | 7.3 |
| [operator-framework/operator-lifecy~](https://github.com/operator-framework/operator-lifecycle-manager/)  | 87.0k | 4.7k | 2.8k | 1.8k | Go | 7.2 |
| [knative/eventing](https://github.com/knative/eventing/)  | 87.0k | 4.4k | 5.4k | 1.5k | Go | 7.2 |
| [moby/swarmkit](https://github.com/moby/swarmkit/)  | 86.8k | 4.4k | 2.4k | 3.5k | Go | 7.2 |
| [kcl-lang/kcl](https://github.com/kcl-lang/kcl/)  | 86.8k | 1.2k | 1.2k | 2.2k | Rust | 7.2 |
| [go-faster/minikube](https://github.com/go-faster/minikube/)  | 86.6k | 23.3k | 0 | 0 | Go | 7.2 |
| [fluent/onigmo](https://github.com/fluent/onigmo/)  | 86.4k | 1.2k | 7 | 6 | C Header | 7.2 |
| [dapr/components-contrib](https://github.com/dapr/components-contrib/)  | 86.0k | 3.5k | 2.4k | 582 | Go | 7.2 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84.4k | 7.0k | 9.0k | 285 | Go | 7.0 |
| [serverless-devs/serverless-registr~](https://github.com/serverless-devs/serverless-registry-website/)  | 82.8k | 31 | 0 | 3 | JavaScript | 6.9 |
| [argoproj/argo-rollouts](https://github.com/argoproj/argo-rollouts/)  | 82.6k | 2.0k | 2.7k | 3.3k | Go | 6.9 |
| [apache/struts-sandbox](https://github.com/apache/struts-sandbox/)  | 82.2k | 1.9k | 13 | 5 | Java | 6.8 |
| [linkerd/linkerd2](https://github.com/linkerd/linkerd2/)  | 82.0k | 8.3k | 9.7k | 11.1k | Go | 6.8 |
| [tremor-rs/tremor-www-docs](https://github.com/tremor-rs/tremor-www-docs/)  | 80.6k | 405 | 153 | 24 | JavaScript | 6.7 |
| [fluent/fluentd](https://github.com/fluent/fluentd/)  | 80.5k | 7.1k | 2.3k | 13.3k | Ruby | 6.7 |
| [apache/httpcomponents-core](https://github.com/apache/httpcomponents-core/)  | 80.4k | 4.0k | 564 | 367 | Java | 6.7 |
| [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker/)  | 79.7k | 7.4k | 3.9k | 30.7k | Rust | 6.6 |
| [open-policy-agent/eopa](https://github.com/open-policy-agent/eopa/)  | 79.3k | 2.4k | 169 | 20 | Go | 6.6 |
| [werf/werf](https://github.com/werf/werf/)  | 79.1k | 14.9k | 6.0k | 4.5k | Go | 6.6 |
| [kubevirt/containerized-data-import~](https://github.com/kubevirt/containerized-data-importer/)  | 79.1k | 3.0k | 3.2k | 491 | Go | 6.6 |
| [litmuschaos/litmus](https://github.com/litmuschaos/litmus/)  | 78.7k | 3.0k | 3.5k | 4.9k | TypeScript | 6.6 |
| [fluent/fluentbit-website-v3](https://github.com/fluent/fluentbit-website-v3/)  | 78.7k | 386 | 27 | 5 | JavaScript | 6.6 |
| [kubernetes/ingress-gce](https://github.com/kubernetes/ingress-gce/)  | 78.0k | 6.0k | 2.5k | 1.3k | Go | 6.5 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77.2k | 6.5k | 4.3k | 324 | Go | 6.4 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76.5k | 1.3k | 1.8k | 4.2k | Go | 6.4 |
| [metrico/loki-apache](https://github.com/metrico/loki-apache/)  | 76.4k | 9 | 4 | 0 | Go | 6.4 |
| [kubevela/velaux](https://github.com/kubevela/velaux/)  | 75.2k | 690 | 623 | 146 | TypeScript | 6.3 |
| [devspace-sh/devspace](https://github.com/devspace-sh/devspace/)  | 74.9k | 6.1k | 1.9k | 4.8k | Go | 6.2 |
| [chaos-mesh/chaos-mesh](https://github.com/chaos-mesh/chaos-mesh/)  | 74.8k | 1.9k | 2.8k | 7.3k | Go | 6.2 |
| [docker/cli](https://github.com/docker/cli/)  | 74.2k | 12.0k | 4.6k | 5.5k | Go | 6.2 |
| [apache/jspwiki](https://github.com/apache/jspwiki/)  | 74.0k | 9.5k | 394 | 110 | Java | 6.2 |
| [easegress-io/easegress](https://github.com/easegress-io/easegress/)  | 74.0k | 2.1k | 1.3k | 5.9k | Go | 6.2 |
| [openservicemesh/osm](https://github.com/openservicemesh/osm/)  | 73.9k | 4.5k | 3.6k | 2.6k | Go | 6.2 |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73.8k | 6.5k | 1.2k | 620 | Go | 6.2 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73.6k | 5.3k | 3.9k | 687 | Go | 6.1 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73.6k | 4.0k | 1.3k | 1.6k | Java | 6.1 |
| [apache/maven-sandbox](https://github.com/apache/maven-sandbox/)  | 73.4k | 2.3k | 2 | 4 | Java | 6.1 |
| [bfenetworks/bfe](https://github.com/bfenetworks/bfe/)  | 73.3k | 1.2k | 775 | 6.2k | Go | 6.1 |
| [buildpacks/pack-private](https://github.com/buildpacks/pack-private/)  | 73.3k | 4.5k | 58 | 0 | Go | 6.1 |
| [linkerd/linkerd](https://github.com/linkerd/linkerd/)  | 73.3k | 1.4k | 1.4k | 5.3k | Scala | 6.1 |
| [grafana/k6](https://github.com/grafana/k6/)  | 72.9k | 10.3k | 2.7k | 28.9k | Go | 6.1 |
| [apache/httpcomponents-client](https://github.com/apache/httpcomponents-client/)  | 72.9k | 3.9k | 725 | 1.5k | Java | 6.1 |
| [megaease/easegress](https://github.com/megaease/easegress/)  | 72.9k | 2.0k | 958 | 5.6k | Go | 6.1 |
| [kubeflow/kfserving-lts](https://github.com/kubeflow/kfserving-lts/)  | 72.5k | 795 | 77 | 11 | Python | 6.0 |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72.4k | 4.8k | 5.3k | 2.4k | TypeScript | 6.0 |
| [dragonflyoss/Dragonfly2](https://github.com/dragonflyoss/Dragonfly2/)  | 72.3k | 2.9k | 3.5k | 2.8k | Go | 6.0 |
| [go-faster/porto](https://github.com/go-faster/porto/)  | 71.9k | 121 | 0 | 1 | C++ | 6.0 |
| [jaegertracing/jaeger](https://github.com/jaegertracing/jaeger/)  | 71.9k | 3.8k | 4.8k | 21.9k | Go | 6.0 |
| [in-toto/toto-pip](https://github.com/in-toto/toto-pip/)  | 71.5k | 4.9k | 0 | 0 | Python | 6.0 |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71.3k | 633 | 232 | 1.1k | JavaScript | 5.9 |
| [falcosecurity/elftoolchain](https://github.com/falcosecurity/elftoolchain/)  | 71.3k | 10 | 2 | 0 | C | 5.9 |
| [project-zot/zot](https://github.com/project-zot/zot/)  | 71.1k | 1.4k | 2.7k | 1.4k | Go | 5.9 |
| [meilisearch/meilisearch](https://github.com/meilisearch/meilisearch/)  | 70.7k | 12.6k | 2.6k | 53.5k | Rust | 5.9 |
| [tikv/migration](https://github.com/tikv/migration/)  | 69.8k | 202 | 241 | 38 | Go | 5.8 |
| [parallaxsecond/parsec-website](https://github.com/parallaxsecond/parsec-website/)  | 69.4k | 5 | 3 | 1 | JavaScript | 5.8 |
| [cncf/obsolete-interactive-landscape](https://github.com/cncf/obsolete-interactive-landscape/)  | 69.1k | 364 | 0 | 4 | JavaScript | 5.8 |
| [longhorn/longhorn-manager](https://github.com/longhorn/longhorn-manager/)  | 68.9k | 4.8k | 4.2k | 189 | Go | 5.7 |
| [ten-nancy/porto](https://github.com/ten-nancy/porto/)  | 68.8k | 281 | 7 | 37 | C++ | 5.7 |
| [runatlantis/atlantis](https://github.com/runatlantis/atlantis/)  | 68.5k | 4.0k | 3.6k | 8.5k | Go | 5.7 |
| [kedacore/keda](https://github.com/kedacore/keda/)  | 68.3k | 2.7k | 3.6k | 9.5k | Go | 5.7 |
| [opencost/opencost](https://github.com/opencost/opencost/)  | 67.9k | 4.9k | 2.6k | 6.1k | Go | 5.7 |
| [kubeedge/kubeedge](https://github.com/kubeedge/kubeedge/)  | 67.5k | 6.1k | 4.0k | 7.2k | Go | 5.6 |
| [apache/commons-collections](https://github.com/apache/commons-collections/)  | 67.2k | 4.9k | 646 | 710 | Java | 5.6 |
| [nocalhost/nocalhost](https://github.com/nocalhost/nocalhost/)  | 66.8k | 3.6k | 1.3k | 1.9k | Go | 5.6 |
| [open-cluster-management-io/ocm](https://github.com/open-cluster-management-io/ocm/)  | 66.8k | 1.9k | 891 | 965 | Go | 5.6 |
| [cdk8s-team/cdk8s-plus-go](https://github.com/cdk8s-team/cdk8s-plus-go/)  | 65.8k | 914 | 0 | 7 | Go | 5.5 |
| [kcp-dev/kcp](https://github.com/kcp-dev/kcp/)  | 65.7k | 6.5k | 2.4k | 2.6k | Go | 5.5 |
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65.6k | 5.0k | 222 | 9.6k | Go | 5.5 |
| [vectordotdev/vrl](https://github.com/vectordotdev/vrl/)  | 65.5k | 10.4k | 1.1k | 195 | Rust | 5.5 |
| [apache/struts1](https://github.com/apache/struts1/)  | 65.4k | 5.3k | 11 | 32 | Java | 5.5 |
| [nats-io/nats.c](https://github.com/nats-io/nats.c/)  | 64.5k | 1.2k | 595 | 441 | C | 5.4 |
| [fluent/fluent-bit-website-old](https://github.com/fluent/fluent-bit-website-old/)  | 63.9k | 19 | 0 | 3 | JavaScript | 5.3 |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63.9k | 2.8k | 1.5k | 2.1k | Go | 5.3 |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63.9k | 1.1k | 1.2k | 4.4k | C# | 5.3 |
| [hwameistor/hwameistor](https://github.com/hwameistor/hwameistor/)  | 63.8k | 2.4k | 1.1k | 603 | Go | 5.3 |
| [megaease/easeagent](https://github.com/megaease/easeagent/)  | 63.5k | 760 | 257 | 568 | Java | 5.3 |
| [ent/ent](https://github.com/ent/ent/)  | 63.2k | 2.3k | 2.5k | 16.6k | Go | 5.3 |
| [ydb-platform/ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk/)  | 62.9k | 6.3k | 1.3k | 171 | Go | 5.2 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 62.7k | 3.5k | 246 | 19 | TypeScript | 5.2 |
| [werf/trdl-vault-actions](https://github.com/werf/trdl-vault-actions/)  | 62.2k | 29 | 12 | 0 | JavaScript | 5.2 |
| [kubeovn/kube-ovn](https://github.com/kubeovn/kube-ovn/)  | 61.8k | 4.9k | 4.4k | 2.2k | Go | 5.2 |
| [knative/pkg](https://github.com/knative/pkg/)  | 61.2k | 2.4k | 2.8k | 272 | Go | 5.1 |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61.1k | 2.3k | 1.9k | 263 | Go | 5.1 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 2.8k | 3.4k | 3.1k | TypeScript | 5.1 |
| [backstage/canon-storybook](https://github.com/backstage/canon-storybook/)  | 60.6k | 249 | 0 | 0 | JavaScript | 5.0 |
| [apache/abdera](https://github.com/apache/abdera/)  | 60.5k | 1.5k | 4 | 19 | Java | 5.0 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60.5k | 5.3k | 3.5k | 8.5k | Go | 5.0 |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60.3k | 3.6k | 2.6k | 209.5k | TypeScript | 5.0 |
| [kubernetes/legacy-cloud-providers](https://github.com/kubernetes/legacy-cloud-providers/)  | 60.2k | 2.0k | 0 | 52 | Go | 5.0 |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60.2k | 2.5k | 1.0k | 38 | Rust | 5.0 |
| [openyurtio/openyurt](https://github.com/openyurtio/openyurt/)  | 60.2k | 1.3k | 1.6k | 1.9k | Go | 5.0 |
| [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/)  | 60.1k | 3.3k | 32 | 886 | Go | 5.0 |
| [getporter/gh-action](https://github.com/getporter/gh-action/)  | 60.0k | 31 | 15 | 1 | JavaScript | 5.0 |
| [tarantool/graphql.0](https://github.com/tarantool/graphql.0/)  | 59.5k | 306 | 113 | 20 | JavaScript | 5.0 |
| [awslabs/aws-config-rules](https://github.com/awslabs/aws-config-rules/)  | 59.4k | 919 | 252 | 1.4k | Python | 5.0 |
| [ydb-platform/ydb-nodejs-genproto](https://github.com/ydb-platform/ydb-nodejs-genproto/)  | 59.3k | 31 | 22 | 2 | JavaScript | 5.0 |
| [apache/nutch](https://github.com/apache/nutch/)  | 58.7k | 3.5k | 850 | 3.1k | Java | 4.9 |
| [apache/maven-plugin-tools](https://github.com/apache/maven-plugin-tools/)  | 58.5k | 1.4k | 428 | 62 | JavaScript | 4.9 |
| [buildpacks/pack](https://github.com/buildpacks/pack/)  | 58.4k | 4.7k | 1.5k | 2.8k | Go | 4.9 |
| [loxilb-io/asn1c](https://github.com/loxilb-io/asn1c/)  | 57.9k | 2.9k | 0 | 0 | C | 4.8 |
| [awslabs/aws-elasticache-cluster-cl~](https://github.com/awslabs/aws-elasticache-cluster-client-libmemcached/)  | 57.8k | 59 | 10 | 22 | C++ | 4.8 |
| [keycloak/terraform-provider-keyclo~](https://github.com/keycloak/terraform-provider-keycloak/)  | 57.5k | 727 | 718 | 847 | Go | 4.8 |
| [envoyproxy/envoy-mobile](https://github.com/envoyproxy/envoy-mobile/)  | 57.5k | 1.8k | 2.2k | 562 | Java | 4.8 |
| [inclavare-containers/Shelter](https://github.com/inclavare-containers/Shelter/)  | 57.3k | 270 | 15 | 2 | Rust | 4.8 |
| [apache/servicemix3](https://github.com/apache/servicemix3/)  | 57.1k | 3.0k | 24 | 3 | Java | 4.8 |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56.9k | 8.2k | 8.0k | 19.0k | Go | 4.7 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.8k | 3.5k | 4.2k | 3.6k | C# | 4.7 |
| [apache/roller](https://github.com/apache/roller/)  | 56.7k | 4.8k | 142 | 126 | Java | 4.7 |
| [kanisterio/kanister](https://github.com/kanisterio/kanister/)  | 56.3k | 3.1k | 3.3k | 844 | Go | 4.7 |
| [googlecontainertools/jib](https://github.com/googlecontainertools/jib/)  | 56.3k | 2.3k | 2.6k | 13.3k | Java | 4.7 |
| [open-telemetry/otel-arrow-collector](https://github.com/open-telemetry/otel-arrow-collector/)  | 56.2k | 4.7k | 48 | 31 | Go | 4.7 |
| [apache/maven-shared](https://github.com/apache/maven-shared/)  | 56.1k | 3.4k | 26 | 28 | Java | 4.7 |
| [nats-io/nats.net.v1](https://github.com/nats-io/nats.net.v1/)  | 55.8k | 878 | 572 | 647 | C# | 4.7 |
| [ariga/atlas](https://github.com/ariga/atlas/)  | 55.7k | 2.7k | 2.8k | 7.4k | Go | 4.6 |
| [dragonflyoss/nydus](https://github.com/dragonflyoss/nydus/)  | 55.7k | 3.4k | 1.4k | 1.4k | Rust | 4.6 |
| [volcano-sh/volcano](https://github.com/volcano-sh/volcano/)  | 55.7k | 6.0k | 2.8k | 5.0k | Go | 4.6 |
| [kusionstack/kusion](https://github.com/kusionstack/kusion/)  | 55.6k | 772 | 834 | 1.2k | Go | 4.6 |
| [apache/harmony-jdktools](https://github.com/apache/harmony-jdktools/)  | 54.8k | 273 | 0 | 4 | Java | 4.6 |
| [prestodb/benchto](https://github.com/prestodb/benchto/)  | 54.6k | 356 | 47 | 23 | JavaScript | 4.5 |
| [moby/libnetwork](https://github.com/moby/libnetwork/)  | 54.1k | 3.1k | 1.9k | 2.2k | Go | 4.5 |
| [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy/)  | 54.1k | 3.8k | 4.2k | 2.1k | Rust | 4.5 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 53.7k | 2.6k | 1.5k | 924 | Rust | 4.5 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53.4k | 4.4k | 37 | 11 | Go | 4.5 |
| [dragonflyoss/image-service](https://github.com/dragonflyoss/image-service/)  | 53.2k | 3.0k | 1.1k | 795 | Rust | 4.4 |
| [gotd/td](https://github.com/gotd/td/)  | 52.6k | 4.1k | 1.3k | 2.0k | Go | 4.4 |
| [kcl-lang/kcl-py](https://github.com/kcl-lang/kcl-py/)  | 52.6k | 92 | 20 | 8 | Python | 4.4 |
| [serverlessworkflow/synapse](https://github.com/serverlessworkflow/synapse/)  | 52.6k | 1.3k | 326 | 287 | C# | 4.4 |
| [parallaxsecond/rust-cryptoki](https://github.com/parallaxsecond/rust-cryptoki/)  | 52.6k | 607 | 223 | 103 | Rust | 4.4 |
| [inclavare-containers/spdm-rs](https://github.com/inclavare-containers/spdm-rs/)  | 52.6k | 28 | 2 | 0 | Rust | 4.4 |
| [cartography-cncf/cartography](https://github.com/cartography-cncf/cartography/)  | 52.5k | 1.1k | 1.4k | 3.6k | Python | 4.4 |
| [rkt/rkt](https://github.com/rkt/rkt/)  | 51.9k | 5.6k | 2.5k | 8.8k | Go | 4.3 |
| [telepresenceio/telepresence](https://github.com/telepresenceio/telepresence/)  | 51.8k | 9.7k | 2.0k | 7.0k | Go | 4.3 |
| [brigadecore/brigade](https://github.com/brigadecore/brigade/)  | 51.5k | 2.0k | 1.2k | 2.4k | Go | 4.3 |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 2.5k | 2.2k | 843 | TypeScript | 4.3 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51.2k | 4.1k | 8 | 256 | Go | 4.3 |
| [kubeslice/nsm-sdk](https://github.com/kubeslice/nsm-sdk/)  | 51.2k | 1.0k | 2 | 0 | Go | 4.3 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 4.0k | 5.2k | 6.1k | Go | 4.3 |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50.8k | 2.5k | 999 | 231 | Go | 4.2 |
| [coredns/coredns](https://github.com/coredns/coredns/)  | 50.7k | 4.4k | 4.9k | 13.4k | Go | 4.2 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50.7k | 233 | 194 | 9 | Go | 4.2 |
| [keycloak/keycloak-client](https://github.com/keycloak/keycloak-client/)  | 50.5k | 93 | 93 | 38 | Java | 4.2 |
| [envoyproxy/toolshed](https://github.com/envoyproxy/toolshed/)  | 50.4k | 2.1k | 2.8k | 10 | Python | 4.2 |
| [trickstercache/trickster](https://github.com/trickstercache/trickster/)  | 50.4k | 1.1k | 614 | 2.1k | Go | 4.2 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector/)  | 50.3k | 7.6k | 9.8k | 5.8k | Go | 4.2 |
| [werf/trdl-actions](https://github.com/werf/trdl-actions/)  | 49.3k | 23 | 9 | 0 | JavaScript | 4.1 |
| [nats-io/nats.java](https://github.com/nats-io/nats.java/)  | 49.2k | 2.3k | 963 | 627 | Java | 4.1 |
| [getporter/porter](https://github.com/getporter/porter/)  | 49.1k | 3.6k | 2.1k | 1.4k | Go | 4.1 |
| [kubernetes-sigs/multi-tenancy](https://github.com/kubernetes-sigs/multi-tenancy/)  | 48.9k | 2.3k | 1.1k | 930 | Go | 4.1 |
| [apache/commons-io](https://github.com/apache/commons-io/)  | 48.9k | 5.9k | 793 | 1.0k | Java | 4.1 |
| [envoyproxy/pytooling](https://github.com/envoyproxy/pytooling/)  | 48.9k | 615 | 619 | 6 | Python | 4.1 |
| [nats-io/nats-streaming-server](https://github.com/nats-io/nats-streaming-server/)  | 48.9k | 1.7k | 662 | 2.5k | Go | 4.1 |
| [jaegertracing/jaeger-ui](https://github.com/jaegertracing/jaeger-ui/)  | 48.9k | 2.1k | 2.5k | 1.3k | JavaScript | 4.1 |
| [cilium/tetragon](https://github.com/cilium/tetragon/)  | 48.5k | 5.4k | 3.5k | 4.2k | Go | 4.0 |
| [external-secrets/external-secrets](https://github.com/external-secrets/external-secrets/)  | 48.4k | 3.9k | 3.3k | 5.9k | Go | 4.0 |
| [grpc/grpc-swift](https://github.com/grpc/grpc-swift/)  | 48.3k | 2.0k | 1.5k | 2.2k | Swift | 4.0 |
| [slimtoolkit/slim](https://github.com/slimtoolkit/slim/)  | 47.4k | 1.1k | 389 | 22.4k | Go | 4.0 |
| [apache/tapestry3](https://github.com/apache/tapestry3/)  | 47.1k | 2.0k | 8 | 5 | Java | 3.9 |
| [kptdev/kpt](https://github.com/kptdev/kpt/)  | 46.8k | 2.7k | 2.4k | 1.8k | Go | 3.9 |
| [paralus/dashboard](https://github.com/paralus/dashboard/)  | 46.8k | 395 | 247 | 14 | JavaScript | 3.9 |
| [nats-io/nats.net](https://github.com/nats-io/nats.net/)  | 46.8k | 747 | 430 | 566 | C# | 3.9 |
| [operator-framework/operator-sdk](https://github.com/operator-framework/operator-sdk/)  | 46.7k | 3.4k | 4.2k | 7.5k | Go | 3.9 |
| [knative/client](https://github.com/knative/client/)  | 46.6k | 1.3k | 1.5k | 364 | Go | 3.9 |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46.5k | 4.5k | 5 | 1 | Go | 3.9 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46.4k | 241 | 358 | 20 | C | 3.9 |
| [vitessio/website](https://github.com/vitessio/website/)  | 45.9k | 3.4k | 1.6k | 60 | JavaScript | 3.8 |
| [networkservicemesh/sdk](https://github.com/networkservicemesh/sdk/)  | 45.9k | 1.1k | 1.3k | 34 | Go | 3.8 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45.8k | 1.2k | 1.1k | 52 | JavaScript | 3.8 |
| [cri-o/cri-o](https://github.com/cri-o/cri-o/)  | 45.8k | 10.3k | 7.8k | 5.5k | Go | 3.8 |
| [headlamp-k8s/headlamp](https://github.com/headlamp-k8s/headlamp/)  | 45.4k | 4.9k | 1.8k | 2.6k | TypeScript | 3.8 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45.3k | 1.2k | 595 | 171 | Go | 3.8 |
| [apache/servicemix4-bundles](https://github.com/apache/servicemix4-bundles/)  | 44.9k | 7.6k | 371 | 17 | Java | 3.7 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44.9k | 5.9k | 7.3k | 15.2k | Go | 3.7 |
| [cockroachdb/gostdlib](https://github.com/cockroachdb/gostdlib/)  | 44.7k | 10 | 5 | 3 | Go | 3.7 |
| [kubeflow/kubeflow](https://github.com/kubeflow/kubeflow/)  | 44.6k | 2.6k | 3.8k | 15.2k | TypeScript | 3.7 |
| [kubeflow/arena](https://github.com/kubeflow/arena/)  | 44.0k | 618 | 1.2k | 797 | Go | 3.7 |
| [skooner-k8s/skooner](https://github.com/skooner-k8s/skooner/)  | 43.8k | 442 | 256 | 1.4k | JavaScript | 3.6 |
| [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper/)  | 43.7k | 1.9k | 2.6k | 4.0k | Go | 3.6 |
| [flatcar/mantle](https://github.com/flatcar/mantle/)  | 43.7k | 4.4k | 667 | 39 | Go | 3.6 |
| [containerd/nerdctl](https://github.com/containerd/nerdctl/)  | 43.6k | 5.6k | 3.0k | 9.4k | Go | 3.6 |
| [tauri-apps/tauri](https://github.com/tauri-apps/tauri/)  | 43.4k | 5.7k | 6.0k | 97.1k | Rust | 3.6 |
| [notaryproject/notary](https://github.com/notaryproject/notary/)  | 43.3k | 3.0k | 1.0k | 3.3k | Go | 3.6 |
| [kptdev/krm-functions-catalog](https://github.com/kptdev/krm-functions-catalog/)  | 43.3k | 739 | 1.2k | 94 | TypeScript | 3.6 |
| [inclavare-containers/inclavare-con~](https://github.com/inclavare-containers/inclavare-containers/)  | 43.0k | 1.4k | 924 | 622 | C | 3.6 |
| [oscal-compass/compliance-trestle](https://github.com/oscal-compass/compliance-trestle/)  | 42.9k | 1.3k | 1.1k | 207 | Python | 3.6 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 42.9k | 2.6k | 2.3k | 908 | Python | 3.6 |
| [distribution/distribution](https://github.com/distribution/distribution/)  | 42.9k | 5.7k | 2.3k | 9.9k | Go | 3.6 |
| [tikv/client-go](https://github.com/tikv/client-go/)  | 42.6k | 925 | 1.4k | 329 | Go | 3.5 |
| [open-telemetry/opentelemetry-netwo~](https://github.com/open-telemetry/opentelemetry-network/)  | 42.6k | 510 | 308 | 385 | C++ | 3.5 |
| [crossplane/crossplane](https://github.com/crossplane/crossplane/)  | 42.5k | 8.1k | 4.0k | 10.9k | Go | 3.5 |
| [istio/old_mixer_repo](https://github.com/istio/old_mixer_repo/)  | 42.5k | 741 | 1.1k | 67 | Go | 3.5 |
| [grpc/grpc-node](https://github.com/grpc/grpc-node/)  | 42.3k | 5.2k | 1.8k | 4.7k | TypeScript | 3.5 |
| [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)  | 42.1k | 3.2k | 2.1k | 2.8k | Go | 3.5 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 42.0k | 52 | 3 | 6 | JavaScript | 3.5 |
| [helm/helm](https://github.com/helm/helm/)  | 41.7k | 8.9k | 6.2k | 28.6k | Go | 3.5 |
| [carvel-dev/ytt](https://github.com/carvel-dev/ytt/)  | 41.7k | 1.3k | 394 | 1.8k | Go | 3.5 |
| [operator-framework/java-operator-s~](https://github.com/operator-framework/java-operator-sdk/)  | 41.6k | 3.0k | 2.0k | 876 | Java | 3.5 |
| [open-telemetry/opentelemetry-ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4k | 268 | 112 | 85 | C++ | 3.5 |
| [operator-framework/operator-regist~](https://github.com/operator-framework/operator-registry/)  | 41.4k | 1.6k | 1.4k | 224 | Go | 3.5 |
| [kubeslice/networkservicemesh](https://github.com/kubeslice/networkservicemesh/)  | 41.3k | 1.4k | 61 | 0 | Go | 3.4 |
| [networkservicemesh/networkservicem~](https://github.com/networkservicemesh/networkservicemesh/)  | 41.1k | 1.4k | 1.7k | 504 | Go | 3.4 |
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40.8k | 2.3k | 22 | 193 | C++ | 3.4 |
| [spidernet-io/spiderpool](https://github.com/spidernet-io/spiderpool/)  | 40.6k | 4.0k | 3.4k | 608 | Go | 3.4 |
| [open-telemetry/otel-arrow](https://github.com/open-telemetry/otel-arrow/)  | 40.6k | 1.2k | 940 | 236 | Go | 3.4 |
| [apache/sanselan](https://github.com/apache/sanselan/)  | 40.5k | 315 | 3 | 35 | Java | 3.4 |
| [capsule-rs/capsule](https://github.com/capsule-rs/capsule/)  | 40.4k | 229 | 117 | 347 | Rust | 3.4 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40.4k | 12.8k | 6.6k | 1.1k | TypeScript | 3.4 |
| [googlecontainertools/kpt-functions~](https://github.com/googlecontainertools/kpt-functions-catalog/)  | 40.3k | 711 | 1.1k | 83 | TypeScript | 3.4 |
| [keptn/tutorials](https://github.com/keptn/tutorials/)  | 40.2k | 624 | 199 | 10 | JavaScript | 3.4 |
| [dapr/java-sdk](https://github.com/dapr/java-sdk/)  | 39.7k | 710 | 949 | 283 | Java | 3.3 |
| [tikv/client-java](https://github.com/tikv/client-java/)  | 39.4k | 295 | 559 | 116 | Java | 3.3 |
| [ytsaurus/ytsaurus-spyt](https://github.com/ytsaurus/ytsaurus-spyt/)  | 39.2k | 471 | 51 | 19 | Scala | 3.3 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39.1k | 177 | 142 | 12 | Go | 3.3 |
| [ydb-platform/yoj-project](https://github.com/ydb-platform/yoj-project/)  | 39.1k | 393 | 131 | 16 | Java | 3.2 |
| [nats-io/nats.go](https://github.com/nats-io/nats.go/)  | 38.8k | 2.2k | 1.1k | 6.2k | Go | 3.2 |
| [open-telemetry/opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python/)  | 38.5k | 2.0k | 2.5k | 2.2k | Python | 3.2 |
| [spectralops/spectral-goat](https://github.com/spectralops/spectral-goat/)  | 38.3k | 3 | 0 | 2 | JavaScript | 3.2 |
| [vscode-kubernetes-tools/vscode-kub~](https://github.com/vscode-kubernetes-tools/vscode-kubernetes-tools/)  | 38.1k | 915 | 895 | 736 | TypeScript | 3.2 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38.0k | 4.9k | 3.2k | 8.7k | Go | 3.2 |
| [open-telemetry/opentelemetry-cpp-c~](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 37.4k | 333 | 390 | 144 | Python | 3.1 |
| [metallb/metallb](https://github.com/metallb/metallb/)  | 37.4k | 2.4k | 1.6k | 7.8k | Go | 3.1 |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37.3k | 1.0k | 381 | 370 | Go | 3.1 |
| [uber/kraken](https://github.com/uber/kraken/)  | 37.2k | 929 | 332 | 6.5k | Go | 3.1 |
| [jquery/jquery](https://github.com/jquery/jquery/)  | 37.2k | 6.8k | 3.1k | 59.6k | JavaScript | 3.1 |
| [longhorn/longhorn-tests](https://github.com/longhorn/longhorn-tests/)  | 36.7k | 2.7k | 2.8k | 20 | Python | 3.1 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36.7k | 4.6k | 1.2k | 1.9k | Go | 3.0 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36.4k | 40 | 10 | 3 | C | 3.0 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36.3k | 298 | 358 | 83 | Go | 3.0 |
| [xline-kv/Xline](https://github.com/xline-kv/Xline/)  | 36.1k | 1.4k | 821 | 681 | Rust | 3.0 |
| [gravity-ui/markdown-editor](https://github.com/gravity-ui/markdown-editor/)  | 36.1k | 760 | 785 | 324 | TypeScript | 3.0 |
| [microcks/microcks](https://github.com/microcks/microcks/)  | 36.0k | 2.2k | 375 | 1.7k | Java | 3.0 |
| [apache/xmlgraphics-commons](https://github.com/apache/xmlgraphics-commons/)  | 36.0k | 659 | 23 | 20 | Java | 3.0 |
| [argoproj/argo-events](https://github.com/argoproj/argo-events/)  | 35.9k | 2.0k | 2.4k | 2.6k | Go | 3.0 |
| [awslabs/aws-ms-deploy-assistant](https://github.com/awslabs/aws-ms-deploy-assistant/)  | 35.8k | 14 | 6 | 8 | JavaScript | 3.0 |
| [metrico/qryn-view](https://github.com/metrico/qryn-view/)  | 35.7k | 1.7k | 568 | 56 | TypeScript | 3.0 |
| [paralus/paralus](https://github.com/paralus/paralus/)  | 35.7k | 684 | 248 | 1.1k | Go | 3.0 |
| [pion/webrtc](https://github.com/pion/webrtc/)  | 35.6k | 2.3k | 2.0k | 15.4k | Go | 3.0 |
| [apache/logging-log4j1](https://github.com/apache/logging-log4j1/)  | 35.6k | 3.2k | 22 | 868 | Java | 3.0 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35.5k | 1.7k | 859 | 509 | Go | 3.0 |
| [openebs/openebs-e2e](https://github.com/openebs/openebs-e2e/)  | 35.4k | 401 | 215 | 0 | Go | 3.0 |
| [ydb-platform/tpcc-postgres](https://github.com/ydb-platform/tpcc-postgres/)  | 35.3k | 2.2k | 25 | 7 | Java | 2.9 |
| [radius-project/eshop](https://github.com/radius-project/eshop/)  | 35.2k | 2 | 13 | 0 | C# | 2.9 |
| [buildpacks/lifecycle-private](https://github.com/buildpacks/lifecycle-private/)  | 35.2k | 1.9k | 25 | 0 | Go | 2.9 |
| [netdata/dashboard](https://github.com/netdata/dashboard/)  | 35.0k | 1.8k | 396 | 70 | JavaScript | 2.9 |
| [cubefs/cubefs-dashboard](https://github.com/cubefs/cubefs-dashboard/)  | 34.8k | 40 | 21 | 26 | Go | 2.9 |
| [opencurve/curve-tgt](https://github.com/opencurve/curve-tgt/)  | 34.8k | 2.1k | 0 | 30 | C | 2.9 |
| [shipwright-io/build](https://github.com/shipwright-io/build/)  | 34.5k | 2.7k | 1.5k | 749 | Go | 2.9 |
| [kubedl-io/kubedl](https://github.com/kubedl-io/kubedl/)  | 34.4k | 389 | 201 | 531 | Go | 2.9 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34.2k | 2.0k | 927 | 3.2k | Go | 2.9 |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34.2k | 13 | 7 | 7 | JavaScript | 2.9 |
| [ratify-project/ratify](https://github.com/ratify-project/ratify/)  | 33.9k | 1.3k | 1.5k | 237 | Go | 2.8 |
| [kmesh-net/kmesh](https://github.com/kmesh-net/kmesh/)  | 33.7k | 3.0k | 1.0k | 678 | Go | 2.8 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33.7k | 2.0k | 789 | 142 | Go | 2.8 |
| [containerssh/libcontainerssh](https://github.com/containerssh/libcontainerssh/)  | 33.7k | 361 | 664 | 64 | Go | 2.8 |
| [kcp-dev/contrib-tmc](https://github.com/kcp-dev/contrib-tmc/)  | 33.6k | 51 | 17 | 7 | Go | 2.8 |
| [flatcar/ignition](https://github.com/flatcar/ignition/)  | 33.6k | 1.6k | 45 | 35 | Go | 2.8 |
| [kubeflow/katib](https://github.com/kubeflow/katib/)  | 33.6k | 1.4k | 1.6k | 1.6k | Go | 2.8 |
| [carvel-dev/kapp-controller](https://github.com/carvel-dev/kapp-controller/)  | 33.6k | 1.7k | 1.2k | 295 | Go | 2.8 |
| [open-telemetry/opentelemetry-cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.6k | 1.7k | 1.9k | 1.1k | C++ | 2.8 |
| [apache/buildr](https://github.com/apache/buildr/)  | 33.5k | 4.4k | 52 | 141 | Ruby | 2.8 |
| [ovn-org/libovsdb](https://github.com/ovn-org/libovsdb/)  | 33.5k | 751 | 287 | 190 | Go | 2.8 |
| [metrico/gigapipe](https://github.com/metrico/gigapipe/)  | 33.4k | 2.0k | 451 | 1.5k | Go | 2.8 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.4k | 2.0k | 2.3k | 595 | C# | 2.8 |
| [opencontainers/runc](https://github.com/opencontainers/runc/)  | 33.3k | 6.3k | 2.7k | 10.1k | Go | 2.8 |
| [in-toto/github-action](https://github.com/in-toto/github-action/)  | 33.2k | 3 | 0 | 8 | JavaScript | 2.8 |
| [kubernetes/perf-tests](https://github.com/kubernetes/perf-tests/)  | 33.2k | 4.7k | 3.2k | 945 | Go | 2.8 |
| [awslabs/amazon-kinesis-client](https://github.com/awslabs/amazon-kinesis-client/)  | 33.2k | 466 | 646 | 608 | Java | 2.8 |
| [meshery/meshery.io](https://github.com/meshery/meshery.io/)  | 33.1k | 8.7k | 1.5k | 649 | JavaScript | 2.8 |
| [open-telemetry/opentelemetry-ebpf-~](https://github.com/open-telemetry/opentelemetry-ebpf-profiler/)  | 33.1k | 507 | 611 | 2.9k | Go | 2.8 |
| [buildpacks/lifecycle](https://github.com/buildpacks/lifecycle/)  | 32.9k | 2.0k | 1.1k | 198 | Go | 2.7 |
| [googlecontainertools/kpt-functions~](https://github.com/googlecontainertools/kpt-functions-sdk/)  | 32.7k | 558 | 668 | 53 | TypeScript | 2.7 |
| [goharbor/harbor-operator](https://github.com/goharbor/harbor-operator/)  | 32.5k | 1.1k | 579 | 377 | Go | 2.7 |
| [keptn/lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit/)  | 32.2k | 2.3k | 2.8k | 398 | Go | 2.7 |
| [cloudevents/sdk-go](https://github.com/cloudevents/sdk-go/)  | 32.0k | 1.1k | 794 | 912 | Go | 2.7 |
| [prometheus/alertmanager](https://github.com/prometheus/alertmanager/)  | 31.9k | 3.5k | 2.6k | 7.2k | Go | 2.7 |
| [ydb-platform/xorm](https://github.com/ydb-platform/xorm/)  | 31.9k | 1.7k | 55 | 0 | Go | 2.7 |
| [youki-dev/youki](https://github.com/youki-dev/youki/)  | 31.8k | 5.6k | 2.9k | 7.0k | Rust | 2.6 |
| [nats-io/nsc](https://github.com/nats-io/nsc/)  | 31.8k | 711 | 499 | 123 | Go | 2.6 |
| [kubeedge/sedna](https://github.com/kubeedge/sedna/)  | 31.7k | 543 | 266 | 524 | Python | 2.6 |
| [clusternet/clusternet](https://github.com/clusternet/clusternet/)  | 31.6k | 871 | 649 | 1.4k | Go | 2.6 |
| [kubevirt/hyperconverged-cluster-op~](https://github.com/kubevirt/hyperconverged-cluster-operator/)  | 31.4k | 2.3k | 3.6k | 168 | Go | 2.6 |
| [litmuschaos/test-tools](https://github.com/litmuschaos/test-tools/)  | 31.3k | 268 | 634 | 29 | C | 2.6 |
| [knative/func](https://github.com/knative/func/)  | 31.2k | 2.1k | 2.3k | 328 | Go | 2.6 |
| [open-cluster-management-io/multicl~](https://github.com/open-cluster-management-io/multicloud-operators-subscription/)  | 31.2k | 915 | 406 | 44 | Go | 2.6 |
| [pravega/pravega-client-rust](https://github.com/pravega/pravega-client-rust/)  | 31.0k | 208 | 241 | 31 | Rust | 2.6 |
| [falcosecurity/plugin-sdk-rs](https://github.com/falcosecurity/plugin-sdk-rs/)  | 31.0k | 560 | 69 | 11 | Rust | 2.6 |
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 820 | 975 | 865 | PHP | 2.6 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30.9k | 1.5k | 460 | 348 | Go | 2.6 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30.9k | 18 | 4 | 1 | C Header | 2.6 |
| [konveyor/move2kube-demos](https://github.com/konveyor/move2kube-demos/)  | 30.9k | 119 | 269 | 13 | JavaScript | 2.6 |
| [cockroachdb/docs](https://github.com/cockroachdb/docs/)  | 30.8k | 17.1k | 9.3k | 201 | Java | 2.6 |
| [kusionstack/kuperator](https://github.com/kusionstack/kuperator/)  | 30.7k | 246 | 252 | 94 | Go | 2.6 |
| [sealerio/sealer](https://github.com/sealerio/sealer/)  | 30.7k | 1.0k | 1.3k | 2.1k | Go | 2.6 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 1.8k | 2.1k | 2.3k | Rust | 2.5 |
| [oras-project/oras-go](https://github.com/oras-project/oras-go/)  | 30.5k | 499 | 647 | 229 | Go | 2.5 |
| [koordinator-sh/yarn-copilot](https://github.com/koordinator-sh/yarn-copilot/)  | 30.4k | 52 | 98 | 13 | Go | 2.5 |
| [loxilb-io/loxilb-ebpf](https://github.com/loxilb-io/loxilb-ebpf/)  | 30.3k | 577 | 71 | 32 | C | 2.5 |
| [superedge/superedge](https://github.com/superedge/superedge/)  | 30.2k | 410 | 409 | 1.1k | Go | 2.5 |
| [connectrpc/connect-es](https://github.com/connectrpc/connect-es/)  | 30.2k | 863 | 1.3k | 1.6k | TypeScript | 2.5 |
| [connectrpc/connect-dart](https://github.com/connectrpc/connect-dart/)  | 30.1k | 24 | 25 | 35 | Dart | 2.5 |
| [fluxcd/source-controller](https://github.com/fluxcd/source-controller/)  | 30.1k | 2.8k | 1.5k | 260 | Go | 2.5 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 30.1k | 3.8k | 2.6k | 400 | Go | 2.5 |
| [kube-logging/logging-operator](https://github.com/kube-logging/logging-operator/)  | 29.9k | 2.8k | 1.4k | 1.7k | Go | 2.5 |
| [docker/machine](https://github.com/docker/machine/)  | 29.9k | 3.5k | 2.0k | 6.6k | Go | 2.5 |
| [cilium/ebpf](https://github.com/cilium/ebpf/)  | 29.9k | 2.1k | 1.3k | 7.2k | Go | 2.5 |
| [metrico/opensearch-action](https://github.com/metrico/opensearch-action/)  | 29.9k | 3 | 0 | 3 | JavaScript | 2.5 |
| [fluxcd/flux](https://github.com/fluxcd/flux/)  | 29.9k | 5.3k | 1.9k | 6.9k | Go | 2.5 |
| [go-faster/yt](https://github.com/go-faster/yt/)  | 29.8k | 9 | 6 | 0 | Go | 2.5 |
| [konveyor/move2kube](https://github.com/konveyor/move2kube/)  | 29.7k | 807 | 922 | 402 | Go | 2.5 |
| [devfile/library](https://github.com/devfile/library/)  | 29.7k | 649 | 222 | 26 | Go | 2.5 |
| [nats-io/nats.deno](https://github.com/nats-io/nats.deno/)  | 29.7k | 928 | 631 | 176 | TypeScript | 2.5 |
| [m3db/m3cluster](https://github.com/m3db/m3cluster/)  | 29.6k | 238 | 227 | 20 | Go | 2.5 |
| [chaosblade-io/chaosblade-space-exp~](https://github.com/chaosblade-io/chaosblade-space-exploration/)  | 29.6k | 40 | 0 | 338 | Java | 2.5 |
| [kubevirt/project-infra](https://github.com/kubevirt/project-infra/)  | 29.5k | 3.9k | 4.2k | 31 | Go | 2.5 |
| [moby/hyperkit](https://github.com/moby/hyperkit/)  | 29.4k | 623 | 220 | 3.7k | C | 2.5 |
| [dexidp/dex](https://github.com/dexidp/dex/)  | 29.3k | 3.5k | 2.9k | 10.3k | Go | 2.4 |
| [ydb-platform/ycsb-ydb](https://github.com/ydb-platform/ycsb-ydb/)  | 29.1k | 1.4k | 3 | 0 | Java | 2.4 |
| [cubefs/compass](https://github.com/cubefs/compass/)  | 29.1k | 489 | 91 | 401 | Java | 2.4 |
| [apache/maven-mercury](https://github.com/apache/maven-mercury/)  | 29.0k | 769 | 0 | 2 | Java | 2.4 |
| [dragonflyoss/Dragonfly](https://github.com/dragonflyoss/Dragonfly/)  | 28.8k | 1.8k | 1.0k | 6.0k | Go | 2.4 |
| [dragonflyoss/dragonfly-archived](https://github.com/dragonflyoss/dragonfly-archived/)  | 28.8k | 1.8k | 1.0k | 6.0k | Go | 2.4 |
| [openkruise/rollouts](https://github.com/openkruise/rollouts/)  | 28.6k | 132 | 203 | 245 | Go | 2.4 |
| [kubevirt/web-ui-components](https://github.com/kubevirt/web-ui-components/)  | 28.5k | 416 | 500 | 8 | JavaScript | 2.4 |
| [litmuschaos/litmus-go](https://github.com/litmuschaos/litmus-go/)  | 28.5k | 516 | 655 | 77 | Go | 2.4 |
| [k3s-io/k3s](https://github.com/k3s-io/k3s/)  | 28.5k | 3.8k | 4.7k | 31.0k | Go | 2.4 |
| [jaegertracing/jaeger-operator](https://github.com/jaegertracing/jaeger-operator/)  | 28.5k | 1.5k | 1.9k | 1.1k | Go | 2.4 |
| [perses/community-dashboards](https://github.com/perses/community-dashboards/)  | 28.4k | 417 | 151 | 23 | Go | 2.4 |
| [loxilb-io/libbpf](https://github.com/loxilb-io/libbpf/)  | 28.4k | 6 | 0 | 1 | C | 2.4 |
| [cncf/devstatscode](https://github.com/cncf/devstatscode/)  | 28.3k | 556 | 20 | 48 | Go | 2.4 |
| [chaosblade-io/chaosblade-box-fe](https://github.com/chaosblade-io/chaosblade-box-fe/)  | 28.2k | 13 | 36 | 10 | TypeScript | 2.4 |
| [open-telemetry/opentelemetry-log-c~](https://github.com/open-telemetry/opentelemetry-log-collection/)  | 28.1k | 267 | 373 | 92 | Go | 2.3 |
| [projectriff/cli](https://github.com/projectriff/cli/)  | 28.0k | 1.5k | 263 | 6 | Go | 2.3 |
| [fnproject/fn](https://github.com/fnproject/fn/)  | 28.0k | 3.4k | 953 | 5.4k | Go | 2.3 |
| [metal3-io/cluster-api-provider-met~](https://github.com/metal3-io/cluster-api-provider-metal3/)  | 27.8k | 3.7k | 2.6k | 244 | Go | 2.3 |
| [notaryproject/ratify](https://github.com/notaryproject/ratify/)  | 27.7k | 1.7k | 1.9k | 279 | Go | 2.3 |
| [kubeflow/training-operator](https://github.com/kubeflow/training-operator/)  | 27.6k | 1.1k | 1.4k | 1.7k | Go | 2.3 |
| [kubernetes-sigs/alibaba-cloud-csi-~](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/)  | 27.6k | 3.1k | 1.3k | 566 | Go | 2.3 |
| [fluxcd/flagger](https://github.com/fluxcd/flagger/)  | 27.6k | 3.0k | 1.0k | 5.2k | Go | 2.3 |
| [fluxcd/pkg](https://github.com/fluxcd/pkg/)  | 27.5k | 2.1k | 929 | 53 | Go | 2.3 |
| [tinkerbell/tinkerbell](https://github.com/tinkerbell/tinkerbell/)  | 27.5k | 6.2k | 367 | 100 | Go | 2.3 |
| [envoyproxy/nighthawk](https://github.com/envoyproxy/nighthawk/)  | 27.4k | 881 | 1.1k | 390 | C++ | 2.3 |
| [kubernetes/release](https://github.com/kubernetes/release/)  | 27.3k | 7.1k | 3.4k | 491 | Go | 2.3 |
| [devfile/integration-tests](https://github.com/devfile/integration-tests/)  | 27.3k | 138 | 108 | 2 | Go | 2.3 |
| [open-telemetry/opentelemetry-go-co~](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.3k | 3.3k | 7.1k | 1.5k | Go | 2.3 |
| [awslabs/lambda-chef-node-cleanup](https://github.com/awslabs/lambda-chef-node-cleanup/)  | 27.2k | 86 | 11 | 72 | Python | 2.3 |
| [devstream-io/devstream](https://github.com/devstream-io/devstream/)  | 27.1k | 7 | 963 | 872 | Go | 2.3 |
| [konveyor/tackle2-ui](https://github.com/konveyor/tackle2-ui/)  | 27.1k | 1.5k | 2.2k | 7 | TypeScript | 2.3 |
| [antrea-io/nephe](https://github.com/antrea-io/nephe/)  | 27.1k | 240 | 262 | 9 | Go | 2.2 |
| [kudobuilder/kudo](https://github.com/kudobuilder/kudo/)  | 26.8k | 1.0k | 1.1k | 1.2k | Go | 2.2 |
| [cilium/hubble-ui](https://github.com/cilium/hubble-ui/)  | 26.8k | 546 | 931 | 475 | TypeScript | 2.2 |
| [ClickHouse/libhdfs3](https://github.com/ClickHouse/libhdfs3/)  | 26.7k | 126 | 56 | 41 | C++ | 2.2 |
| [kubevirt/vm-import-operator](https://github.com/kubevirt/vm-import-operator/)  | 26.6k | 900 | 416 | 16 | Go | 2.2 |
| [dapr/dotnet-sdk](https://github.com/dapr/dotnet-sdk/)  | 26.5k | 804 | 756 | 1.2k | C# | 2.2 |
| [apache/chainsaw](https://github.com/apache/chainsaw/)  | 26.4k | 934 | 2 | 18 | Java | 2.2 |
| [gravity-ui/charts](https://github.com/gravity-ui/charts/)  | 26.3k | 149 | 159 | 8 | JavaScript | 2.2 |
| [containernetworking/plugins](https://github.com/containernetworking/plugins/)  | 26.2k | 1.9k | 750 | 2.4k | Go | 2.2 |
| [tarantool/vshard](https://github.com/tarantool/vshard/)  | 26.0k | 701 | 138 | 107 | Lua | 2.2 |
| [open-telemetry/opentelemetry-swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9k | 1.0k | 584 | 289 | Swift | 2.2 |
| [cockroachdb/cockroach-operator](https://github.com/cockroachdb/cockroach-operator/)  | 25.8k | 672 | 575 | 305 | Go | 2.1 |
| [openfga/openfga](https://github.com/openfga/openfga/)  | 25.8k | 1.7k | 2.2k | 4.2k | Go | 2.1 |
| [carvel-dev/kapp](https://github.com/carvel-dev/kapp/)  | 25.7k | 1.0k | 704 | 1.0k | Go | 2.1 |
| [flatcar/nebraska](https://github.com/flatcar/nebraska/)  | 25.7k | 2.3k | 934 | 193 | Go | 2.1 |
| [longhorn/longhorn-ui](https://github.com/longhorn/longhorn-ui/)  | 25.6k | 1.1k | 959 | 114 | JavaScript | 2.1 |
| [grpc/grpc-swift-2](https://github.com/grpc/grpc-swift-2/)  | 25.6k | 2.1k | 12 | 53 | Swift | 2.1 |
| [pravega/www.pravega.io](https://github.com/pravega/www.pravega.io/)  | 25.6k | 13 | 0 | 0 | JavaScript | 2.1 |
| [netdata/libjudy](https://github.com/netdata/libjudy/)  | 25.5k | 25 | 4 | 33 | C | 2.1 |
| [envoyproxy/java-control-plane](https://github.com/envoyproxy/java-control-plane/)  | 25.5k | 332 | 387 | 308 | Protocol Buffers | 2.1 |
| [opengemini/pr-title-checker](https://github.com/opengemini/pr-title-checker/)  | 25.4k | 9 | 5 | 2 | JavaScript | 2.1 |
| [docker/docker-py](https://github.com/docker/docker-py/)  | 25.4k | 3.4k | 1.7k | 7.1k | Python | 2.1 |
| [grafana/grafana-plugin-sdk-go](https://github.com/grafana/grafana-plugin-sdk-go/)  | 25.4k | 905 | 1.2k | 237 | Go | 2.1 |
| [apache/maven-scm](https://github.com/apache/maven-scm/)  | 25.4k | 2.3k | 294 | 96 | Java | 2.1 |
| [kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)  | 25.3k | 1.1k | 600 | 1.2k | Go | 2.1 |
| [gravity-ui/uikit](https://github.com/gravity-ui/uikit/)  | 25.1k | 1.7k | 1.9k | 847 | TypeScript | 2.1 |
| [containerd/stargz-snapshotter](https://github.com/containerd/stargz-snapshotter/)  | 25.1k | 2.0k | 1.9k | 1.4k | Go | 2.1 |
| [kubernetes-sigs/kubefed](https://github.com/kubernetes-sigs/kubefed/)  | 25.0k | 2.7k | 978 | 2.5k | Go | 2.1 |
| [fluent/cmetrics](https://github.com/fluent/cmetrics/)  | 25.0k | 584 | 211 | 19 | C | 2.1 |


### CNCF Projects
| Project | SLOC | Commits | PRs | Stars | Language | Effort |
|---------|------|---------|-----|-------|----------|--------|
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 489.9k | 248.2k | 301.2k | Go | 286.8 |
| [envoyproxy](https://github.com/envoyproxy)  | 2.9M | 100.9k | 46.5k | 37.1k | JavaScript | 238.6 |
| [open-telemetry](https://github.com/open-telemetry)  | 2.3M | 128.4k | 138.2k | 60.4k | Go | 191.3 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.1M | 190.0k | 126.2k | 137.5k | Go | 175.5 |
| [grpc](https://github.com/grpc)  | 1.7M | 87.0k | 52.3k | 104.5k | C++ | 145.0 |
| [konveyor](https://github.com/konveyor)  | 1.6M | 14.3k | 14.6k | 1.2k | JavaScript | 136.4 |
| [inclavare-containers](https://github.com/inclavare-containers)  | 1.6M | 25.6k | 1.3k | 748 | C++ | 131.7 |
| [openkruise](https://github.com/openkruise)  | 1.4M | 2.2k | 2.5k | 6.5k | JavaScript | 113.0 |
| [kubewarden](https://github.com/kubewarden)  | 1.1M | 32.9k | 15.4k | 804 | JavaScript | 92.4 |
| [kubeflow](https://github.com/kubeflow)  | 936.3k | 26.8k | 32.2k | 33.7k | Go | 78.0 |
| [brigadecore](https://github.com/brigadecore)  | 872.9k | 6.3k | 2.8k | 3.0k | JavaScript | 72.7 |
| [tikv](https://github.com/tikv)  | 862.5k | 20.1k | 26.7k | 30.1k | Rust | 71.9 |
| [nats-io](https://github.com/nats-io)  | 858.5k | 50.2k | 20.8k | 44.9k | Go | 71.5 |
| [keycloak](https://github.com/keycloak)  | 849.2k | 44.6k | 31.8k | 38.0k | Java | 70.8 |
| [cdk8s-team](https://github.com/cdk8s-team)  | 805.9k | 36.8k | 38.1k | 5.2k | Go | 67.2 |
| [cilium](https://github.com/cilium)  | 733.4k | 61.9k | 46.4k | 43.6k | Go | 61.1 |
| [falcosecurity](https://github.com/falcosecurity)  | 705.2k | 33.3k | 14.9k | 11.0k | C Header | 58.8 |
| [kubevirt](https://github.com/kubevirt)  | 683.0k | 122.3k | 38.7k | 8.8k | Go | 56.9 |
| [vitessio](https://github.com/vitessio)  | 678.5k | 42.1k | 16.7k | 20.6k | Go | 56.5 |
| [fluent](https://github.com/fluent)  | 660.2k | 47.7k | 17.8k | 29.7k | C | 55.0 |
| [cubefs](https://github.com/cubefs)  | 657.8k | 7.0k | 3.2k | 6.0k | Go | 54.8 |
| [istio](https://github.com/istio)  | 583.2k | 61.8k | 84.2k | 44.7k | Go | 48.6 |
| [openebs](https://github.com/openebs)  | 559.4k | 21.7k | 12.9k | 12.8k | Go | 46.6 |
| [argoproj](https://github.com/argoproj)  | 508.3k | 29.6k | 31.1k | 51.3k | Go | 42.4 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 503.3k | 18.2k | 15.2k | 22.1k | Go | 41.9 |
| [opengemini](https://github.com/opengemini)  | 500.0k | 1.2k | 1.3k | 1.3k | Go | 41.7 |
| [cncf](https://github.com/cncf)  | 496.7k | 78.5k | 23.6k | 36.4k | JavaScript | 41.4 |
| [pravega](https://github.com/pravega)  | 471.3k | 6.7k | 6.0k | 2.8k | Java | 39.3 |
| [backstage](https://github.com/backstage)  | 470.3k | 81.1k | 30.2k | 33.2k | TypeScript | 39.2 |
| [kubevela](https://github.com/kubevela)  | 448.4k | 9.2k | 8.1k | 8.5k | Go | 37.4 |
| [athenz](https://github.com/athenz)  | 433.0k | 5.0k | 3.6k | 1.0k | Java | 36.1 |
| [knative](https://github.com/knative)  | 432.7k | 40.5k | 40.8k | 15.1k | Go | 36.1 |
| [opencurve](https://github.com/opencurve)  | 418.8k | 6.4k | 2.7k | 2.6k | C++ | 34.9 |
| [werf](https://github.com/werf)  | 410.4k | 20.1k | 8.1k | 6.8k | JavaScript | 34.2 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762 | Java | 33.8 |
| [koordinator-sh](https://github.com/koordinator-sh)  | 399.1k | 2.0k | 2.4k | 1.6k | Go | 33.2 |
| [thanos-io](https://github.com/thanos-io)  | 383.9k | 21.5k | 6.3k | 14.7k | Go | 32.0 |
| [prometheus](https://github.com/prometheus)  | 381.1k | 44.4k | 30.1k | 120.9k | Go | 31.8 |
| [kubeovn](https://github.com/kubeovn)  | 377.0k | 24.1k | 4.8k | 2.2k | Go | 31.4 |
| [pixie-io](https://github.com/pixie-io)  | 376.8k | 13.7k | 2.5k | 6.5k | C++ | 31.4 |
| [dapr](https://github.com/dapr)  | 360.5k | 28.0k | 17.7k | 32.5k | Go | 30.0 |
| [containerd](https://github.com/containerd)  | 355.2k | 35.7k | 19.7k | 37.7k | Go | 29.6 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 347.2k | 19.2k | 15.2k | 20.0k | Go | 28.9 |
| [longhorn](https://github.com/longhorn)  | 345.1k | 19.0k | 17.2k | 8.2k | C | 28.8 |
| [chaosblade-io](https://github.com/chaosblade-io)  | 340.5k | 1.7k | 1.4k | 7.7k | JavaScript | 28.4 |
| [operator-framework](https://github.com/operator-framework)  | 338.1k | 26.6k | 21.7k | 16.3k | Go | 28.2 |
| [radius-project](https://github.com/radius-project)  | 315.1k | 10.1k | 6.8k | 1.8k | C# | 26.3 |
| [kyverno](https://github.com/kyverno)  | 305.0k | 18.6k | 18.1k | 8.5k | Go | 25.4 |
| [open-cluster-management-io](https://github.com/open-cluster-management-io)  | 303.3k | 10.2k | 6.8k | 2.6k | Go | 25.3 |
| [parallaxsecond](https://github.com/parallaxsecond)  | 292.2k | 4.7k | 2.0k | 813 | Rust | 24.4 |
| [confidential-containers](https://github.com/confidential-containers)  | 283.0k | 29.3k | 5.4k | 1.2k | Go | 23.6 |
| [inspektor-gadget](https://github.com/inspektor-gadget)  | 263.3k | 16.2k | 4.4k | 2.6k | C Header | 21.9 |
| [antrea-io](https://github.com/antrea-io)  | 262.3k | 6.7k | 8.4k | 1.8k | Go | 21.9 |
| [keptn](https://github.com/keptn)  | 256.1k | 14.0k | 11.0k | 2.3k | Go | 21.3 |
| [strimzi](https://github.com/strimzi)  | 253.2k | 9.6k | 8.9k | 6.4k | Java | 21.1 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.0k | 12.6k | 13.0k | 29.4k | Go | 20.8 |
| [kusionstack](https://github.com/kusionstack)  | 247.3k | 3.4k | 3.1k | 3.0k | Go | 20.6 |
| [goharbor](https://github.com/goharbor)  | 245.9k | 19.2k | 13.6k | 29.0k | Go | 20.5 |
| [kumahq](https://github.com/kumahq)  | 241.2k | 16.2k | 18.6k | 4.2k | Go | 20.1 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 18.7k | 20.5k | 20.2k | Rust | 19.8 |
| [buildpacks](https://github.com/buildpacks)  | 225.1k | 33.7k | 5.4k | 3.9k | Go | 18.8 |
| [spiffe](https://github.com/spiffe)  | 220.8k | 21.8k | 9.0k | 4.6k | Go | 18.4 |
| [litmuschaos](https://github.com/litmuschaos)  | 211.6k | 8.8k | 9.2k | 5.6k | TypeScript | 17.6 |
| [kcl-lang](https://github.com/kcl-lang)  | 200.3k | 9.0k | 4.3k | 2.6k | Rust | 16.7 |
| [fluxcd](https://github.com/fluxcd)  | 200.1k | 37.0k | 16.1k | 25.4k | Go | 16.7 |
| [wasmcloud](https://github.com/wasmcloud)  | 192.4k | 19.0k | 9.5k | 3.5k | Rust | 16.0 |
| [flatcar](https://github.com/flatcar)  | 192.3k | 54.1k | 6.8k | 1.9k | Go | 16.0 |
| [tremor-rs](https://github.com/tremor-rs)  | 187.6k | 8.1k | 4.3k | 1.0k | Rust | 15.6 |
| [etcd-io](https://github.com/etcd-io)  | 186.5k | 34.2k | 17.3k | 64.2k | Go | 15.6 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 181.9k | 5.5k | 6.0k | 6.1k | Python | 15.2 |
| [carvel-dev](https://github.com/carvel-dev)  | 180.0k | 10.4k | 5.3k | 5.0k | Go | 15.0 |
| [cert-manager](https://github.com/cert-manager)  | 165.2k | 27.9k | 12.4k | 14.9k | Go | 13.8 |
| [loxilb-io](https://github.com/loxilb-io)  | 162.6k | 9.2k | 1.2k | 2.0k | C | 13.6 |
| [kubeslice](https://github.com/kubeslice)  | 157.9k | 7.7k | 1.7k | 510 | Go | 13.2 |
| [meshery](https://github.com/meshery)  | 155.9k | 69.4k | 18.3k | 10.6k | JavaScript | 13.0 |
| [kuadrant](https://github.com/kuadrant)  | 154.5k | 10.9k | 4.5k | 493 | Go | 12.9 |
| [crossplane](https://github.com/crossplane)  | 154.3k | 19.7k | 7.5k | 12.7k | Go | 12.8 |
| [networkservicemesh](https://github.com/networkservicemesh)  | 152.1k | 55.2k | 43.5k | 704 | Go | 12.7 |
| [in-toto](https://github.com/in-toto)  | 151.9k | 13.6k | 4.0k | 2.5k | Python | 12.7 |
| [perses](https://github.com/perses)  | 150.4k | 5.3k | 4.3k | 1.8k | TypeScript | 12.5 |
| [connectrpc](https://github.com/connectrpc)  | 149.0k | 5.1k | 8.2k | 7.2k | Go | 12.4 |
| [kubeedge](https://github.com/kubeedge)  | 148.3k | 9.8k | 5.8k | 8.8k | Go | 12.4 |
| [kube-vip](https://github.com/kube-vip)  | 145.0k | 2.2k | 1.0k | 2.8k | C Header | 12.1 |
| [devfile](https://github.com/devfile)  | 138.2k | 6.7k | 4.3k | 607 | Go | 11.5 |
| [getporter](https://github.com/getporter)  | 136.8k | 8.5k | 3.9k | 1.5k | Go | 11.4 |
| [ovn-org](https://github.com/ovn-org)  | 136.7k | 24.6k | 990 | 890 | C | 11.4 |
| [armadaproject](https://github.com/armadaproject)  | 136.5k | 8.0k | 4.0k | 579 | Go | 11.4 |
| [fluid-cloudnative](https://github.com/fluid-cloudnative)  | 136.3k | 3.7k | 4.0k | 1.8k | Go | 11.4 |
| [kubescape](https://github.com/kubescape)  | 133.7k | 21.9k | 5.6k | 11.5k | Go | 11.1 |
| [wasmedge](https://github.com/wasmedge)  | 133.5k | 6.0k | 3.2k | 10.7k | Rust | 11.1 |
| [rook](https://github.com/rook)  | 131.8k | 28.4k | 10.4k | 13.5k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 7.3k | 5.7k | 4.6k | Go | 10.9 |
| [sermant-io](https://github.com/sermant-io)  | 130.4k | 2.5k | 1.4k | 1.4k | Java | 10.9 |
| [kcp-dev](https://github.com/kcp-dev)  | 128.5k | 8.9k | 3.4k | 2.8k | Go | 10.7 |
| [keylime](https://github.com/keylime)  | 125.0k | 3.5k | 2.5k | 661 | C | 10.4 |
| [nocalhost](https://github.com/nocalhost)  | 122.6k | 6.2k | 2.1k | 2.0k | Go | 10.2 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 6.5k | 5.5k | 5.9k | Go | 10.0 |
| [artifacthub](https://github.com/artifacthub)  | 116.7k | 2.1k | 3.5k | 1.9k | TypeScript | 9.7 |
| [openelb](https://github.com/openelb)  | 116.3k | 695 | 397 | 1.8k | JavaScript | 9.7 |
| [pipe-cd](https://github.com/pipe-cd)  | 115.1k | 7.3k | 5.3k | 1.4k | Go | 9.6 |
| [serverless-devs](https://github.com/serverless-devs)  | 114.0k | 5.3k | 772 | 1.9k | JavaScript | 9.5 |
| [kptdev](https://github.com/kptdev)  | 113.6k | 4.3k | 4.6k | 2.0k | Go | 9.5 |
| [paralus](https://github.com/paralus)  | 112.6k | 2.1k | 1.3k | 1.2k | Go | 9.4 |
| [notaryproject](https://github.com/notaryproject)  | 107.7k | 6.8k | 5.6k | 4.4k | Go | 9.0 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 107.1k | 3.7k | 3.7k | 7.8k | Go | 8.9 |
| [kedacore](https://github.com/kedacore)  | 105.3k | 5.9k | 7.2k | 10.8k | Go | 8.8 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 103.5k | 19.6k | 4.1k | 4.5k | Python | 8.6 |
| [karmada-io](https://github.com/karmada-io)  | 100.5k | 8.9k | 5.9k | 5.3k | Go | 8.4 |
| [kubeclipper](https://github.com/kubeclipper)  | 99.5k | 686 | 1.0k | 510 | Go | 8.3 |
| [openyurtio](https://github.com/openyurtio)  | 98.6k | 2.9k | 2.6k | 2.1k | Go | 8.2 |
| [bfenetworks](https://github.com/bfenetworks)  | 98.2k | 1.5k | 875 | 6.4k | Go | 8.2 |
| [v6d-io](https://github.com/v6d-io)  | 92.2k | 1.5k | 1.4k | 928 | C++ | 7.7 |
| [spidernet-io](https://github.com/spidernet-io)  | 92.1k | 7.3k | 6.7k | 912 | Go | 7.7 |
| [serverlessworkflow](https://github.com/serverlessworkflow)  | 90.4k | 4.5k | 2.1k | 1.5k | C# | 7.5 |
| [kairos-io](https://github.com/kairos-io)  | 88.2k | 14.2k | 9.7k | 1.7k | Go | 7.3 |
| [opencost](https://github.com/opencost)  | 82.3k | 6.7k | 3.5k | 6.3k | Go | 6.9 |
| [openservicemesh](https://github.com/openservicemesh)  | 81.9k | 5.2k | 4.1k | 2.6k | Go | 6.8 |
| [hwameistor](https://github.com/hwameistor)  | 80.4k | 3.4k | 1.5k | 627 | Go | 6.7 |
| [easegress-io](https://github.com/easegress-io)  | 79.9k | 2.2k | 1.3k | 5.9k | Go | 6.7 |
| [cloudevents](https://github.com/cloudevents)  | 79.0k | 4.8k | 3.3k | 8.3k | Go | 6.6 |
| [oras-project](https://github.com/oras-project)  | 76.9k | 5.7k | 3.3k | 2.5k | Go | 6.4 |
| [coredns](https://github.com/coredns)  | 76.2k | 6.6k | 6.1k | 14.9k | Go | 6.3 |
| [project-zot](https://github.com/project-zot)  | 75.8k | 1.9k | 3.1k | 1.5k | Go | 6.3 |
| [devspace-sh](https://github.com/devspace-sh)  | 74.9k | 6.1k | 1.9k | 4.8k | Go | 6.2 |
| [volcano-sh](https://github.com/volcano-sh)  | 74.6k | 7.3k | 3.5k | 5.5k | Go | 6.2 |
| [tinkerbell](https://github.com/tinkerbell)  | 71.4k | 18.0k | 5.1k | 2.7k | Go | 6.0 |
| [helm](https://github.com/helm)  | 71.3k | 32.2k | 29.5k | 56.4k | Go | 5.9 |
| [cri-o](https://github.com/cri-o)  | 71.1k | 12.8k | 8.5k | 5.7k | Go | 5.9 |
| [oscal-compass](https://github.com/oscal-compass)  | 71.0k | 2.4k | 1.6k | 366 | Python | 5.9 |
| [metal3-io](https://github.com/metal3-io)  | 70.9k | 17.4k | 10.5k | 1.6k | Go | 5.9 |
| [submariner-io](https://github.com/submariner-io)  | 69.8k | 13.9k | 16.6k | 3.0k | Go | 5.8 |
| [runatlantis](https://github.com/runatlantis)  | 68.5k | 4.9k | 20.6k | 8.7k | Go | 5.7 |
| [open-feature](https://github.com/open-feature)  | 67.6k | 15.4k | 14.2k | 3.5k | Go | 5.6 |
| [microcks](https://github.com/microcks)  | 66.0k | 6.9k | 2.8k | 2.1k | Java | 5.5 |
| [openfga](https://github.com/openfga)  | 65.8k | 9.3k | 7.2k | 5.0k | Go | 5.5 |
| [kanisterio](https://github.com/kanisterio)  | 63.6k | 3.3k | 3.3k | 847 | Go | 5.3 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 2.0k | 2.6k | 3.0k | Go | 5.2 |
| [telepresenceio](https://github.com/telepresenceio)  | 61.1k | 10.8k | 2.4k | 7.0k | Go | 5.1 |
| [external-secrets](https://github.com/external-secrets)  | 57.3k | 4.5k | 3.9k | 8.6k | Go | 4.8 |
| [cartography-cncf](https://github.com/cartography-cncf)  | 57.2k | 1.2k | 1.4k | 3.6k | Python | 4.8 |
| [project-hami](https://github.com/project-hami)  | 53.2k | 1.9k | 1.1k | 2.6k | Go | 4.4 |
| [rkt](https://github.com/rkt)  | 52.4k | 5.6k | 2.5k | 8.8k | Go | 4.4 |
| [trickstercache](https://github.com/trickstercache)  | 52.1k | 1.3k | 640 | 2.1k | Go | 4.3 |
| [headlamp-k8s](https://github.com/headlamp-k8s)  | 51.4k | 6.0k | 2.2k | 2.7k | TypeScript | 4.3 |
| [kaito-project](https://github.com/kaito-project)  | 51.0k | 6.1k | 2.1k | 1.3k | Go | 4.2 |
| [kubestellar](https://github.com/kubestellar)  | 50.6k | 9.6k | 4.1k | 823 | Go | 4.2 |
| [slimtoolkit](https://github.com/slimtoolkit)  | 49.7k | 1.3k | 451 | 22.5k | Go | 4.1 |
| [shipwright-io](https://github.com/shipwright-io)  | 48.8k | 4.3k | 2.4k | 835 | Go | 4.1 |
| [theupdateframework](https://github.com/theupdateframework)  | 48.6k | 11.8k | 4.8k | 3.1k | Python | 4.0 |
| [superedge](https://github.com/superedge)  | 48.2k | 524 | 461 | 1.1k | Go | 4.0 |
| [virtual-kubelet](https://github.com/virtual-kubelet)  | 48.1k | 3.0k | 1.7k | 5.2k | Go | 4.0 |
| [kudobuilder](https://github.com/kudobuilder)  | 47.9k | 2.3k | 2.2k | 2.2k | Go | 4.0 |
| [xline-kv](https://github.com/xline-kv)  | 47.6k | 1.6k | 1.0k | 713 | Rust | 4.0 |
| [distribution](https://github.com/distribution)  | 45.7k | 6.1k | 2.4k | 10.3k | Go | 3.8 |
| [k3s-io](https://github.com/k3s-io)  | 44.5k | 5.9k | 6.5k | 37.0k | Go | 3.7 |
| [openfunction](https://github.com/openfunction)  | 44.4k | 1.7k | 1.2k | 1.7k | Go | 3.7 |
| [skooner-k8s](https://github.com/skooner-k8s)  | 44.0k | 543 | 259 | 1.4k | JavaScript | 3.7 |
| [kubedl-io](https://github.com/kubedl-io)  | 44.0k | 1.4k | 409 | 618 | Go | 3.7 |
| [krkn-chaos](https://github.com/krkn-chaos)  | 43.2k | 2.0k | 1.6k | 507 | Python | 3.6 |
| [kubearmor](https://github.com/kubearmor)  | 41.6k | 7.4k | 3.3k | 2.2k | Go | 3.5 |
| [metallb](https://github.com/metallb)  | 40.9k | 3.5k | 2.4k | 8.0k | Go | 3.4 |
| [vscode-kubernetes-tools](https://github.com/vscode-kubernetes-tools)  | 39.4k | 956 | 983 | 752 | TypeScript | 3.3 |
| [devstream-io](https://github.com/devstream-io)  | 39.3k | 663 | 1.1k | 887 | Go | 3.3 |
| [aeraki-mesh](https://github.com/aeraki-mesh)  | 39.1k | 1.1k | 503 | 1.0k | Go | 3.2 |
| [youki-dev](https://github.com/youki-dev)  | 37.9k | 6.1k | 3.1k | 7.2k | Rust | 3.2 |
| [kube-logging](https://github.com/kube-logging)  | 35.9k | 4.7k | 2.3k | 1.8k | Go | 3.0 |
| [kuasar-io](https://github.com/kuasar-io)  | 35.7k | 379 | 186 | 1.3k | Rust | 3.0 |
| [bpfman](https://github.com/bpfman)  | 35.5k | 3.4k | 1.6k | 711 | Rust | 3.0 |
| [score-spec](https://github.com/score-spec)  | 35.5k | 2.6k | 1.3k | 8.9k | Go | 3.0 |
| [sustainable-computing-io](https://github.com/sustainable-computing-io)  | 35.1k | 5.6k | 3.5k | 1.6k | Go | 2.9 |
| [kmesh-net](https://github.com/kmesh-net)  | 35.1k | 3.7k | 1.3k | 708 | Go | 2.9 |
| [ratify-project](https://github.com/ratify-project)  | 35.0k | 1.5k | 1.7k | 255 | Go | 2.9 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 3.4k | 1.6k | 8.4k | Go | 2.9 |
| [clusternet](https://github.com/clusternet)  | 33.4k | 1.2k | 801 | 1.4k | Go | 2.8 |
| [sealerio](https://github.com/sealerio)  | 32.6k | 1.2k | 1.4k | 2.1k | Go | 2.7 |
| [kube-rs](https://github.com/kube-rs)  | 31.9k | 4.8k | 2.0k | 4.0k | Rust | 2.7 |
| [bank-vaults](https://github.com/bank-vaults)  | 31.7k | 13.6k | 7.9k | 2.5k | Go | 2.6 |
| [clusterpedia-io](https://github.com/clusterpedia-io)  | 31.6k | 1.6k | 871 | 886 | Go | 2.6 |
| [dexidp](https://github.com/dexidp)  | 29.3k | 3.9k | 3.1k | 10.3k | Go | 2.4 |
| [piraeusdatastore](https://github.com/piraeusdatastore)  | 28.9k | 2.4k | 1.2k | 1.2k | Go | 2.4 |
| [opentracing](https://github.com/opentracing)  | 27.2k | 2.7k | 1.5k | 10.2k | Java | 2.3 |
| [krustlet](https://github.com/krustlet)  | 23.6k | 2.4k | 535 | 3.6k | Rust | 2.0 |
| [project-stacker](https://github.com/project-stacker)  | 22.4k | 1.5k | 721 | 327 | Go | 1.9 |
| [curiefense](https://github.com/curiefense)  | 22.3k | 2.9k | 832 | 617 | Rust | 1.9 |
| [projectcapsule](https://github.com/projectcapsule)  | 21.9k | 2.3k | 2.0k | 2.0k | Go | 1.8 |
| [fabedge](https://github.com/fabedge)  | 21.6k | 1.1k | 529 | 535 | Go | 1.8 |
| [project-akri](https://github.com/project-akri)  | 20.5k | 1.0k | 604 | 1.2k | Rust | 1.7 |
| [k8gb-io](https://github.com/k8gb-io)  | 19.4k | 2.0k | 1.8k | 1.1k | Go | 1.6 |
| [foniod](https://github.com/foniod)  | 18.1k | 1.7k | 495 | 2.1k | Rust | 1.5 |
| [lima-vm](https://github.com/lima-vm)  | 17.4k | 5.9k | 2.9k | 19.1k | Go | 1.4 |
| [schemahero](https://github.com/schemahero)  | 17.0k | 1.7k | 1.1k | 1.1k | Go | 1.4 |
| [hexa-org](https://github.com/hexa-org)  | 16.5k | 956 | 181 | 125 | Go | 1.4 |
| [getsops](https://github.com/getsops)  | 15.3k | 2.4k | 1.1k | 19.7k | Go | 1.3 |
| [k8up-io](https://github.com/k8up-io)  | 14.1k | 2.4k | 864 | 874 | Go | 1.2 |
| [kubean-io](https://github.com/kubean-io)  | 13.2k | 2.8k | 1.3k | 552 | Go | 1.1 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 824 | 150 | 712 | Go | 1.1 |
| [kuberhealthy](https://github.com/kuberhealthy)  | 11.8k | 3.3k | 1.0k | 2.2k | Go | 990.0m |
| [servicemeshinterface](https://github.com/servicemeshinterface)  | 11.6k | 974 | 375 | 1.2k | Go | 970.0m |
| [k8sgpt-ai](https://github.com/k8sgpt-ai)  | 11.0k | 2.6k | 2.0k | 7.5k | Go | 920.0m |
| [ko-build](https://github.com/ko-build)  | 9.5k | 1.8k | 1.4k | 8.2k | Go | 790.0m |
| [eraser-dev](https://github.com/eraser-dev)  | 9.4k | 724 | 935 | 568 | Go | 780.0m |
| [kube-burner](https://github.com/kube-burner)  | 8.0k | 1.7k | 1.0k | 675 | Go | 660.0m |
| [tellerops](https://github.com/tellerops)  | 7.7k | 246 | 167 | 3.2k | Go | 640.0m |
| [cni-genie](https://github.com/cni-genie)  | 5.4k | 490 | 136 | 538 | Go | 450.0m |
| [project-copacetic](https://github.com/project-copacetic)  | 5.4k | 1.0k | 1.2k | 1.5k | Go | 450.0m |
| [merbridge](https://github.com/merbridge)  | 4.8k | 296 | 386 | 806 | Go | 400.0m |
| [opcr-io](https://github.com/opcr-io)  | 3.6k | 734 | 257 | 279 | Go | 300.0m |
| [krator-rs](https://github.com/krator-rs)  | 2.9k | 498 | 34 | 167 | Rust | 250.0m |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251 | 117 | 2.0k | Go | 190.0m |
| [openobservability](https://github.com/openobservability)  | 2.3k | 251 | 117 | 2.0k | Go | 190.0m |
| [service-mesh-performance](https://github.com/service-mesh-performance)  | 2.2k | 906 | 261 | 305 | JavaScript | 180.0m |
| [kubereboot](https://github.com/kubereboot)  | 2.0k | 1.9k | 1.1k | 2.4k | Go | 170.0m |


### OTEL Instrumentations
| Project | SLOC | Commits | PRs | Stars | Effort |
|---------|------|---------|-----|-------|--------|
| [**OTEL**](https://opentelemetry.io/docs/instrumentation/)  | 1.5M | 64.8k | 67.3k | 35.7k | 126.6 |
| [ebpf-instrumentation](https://github.com/open-telemetry/opentelemetry-ebpf-instrumentation/)  | 367.3k | 592 | 654 | 174 | 30.61 |
| [java-instrumentation](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.8k | 12.6k | 9.7k | 2.3k | 13.99 |
| [dotnet-instrumentation](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.1k | 2.4k | 3.7k | 436 | 11.76 |
| [java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 4.8k | 5.4k | 2.3k | 8.03 |
| [js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 2.8k | 3.4k | 3.1k | 5.08 |
| [dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.8k | 3.5k | 4.2k | 3.6k | 4.73 |
| [js-contrib](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 2.5k | 2.2k | 843 | 4.28 |
| [go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 4.0k | 5.2k | 6.1k | 4.26 |
| [python-contrib](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 42.9k | 2.6k | 2.3k | 908 | 3.57 |
| [network](https://github.com/open-telemetry/opentelemetry-network/)  | 42.6k | 510 | 308 | 385 | 3.55 |
| [ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4k | 268 | 112 | 85 | 3.45 |
| [python](https://github.com/open-telemetry/opentelemetry-python/)  | 38.5k | 2.0k | 2.5k | 2.2k | 3.21 |
| [cpp-contrib](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 37.4k | 333 | 390 | 144 | 3.12 |
| [cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.6k | 1.7k | 1.9k | 1.1k | 2.8 |
| [dotnet-contrib](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.4k | 2.0k | 2.3k | 595 | 2.78 |
| [ebpf-profiler](https://github.com/open-telemetry/opentelemetry-ebpf-profiler/)  | 33.1k | 507 | 611 | 2.9k | 2.76 |
| [php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 820 | 975 | 865 | 2.58 |
| [rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 1.8k | 2.1k | 2.3k | 2.54 |
| [go-contrib](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.3k | 3.3k | 7.1k | 1.5k | 2.27 |
| [swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9k | 1.0k | 584 | 289 | 2.16 |
| [swift-core](https://github.com/open-telemetry/opentelemetry-swift-core/)  | 20.4k | 1.0k | 16 | 2 | 1.7 |
| [ruby-contrib](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 20.0k | 1.6k | 1.4k | 113 | 1.66 |
| [ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16.4k | 993 | 1.2k | 544 | 1.37 |
| [php-contrib](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12.8k | 594 | 430 | 92 | 1.06 |
| [erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11.5k | 1.9k | 619 | 371 | 0.96 |
| [rust-contrib](https://github.com/open-telemetry/opentelemetry-rust-contrib/)  | 7.4k | 1.1k | 370 | 66 | 0.62 |
| [android](https://github.com/open-telemetry/opentelemetry-android/)  | 7.1k | 1.8k | 1.0k | 238 | 0.6 |
| [go-compile-instrumentation](https://github.com/open-telemetry/opentelemetry-go-compile-instrumentation/)  | 3.8k | 55 | 62 | 149 | 0.32 |
| [go-instrumentation](https://github.com/open-telemetry/opentelemetry-go-instrumentation/)  | 2.9k | 2.3k | 2.6k | 919 | 0.24 |
| [java-contrib](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 2.8k | 1.8k | 2.0k | 240 | 0.24 |
| [erlang-contrib](https://github.com/open-telemetry/opentelemetry-erlang-contrib/)  | 1.6k | 371 | 455 | 192 | 0.14 |
| [java-examples](https://github.com/open-telemetry/opentelemetry-java-examples/)  | 1.3k | 762 | 796 | 319 | 0.11 |
| [php-instrumentation](https://github.com/open-telemetry/opentelemetry-php-instrumentation/)  | 556 | 161 | 170 | 123 | 0.05 |
| [injector](https://github.com/open-telemetry/opentelemetry-injector/)  | 143 | 60 | 71 | 60 | 0.01 |
| [browser](https://github.com/open-telemetry/opentelemetry-browser/)  | 0 | 2 | 4 | 5 | 0 |
| [configuration](https://github.com/open-telemetry/opentelemetry-configuration/)  | 0 | 199 | 217 | 66 | 0 |
| [for-beginners](https://github.com/open-telemetry/opentelemetry-for-beginners/)  | 0 | 1 | 0 | 0 | 0 |
| [swift-grpc](https://github.com/open-telemetry/opentelemetry-swift-grpc/)  | 0 | 4 | 0 | 0 | 0 |


### go-faster

| Project | SLOC | Commits | PRs | Stars | Effort |
|---------|------|---------|-----|-------|--------|
| [**go-faster**](https://github.com/go-faster)  | 196.9k | 29.8k | 10.2k | 5.8k | 16.41 |
| [td](https://github.com/gotd/td/)  | 52.6k | 4.1k | 1.3k | 2.0k | 4.38 |
| [ogen](https://github.com/ogen-go/ogen/)  | 36.7k | 4.6k | 1.2k | 1.9k | 3.05 |
| [yaml](https://github.com/go-faster/yaml/)  | 15.6k | 613 | 64 | 10 | 1.3 |
| [ch-go](https://github.com/ClickHouse/ch-go/)  | 15.0k | 1.5k | 435 | 386 | 1.25 |
| [hx](https://github.com/go-faster/hx/)  | 13.0k | 1.5k | 20 | 3 | 1.09 |
| [jx](https://github.com/go-faster/jx/)  | 9.7k | 1.5k | 88 | 236 | 0.81 |
| [stun](https://github.com/pion/stun/)  | 8.7k | 856 | 219 | 769 | 0.72 |
| [botapi](https://github.com/gotd/botapi/)  | 6.4k | 782 | 439 | 26 | 0.53 |
| [contrib](https://github.com/gotd/contrib/)  | 4.7k | 1.2k | 707 | 22 | 0.39 |
| [oteldb](https://github.com/go-faster/oteldb/)  | 4.5k | 3.3k | 735 | 58 | 0.38 |
| [bot](https://github.com/gotd/bot/)  | 3.1k | 732 | 322 | 15 | 0.26 |
| [city](https://github.com/go-faster/city/)  | 2.9k | 60 | 3 | 39 | 0.24 |
| [gha](https://github.com/go-faster/gha/)  | 2.8k | 319 | 213 | 0 | 0.24 |
| [bot](https://github.com/go-faster/bot/)  | 2.7k | 1.4k | 585 | 22 | 0.22 |
| [ref](https://github.com/gotd/ref/)  | 2.4k | 242 | 156 | 0 | 0.2 |
| [tl](https://github.com/gotd/tl/)  | 1.5k | 153 | 57 | 28 | 0.12 |
| [errors](https://github.com/go-faster/errors/)  | 1.4k | 122 | 26 | 62 | 0.12 |
| [sdk](https://github.com/go-faster/sdk/)  | 1.3k | 446 | 224 | 16 | 0.11 |
| [protoc-gen-oas](https://github.com/ogen-go/protoc-gen-oas/)  | 1.1k | 336 | 200 | 25 | 0.09 |
| [getdoc](https://github.com/gotd/getdoc/)  | 1.1k | 1.2k | 575 | 8 | 0.09 |
| [gh-archive-clickhouse](https://github.com/go-faster/gh-archive-clickhouse/)  | 983 | 397 | 241 | 9 | 0.08 |
| [tail](https://github.com/go-faster/tail/)  | 943 | 119 | 37 | 23 | 0.08 |
| [ent2ogen](https://github.com/ogen-go/ent2ogen/)  | 924 | 174 | 131 | 10 | 0.08 |
| [neo](https://github.com/gotd/neo/)  | 795 | 103 | 38 | 10 | 0.07 |
| [simon](https://github.com/go-faster/simon/)  | 625 | 414 | 171 | 5 | 0.05 |
| [teled](https://github.com/gotd/teled/)  | 623 | 175 | 130 | 15 | 0.05 |
| [ch-bench](https://github.com/ClickHouse/ch-bench/)  | 620 | 92 | 21 | 33 | 0.05 |
| [cli](https://github.com/gotd/cli/)  | 593 | 332 | 223 | 22 | 0.05 |
| [tgstatus](https://github.com/gotd/tgstatus/)  | 408 | 253 | 121 | 13 | 0.03 |
| [flightorder](https://github.com/go-faster/flightorder/)  | 388 | 79 | 37 | 3 | 0.03 |
| [example](https://github.com/ogen-go/example/)  | 358 | 258 | 217 | 32 | 0.03 |
| [web](https://github.com/go-faster/web/)  | 286 | 402 | 179 | 0 | 0.02 |
| [cilium-mesh-testing](https://github.com/go-faster/cilium-mesh-testing/)  | 253 | 3 | 2 | 1 | 0.02 |
| [stat](https://github.com/go-faster/stat/)  | 242 | 369 | 185 | 1 | 0.02 |
| [prio](https://github.com/go-faster/prio/)  | 240 | 79 | 57 | 0 | 0.02 |
| [web](https://github.com/ogen-go/web/)  | 200 | 445 | 268 | 3 | 0.02 |
| [docs](https://github.com/gotd/docs/)  | 193 | 209 | 211 | 0 | 0.02 |
| [gh-archive-yt](https://github.com/go-faster/gh-archive-yt/)  | 192 | 134 | 65 | 5 | 0.02 |
| [testbot](https://github.com/gotd/testbot/)  | 189 | 34 | 15 | 0 | 0.02 |
| [vega](https://github.com/go-faster/vega/)  | 185 | 466 | 136 | 1 | 0.02 |
| [xor](https://github.com/go-faster/xor/)  | 178 | 212 | 108 | 16 | 0.01 |
| [ytst](https://github.com/go-faster/ytst/)  | 120 | 131 | 4 | 0 | 0.01 |
| [schedpolicy](https://github.com/go-faster/schedpolicy/)  | 88 | 4 | 0 | 1 | 0.01 |
| [bfj](https://github.com/go-faster/bfj/)  | 82 | 7 | 0 | 0 | 0.01 |
| [jx-bench-sample](https://github.com/go-faster/jx-bench-sample/)  | 59 | 8 | 0 | 0 | 0 |
| [ki](https://github.com/go-faster/ki/)  | 55 | 3 | 0 | 0 | 0 |
| [gotdlog](https://github.com/gotd/gotdlog/)  | 26 | 11 | 3 | 0 | 0 |
| [charts](https://github.com/go-faster/charts/)  | 3 | 25 | 3 | 0 | 0 |
| [hostbench](https://github.com/go-faster/hostbench/)  | 1 | 12 | 0 | 0 | 0 |
| [otlab](https://github.com/go-faster/otlab/)  | 1 | 24 | 7 | 0 | 0 |
| [template](https://github.com/go-faster/template/)  | 1 | 46 | 4 | 0 | 0 |
