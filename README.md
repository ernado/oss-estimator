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
| [**CNCF**](https://www.cncf.io/)  | 48.8M | 2.3M | 1.2M | 1.7M | Go | 4.1k |
| [torvalds](https://github.com/torvalds)  | 19.1M | 1.3M | 792.0 | 168.9k | C | 1.6k |
| [apache](https://github.com/apache)  | 15.7M | 610.1k | 45.2k | 88.9k | Java | 1.3k |
| [ydb-platform](https://github.com/ydb-platform)  | 7.1M | 36.9k | 2.0k | 3.2k | C++ | 589.2 |
| [**K8s**](https://kubernetes.io/)  | 5.5M | 502.6k | 262.8k | 341.9k | Go | 454.7 |
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1k | 190.5k | 243.1k | Go | 287.0 |
| [tensorflow](https://github.com/tensorflow)  | 3.4M | 162.5k | 25.5k | 182.0k | C++ | 279.2 |
| [ytsaurus](https://github.com/ytsaurus)  | 3.3M | 68.3k | 354.0 | 880.0 | C++ | 274.9 |
| [elastic](https://github.com/elastic)  | 3.1M | 104.3k | 109.9k | 93.2k | Java | 261.0 |
| [envoyproxy](https://github.com/envoyproxy)  | 2.9M | 52.5k | 25.1k | 28.1k | JavaScript | 243.5 |
| [godotengine](https://github.com/godotengine)  | 2.1M | 78.5k | 48.0k | 99.9k | C++ | 171.0 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.0M | 121.4k | 72.3k | 98.8k | Go | 167.7 |
| [open-telemetry](https://github.com/open-telemetry)  | 2.0M | 69.7k | 61.8k | 31.5k | Go | 165.4 |
| [cockroachdb](https://github.com/cockroachdb)  | 2.0M | 100.8k | 59.3k | 33.7k | Go | 163.2 |
| [grpc](https://github.com/grpc)  | 1.7M | 75.0k | 38.9k | 86.8k | C++ | 143.8 |
| [ClickHouse](https://github.com/ClickHouse)  | 1.7M | 120.5k | 35.2k | 34.5k | C++ | 142.1 |
| [grafana](https://github.com/grafana)  | 1.7M | 80.4k | 51.2k | 105.8k | Go | 142.0 |
| [konveyor](https://github.com/konveyor)  | 1.6M | 8.1k | 7.7k | 858.0 | JavaScript | 136.6 |
| [inclavare-containers](https://github.com/inclavare-containers)  | 1.6M | 25.1k | 1.1k | 571.0 | C++ | 130.7 |
| [rust-lang](https://github.com/rust-lang)  | 1.5M | 251.7k | 71.0k | 92.3k | Rust | 126.5 |
| [python](https://github.com/python)  | 1.4M | 121.5k | 49.8k | 59.3k | Python | 114.4 |
| [openkruise](https://github.com/openkruise)  | 1.4M | 1.3k | 1.3k | 4.0k | JavaScript | 113.7 |
| [docker](https://github.com/docker)  | 1.3M | 177.2k | 30.8k | 146.8k | Go | 107.4 |
| [netdata](https://github.com/netdata)  | 1.2M | 25.8k | 12.1k | 63.1k | JavaScript | 103.7 |
| [prestodb](https://github.com/prestodb)  | 1.2M | 26.0k | 15.3k | 15.7k | Java | 97.0 |
| [kubewarden](https://github.com/kubewarden)  | 1.1M | 8.8k | 3.1k | 494.0 | JavaScript | 92.6 |
| [golang](https://github.com/golang)  | 1.1M | 59.7k | 3.9k | 119.1k | Go | 88.4 |
| [nodejs](https://github.com/nodejs)  | 966.7k | 41.0k | 33.0k | 103.3k | JavaScript | 80.6 |
| [kubeflow](https://github.com/kubeflow)  | 938.8k | 20.8k | 24.2k | 27.2k | Go | 78.2 |
| [ziglang](https://github.com/ziglang)  | 888.6k | 28.9k | 9.2k | 30.0k | Zig | 74.0 |
| [brigadecore](https://github.com/brigadecore)  | 877.2k | 6.3k | 2.8k | 3.0k | JavaScript | 73.1 |
| [tikv](https://github.com/tikv)  | 862.8k | 17.2k | 21.3k | 25.2k | Rust | 71.9 |
| [nats-io](https://github.com/nats-io)  | 862.6k | 37.8k | 13.8k | 33.3k | Go | 71.9 |
| [keycloak](https://github.com/keycloak)  | 842.3k | 34.8k | 22.5k | 25.4k | Java | 70.2 |
| [cdk8s-team](https://github.com/cdk8s-team)  | 804.0k | 9.3k | 7.9k | 3.8k | Go | 67.0 |
| [m3db](https://github.com/m3db)  | 736.8k | 10.4k | 5.1k | 4.6k | Go | 61.4 |
| [cilium](https://github.com/cilium)  | 734.5k | 34.6k | 23.4k | 26.3k | Go | 61.2 |
| [kubevirt](https://github.com/kubevirt)  | 683.4k | 61.0k | 22.7k | 5.5k | Go | 57.0 |
| [vitessio](https://github.com/vitessio)  | 677.6k | 37.2k | 11.1k | 15.9k | Go | 56.5 |
| [cubefs](https://github.com/cubefs)  | 659.3k | 3.1k | 1.5k | 3.6k | Go | 54.9 |
| [fluent](https://github.com/fluent)  | 655.2k | 38.1k | 12.0k | 24.2k | C | 54.6 |
| [moby](https://github.com/moby)  | 652.1k | 75.1k | 31.5k | 83.2k | Go | 54.3 |
| [falcosecurity](https://github.com/falcosecurity)  | 625.9k | 22.9k | 7.9k | 8.6k | C Header | 52.2 |
| [istio](https://github.com/istio)  | 583.6k | 47.1k | 62.4k | 38.5k | Go | 48.6 |
| [openebs](https://github.com/openebs)  | 560.6k | 16.1k | 9.7k | 10.1k | Go | 46.7 |
| [argoproj](https://github.com/argoproj)  | 517.8k | 18.1k | 16.9k | 34.2k | Go | 43.1 |
| [denoland](https://github.com/denoland)  | 498.2k | 11.1k | 12.5k | 92.8k | JavaScript | 41.5 |
| [cncf](https://github.com/cncf)  | 497.9k | 68.0k | 16.3k | 28.6k | JavaScript | 41.5 |
| [googlecontainertools](https://github.com/googlecontainertools)  | 493.2k | 21.6k | 34.6k | 66.7k | Go | 41.1 |
| [pravega](https://github.com/pravega)  | 474.4k | 6.4k | 5.5k | 2.5k | Java | 39.5 |
| [metrico](https://github.com/metrico)  | 467.2k | 19.8k | 546.0 | 884.0 | Rust | 38.9 |
| [kubevela](https://github.com/kubevela)  | 448.4k | 8.3k | 6.8k | 5.6k | Go | 37.4 |
| [athenz](https://github.com/athenz)  | 433.0k | 3.8k | 2.2k | 808.0 | Java | 36.1 |
| [knative](https://github.com/knative)  | 432.8k | 30.6k | 32.3k | 12.4k | Go | 36.1 |
| [go-faster](https://github.com/go-faster)  | 432.2k | 30.7k | 906.0 | 169.0 | Go | 36.0 |
| [awslabs](https://github.com/awslabs)  | 427.5k | 11.1k | 4.5k | 45.6k | Java | 35.6 |
| [opencurve](https://github.com/opencurve)  | 419.9k | 5.7k | 1.9k | 1.9k | C++ | 35.0 |
| [backstage](https://github.com/backstage)  | 409.3k | 40.6k | 13.0k | 21.8k | TypeScript | 34.1 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762.0 | Java | 33.8 |
| [prometheus](https://github.com/prometheus)  | 379.9k | 31.2k | 19.3k | 91.9k | Go | 31.7 |
| [pixie-io](https://github.com/pixie-io)  | 378.2k | 13.4k | 2.1k | 5.5k | C++ | 31.5 |
| [kubeovn](https://github.com/kubeovn)  | 376.4k | 21.7k | 2.0k | 1.5k | Go | 31.4 |
| [django](https://github.com/django)  | 365.8k | 32.6k | 18.0k | 76.5k | Python | 30.5 |
| [werf](https://github.com/werf)  | 358.3k | 16.4k | 5.8k | 4.8k | JavaScript | 29.9 |
| [containerd](https://github.com/containerd)  | 358.1k | 23.6k | 11.7k | 24.8k | Go | 29.8 |
| [facebook](https://github.com/facebook)  | 357.5k | 16.6k | 14.8k | 221.0k | JavaScript | 29.8 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.0k | 8.5k | 7.2k | 15.4k | Go | 28.6 |
| [dapr](https://github.com/dapr)  | 340.1k | 18.1k | 11.0k | 26.1k | Go | 28.4 |
| [operator-framework](https://github.com/operator-framework)  | 326.8k | 21.6k | 15.5k | 14.0k | Go | 27.2 |
| [chaosblade-io](https://github.com/chaosblade-io)  | 310.9k | 1.3k | 979.0 | 6.2k | JavaScript | 25.9 |
| [LINBIT](https://github.com/LINBIT)  | 310.9k | 4.9k | 14.0 | 812.0 | Java | 25.9 |
| [vectordotdev](https://github.com/vectordotdev)  | 309.0k | 14.4k | 10.2k | 14.4k | Rust | 25.8 |
| [kyverno](https://github.com/kyverno)  | 300.4k | 9.4k | 5.6k | 4.0k | Go | 25.0 |
| [parallaxsecond](https://github.com/parallaxsecond)  | 292.2k | 3.7k | 1.5k | 553.0 | Rust | 24.4 |
| [open-cluster-management-io](https://github.com/open-cluster-management-io)  | 289.8k | 7.0k | 3.5k | 1.9k | Go | 24.1 |
| [confidential-containers](https://github.com/confidential-containers)  | 282.4k | 23.7k | 1.8k | 425.0 | Go | 23.5 |
| [antrea-io](https://github.com/antrea-io)  | 262.8k | 3.5k | 3.9k | 1.5k | Go | 21.9 |
| [keptn](https://github.com/keptn)  | 256.8k | 11.8k | 8.4k | 1.9k | Go | 21.4 |
| [uber](https://github.com/uber)  | 253.8k | 1.6k | 254.0 | 6.5k | Go | 21.1 |
| [inspektor-gadget](https://github.com/inspektor-gadget)  | 253.6k | 3.7k | 1.0k | 1.3k | C Header | 21.1 |
| [siderolabs](https://github.com/siderolabs)  | 253.4k | 7.4k | 9.2k | 4.9k | Go | 21.1 |
| [strimzi](https://github.com/strimzi)  | 253.2k | 7.1k | 6.0k | 4.3k | Java | 21.1 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.5k | 8.7k | 7.6k | 23.7k | Go | 20.9 |
| [gravity-ui](https://github.com/gravity-ui)  | 249.4k | 4.1k | 2.2k | 576.0 | TypeScript | 20.8 |
| [goharbor](https://github.com/goharbor)  | 246.8k | 17.0k | 10.4k | 21.1k | Go | 20.6 |
| [kumahq](https://github.com/kumahq)  | 240.3k | 7.2k | 7.1k | 3.3k | Go | 20.0 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 13.0k | 13.4k | 18.0k | Rust | 19.8 |
| [matplotlib](https://github.com/matplotlib)  | 235.9k | 50.1k | 17.9k | 19.1k | Python | 19.6 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 233.9k | 8.6k | 2.7k | 11.2k | Go | 19.5 |
| [kata-containers](https://github.com/kata-containers)  | 225.2k | 22.8k | 10.3k | 7.0k | Go | 18.8 |
| [buildpacks](https://github.com/buildpacks)  | 225.2k | 22.5k | 3.6k | 2.8k | Go | 18.8 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 224.7k | 11.1k | 6.7k | 9.8k | Rust | 18.7 |
| [spiffe](https://github.com/spiffe)  | 216.5k | 9.4k | 4.1k | 3.0k | Go | 18.0 |
| [Netflix](https://github.com/Netflix)  | 208.1k | 4.2k | 2.3k | 549.0 | Java | 17.4 |
| [litmuschaos](https://github.com/litmuschaos)  | 205.9k | 7.5k | 7.1k | 4.1k | TypeScript | 17.1 |
| [alibaba](https://github.com/alibaba)  | 205.2k | 4.1k | 859.0 | 27.2k | Java | 17.1 |
| [kcl-lang](https://github.com/kcl-lang)  | 201.3k | 3.9k | 1.5k | 1.2k | Rust | 16.8 |
| [fluxcd](https://github.com/fluxcd)  | 200.2k | 27.8k | 9.5k | 19.9k | Go | 16.7 |
| [longhorn](https://github.com/longhorn)  | 199.6k | 10.3k | 6.6k | 5.3k | Go | 16.6 |
| [megaease](https://github.com/megaease)  | 198.9k | 4.5k | 1.9k | 11.5k | Go | 16.6 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4k | 4.5k | 4.6k | 4.8k | Python | 16.4 |
| [tremor-rs](https://github.com/tremor-rs)  | 193.9k | 7.4k | 3.6k | 837.0 | Rust | 16.2 |
| [etcd-io](https://github.com/etcd-io)  | 187.8k | 26.5k | 11.3k | 52.3k | Go | 15.7 |
| [carvel-dev](https://github.com/carvel-dev)  | 181.2k | 8.5k | 3.3k | 3.5k | Go | 15.1 |
| [containers](https://github.com/containers)  | 173.6k | 22.2k | 12.5k | 21.5k | Go | 14.5 |
| [cert-manager](https://github.com/cert-manager)  | 163.7k | 15.7k | 5.9k | 11.0k | Go | 13.6 |
| [wasmcloud](https://github.com/wasmcloud)  | 162.0k | 5.7k | 2.4k | 1.9k | Rust | 13.5 |
| [meshery](https://github.com/meshery)  | 155.9k | 30.4k | 9.2k | 2.9k | JavaScript | 13.0 |
| [crossplane](https://github.com/crossplane)  | 154.1k | 12.8k | 3.8k | 7.9k | Go | 12.8 |
| [networkservicemesh](https://github.com/networkservicemesh)  | 152.2k | 37.4k | 30.5k | 660.0 | Go | 12.7 |
| [kubeedge](https://github.com/kubeedge)  | 150.1k | 7.3k | 3.9k | 6.8k | Go | 12.5 |
| [thanos-io](https://github.com/thanos-io)  | 144.6k | 4.3k | 4.2k | 12.1k | Go | 12.1 |
| [drakkan](https://github.com/drakkan)  | 142.1k | 1.8k | 236.0 | 8.0k | Go | 11.8 |
| [devfile](https://github.com/devfile)  | 138.4k | 4.7k | 2.2k | 322.0 | Go | 11.5 |
| [getporter](https://github.com/getporter)  | 136.8k | 6.2k | 2.5k | 1.1k | Go | 11.4 |
| [fluid-cloudnative](https://github.com/fluid-cloudnative)  | 136.8k | 2.0k | 2.0k | 1.3k | Go | 11.4 |
| [wasmedge](https://github.com/wasmedge)  | 132.5k | 3.4k | 1.5k | 5.8k | Rust | 11.1 |
| [rook](https://github.com/rook)  | 132.1k | 24.8k | 7.1k | 11.2k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 4.9k | 3.8k | 4.1k | Go | 10.9 |
| [armadaproject](https://github.com/armadaproject)  | 128.9k | 2.2k | 1.8k | 286.0 | Go | 10.7 |
| [keylime](https://github.com/keylime)  | 127.6k | 2.4k | 1.5k | 399.0 | C | 10.6 |
| [nocalhost](https://github.com/nocalhost)  | 123.0k | 6.1k | 2.1k | 1.6k | Go | 10.2 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 4.9k | 3.8k | 5.1k | Go | 10.0 |
| [kcp-dev](https://github.com/kcp-dev)  | 117.7k | 6.5k | 2.5k | 2.2k | Go | 9.8 |
| [artifacthub](https://github.com/artifacthub)  | 116.7k | 1.8k | 2.3k | 1.2k | TypeScript | 9.7 |
| [openelb](https://github.com/openelb)  | 116.4k | 498.0 | 283.0 | 1.3k | JavaScript | 9.7 |
| [kubescape](https://github.com/kubescape)  | 115.3k | 12.5k | 2.0k | 8.4k | Go | 9.6 |
| [serverless-devs](https://github.com/serverless-devs)  | 115.1k | 3.5k | 460.0 | 1.5k | JavaScript | 9.6 |
| [paralus](https://github.com/paralus)  | 112.6k | 1.6k | 437.0 | 758.0 | Go | 9.4 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 111.1k | 19.1k | 3.3k | 4.0k | Python | 9.3 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 107.2k | 3.3k | 3.2k | 5.9k | Go | 8.9 |
| [in-toto](https://github.com/in-toto)  | 106.7k | 10.2k | 1.4k | 1.5k | Python | 8.9 |
| [kedacore](https://github.com/kedacore)  | 105.4k | 3.9k | 4.2k | 6.8k | Go | 8.8 |
| [pipe-cd](https://github.com/pipe-cd)  | 104.6k | 4.8k | 3.8k | 1.0k | Go | 8.7 |
| [VKCOM](https://github.com/VKCOM)  | 101.5k | 12.1k | 5.8k | 1.1k | TypeScript | 8.4 |
| [karmada-io](https://github.com/karmada-io)  | 100.5k | 4.1k | 2.7k | 3.2k | Go | 8.4 |
| [kubeclipper](https://github.com/kubeclipper)  | 99.5k | 575.0 | 829.0 | 368.0 | Go | 8.3 |
| [openyurtio](https://github.com/openyurtio)  | 98.6k | 1.9k | 1.5k | 1.6k | Go | 8.2 |
| [bfenetworks](https://github.com/bfenetworks)  | 95.8k | 1.4k | 825.0 | 6.0k | Go | 8.0 |
| [v6d-io](https://github.com/v6d-io)  | 94.8k | 1.0k | 882.0 | 689.0 | C++ | 7.9 |
| [spidernet-io](https://github.com/spidernet-io)  | 88.5k | 5.9k | 4.4k | 582.0 | Go | 7.4 |
| [serverlessworkflow](https://github.com/serverlessworkflow)  | 88.4k | 2.4k | 1.2k | 951.0 | C# | 7.4 |
| [fnproject](https://github.com/fnproject)  | 87.4k | 8.5k | 2.5k | 6.4k | Go | 7.3 |
| [firecracker-microvm](https://github.com/firecracker-microvm)  | 82.5k | 5.8k | 3.1k | 23.9k | Rust | 6.9 |
| [openservicemesh](https://github.com/openservicemesh)  | 81.9k | 5.2k | 4.1k | 2.6k | Go | 6.8 |
| [hwameistor](https://github.com/hwameistor)  | 80.4k | 2.8k | 1.2k | 474.0 | Go | 6.7 |
| [cloudevents](https://github.com/cloudevents)  | 79.5k | 4.0k | 2.6k | 5.8k | Go | 6.6 |
| [opencontainers](https://github.com/opencontainers)  | 78.9k | 14.0k | 6.4k | 18.2k | Go | 6.6 |
| [coredns](https://github.com/coredns)  | 76.2k | 5.5k | 4.7k | 11.5k | Go | 6.3 |
| [project-zot](https://github.com/project-zot)  | 75.8k | 1.1k | 1.2k | 333.0 | Go | 6.3 |
| [gotd](https://github.com/gotd)  | 74.5k | 6.2k | 2.4k | 1.2k | Go | 6.2 |
| [ten-nancy](https://github.com/ten-nancy)  | 73.1k | 9.0 | 0.0 | 0.0 | C++ | 6.1 |
| [spectralops](https://github.com/spectralops)  | 73.0k | 359.0 | 63.0 | 1.2k | JavaScript | 6.1 |
| [opencost](https://github.com/opencost)  | 72.8k | 3.9k | 1.5k | 3.6k | Go | 6.1 |
| [notaryproject](https://github.com/notaryproject)  | 72.3k | 3.8k | 2.1k | 3.4k | Go | 6.0 |
| [helm](https://github.com/helm)  | 71.5k | 28.9k | 26.9k | 48.8k | Go | 6.0 |
| [cri-o](https://github.com/cri-o)  | 71.5k | 9.2k | 5.8k | 4.5k | Go | 6.0 |
| [metal3-io](https://github.com/metal3-io)  | 71.0k | 9.5k | 4.6k | 1.0k | Go | 5.9 |
| [meilisearch](https://github.com/meilisearch)  | 70.7k | 9.3k | 1.9k | 42.9k | Rust | 5.9 |
| [volcano-sh](https://github.com/volcano-sh)  | 70.0k | 5.5k | 2.0k | 3.2k | Go | 5.8 |
| [submariner-io](https://github.com/submariner-io)  | 69.8k | 10.0k | 9.1k | 2.3k | Go | 5.8 |
| [oras-project](https://github.com/oras-project)  | 66.1k | 3.6k | 1.2k | 1.2k | Go | 5.5 |
| [microcks](https://github.com/microcks)  | 65.1k | 2.8k | 368.0 | 1.3k | Java | 5.4 |
| [ent](https://github.com/ent)  | 63.2k | 2.2k | 2.3k | 14.8k | Go | 5.3 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 1.8k | 2.2k | 2.2k | Go | 5.2 |
| [open-feature](https://github.com/open-feature)  | 60.8k | 5.6k | 3.5k | 800.0 | Go | 5.1 |
| [vuejs](https://github.com/vuejs)  | 60.3k | 3.6k | 2.5k | 206.7k | TypeScript | 5.0 |
| [openfga](https://github.com/openfga)  | 60.2k | 1.8k | 1.5k | 1.0k | Go | 5.0 |
| [kanisterio](https://github.com/kanisterio)  | 59.8k | 2.5k | 2.3k | 707.0 | Go | 5.0 |
| [openfaas](https://github.com/openfaas)  | 59.0k | 9.0k | 4.1k | 31.9k | Go | 4.9 |
| [telepresenceio](https://github.com/telepresenceio)  | 58.0k | 8.1k | 1.7k | 5.7k | Go | 4.8 |
| [projectriff](https://github.com/projectriff)  | 57.7k | 6.6k | 4.2k | 887.0 | Go | 4.8 |
| [external-secrets](https://github.com/external-secrets)  | 57.3k | 2.8k | 1.8k | 5.1k | Go | 4.8 |
| [ariga](https://github.com/ariga)  | 55.7k | 2.2k | 2.2k | 4.8k | Go | 4.6 |
| [rkt](https://github.com/rkt)  | 52.4k | 5.6k | 2.5k | 8.9k | Go | 4.4 |
| [trickstercache](https://github.com/trickstercache)  | 52.1k | 1.1k | 458.0 | 1.9k | Go | 4.3 |
| [cloud-bulldozer](https://github.com/cloud-bulldozer)  | 51.9k | 7.9k | 3.6k | 914.0 | Python | 4.3 |
| [headlamp-k8s](https://github.com/headlamp-k8s)  | 51.4k | 3.7k | 1.1k | 1.3k | TypeScript | 4.3 |
| [theupdateframework](https://github.com/theupdateframework)  | 50.3k | 9.1k | 2.8k | 2.6k | Python | 4.2 |
| [slimtoolkit](https://github.com/slimtoolkit)  | 49.7k | 1.3k | 359.0 | 18.3k | Go | 4.1 |
| [superedge](https://github.com/superedge)  | 48.2k | 487.0 | 425.0 | 913.0 | Go | 4.0 |
| [virtual-kubelet](https://github.com/virtual-kubelet)  | 48.1k | 2.8k | 1.3k | 4.4k | Go | 4.0 |
| [kudobuilder](https://github.com/kudobuilder)  | 47.9k | 2.2k | 2.0k | 1.8k | Go | 4.0 |
| [xline-kv](https://github.com/xline-kv)  | 47.9k | 1.3k | 723.0 | 539.0 | Rust | 4.0 |
| [distribution](https://github.com/distribution)  | 45.7k | 5.1k | 1.9k | 7.6k | Go | 3.8 |
| [k3s-io](https://github.com/k3s-io)  | 45.0k | 3.6k | 3.3k | 25.5k | Go | 3.8 |
| [openfunction](https://github.com/openfunction)  | 44.5k | 1.5k | 915.0 | 1.2k | Go | 3.7 |
| [tinkerbell](https://github.com/tinkerbell)  | 44.3k | 8.3k | 2.9k | 1.9k | Go | 3.7 |
| [kubedl-io](https://github.com/kubedl-io)  | 44.1k | 1.3k | 369.0 | 488.0 | Go | 3.7 |
| [skooner-k8s](https://github.com/skooner-k8s)  | 44.0k | 507.0 | 213.0 | 1.1k | JavaScript | 3.7 |
| [tauri-apps](https://github.com/tauri-apps)  | 43.4k | 4.5k | 4.3k | 76.7k | Rust | 3.6 |
| [metallb](https://github.com/metallb)  | 40.9k | 2.3k | 1.3k | 5.7k | Go | 3.4 |
| [vitalif](https://github.com/vitalif)  | 40.9k | 1.7k | 14.0 | 105.0 | C++ | 3.4 |
| [aeraki-mesh](https://github.com/aeraki-mesh)  | 40.5k | 992.0 | 368.0 | 868.0 | Go | 3.4 |
| [capsule-rs](https://github.com/capsule-rs)  | 40.4k | 252.0 | 140.0 | 356.0 | Rust | 3.4 |
| [kubearmor](https://github.com/kubearmor)  | 40.3k | 5.7k | 2.2k | 733.0 | Go | 3.4 |
| [devstream-io](https://github.com/devstream-io)  | 39.3k | 2.9k | 1.1k | 747.0 | Go | 3.3 |
| [vscode-kubernetes-tools](https://github.com/vscode-kubernetes-tools)  | 39.3k | 634.0 | 613.0 | 602.0 | TypeScript | 3.3 |
| [ogen-go](https://github.com/ogen-go)  | 39.2k | 4.0k | 1.1k | 473.0 | Go | 3.3 |
| [jquery](https://github.com/jquery)  | 37.2k | 6.7k | 3.0k | 58.8k | JavaScript | 3.1 |
| [kube-logging](https://github.com/kube-logging)  | 35.9k | 2.9k | 1.4k | 1.4k | Go | 3.0 |
| [kuasar-io](https://github.com/kuasar-io)  | 35.7k | 156.0 | 78.0 | 972.0 | Rust | 3.0 |
| [sustainable-computing-io](https://github.com/sustainable-computing-io)  | 35.5k | 4.6k | 1.7k | 944.0 | Go | 3.0 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 2.8k | 1.2k | 6.5k | Go | 2.9 |
| [clusternet](https://github.com/clusternet)  | 33.4k | 948.0 | 545.0 | 1.1k | Go | 2.8 |
| [sealerio](https://github.com/sealerio)  | 32.7k | 1.2k | 1.3k | 1.8k | Go | 2.7 |
| [kubestellar](https://github.com/kubestellar)  | 32.6k | 4.4k | 1.1k | 215.0 | Go | 2.7 |
| [clusterpedia-io](https://github.com/clusterpedia-io)  | 31.6k | 1.1k | 541.0 | 597.0 | Go | 2.6 |
| [kube-rs](https://github.com/kube-rs)  | 31.6k | 3.0k | 775.0 | 2.2k | Rust | 2.6 |
| [dexidp](https://github.com/dexidp)  | 29.3k | 2.6k | 1.8k | 8.0k | Go | 2.4 |
| [piraeusdatastore](https://github.com/piraeusdatastore)  | 28.7k | 1.4k | 522.0 | 748.0 | Go | 2.4 |
| [opentracing](https://github.com/opentracing)  | 27.4k | 2.7k | 1.4k | 10.1k | Java | 2.3 |
| [krustlet](https://github.com/krustlet)  | 23.6k | 2.4k | 530.0 | 3.4k | Rust | 2.0 |
| [alacritty](https://github.com/alacritty)  | 23.2k | 2.3k | 2.3k | 52.3k | Rust | 1.9 |
| [curiefense](https://github.com/curiefense)  | 22.3k | 2.9k | 817.0 | 617.0 | Rust | 1.9 |
| [redhat-chaos](https://github.com/redhat-chaos)  | 22.3k | 911.0 | 784.0 | 317.0 | Python | 1.9 |
| [projectcapsule](https://github.com/projectcapsule)  | 21.9k | 1.2k | 744.0 | 1.4k | Go | 1.8 |
| [fabedge](https://github.com/fabedge)  | 21.6k | 941.0 | 429.0 | 482.0 | Go | 1.8 |
| [project-akri](https://github.com/project-akri)  | 20.6k | 702.0 | 448.0 | 946.0 | Rust | 1.7 |
| [k8gb-io](https://github.com/k8gb-io)  | 19.4k | 1.1k | 918.0 | 571.0 | Go | 1.6 |
| [foniod](https://github.com/foniod)  | 18.1k | 1.7k | 493.0 | 2.0k | Rust | 1.5 |
| [lima-vm](https://github.com/lima-vm)  | 17.4k | 2.1k | 934.0 | 11.7k | Go | 1.4 |
| [schemahero](https://github.com/schemahero)  | 17.0k | 1.4k | 773.0 | 765.0 | Go | 1.4 |
| [hexa-org](https://github.com/hexa-org)  | 16.5k | 491.0 | 118.0 | 70.0 | Go | 1.4 |
| [getsops](https://github.com/getsops)  | 15.4k | 1.6k | 643.0 | 14.5k | Go | 1.3 |
| [pallets](https://github.com/pallets)  | 15.0k | 5.2k | 2.5k | 66.2k | Python | 1.2 |
| [k8up-io](https://github.com/k8up-io)  | 14.1k | 2.0k | 716.0 | 412.0 | Go | 1.2 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 749.0 | 118.0 | 588.0 | Go | 1.1 |
| [servicemeshinterface](https://github.com/servicemeshinterface)  | 11.6k | 966.0 | 372.0 | 1.2k | Go | 970.0m |
| [kube-vip](https://github.com/kube-vip)  | 10.9k | 918.0 | 320.0 | 1.3k | Go | 910.0m |
| [kuberhealthy](https://github.com/kuberhealthy)  | 10.8k | 3.1k | 721.0 | 1.7k | Go | 900.0m |
| [k8sgpt-ai](https://github.com/k8sgpt-ai)  | 10.7k | 1.5k | 989.0 | 3.9k | Go | 890.0m |
| [ko-build](https://github.com/ko-build)  | 9.5k | 793.0 | 721.0 | 5.8k | Go | 790.0m |
| [eraser-dev](https://github.com/eraser-dev)  | 9.4k | 587.0 | 728.0 | 330.0 | Go | 780.0m |
| [devspace-cloud](https://github.com/devspace-cloud)  | 8.7k | 348.0 | 83.0 | 113.0 | Go | 720.0m |
| [tellerops](https://github.com/tellerops)  | 7.7k | 179.0 | 80.0 | 1.6k | Go | 640.0m |
| [cni-genie](https://github.com/cni-genie)  | 5.9k | 484.0 | 132.0 | 499.0 | Go | 490.0m |
| [project-copacetic](https://github.com/project-copacetic)  | 5.4k | 483.0 | 475.0 | 675.0 | Go | 450.0m |
| [merbridge](https://github.com/merbridge)  | 4.8k | 271.0 | 279.0 | 586.0 | Go | 400.0m |
| [opcr-io](https://github.com/opcr-io)  | 3.6k | 597.0 | 153.0 | 181.0 | Go | 300.0m |
| [krator-rs](https://github.com/krator-rs)  | 2.9k | 498.0 | 33.0 | 113.0 | Rust | 250.0m |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [openobservability](https://github.com/openobservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [service-mesh-performance](https://github.com/service-mesh-performance)  | 2.2k | 797.0 | 217.0 | 237.0 | JavaScript | 180.0m |
| [kubereboot](https://github.com/kubereboot)  | 2.0k | 1.0k | 558.0 | 1.8k | Go | 170.0m |
| [g-research](https://github.com/g-research)  | 870.0 | 9.0 | 0.0 | 4.0 | Go | 70.0m |
| [gitops-working-group](https://github.com/gitops-working-group)  | 0.0 | 43.0 | 57.0 | 552.0 | N/A | 0.0 |


### Repositories
| Repository | SLOC | Commits | PRs | Stars | Language | Effort |
|------------|------|---------|-----|-------|----------|--------|
| [torvalds/linux](https://github.com/torvalds/linux/)  | 19.1M | 1.3M | 792.0 | 168.9k | C | 1.6k |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3.4M | 162.5k | 25.5k | 182.0k | C++ | 279.2 |
| [ydb-platform/nbs](https://github.com/ydb-platform/nbs/)  | 3.2M | 802.0 | 0.0 | 14.0 | C++ | 267.1 |
| [ydb-platform/ydb](https://github.com/ydb-platform/ydb/)  | 2.8M | 9.4k | 35.0 | 2.8k | C++ | 237.4 |
| [ytsaurus/ytsaurus](https://github.com/ytsaurus/ytsaurus/)  | 2.5M | 65.5k | 5.0 | 799.0 | C++ | 204.8 |
| [elastic/elasticsearch](https://github.com/elastic/elasticsearch/)  | 2.4M | 76.3k | 72.2k | 67.3k | Java | 201.9 |
| [apache/hadoop](https://github.com/apache/hadoop/)  | 2.0M | 27.2k | 6.7k | 14.3k | Java | 165.6 |
| [godotengine/godot](https://github.com/godotengine/godot/)  | 2.0M | 62.7k | 40.2k | 82.5k | C++ | 164.8 |
| [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/)  | 1.7M | 78.5k | 48.5k | 26.8k | Go | 139.7 |
| [rust-lang/rust](https://github.com/rust-lang/rust/)  | 1.5M | 251.7k | 71.0k | 92.3k | Rust | 126.5 |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1.5M | 114.5k | 73.6k | 96.5k | Go | 124.0 |
| [apache/camel](https://github.com/apache/camel/)  | 1.4M | 64.6k | 10.1k | 4.8k | Java | 117.9 |
| [python/cpython](https://github.com/python/cpython/)  | 1.4M | 121.5k | 49.8k | 59.3k | Python | 114.4 |
| [openkruise/game.openkruise.io](https://github.com/openkruise/game.openkruise.io/)  | 1.2M | 4.0 | 0.0 | 0.0 | JavaScript | 99.0 |
| [apache/hive](https://github.com/apache/hive/)  | 1.1M | 16.8k | 4.3k | 4.8k | Java | 93.6 |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1.1M | 109.8k | 32.2k | 27.4k | C++ | 92.7 |
| [apache/harmony-classlib](https://github.com/apache/harmony-classlib/)  | 1.1M | 7.5k | 3.0 | 6.0 | Java | 92.4 |
| [golang/go](https://github.com/golang/go/)  | 1.1M | 59.7k | 3.9k | 119.1k | Go | 88.4 |
| [kubewarden/kubecon-24-eu-kubewarden](https://github.com/kubewarden/kubecon-24-eu-kubewarden/)  | 1.0M | 41.0 | 15.0 | 1.0 | JavaScript | 87.2 |
| [prestodb/presto](https://github.com/prestodb/presto/)  | 1.0M | 20.8k | 13.9k | 14.6k | Java | 86.0 |
| [nodejs/node](https://github.com/nodejs/node/)  | 966.7k | 41.0k | 33.0k | 103.3k | JavaScript | 80.6 |
| [envoyproxy/mobile-website](https://github.com/envoyproxy/mobile-website/)  | 916.9k | 4.7k | 15.0 | 0.0 | JavaScript | 76.4 |
| [konveyor/github-actions](https://github.com/konveyor/github-actions/)  | 897.5k | 17.0 | 13.0 | 0.0 | JavaScript | 74.8 |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889.8k | 41.4k | 33.6k | 54.3k | TypeScript | 74.2 |
| [ziglang/zig](https://github.com/ziglang/zig/)  | 888.6k | 28.9k | 9.2k | 30.0k | Zig | 74.0 |
| [inclavare-containers/tng-envoy](https://github.com/inclavare-containers/tng-envoy/)  | 868.5k | 21.5k | 0.0 | 0.0 | C++ | 72.4 |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 807.6k | 17.3k | 17.1k | 21.6k | C++ | 67.3 |
| [brigadecore/go-libgit2](https://github.com/brigadecore/go-libgit2/)  | 782.3k | 3.0 | 0.0 | 0.0 | JavaScript | 65.2 |
| [keycloak/keycloak](https://github.com/keycloak/keycloak/)  | 675.4k | 23.6k | 13.8k | 18.5k | Java | 56.3 |
| [apache/cassandra](https://github.com/apache/cassandra/)  | 632.5k | 28.6k | 2.3k | 8.0k | Java | 52.7 |
| [ytsaurus/ytsaurus-cpp-sdk](https://github.com/ytsaurus/ytsaurus-cpp-sdk/)  | 623.2k | 1.9k | 3.0 | 2.0 | C++ | 51.9 |
| [elastic/beats](https://github.com/elastic/beats/)  | 616.2k | 17.3k | 28.4k | 12.0k | Go | 51.4 |
| [vitessio/vitess](https://github.com/vitessio/vitess/)  | 611.0k | 32.3k | 9.6k | 15.7k | Go | 50.9 |
| [apache/lucenenet](https://github.com/apache/lucenenet/)  | 608.0k | 6.7k | 619.0 | 2.0k | C# | 50.7 |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 514.8k | 52.7k | 21.6k | 37.2k | C++ | 42.9 |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 512.3k | 96.0 | 103.0 | 22.0 | C++ | 42.7 |
| [falcosecurity/libs](https://github.com/falcosecurity/libs/)  | 504.5k | 7.4k | 964.0 | 156.0 | C Header | 42.0 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 504.5k | 8.5k | 454.0 | 205.0 | C++ | 42.0 |
| [denoland/deno](https://github.com/denoland/deno/)  | 498.2k | 11.1k | 12.5k | 92.8k | JavaScript | 41.5 |
| [cubefs/cubefs](https://github.com/cubefs/cubefs/)  | 481.6k | 2.5k | 1.4k | 3.2k | Go | 40.1 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477.3k | 4.2k | 3.6k | 4.4k | Go | 39.8 |
| [apache/openjpa](https://github.com/apache/openjpa/)  | 464.1k | 5.3k | 98.0 | 114.0 | Java | 38.7 |
| [apache/activemq](https://github.com/apache/activemq/)  | 444.6k | 11.4k | 999.0 | 2.1k | Java | 37.0 |
| [apache/poi](https://github.com/apache/poi/)  | 424.4k | 12.6k | 457.0 | 1.6k | Java | 35.4 |
| [konveyor/tackle-diva](https://github.com/konveyor/tackle-diva/)  | 400.7k | 106.0 | 88.0 | 39.0 | Java | 33.4 |
| [tikv/tikv](https://github.com/tikv/tikv/)  | 395.8k | 7.7k | 11.8k | 14.4k | Rust | 33.0 |
| [netdata/charts](https://github.com/netdata/charts/)  | 395.7k | 582.0 | 167.0 | 9.0 | JavaScript | 33.0 |
| [backstage/backstage](https://github.com/backstage/backstage/)  | 393.7k | 38.2k | 12.3k | 21.1k | TypeScript | 32.8 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389.2k | 9.6k | 16.2k | 1.5k | Go | 32.4 |
| [netdata/netdata](https://github.com/netdata/netdata/)  | 380.5k | 15.2k | 7.3k | 62.4k | C | 31.7 |
| [envoyproxy/archive](https://github.com/envoyproxy/archive/)  | 377.7k | 523.0 | 15.0 | 0.0 | JavaScript | 31.5 |
| [django/django](https://github.com/django/django/)  | 365.8k | 32.6k | 18.0k | 76.5k | Python | 30.5 |
| [cdk8s-team/cdk8s-examples](https://github.com/cdk8s-team/cdk8s-examples/)  | 360.6k | 34.0 | 36.0 | 4.0 | Java | 30.1 |
| [facebook/react](https://github.com/facebook/react/)  | 357.5k | 16.6k | 14.8k | 221.0k | JavaScript | 29.8 |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 352.4k | 12.4k | 1.3k | 5.2k | C++ | 29.4 |
| [athenz/athenz](https://github.com/athenz/athenz/)  | 348.7k | 2.6k | 1.9k | 760.0 | Java | 29.1 |
| [apache/ofbiz](https://github.com/apache/ofbiz/)  | 348.4k | 24.4k | 27.0 | 753.0 | Java | 29.0 |
| [AthenZ/athenz](https://github.com/AthenZ/athenz/)  | 346.1k | 2.6k | 1.9k | 751.0 | Java | 28.8 |
| [opencurve/curve](https://github.com/opencurve/curve/)  | 339.3k | 2.8k | 1.6k | 1.8k | C++ | 28.3 |
| [istio/istio](https://github.com/istio/istio/)  | 339.1k | 19.9k | 26.5k | 32.5k | Go | 28.3 |
| [kubernetes/autoscaler](https://github.com/kubernetes/autoscaler/)  | 332.8k | 6.5k | 3.7k | 6.6k | Go | 27.7 |
| [cilium/cilium](https://github.com/cilium/cilium/)  | 324.8k | 22.3k | 17.0k | 14.6k | Go | 27.1 |
| [docker/docker-ce](https://github.com/docker/docker-ce/)  | 324.4k | 54.3k | 662.0 | 5.6k | Go | 27.0 |
| [apache/harmony-drlvm](https://github.com/apache/harmony-drlvm/)  | 316.3k | 2.2k | 0.0 | 12.0 | C++ | 26.4 |
| [apache/jackrabbit](https://github.com/apache/jackrabbit/)  | 315.9k | 9.1k | 112.0 | 305.0 | Java | 26.3 |
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310.9k | 4.9k | 14.0 | 812.0 | Java | 25.9 |
| [apache/mesos](https://github.com/apache/mesos/)  | 305.6k | 18.3k | 547.0 | 5.2k | C++ | 25.5 |
| [docker/labs](https://github.com/docker/labs/)  | 304.4k | 718.0 | 398.0 | 11.1k | PHP | 25.4 |
| [kubeovn/kubevirt-dpdk](https://github.com/kubeovn/kubevirt-dpdk/)  | 299.8k | 17.1k | 0.0 | 2.0 | Go | 25.0 |
| [kubeflow/pipelines](https://github.com/kubeflow/pipelines/)  | 296.1k | 5.6k | 6.8k | 3.4k | Python | 24.7 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284.4k | 1.6k | 1.3k | 465.0 | C Header | 23.7 |
| [pravega/pravega](https://github.com/pravega/pravega/)  | 283.9k | 3.2k | 3.2k | 1.9k | Java | 23.6 |
| [kubevirt/kubevirt](https://github.com/kubevirt/kubevirt/)  | 281.1k | 16.3k | 7.0k | 3.9k | Go | 23.4 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274.6k | 9.4k | 9.4k | 12.9k | Rust | 22.9 |
| [ydb-platform/ydb-slo-action](https://github.com/ydb-platform/ydb-slo-action/)  | 272.2k | 37.0 | 8.0 | 1.0 | JavaScript | 22.7 |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 270.6k | 51.9k | 24.6k | 3.6k | Go | 22.6 |
| [cdk8s-team/cdk8s](https://github.com/cdk8s-team/cdk8s/)  | 266.0k | 1.0k | 888.0 | 3.5k | TypeScript | 22.2 |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 264.4k | 1.6k | 36.0 | 24.0 | JavaScript | 22.0 |
| [metrico/influxdb_iox](https://github.com/metrico/influxdb_iox/)  | 262.3k | 14.0k | 0.0 | 0.0 | Rust | 21.9 |
| [apache/cayenne](https://github.com/apache/cayenne/)  | 261.7k | 7.1k | 580.0 | 287.0 | Java | 21.8 |
| [ydb-platform/ydb-cpp-sdk](https://github.com/ydb-platform/ydb-cpp-sdk/)  | 260.0k | 10.5k | 56.0 | 6.0 | C++ | 21.7 |
| [kyverno/kyverno](https://github.com/kyverno/kyverno/)  | 259.5k | 5.5k | 3.9k | 3.6k | Go | 21.6 |
| [inclavare-containers/deterministic~](https://github.com/inclavare-containers/deterministic-builds/)  | 259.3k | 28.0 | 4.0 | 1.0 | C Header | 21.6 |
| [apache/wicket](https://github.com/apache/wicket/)  | 258.1k | 21.4k | 550.0 | 684.0 | Java | 21.5 |
| [apache/httpd](https://github.com/apache/httpd/)  | 256.2k | 33.5k | 308.0 | 3.2k | C | 21.4 |
| [apache/pig](https://github.com/apache/pig/)  | 254.4k | 3.7k | 38.0 | 655.0 | Java | 21.2 |
| [inspektor-gadget/inspektor-gadget](https://github.com/inspektor-gadget/inspektor-gadget/)  | 251.3k | 2.8k | 1.0k | 1.3k | C Header | 20.9 |
| [moby/moby](https://github.com/moby/moby/)  | 244.5k | 45.7k | 23.1k | 65.6k | Go | 20.4 |
| [apache/cocoon](https://github.com/apache/cocoon/)  | 241.8k | 13.2k | 44.0 | 22.0 | Java | 20.1 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241.1k | 4.7k | 5.2k | 18.5k | Go | 20.1 |
| [inclavare-containers/kong_coco](https://github.com/inclavare-containers/kong_coco/)  | 239.3k | 2.0 | 11.0 | 0.0 | Lua | 19.9 |
| [matplotlib/matplotlib](https://github.com/matplotlib/matplotlib/)  | 235.9k | 50.1k | 17.9k | 19.1k | Python | 19.6 |
| [open-policy-agent/opa](https://github.com/open-policy-agent/opa/)  | 235.4k | 4.5k | 3.5k | 7.8k | Go | 19.6 |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235.2k | 5.7k | 6.9k | 10.4k | Java | 19.6 |
| [netdata/ebpf-co-re](https://github.com/netdata/ebpf-co-re/)  | 231.9k | 35.0 | 38.0 | 6.0 | C Header | 19.3 |
| [kumahq/kuma](https://github.com/kumahq/kuma/)  | 221.4k | 4.0k | 4.8k | 3.1k | Go | 18.4 |
| [apache/xmlgraphics-fop](https://github.com/apache/xmlgraphics-fop/)  | 219.8k | 8.5k | 44.0 | 154.0 | Java | 18.3 |
| [chaosblade-io/chaosblade-box](https://github.com/chaosblade-io/chaosblade-box/)  | 219.5k | 40.0 | 41.0 | 186.0 | JavaScript | 18.3 |
| [apache/directory-studio](https://github.com/apache/directory-studio/)  | 216.5k | 5.7k | 27.0 | 101.0 | Java | 18.0 |
| [uber/peloton](https://github.com/uber/peloton/)  | 216.4k | 705.0 | 10.0 | 636.0 | Go | 18.0 |
| [antrea-io/antrea](https://github.com/antrea-io/antrea/)  | 198.7k | 2.7k | 3.5k | 1.4k | Go | 16.6 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194.2k | 181.0 | 125.0 | 1.2k | C Header | 16.2 |
| [apache/xmlgraphics-batik](https://github.com/apache/xmlgraphics-batik/)  | 193.7k | 3.6k | 36.0 | 173.0 | Java | 16.1 |
| [goharbor/harbor](https://github.com/goharbor/harbor/)  | 191.5k | 11.6k | 8.0k | 19.5k | Go | 16.0 |
| [apache/tapestry4](https://github.com/apache/tapestry4/)  | 190.8k | 3.8k | 11.0 | 7.0 | Java | 15.9 |
| [nats-io/nats-server](https://github.com/nats-io/nats-server/)  | 190.7k | 6.8k | 2.6k | 12.4k | Go | 15.9 |
| [apache/avro](https://github.com/apache/avro/)  | 189.0k | 3.7k | 2.2k | 2.5k | Java | 15.8 |
| [alibaba/fastjson](https://github.com/alibaba/fastjson/)  | 188.1k | 4.0k | 848.0 | 25.2k | Java | 15.7 |
| [googlecontainertools/skaffold](https://github.com/googlecontainertools/skaffold/)  | 186.7k | 8.9k | 5.3k | 14.6k | Go | 15.6 |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186.6k | 1.0 | 2.0 | 2.0 | JavaScript | 15.6 |
| [fluent/fluent-bit](https://github.com/fluent/fluent-bit/)  | 182.4k | 9.3k | 3.3k | 4.4k | C | 15.2 |
| [apache/directory-server](https://github.com/apache/directory-server/)  | 181.5k | 10.0k | 80.0 | 118.0 | Java | 15.1 |
| [cncf/cncf.io](https://github.com/cncf/cncf.io/)  | 180.9k | 1.9k | 88.0 | 66.0 | PHP | 15.1 |
| [cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian/)  | 175.1k | 4.1k | 4.4k | 4.7k | Python | 14.6 |
| [strimzi/strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator/)  | 175.0k | 5.1k | 4.7k | 3.7k | Java | 14.6 |
| [containers/podman](https://github.com/containers/podman/)  | 173.6k | 22.2k | 12.5k | 21.5k | Go | 14.5 |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173.3k | 8.9k | 5.3k | 2.8k | Go | 14.4 |
| [apache/xalan-j](https://github.com/apache/xalan-j/)  | 170.9k | 4.6k | 6.0 | 25.0 | Java | 14.2 |
| [argoproj/argo-workflows](https://github.com/argoproj/argo-workflows/)  | 169.8k | 4.1k | 4.8k | 12.6k | Go | 14.2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.9k | 8.8k | 5.1k | 1.2k | Java | 14.0 |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167.6k | 19.4k | 10.4k | 14.8k | Go | 14.0 |
| [apache/tika](https://github.com/apache/tika/)  | 167.3k | 7.3k | 1.1k | 1.7k | Java | 13.9 |
| [keptn/keptn](https://github.com/keptn/keptn/)  | 162.3k | 8.3k | 6.1k | 1.7k | Go | 13.5 |
| [apache/couchdb](https://github.com/apache/couchdb/)  | 159.2k | 13.2k | 2.8k | 5.7k | Erlang | 13.3 |
| [apache/shindig](https://github.com/apache/shindig/)  | 158.2k | 4.8k | 7.0 | 69.0 | Java | 13.2 |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157.3k | 1.7k | 1.3k | 316.0 | Java | 13.1 |
| [spiffe/spire](https://github.com/spiffe/spire/)  | 156.7k | 5.8k | 2.8k | 1.4k | Go | 13.1 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155.1k | 10.9k | 6.8k | 47.1k | Go | 12.9 |
| [apache/etch](https://github.com/apache/etch/)  | 153.2k | 549.0 | 0.0 | 17.0 | Java | 12.8 |
| [werf/actions](https://github.com/werf/actions/)  | 153.1k | 118.0 | 60.0 | 68.0 | JavaScript | 12.8 |
| [go-faster/ytsaurus-ui](https://github.com/go-faster/ytsaurus-ui/)  | 151.4k | 636.0 | 0.0 | 0.0 | TypeScript | 12.6 |
| [apache/stdcxx](https://github.com/apache/stdcxx/)  | 150.4k | 2.2k | 0.0 | 57.0 | C++ | 12.5 |
| [containerd/containerd](https://github.com/containerd/containerd/)  | 150.2k | 12.0k | 5.9k | 13.3k | Go | 12.5 |
| [argoproj/argo-cd](https://github.com/argoproj/argo-cd/)  | 149.4k | 5.1k | 5.8k | 12.3k | Go | 12.4 |
| [kubevela/samples](https://github.com/kubevela/samples/)  | 146.8k | 219.0 | 118.0 | 93.0 | JavaScript | 12.2 |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144.5k | 5.9k | 1.8k | 8.1k | Go | 12.0 |
| [kubevela/kubevela](https://github.com/kubevela/kubevela/)  | 144.3k | 3.6k | 3.8k | 5.0k | Go | 12.0 |
| [kubeflow/kfp-tekton](https://github.com/kubeflow/kfp-tekton/)  | 143.9k | 798.0 | 1.0k | 152.0 | TypeScript | 12.0 |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143.2k | 4.5k | 3.9k | 17.7k | Go | 11.9 |
| [apache/maven-plugins](https://github.com/apache/maven-plugins/)  | 142.5k | 15.1k | 146.0 | 242.0 | Java | 11.9 |
| [drakkan/sftpgo](https://github.com/drakkan/sftpgo/)  | 142.1k | 1.8k | 236.0 | 8.0k | Go | 11.8 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.1k | 929.0 | 1.8k | 224.0 | C Header | 11.8 |
| [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize/)  | 139.5k | 6.3k | 3.0k | 9.4k | Go | 11.6 |
| [tikv/pd](https://github.com/tikv/pd/)  | 135.9k | 3.3k | 4.6k | 964.0 | Go | 11.3 |
| [apache/zookeeper](https://github.com/apache/zookeeper/)  | 135.8k | 2.5k | 2.0k | 11.3k | Java | 11.3 |
| [fluid-cloudnative/fluid](https://github.com/fluid-cloudnative/fluid/)  | 134.8k | 1.9k | 1.9k | 1.3k | Go | 11.2 |
| [confidential-containers/kata-conta~](https://github.com/confidential-containers/kata-containers-CCv0/)  | 133.2k | 10.9k | 2.0 | 5.0 | Go | 11.1 |
| [apache/apr-iconv](https://github.com/apache/apr-iconv/)  | 132.0k | 210.0 | 0.0 | 19.0 | C | 11.0 |
| [kubernetes/apiserver](https://github.com/kubernetes/apiserver/)  | 131.3k | 6.2k | 25.0 | 494.0 | Go | 10.9 |
| [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers/)  | 130.8k | 10.1k | 3.0k | 3.1k | Go | 10.9 |
| [ytsaurus/ytsaurus-ui](https://github.com/ytsaurus/ytsaurus-ui/)  | 130.7k | 2.0 | 2.0 | 8.0 | TypeScript | 10.9 |
| [siderolabs/talos](https://github.com/siderolabs/talos/)  | 130.2k | 3.9k | 5.2k | 3.7k | Go | 10.8 |
| [etcd-io/etcd](https://github.com/etcd-io/etcd/)  | 125.7k | 19.3k | 9.2k | 42.8k | Go | 10.5 |
| [cncf/demo](https://github.com/cncf/demo/)  | 122.9k | 670.0 | 161.0 | 76.0 | JavaScript | 10.2 |
| [armadaproject/armada](https://github.com/armadaproject/armada/)  | 119.4k | 2.0k | 1.7k | 278.0 | Go | 9.9 |
| [cortexproject/cortex](https://github.com/cortexproject/cortex/)  | 119.0k | 4.6k | 3.5k | 5.0k | Go | 9.9 |
| [artifacthub/hub](https://github.com/artifacthub/hub/)  | 116.7k | 1.7k | 2.3k | 1.2k | TypeScript | 9.7 |
| [parallaxsecond/rust-tss-esapi](https://github.com/parallaxsecond/rust-tss-esapi/)  | 116.6k | 658.0 | 280.0 | 44.0 | Rust | 9.7 |
| [apache/aurora](https://github.com/apache/aurora/)  | 114.8k | 4.1k | 71.0 | 634.0 | Java | 9.6 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 113.7k | 2.8k | 63.0 | 10.0 | TypeScript | 9.5 |
| [thanos-io/thanos](https://github.com/thanos-io/thanos/)  | 113.2k | 3.1k | 3.6k | 11.5k | Go | 9.4 |
| [projectcontour/contour](https://github.com/projectcontour/contour/)  | 111.5k | 4.0k | 3.2k | 3.3k | Go | 9.3 |
| [netdata/go.d.plugin](https://github.com/netdata/go.d.plugin/)  | 110.3k | 1.4k | 958.0 | 145.0 | Go | 9.2 |
| [openelb/website](https://github.com/openelb/website/)  | 107.6k | 153.0 | 73.0 | 3.0 | JavaScript | 9.0 |
| [moby/buildkit](https://github.com/moby/buildkit/)  | 107.3k | 5.2k | 2.2k | 6.5k | Go | 8.9 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 107.0k | 2.3k | 1.8k | 218.0 | Go | 8.9 |
| [kubeflow/kfp-tekton-backend](https://github.com/kubeflow/kfp-tekton-backend/)  | 106.1k | 14.0 | 71.0 | 8.0 | TypeScript | 8.8 |
| [apache/synapse](https://github.com/apache/synapse/)  | 106.1k | 4.4k | 80.0 | 52.0 | Java | 8.8 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105.9k | 1.5k | 1.8k | 980.0 | Go | 8.8 |
| [emissary-ingress/emissary](https://github.com/emissary-ingress/emissary/)  | 105.3k | 17.9k | 3.3k | 4.0k | Python | 8.8 |
| [googlecontainertools/kpt-config-sy~](https://github.com/googlecontainertools/kpt-config-sync/)  | 105.0k | 983.0 | 1.2k | 222.0 | Go | 8.8 |
| [pipe-cd/pipecd](https://github.com/pipe-cd/pipecd/)  | 103.3k | 3.7k | 3.6k | 908.0 | Go | 8.6 |
| [apache/ode](https://github.com/apache/ode/)  | 101.9k | 3.8k | 8.0 | 44.0 | Java | 8.5 |
| [cockroachdb/pebble](https://github.com/cockroachdb/pebble/)  | 101.1k | 2.8k | 1.7k | 3.5k | Go | 8.4 |
| [dapr/dapr](https://github.com/dapr/dapr/)  | 101.0k | 4.0k | 3.1k | 20.7k | Go | 8.4 |
| [cert-manager/cert-manager](https://github.com/cert-manager/cert-manager/)  | 100.4k | 7.5k | 2.8k | 10.1k | Go | 8.4 |
| [openebs/maya](https://github.com/openebs/maya/)  | 100.1k | 1.8k | 1.7k | 180.0 | Go | 8.3 |
| [apache/apr](https://github.com/apache/apr/)  | 99.5k | 9.3k | 35.0 | 413.0 | C | 8.3 |
| [openkruise/kruise](https://github.com/openkruise/kruise/)  | 98.1k | 712.0 | 770.0 | 3.7k | Go | 8.2 |
| [keylime/tpm4720-keylime](https://github.com/keylime/tpm4720-keylime/)  | 97.6k | 10.0 | 0.0 | 6.0 | C | 8.1 |
| [moby/containerd](https://github.com/moby/containerd/)  | 97.5k | 8.4k | 2.0 | 16.0 | Go | 8.1 |
| [apache/continuum](https://github.com/apache/continuum/)  | 97.2k | 4.9k | 26.0 | 10.0 | Java | 8.1 |
| [knative/serving](https://github.com/knative/serving/)  | 96.9k | 8.0k | 9.4k | 4.9k | Go | 8.1 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 3.3k | 3.6k | 1.5k | Java | 8.0 |
| [apache/spamassassin](https://github.com/apache/spamassassin/)  | 96.0k | 28.4k | 2.0 | 242.0 | Perl | 8.0 |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 96.0k | 3.0k | 313.0 | 2.2k | Go | 8.0 |
| [rook/rook](https://github.com/rook/rook/)  | 95.8k | 9.6k | 7.0k | 10.8k | Go | 8.0 |
| [v6d-io/v6d](https://github.com/v6d-io/v6d/)  | 94.8k | 1.0k | 882.0 | 689.0 | C++ | 7.9 |
| [karmada-io/karmada](https://github.com/karmada-io/karmada/)  | 94.6k | 3.6k | 2.4k | 3.1k | Go | 7.9 |
| [elastic/logstash](https://github.com/elastic/logstash/)  | 92.8k | 10.7k | 9.3k | 14.0k | Ruby | 7.7 |
| [tremor-rs/tremor-runtime](https://github.com/tremor-rs/tremor-runtime/)  | 91.6k | 4.0k | 1.8k | 729.0 | Rust | 7.6 |
| [apache/commons-lang](https://github.com/apache/commons-lang/)  | 91.1k | 7.2k | 1.1k | 2.5k | Java | 7.6 |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90.1k | 21.0k | 7.6k | 26.0k | Go | 7.5 |
| [kubeclipper/kubeclipper](https://github.com/kubeclipper/kubeclipper/)  | 90.0k | 329.0 | 504.0 | 242.0 | Go | 7.5 |
| [apache/maven](https://github.com/apache/maven/)  | 90.0k | 11.8k | 1.1k | 3.6k | Java | 7.5 |
| [meshery/meshery](https://github.com/meshery/meshery/)  | 87.6k | 19.3k | 4.9k | 1.9k | Go | 7.3 |
| [wasmedge/WasmEdge](https://github.com/wasmedge/WasmEdge/)  | 87.4k | 2.5k | 1.4k | 5.6k | C++ | 7.3 |
| [cubefs/cubefs-blobstore](https://github.com/cubefs/cubefs-blobstore/)  | 87.4k | 82.0 | 3.0 | 14.0 | Go | 7.3 |
| [kcl-lang/kcl](https://github.com/kcl-lang/kcl/)  | 87.2k | 633.0 | 625.0 | 1.0k | Rust | 7.3 |
| [tremor-rs/tremor-www-docs](https://github.com/tremor-rs/tremor-www-docs/)  | 87.1k | 405.0 | 153.0 | 24.0 | JavaScript | 7.3 |
| [knative/eventing](https://github.com/knative/eventing/)  | 87.1k | 3.5k | 4.3k | 1.3k | Go | 7.3 |
| [operator-framework/operator-lifecy~](https://github.com/operator-framework/operator-lifecycle-manager/)  | 87.0k | 4.2k | 2.2k | 1.5k | Go | 7.2 |
| [moby/swarmkit](https://github.com/moby/swarmkit/)  | 86.8k | 4.2k | 2.3k | 3.0k | Go | 7.2 |
| [go-faster/minikube](https://github.com/go-faster/minikube/)  | 86.6k | 23.3k | 0.0 | 0.0 | Go | 7.2 |
| [fluent/onigmo](https://github.com/fluent/onigmo/)  | 86.6k | 1.2k | 5.0 | 2.0 | C Header | 7.2 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84.5k | 3.4k | 3.2k | 210.0 | Go | 7.0 |
| [serverless-devs/serverless-registr~](https://github.com/serverless-devs/serverless-registry-website/)  | 84.3k | 31.0 | 0.0 | 1.0 | JavaScript | 7.0 |
| [dapr/components-contrib](https://github.com/dapr/components-contrib/)  | 83.8k | 2.9k | 1.5k | 426.0 | Go | 7.0 |
| [argoproj/argo-rollouts](https://github.com/argoproj/argo-rollouts/)  | 82.6k | 1.2k | 1.5k | 1.9k | Go | 6.9 |
| [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker/)  | 82.5k | 5.8k | 3.1k | 23.9k | Rust | 6.9 |
| [apache/struts-sandbox](https://github.com/apache/struts-sandbox/)  | 82.3k | 1.9k | 13.0 | 5.0 | Java | 6.9 |
| [linkerd/linkerd2](https://github.com/linkerd/linkerd2/)  | 82.0k | 5.5k | 6.4k | 9.4k | Go | 6.8 |
| [fluent/fluentd](https://github.com/fluent/fluentd/)  | 80.5k | 6.5k | 1.9k | 11.8k | Ruby | 6.7 |
| [apache/httpcomponents-core](https://github.com/apache/httpcomponents-core/)  | 80.4k | 3.7k | 402.0 | 313.0 | Java | 6.7 |
| [werf/werf](https://github.com/werf/werf/)  | 79.3k | 12.6k | 4.5k | 3.6k | Go | 6.6 |
| [kubevirt/containerized-data-import~](https://github.com/kubevirt/containerized-data-importer/)  | 79.1k | 2.5k | 2.1k | 311.0 | Go | 6.6 |
| [docker/cli](https://github.com/docker/cli/)  | 78.9k | 8.5k | 2.7k | 3.9k | Go | 6.6 |
| [litmuschaos/litmus](https://github.com/litmuschaos/litmus/)  | 78.7k | 2.5k | 2.6k | 3.5k | TypeScript | 6.6 |
| [fluent/fluentbit-website-v3](https://github.com/fluent/fluentbit-website-v3/)  | 78.2k | 386.0 | 27.0 | 4.0 | JavaScript | 6.5 |
| [kubernetes/ingress-gce](https://github.com/kubernetes/ingress-gce/)  | 78.0k | 4.6k | 1.6k | 1.2k | Go | 6.5 |
| [kubeflow/kfserving-lts](https://github.com/kubeflow/kfserving-lts/)  | 78.0k | 795.0 | 33.0 | 10.0 | Python | 6.5 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77.2k | 3.6k | 2.2k | 244.0 | Go | 6.4 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76.5k | 630.0 | 1.2k | 3.3k | Go | 6.4 |
| [metrico/loki-apache](https://github.com/metrico/loki-apache/)  | 76.4k | 9.0 | 0.0 | 0.0 | Go | 6.4 |
| [kubevela/velaux](https://github.com/kubevela/velaux/)  | 75.2k | 589.0 | 478.0 | 95.0 | TypeScript | 6.3 |
| [chaos-mesh/chaos-mesh](https://github.com/chaos-mesh/chaos-mesh/)  | 74.9k | 1.7k | 2.4k | 5.5k | Go | 6.2 |
| [apache/jspwiki](https://github.com/apache/jspwiki/)  | 74.4k | 9.3k | 271.0 | 93.0 | Java | 6.2 |
| [go-faster/porto](https://github.com/go-faster/porto/)  | 74.2k | 121.0 | 0.0 | 1.0 | C++ | 6.2 |
| [openservicemesh/osm](https://github.com/openservicemesh/osm/)  | 73.9k | 4.5k | 3.6k | 2.6k | Go | 6.2 |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73.8k | 6.5k | 1.2k | 628.0 | Go | 6.2 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73.8k | 1.7k | 634.0 | 1.2k | Java | 6.2 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73.7k | 3.8k | 2.7k | 543.0 | Go | 6.1 |
| [apache/maven-sandbox](https://github.com/apache/maven-sandbox/)  | 73.4k | 2.3k | 2.0 | 5.0 | Java | 6.1 |
| [buildpacks/pack-private](https://github.com/buildpacks/pack-private/)  | 73.3k | 4.5k | 58.0 | 0.0 | Go | 6.1 |
| [linkerd/linkerd](https://github.com/linkerd/linkerd/)  | 73.3k | 1.4k | 1.4k | 5.4k | Scala | 6.1 |
| [bfenetworks/bfe](https://github.com/bfenetworks/bfe/)  | 73.3k | 1.1k | 730.0 | 5.8k | Go | 6.1 |
| [grafana/k6](https://github.com/grafana/k6/)  | 73.0k | 5.1k | 1.5k | 19.7k | Go | 6.1 |
| [apache/httpcomponents-client](https://github.com/apache/httpcomponents-client/)  | 72.9k | 3.4k | 445.0 | 1.3k | Java | 6.1 |
| [megaease/easegress](https://github.com/megaease/easegress/)  | 72.9k | 2.0k | 958.0 | 5.6k | Go | 6.1 |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72.4k | 4.8k | 5.3k | 2.4k | TypeScript | 6.0 |
| [dragonflyoss/Dragonfly2](https://github.com/dragonflyoss/Dragonfly2/)  | 72.3k | 2.1k | 2.5k | 2.0k | Go | 6.0 |
| [jaegertracing/jaeger](https://github.com/jaegertracing/jaeger/)  | 71.9k | 2.0k | 2.3k | 17.3k | Go | 6.0 |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71.9k | 633.0 | 232.0 | 1.1k | JavaScript | 6.0 |
| [project-zot/zot](https://github.com/project-zot/zot/)  | 71.1k | 757.0 | 928.0 | 329.0 | Go | 5.9 |
| [ten-nancy/porto](https://github.com/ten-nancy/porto/)  | 70.9k | 6.0 | 0.0 | 0.0 | C++ | 5.9 |
| [meilisearch/meilisearch](https://github.com/meilisearch/meilisearch/)  | 70.7k | 9.3k | 1.9k | 42.9k | Rust | 5.9 |
| [tikv/migration](https://github.com/tikv/migration/)  | 69.9k | 165.0 | 188.0 | 29.0 | Go | 5.8 |
| [parallaxsecond/parsec-website](https://github.com/parallaxsecond/parsec-website/)  | 69.4k | 5.0 | 3.0 | 1.0 | JavaScript | 5.8 |
| [cncf/obsolete-interactive-landscape](https://github.com/cncf/obsolete-interactive-landscape/)  | 69.1k | 364.0 | 0.0 | 5.0 | JavaScript | 5.8 |
| [longhorn/longhorn-manager](https://github.com/longhorn/longhorn-manager/)  | 68.9k | 3.0k | 1.8k | 132.0 | Go | 5.7 |
| [kedacore/keda](https://github.com/kedacore/keda/)  | 68.3k | 1.9k | 2.1k | 6.0k | Go | 5.7 |
| [opencost/opencost](https://github.com/opencost/opencost/)  | 67.9k | 3.4k | 1.3k | 3.5k | Go | 5.7 |
| [apache/commons-collections](https://github.com/apache/commons-collections/)  | 67.2k | 3.9k | 389.0 | 609.0 | Java | 5.6 |
| [nocalhost/nocalhost](https://github.com/nocalhost/nocalhost/)  | 66.8k | 3.6k | 1.3k | 1.5k | Go | 5.6 |
| [open-cluster-management-io/ocm](https://github.com/open-cluster-management-io/ocm/)  | 66.8k | 1.3k | 114.0 | 566.0 | Go | 5.6 |
| [kubeedge/kubeedge](https://github.com/kubeedge/kubeedge/)  | 66.7k | 4.7k | 2.9k | 5.7k | Go | 5.6 |
| [cdk8s-team/cdk8s-plus-go](https://github.com/cdk8s-team/cdk8s-plus-go/)  | 65.8k | 574.0 | 0.0 | 4.0 | Go | 5.5 |
| [kcp-dev/kcp](https://github.com/kcp-dev/kcp/)  | 65.7k | 5.4k | 2.0k | 2.1k | Go | 5.5 |
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65.6k | 3.8k | 204.0 | 7.4k | Go | 5.5 |
| [apache/struts1](https://github.com/apache/struts1/)  | 65.4k | 5.3k | 10.0 | 27.0 | Java | 5.5 |
| [nats-io/nats.c](https://github.com/nats-io/nats.c/)  | 64.9k | 972.0 | 399.0 | 301.0 | C | 5.4 |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63.9k | 2.8k | 1.5k | 2.1k | Go | 5.3 |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63.9k | 859.0 | 941.0 | 3.5k | C# | 5.3 |
| [hwameistor/hwameistor](https://github.com/hwameistor/hwameistor/)  | 63.8k | 2.0k | 872.0 | 456.0 | Go | 5.3 |
| [megaease/easeagent](https://github.com/megaease/easeagent/)  | 63.5k | 760.0 | 257.0 | 568.0 | Java | 5.3 |
| [ent/ent](https://github.com/ent/ent/)  | 63.2k | 2.2k | 2.3k | 14.8k | Go | 5.3 |
| [fluent/fluent-bit-website-old](https://github.com/fluent/fluent-bit-website-old/)  | 63.2k | 19.0 | 0.0 | 2.0 | JavaScript | 5.3 |
| [ydb-platform/ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk/)  | 62.9k | 2.9k | 465.0 | 80.0 | Go | 5.2 |
| [werf/trdl-vault-actions](https://github.com/werf/trdl-vault-actions/)  | 62.2k | 13.0 | 7.0 | 0.0 | JavaScript | 5.2 |
| [kubeovn/kube-ovn](https://github.com/kubeovn/kube-ovn/)  | 61.8k | 3.0k | 1.9k | 1.5k | Go | 5.2 |
| [knative/pkg](https://github.com/knative/pkg/)  | 61.2k | 2.0k | 2.3k | 234.0 | Go | 5.1 |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61.1k | 1.1k | 995.0 | 188.0 | Go | 5.1 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 1.7k | 2.0k | 1.8k | TypeScript | 5.1 |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60.9k | 1.2k | 447.0 | 27.0 | Rust | 5.1 |
| [apache/abdera](https://github.com/apache/abdera/)  | 60.5k | 1.5k | 4.0 | 18.0 | Java | 5.0 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60.5k | 3.0k | 1.9k | 6.1k | Go | 5.0 |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60.3k | 3.6k | 2.5k | 206.7k | TypeScript | 5.0 |
| [kubernetes/legacy-cloud-providers](https://github.com/kubernetes/legacy-cloud-providers/)  | 60.2k | 1.8k | 0.0 | 47.0 | Go | 5.0 |
| [openyurtio/openyurt](https://github.com/openyurtio/openyurt/)  | 60.2k | 850.0 | 885.0 | 1.4k | Go | 5.0 |
| [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/)  | 60.1k | 2.5k | 30.0 | 661.0 | Go | 5.0 |
| [getporter/gh-action](https://github.com/getporter/gh-action/)  | 60.0k | 27.0 | 12.0 | 1.0 | JavaScript | 5.0 |
| [awslabs/aws-config-rules](https://github.com/awslabs/aws-config-rules/)  | 59.4k | 919.0 | 252.0 | 1.4k | Python | 5.0 |
| [ydb-platform/ydb-nodejs-genproto](https://github.com/ydb-platform/ydb-nodejs-genproto/)  | 59.3k | 9.0 | 3.0 | 1.0 | JavaScript | 5.0 |
| [apache/nutch](https://github.com/apache/nutch/)  | 58.7k | 3.4k | 748.0 | 2.6k | Java | 4.9 |
| [apache/maven-plugin-tools](https://github.com/apache/maven-plugin-tools/)  | 58.4k | 1.2k | 210.0 | 45.0 | JavaScript | 4.9 |
| [buildpacks/pack](https://github.com/buildpacks/pack/)  | 58.4k | 3.4k | 957.0 | 1.9k | Go | 4.9 |
| [awslabs/aws-elasticache-cluster-cl~](https://github.com/awslabs/aws-elasticache-cluster-client-libmemcached/)  | 57.8k | 59.0 | 10.0 | 22.0 | C++ | 4.8 |
| [envoyproxy/envoy-mobile](https://github.com/envoyproxy/envoy-mobile/)  | 57.6k | 1.8k | 2.2k | 550.0 | Java | 4.8 |
| [keycloak/terraform-provider-keyclo~](https://github.com/keycloak/terraform-provider-keycloak/)  | 57.5k | 581.0 | 526.0 | 649.0 | Go | 4.8 |
| [inclavare-containers/Shelter](https://github.com/inclavare-containers/Shelter/)  | 57.3k | 192.0 | 12.0 | 0.0 | Rust | 4.8 |
| [apache/servicemix3](https://github.com/apache/servicemix3/)  | 57.1k | 3.0k | 22.0 | 6.0 | Java | 4.8 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9k | 2.3k | 2.8k | 2.3k | C# | 4.7 |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56.9k | 6.9k | 4.9k | 14.5k | Go | 4.7 |
| [apache/roller](https://github.com/apache/roller/)  | 56.7k | 4.7k | 117.0 | 115.0 | Java | 4.7 |
| [apache/maven-shared](https://github.com/apache/maven-shared/)  | 56.5k | 3.4k | 26.0 | 29.0 | Java | 4.7 |
| [kanisterio/kanister](https://github.com/kanisterio/kanister/)  | 56.3k | 2.4k | 2.3k | 705.0 | Go | 4.7 |
| [googlecontainertools/jib](https://github.com/googlecontainertools/jib/)  | 56.3k | 2.3k | 2.6k | 13.3k | Java | 4.7 |
| [open-telemetry/otel-arrow-collector](https://github.com/open-telemetry/otel-arrow-collector/)  | 56.2k | 4.7k | 44.0 | 20.0 | Go | 4.7 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 56.1k | 1.8k | 1.1k | 286.0 | Go | 4.7 |
| [dragonflyoss/nydus](https://github.com/dragonflyoss/nydus/)  | 55.9k | 3.1k | 1.2k | 908.0 | Rust | 4.7 |
| [ariga/atlas](https://github.com/ariga/atlas/)  | 55.7k | 2.2k | 2.2k | 4.8k | Go | 4.6 |
| [nats-io/nats.net.v1](https://github.com/nats-io/nats.net.v1/)  | 55.7k | 853.0 | 548.0 | 647.0 | C# | 4.6 |
| [volcano-sh/volcano](https://github.com/volcano-sh/volcano/)  | 55.7k | 4.6k | 1.6k | 2.9k | Go | 4.6 |
| [apache/harmony-jdktools](https://github.com/apache/harmony-jdktools/)  | 55.1k | 273.0 | 0.0 | 4.0 | Java | 4.6 |
| [prestodb/benchto](https://github.com/prestodb/benchto/)  | 54.6k | 356.0 | 47.0 | 23.0 | JavaScript | 4.5 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 54.2k | 1.9k | 1.1k | 457.0 | Rust | 4.5 |
| [moby/libnetwork](https://github.com/moby/libnetwork/)  | 54.1k | 3.1k | 1.9k | 2.0k | Go | 4.5 |
| [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy/)  | 54.1k | 2.2k | 2.3k | 1.7k | Rust | 4.5 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53.4k | 4.4k | 37.0 | 11.0 | Go | 4.5 |
| [dragonflyoss/image-service](https://github.com/dragonflyoss/image-service/)  | 53.2k | 3.0k | 1.1k | 795.0 | Rust | 4.4 |
| [gotd/td](https://github.com/gotd/td/)  | 52.6k | 3.1k | 788.0 | 1.1k | Go | 4.4 |
| [kcl-lang/kcl-py](https://github.com/kcl-lang/kcl-py/)  | 52.6k | 92.0 | 20.0 | 8.0 | Python | 4.4 |
| [serverlessworkflow/synapse](https://github.com/serverlessworkflow/synapse/)  | 52.6k | 814.0 | 204.0 | 145.0 | C# | 4.4 |
| [inclavare-containers/spdm-rs](https://github.com/inclavare-containers/spdm-rs/)  | 52.6k | 28.0 | 2.0 | 0.0 | Rust | 4.4 |
| [parallaxsecond/rust-cryptoki](https://github.com/parallaxsecond/rust-cryptoki/)  | 52.6k | 252.0 | 94.0 | 30.0 | Rust | 4.4 |
| [rkt/rkt](https://github.com/rkt/rkt/)  | 51.9k | 5.6k | 2.5k | 8.8k | Go | 4.3 |
| [telepresenceio/telepresence](https://github.com/telepresenceio/telepresence/)  | 51.8k | 7.3k | 1.5k | 5.7k | Go | 4.3 |
| [brigadecore/brigade](https://github.com/brigadecore/brigade/)  | 51.5k | 2.0k | 1.2k | 2.4k | Go | 4.3 |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 1.6k | 990.0 | 420.0 | TypeScript | 4.3 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51.2k | 3.2k | 8.0 | 206.0 | Go | 4.3 |
| [envoyproxy/toolshed](https://github.com/envoyproxy/toolshed/)  | 51.2k | 651.0 | 658.0 | 6.0 | Python | 4.3 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 1.9k | 2.4k | 3.7k | Go | 4.3 |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50.8k | 2.5k | 999.0 | 233.0 | Go | 4.2 |
| [coredns/coredns](https://github.com/coredns/coredns/)  | 50.7k | 3.6k | 3.8k | 10.4k | Go | 4.2 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50.7k | 233.0 | 194.0 | 9.0 | Go | 4.2 |
| [keycloak/keycloak-client](https://github.com/keycloak/keycloak-client/)  | 50.5k | 65.0 | 60.0 | 8.0 | Java | 4.2 |
| [trickstercache/trickster](https://github.com/trickstercache/trickster/)  | 50.4k | 927.0 | 438.0 | 1.9k | Go | 4.2 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector/)  | 50.3k | 4.4k | 5.1k | 2.8k | Go | 4.2 |
| [nats-io/nats.java](https://github.com/nats-io/nats.java/)  | 49.2k | 1.5k | 510.0 | 461.0 | Java | 4.1 |
| [getporter/porter](https://github.com/getporter/porter/)  | 49.1k | 2.9k | 1.5k | 969.0 | Go | 4.1 |
| [apache/commons-io](https://github.com/apache/commons-io/)  | 48.9k | 4.1k | 454.0 | 908.0 | Java | 4.1 |
| [kubernetes-sigs/multi-tenancy](https://github.com/kubernetes-sigs/multi-tenancy/)  | 48.9k | 2.3k | 1.1k | 930.0 | Go | 4.1 |
| [envoyproxy/pytooling](https://github.com/envoyproxy/pytooling/)  | 48.9k | 615.0 | 619.0 | 6.0 | Python | 4.1 |
| [jaegertracing/jaeger-ui](https://github.com/jaegertracing/jaeger-ui/)  | 48.9k | 1.1k | 781.0 | 894.0 | JavaScript | 4.1 |
| [nats-io/nats-streaming-server](https://github.com/nats-io/nats-streaming-server/)  | 48.9k | 1.7k | 651.0 | 2.4k | Go | 4.1 |
| [external-secrets/external-secrets](https://github.com/external-secrets/external-secrets/)  | 48.4k | 2.2k | 1.3k | 2.5k | Go | 4.0 |
| [grpc/grpc-swift](https://github.com/grpc/grpc-swift/)  | 48.3k | 1.6k | 981.0 | 1.7k | Swift | 4.0 |
| [cilium/tetragon](https://github.com/cilium/tetragon/)  | 47.5k | 1.8k | 1.1k | 2.5k | Go | 4.0 |
| [cri-o/cri-o](https://github.com/cri-o/cri-o/)  | 47.4k | 7.6k | 5.5k | 4.4k | Go | 4.0 |
| [slimtoolkit/slim](https://github.com/slimtoolkit/slim/)  | 47.4k | 1.1k | 300.0 | 18.1k | Go | 4.0 |
| [apache/tapestry3](https://github.com/apache/tapestry3/)  | 47.2k | 2.0k | 6.0 | 5.0 | Java | 3.9 |
| [paralus/dashboard](https://github.com/paralus/dashboard/)  | 46.8k | 299.0 | 96.0 | 9.0 | JavaScript | 3.9 |
| [nats-io/nats.net](https://github.com/nats-io/nats.net/)  | 46.8k | 747.0 | 430.0 | 566.0 | C# | 3.9 |
| [operator-framework/operator-sdk](https://github.com/operator-framework/operator-sdk/)  | 46.7k | 3.2k | 3.9k | 6.3k | Go | 3.9 |
| [knative/client](https://github.com/knative/client/)  | 46.6k | 1.1k | 1.3k | 303.0 | Go | 3.9 |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46.6k | 4.5k | 5.0 | 1.0 | Go | 3.9 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46.4k | 241.0 | 358.0 | 20.0 | C | 3.9 |
| [vitessio/website](https://github.com/vitessio/website/)  | 46.0k | 2.9k | 1.1k | 36.0 | JavaScript | 3.8 |
| [networkservicemesh/sdk](https://github.com/networkservicemesh/sdk/)  | 45.9k | 932.0 | 1.1k | 25.0 | Go | 3.8 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45.8k | 790.0 | 309.0 | 38.0 | JavaScript | 3.8 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.8k | 1.8k | 1.0k | 413.0 | Python | 3.8 |
| [headlamp-k8s/headlamp](https://github.com/headlamp-k8s/headlamp/)  | 45.4k | 3.5k | 1.1k | 1.3k | TypeScript | 3.8 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45.3k | 1.1k | 543.0 | 113.0 | Go | 3.8 |
| [kubeflow/kubeflow](https://github.com/kubeflow/kubeflow/)  | 45.1k | 2.5k | 3.6k | 13.3k | TypeScript | 3.8 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44.9k | 4.4k | 4.9k | 12.3k | Go | 3.7 |
| [cockroachdb/gostdlib](https://github.com/cockroachdb/gostdlib/)  | 44.7k | 10.0 | 5.0 | 2.0 | Go | 3.7 |
| [kubeflow/arena](https://github.com/kubeflow/arena/)  | 44.2k | 436.0 | 880.0 | 678.0 | Go | 3.7 |
| [apache/servicemix4-bundles](https://github.com/apache/servicemix4-bundles/)  | 43.9k | 7.6k | 347.0 | 18.0 | Java | 3.7 |
| [notaryproject/notary](https://github.com/notaryproject/notary/)  | 43.8k | 3.0k | 992.0 | 3.0k | Go | 3.6 |
| [skooner-k8s/skooner](https://github.com/skooner-k8s/skooner/)  | 43.8k | 406.0 | 213.0 | 1.1k | JavaScript | 3.6 |
| [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper/)  | 43.7k | 1.2k | 1.5k | 3.0k | Go | 3.6 |
| [containerd/nerdctl](https://github.com/containerd/nerdctl/)  | 43.6k | 2.4k | 1.3k | 5.7k | Go | 3.6 |
| [tauri-apps/tauri](https://github.com/tauri-apps/tauri/)  | 43.4k | 4.5k | 4.3k | 76.7k | Rust | 3.6 |
| [inclavare-containers/inclavare-con~](https://github.com/inclavare-containers/inclavare-containers/)  | 43.1k | 1.4k | 923.0 | 532.0 | C | 3.6 |
| [operator-framework/java-operator-s~](https://github.com/operator-framework/java-operator-sdk/)  | 42.9k | 2.4k | 1.3k | 595.0 | Java | 3.6 |
| [distribution/distribution](https://github.com/distribution/distribution/)  | 42.9k | 4.8k | 1.8k | 7.4k | Go | 3.6 |
| [tikv/client-go](https://github.com/tikv/client-go/)  | 42.6k | 474.0 | 598.0 | 227.0 | Go | 3.5 |
| [open-telemetry/opentelemetry-netwo~](https://github.com/open-telemetry/opentelemetry-network/)  | 42.6k | 380.0 | 207.0 | 181.0 | C++ | 3.5 |
| [istio/old_mixer_repo](https://github.com/istio/old_mixer_repo/)  | 42.5k | 741.0 | 1.1k | 68.0 | Go | 3.5 |
| [crossplane/crossplane](https://github.com/crossplane/crossplane/)  | 42.5k | 4.6k | 2.2k | 6.7k | Go | 3.5 |
| [longhorn/longhorn-tests](https://github.com/longhorn/longhorn-tests/)  | 42.4k | 1.5k | 1.3k | 13.0 | Python | 3.5 |
| [grpc/grpc-node](https://github.com/grpc/grpc-node/)  | 42.3k | 4.2k | 1.4k | 3.9k | TypeScript | 3.5 |
| [carvel-dev/ytt](https://github.com/carvel-dev/ytt/)  | 42.3k | 1.2k | 298.0 | 1.3k | Go | 3.5 |
| [open-telemetry/opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python/)  | 42.2k | 1.4k | 1.7k | 1.2k | Python | 3.5 |
| [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)  | 42.1k | 2.2k | 1.3k | 1.9k | Go | 3.5 |
| [open-telemetry/opentelemetry-cpp-c~](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42.1k | 157.0 | 180.0 | 78.0 | Python | 3.5 |
| [helm/helm](https://github.com/helm/helm/)  | 41.7k | 6.8k | 4.8k | 23.9k | Go | 3.5 |
| [open-telemetry/opentelemetry-ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4k | 268.0 | 112.0 | 85.0 | C++ | 3.5 |
| [networkservicemesh/networkservicem~](https://github.com/networkservicemesh/networkservicemesh/)  | 41.1k | 1.4k | 1.7k | 509.0 | Go | 3.4 |
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40.9k | 1.7k | 14.0 | 105.0 | C++ | 3.4 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 40.8k | 52.0 | 3.0 | 5.0 | JavaScript | 3.4 |
| [keptn/tutorials](https://github.com/keptn/tutorials/)  | 40.8k | 622.0 | 193.0 | 10.0 | JavaScript | 3.4 |
| [spidernet-io/spiderpool](https://github.com/spidernet-io/spiderpool/)  | 40.6k | 3.1k | 2.3k | 401.0 | Go | 3.4 |
| [open-telemetry/otel-arrow](https://github.com/open-telemetry/otel-arrow/)  | 40.6k | 339.0 | 0.0 | 0.0 | Go | 3.4 |
| [apache/sanselan](https://github.com/apache/sanselan/)  | 40.5k | 315.0 | 3.0 | 37.0 | Java | 3.4 |
| [capsule-rs/capsule](https://github.com/capsule-rs/capsule/)  | 40.4k | 229.0 | 117.0 | 347.0 | Rust | 3.4 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40.4k | 11.1k | 4.8k | 952.0 | TypeScript | 3.4 |
| [googlecontainertools/kpt-functions~](https://github.com/googlecontainertools/kpt-functions-catalog/)  | 40.3k | 711.0 | 1.1k | 83.0 | TypeScript | 3.4 |
| [dapr/java-sdk](https://github.com/dapr/java-sdk/)  | 39.7k | 409.0 | 490.0 | 223.0 | Java | 3.3 |
| [ytsaurus/ytsaurus-spyt](https://github.com/ytsaurus/ytsaurus-spyt/)  | 39.4k | 327.0 | 28.0 | 12.0 | Scala | 3.3 |
| [tikv/client-java](https://github.com/tikv/client-java/)  | 39.4k | 287.0 | 534.0 | 91.0 | Java | 3.3 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39.1k | 177.0 | 142.0 | 13.0 | Go | 3.3 |
| [ydb-platform/yoj-project](https://github.com/ydb-platform/yoj-project/)  | 39.1k | 22.0 | 5.0 | 4.0 | Java | 3.2 |
| [nats-io/nats.go](https://github.com/nats-io/nats.go/)  | 38.8k | 1.8k | 716.0 | 4.4k | Go | 3.2 |
| [spectralops/spectral-goat](https://github.com/spectralops/spectral-goat/)  | 38.3k | 3.0 | 0.0 | 2.0 | JavaScript | 3.2 |
| [vscode-kubernetes-tools/vscode-kub~](https://github.com/vscode-kubernetes-tools/vscode-kubernetes-tools/)  | 38.0k | 593.0 | 527.0 | 587.0 | TypeScript | 3.2 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38.0k | 3.0k | 1.8k | 6.2k | Go | 3.2 |
| [metallb/metallb](https://github.com/metallb/metallb/)  | 37.4k | 1.8k | 958.0 | 5.7k | Go | 3.1 |
| [uber/kraken](https://github.com/uber/kraken/)  | 37.4k | 870.0 | 244.0 | 5.8k | Go | 3.1 |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37.3k | 799.0 | 283.0 | 321.0 | Go | 3.1 |
| [jquery/jquery](https://github.com/jquery/jquery/)  | 37.2k | 6.7k | 3.0k | 58.8k | JavaScript | 3.1 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36.7k | 3.3k | 700.0 | 454.0 | Go | 3.0 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36.4k | 35.0 | 7.0 | 1.0 | C | 3.0 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36.3k | 298.0 | 358.0 | 83.0 | Go | 3.0 |
| [apache/xmlgraphics-commons](https://github.com/apache/xmlgraphics-commons/)  | 36.2k | 625.0 | 16.0 | 18.0 | Java | 3.0 |
| [xline-kv/Xline](https://github.com/xline-kv/Xline/)  | 36.1k | 1.0k | 548.0 | 510.0 | Rust | 3.0 |
| [gravity-ui/markdown-editor](https://github.com/gravity-ui/markdown-editor/)  | 36.1k | 483.0 | 470.0 | 180.0 | TypeScript | 3.0 |
| [microcks/microcks](https://github.com/microcks/microcks/)  | 36.0k | 1.4k | 131.0 | 1.1k | Java | 3.0 |
| [argoproj/argo-events](https://github.com/argoproj/argo-events/)  | 35.9k | 1.3k | 1.5k | 1.9k | Go | 3.0 |
| [awslabs/aws-ms-deploy-assistant](https://github.com/awslabs/aws-ms-deploy-assistant/)  | 35.8k | 14.0 | 6.0 | 8.0 | JavaScript | 3.0 |
| [metrico/qryn-view](https://github.com/metrico/qryn-view/)  | 35.7k | 1.0k | 165.0 | 22.0 | TypeScript | 3.0 |
| [paralus/paralus](https://github.com/paralus/paralus/)  | 35.7k | 572.0 | 98.0 | 718.0 | Go | 3.0 |
| [apache/logging-log4j1](https://github.com/apache/logging-log4j1/)  | 35.6k | 3.2k | 22.0 | 877.0 | Java | 3.0 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35.5k | 662.0 | 313.0 | 278.0 | Go | 3.0 |
| [openebs/openebs-e2e](https://github.com/openebs/openebs-e2e/)  | 35.4k | 197.0 | 93.0 | 0.0 | Go | 3.0 |
| [ydb-platform/tpcc-postgres](https://github.com/ydb-platform/tpcc-postgres/)  | 35.3k | 2.2k | 21.0 | 1.0 | Java | 2.9 |
| [buildpacks/lifecycle-private](https://github.com/buildpacks/lifecycle-private/)  | 35.2k | 1.9k | 25.0 | 0.0 | Go | 2.9 |
| [netdata/dashboard](https://github.com/netdata/dashboard/)  | 35.0k | 1.8k | 394.0 | 63.0 | JavaScript | 2.9 |
| [cubefs/cubefs-dashboard](https://github.com/cubefs/cubefs-dashboard/)  | 34.8k | 8.0 | 5.0 | 4.0 | Go | 2.9 |
| [opencurve/curve-tgt](https://github.com/opencurve/curve-tgt/)  | 34.8k | 2.1k | 0.0 | 18.0 | C | 2.9 |
| [kubedl-io/kubedl](https://github.com/kubedl-io/kubedl/)  | 34.4k | 377.0 | 180.0 | 417.0 | Go | 2.9 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34.2k | 1.3k | 462.0 | 2.3k | Go | 2.9 |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34.1k | 13.0 | 0.0 | 6.0 | JavaScript | 2.8 |
| [kubeflow/katib](https://github.com/kubeflow/katib/)  | 34.0k | 1.2k | 1.4k | 1.4k | Go | 2.8 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33.7k | 1.5k | 420.0 | 78.0 | Go | 2.8 |
| [kcp-dev/contrib-tmc](https://github.com/kcp-dev/contrib-tmc/)  | 33.7k | 51.0 | 17.0 | 5.0 | Go | 2.8 |
| [containerssh/libcontainerssh](https://github.com/containerssh/libcontainerssh/)  | 33.7k | 324.0 | 384.0 | 28.0 | Go | 2.8 |
| [apache/buildr](https://github.com/apache/buildr/)  | 33.6k | 4.4k | 52.0 | 142.0 | Ruby | 2.8 |
| [carvel-dev/kapp-controller](https://github.com/carvel-dev/kapp-controller/)  | 33.6k | 1.4k | 704.0 | 160.0 | Go | 2.8 |
| [open-telemetry/opentelemetry-cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.6k | 1.0k | 1.2k | 432.0 | C++ | 2.8 |
| [kubeedge/sedna](https://github.com/kubeedge/sedna/)  | 33.4k | 475.0 | 215.0 | 428.0 | Python | 2.8 |
| [kubernetes/perf-tests](https://github.com/kubernetes/perf-tests/)  | 33.4k | 3.1k | 1.9k | 764.0 | Go | 2.8 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.4k | 732.0 | 850.0 | 230.0 | C# | 2.8 |
| [opencontainers/runc](https://github.com/opencontainers/runc/)  | 33.3k | 6.3k | 2.7k | 10.1k | Go | 2.8 |
| [in-toto/github-action](https://github.com/in-toto/github-action/)  | 33.2k | 3.0 | 0.0 | 4.0 | JavaScript | 2.8 |
| [awslabs/amazon-kinesis-client](https://github.com/awslabs/amazon-kinesis-client/)  | 33.2k | 466.0 | 646.0 | 608.0 | Java | 2.8 |
| [open-telemetry/opentelemetry-ebpf-~](https://github.com/open-telemetry/opentelemetry-ebpf-profiler/)  | 33.1k | 142.0 | 178.0 | 2.5k | Go | 2.8 |
| [meshery/meshery.io](https://github.com/meshery/meshery.io/)  | 33.1k | 2.6k | 646.0 | 188.0 | JavaScript | 2.8 |
| [buildpacks/lifecycle](https://github.com/buildpacks/lifecycle/)  | 32.9k | 1.6k | 665.0 | 158.0 | Go | 2.7 |
| [googlecontainertools/kpt-functions~](https://github.com/googlecontainertools/kpt-functions-sdk/)  | 32.7k | 558.0 | 668.0 | 53.0 | TypeScript | 2.7 |
| [operator-framework/operator-regist~](https://github.com/operator-framework/operator-registry/)  | 32.6k | 1.1k | 780.0 | 181.0 | Go | 2.7 |
| [goharbor/harbor-operator](https://github.com/goharbor/harbor-operator/)  | 32.5k | 1.0k | 537.0 | 293.0 | Go | 2.7 |
| [keptn/lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit/)  | 32.2k | 497.0 | 666.0 | 56.0 | Go | 2.7 |
| [cloudevents/sdk-go](https://github.com/cloudevents/sdk-go/)  | 32.2k | 624.0 | 542.0 | 613.0 | Go | 2.7 |
| [nats-io/nsc](https://github.com/nats-io/nsc/)  | 31.9k | 604.0 | 358.0 | 61.0 | Go | 2.7 |
| [prometheus/alertmanager](https://github.com/prometheus/alertmanager/)  | 31.9k | 2.8k | 1.7k | 5.5k | Go | 2.7 |
| [ydb-platform/xorm](https://github.com/ydb-platform/xorm/)  | 31.9k | 1.7k | 9.0 | 0.0 | Go | 2.7 |
| [litmuschaos/test-tools](https://github.com/litmuschaos/test-tools/)  | 31.8k | 226.0 | 300.0 | 27.0 | C | 2.6 |
| [clusternet/clusternet](https://github.com/clusternet/clusternet/)  | 31.6k | 672.0 | 449.0 | 1.1k | Go | 2.6 |
| [kubevirt/hyperconverged-cluster-op~](https://github.com/kubevirt/hyperconverged-cluster-operator/)  | 31.4k | 1.4k | 2.2k | 106.0 | Go | 2.6 |
| [open-cluster-management-io/multicl~](https://github.com/open-cluster-management-io/multicloud-operators-subscription/)  | 31.2k | 801.0 | 298.0 | 32.0 | Go | 2.6 |
| [knative/func](https://github.com/knative/func/)  | 31.2k | 1.2k | 1.1k | 156.0 | Go | 2.6 |
| [pravega/pravega-client-rust](https://github.com/pravega/pravega-client-rust/)  | 31.1k | 173.0 | 202.0 | 27.0 | Rust | 2.6 |
| [cockroachdb/docs](https://github.com/cockroachdb/docs/)  | 31.0k | 14.0k | 6.3k | 169.0 | Java | 2.6 |
| [falcosecurity/plugin-sdk-rs](https://github.com/falcosecurity/plugin-sdk-rs/)  | 31.0k | 296.0 | 24.0 | 6.0 | Rust | 2.6 |
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 508.0 | 591.0 | 519.0 | PHP | 2.6 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30.9k | 1.3k | 308.0 | 232.0 | Go | 2.6 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30.9k | 18.0 | 4.0 | 0.0 | C Header | 2.6 |
| [konveyor/move2kube-demos](https://github.com/konveyor/move2kube-demos/)  | 30.9k | 116.0 | 240.0 | 9.0 | JavaScript | 2.6 |
| [sealerio/sealer](https://github.com/sealerio/sealer/)  | 30.8k | 988.0 | 1.2k | 1.8k | Go | 2.6 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 574.0 | 597.0 | 997.0 | Rust | 2.5 |
| [oras-project/oras-go](https://github.com/oras-project/oras-go/)  | 30.5k | 366.0 | 252.0 | 113.0 | Go | 2.5 |
| [superedge/superedge](https://github.com/superedge/superedge/)  | 30.2k | 378.0 | 375.0 | 903.0 | Go | 2.5 |
| [fluxcd/source-controller](https://github.com/fluxcd/source-controller/)  | 30.1k | 1.9k | 759.0 | 191.0 | Go | 2.5 |
| [kube-logging/logging-operator](https://github.com/kube-logging/logging-operator/)  | 29.9k | 2.0k | 1.1k | 1.3k | Go | 2.5 |
| [docker/machine](https://github.com/docker/machine/)  | 29.9k | 3.5k | 2.0k | 6.6k | Go | 2.5 |
| [cilium/ebpf](https://github.com/cilium/ebpf/)  | 29.9k | 1.3k | 686.0 | 4.2k | Go | 2.5 |
| [metrico/opensearch-action](https://github.com/metrico/opensearch-action/)  | 29.9k | 3.0 | 0.0 | 3.0 | JavaScript | 2.5 |
| [fluxcd/flux](https://github.com/fluxcd/flux/)  | 29.9k | 5.3k | 1.9k | 6.9k | Go | 2.5 |
| [go-faster/yt](https://github.com/go-faster/yt/)  | 29.8k | 9.0 | 6.0 | 1.0 | Go | 2.5 |
| [konveyor/move2kube](https://github.com/konveyor/move2kube/)  | 29.7k | 702.0 | 760.0 | 287.0 | Go | 2.5 |
| [nats-io/nats.deno](https://github.com/nats-io/nats.deno/)  | 29.7k | 442.0 | 413.0 | 115.0 | TypeScript | 2.5 |
| [devfile/library](https://github.com/devfile/library/)  | 29.7k | 530.0 | 155.0 | 16.0 | Go | 2.5 |
| [m3db/m3cluster](https://github.com/m3db/m3cluster/)  | 29.6k | 238.0 | 227.0 | 21.0 | Go | 2.5 |
| [kubevirt/project-infra](https://github.com/kubevirt/project-infra/)  | 29.5k | 2.4k | 2.5k | 21.0 | Go | 2.5 |
| [moby/hyperkit](https://github.com/moby/hyperkit/)  | 29.4k | 623.0 | 220.0 | 3.5k | C | 2.5 |
| [dexidp/dex](https://github.com/dexidp/dex/)  | 29.3k | 2.4k | 1.7k | 7.9k | Go | 2.4 |
| [k3s-io/k3s](https://github.com/k3s-io/k3s/)  | 29.3k | 2.6k | 2.6k | 22.5k | Go | 2.4 |
| [ydb-platform/ycsb-ydb](https://github.com/ydb-platform/ycsb-ydb/)  | 29.2k | 1.4k | 0.0 | 0.0 | Java | 2.4 |
| [cubefs/compass](https://github.com/cubefs/compass/)  | 29.1k | 86.0 | 5.0 | 171.0 | Java | 2.4 |
| [apache/maven-mercury](https://github.com/apache/maven-mercury/)  | 29.0k | 769.0 | 0.0 | 4.0 | Java | 2.4 |
| [dragonflyoss/Dragonfly](https://github.com/dragonflyoss/Dragonfly/)  | 28.8k | 1.8k | 1.0k | 6.0k | Go | 2.4 |
| [openkruise/rollouts](https://github.com/openkruise/rollouts/)  | 28.6k | 58.0 | 88.0 | 107.0 | Go | 2.4 |
| [kubevirt/web-ui-components](https://github.com/kubevirt/web-ui-components/)  | 28.5k | 416.0 | 500.0 | 9.0 | JavaScript | 2.4 |
| [litmuschaos/litmus-go](https://github.com/litmuschaos/litmus-go/)  | 28.5k | 425.0 | 558.0 | 51.0 | Go | 2.4 |
| [jaegertracing/jaeger-operator](https://github.com/jaegertracing/jaeger-operator/)  | 28.5k | 1.1k | 1.3k | 904.0 | Go | 2.4 |
| [cncf/devstatscode](https://github.com/cncf/devstatscode/)  | 28.3k | 405.0 | 13.0 | 30.0 | Go | 2.4 |
| [chaosblade-io/chaosblade-box-fe](https://github.com/chaosblade-io/chaosblade-box-fe/)  | 28.2k | 10.0 | 32.0 | 5.0 | TypeScript | 2.4 |
| [open-telemetry/opentelemetry-log-c~](https://github.com/open-telemetry/opentelemetry-log-collection/)  | 28.1k | 267.0 | 373.0 | 92.0 | Go | 2.3 |
| [projectriff/cli](https://github.com/projectriff/cli/)  | 28.0k | 1.5k | 263.0 | 6.0 | Go | 2.3 |
| [fnproject/fn](https://github.com/fnproject/fn/)  | 28.0k | 3.4k | 953.0 | 5.4k | Go | 2.3 |
| [envoyproxy/nighthawk](https://github.com/envoyproxy/nighthawk/)  | 27.9k | 578.0 | 655.0 | 272.0 | C++ | 2.3 |
| [metal3-io/cluster-api-provider-met~](https://github.com/metal3-io/cluster-api-provider-metal3/)  | 27.8k | 1.7k | 801.0 | 138.0 | Go | 2.3 |
| [kubeflow/training-operator](https://github.com/kubeflow/training-operator/)  | 27.6k | 950.0 | 1.1k | 1.4k | Go | 2.3 |
| [kubernetes-sigs/alibaba-cloud-csi-~](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/)  | 27.6k | 1.3k | 590.0 | 444.0 | Go | 2.3 |
| [fluxcd/flagger](https://github.com/fluxcd/flagger/)  | 27.6k | 2.7k | 753.0 | 4.1k | Go | 2.3 |
| [fluxcd/pkg](https://github.com/fluxcd/pkg/)  | 27.5k | 1.1k | 454.0 | 30.0 | Go | 2.3 |
| [kubernetes/release](https://github.com/kubernetes/release/)  | 27.3k | 5.2k | 2.4k | 436.0 | Go | 2.3 |
| [devfile/integration-tests](https://github.com/devfile/integration-tests/)  | 27.3k | 108.0 | 95.0 | 1.0 | Go | 2.3 |
| [docker/docker-py](https://github.com/docker/docker-py/)  | 27.3k | 3.3k | 1.5k | 6.1k | Python | 2.3 |
| [open-telemetry/opentelemetry-go-co~](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.3k | 1.0k | 3.2k | 684.0 | Go | 2.3 |
| [awslabs/lambda-chef-node-cleanup](https://github.com/awslabs/lambda-chef-node-cleanup/)  | 27.2k | 86.0 | 11.0 | 72.0 | Python | 2.3 |
| [devstream-io/devstream](https://github.com/devstream-io/devstream/)  | 27.1k | 2.2k | 954.0 | 736.0 | Go | 2.3 |
| [konveyor/tackle2-ui](https://github.com/konveyor/tackle2-ui/)  | 27.1k | 574.0 | 675.0 | 3.0 | TypeScript | 2.3 |
| [antrea-io/nephe](https://github.com/antrea-io/nephe/)  | 27.1k | 124.0 | 144.0 | 8.0 | Go | 2.3 |
| [kudobuilder/kudo](https://github.com/kudobuilder/kudo/)  | 26.8k | 1.0k | 1.1k | 1.1k | Go | 2.2 |
| [cilium/hubble-ui](https://github.com/cilium/hubble-ui/)  | 26.8k | 385.0 | 426.0 | 237.0 | TypeScript | 2.2 |
| [ClickHouse/libhdfs3](https://github.com/ClickHouse/libhdfs3/)  | 26.7k | 64.0 | 34.0 | 23.0 | C++ | 2.2 |
| [ClickHouse/clickhouse-cpp](https://github.com/ClickHouse/clickhouse-cpp/)  | 26.7k | 673.0 | 172.0 | 215.0 | C++ | 2.2 |
| [kubevirt/vm-import-operator](https://github.com/kubevirt/vm-import-operator/)  | 26.6k | 900.0 | 415.0 | 13.0 | Go | 2.2 |
| [dapr/dotnet-sdk](https://github.com/dapr/dotnet-sdk/)  | 26.5k | 550.0 | 465.0 | 967.0 | C# | 2.2 |
| [apache/chainsaw](https://github.com/apache/chainsaw/)  | 26.4k | 934.0 | 2.0 | 18.0 | Java | 2.2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26.3k | 562.0 | 644.0 | 80.0 | Java | 2.2 |
| [gravity-ui/charts](https://github.com/gravity-ui/charts/)  | 26.3k | 14.0 | 13.0 | 2.0 | JavaScript | 2.2 |
| [containernetworking/plugins](https://github.com/containernetworking/plugins/)  | 26.2k | 1.6k | 484.0 | 1.7k | Go | 2.2 |
| [open-telemetry/opentelemetry-swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9k | 618.0 | 252.0 | 126.0 | Swift | 2.2 |
| [cockroachdb/cockroach-operator](https://github.com/cockroachdb/cockroach-operator/)  | 25.8k | 589.0 | 461.0 | 216.0 | Go | 2.1 |
| [openfga/openfga](https://github.com/openfga/openfga/)  | 25.8k | 471.0 | 536.0 | 843.0 | Go | 2.1 |
| [carvel-dev/kapp](https://github.com/carvel-dev/kapp/)  | 25.7k | 873.0 | 403.0 | 734.0 | Go | 2.1 |
| [longhorn/longhorn-ui](https://github.com/longhorn/longhorn-ui/)  | 25.6k | 817.0 | 595.0 | 92.0 | JavaScript | 2.1 |
| [pravega/www.pravega.io](https://github.com/pravega/www.pravega.io/)  | 25.6k | 13.0 | 0.0 | 0.0 | JavaScript | 2.1 |
| [netdata/libjudy](https://github.com/netdata/libjudy/)  | 25.5k | 25.0 | 3.0 | 14.0 | C | 2.1 |
| [envoyproxy/java-control-plane](https://github.com/envoyproxy/java-control-plane/)  | 25.5k | 235.0 | 206.0 | 246.0 | Protocol Buffers | 2.1 |
| [apache/maven-scm](https://github.com/apache/maven-scm/)  | 25.4k | 2.2k | 175.0 | 88.0 | Java | 2.1 |
| [grafana/grafana-plugin-sdk-go](https://github.com/grafana/grafana-plugin-sdk-go/)  | 25.4k | 416.0 | 531.0 | 157.0 | Go | 2.1 |
| [kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)  | 25.3k | 642.0 | 349.0 | 733.0 | Go | 2.1 |
| [antrea-io/theia](https://github.com/antrea-io/theia/)  | 25.1k | 123.0 | 161.0 | 25.0 | Go | 2.1 |
| [containerd/stargz-snapshotter](https://github.com/containerd/stargz-snapshotter/)  | 25.1k | 1.4k | 1.0k | 796.0 | Go | 2.1 |
| [gravity-ui/uikit](https://github.com/gravity-ui/uikit/)  | 25.1k | 481.0 | 501.0 | 171.0 | TypeScript | 2.1 |
| [kubernetes-sigs/kubefed](https://github.com/kubernetes-sigs/kubefed/)  | 25.0k | 2.7k | 978.0 | 2.5k | Go | 2.1 |


### CNCF Projects
| Project | SLOC | Commits | PRs | Stars | Language | Effort |
|---------|------|---------|-----|-------|----------|--------|
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1k | 190.5k | 243.1k | Go | 287.0 |
| [envoyproxy](https://github.com/envoyproxy)  | 2.9M | 52.5k | 25.1k | 28.1k | JavaScript | 243.5 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.0M | 121.4k | 72.3k | 98.8k | Go | 167.7 |
| [open-telemetry](https://github.com/open-telemetry)  | 2.0M | 69.7k | 61.8k | 31.5k | Go | 165.4 |
| [grpc](https://github.com/grpc)  | 1.7M | 75.0k | 38.9k | 86.8k | C++ | 143.8 |
| [konveyor](https://github.com/konveyor)  | 1.6M | 8.1k | 7.7k | 858.0 | JavaScript | 136.6 |
| [inclavare-containers](https://github.com/inclavare-containers)  | 1.6M | 25.1k | 1.1k | 571.0 | C++ | 130.7 |
| [openkruise](https://github.com/openkruise)  | 1.4M | 1.3k | 1.3k | 4.0k | JavaScript | 113.7 |
| [kubewarden](https://github.com/kubewarden)  | 1.1M | 8.8k | 3.1k | 494.0 | JavaScript | 92.6 |
| [kubeflow](https://github.com/kubeflow)  | 938.8k | 20.8k | 24.2k | 27.2k | Go | 78.2 |
| [brigadecore](https://github.com/brigadecore)  | 877.2k | 6.3k | 2.8k | 3.0k | JavaScript | 73.1 |
| [tikv](https://github.com/tikv)  | 862.8k | 17.2k | 21.3k | 25.2k | Rust | 71.9 |
| [nats-io](https://github.com/nats-io)  | 862.6k | 37.8k | 13.8k | 33.3k | Go | 71.9 |
| [keycloak](https://github.com/keycloak)  | 842.3k | 34.8k | 22.5k | 25.4k | Java | 70.2 |
| [cdk8s-team](https://github.com/cdk8s-team)  | 804.0k | 9.3k | 7.9k | 3.8k | Go | 67.0 |
| [cilium](https://github.com/cilium)  | 734.5k | 34.6k | 23.4k | 26.3k | Go | 61.2 |
| [kubevirt](https://github.com/kubevirt)  | 683.4k | 61.0k | 22.7k | 5.5k | Go | 57.0 |
| [vitessio](https://github.com/vitessio)  | 677.6k | 37.2k | 11.1k | 15.9k | Go | 56.5 |
| [cubefs](https://github.com/cubefs)  | 659.3k | 3.1k | 1.5k | 3.6k | Go | 54.9 |
| [fluent](https://github.com/fluent)  | 655.2k | 38.1k | 12.0k | 24.2k | C | 54.6 |
| [falcosecurity](https://github.com/falcosecurity)  | 625.9k | 22.9k | 7.9k | 8.6k | C Header | 52.2 |
| [istio](https://github.com/istio)  | 583.6k | 47.1k | 62.4k | 38.5k | Go | 48.6 |
| [openebs](https://github.com/openebs)  | 560.6k | 16.1k | 9.7k | 10.1k | Go | 46.7 |
| [argoproj](https://github.com/argoproj)  | 517.8k | 18.1k | 16.9k | 34.2k | Go | 43.1 |
| [cncf](https://github.com/cncf)  | 497.9k | 68.0k | 16.3k | 28.6k | JavaScript | 41.5 |
| [googlecontainertools](https://github.com/googlecontainertools)  | 493.2k | 21.6k | 34.6k | 66.7k | Go | 41.1 |
| [pravega](https://github.com/pravega)  | 474.4k | 6.4k | 5.5k | 2.5k | Java | 39.5 |
| [kubevela](https://github.com/kubevela)  | 448.4k | 8.3k | 6.8k | 5.6k | Go | 37.4 |
| [athenz](https://github.com/athenz)  | 433.0k | 3.8k | 2.2k | 808.0 | Java | 36.1 |
| [knative](https://github.com/knative)  | 432.8k | 30.6k | 32.3k | 12.4k | Go | 36.1 |
| [opencurve](https://github.com/opencurve)  | 419.9k | 5.7k | 1.9k | 1.9k | C++ | 35.0 |
| [backstage](https://github.com/backstage)  | 409.3k | 40.6k | 13.0k | 21.8k | TypeScript | 34.1 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762.0 | Java | 33.8 |
| [prometheus](https://github.com/prometheus)  | 379.9k | 31.2k | 19.3k | 91.9k | Go | 31.7 |
| [pixie-io](https://github.com/pixie-io)  | 378.2k | 13.4k | 2.1k | 5.5k | C++ | 31.5 |
| [kubeovn](https://github.com/kubeovn)  | 376.4k | 21.7k | 2.0k | 1.5k | Go | 31.4 |
| [werf](https://github.com/werf)  | 358.3k | 16.4k | 5.8k | 4.8k | JavaScript | 29.9 |
| [containerd](https://github.com/containerd)  | 358.1k | 23.6k | 11.7k | 24.8k | Go | 29.8 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.0k | 8.5k | 7.2k | 15.4k | Go | 28.6 |
| [dapr](https://github.com/dapr)  | 340.1k | 18.1k | 11.0k | 26.1k | Go | 28.4 |
| [operator-framework](https://github.com/operator-framework)  | 326.8k | 21.6k | 15.5k | 14.0k | Go | 27.2 |
| [chaosblade-io](https://github.com/chaosblade-io)  | 310.9k | 1.3k | 979.0 | 6.2k | JavaScript | 25.9 |
| [kyverno](https://github.com/kyverno)  | 300.4k | 9.4k | 5.6k | 4.0k | Go | 25.0 |
| [parallaxsecond](https://github.com/parallaxsecond)  | 292.2k | 3.7k | 1.5k | 553.0 | Rust | 24.4 |
| [open-cluster-management-io](https://github.com/open-cluster-management-io)  | 289.8k | 7.0k | 3.5k | 1.9k | Go | 24.1 |
| [confidential-containers](https://github.com/confidential-containers)  | 282.4k | 23.7k | 1.8k | 425.0 | Go | 23.5 |
| [antrea-io](https://github.com/antrea-io)  | 262.8k | 3.5k | 3.9k | 1.5k | Go | 21.9 |
| [keptn](https://github.com/keptn)  | 256.8k | 11.8k | 8.4k | 1.9k | Go | 21.4 |
| [inspektor-gadget](https://github.com/inspektor-gadget)  | 253.6k | 3.7k | 1.0k | 1.3k | C Header | 21.1 |
| [strimzi](https://github.com/strimzi)  | 253.2k | 7.1k | 6.0k | 4.3k | Java | 21.1 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.5k | 8.7k | 7.6k | 23.7k | Go | 20.9 |
| [goharbor](https://github.com/goharbor)  | 246.8k | 17.0k | 10.4k | 21.1k | Go | 20.6 |
| [kumahq](https://github.com/kumahq)  | 240.3k | 7.2k | 7.1k | 3.3k | Go | 20.0 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 13.0k | 13.4k | 18.0k | Rust | 19.8 |
| [buildpacks](https://github.com/buildpacks)  | 225.2k | 22.5k | 3.6k | 2.8k | Go | 18.8 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 224.7k | 11.1k | 6.7k | 9.8k | Rust | 18.7 |
| [spiffe](https://github.com/spiffe)  | 216.5k | 9.4k | 4.1k | 3.0k | Go | 18.0 |
| [litmuschaos](https://github.com/litmuschaos)  | 205.9k | 7.5k | 7.1k | 4.1k | TypeScript | 17.1 |
| [kcl-lang](https://github.com/kcl-lang)  | 201.3k | 3.9k | 1.5k | 1.2k | Rust | 16.8 |
| [fluxcd](https://github.com/fluxcd)  | 200.2k | 27.8k | 9.5k | 19.9k | Go | 16.7 |
| [longhorn](https://github.com/longhorn)  | 199.6k | 10.3k | 6.6k | 5.3k | Go | 16.6 |
| [megaease](https://github.com/megaease)  | 198.9k | 4.5k | 1.9k | 11.5k | Go | 16.6 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4k | 4.5k | 4.6k | 4.8k | Python | 16.4 |
| [tremor-rs](https://github.com/tremor-rs)  | 193.9k | 7.4k | 3.6k | 837.0 | Rust | 16.2 |
| [etcd-io](https://github.com/etcd-io)  | 187.8k | 26.5k | 11.3k | 52.3k | Go | 15.7 |
| [carvel-dev](https://github.com/carvel-dev)  | 181.2k | 8.5k | 3.3k | 3.5k | Go | 15.1 |
| [cert-manager](https://github.com/cert-manager)  | 163.7k | 15.7k | 5.9k | 11.0k | Go | 13.6 |
| [wasmcloud](https://github.com/wasmcloud)  | 162.0k | 5.7k | 2.4k | 1.9k | Rust | 13.5 |
| [meshery](https://github.com/meshery)  | 155.9k | 30.4k | 9.2k | 2.9k | JavaScript | 13.0 |
| [crossplane](https://github.com/crossplane)  | 154.1k | 12.8k | 3.8k | 7.9k | Go | 12.8 |
| [networkservicemesh](https://github.com/networkservicemesh)  | 152.2k | 37.4k | 30.5k | 660.0 | Go | 12.7 |
| [kubeedge](https://github.com/kubeedge)  | 150.1k | 7.3k | 3.9k | 6.8k | Go | 12.5 |
| [thanos-io](https://github.com/thanos-io)  | 144.6k | 4.3k | 4.2k | 12.1k | Go | 12.1 |
| [devfile](https://github.com/devfile)  | 138.4k | 4.7k | 2.2k | 322.0 | Go | 11.5 |
| [getporter](https://github.com/getporter)  | 136.8k | 6.2k | 2.5k | 1.1k | Go | 11.4 |
| [fluid-cloudnative](https://github.com/fluid-cloudnative)  | 136.8k | 2.0k | 2.0k | 1.3k | Go | 11.4 |
| [wasmedge](https://github.com/wasmedge)  | 132.5k | 3.4k | 1.5k | 5.8k | Rust | 11.1 |
| [rook](https://github.com/rook)  | 132.1k | 24.8k | 7.1k | 11.2k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 4.9k | 3.8k | 4.1k | Go | 10.9 |
| [armadaproject](https://github.com/armadaproject)  | 128.9k | 2.2k | 1.8k | 286.0 | Go | 10.7 |
| [keylime](https://github.com/keylime)  | 127.6k | 2.4k | 1.5k | 399.0 | C | 10.6 |
| [nocalhost](https://github.com/nocalhost)  | 123.0k | 6.1k | 2.1k | 1.6k | Go | 10.2 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 4.9k | 3.8k | 5.1k | Go | 10.0 |
| [kcp-dev](https://github.com/kcp-dev)  | 117.7k | 6.5k | 2.5k | 2.2k | Go | 9.8 |
| [artifacthub](https://github.com/artifacthub)  | 116.7k | 1.8k | 2.3k | 1.2k | TypeScript | 9.7 |
| [openelb](https://github.com/openelb)  | 116.4k | 498.0 | 283.0 | 1.3k | JavaScript | 9.7 |
| [kubescape](https://github.com/kubescape)  | 115.3k | 12.5k | 2.0k | 8.4k | Go | 9.6 |
| [serverless-devs](https://github.com/serverless-devs)  | 115.1k | 3.5k | 460.0 | 1.5k | JavaScript | 9.6 |
| [paralus](https://github.com/paralus)  | 112.6k | 1.6k | 437.0 | 758.0 | Go | 9.4 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 111.1k | 19.1k | 3.3k | 4.0k | Python | 9.3 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 107.2k | 3.3k | 3.2k | 5.9k | Go | 8.9 |
| [in-toto](https://github.com/in-toto)  | 106.7k | 10.2k | 1.4k | 1.5k | Python | 8.9 |
| [kedacore](https://github.com/kedacore)  | 105.4k | 3.9k | 4.2k | 6.8k | Go | 8.8 |
| [pipe-cd](https://github.com/pipe-cd)  | 104.6k | 4.8k | 3.8k | 1.0k | Go | 8.7 |
| [karmada-io](https://github.com/karmada-io)  | 100.5k | 4.1k | 2.7k | 3.2k | Go | 8.4 |
| [kubeclipper](https://github.com/kubeclipper)  | 99.5k | 575.0 | 829.0 | 368.0 | Go | 8.3 |
| [openyurtio](https://github.com/openyurtio)  | 98.6k | 1.9k | 1.5k | 1.6k | Go | 8.2 |
| [bfenetworks](https://github.com/bfenetworks)  | 95.8k | 1.4k | 825.0 | 6.0k | Go | 8.0 |
| [v6d-io](https://github.com/v6d-io)  | 94.8k | 1.0k | 882.0 | 689.0 | C++ | 7.9 |
| [spidernet-io](https://github.com/spidernet-io)  | 88.5k | 5.9k | 4.4k | 582.0 | Go | 7.4 |
| [serverlessworkflow](https://github.com/serverlessworkflow)  | 88.4k | 2.4k | 1.2k | 951.0 | C# | 7.4 |
| [openservicemesh](https://github.com/openservicemesh)  | 81.9k | 5.2k | 4.1k | 2.6k | Go | 6.8 |
| [hwameistor](https://github.com/hwameistor)  | 80.4k | 2.8k | 1.2k | 474.0 | Go | 6.7 |
| [cloudevents](https://github.com/cloudevents)  | 79.5k | 4.0k | 2.6k | 5.8k | Go | 6.6 |
| [coredns](https://github.com/coredns)  | 76.2k | 5.5k | 4.7k | 11.5k | Go | 6.3 |
| [project-zot](https://github.com/project-zot)  | 75.8k | 1.1k | 1.2k | 333.0 | Go | 6.3 |
| [opencost](https://github.com/opencost)  | 72.8k | 3.9k | 1.5k | 3.6k | Go | 6.1 |
| [notaryproject](https://github.com/notaryproject)  | 72.3k | 3.8k | 2.1k | 3.4k | Go | 6.0 |
| [helm](https://github.com/helm)  | 71.5k | 28.9k | 26.9k | 48.8k | Go | 6.0 |
| [cri-o](https://github.com/cri-o)  | 71.5k | 9.2k | 5.8k | 4.5k | Go | 6.0 |
| [metal3-io](https://github.com/metal3-io)  | 71.0k | 9.5k | 4.6k | 1.0k | Go | 5.9 |
| [volcano-sh](https://github.com/volcano-sh)  | 70.0k | 5.5k | 2.0k | 3.2k | Go | 5.8 |
| [submariner-io](https://github.com/submariner-io)  | 69.8k | 10.0k | 9.1k | 2.3k | Go | 5.8 |
| [oras-project](https://github.com/oras-project)  | 66.1k | 3.6k | 1.2k | 1.2k | Go | 5.5 |
| [microcks](https://github.com/microcks)  | 65.1k | 2.8k | 368.0 | 1.3k | Java | 5.4 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 1.8k | 2.2k | 2.2k | Go | 5.2 |
| [open-feature](https://github.com/open-feature)  | 60.8k | 5.6k | 3.5k | 800.0 | Go | 5.1 |
| [openfga](https://github.com/openfga)  | 60.2k | 1.8k | 1.5k | 1.0k | Go | 5.0 |
| [kanisterio](https://github.com/kanisterio)  | 59.8k | 2.5k | 2.3k | 707.0 | Go | 5.0 |
| [telepresenceio](https://github.com/telepresenceio)  | 58.0k | 8.1k | 1.7k | 5.7k | Go | 4.8 |
| [external-secrets](https://github.com/external-secrets)  | 57.3k | 2.8k | 1.8k | 5.1k | Go | 4.8 |
| [rkt](https://github.com/rkt)  | 52.4k | 5.6k | 2.5k | 8.9k | Go | 4.4 |
| [trickstercache](https://github.com/trickstercache)  | 52.1k | 1.1k | 458.0 | 1.9k | Go | 4.3 |
| [cloud-bulldozer](https://github.com/cloud-bulldozer)  | 51.9k | 7.9k | 3.6k | 914.0 | Python | 4.3 |
| [headlamp-k8s](https://github.com/headlamp-k8s)  | 51.4k | 3.7k | 1.1k | 1.3k | TypeScript | 4.3 |
| [theupdateframework](https://github.com/theupdateframework)  | 50.3k | 9.1k | 2.8k | 2.6k | Python | 4.2 |
| [slimtoolkit](https://github.com/slimtoolkit)  | 49.7k | 1.3k | 359.0 | 18.3k | Go | 4.1 |
| [superedge](https://github.com/superedge)  | 48.2k | 487.0 | 425.0 | 913.0 | Go | 4.0 |
| [virtual-kubelet](https://github.com/virtual-kubelet)  | 48.1k | 2.8k | 1.3k | 4.4k | Go | 4.0 |
| [kudobuilder](https://github.com/kudobuilder)  | 47.9k | 2.2k | 2.0k | 1.8k | Go | 4.0 |
| [xline-kv](https://github.com/xline-kv)  | 47.9k | 1.3k | 723.0 | 539.0 | Rust | 4.0 |
| [distribution](https://github.com/distribution)  | 45.7k | 5.1k | 1.9k | 7.6k | Go | 3.8 |
| [k3s-io](https://github.com/k3s-io)  | 45.0k | 3.6k | 3.3k | 25.5k | Go | 3.8 |
| [openfunction](https://github.com/openfunction)  | 44.5k | 1.5k | 915.0 | 1.2k | Go | 3.7 |
| [tinkerbell](https://github.com/tinkerbell)  | 44.3k | 8.3k | 2.9k | 1.9k | Go | 3.7 |
| [kubedl-io](https://github.com/kubedl-io)  | 44.1k | 1.3k | 369.0 | 488.0 | Go | 3.7 |
| [skooner-k8s](https://github.com/skooner-k8s)  | 44.0k | 507.0 | 213.0 | 1.1k | JavaScript | 3.7 |
| [metallb](https://github.com/metallb)  | 40.9k | 2.3k | 1.3k | 5.7k | Go | 3.4 |
| [aeraki-mesh](https://github.com/aeraki-mesh)  | 40.5k | 992.0 | 368.0 | 868.0 | Go | 3.4 |
| [kubearmor](https://github.com/kubearmor)  | 40.3k | 5.7k | 2.2k | 733.0 | Go | 3.4 |
| [devstream-io](https://github.com/devstream-io)  | 39.3k | 2.9k | 1.1k | 747.0 | Go | 3.3 |
| [vscode-kubernetes-tools](https://github.com/vscode-kubernetes-tools)  | 39.3k | 634.0 | 613.0 | 602.0 | TypeScript | 3.3 |
| [kube-logging](https://github.com/kube-logging)  | 35.9k | 2.9k | 1.4k | 1.4k | Go | 3.0 |
| [kuasar-io](https://github.com/kuasar-io)  | 35.7k | 156.0 | 78.0 | 972.0 | Rust | 3.0 |
| [sustainable-computing-io](https://github.com/sustainable-computing-io)  | 35.5k | 4.6k | 1.7k | 944.0 | Go | 3.0 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 2.8k | 1.2k | 6.5k | Go | 2.9 |
| [clusternet](https://github.com/clusternet)  | 33.4k | 948.0 | 545.0 | 1.1k | Go | 2.8 |
| [sealerio](https://github.com/sealerio)  | 32.7k | 1.2k | 1.3k | 1.8k | Go | 2.7 |
| [kubestellar](https://github.com/kubestellar)  | 32.6k | 4.4k | 1.1k | 215.0 | Go | 2.7 |
| [clusterpedia-io](https://github.com/clusterpedia-io)  | 31.6k | 1.1k | 541.0 | 597.0 | Go | 2.6 |
| [kube-rs](https://github.com/kube-rs)  | 31.6k | 3.0k | 775.0 | 2.2k | Rust | 2.6 |
| [dexidp](https://github.com/dexidp)  | 29.3k | 2.6k | 1.8k | 8.0k | Go | 2.4 |
| [piraeusdatastore](https://github.com/piraeusdatastore)  | 28.7k | 1.4k | 522.0 | 748.0 | Go | 2.4 |
| [opentracing](https://github.com/opentracing)  | 27.4k | 2.7k | 1.4k | 10.1k | Java | 2.3 |
| [krustlet](https://github.com/krustlet)  | 23.6k | 2.4k | 530.0 | 3.4k | Rust | 2.0 |
| [curiefense](https://github.com/curiefense)  | 22.3k | 2.9k | 817.0 | 617.0 | Rust | 1.9 |
| [redhat-chaos](https://github.com/redhat-chaos)  | 22.3k | 911.0 | 784.0 | 317.0 | Python | 1.9 |
| [projectcapsule](https://github.com/projectcapsule)  | 21.9k | 1.2k | 744.0 | 1.4k | Go | 1.8 |
| [fabedge](https://github.com/fabedge)  | 21.6k | 941.0 | 429.0 | 482.0 | Go | 1.8 |
| [project-akri](https://github.com/project-akri)  | 20.6k | 702.0 | 448.0 | 946.0 | Rust | 1.7 |
| [k8gb-io](https://github.com/k8gb-io)  | 19.4k | 1.1k | 918.0 | 571.0 | Go | 1.6 |
| [foniod](https://github.com/foniod)  | 18.1k | 1.7k | 493.0 | 2.0k | Rust | 1.5 |
| [lima-vm](https://github.com/lima-vm)  | 17.4k | 2.1k | 934.0 | 11.7k | Go | 1.4 |
| [schemahero](https://github.com/schemahero)  | 17.0k | 1.4k | 773.0 | 765.0 | Go | 1.4 |
| [hexa-org](https://github.com/hexa-org)  | 16.5k | 491.0 | 118.0 | 70.0 | Go | 1.4 |
| [getsops](https://github.com/getsops)  | 15.4k | 1.6k | 643.0 | 14.5k | Go | 1.3 |
| [k8up-io](https://github.com/k8up-io)  | 14.1k | 2.0k | 716.0 | 412.0 | Go | 1.2 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 749.0 | 118.0 | 588.0 | Go | 1.1 |
| [servicemeshinterface](https://github.com/servicemeshinterface)  | 11.6k | 966.0 | 372.0 | 1.2k | Go | 970.0m |
| [kube-vip](https://github.com/kube-vip)  | 10.9k | 918.0 | 320.0 | 1.3k | Go | 910.0m |
| [kuberhealthy](https://github.com/kuberhealthy)  | 10.8k | 3.1k | 721.0 | 1.7k | Go | 900.0m |
| [k8sgpt-ai](https://github.com/k8sgpt-ai)  | 10.7k | 1.5k | 989.0 | 3.9k | Go | 890.0m |
| [ko-build](https://github.com/ko-build)  | 9.5k | 793.0 | 721.0 | 5.8k | Go | 790.0m |
| [eraser-dev](https://github.com/eraser-dev)  | 9.4k | 587.0 | 728.0 | 330.0 | Go | 780.0m |
| [devspace-cloud](https://github.com/devspace-cloud)  | 8.7k | 348.0 | 83.0 | 113.0 | Go | 720.0m |
| [tellerops](https://github.com/tellerops)  | 7.7k | 179.0 | 80.0 | 1.6k | Go | 640.0m |
| [cni-genie](https://github.com/cni-genie)  | 5.9k | 484.0 | 132.0 | 499.0 | Go | 490.0m |
| [project-copacetic](https://github.com/project-copacetic)  | 5.4k | 483.0 | 475.0 | 675.0 | Go | 450.0m |
| [merbridge](https://github.com/merbridge)  | 4.8k | 271.0 | 279.0 | 586.0 | Go | 400.0m |
| [opcr-io](https://github.com/opcr-io)  | 3.6k | 597.0 | 153.0 | 181.0 | Go | 300.0m |
| [krator-rs](https://github.com/krator-rs)  | 2.9k | 498.0 | 33.0 | 113.0 | Rust | 250.0m |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [openobservability](https://github.com/openobservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [service-mesh-performance](https://github.com/service-mesh-performance)  | 2.2k | 797.0 | 217.0 | 237.0 | JavaScript | 180.0m |
| [kubereboot](https://github.com/kubereboot)  | 2.0k | 1.0k | 558.0 | 1.8k | Go | 170.0m |


### OTEL Instrumentations
| Project | SLOC | Commits | PRs | Stars | Effort |
|---------|------|---------|-----|-------|--------|
| [**OTEL**](https://opentelemetry.io/docs/instrumentation/)  | 1.2M | 34.4k | 31.4k | 19.5k | 96.78 |
| [java-instrumentation](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.9k | 8.8k | 5.1k | 1.2k | 13.99 |
| [dotnet-instrumentation](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.1k | 929.0 | 1.8k | 224.0 | 11.76 |
| [java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 3.3k | 3.6k | 1.5k | 8.03 |
| [js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 1.7k | 2.0k | 1.8k | 5.08 |
| [dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9k | 2.3k | 2.8k | 2.3k | 4.74 |
| [js-contrib](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 1.6k | 990.0 | 420.0 | 4.28 |
| [go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 1.9k | 2.4k | 3.7k | 4.26 |
| [python-contrib](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.8k | 1.8k | 1.0k | 413.0 | 3.82 |
| [network](https://github.com/open-telemetry/opentelemetry-network/)  | 42.6k | 380.0 | 207.0 | 181.0 | 3.55 |
| [python](https://github.com/open-telemetry/opentelemetry-python/)  | 42.2k | 1.4k | 1.7k | 1.2k | 3.52 |
| [cpp-contrib](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42.1k | 157.0 | 180.0 | 78.0 | 3.5 |
| [ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4k | 268.0 | 112.0 | 85.0 | 3.45 |
| [cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.6k | 1.0k | 1.2k | 432.0 | 2.8 |
| [dotnet-contrib](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.4k | 732.0 | 850.0 | 230.0 | 2.78 |
| [ebpf-profiler](https://github.com/open-telemetry/opentelemetry-ebpf-profiler/)  | 33.1k | 142.0 | 178.0 | 2.5k | 2.76 |
| [php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 508.0 | 591.0 | 519.0 | 2.58 |
| [rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 574.0 | 597.0 | 997.0 | 2.54 |
| [go-contrib](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.3k | 1.0k | 3.2k | 684.0 | 2.27 |
| [java-contrib](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26.3k | 562.0 | 644.0 | 80.0 | 2.2 |
| [swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9k | 618.0 | 252.0 | 126.0 | 2.16 |
| [ruby-contrib](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 20.0k | 895.0 | 241.0 | 29.0 | 1.66 |
| [ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16.4k | 763.0 | 919.0 | 375.0 | 1.37 |
| [php-contrib](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12.8k | 322.0 | 112.0 | 22.0 | 1.06 |
| [erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11.5k | 1.1k | 355.0 | 253.0 | 0.96 |
| [android](https://github.com/open-telemetry/opentelemetry-android/)  | 7.1k | 532.0 | 6.0 | 14.0 | 0.6 |
| [rust-contrib](https://github.com/open-telemetry/opentelemetry-rust-contrib/)  | 6.4k | 797.0 | 4.0 | 3.0 | 0.54 |
| [go-instrumentation](https://github.com/open-telemetry/opentelemetry-go-instrumentation/)  | 2.9k | 23.0 | 23.0 | 54.0 | 0.24 |
| [erlang-contrib](https://github.com/open-telemetry/opentelemetry-erlang-contrib/)  | 1.5k | 96.0 | 117.0 | 73.0 | 0.12 |
| [java-examples](https://github.com/open-telemetry/opentelemetry-java-examples/)  | 1.3k | 123.0 | 124.0 | 90.0 | 0.11 |
| [php-instrumentation](https://github.com/open-telemetry/opentelemetry-php-instrumentation/)  | 556.0 | 30.0 | 30.0 | 24.0 | 0.05 |
| [configuration](https://github.com/open-telemetry/opentelemetry-configuration/)  | 0.0 | 9.0 | 8.0 | 7.0 | 0 |

