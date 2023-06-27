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
| [**CNCF**](https://www.cncf.io/)  | 39.5M | 1.9M | 1.1M | 1.5M | Go | 3.3k |
| [zephyrproject-rtos](https://github.com/zephyrproject-rtos)  | 20.3M | 269.8k | 42.3k | 8.6k | C Header | 1.7k |
| [torvalds](https://github.com/torvalds)  | 19.1M | 1.2M | 761.0 | 153.7k | C | 1.6k |
| [apache](https://github.com/apache)  | 15.7M | 609.5k | 44.2k | 88.2k | Java | 1.3k |
| [ydb-platform](https://github.com/ydb-platform)  | 6.4M | 19.6k | 1.6k | 3.2k | C++ | 530.4 |
| [**K8s**](https://kubernetes.io/)  | 5.4M | 499.3k | 261.8k | 341.6k | Go | 453.2 |
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1k | 190.5k | 243.1k | Go | 286.4 |
| [tensorflow](https://github.com/tensorflow)  | 3.4M | 149.9k | 23.1k | 175.7k | C++ | 279.2 |
| [elastic](https://github.com/elastic)  | 3.1M | 96.3k | 99.7k | 89.5k | Java | 261.0 |
| [ytsaurus](https://github.com/ytsaurus)  | 2.6M | 65.5k | 9.0 | 818.0 | C++ | 216.1 |
| [godotengine](https://github.com/godotengine)  | 2.1M | 71.2k | 42.7k | 79.6k | C++ | 171.0 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.0M | 118.2k | 71.3k | 98.5k | Go | 166.8 |
| [cockroachdb](https://github.com/cockroachdb)  | 2.0M | 100.8k | 59.3k | 33.7k | Go | 163.1 |
| [open-telemetry](https://github.com/open-telemetry)  | 1.8M | 62.5k | 61.2k | 28.7k | Go | 148.7 |
| [grpc](https://github.com/grpc)  | 1.7M | 74.1k | 38.8k | 86.8k | C++ | 140.3 |
| [ClickHouse](https://github.com/ClickHouse)  | 1.7M | 119.4k | 34.9k | 34.0k | C++ | 140.1 |
| [grafana](https://github.com/grafana)  | 1.7M | 76.6k | 50.2k | 103.2k | Go | 139.9 |
| [envoyproxy](https://github.com/envoyproxy)  | 1.6M | 47.0k | 24.7k | 28.1k | C++ | 135.3 |
| [konveyor](https://github.com/konveyor)  | 1.6M | 7.1k | 6.3k | 806.0 | JavaScript | 131.0 |
| [rust-lang](https://github.com/rust-lang)  | 1.5M | 227.8k | 63.9k | 83.0k | Rust | 126.5 |
| [python](https://github.com/python)  | 1.4M | 117.6k | 41.8k | 54.2k | Python | 114.4 |
| [openkruise](https://github.com/openkruise)  | 1.4M | 1.3k | 1.2k | 4.0k | JavaScript | 113.5 |
| [docker](https://github.com/docker)  | 1.3M | 176.8k | 30.6k | 146.4k | Go | 106.0 |
| [netdata](https://github.com/netdata)  | 1.2M | 25.8k | 12.1k | 63.1k | JavaScript | 103.4 |
| [prestodb](https://github.com/prestodb)  | 1.2M | 26.0k | 15.3k | 15.7k | Java | 97.0 |
| [golang](https://github.com/golang)  | 1.1M | 57.1k | 3.0k | 112.4k | Go | 88.4 |
| [nodejs](https://github.com/nodejs)  | 966.7k | 39.2k | 30.7k | 96.2k | JavaScript | 80.6 |
| [ziglang](https://github.com/ziglang)  | 888.6k | 24.8k | 7.5k | 23.5k | Zig | 74.0 |
| [brigadecore](https://github.com/brigadecore)  | 877.2k | 6.3k | 2.8k | 3.0k | JavaScript | 73.1 |
| [tikv](https://github.com/tikv)  | 860.8k | 16.8k | 20.1k | 24.1k | Rust | 71.7 |
| [cdk8s-team](https://github.com/cdk8s-team)  | 803.7k | 9.3k | 7.9k | 3.8k | Go | 67.0 |
| [nats-io](https://github.com/nats-io)  | 802.5k | 36.2k | 12.9k | 31.0k | Go | 66.9 |
| [m3db](https://github.com/m3db)  | 736.8k | 10.4k | 5.1k | 4.6k | Go | 61.4 |
| [cilium](https://github.com/cilium)  | 681.7k | 31.0k | 21.9k | 23.6k | Go | 56.8 |
| [vitessio](https://github.com/vitessio)  | 671.3k | 36.7k | 11.0k | 15.9k | Go | 56.0 |
| [moby](https://github.com/moby)  | 648.2k | 74.7k | 31.3k | 83.2k | Go | 54.0 |
| [fluent](https://github.com/fluent)  | 639.2k | 38.0k | 12.0k | 24.2k | C | 53.3 |
| [kubevirt](https://github.com/kubevirt)  | 637.7k | 60.6k | 22.6k | 5.5k | Go | 53.1 |
| [cubefs](https://github.com/cubefs)  | 623.9k | 3.0k | 1.5k | 3.6k | Go | 52.0 |
| [istio](https://github.com/istio)  | 583.6k | 47.1k | 62.4k | 38.5k | Go | 48.6 |
| [falcosecurity](https://github.com/falcosecurity)  | 578.3k | 20.8k | 6.8k | 7.6k | C Header | 48.2 |
| [openebs](https://github.com/openebs)  | 524.7k | 15.5k | 9.6k | 10.1k | Go | 43.7 |
| [argoproj](https://github.com/argoproj)  | 517.8k | 18.1k | 16.9k | 34.2k | Go | 43.1 |
| [denoland](https://github.com/denoland)  | 498.2k | 9.6k | 10.7k | 89.9k | JavaScript | 41.5 |
| [cncf](https://github.com/cncf)  | 481.4k | 53.9k | 14.9k | 28.0k | JavaScript | 40.1 |
| [pravega](https://github.com/pravega)  | 472.7k | 6.4k | 5.5k | 2.5k | Java | 39.4 |
| [kubevela](https://github.com/kubevela)  | 448.4k | 8.3k | 6.8k | 5.6k | Go | 37.4 |
| [knative](https://github.com/knative)  | 432.8k | 30.6k | 32.3k | 12.4k | Go | 36.1 |
| [awslabs](https://github.com/awslabs)  | 427.5k | 11.1k | 4.5k | 45.6k | Java | 35.6 |
| [athenz](https://github.com/athenz)  | 408.2k | 3.5k | 2.1k | 771.0 | Java | 34.0 |
| [opencurve](https://github.com/opencurve)  | 406.3k | 5.6k | 1.9k | 1.9k | C++ | 33.9 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762.0 | Java | 33.8 |
| [backstage](https://github.com/backstage)  | 403.5k | 40.5k | 12.9k | 21.8k | TypeScript | 33.6 |
| [pixie-io](https://github.com/pixie-io)  | 378.2k | 13.2k | 1.9k | 4.9k | C++ | 31.5 |
| [prometheus](https://github.com/prometheus)  | 377.3k | 30.9k | 19.1k | 89.5k | Go | 31.4 |
| [django](https://github.com/django)  | 365.8k | 31.8k | 17.0k | 71.7k | Python | 30.5 |
| [facebook](https://github.com/facebook)  | 357.5k | 15.9k | 13.7k | 209.6k | JavaScript | 29.8 |
| [containerd](https://github.com/containerd)  | 353.4k | 23.2k | 11.7k | 24.7k | Go | 29.4 |
| [werf](https://github.com/werf)  | 349.4k | 16.3k | 5.8k | 4.7k | JavaScript | 29.1 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.0k | 8.5k | 7.2k | 15.4k | Go | 28.6 |
| [inclavare-containers](https://github.com/inclavare-containers)  | 341.1k | 3.3k | 1.1k | 569.0 | C Header | 28.4 |
| [dapr](https://github.com/dapr)  | 335.5k | 17.4k | 10.7k | 26.0k | Go | 28.0 |
| [operator-framework](https://github.com/operator-framework)  | 312.6k | 21.6k | 15.5k | 14.0k | Go | 26.1 |
| [chaosblade-io](https://github.com/chaosblade-io)  | 310.9k | 1.3k | 979.0 | 6.2k | JavaScript | 25.9 |
| [LINBIT](https://github.com/LINBIT)  | 310.9k | 4.5k | 12.0 | 701.0 | Java | 25.9 |
| [vectordotdev](https://github.com/vectordotdev)  | 309.0k | 14.4k | 10.2k | 14.4k | Rust | 25.8 |
| [parallaxsecond](https://github.com/parallaxsecond)  | 289.8k | 3.5k | 1.4k | 552.0 | Rust | 24.1 |
| [kyverno](https://github.com/kyverno)  | 281.4k | 8.9k | 5.3k | 4.0k | Go | 23.4 |
| [open-cluster-management-io](https://github.com/open-cluster-management-io)  | 276.0k | 7.0k | 3.5k | 1.9k | Go | 23.0 |
| [antrea-io](https://github.com/antrea-io)  | 262.8k | 3.5k | 3.9k | 1.5k | Go | 21.9 |
| [keptn](https://github.com/keptn)  | 256.8k | 11.8k | 8.4k | 1.9k | Go | 21.4 |
| [uber](https://github.com/uber)  | 253.8k | 1.6k | 249.0 | 6.1k | Go | 21.1 |
| [inspektor-gadget](https://github.com/inspektor-gadget)  | 253.5k | 2.9k | 1.0k | 1.3k | C Header | 21.1 |
| [strimzi](https://github.com/strimzi)  | 252.3k | 7.1k | 6.0k | 4.3k | Java | 21.0 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.5k | 8.7k | 7.6k | 23.7k | Go | 20.9 |
| [siderolabs](https://github.com/siderolabs)  | 243.1k | 7.1k | 9.1k | 4.8k | Go | 20.3 |
| [goharbor](https://github.com/goharbor)  | 242.5k | 16.9k | 10.4k | 21.1k | Go | 20.2 |
| [kumahq](https://github.com/kumahq)  | 240.3k | 7.2k | 7.1k | 3.3k | Go | 20.0 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 13.0k | 13.4k | 18.0k | Rust | 19.8 |
| [matplotlib](https://github.com/matplotlib)  | 235.9k | 47.5k | 16.6k | 17.6k | Python | 19.6 |
| [kata-containers](https://github.com/kata-containers)  | 225.2k | 22.8k | 10.3k | 7.0k | Go | 18.8 |
| [Netflix](https://github.com/Netflix)  | 208.1k | 4.2k | 2.3k | 549.0 | Java | 17.4 |
| [spiffe](https://github.com/spiffe)  | 206.9k | 8.6k | 4.0k | 2.9k | Go | 17.2 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 206.0k | 8.1k | 2.6k | 11.0k | Go | 17.2 |
| [alibaba](https://github.com/alibaba)  | 205.2k | 4.1k | 859.0 | 27.2k | Java | 17.1 |
| [litmuschaos](https://github.com/litmuschaos)  | 203.2k | 7.5k | 7.1k | 4.1k | TypeScript | 16.9 |
| [confidential-containers](https://github.com/confidential-containers)  | 201.9k | 18.7k | 1.5k | 383.0 | Rust | 16.8 |
| [fluxcd](https://github.com/fluxcd)  | 200.2k | 27.7k | 9.5k | 19.9k | Go | 16.7 |
| [metrico](https://github.com/metrico)  | 199.2k | 5.3k | 519.0 | 819.0 | Go | 16.6 |
| [longhorn](https://github.com/longhorn)  | 197.7k | 10.0k | 6.6k | 5.1k | Go | 16.5 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4k | 4.5k | 4.6k | 4.8k | Python | 16.4 |
| [tremor-rs](https://github.com/tremor-rs)  | 193.6k | 7.3k | 3.6k | 837.0 | Rust | 16.1 |
| [etcd-io](https://github.com/etcd-io)  | 185.3k | 26.4k | 11.3k | 52.3k | Go | 15.4 |
| [carvel-dev](https://github.com/carvel-dev)  | 180.0k | 8.4k | 3.3k | 3.4k | Go | 15.0 |
| [containers](https://github.com/containers)  | 173.6k | 19.3k | 10.8k | 18.3k | Go | 14.5 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 168.8k | 7.5k | 4.9k | 8.3k | Go | 14.1 |
| [meshery](https://github.com/meshery)  | 154.9k | 30.0k | 9.0k | 2.8k | JavaScript | 12.9 |
| [networkservicemesh](https://github.com/networkservicemesh)  | 150.4k | 37.4k | 30.5k | 660.0 | Go | 12.5 |
| [cert-manager](https://github.com/cert-manager)  | 146.2k | 13.1k | 4.8k | 10.9k | Go | 12.2 |
| [wasmcloud](https://github.com/wasmcloud)  | 145.4k | 4.9k | 2.0k | 1.8k | Rust | 12.1 |
| [thanos-io](https://github.com/thanos-io)  | 144.3k | 4.3k | 4.2k | 12.1k | Go | 12.0 |
| [drakkan](https://github.com/drakkan)  | 142.1k | 1.5k | 221.0 | 6.4k | Go | 11.8 |
| [kubeedge](https://github.com/kubeedge)  | 141.3k | 7.3k | 3.9k | 6.8k | Go | 11.8 |
| [gravity-ui](https://github.com/gravity-ui)  | 138.7k | 3.1k | 1.5k | 330.0 | TypeScript | 11.6 |
| [fluid-cloudnative](https://github.com/fluid-cloudnative)  | 136.8k | 2.0k | 2.0k | 1.3k | Go | 11.4 |
| [getporter](https://github.com/getporter)  | 134.4k | 6.1k | 2.5k | 1.1k | Go | 11.2 |
| [rook](https://github.com/rook)  | 132.1k | 24.8k | 7.1k | 11.2k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 4.9k | 3.8k | 4.1k | Go | 10.9 |
| [wasmedge](https://github.com/wasmedge)  | 128.4k | 3.3k | 1.5k | 5.8k | Rust | 10.7 |
| [crossplane](https://github.com/crossplane)  | 127.7k | 11.2k | 3.5k | 7.7k | Go | 10.6 |
| [armadaproject](https://github.com/armadaproject)  | 127.7k | 2.2k | 1.8k | 286.0 | Go | 10.6 |
| [keylime](https://github.com/keylime)  | 127.1k | 2.4k | 1.5k | 399.0 | C | 10.6 |
| [nocalhost](https://github.com/nocalhost)  | 123.0k | 6.1k | 2.1k | 1.6k | Go | 10.2 |
| [devfile](https://github.com/devfile)  | 121.0k | 4.3k | 2.0k | 310.0 | Go | 10.1 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 4.9k | 3.8k | 5.1k | Go | 10.0 |
| [artifacthub](https://github.com/artifacthub)  | 116.7k | 1.8k | 2.3k | 1.2k | TypeScript | 9.7 |
| [buildpacks](https://github.com/buildpacks)  | 116.7k | 16.1k | 3.5k | 2.8k | Go | 9.7 |
| [openelb](https://github.com/openelb)  | 116.4k | 498.0 | 283.0 | 1.3k | JavaScript | 9.7 |
| [serverless-devs](https://github.com/serverless-devs)  | 115.1k | 3.5k | 456.0 | 1.5k | JavaScript | 9.6 |
| [paralus](https://github.com/paralus)  | 112.6k | 1.6k | 437.0 | 758.0 | Go | 9.4 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 107.2k | 3.3k | 3.2k | 5.9k | Go | 8.9 |
| [kubescape](https://github.com/kubescape)  | 105.7k | 12.3k | 2.0k | 8.3k | Go | 8.8 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 105.7k | 18.0k | 3.3k | 4.0k | Python | 8.8 |
| [kedacore](https://github.com/kedacore)  | 105.0k | 3.9k | 4.2k | 6.8k | Go | 8.8 |
| [go-faster](https://github.com/go-faster)  | 103.0k | 6.0k | 880.0 | 166.0 | Go | 8.6 |
| [VKCOM](https://github.com/VKCOM)  | 101.5k | 10.7k | 4.2k | 956.0 | TypeScript | 8.4 |
| [karmada-io](https://github.com/karmada-io)  | 100.5k | 4.1k | 2.7k | 3.2k | Go | 8.4 |
| [bfenetworks](https://github.com/bfenetworks)  | 95.8k | 1.4k | 825.0 | 6.0k | Go | 8.0 |
| [openyurtio](https://github.com/openyurtio)  | 95.6k | 1.8k | 1.5k | 1.6k | Go | 8.0 |
| [v6d-io](https://github.com/v6d-io)  | 94.8k | 1.0k | 882.0 | 689.0 | C++ | 7.9 |
| [in-toto](https://github.com/in-toto)  | 92.3k | 9.5k | 1.0k | 1.1k | Python | 7.7 |
| [fnproject](https://github.com/fnproject)  | 87.4k | 8.5k | 2.5k | 6.4k | Go | 7.3 |
| [serverlessworkflow](https://github.com/serverlessworkflow)  | 87.3k | 2.3k | 1.2k | 944.0 | C# | 7.3 |
| [firecracker-microvm](https://github.com/firecracker-microvm)  | 82.5k | 4.6k | 2.5k | 21.9k | Rust | 6.9 |
| [openservicemesh](https://github.com/openservicemesh)  | 81.9k | 5.2k | 4.1k | 2.6k | Go | 6.8 |
| [cloudevents](https://github.com/cloudevents)  | 79.5k | 4.0k | 2.6k | 5.8k | Go | 6.6 |
| [opencontainers](https://github.com/opencontainers)  | 78.9k | 14.0k | 6.4k | 18.2k | Go | 6.6 |
| [kubeovn](https://github.com/kubeovn)  | 76.6k | 4.5k | 2.0k | 1.5k | Go | 6.4 |
| [coredns](https://github.com/coredns)  | 76.2k | 5.5k | 4.7k | 11.5k | Go | 6.3 |
| [project-zot](https://github.com/project-zot)  | 75.8k | 1.1k | 1.2k | 333.0 | Go | 6.3 |
| [gotd](https://github.com/gotd)  | 74.5k | 6.2k | 2.4k | 1.2k | Go | 6.2 |
| [ten-nancy](https://github.com/ten-nancy)  | 73.1k | 9.0 | 0.0 | 0.0 | C++ | 6.1 |
| [spectralops](https://github.com/spectralops)  | 73.0k | 359.0 | 63.0 | 1.2k | JavaScript | 6.1 |
| [helm](https://github.com/helm)  | 71.5k | 28.9k | 26.9k | 48.8k | Go | 6.0 |
| [cri-o](https://github.com/cri-o)  | 71.3k | 9.2k | 5.8k | 4.5k | Go | 5.9 |
| [meilisearch](https://github.com/meilisearch)  | 70.7k | 8.1k | 1.6k | 37.3k | Rust | 5.9 |
| [submariner-io](https://github.com/submariner-io)  | 69.8k | 10.0k | 9.1k | 2.3k | Go | 5.8 |
| [opencost](https://github.com/opencost)  | 68.6k | 3.7k | 1.4k | 3.6k | Go | 5.7 |
| [volcano-sh](https://github.com/volcano-sh)  | 68.5k | 5.5k | 2.0k | 3.2k | Go | 5.7 |
| [metal3-io](https://github.com/metal3-io)  | 68.4k | 9.4k | 4.6k | 1.0k | Go | 5.7 |
| [notaryproject](https://github.com/notaryproject)  | 64.4k | 3.7k | 1.9k | 3.3k | Go | 5.4 |
| [ent](https://github.com/ent)  | 63.2k | 2.2k | 2.1k | 13.7k | Go | 5.3 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 1.8k | 2.2k | 2.2k | Go | 5.2 |
| [oras-project](https://github.com/oras-project)  | 61.5k | 2.2k | 1.0k | 1.1k | Go | 5.1 |
| [vuejs](https://github.com/vuejs)  | 60.3k | 3.5k | 2.4k | 204.2k | TypeScript | 5.0 |
| [openfaas](https://github.com/openfaas)  | 59.0k | 9.0k | 4.1k | 31.9k | Go | 4.9 |
| [projectriff](https://github.com/projectriff)  | 57.7k | 6.6k | 4.2k | 887.0 | Go | 4.8 |
| [telepresenceio](https://github.com/telepresenceio)  | 57.4k | 8.0k | 1.7k | 5.7k | Go | 4.8 |
| [external-secrets](https://github.com/external-secrets)  | 56.9k | 2.8k | 1.8k | 5.1k | Go | 4.7 |
| [ariga](https://github.com/ariga)  | 55.7k | 1.6k | 1.5k | 3.2k | Go | 4.6 |
| [kubewarden](https://github.com/kubewarden)  | 55.7k | 8.5k | 3.0k | 489.0 | Rust | 4.6 |
| [open-feature](https://github.com/open-feature)  | 53.0k | 5.1k | 3.2k | 754.0 | Go | 4.4 |
| [rkt](https://github.com/rkt)  | 52.4k | 5.6k | 2.5k | 8.9k | Go | 4.4 |
| [trickstercache](https://github.com/trickstercache)  | 52.1k | 1.1k | 458.0 | 1.9k | Go | 4.3 |
| [superedge](https://github.com/superedge)  | 48.2k | 487.0 | 425.0 | 913.0 | Go | 4.0 |
| [virtual-kubelet](https://github.com/virtual-kubelet)  | 48.1k | 2.8k | 1.3k | 4.4k | Go | 4.0 |
| [kudobuilder](https://github.com/kudobuilder)  | 47.9k | 2.2k | 2.0k | 1.8k | Go | 4.0 |
| [theupdateframework](https://github.com/theupdateframework)  | 45.4k | 8.6k | 2.6k | 2.6k | Rust | 3.8 |
| [openfga](https://github.com/openfga)  | 45.2k | 1.2k | 1.2k | 979.0 | Go | 3.8 |
| [kubedl-io](https://github.com/kubedl-io)  | 44.1k | 1.3k | 369.0 | 488.0 | Go | 3.7 |
| [skooner-k8s](https://github.com/skooner-k8s)  | 44.0k | 507.0 | 213.0 | 1.1k | JavaScript | 3.7 |
| [tauri-apps](https://github.com/tauri-apps)  | 43.4k | 3.8k | 3.6k | 65.8k | Rust | 3.6 |
| [distribution](https://github.com/distribution)  | 42.9k | 4.8k | 1.8k | 7.4k | Go | 3.6 |
| [openfunction](https://github.com/openfunction)  | 41.4k | 1.4k | 915.0 | 1.2k | Go | 3.5 |
| [metallb](https://github.com/metallb)  | 40.9k | 2.3k | 1.3k | 5.7k | Go | 3.4 |
| [vitalif](https://github.com/vitalif)  | 40.9k | 1.4k | 12.0 | 75.0 | C++ | 3.4 |
| [tinkerbell](https://github.com/tinkerbell)  | 40.5k | 6.8k | 2.4k | 1.6k | Go | 3.4 |
| [capsule-rs](https://github.com/capsule-rs)  | 40.4k | 252.0 | 140.0 | 356.0 | Rust | 3.4 |
| [devstream-io](https://github.com/devstream-io)  | 39.3k | 2.9k | 1.1k | 747.0 | Go | 3.3 |
| [vscode-kubernetes-tools](https://github.com/vscode-kubernetes-tools)  | 39.3k | 634.0 | 613.0 | 602.0 | TypeScript | 3.3 |
| [ogen-go](https://github.com/ogen-go)  | 39.2k | 4.0k | 1.1k | 473.0 | Go | 3.3 |
| [k3s-io](https://github.com/k3s-io)  | 38.5k | 3.5k | 3.3k | 25.3k | Go | 3.2 |
| [jquery](https://github.com/jquery)  | 37.2k | 6.6k | 2.9k | 57.6k | JavaScript | 3.1 |
| [aeraki-mesh](https://github.com/aeraki-mesh)  | 36.7k | 988.0 | 368.0 | 866.0 | Go | 3.1 |
| [kubearmor](https://github.com/kubearmor)  | 35.8k | 5.5k | 2.1k | 718.0 | Go | 3.0 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 2.8k | 1.2k | 6.5k | Go | 2.9 |
| [clusternet](https://github.com/clusternet)  | 33.4k | 948.0 | 545.0 | 1.1k | Go | 2.8 |
| [sealerio](https://github.com/sealerio)  | 32.7k | 1.2k | 1.3k | 1.8k | Go | 2.7 |
| [kube-rs](https://github.com/kube-rs)  | 31.6k | 3.0k | 775.0 | 2.2k | Rust | 2.6 |
| [dexidp](https://github.com/dexidp)  | 29.3k | 2.6k | 1.8k | 8.0k | Go | 2.4 |
| [clusterpedia-io](https://github.com/clusterpedia-io)  | 29.1k | 1.1k | 541.0 | 596.0 | Go | 2.4 |
| [opentracing](https://github.com/opentracing)  | 27.4k | 2.7k | 1.4k | 10.1k | Java | 2.3 |
| [krustlet](https://github.com/krustlet)  | 23.6k | 2.4k | 530.0 | 3.4k | Rust | 2.0 |
| [alacritty](https://github.com/alacritty)  | 23.2k | 2.1k | 2.0k | 47.4k | Rust | 1.9 |
| [curiefense](https://github.com/curiefense)  | 22.3k | 2.9k | 817.0 | 617.0 | Rust | 1.9 |
| [fabedge](https://github.com/fabedge)  | 21.6k | 941.0 | 429.0 | 482.0 | Go | 1.8 |
| [project-akri](https://github.com/project-akri)  | 20.6k | 701.0 | 448.0 | 946.0 | Rust | 1.7 |
| [piraeusdatastore](https://github.com/piraeusdatastore)  | 19.7k | 1.3k | 522.0 | 748.0 | Go | 1.6 |
| [k8gb-io](https://github.com/k8gb-io)  | 19.4k | 1.1k | 918.0 | 571.0 | Go | 1.6 |
| [foniod](https://github.com/foniod)  | 18.1k | 1.7k | 493.0 | 2.0k | Rust | 1.5 |
| [lima-vm](https://github.com/lima-vm)  | 17.4k | 2.1k | 934.0 | 11.7k | Go | 1.4 |
| [schemahero](https://github.com/schemahero)  | 17.0k | 1.4k | 773.0 | 765.0 | Go | 1.4 |
| [hexa-org](https://github.com/hexa-org)  | 16.5k | 491.0 | 118.0 | 70.0 | Go | 1.4 |
| [pallets](https://github.com/pallets)  | 15.0k | 5.1k | 2.4k | 63.4k | Python | 1.2 |
| [k8up-io](https://github.com/k8up-io)  | 13.5k | 2.0k | 699.0 | 412.0 | Go | 1.1 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 749.0 | 118.0 | 588.0 | Go | 1.1 |
| [servicemeshinterface](https://github.com/servicemeshinterface)  | 11.6k | 966.0 | 372.0 | 1.2k | Go | 970.0m |
| [kuberhealthy](https://github.com/kuberhealthy)  | 10.8k | 3.1k | 721.0 | 1.7k | Go | 900.0m |
| [kube-vip](https://github.com/kube-vip)  | 9.7k | 916.0 | 320.0 | 1.3k | Go | 800.0m |
| [ko-build](https://github.com/ko-build)  | 9.5k | 793.0 | 721.0 | 5.8k | Go | 790.0m |
| [devspace-cloud](https://github.com/devspace-cloud)  | 8.7k | 348.0 | 83.0 | 113.0 | Go | 720.0m |
| [tellerops](https://github.com/tellerops)  | 7.7k | 179.0 | 80.0 | 1.6k | Go | 640.0m |
| [cni-genie](https://github.com/cni-genie)  | 5.9k | 484.0 | 132.0 | 499.0 | Go | 490.0m |
| [merbridge](https://github.com/merbridge)  | 4.8k | 271.0 | 279.0 | 586.0 | Go | 400.0m |
| [opcr-io](https://github.com/opcr-io)  | 3.6k | 597.0 | 153.0 | 181.0 | Go | 300.0m |
| [krator-rs](https://github.com/krator-rs)  | 2.9k | 498.0 | 33.0 | 113.0 | Rust | 250.0m |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [openobservability](https://github.com/openobservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [service-mesh-performance](https://github.com/service-mesh-performance)  | 2.2k | 797.0 | 217.0 | 237.0 | JavaScript | 180.0m |
| [kubereboot](https://github.com/kubereboot)  | 2.0k | 997.0 | 558.0 | 1.8k | Go | 170.0m |
| [g-research](https://github.com/g-research)  | 870.0 | 9.0 | 0.0 | 4.0 | Go | 70.0m |
| [gitops-working-group](https://github.com/gitops-working-group)  | 0.0 | 43.0 | 57.0 | 552.0 | N/A | 0.0 |
| [zephyrproject](https://github.com/zephyrproject)  | 0.0 | 0.0 | 0.0 | 0.0 | N/A | 0.0 |


### Repositories
| Repository | SLOC | Commits | PRs | Stars | Language | Effort |
|------------|------|---------|-----|-------|----------|--------|
| [torvalds/linux](https://github.com/torvalds/linux/)  | 19.1M | 1.2M | 761.0 | 153.7k | C | 1.6k |
| [zephyrproject-rtos/hal_nxp](https://github.com/zephyrproject-rtos/hal_nxp/)  | 5.8M | 611.0 | 228.0 | 17.0 | C Header | 482.4 |
| [zephyrproject-rtos/hal_stm32](https://github.com/zephyrproject-rtos/hal_stm32/)  | 4.6M | 463.0 | 167.0 | 55.0 | C Header | 382.3 |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3.4M | 149.9k | 23.1k | 175.7k | C++ | 279.2 |
| [ydb-platform/nbs](https://github.com/ydb-platform/nbs/)  | 3.2M | 802.0 | 0.0 | 14.0 | C++ | 267.1 |
| [ydb-platform/ydb](https://github.com/ydb-platform/ydb/)  | 2.8M | 9.4k | 35.0 | 2.8k | C++ | 237.4 |
| [ytsaurus/ytsaurus](https://github.com/ytsaurus/ytsaurus/)  | 2.5M | 65.5k | 5.0 | 799.0 | C++ | 204.8 |
| [elastic/elasticsearch](https://github.com/elastic/elasticsearch/)  | 2.4M | 69.3k | 64.5k | 64.3k | Java | 201.9 |
| [apache/hadoop](https://github.com/apache/hadoop/)  | 2.0M | 26.7k | 5.8k | 13.6k | Java | 165.6 |
| [godotengine/godot](https://github.com/godotengine/godot/)  | 2.0M | 55.4k | 35.0k | 62.2k | C++ | 164.8 |
| [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/)  | 1.7M | 78.5k | 48.5k | 26.8k | Go | 139.7 |
| [rust-lang/rust](https://github.com/rust-lang/rust/)  | 1.5M | 227.8k | 63.9k | 83.0k | Rust | 126.5 |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1.5M | 114.5k | 73.6k | 96.5k | Go | 124.0 |
| [apache/camel](https://github.com/apache/camel/)  | 1.4M | 64.6k | 10.1k | 4.8k | Java | 117.9 |
| [python/cpython](https://github.com/python/cpython/)  | 1.4M | 117.6k | 41.8k | 54.2k | Python | 114.4 |
| [zephyrproject-rtos/zephyr](https://github.com/zephyrproject-rtos/zephyr/)  | 1.3M | 76.7k | 38.1k | 7.5k | C | 104.8 |
| [zephyrproject-rtos/zephyr-testing](https://github.com/zephyrproject-rtos/zephyr-testing/)  | 1.3M | 76.6k | 199.0 | 6.0 | C | 104.7 |
| [openkruise/game.openkruise.io](https://github.com/openkruise/game.openkruise.io/)  | 1.2M | 4.0 | 0.0 | 0.0 | JavaScript | 99.0 |
| [apache/hive](https://github.com/apache/hive/)  | 1.1M | 16.8k | 4.3k | 4.8k | Java | 93.6 |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1.1M | 109.8k | 32.2k | 27.4k | C++ | 92.7 |
| [apache/harmony-classlib](https://github.com/apache/harmony-classlib/)  | 1.1M | 7.5k | 3.0 | 6.0 | Java | 92.4 |
| [golang/go](https://github.com/golang/go/)  | 1.1M | 57.1k | 3.0k | 112.4k | Go | 88.4 |
| [prestodb/presto](https://github.com/prestodb/presto/)  | 1.0M | 20.8k | 13.9k | 14.6k | Java | 86.0 |
| [zephyrproject-rtos/hal_silabs](https://github.com/zephyrproject-rtos/hal_silabs/)  | 1.0M | 54.0 | 24.0 | 4.0 | C Header | 84.6 |
| [zephyrproject-rtos/sdk-ng](https://github.com/zephyrproject-rtos/sdk-ng/)  | 980.2k | 730.0 | 453.0 | 110.0 | C | 81.7 |
| [nodejs/node](https://github.com/nodejs/node/)  | 966.7k | 39.2k | 30.7k | 96.2k | JavaScript | 80.6 |
| [konveyor/github-actions](https://github.com/konveyor/github-actions/)  | 897.5k | 17.0 | 13.0 | 0.0 | JavaScript | 74.8 |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889.8k | 41.4k | 33.6k | 54.3k | TypeScript | 74.2 |
| [ziglang/zig](https://github.com/ziglang/zig/)  | 888.6k | 24.8k | 7.5k | 23.5k | Zig | 74.0 |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 807.6k | 17.3k | 17.1k | 21.6k | C++ | 67.3 |
| [brigadecore/go-libgit2](https://github.com/brigadecore/go-libgit2/)  | 782.3k | 3.0 | 0.0 | 0.0 | JavaScript | 65.2 |
| [zephyrproject-rtos/hal_espressif](https://github.com/zephyrproject-rtos/hal_espressif/)  | 765.0k | 21.1k | 187.0 | 25.0 | C | 63.8 |
| [apache/cassandra](https://github.com/apache/cassandra/)  | 632.5k | 28.6k | 2.3k | 8.0k | Java | 52.7 |
| [elastic/beats](https://github.com/elastic/beats/)  | 616.2k | 16.6k | 26.5k | 11.7k | Go | 51.4 |
| [vitessio/vitess](https://github.com/vitessio/vitess/)  | 611.0k | 32.3k | 9.6k | 15.7k | Go | 50.9 |
| [apache/lucenenet](https://github.com/apache/lucenenet/)  | 608.0k | 6.7k | 619.0 | 2.0k | C# | 50.7 |
| [zephyrproject-rtos/newlib-cygwin](https://github.com/zephyrproject-rtos/newlib-cygwin/)  | 567.8k | 19.2k | 6.0 | 1.0 | C | 47.3 |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 514.8k | 52.7k | 21.6k | 37.2k | C++ | 42.9 |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 512.3k | 96.0 | 103.0 | 22.0 | C++ | 42.7 |
| [falcosecurity/libs](https://github.com/falcosecurity/libs/)  | 504.5k | 7.4k | 964.0 | 156.0 | C Header | 42.0 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 504.5k | 8.5k | 454.0 | 205.0 | C++ | 42.0 |
| [denoland/deno](https://github.com/denoland/deno/)  | 498.2k | 9.6k | 10.7k | 89.9k | JavaScript | 41.5 |
| [cubefs/cubefs](https://github.com/cubefs/cubefs/)  | 481.6k | 2.5k | 1.4k | 3.2k | Go | 40.1 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477.3k | 4.2k | 3.6k | 4.4k | Go | 39.8 |
| [zephyrproject-rtos/trusted-firmwar~](https://github.com/zephyrproject-rtos/trusted-firmware-m/)  | 471.1k | 3.9k | 89.0 | 13.0 | C | 39.2 |
| [apache/openjpa](https://github.com/apache/openjpa/)  | 464.1k | 5.3k | 98.0 | 114.0 | Java | 38.7 |
| [apache/activemq](https://github.com/apache/activemq/)  | 444.6k | 11.4k | 999.0 | 2.1k | Java | 37.0 |
| [zephyrproject-rtos/loramac-node](https://github.com/zephyrproject-rtos/loramac-node/)  | 435.6k | 1.6k | 11.0 | 16.0 | C Header | 36.3 |
| [apache/poi](https://github.com/apache/poi/)  | 424.4k | 12.6k | 457.0 | 1.6k | Java | 35.4 |
| [konveyor/tackle-diva](https://github.com/konveyor/tackle-diva/)  | 400.7k | 106.0 | 88.0 | 39.0 | Java | 33.4 |
| [tikv/tikv](https://github.com/tikv/tikv/)  | 395.8k | 7.3k | 10.7k | 13.3k | Rust | 33.0 |
| [netdata/charts](https://github.com/netdata/charts/)  | 395.7k | 582.0 | 167.0 | 9.0 | JavaScript | 33.0 |
| [backstage/backstage](https://github.com/backstage/backstage/)  | 393.7k | 38.2k | 12.3k | 21.1k | TypeScript | 32.8 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389.2k | 9.6k | 16.2k | 1.5k | Go | 32.4 |
| [netdata/netdata](https://github.com/netdata/netdata/)  | 380.5k | 15.2k | 7.3k | 62.4k | C | 31.7 |
| [django/django](https://github.com/django/django/)  | 365.8k | 31.8k | 17.0k | 71.7k | Python | 30.5 |
| [zephyrproject-rtos/esp-idf](https://github.com/zephyrproject-rtos/esp-idf/)  | 363.8k | 5.6k | 4.0 | 0.0 | C | 30.3 |
| [cdk8s-team/cdk8s-examples](https://github.com/cdk8s-team/cdk8s-examples/)  | 360.6k | 34.0 | 36.0 | 4.0 | Java | 30.1 |
| [facebook/react](https://github.com/facebook/react/)  | 357.5k | 15.9k | 13.7k | 209.6k | JavaScript | 29.8 |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 352.4k | 12.2k | 1.0k | 4.7k | C++ | 29.4 |
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
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310.9k | 4.5k | 12.0 | 701.0 | Java | 25.9 |
| [apache/mesos](https://github.com/apache/mesos/)  | 305.6k | 18.2k | 451.0 | 5.1k | C++ | 25.5 |
| [docker/labs](https://github.com/docker/labs/)  | 304.4k | 718.0 | 398.0 | 11.1k | PHP | 25.4 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284.4k | 1.6k | 1.3k | 465.0 | C Header | 23.7 |
| [pravega/pravega](https://github.com/pravega/pravega/)  | 283.9k | 3.2k | 3.2k | 1.9k | Java | 23.6 |
| [kubevirt/kubevirt](https://github.com/kubevirt/kubevirt/)  | 281.1k | 16.3k | 7.0k | 3.9k | Go | 23.4 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274.6k | 9.4k | 9.4k | 12.9k | Rust | 22.9 |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 270.6k | 51.9k | 24.6k | 3.6k | Go | 22.6 |
| [cdk8s-team/cdk8s](https://github.com/cdk8s-team/cdk8s/)  | 266.0k | 1.0k | 888.0 | 3.5k | TypeScript | 22.2 |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 264.4k | 1.6k | 36.0 | 24.0 | JavaScript | 22.0 |
| [apache/cayenne](https://github.com/apache/cayenne/)  | 261.7k | 7.1k | 580.0 | 287.0 | Java | 21.8 |
| [kyverno/kyverno](https://github.com/kyverno/kyverno/)  | 259.5k | 5.5k | 3.9k | 3.6k | Go | 21.6 |
| [inclavare-containers/deterministic~](https://github.com/inclavare-containers/deterministic-builds/)  | 259.3k | 28.0 | 4.0 | 1.0 | C Header | 21.6 |
| [apache/wicket](https://github.com/apache/wicket/)  | 258.1k | 21.4k | 550.0 | 684.0 | Java | 21.5 |
| [apache/httpd](https://github.com/apache/httpd/)  | 256.2k | 33.5k | 308.0 | 3.2k | C | 21.4 |
| [apache/pig](https://github.com/apache/pig/)  | 254.4k | 3.7k | 38.0 | 655.0 | Java | 21.2 |
| [inspektor-gadget/inspektor-gadget](https://github.com/inspektor-gadget/inspektor-gadget/)  | 251.3k | 2.8k | 1.0k | 1.3k | C Header | 20.9 |
| [zephyrproject-rtos/hal_ti](https://github.com/zephyrproject-rtos/hal_ti/)  | 249.5k | 102.0 | 45.0 | 5.0 | C Header | 20.8 |
| [moby/moby](https://github.com/moby/moby/)  | 244.5k | 45.7k | 23.1k | 65.6k | Go | 20.4 |
| [zephyrproject-rtos/lvgl](https://github.com/zephyrproject-rtos/lvgl/)  | 242.5k | 7.8k | 35.0 | 38.0 | C | 20.2 |
| [apache/cocoon](https://github.com/apache/cocoon/)  | 241.8k | 13.2k | 44.0 | 22.0 | Java | 20.1 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241.1k | 4.7k | 5.2k | 18.5k | Go | 20.1 |
| [zephyrproject-rtos/openocd](https://github.com/zephyrproject-rtos/openocd/)  | 240.0k | 9.6k | 39.0 | 20.0 | C | 20.0 |
| [zephyrproject-rtos/hal_st](https://github.com/zephyrproject-rtos/hal_st/)  | 239.6k | 18.0 | 12.0 | 8.0 | C | 20.0 |
| [matplotlib/matplotlib](https://github.com/matplotlib/matplotlib/)  | 235.9k | 47.5k | 16.6k | 17.6k | Python | 19.6 |
| [open-policy-agent/opa](https://github.com/open-policy-agent/opa/)  | 235.4k | 4.5k | 3.5k | 7.8k | Go | 19.6 |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235.2k | 5.7k | 6.9k | 10.4k | Java | 19.6 |
| [netdata/ebpf-co-re](https://github.com/netdata/ebpf-co-re/)  | 231.9k | 35.0 | 38.0 | 6.0 | C Header | 19.3 |
| [kumahq/kuma](https://github.com/kumahq/kuma/)  | 221.4k | 4.0k | 4.8k | 3.1k | Go | 18.4 |
| [apache/xmlgraphics-fop](https://github.com/apache/xmlgraphics-fop/)  | 219.8k | 8.5k | 44.0 | 154.0 | Java | 18.3 |
| [chaosblade-io/chaosblade-box](https://github.com/chaosblade-io/chaosblade-box/)  | 219.5k | 40.0 | 41.0 | 186.0 | JavaScript | 18.3 |
| [apache/directory-studio](https://github.com/apache/directory-studio/)  | 216.5k | 5.7k | 27.0 | 101.0 | Java | 18.0 |
| [uber/peloton](https://github.com/uber/peloton/)  | 216.4k | 705.0 | 10.0 | 586.0 | Go | 18.0 |
| [zephyrproject-rtos/hal_infineon](https://github.com/zephyrproject-rtos/hal_infineon/)  | 216.1k | 14.0 | 8.0 | 3.0 | C | 18.0 |
| [zephyrproject-rtos/cmsis](https://github.com/zephyrproject-rtos/cmsis/)  | 203.3k | 31.0 | 20.0 | 4.0 | C | 16.9 |
| [zephyrproject-rtos/openthread](https://github.com/zephyrproject-rtos/openthread/)  | 199.9k | 7.0k | 83.0 | 23.0 | C++ | 16.7 |
| [antrea-io/antrea](https://github.com/antrea-io/antrea/)  | 198.7k | 2.7k | 3.5k | 1.4k | Go | 16.6 |
| [zephyrproject-rtos/hal_nordic](https://github.com/zephyrproject-rtos/hal_nordic/)  | 196.0k | 227.0 | 127.0 | 15.0 | C Header | 16.3 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194.2k | 181.0 | 125.0 | 1.2k | C Header | 16.2 |
| [apache/xmlgraphics-batik](https://github.com/apache/xmlgraphics-batik/)  | 193.7k | 3.6k | 36.0 | 173.0 | Java | 16.1 |
| [goharbor/harbor](https://github.com/goharbor/harbor/)  | 191.5k | 11.6k | 8.0k | 19.5k | Go | 16.0 |
| [apache/tapestry4](https://github.com/apache/tapestry4/)  | 190.8k | 3.8k | 11.0 | 7.0 | Java | 15.9 |
| [nats-io/nats-server](https://github.com/nats-io/nats-server/)  | 190.7k | 6.8k | 2.6k | 12.4k | Go | 15.9 |
| [apache/avro](https://github.com/apache/avro/)  | 189.0k | 3.7k | 2.2k | 2.5k | Java | 15.8 |
| [alibaba/fastjson](https://github.com/alibaba/fastjson/)  | 188.1k | 4.0k | 848.0 | 25.2k | Java | 15.7 |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186.6k | 1.0 | 2.0 | 2.0 | JavaScript | 15.6 |
| [fluent/fluent-bit](https://github.com/fluent/fluent-bit/)  | 182.4k | 9.3k | 3.3k | 4.4k | C | 15.2 |
| [apache/directory-server](https://github.com/apache/directory-server/)  | 181.5k | 10.0k | 80.0 | 118.0 | Java | 15.1 |
| [cncf/cncf.io](https://github.com/cncf/cncf.io/)  | 180.9k | 1.9k | 88.0 | 66.0 | PHP | 15.1 |
| [cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian/)  | 175.1k | 4.1k | 4.4k | 4.7k | Python | 14.6 |
| [strimzi/strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator/)  | 175.0k | 5.1k | 4.7k | 3.7k | Java | 14.6 |
| [containers/podman](https://github.com/containers/podman/)  | 173.6k | 19.3k | 10.8k | 18.3k | Go | 14.5 |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173.3k | 8.9k | 5.3k | 2.8k | Go | 14.4 |
| [apache/xalan-j](https://github.com/apache/xalan-j/)  | 170.9k | 4.6k | 6.0 | 25.0 | Java | 14.2 |
| [argoproj/argo-workflows](https://github.com/argoproj/argo-workflows/)  | 169.8k | 4.1k | 4.8k | 12.6k | Go | 14.2 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.9k | 8.8k | 5.1k | 1.2k | Java | 14.0 |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167.6k | 19.4k | 10.4k | 14.8k | Go | 14.0 |
| [apache/tika](https://github.com/apache/tika/)  | 167.3k | 7.3k | 1.1k | 1.7k | Java | 13.9 |
| [keptn/keptn](https://github.com/keptn/keptn/)  | 162.3k | 8.3k | 6.1k | 1.7k | Go | 13.5 |
| [apache/couchdb](https://github.com/apache/couchdb/)  | 159.2k | 13.2k | 2.8k | 5.7k | Erlang | 13.3 |
| [apache/shindig](https://github.com/apache/shindig/)  | 158.2k | 4.8k | 7.0 | 69.0 | Java | 13.2 |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157.3k | 1.7k | 1.3k | 318.0 | Java | 13.1 |
| [spiffe/spire](https://github.com/spiffe/spire/)  | 156.7k | 5.8k | 2.8k | 1.4k | Go | 13.1 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155.1k | 10.9k | 6.8k | 47.1k | Go | 12.9 |
| [apache/etch](https://github.com/apache/etch/)  | 153.2k | 549.0 | 0.0 | 17.0 | Java | 12.8 |
| [werf/actions](https://github.com/werf/actions/)  | 153.1k | 118.0 | 60.0 | 68.0 | JavaScript | 12.8 |
| [apache/stdcxx](https://github.com/apache/stdcxx/)  | 150.4k | 2.2k | 0.0 | 57.0 | C++ | 12.5 |
| [containerd/containerd](https://github.com/containerd/containerd/)  | 150.2k | 12.0k | 5.9k | 13.3k | Go | 12.5 |
| [argoproj/argo-cd](https://github.com/argoproj/argo-cd/)  | 149.4k | 5.1k | 5.8k | 12.3k | Go | 12.4 |
| [kubevela/samples](https://github.com/kubevela/samples/)  | 146.8k | 219.0 | 118.0 | 93.0 | JavaScript | 12.2 |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144.5k | 5.9k | 1.8k | 8.1k | Go | 12.0 |
| [kubevela/kubevela](https://github.com/kubevela/kubevela/)  | 144.3k | 3.6k | 3.8k | 5.0k | Go | 12.0 |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143.2k | 4.5k | 3.9k | 17.7k | Go | 11.9 |
| [apache/maven-plugins](https://github.com/apache/maven-plugins/)  | 142.5k | 15.1k | 146.0 | 242.0 | Java | 11.9 |
| [drakkan/sftpgo](https://github.com/drakkan/sftpgo/)  | 142.1k | 1.5k | 221.0 | 6.4k | Go | 11.8 |
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
| [zephyrproject-rtos/mbedtls](https://github.com/zephyrproject-rtos/mbedtls/)  | 119.9k | 20.5k | 42.0 | 10.0 | C | 10.0 |
| [armadaproject/armada](https://github.com/armadaproject/armada/)  | 119.4k | 2.0k | 1.7k | 278.0 | Go | 9.9 |
| [cortexproject/cortex](https://github.com/cortexproject/cortex/)  | 119.0k | 4.6k | 3.5k | 5.0k | Go | 9.9 |
| [artifacthub/hub](https://github.com/artifacthub/hub/)  | 116.7k | 1.7k | 2.3k | 1.2k | TypeScript | 9.7 |
| [parallaxsecond/rust-tss-esapi](https://github.com/parallaxsecond/rust-tss-esapi/)  | 116.6k | 658.0 | 280.0 | 44.0 | Rust | 9.7 |
| [apache/aurora](https://github.com/apache/aurora/)  | 114.8k | 4.1k | 71.0 | 630.0 | Java | 9.6 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 113.7k | 2.8k | 63.0 | 10.0 | TypeScript | 9.5 |
| [thanos-io/thanos](https://github.com/thanos-io/thanos/)  | 113.2k | 3.1k | 3.6k | 11.5k | Go | 9.4 |
| [projectcontour/contour](https://github.com/projectcontour/contour/)  | 111.5k | 4.0k | 3.2k | 3.3k | Go | 9.3 |
| [netdata/go.d.plugin](https://github.com/netdata/go.d.plugin/)  | 110.3k | 1.4k | 958.0 | 145.0 | Go | 9.2 |
| [openelb/website](https://github.com/openelb/website/)  | 107.6k | 153.0 | 73.0 | 3.0 | JavaScript | 9.0 |
| [moby/buildkit](https://github.com/moby/buildkit/)  | 107.3k | 5.2k | 2.2k | 6.5k | Go | 8.9 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 107.0k | 2.3k | 1.8k | 218.0 | Go | 8.9 |
| [apache/synapse](https://github.com/apache/synapse/)  | 106.1k | 4.4k | 80.0 | 52.0 | Java | 8.8 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105.9k | 1.5k | 1.8k | 980.0 | Go | 8.8 |
| [emissary-ingress/emissary](https://github.com/emissary-ingress/emissary/)  | 105.3k | 17.9k | 3.3k | 4.0k | Python | 8.8 |
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
| [zephyrproject-rtos/hal_microchip](https://github.com/zephyrproject-rtos/hal_microchip/)  | 96.6k | 27.0 | 15.0 | 3.0 | C Header | 8.1 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.4k | 3.3k | 3.6k | 1.5k | Java | 8.0 |
| [apache/spamassassin](https://github.com/apache/spamassassin/)  | 96.0k | 28.4k | 2.0 | 242.0 | Perl | 8.0 |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 96.0k | 3.0k | 313.0 | 2.2k | Go | 8.0 |
| [rook/rook](https://github.com/rook/rook/)  | 95.8k | 9.6k | 7.0k | 10.8k | Go | 8.0 |
| [v6d-io/v6d](https://github.com/v6d-io/v6d/)  | 94.8k | 1.0k | 882.0 | 689.0 | C++ | 7.9 |
| [karmada-io/karmada](https://github.com/karmada-io/karmada/)  | 94.6k | 3.6k | 2.4k | 3.1k | Go | 7.9 |
| [elastic/logstash](https://github.com/elastic/logstash/)  | 92.8k | 10.4k | 8.6k | 13.5k | Ruby | 7.7 |
| [zephyrproject-rtos/hal_nrfx](https://github.com/zephyrproject-rtos/hal_nrfx/)  | 92.0k | 89.0 | 0.0 | 1.0 | C Header | 7.7 |
| [tremor-rs/tremor-runtime](https://github.com/tremor-rs/tremor-runtime/)  | 91.6k | 4.0k | 1.8k | 729.0 | Rust | 7.6 |
| [zephyrproject-rtos/hal_nuvoton](https://github.com/zephyrproject-rtos/hal_nuvoton/)  | 91.6k | 17.0 | 5.0 | 0.0 | C Header | 7.6 |
| [apache/commons-lang](https://github.com/apache/commons-lang/)  | 91.1k | 7.2k | 1.1k | 2.5k | Java | 7.6 |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90.1k | 21.0k | 7.6k | 26.0k | Go | 7.5 |
| [apache/maven](https://github.com/apache/maven/)  | 90.0k | 11.8k | 1.1k | 3.6k | Java | 7.5 |
| [meshery/meshery](https://github.com/meshery/meshery/)  | 87.6k | 19.3k | 4.9k | 1.9k | Go | 7.3 |
| [wasmedge/WasmEdge](https://github.com/wasmedge/WasmEdge/)  | 87.4k | 2.5k | 1.4k | 5.6k | C++ | 7.3 |
| [cubefs/cubefs-blobstore](https://github.com/cubefs/cubefs-blobstore/)  | 87.4k | 82.0 | 3.0 | 14.0 | Go | 7.3 |
| [tremor-rs/tremor-www-docs](https://github.com/tremor-rs/tremor-www-docs/)  | 87.1k | 405.0 | 153.0 | 24.0 | JavaScript | 7.3 |
| [knative/eventing](https://github.com/knative/eventing/)  | 87.1k | 3.5k | 4.3k | 1.3k | Go | 7.3 |
| [operator-framework/operator-lifecy~](https://github.com/operator-framework/operator-lifecycle-manager/)  | 87.0k | 4.2k | 2.2k | 1.5k | Go | 7.2 |
| [moby/swarmkit](https://github.com/moby/swarmkit/)  | 86.8k | 4.2k | 2.3k | 3.0k | Go | 7.2 |
| [fluent/onigmo](https://github.com/fluent/onigmo/)  | 86.6k | 1.2k | 5.0 | 2.0 | C Header | 7.2 |
| [zephyrproject-rtos/net-tools](https://github.com/zephyrproject-rtos/net-tools/)  | 86.4k | 138.0 | 57.0 | 24.0 | C | 7.2 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84.5k | 3.4k | 3.2k | 210.0 | Go | 7.0 |
| [serverless-devs/serverless-registr~](https://github.com/serverless-devs/serverless-registry-website/)  | 84.3k | 31.0 | 0.0 | 1.0 | JavaScript | 7.0 |
| [dapr/components-contrib](https://github.com/dapr/components-contrib/)  | 83.8k | 2.9k | 1.5k | 426.0 | Go | 7.0 |
| [argoproj/argo-rollouts](https://github.com/argoproj/argo-rollouts/)  | 82.6k | 1.2k | 1.5k | 1.9k | Go | 6.9 |
| [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker/)  | 82.5k | 4.6k | 2.5k | 21.9k | Rust | 6.9 |
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
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77.2k | 3.6k | 2.2k | 244.0 | Go | 6.4 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76.5k | 630.0 | 1.2k | 3.3k | Go | 6.4 |
| [metrico/loki-apache](https://github.com/metrico/loki-apache/)  | 76.4k | 9.0 | 0.0 | 0.0 | Go | 6.4 |
| [kubevela/velaux](https://github.com/kubevela/velaux/)  | 75.2k | 589.0 | 478.0 | 95.0 | TypeScript | 6.3 |
| [chaos-mesh/chaos-mesh](https://github.com/chaos-mesh/chaos-mesh/)  | 74.9k | 1.7k | 2.4k | 5.5k | Go | 6.2 |
| [apache/jspwiki](https://github.com/apache/jspwiki/)  | 74.4k | 9.3k | 271.0 | 93.0 | Java | 6.2 |
| [openservicemesh/osm](https://github.com/openservicemesh/osm/)  | 73.9k | 4.5k | 3.6k | 2.6k | Go | 6.2 |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73.8k | 6.5k | 1.2k | 628.0 | Go | 6.2 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73.8k | 1.7k | 634.0 | 1.2k | Java | 6.2 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73.7k | 3.8k | 2.7k | 543.0 | Go | 6.1 |
| [apache/maven-sandbox](https://github.com/apache/maven-sandbox/)  | 73.4k | 2.3k | 2.0 | 5.0 | Java | 6.1 |
| [linkerd/linkerd](https://github.com/linkerd/linkerd/)  | 73.3k | 1.4k | 1.4k | 5.4k | Scala | 6.1 |
| [bfenetworks/bfe](https://github.com/bfenetworks/bfe/)  | 73.3k | 1.1k | 730.0 | 5.8k | Go | 6.1 |
| [grafana/k6](https://github.com/grafana/k6/)  | 73.0k | 5.1k | 1.5k | 19.7k | Go | 6.1 |
| [apache/httpcomponents-client](https://github.com/apache/httpcomponents-client/)  | 72.9k | 3.4k | 445.0 | 1.3k | Java | 6.1 |
| [zephyrproject-rtos/hal_openisa](https://github.com/zephyrproject-rtos/hal_openisa/)  | 72.9k | 15.0 | 7.0 | 0.0 | C Header | 6.1 |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72.4k | 4.8k | 5.3k | 2.4k | TypeScript | 6.0 |
| [dragonflyoss/Dragonfly2](https://github.com/dragonflyoss/Dragonfly2/)  | 72.3k | 1.7k | 2.0k | 1.4k | Go | 6.0 |
| [jaegertracing/jaeger](https://github.com/jaegertracing/jaeger/)  | 71.9k | 2.0k | 2.3k | 17.3k | Go | 6.0 |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71.9k | 633.0 | 232.0 | 1.1k | JavaScript | 6.0 |
| [project-zot/zot](https://github.com/project-zot/zot/)  | 71.1k | 757.0 | 928.0 | 329.0 | Go | 5.9 |
| [ten-nancy/porto](https://github.com/ten-nancy/porto/)  | 70.9k | 6.0 | 0.0 | 0.0 | C++ | 5.9 |
| [meilisearch/meilisearch](https://github.com/meilisearch/meilisearch/)  | 70.7k | 8.1k | 1.6k | 37.3k | Rust | 5.9 |
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
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65.6k | 3.8k | 204.0 | 7.4k | Go | 5.5 |
| [apache/struts1](https://github.com/apache/struts1/)  | 65.4k | 5.3k | 10.0 | 27.0 | Java | 5.5 |
| [nats-io/nats.c](https://github.com/nats-io/nats.c/)  | 64.9k | 972.0 | 399.0 | 301.0 | C | 5.4 |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63.9k | 2.8k | 1.5k | 2.1k | Go | 5.3 |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63.9k | 859.0 | 941.0 | 3.5k | C# | 5.3 |
| [ent/ent](https://github.com/ent/ent/)  | 63.2k | 2.2k | 2.1k | 13.7k | Go | 5.3 |
| [fluent/fluent-bit-website-old](https://github.com/fluent/fluent-bit-website-old/)  | 63.2k | 19.0 | 0.0 | 2.0 | JavaScript | 5.3 |
| [ydb-platform/ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk/)  | 62.9k | 2.9k | 465.0 | 80.0 | Go | 5.2 |
| [werf/trdl-vault-actions](https://github.com/werf/trdl-vault-actions/)  | 62.2k | 13.0 | 7.0 | 0.0 | JavaScript | 5.2 |
| [kubeovn/kube-ovn](https://github.com/kubeovn/kube-ovn/)  | 61.8k | 3.0k | 1.9k | 1.5k | Go | 5.2 |
| [knative/pkg](https://github.com/knative/pkg/)  | 61.2k | 2.0k | 2.3k | 234.0 | Go | 5.1 |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61.1k | 551.0 | 493.0 | 149.0 | Go | 5.1 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.0k | 1.7k | 2.0k | 1.8k | TypeScript | 5.1 |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60.9k | 1.2k | 447.0 | 27.0 | Rust | 5.1 |
| [zephyrproject-rtos/hal_gigadevice](https://github.com/zephyrproject-rtos/hal_gigadevice/)  | 60.8k | 90.0 | 32.0 | 3.0 | C | 5.1 |
| [apache/abdera](https://github.com/apache/abdera/)  | 60.5k | 1.5k | 4.0 | 18.0 | Java | 5.0 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60.5k | 3.0k | 1.9k | 6.1k | Go | 5.0 |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60.3k | 3.5k | 2.4k | 204.2k | TypeScript | 5.0 |
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
| [apache/servicemix3](https://github.com/apache/servicemix3/)  | 57.1k | 3.0k | 22.0 | 6.0 | Java | 4.8 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9k | 2.3k | 2.8k | 2.3k | C# | 4.7 |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56.9k | 6.9k | 4.9k | 14.5k | Go | 4.7 |
| [apache/roller](https://github.com/apache/roller/)  | 56.7k | 4.7k | 117.0 | 115.0 | Java | 4.7 |
| [apache/maven-shared](https://github.com/apache/maven-shared/)  | 56.5k | 3.4k | 26.0 | 29.0 | Java | 4.7 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 56.1k | 1.8k | 1.1k | 286.0 | Go | 4.7 |
| [ariga/atlas](https://github.com/ariga/atlas/)  | 55.7k | 1.6k | 1.5k | 3.2k | Go | 4.6 |
| [volcano-sh/volcano](https://github.com/volcano-sh/volcano/)  | 55.7k | 4.6k | 1.6k | 2.9k | Go | 4.6 |
| [apache/harmony-jdktools](https://github.com/apache/harmony-jdktools/)  | 55.1k | 273.0 | 0.0 | 4.0 | Java | 4.6 |
| [prestodb/benchto](https://github.com/prestodb/benchto/)  | 54.6k | 356.0 | 47.0 | 23.0 | JavaScript | 4.5 |
| [zephyrproject-rtos/chre](https://github.com/zephyrproject-rtos/chre/)  | 54.5k | 6.4k | 8.0 | 4.0 | C++ | 4.5 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 54.2k | 1.9k | 1.1k | 457.0 | Rust | 4.5 |
| [moby/libnetwork](https://github.com/moby/libnetwork/)  | 54.1k | 3.1k | 1.9k | 2.0k | Go | 4.5 |
| [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy/)  | 54.1k | 2.2k | 2.3k | 1.7k | Rust | 4.5 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53.4k | 4.4k | 37.0 | 11.0 | Go | 4.5 |
| [dragonflyoss/image-service](https://github.com/dragonflyoss/image-service/)  | 53.2k | 3.0k | 1.1k | 795.0 | Rust | 4.4 |
| [gotd/td](https://github.com/gotd/td/)  | 52.6k | 3.1k | 788.0 | 1.1k | Go | 4.4 |
| [serverlessworkflow/synapse](https://github.com/serverlessworkflow/synapse/)  | 52.6k | 814.0 | 204.0 | 145.0 | C# | 4.4 |
| [parallaxsecond/rust-cryptoki](https://github.com/parallaxsecond/rust-cryptoki/)  | 52.6k | 252.0 | 94.0 | 30.0 | Rust | 4.4 |
| [rkt/rkt](https://github.com/rkt/rkt/)  | 51.9k | 5.6k | 2.5k | 8.8k | Go | 4.3 |
| [telepresenceio/telepresence](https://github.com/telepresenceio/telepresence/)  | 51.8k | 7.3k | 1.5k | 5.7k | Go | 4.3 |
| [zephyrproject-rtos/hal_ethos_u](https://github.com/zephyrproject-rtos/hal_ethos_u/)  | 51.7k | 96.0 | 0.0 | 0.0 | C Header | 4.3 |
| [brigadecore/brigade](https://github.com/brigadecore/brigade/)  | 51.5k | 2.0k | 1.2k | 2.4k | Go | 4.3 |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.4k | 1.6k | 990.0 | 420.0 | TypeScript | 4.3 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51.2k | 3.2k | 8.0 | 206.0 | Go | 4.3 |
| [envoyproxy/toolshed](https://github.com/envoyproxy/toolshed/)  | 51.2k | 651.0 | 658.0 | 6.0 | Python | 4.3 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.1k | 1.9k | 2.4k | 3.7k | Go | 4.3 |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50.8k | 2.5k | 997.0 | 231.0 | Go | 4.2 |
| [coredns/coredns](https://github.com/coredns/coredns/)  | 50.7k | 3.6k | 3.8k | 10.4k | Go | 4.2 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50.7k | 233.0 | 194.0 | 9.0 | Go | 4.2 |
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
| [cri-o/cri-o](https://github.com/cri-o/cri-o/)  | 47.4k | 7.6k | 5.5k | 4.4k | Go | 4.0 |
| [apache/tapestry3](https://github.com/apache/tapestry3/)  | 47.2k | 2.0k | 6.0 | 5.0 | Java | 3.9 |
| [paralus/dashboard](https://github.com/paralus/dashboard/)  | 46.8k | 299.0 | 96.0 | 9.0 | JavaScript | 3.9 |
| [nats-io/nats.net](https://github.com/nats-io/nats.net/)  | 46.8k | 747.0 | 430.0 | 566.0 | C# | 3.9 |
| [operator-framework/operator-sdk](https://github.com/operator-framework/operator-sdk/)  | 46.7k | 3.2k | 3.9k | 6.3k | Go | 3.9 |
| [knative/client](https://github.com/knative/client/)  | 46.6k | 1.1k | 1.3k | 303.0 | Go | 3.9 |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46.6k | 4.5k | 5.0 | 1.0 | Go | 3.9 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46.4k | 241.0 | 358.0 | 20.0 | C | 3.9 |
| [zephyrproject-rtos/windows-curses](https://github.com/zephyrproject-rtos/windows-curses/)  | 46.1k | 75.0 | 11.0 | 120.0 | C | 3.8 |
| [vitessio/website](https://github.com/vitessio/website/)  | 46.0k | 2.9k | 1.1k | 36.0 | JavaScript | 3.8 |
| [networkservicemesh/sdk](https://github.com/networkservicemesh/sdk/)  | 45.9k | 932.0 | 1.1k | 25.0 | Go | 3.8 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45.8k | 790.0 | 309.0 | 38.0 | JavaScript | 3.8 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.8k | 1.8k | 1.0k | 413.0 | Python | 3.8 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45.3k | 1.1k | 543.0 | 113.0 | Go | 3.8 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44.9k | 4.4k | 4.9k | 12.3k | Go | 3.7 |
| [cockroachdb/gostdlib](https://github.com/cockroachdb/gostdlib/)  | 44.7k | 10.0 | 5.0 | 2.0 | Go | 3.7 |
| [apache/servicemix4-bundles](https://github.com/apache/servicemix4-bundles/)  | 43.9k | 7.6k | 347.0 | 18.0 | Java | 3.7 |
| [notaryproject/notary](https://github.com/notaryproject/notary/)  | 43.8k | 3.0k | 992.0 | 3.0k | Go | 3.6 |
| [skooner-k8s/skooner](https://github.com/skooner-k8s/skooner/)  | 43.8k | 406.0 | 213.0 | 1.1k | JavaScript | 3.6 |
| [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper/)  | 43.7k | 1.2k | 1.5k | 3.0k | Go | 3.6 |
| [containerd/nerdctl](https://github.com/containerd/nerdctl/)  | 43.6k | 2.4k | 1.3k | 5.7k | Go | 3.6 |
| [tauri-apps/tauri](https://github.com/tauri-apps/tauri/)  | 43.4k | 3.8k | 3.6k | 65.8k | Rust | 3.6 |
| [inclavare-containers/inclavare-con~](https://github.com/inclavare-containers/inclavare-containers/)  | 43.1k | 1.4k | 923.0 | 532.0 | C | 3.6 |
| [operator-framework/java-operator-s~](https://github.com/operator-framework/java-operator-sdk/)  | 42.9k | 2.4k | 1.3k | 595.0 | Java | 3.6 |
| [distribution/distribution](https://github.com/distribution/distribution/)  | 42.9k | 4.8k | 1.8k | 7.4k | Go | 3.6 |
| [tikv/client-go](https://github.com/tikv/client-go/)  | 42.6k | 474.0 | 598.0 | 227.0 | Go | 3.5 |
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
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40.9k | 1.4k | 12.0 | 75.0 | C++ | 3.4 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 40.8k | 52.0 | 3.0 | 5.0 | JavaScript | 3.4 |
| [keptn/tutorials](https://github.com/keptn/tutorials/)  | 40.8k | 622.0 | 193.0 | 10.0 | JavaScript | 3.4 |
| [apache/sanselan](https://github.com/apache/sanselan/)  | 40.5k | 315.0 | 3.0 | 37.0 | Java | 3.4 |
| [capsule-rs/capsule](https://github.com/capsule-rs/capsule/)  | 40.4k | 229.0 | 117.0 | 347.0 | Rust | 3.4 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40.4k | 10.2k | 3.7k | 807.0 | TypeScript | 3.4 |
| [dapr/java-sdk](https://github.com/dapr/java-sdk/)  | 39.7k | 409.0 | 490.0 | 223.0 | Java | 3.3 |
| [tikv/client-java](https://github.com/tikv/client-java/)  | 39.4k | 287.0 | 534.0 | 91.0 | Java | 3.3 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39.1k | 177.0 | 142.0 | 13.0 | Go | 3.3 |
| [nats-io/nats.go](https://github.com/nats-io/nats.go/)  | 38.8k | 1.8k | 716.0 | 4.4k | Go | 3.2 |
| [spectralops/spectral-goat](https://github.com/spectralops/spectral-goat/)  | 38.3k | 3.0 | 0.0 | 2.0 | JavaScript | 3.2 |
| [vscode-kubernetes-tools/vscode-kub~](https://github.com/vscode-kubernetes-tools/vscode-kubernetes-tools/)  | 38.0k | 593.0 | 527.0 | 587.0 | TypeScript | 3.2 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38.0k | 3.0k | 1.8k | 6.2k | Go | 3.2 |
| [metallb/metallb](https://github.com/metallb/metallb/)  | 37.4k | 1.8k | 958.0 | 5.7k | Go | 3.1 |
| [uber/kraken](https://github.com/uber/kraken/)  | 37.4k | 868.0 | 239.0 | 5.5k | Go | 3.1 |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37.3k | 799.0 | 283.0 | 321.0 | Go | 3.1 |
| [jquery/jquery](https://github.com/jquery/jquery/)  | 37.2k | 6.6k | 2.9k | 57.6k | JavaScript | 3.1 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36.7k | 3.3k | 700.0 | 454.0 | Go | 3.0 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36.4k | 35.0 | 7.0 | 1.0 | C | 3.0 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36.3k | 298.0 | 358.0 | 83.0 | Go | 3.0 |
| [apache/xmlgraphics-commons](https://github.com/apache/xmlgraphics-commons/)  | 36.2k | 625.0 | 16.0 | 18.0 | Java | 3.0 |
| [argoproj/argo-events](https://github.com/argoproj/argo-events/)  | 35.9k | 1.3k | 1.5k | 1.9k | Go | 3.0 |
| [awslabs/aws-ms-deploy-assistant](https://github.com/awslabs/aws-ms-deploy-assistant/)  | 35.8k | 14.0 | 6.0 | 8.0 | JavaScript | 3.0 |
| [metrico/qryn-view](https://github.com/metrico/qryn-view/)  | 35.7k | 1.0k | 165.0 | 22.0 | TypeScript | 3.0 |
| [paralus/paralus](https://github.com/paralus/paralus/)  | 35.7k | 572.0 | 98.0 | 718.0 | Go | 3.0 |
| [apache/logging-log4j1](https://github.com/apache/logging-log4j1/)  | 35.6k | 3.2k | 22.0 | 877.0 | Java | 3.0 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35.5k | 662.0 | 313.0 | 278.0 | Go | 3.0 |
| [netdata/dashboard](https://github.com/netdata/dashboard/)  | 35.0k | 1.8k | 394.0 | 63.0 | JavaScript | 2.9 |
| [opencurve/curve-tgt](https://github.com/opencurve/curve-tgt/)  | 34.8k | 2.1k | 0.0 | 18.0 | C | 2.9 |
| [kubedl-io/kubedl](https://github.com/kubedl-io/kubedl/)  | 34.4k | 377.0 | 180.0 | 417.0 | Go | 2.9 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34.2k | 1.3k | 462.0 | 2.3k | Go | 2.9 |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34.1k | 13.0 | 0.0 | 6.0 | JavaScript | 2.8 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33.7k | 1.5k | 420.0 | 78.0 | Go | 2.8 |
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
| [meshery/meshery.io](https://github.com/meshery/meshery.io/)  | 33.1k | 2.6k | 646.0 | 188.0 | JavaScript | 2.8 |
| [buildpacks/lifecycle](https://github.com/buildpacks/lifecycle/)  | 32.9k | 1.6k | 665.0 | 158.0 | Go | 2.7 |
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
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 31.0k | 508.0 | 591.0 | 519.0 | PHP | 2.6 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30.9k | 1.3k | 308.0 | 232.0 | Go | 2.6 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30.9k | 18.0 | 4.0 | 0.0 | C Header | 2.6 |
| [konveyor/move2kube-demos](https://github.com/konveyor/move2kube-demos/)  | 30.9k | 116.0 | 240.0 | 9.0 | JavaScript | 2.6 |
| [sealerio/sealer](https://github.com/sealerio/sealer/)  | 30.8k | 988.0 | 1.2k | 1.8k | Go | 2.6 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5k | 574.0 | 597.0 | 997.0 | Rust | 2.5 |
| [oras-project/oras-go](https://github.com/oras-project/oras-go/)  | 30.5k | 366.0 | 252.0 | 113.0 | Go | 2.5 |
| [superedge/superedge](https://github.com/superedge/superedge/)  | 30.2k | 378.0 | 375.0 | 903.0 | Go | 2.5 |
| [fluxcd/source-controller](https://github.com/fluxcd/source-controller/)  | 30.1k | 1.9k | 759.0 | 191.0 | Go | 2.5 |
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
| [containernetworking/plugins](https://github.com/containernetworking/plugins/)  | 26.2k | 1.6k | 484.0 | 1.7k | Go | 2.2 |
| [zephyrproject-rtos/edtt](https://github.com/zephyrproject-rtos/edtt/)  | 26.1k | 187.0 | 0.0 | 0.0 | Python | 2.2 |
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
| [kubernetes](https://github.com/kubernetes)  | 3.4M | 381.1k | 190.5k | 243.1k | Go | 286.4 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 2.0M | 118.2k | 71.3k | 98.5k | Go | 166.8 |
| [open-telemetry](https://github.com/open-telemetry)  | 1.8M | 62.5k | 61.2k | 28.7k | Go | 148.7 |
| [grpc](https://github.com/grpc)  | 1.7M | 74.1k | 38.8k | 86.8k | C++ | 140.3 |
| [envoyproxy](https://github.com/envoyproxy)  | 1.6M | 47.0k | 24.7k | 28.1k | C++ | 135.3 |
| [konveyor](https://github.com/konveyor)  | 1.6M | 7.1k | 6.3k | 806.0 | JavaScript | 131.0 |
| [openkruise](https://github.com/openkruise)  | 1.4M | 1.3k | 1.2k | 4.0k | JavaScript | 113.5 |
| [brigadecore](https://github.com/brigadecore)  | 877.2k | 6.3k | 2.8k | 3.0k | JavaScript | 73.1 |
| [tikv](https://github.com/tikv)  | 860.8k | 16.8k | 20.1k | 24.1k | Rust | 71.7 |
| [cdk8s-team](https://github.com/cdk8s-team)  | 803.7k | 9.3k | 7.9k | 3.8k | Go | 67.0 |
| [nats-io](https://github.com/nats-io)  | 802.5k | 36.2k | 12.9k | 31.0k | Go | 66.9 |
| [cilium](https://github.com/cilium)  | 681.7k | 31.0k | 21.9k | 23.6k | Go | 56.8 |
| [vitessio](https://github.com/vitessio)  | 671.3k | 36.7k | 11.0k | 15.9k | Go | 56.0 |
| [fluent](https://github.com/fluent)  | 639.2k | 38.0k | 12.0k | 24.2k | C | 53.3 |
| [kubevirt](https://github.com/kubevirt)  | 637.7k | 60.6k | 22.6k | 5.5k | Go | 53.1 |
| [cubefs](https://github.com/cubefs)  | 623.9k | 3.0k | 1.5k | 3.6k | Go | 52.0 |
| [istio](https://github.com/istio)  | 583.6k | 47.1k | 62.4k | 38.5k | Go | 48.6 |
| [falcosecurity](https://github.com/falcosecurity)  | 578.3k | 20.8k | 6.8k | 7.6k | C Header | 48.2 |
| [openebs](https://github.com/openebs)  | 524.7k | 15.5k | 9.6k | 10.1k | Go | 43.7 |
| [argoproj](https://github.com/argoproj)  | 517.8k | 18.1k | 16.9k | 34.2k | Go | 43.1 |
| [pravega](https://github.com/pravega)  | 472.7k | 6.4k | 5.5k | 2.5k | Java | 39.4 |
| [kubevela](https://github.com/kubevela)  | 448.4k | 8.3k | 6.8k | 5.6k | Go | 37.4 |
| [knative](https://github.com/knative)  | 432.8k | 30.6k | 32.3k | 12.4k | Go | 36.1 |
| [athenz](https://github.com/athenz)  | 408.2k | 3.5k | 2.1k | 771.0 | Java | 34.0 |
| [opencurve](https://github.com/opencurve)  | 406.3k | 5.6k | 1.9k | 1.9k | C++ | 33.9 |
| [AthenZ](https://github.com/AthenZ)  | 405.6k | 3.5k | 2.0k | 762.0 | Java | 33.8 |
| [backstage](https://github.com/backstage)  | 403.5k | 40.5k | 12.9k | 21.8k | TypeScript | 33.6 |
| [pixie-io](https://github.com/pixie-io)  | 378.2k | 13.2k | 1.9k | 4.9k | C++ | 31.5 |
| [prometheus](https://github.com/prometheus)  | 377.3k | 30.9k | 19.1k | 89.5k | Go | 31.4 |
| [containerd](https://github.com/containerd)  | 353.4k | 23.2k | 11.7k | 24.7k | Go | 29.4 |
| [werf](https://github.com/werf)  | 349.4k | 16.3k | 5.8k | 4.7k | JavaScript | 29.1 |
| [open-policy-agent](https://github.com/open-policy-agent)  | 343.0k | 8.5k | 7.2k | 15.4k | Go | 28.6 |
| [inclavare-containers](https://github.com/inclavare-containers)  | 341.1k | 3.3k | 1.1k | 569.0 | C Header | 28.4 |
| [dapr](https://github.com/dapr)  | 335.5k | 17.4k | 10.7k | 26.0k | Go | 28.0 |
| [operator-framework](https://github.com/operator-framework)  | 312.6k | 21.6k | 15.5k | 14.0k | Go | 26.1 |
| [chaosblade-io](https://github.com/chaosblade-io)  | 310.9k | 1.3k | 979.0 | 6.2k | JavaScript | 25.9 |
| [parallaxsecond](https://github.com/parallaxsecond)  | 289.8k | 3.5k | 1.4k | 552.0 | Rust | 24.1 |
| [kyverno](https://github.com/kyverno)  | 281.4k | 8.9k | 5.3k | 4.0k | Go | 23.4 |
| [open-cluster-management-io](https://github.com/open-cluster-management-io)  | 276.0k | 7.0k | 3.5k | 1.9k | Go | 23.0 |
| [antrea-io](https://github.com/antrea-io)  | 262.8k | 3.5k | 3.9k | 1.5k | Go | 21.9 |
| [keptn](https://github.com/keptn)  | 256.8k | 11.8k | 8.4k | 1.9k | Go | 21.4 |
| [inspektor-gadget](https://github.com/inspektor-gadget)  | 253.5k | 2.9k | 1.0k | 1.3k | C Header | 21.1 |
| [strimzi](https://github.com/strimzi)  | 252.3k | 7.1k | 6.0k | 4.3k | Java | 21.0 |
| [jaegertracing](https://github.com/jaegertracing)  | 250.5k | 8.7k | 7.6k | 23.7k | Go | 20.9 |
| [goharbor](https://github.com/goharbor)  | 242.5k | 16.9k | 10.4k | 21.1k | Go | 20.2 |
| [kumahq](https://github.com/kumahq)  | 240.3k | 7.2k | 7.1k | 3.3k | Go | 20.0 |
| [linkerd](https://github.com/linkerd)  | 237.9k | 13.0k | 13.4k | 18.0k | Rust | 19.8 |
| [spiffe](https://github.com/spiffe)  | 206.9k | 8.6k | 4.0k | 2.9k | Go | 17.2 |
| [litmuschaos](https://github.com/litmuschaos)  | 203.2k | 7.5k | 7.1k | 4.1k | TypeScript | 16.9 |
| [confidential-containers](https://github.com/confidential-containers)  | 201.9k | 18.7k | 1.5k | 383.0 | Rust | 16.8 |
| [fluxcd](https://github.com/fluxcd)  | 200.2k | 27.7k | 9.5k | 19.9k | Go | 16.7 |
| [longhorn](https://github.com/longhorn)  | 197.7k | 10.0k | 6.6k | 5.1k | Go | 16.5 |
| [cloud-custodian](https://github.com/cloud-custodian)  | 196.4k | 4.5k | 4.6k | 4.8k | Python | 16.4 |
| [tremor-rs](https://github.com/tremor-rs)  | 193.6k | 7.3k | 3.6k | 837.0 | Rust | 16.1 |
| [etcd-io](https://github.com/etcd-io)  | 185.3k | 26.4k | 11.3k | 52.3k | Go | 15.4 |
| [carvel-dev](https://github.com/carvel-dev)  | 180.0k | 8.4k | 3.3k | 3.4k | Go | 15.0 |
| [dragonflyoss](https://github.com/dragonflyoss)  | 168.8k | 7.5k | 4.9k | 8.3k | Go | 14.1 |
| [meshery](https://github.com/meshery)  | 154.9k | 30.0k | 9.0k | 2.8k | JavaScript | 12.9 |
| [networkservicemesh](https://github.com/networkservicemesh)  | 150.4k | 37.4k | 30.5k | 660.0 | Go | 12.5 |
| [cert-manager](https://github.com/cert-manager)  | 146.2k | 13.1k | 4.8k | 10.9k | Go | 12.2 |
| [wasmcloud](https://github.com/wasmcloud)  | 145.4k | 4.9k | 2.0k | 1.8k | Rust | 12.1 |
| [thanos-io](https://github.com/thanos-io)  | 144.3k | 4.3k | 4.2k | 12.1k | Go | 12.0 |
| [kubeedge](https://github.com/kubeedge)  | 141.3k | 7.3k | 3.9k | 6.8k | Go | 11.8 |
| [fluid-cloudnative](https://github.com/fluid-cloudnative)  | 136.8k | 2.0k | 2.0k | 1.3k | Go | 11.4 |
| [getporter](https://github.com/getporter)  | 134.4k | 6.1k | 2.5k | 1.1k | Go | 11.2 |
| [rook](https://github.com/rook)  | 132.1k | 24.8k | 7.1k | 11.2k | Go | 11.0 |
| [projectcontour](https://github.com/projectcontour)  | 130.8k | 4.9k | 3.8k | 4.1k | Go | 10.9 |
| [wasmedge](https://github.com/wasmedge)  | 128.4k | 3.3k | 1.5k | 5.8k | Rust | 10.7 |
| [crossplane](https://github.com/crossplane)  | 127.7k | 11.2k | 3.5k | 7.7k | Go | 10.6 |
| [armadaproject](https://github.com/armadaproject)  | 127.7k | 2.2k | 1.8k | 286.0 | Go | 10.6 |
| [keylime](https://github.com/keylime)  | 127.1k | 2.4k | 1.5k | 399.0 | C | 10.6 |
| [nocalhost](https://github.com/nocalhost)  | 123.0k | 6.1k | 2.1k | 1.6k | Go | 10.2 |
| [devfile](https://github.com/devfile)  | 121.0k | 4.3k | 2.0k | 310.0 | Go | 10.1 |
| [cortexproject](https://github.com/cortexproject)  | 120.0k | 4.9k | 3.8k | 5.1k | Go | 10.0 |
| [artifacthub](https://github.com/artifacthub)  | 116.7k | 1.8k | 2.3k | 1.2k | TypeScript | 9.7 |
| [buildpacks](https://github.com/buildpacks)  | 116.7k | 16.1k | 3.5k | 2.8k | Go | 9.7 |
| [openelb](https://github.com/openelb)  | 116.4k | 498.0 | 283.0 | 1.3k | JavaScript | 9.7 |
| [serverless-devs](https://github.com/serverless-devs)  | 115.1k | 3.5k | 456.0 | 1.5k | JavaScript | 9.6 |
| [paralus](https://github.com/paralus)  | 112.6k | 1.6k | 437.0 | 758.0 | Go | 9.4 |
| [chaos-mesh](https://github.com/chaos-mesh)  | 107.2k | 3.3k | 3.2k | 5.9k | Go | 8.9 |
| [kubescape](https://github.com/kubescape)  | 105.7k | 12.3k | 2.0k | 8.3k | Go | 8.8 |
| [emissary-ingress](https://github.com/emissary-ingress)  | 105.7k | 18.0k | 3.3k | 4.0k | Python | 8.8 |
| [kedacore](https://github.com/kedacore)  | 105.0k | 3.9k | 4.2k | 6.8k | Go | 8.8 |
| [karmada-io](https://github.com/karmada-io)  | 100.5k | 4.1k | 2.7k | 3.2k | Go | 8.4 |
| [bfenetworks](https://github.com/bfenetworks)  | 95.8k | 1.4k | 825.0 | 6.0k | Go | 8.0 |
| [openyurtio](https://github.com/openyurtio)  | 95.6k | 1.8k | 1.5k | 1.6k | Go | 8.0 |
| [v6d-io](https://github.com/v6d-io)  | 94.8k | 1.0k | 882.0 | 689.0 | C++ | 7.9 |
| [in-toto](https://github.com/in-toto)  | 92.3k | 9.5k | 1.0k | 1.1k | Python | 7.7 |
| [serverlessworkflow](https://github.com/serverlessworkflow)  | 87.3k | 2.3k | 1.2k | 944.0 | C# | 7.3 |
| [openservicemesh](https://github.com/openservicemesh)  | 81.9k | 5.2k | 4.1k | 2.6k | Go | 6.8 |
| [cloudevents](https://github.com/cloudevents)  | 79.5k | 4.0k | 2.6k | 5.8k | Go | 6.6 |
| [kubeovn](https://github.com/kubeovn)  | 76.6k | 4.5k | 2.0k | 1.5k | Go | 6.4 |
| [coredns](https://github.com/coredns)  | 76.2k | 5.5k | 4.7k | 11.5k | Go | 6.3 |
| [project-zot](https://github.com/project-zot)  | 75.8k | 1.1k | 1.2k | 333.0 | Go | 6.3 |
| [helm](https://github.com/helm)  | 71.5k | 28.9k | 26.9k | 48.8k | Go | 6.0 |
| [cri-o](https://github.com/cri-o)  | 71.3k | 9.2k | 5.8k | 4.5k | Go | 5.9 |
| [submariner-io](https://github.com/submariner-io)  | 69.8k | 10.0k | 9.1k | 2.3k | Go | 5.8 |
| [opencost](https://github.com/opencost)  | 68.6k | 3.7k | 1.4k | 3.6k | Go | 5.7 |
| [volcano-sh](https://github.com/volcano-sh)  | 68.5k | 5.5k | 2.0k | 3.2k | Go | 5.7 |
| [metal3-io](https://github.com/metal3-io)  | 68.4k | 9.4k | 4.6k | 1.0k | Go | 5.7 |
| [notaryproject](https://github.com/notaryproject)  | 64.4k | 3.7k | 1.9k | 3.3k | Go | 5.4 |
| [containerssh](https://github.com/containerssh)  | 62.2k | 1.8k | 2.2k | 2.2k | Go | 5.2 |
| [oras-project](https://github.com/oras-project)  | 61.5k | 2.2k | 1.0k | 1.1k | Go | 5.1 |
| [telepresenceio](https://github.com/telepresenceio)  | 57.4k | 8.0k | 1.7k | 5.7k | Go | 4.8 |
| [external-secrets](https://github.com/external-secrets)  | 56.9k | 2.8k | 1.8k | 5.1k | Go | 4.7 |
| [kubewarden](https://github.com/kubewarden)  | 55.7k | 8.5k | 3.0k | 489.0 | Rust | 4.6 |
| [open-feature](https://github.com/open-feature)  | 53.0k | 5.1k | 3.2k | 754.0 | Go | 4.4 |
| [rkt](https://github.com/rkt)  | 52.4k | 5.6k | 2.5k | 8.9k | Go | 4.4 |
| [trickstercache](https://github.com/trickstercache)  | 52.1k | 1.1k | 458.0 | 1.9k | Go | 4.3 |
| [superedge](https://github.com/superedge)  | 48.2k | 487.0 | 425.0 | 913.0 | Go | 4.0 |
| [virtual-kubelet](https://github.com/virtual-kubelet)  | 48.1k | 2.8k | 1.3k | 4.4k | Go | 4.0 |
| [kudobuilder](https://github.com/kudobuilder)  | 47.9k | 2.2k | 2.0k | 1.8k | Go | 4.0 |
| [theupdateframework](https://github.com/theupdateframework)  | 45.4k | 8.6k | 2.6k | 2.6k | Rust | 3.8 |
| [openfga](https://github.com/openfga)  | 45.2k | 1.2k | 1.2k | 979.0 | Go | 3.8 |
| [kubedl-io](https://github.com/kubedl-io)  | 44.1k | 1.3k | 369.0 | 488.0 | Go | 3.7 |
| [skooner-k8s](https://github.com/skooner-k8s)  | 44.0k | 507.0 | 213.0 | 1.1k | JavaScript | 3.7 |
| [distribution](https://github.com/distribution)  | 42.9k | 4.8k | 1.8k | 7.4k | Go | 3.6 |
| [openfunction](https://github.com/openfunction)  | 41.4k | 1.4k | 915.0 | 1.2k | Go | 3.5 |
| [metallb](https://github.com/metallb)  | 40.9k | 2.3k | 1.3k | 5.7k | Go | 3.4 |
| [tinkerbell](https://github.com/tinkerbell)  | 40.5k | 6.8k | 2.4k | 1.6k | Go | 3.4 |
| [capsule-rs](https://github.com/capsule-rs)  | 40.4k | 252.0 | 140.0 | 356.0 | Rust | 3.4 |
| [devstream-io](https://github.com/devstream-io)  | 39.3k | 2.9k | 1.1k | 747.0 | Go | 3.3 |
| [vscode-kubernetes-tools](https://github.com/vscode-kubernetes-tools)  | 39.3k | 634.0 | 613.0 | 602.0 | TypeScript | 3.3 |
| [k3s-io](https://github.com/k3s-io)  | 38.5k | 3.5k | 3.3k | 25.3k | Go | 3.2 |
| [aeraki-mesh](https://github.com/aeraki-mesh)  | 36.7k | 988.0 | 368.0 | 866.0 | Go | 3.1 |
| [kubearmor](https://github.com/kubearmor)  | 35.8k | 5.5k | 2.1k | 718.0 | Go | 3.0 |
| [containernetworking](https://github.com/containernetworking)  | 34.8k | 2.8k | 1.2k | 6.5k | Go | 2.9 |
| [clusternet](https://github.com/clusternet)  | 33.4k | 948.0 | 545.0 | 1.1k | Go | 2.8 |
| [sealerio](https://github.com/sealerio)  | 32.7k | 1.2k | 1.3k | 1.8k | Go | 2.7 |
| [kube-rs](https://github.com/kube-rs)  | 31.6k | 3.0k | 775.0 | 2.2k | Rust | 2.6 |
| [dexidp](https://github.com/dexidp)  | 29.3k | 2.6k | 1.8k | 8.0k | Go | 2.4 |
| [clusterpedia-io](https://github.com/clusterpedia-io)  | 29.1k | 1.1k | 541.0 | 596.0 | Go | 2.4 |
| [opentracing](https://github.com/opentracing)  | 27.4k | 2.7k | 1.4k | 10.1k | Java | 2.3 |
| [krustlet](https://github.com/krustlet)  | 23.6k | 2.4k | 530.0 | 3.4k | Rust | 2.0 |
| [curiefense](https://github.com/curiefense)  | 22.3k | 2.9k | 817.0 | 617.0 | Rust | 1.9 |
| [fabedge](https://github.com/fabedge)  | 21.6k | 941.0 | 429.0 | 482.0 | Go | 1.8 |
| [project-akri](https://github.com/project-akri)  | 20.6k | 701.0 | 448.0 | 946.0 | Rust | 1.7 |
| [piraeusdatastore](https://github.com/piraeusdatastore)  | 19.7k | 1.3k | 522.0 | 748.0 | Go | 1.6 |
| [k8gb-io](https://github.com/k8gb-io)  | 19.4k | 1.1k | 918.0 | 571.0 | Go | 1.6 |
| [foniod](https://github.com/foniod)  | 18.1k | 1.7k | 493.0 | 2.0k | Rust | 1.5 |
| [lima-vm](https://github.com/lima-vm)  | 17.4k | 2.1k | 934.0 | 11.7k | Go | 1.4 |
| [schemahero](https://github.com/schemahero)  | 17.0k | 1.4k | 773.0 | 765.0 | Go | 1.4 |
| [hexa-org](https://github.com/hexa-org)  | 16.5k | 491.0 | 118.0 | 70.0 | Go | 1.4 |
| [k8up-io](https://github.com/k8up-io)  | 13.5k | 2.0k | 699.0 | 412.0 | Go | 1.1 |
| [carina-io](https://github.com/carina-io)  | 12.9k | 749.0 | 118.0 | 588.0 | Go | 1.1 |
| [servicemeshinterface](https://github.com/servicemeshinterface)  | 11.6k | 966.0 | 372.0 | 1.2k | Go | 970.0m |
| [kuberhealthy](https://github.com/kuberhealthy)  | 10.8k | 3.1k | 721.0 | 1.7k | Go | 900.0m |
| [kube-vip](https://github.com/kube-vip)  | 9.7k | 916.0 | 320.0 | 1.3k | Go | 800.0m |
| [ko-build](https://github.com/ko-build)  | 9.5k | 793.0 | 721.0 | 5.8k | Go | 790.0m |
| [devspace-cloud](https://github.com/devspace-cloud)  | 8.7k | 348.0 | 83.0 | 113.0 | Go | 720.0m |
| [tellerops](https://github.com/tellerops)  | 7.7k | 179.0 | 80.0 | 1.6k | Go | 640.0m |
| [cni-genie](https://github.com/cni-genie)  | 5.9k | 484.0 | 132.0 | 499.0 | Go | 490.0m |
| [merbridge](https://github.com/merbridge)  | 4.8k | 271.0 | 279.0 | 586.0 | Go | 400.0m |
| [opcr-io](https://github.com/opcr-io)  | 3.6k | 597.0 | 153.0 | 181.0 | Go | 300.0m |
| [krator-rs](https://github.com/krator-rs)  | 2.9k | 498.0 | 33.0 | 113.0 | Rust | 250.0m |
| [OpenObservability](https://github.com/OpenObservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [openobservability](https://github.com/openobservability)  | 2.3k | 251.0 | 117.0 | 2.0k | Go | 190.0m |
| [service-mesh-performance](https://github.com/service-mesh-performance)  | 2.2k | 797.0 | 217.0 | 237.0 | JavaScript | 180.0m |
| [kubereboot](https://github.com/kubereboot)  | 2.0k | 997.0 | 558.0 | 1.8k | Go | 170.0m |
| [gitops-working-group](https://github.com/gitops-working-group)  | 0.0 | 43.0 | 57.0 | 552.0 | N/A | 0.0 |


### OTEL Instrumentations
| Project | SLOC | Commits | PRs | Stars | Effort |
|---------|------|---------|-----|-------|--------|
| [**OTEL**](https://opentelemetry.io/docs/instrumentation/)  | 1.1M | 32.6k | 31.0k | 16.8k | 89.34 |
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
| [java-examples](https://github.com/open-telemetry/opentelemetry-java-examples/)  | 1.3k | 123.0 | 124.0 | 90.0 | 0.11 |
| [php-instrumentation](https://github.com/open-telemetry/opentelemetry-php-instrumentation/)  | 556.0 | 30.0 | 30.0 | 24.0 | 0.05 |
| [configuration](https://github.com/open-telemetry/opentelemetry-configuration/)  | 0.0 | 9.0 | 8.0 | 7.0 | 0 |

