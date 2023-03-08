# oss-estimator

Set of tools to get info about GitHub project that is needed to estimate the cost of running it.

## Tools

| Tool                                                  | Description                                                                 |
|-------------------------------------------------------|-----------------------------------------------------------------------------|
| [estimator-dl](../estimator-dl/main.go)               | Download from [gharchive.org](https://gharchive.org)                        |
| [estimator-sloc](../estimator-sloc/main.go)           | SLOC count using [scc](https://github.com/boyter/scc/) (should be in $PATH) |
| [estimator-list](../estimator-list/main.go)           | Concurrently fetch popular oss repos and stat for them                      |
| [estimator-aggregate](../estimator-aggregate/main.go) | Aggregate stats                                                             |
| [estimator-gen](../estimator-gen/main.go)             | Generate README.md                                                          |


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

| Organization | SLOC | Commits | PRs | Stars |
|--------------|------|---------|-----|-------|
| [torvalds](https://github.com/torvalds)  | 19.1M | 1.1M | 761 | 147.5K |
| [kubernetes](https://github.com/kubernetes)  | 5.4M | 498.4K | 261.5K | 341.3K |
| [tensorflow](https://github.com/tensorflow)  | 3.3M | 144.4K | 22.7K | 171.8K |
| [ClickHouse](https://github.com/ClickHouse)  | 2.3M | 166.2K | 34.8K | 33.8K |
| [envoyproxy](https://github.com/envoyproxy)  | 1.8M | 46.7K | 24.3K | 28.1K |
| [open-telemetry](https://github.com/open-telemetry)  | 1.7M | 61.8K | 61.1K | 28.6K |
| [grpc](https://github.com/grpc)  | 1.6M | 74.1K | 38.7K | 86.8K |
| [grafana](https://github.com/grafana)  | 1.6M | 76.6K | 50.2K | 103.1K |
| [rust-lang](https://github.com/rust-lang)  | 1.5M | 219.1K | 61.1K | 78.6K |
| [python](https://github.com/python)  | 1.3M | 116.2K | 39.4K | 51.6K |
| [docker](https://github.com/docker)  | 1.2M | 173.7K | 30.5K | 146.4K |
| [golang](https://github.com/golang)  | 1.6M | 55.7K | 2.8K | 109.1K |
| [cilium](https://github.com/cilium)  | 867.3K | 47.5K | 21.8K | 23.6K |
| [m3db](https://github.com/m3db)  | 736.7K | 10.4K | 5.7K | 4.6K |
| [istio](https://github.com/istio)  | 583.5K | 47.4K | 62.3K | 38.4K |
| [openebs](https://github.com/openebs)  | 524.6K | 15.5K | 9.6K | 10.1K |
| [apache](https://github.com/apache)  | 420.4K | 22.2K | 521 | 5.6K |
| [prometheus](https://github.com/prometheus)  | 377.3K | 30.8K | 19.1K | 89.5K |
| [facebook](https://github.com/facebook)  | 357.5K | 15.5K | 13.2K | 203.4K |
| [pixie-io](https://github.com/pixie-io)  | 352.4K | 11.7K | 549 | 4.3K |
| [LINBIT](https://github.com/LINBIT)  | 310.8K | 4.3K | 11 | 649 |
| [vectordotdev](https://github.com/vectordotdev)  | 308.9K | 14.3K | 10.1K | 14.3K |
| [uber](https://github.com/uber)  | 253.7K | 1.5K | 248 | 5.9K |
| [siderolabs](https://github.com/siderolabs)  | 232.7K | 7.1K | 9.6K | 4.7K |
| [kata-containers](https://github.com/kata-containers)  | 223.8K | 22.7K | 10.2K | 7.0K |
| [Netflix](https://github.com/Netflix)  | 208.1K | 4.1K | 2.2K | 551 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 205.9K | 8.1K | 2.5K | 10.9K |
| [etcd-io](https://github.com/etcd-io)  | 185.2K | 26.3K | 11.2K | 52.2K |
| [containers](https://github.com/containers)  | 174.4K | 18.2K | 10.1K | 17.8K |
| [VKCOM](https://github.com/VKCOM)  | 101.4K | 9.8K | 3.8K | 877 |
| [gotd](https://github.com/gotd)  | 74.3K | 6.2K | 2.4K | 1.1K |
| [helm](https://github.com/helm)  | 71.5K | 28.9K | 26.8K | 48.7K |
| [vuejs](https://github.com/vuejs)  | 60.2K | 3.5K | 2.3K | 202.6K |
| [go-faster](https://github.com/go-faster)  | 50.7K | 4.7K | 670 | 151 |
| [vitalif](https://github.com/vitalif)  | 40.8K | 1.2K | 12 | 69 |
| [ogen-go](https://github.com/ogen-go)  | 38.1K | 3.8K | 1.5K | 472 |
| [falcosecurity](https://github.com/falcosecurity)  | 13.9K | 3.5K | 1.4K | 5.6K |


| Project                                        | SLOC | Commits | PRs | Stars |
|------------------------------------------------|------|---------|-----|-------|
| [torvalds/linux](https://github.com/torvalds/linux/)  | 19.1M | 1.1M | 761 | 147.5K |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3.3M | 144.4K | 22.7K | 171.8K |
| [rust-lang/rust](https://github.com/rust-lang/rust/)  | 1.5M | 219.1K | 61.1K | 78.6K |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1.4M | 114.4K | 73.5K | 96.4K |
| [python/cpython](https://github.com/python/cpython/)  | 1.3M | 116.2K | 39.4K | 51.6K |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1.1M | 109.7K | 32.2K | 27.4K |
| [golang/go](https://github.com/golang/go/)  | 1.6M | 55.7K | 2.8K | 109.1K |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889.8K | 41.3K | 33.5K | 54.2K |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 807.5K | 17.3K | 17.1K | 21.5K |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 514.8K | 52.6K | 21.6K | 37.2K |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 512.4K | 96 | 103 | 22 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 504.5K | 8.4K | 454 | 205 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477.3K | 4.2K | 3.5K | 4.3K |
| [ClickHouse/grpc](https://github.com/ClickHouse/grpc/)  | 397.2K | 47.6K | 13 | 1 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389.2K | 9.6K | 16.1K | 1.5K |
| [facebook/react](https://github.com/facebook/react/)  | 357.5K | 15.5K | 13.2K | 203.4K |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 352.4K | 11.7K | 549 | 4.3K |
| [istio/istio](https://github.com/istio/istio/)  | 339.8K | 19.8K | 26.4K | 32.5K |
| [kubernetes/autoscaler](https://github.com/kubernetes/autoscaler/)  | 332.8K | 6.4K | 3.7K | 6.5K |
| [cilium/cilium](https://github.com/cilium/cilium/)  | 324.8K | 22.2K | 16.9K | 14.5K |
| [docker/docker-ce](https://github.com/docker/docker-ce/)  | 324.4K | 54.3K | 662 | 5.5K |
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310.8K | 4.3K | 11 | 649 |
| [apache/mesos](https://github.com/apache/mesos/)  | 305.6K | 18.1K | 450 | 5.2K |
| [docker/labs](https://github.com/docker/labs/)  | 304.4K | 718 | 398 | 11.1K |
| [envoyproxy/envoy-website](https://github.com/envoyproxy/envoy-website/)  | 299.3K | 428 | 252 | 33 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284.4K | 1.5K | 1.3K | 465 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274.6K | 9.3K | 9.3K | 12.9K |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 270.5K | 51.9K | 24.6K | 3.5K |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 264.4K | 1.6K | 36 | 24 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241.7K | 4.7K | 5.1K | 18.4K |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235.1K | 5.7K | 6.8K | 10.4K |
| [uber/peloton](https://github.com/uber/peloton/)  | 216.3K | 705 | 10 | 582 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194.1K | 181 | 125 | 1.1K |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186.6K | 1 | 2 | 2 |
| [cilium/busybox](https://github.com/cilium/busybox/)  | 185.7K | 16.6K | 0 | 0 |
| [containers/podman](https://github.com/containers/podman/)  | 174.4K | 18.2K | 10.1K | 17.8K |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173.3K | 8.8K | 5.3K | 2.7K |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167.8K | 8.8K | 5.6K | 1.2K |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167.6K | 19.3K | 10.3K | 14.7K |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157.3K | 1.6K | 1.2K | 319 |
| [ClickHouse/ssl](https://github.com/ClickHouse/ssl/)  | 156.6K | 16 | 2 | 2 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155.1K | 10.9K | 6.7K | 47.8K |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144.5K | 5.9K | 1.8K | 8.1K |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143.2K | 4.4K | 3.9K | 17.6K |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141.8K | 927 | 1.7K | 224 |
| [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize/)  | 139.4K | 6.3K | 3.2K | 9.4K |
| [kubernetes/apiserver](https://github.com/kubernetes/apiserver/)  | 131.3K | 6.1K | 25 | 494 |
| [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers/)  | 130.8K | 10.1K | 3.7K | 3.8K |
| [siderolabs/talos](https://github.com/siderolabs/talos/)  | 130.1K | 3.8K | 5.1K | 3.6K |
| [etcd-io/etcd](https://github.com/etcd-io/etcd/)  | 125.6K | 19.3K | 9.1K | 42.8K |
| [apache/aurora](https://github.com/apache/aurora/)  | 114.7K | 4.9K | 71 | 628 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 113.6K | 2.8K | 60 | 10 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 106.9K | 2.2K | 1.8K | 218 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105.9K | 1.4K | 1.7K | 980 |
| [openebs/maya](https://github.com/openebs/maya/)  | 100.7K | 1.7K | 1.6K | 180 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96.3K | 3.2K | 3.6K | 1.4K |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 95.9K | 2.9K | 313 | 2.2K |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90.6K | 20.9K | 7.6K | 26.1K |
| [ClickHouse/UnixODBC](https://github.com/ClickHouse/UnixODBC/)  | 88.9K | 5 | 2 | 0 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84.4K | 3.4K | 3.1K | 210 |
| [docker/cli](https://github.com/docker/cli/)  | 78.9K | 8.4K | 2.6K | 3.9K |
| [kubernetes/ingress-gce](https://github.com/kubernetes/ingress-gce/)  | 78.2K | 4.5K | 1.5K | 1.1K |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77.2K | 3.6K | 2.1K | 244 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76.4K | 630 | 1.1K | 3.2K |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73.7K | 6.5K | 1.2K | 628 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73.7K | 1.7K | 634 | 1.1K |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73.6K | 3.8K | 2.6K | 543 |
| [grafana/k6](https://github.com/grafana/k6/)  | 73.4K | 5.5K | 1.4K | 19.6K |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72.3K | 4.8K | 5.2K | 2.4K |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71.8K | 633 | 232 | 1.6K |
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65.6K | 3.8K | 204 | 7.3K |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63.9K | 2.7K | 1.4K | 2.1K |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63.8K | 859 | 941 | 3.5K |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61.5K | 236 | 214 | 120 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61.1K | 1.7K | 2.5K | 1.7K |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60.9K | 1.2K | 447 | 27 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60.5K | 3.1K | 1.8K | 6.1K |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60.2K | 3.5K | 2.3K | 202.6K |
| [kubernetes/legacy-cloud-providers](https://github.com/kubernetes/legacy-cloud-providers/)  | 60.2K | 1.8K | 0 | 47 |
| [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/)  | 60.1K | 2.5K | 30 | 661 |
| [envoyproxy/envoy-mobile](https://github.com/envoyproxy/envoy-mobile/)  | 57.6K | 1.7K | 2.2K | 550 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56.9K | 2.2K | 2.7K | 2.2K |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56.8K | 6.9K | 4.9K | 14.5K |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 56.9K | 1.8K | 1.6K | 286 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 54.1K | 1.8K | 1.1K | 457 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53.4K | 4.4K | 37 | 11 |
| [gotd/td](https://github.com/gotd/td/)  | 52.6K | 3.8K | 788 | 1.5K |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51.3K | 1.5K | 990 | 420 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51.2K | 3.1K | 8 | 206 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51.9K | 1.9K | 2.3K | 3.6K |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50.8K | 2.4K | 997 | 232 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50.6K | 233 | 194 | 9 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector/)  | 50.3K | 4.4K | 5.9K | 2.7K |
| [kubernetes-sigs/multi-tenancy](https://github.com/kubernetes-sigs/multi-tenancy/)  | 48.9K | 2.3K | 1.6K | 930 |
| [envoyproxy/pytooling](https://github.com/envoyproxy/pytooling/)  | 48.9K | 615 | 619 | 6 |
| [grpc/grpc-swift](https://github.com/grpc/grpc-swift/)  | 48.3K | 1.5K | 981 | 1.7K |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46.5K | 4.4K | 5 | 1 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46.4K | 241 | 358 | 20 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45.7K | 790 | 309 | 38 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45.7K | 1.7K | 1.2K | 413 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45.2K | 1.8K | 543 | 113 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44.8K | 4.4K | 4.9K | 12.3K |
| [istio/old_mixer_repo](https://github.com/istio/old_mixer_repo/)  | 42.5K | 741 | 1.9K | 68 |
| [grpc/grpc-node](https://github.com/grpc/grpc-node/)  | 42.3K | 4.2K | 1.3K | 3.8K |
| [open-telemetry/opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python/)  | 42.2K | 1.4K | 1.7K | 1.2K |
| [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)  | 42.1K | 2.2K | 1.3K | 1.8K |
| [open-telemetry/opentelemetry-cpp-c~](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42.5K | 157 | 180 | 78 |
| [helm/helm](https://github.com/helm/helm/)  | 41.7K | 6.8K | 4.8K | 23.8K |
| [open-telemetry/opentelemetry-ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41.4K | 267 | 110 | 85 |
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40.8K | 1.2K | 12 | 69 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 40.8K | 52 | 3 | 5 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40.4K | 9.5K | 2.8K | 757 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39.8K | 177 | 142 | 13 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38.1K | 2.9K | 1.8K | 6.1K |
| [uber/kraken](https://github.com/uber/kraken/)  | 37.4K | 867 | 238 | 5.3K |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37.3K | 799 | 283 | 321 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36.6K | 3.2K | 700 | 454 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36.3K | 35 | 7 | 1 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36.3K | 298 | 358 | 83 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35.4K | 662 | 313 | 278 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34.2K | 1.2K | 459 | 2.3K |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34.1K | 13 | 0 | 6 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33.6K | 1.4K | 420 | 78 |
| [open-telemetry/opentelemetry-cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33.5K | 1.5K | 1.1K | 431 |
| [kubernetes/perf-tests](https://github.com/kubernetes/perf-tests/)  | 33.3K | 3.1K | 1.8K | 764 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33.3K | 732 | 850 | 230 |
| [prometheus/alertmanager](https://github.com/prometheus/alertmanager/)  | 31.8K | 2.8K | 1.6K | 5.5K |
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 30.9K | 508 | 591 | 519 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30.9K | 1.2K | 308 | 232 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30.9K | 18 | 4 | 0 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30.5K | 574 | 597 | 997 |
| [docker/machine](https://github.com/docker/machine/)  | 29.9K | 3.4K | 1.9K | 6.5K |
| [cilium/ebpf](https://github.com/cilium/ebpf/)  | 29.9K | 1.2K | 686 | 4.1K |
| [m3db/m3cluster](https://github.com/m3db/m3cluster/)  | 29.6K | 238 | 227 | 21 |
| [open-telemetry/opentelemetry-log-c~](https://github.com/open-telemetry/opentelemetry-log-collection/)  | 28.5K | 267 | 373 | 92 |
| [envoyproxy/nighthawk](https://github.com/envoyproxy/nighthawk/)  | 27.8K | 578 | 655 | 272 |
| [kubernetes-sigs/alibaba-cloud-csi-~](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/)  | 27.6K | 1.3K | 590 | 444 |
| [kubernetes/release](https://github.com/kubernetes/release/)  | 27.3K | 5.1K | 2.3K | 436 |
| [docker/docker-py](https://github.com/docker/docker-py/)  | 27.3K | 3.2K | 1.5K | 6.1K |
| [open-telemetry/opentelemetry-go-co~](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27.2K | 1.4K | 3.2K | 684 |
| [cilium/hubble-ui](https://github.com/cilium/hubble-ui/)  | 26.7K | 385 | 426 | 237 |
| [ClickHouse/libhdfs3](https://github.com/ClickHouse/libhdfs3/)  | 26.7K | 64 | 34 | 23 |
| [ClickHouse/clickhouse-cpp](https://github.com/ClickHouse/clickhouse-cpp/)  | 26.6K | 673 | 172 | 215 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26.3K | 562 | 644 | 80 |
| [open-telemetry/opentelemetry-swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25.9K | 618 | 252 | 126 |
| [envoyproxy/java-control-plane](https://github.com/envoyproxy/java-control-plane/)  | 25.4K | 235 | 206 | 246 |
| [grafana/grafana-plugin-sdk-go](https://github.com/grafana/grafana-plugin-sdk-go/)  | 25.3K | 415 | 531 | 157 |
| [kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)  | 25.3K | 642 | 349 | 733 |
| [kubernetes-sigs/kubefed](https://github.com/kubernetes-sigs/kubefed/)  | 25.4K | 2.7K | 978 | 2.4K |
| [docker/compose-cli](https://github.com/docker/compose-cli/)  | 24.9K | 3.2K | 1.2K | 910 |
| [grafana/terraform-provider-grafana](https://github.com/grafana/terraform-provider-grafana/)  | 24.6K | 697 | 522 | 312 |
| [siderolabs/theila](https://github.com/siderolabs/theila/)  | 24.3K | 100 | 114 | 45 |
| [kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/)  | 23.6K | 2.5K | 1.2K | 4.3K |
| [open-telemetry/opentelemetry-opera~](https://github.com/open-telemetry/opentelemetry-operator/)  | 23.6K | 781 | 1.1K | 675 |
| [openebs/elves](https://github.com/openebs/elves/)  | 23.2K | 234 | 47 | 17 |
| [openebs/node-disk-manager](https://github.com/openebs/node-disk-manager/)  | 23.2K | 451 | 533 | 170 |
| [istio/old_pilot_repo](https://github.com/istio/old_pilot_repo/)  | 22.9K | 688 | 1.2K | 138 |
| [istio/proxy](https://github.com/istio/proxy/)  | 22.4K | 2.4K | 4.2K | 700 |
| [kubernetes-sigs/azuredisk-csi-driv~](https://github.com/kubernetes-sigs/azuredisk-csi-driver/)  | 21.7K | 2.7K | 1.3K | 113 |
| [envoyproxy/gateway](https://github.com/envoyproxy/gateway/)  | 21.6K | 539 | 632 | 886 |
| [istio/operator](https://github.com/istio/operator/)  | 21.2K | 508 | 780 | 174 |
| [kubernetes-sigs/structured-merge-d~](https://github.com/kubernetes-sigs/structured-merge-diff/)  | 20.9K | 495 | 208 | 74 |
| [prometheus/procfs](https://github.com/prometheus/procfs/)  | 20.8K | 572 | 372 | 618 |
| [prometheus/client_golang](https://github.com/prometheus/client_golang/)  | 20.7K | 1.4K | 741 | 4.4K |
| [docker/buildx](https://github.com/docker/buildx/)  | 20.6K | 1.4K | 722 | 2.4K |
| [kubernetes/cloud-provider-vsphere](https://github.com/kubernetes/cloud-provider-vsphere/)  | 20.4K | 1.8K | 472 | 189 |
| [kubernetes/cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack/)  | 20.2K | 2.6K | 1.3K | 507 |
| [open-telemetry/opentelemetry-ruby-~](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 19.9K | 895 | 241 | 29 |
| [prometheus/node_exporter](https://github.com/prometheus/node_exporter/)  | 19.2K | 1.9K | 1.4K | 8.5K |
| [envoyproxy/data-plane-api](https://github.com/envoyproxy/data-plane-api/)  | 19.2K | 2.3K | 540 | 510 |
| [kubernetes-sigs/descheduler](https://github.com/kubernetes-sigs/descheduler/)  | 19.7K | 1.3K | 667 | 3.2K |
| [cilium/cilium-olm](https://github.com/cilium/cilium-olm/)  | 19.2K | 341 | 55 | 11 |
| [ClickHouse/libgsasl](https://github.com/ClickHouse/libgsasl/)  | 18.9K | 21 | 8 | 0 |
| [grpc/grpc-dart](https://github.com/grpc/grpc-dart/)  | 18.7K | 228 | 297 | 759 |
| [cilium/cilium-cli](https://github.com/cilium/cilium-cli/)  | 18.3K | 1.3K | 1.1K | 205 |
| [kubernetes-sigs/gcp-filestore-csi-~](https://github.com/kubernetes-sigs/gcp-filestore-csi-driver/)  | 17.8K | 593 | 403 | 68 |
| [istio/tools](https://github.com/istio/tools/)  | 17.4K | 1.7K | 2.3K | 302 |
| [openebs/dynamic-nfs-provisioner](https://github.com/openebs/dynamic-nfs-provisioner/)  | 17.3K | 112 | 109 | 103 |
| [docker/compose](https://github.com/docker/compose/)  | 17.3K | 4.1K | 3.5K | 28.7K |
| [docker/compose-on-kubernetes](https://github.com/docker/compose-on-kubernetes/)  | 17.2K | 236 | 114 | 1.4K |
| [kubernetes-sigs/etcdadm](https://github.com/kubernetes-sigs/etcdadm/)  | 16.8K | 1.2K | 248 | 652 |
| [kubernetes/api](https://github.com/kubernetes/api/)  | 16.8K | 6.7K | 10 | 514 |
| [ClickHouse/libcxx](https://github.com/ClickHouse/libcxx/)  | 16.5K | 38 | 14 | 0 |
| [open-telemetry/opentelemetry-ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16.4K | 763 | 919 | 375 |
| [ClickHouse/clickhouse-odbc](https://github.com/ClickHouse/clickhouse-odbc/)  | 16.2K | 1.1K | 172 | 222 |
| [kubernetes-sigs/aws-ebs-csi-driver](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/)  | 16.2K | 1.6K | 882 | 720 |
| [ClickHouse/boost](https://github.com/ClickHouse/boost/)  | 16.1K | 83 | 28 | 1 |
| [ClickHouse/antlr4-runtime](https://github.com/ClickHouse/antlr4-runtime/)  | 16.1K | 306 | 0 | 1 |
| [kubernetes-sigs/boskos](https://github.com/kubernetes-sigs/boskos/)  | 15.9K | 923 | 129 | 93 |
| [openebs/dynamic-localpv-provisioner](https://github.com/openebs/dynamic-localpv-provisioner/)  | 15.8K | 154 | 123 | 88 |
| [kubernetes-sigs/gcp-compute-persis~](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver/)  | 15.7K | 1.3K | 824 | 133 |
| [prometheus/client_java](https://github.com/prometheus/client_java/)  | 15.6K | 578 | 490 | 1.9K |
| [grafana/carbon-relay-ng](https://github.com/grafana/carbon-relay-ng/)  | 15.5K | 1.1K | 258 | 454 |
| [go-faster/yamlx](https://github.com/go-faster/yamlx/)  | 15.4K | 540 | 38 | 5 |
| [kubernetes-sigs/azurefile-csi-driv~](https://github.com/kubernetes-sigs/azurefile-csi-driver/)  | 15.2K | 2.4K | 886 | 113 |
| [ClickHouse/boost-extra](https://github.com/ClickHouse/boost-extra/)  | 15.2K | 34 | 0 | 0 |
| [kubernetes-sigs/promo-tools](https://github.com/kubernetes-sigs/promo-tools/)  | 15.1K | 1.6K | 602 | 125 |
| [etcd-io/raft](https://github.com/etcd-io/raft/)  | 15.2K | 1.1K | 19 | 139 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-openstack/)  | 15.1K | 1.4K | 977 | 220 |
| [ClickHouse/ch-go](https://github.com/ClickHouse/ch-go/)  | 14.9K | 1.9K | 230 | 230 |
| [kubernetes/cloud-provider-aws](https://github.com/kubernetes/cloud-provider-aws/)  | 14.9K | 975 | 401 | 294 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-ibmcloud/)  | 14.5K | 727 | 840 | 55 |
| [openebs/jiva](https://github.com/openebs/jiva/)  | 14.5K | 804 | 337 | 133 |
| [kubernetes-sigs/kind](https://github.com/kubernetes-sigs/kind/)  | 14.5K | 3.8K | 1.5K | 11.2K |
| [docker/libcompose](https://github.com/docker/libcompose/)  | 14.3K | 673 | 361 | 587 |
| [docker/go-docker](https://github.com/docker/go-docker/)  | 14.3K | 7 | 4 | 179 |
| [kubernetes/cloud-provider](https://github.com/kubernetes/cloud-provider/)  | 14.1K | 1.4K | 0 | 182 |
| [VictoriaMetrics/grafana-datasource](https://github.com/VictoriaMetrics/grafana-datasource/)  | 14.1K | 189 | 29 | 23 |
| [falcosecurity/falco](https://github.com/falcosecurity/falco/)  | 13.9K | 3.5K | 1.4K | 5.6K |
| [kubernetes/component-base](https://github.com/kubernetes/component-base/)  | 13.7K | 1.2K | 0 | 84 |
| [kubernetes-sigs/node-feature-disco~](https://github.com/kubernetes-sigs/node-feature-discovery/)  | 13.6K | 1.5K | 832 | 541 |
| [docker/awesome-compose](https://github.com/docker/awesome-compose/)  | 13.5K | 267 | 255 | 21.8K |
| [prometheus/promlens](https://github.com/prometheus/promlens/)  | 13.5K | 78 | 81 | 307 |
| [grafana/har-to-k6](https://github.com/grafana/har-to-k6/)  | 13.4K | 627 | 84 | 77 |
| [istio/bots](https://github.com/istio/bots/)  | 13.4K | 638 | 652 | 10 |
| [m3db/m3coordinator](https://github.com/m3db/m3coordinator/)  | 13.2K | 68 | 54 | 4 |
| [m3db/m3ninx](https://github.com/m3db/m3ninx/)  | 13.2K | 67 | 77 | 3 |
| [envoyproxy/playground](https://github.com/envoyproxy/playground/)  | 13.8K | 192 | 168 | 7 |
| [go-faster/hx](https://github.com/go-faster/hx/)  | 13.3K | 1.4K | 12 | 4 |
| [prometheus/common](https://github.com/prometheus/common/)  | 12.9K | 497 | 362 | 224 |
| [open-telemetry/opentelemetry-php-c~](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12.7K | 320 | 112 | 22 |
| [m3db/m3x](https://github.com/m3db/m3x/)  | 12.6K | 208 | 206 | 17 |
| [etcd-io/bbolt](https://github.com/etcd-io/bbolt/)  | 12.5K | 1.1K | 220 | 6.2K |
| [siderolabs/sidero](https://github.com/siderolabs/sidero/)  | 12.3K | 366 | 839 | 257 |
| [docker/app](https://github.com/docker/app/)  | 12.3K | 1.5K | 650 | 1.5K |
| [grafana/ksonnet](https://github.com/grafana/ksonnet/)  | 12.1K | 434 | 0 | 0 |
| [m3db/m3db-operator](https://github.com/m3db/m3db-operator/)  | 11.9K | 230 | 247 | 134 |
| [openebs/openebsctl](https://github.com/openebs/openebsctl/)  | 11.9K | 112 | 105 | 27 |
| [kata-containers/agent](https://github.com/kata-containers/agent/)  | 11.8K | 833 | 515 | 237 |
| [openebs/website](https://github.com/openebs/website/)  | 11.8K | 1.1K | 342 | 11 |
| [grafana/jslib.k6.io](https://github.com/grafana/jslib.k6.io/)  | 11.5K | 189 | 71 | 32 |
| [open-telemetry/opentelemetry-erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11.5K | 1.1K | 355 | 253 |
| [prometheus/pushgateway](https://github.com/prometheus/pushgateway/)  | 11.5K | 669 | 267 | 2.5K |
| [kubernetes-sigs/blob-csi-driver](https://github.com/kubernetes-sigs/blob-csi-driver/)  | 11.3K | 1.6K | 645 | 97 |
| [openebs/zfs-localpv](https://github.com/openebs/zfs-localpv/)  | 11.3K | 275 | 320 | 261 |
| [docker/getting-started](https://github.com/docker/getting-started/)  | 11.3K | 190 | 212 | 2.5K |
| [ClickHouse/clickhouse-presentations](https://github.com/ClickHouse/clickhouse-presentations/)  | 11.2K | 515 | 39 | 834 |
| [m3db/m3msg](https://github.com/m3db/m3msg/)  | 11.1K | 62 | 55 | 15 |
| [docker/engine-api](https://github.com/docker/engine-api/)  | 11.1K | 9.1K | 327 | 267 |
| [kubernetes-sigs/kubebuilder-declar~](https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/)  | 11.8K | 591 | 262 | 188 |
| [cilium/proxy](https://github.com/cilium/proxy/)  | 10.9K | 567 | 116 | 82 |
| [kubernetes-sigs/controller-tools](https://github.com/kubernetes-sigs/controller-tools/)  | 10.9K | 773 | 455 | 549 |
| [envoyproxy/xds-relay](https://github.com/envoyproxy/xds-relay/)  | 10.8K | 127 | 162 | 128 |
| [openebs/lvm-localpv](https://github.com/openebs/lvm-localpv/)  | 10.7K | 182 | 144 | 141 |
| [open-telemetry/assign-reviewers-ac~](https://github.com/open-telemetry/assign-reviewers-action/)  | 10.7K | 5 | 6 | 6 |
| [kubernetes/code-generator](https://github.com/kubernetes/code-generator/)  | 10.6K | 1.7K | 17 | 1.3K |
| [grafana/jmeter-to-k6](https://github.com/grafana/jmeter-to-k6/)  | 10.6K | 347 | 24 | 57 |
| [kubernetes/utils](https://github.com/kubernetes/utils/)  | 10.6K | 893 | 233 | 256 |
| [kubernetes/node-problem-detector](https://github.com/kubernetes/node-problem-detector/)  | 10.5K | 753 | 434 | 2.2K |
| [grpc/grpc-kotlin](https://github.com/grpc/grpc-kotlin/)  | 10.4K | 230 | 184 | 977 |
| [kubernetes/cli-runtime](https://github.com/kubernetes/cli-runtime/)  | 10.4K | 913 | 5 | 240 |
| [grafana/tanka](https://github.com/grafana/tanka/)  | 10.3K | 449 | 469 | 1.8K |
| [istio/pkg](https://github.com/istio/pkg/)  | 10.1K | 859 | 763 | 47 |
| [kubernetes-sigs/kube-batch](https://github.com/kubernetes-sigs/kube-batch/)  | 10.1K | 1.3K | 691 | 1.3K |
| [istio/istio.io](https://github.com/istio/istio.io/)  | 10.1K | 8.3K | 10.9K | 681 |
| [istio/old_mixerclient_repo](https://github.com/istio/old_mixerclient_repo/)  | 10.4K | 228 | 409 | 15 |
| [kubernetes/kubeadm](https://github.com/kubernetes/kubeadm/)  | 9.9K | 1.2K | 520 | 3.3K |
| [etcd-io/jetcd](https://github.com/etcd-io/jetcd/)  | 9.8K | 1.1K | 742 | 970 |
| [kubernetes-sigs/secrets-store-csi-~](https://github.com/kubernetes-sigs/secrets-store-csi-driver/)  | 9.6K | 1.1K | 747 | 939 |
| [istio/ztunnel](https://github.com/istio/ztunnel/)  | 9.6K | 275 | 301 | 90 |
| [go-faster/jx](https://github.com/go-faster/jx/)  | 9.6K | 1.4K | 65 | 80 |
| [cilium/kube-apate](https://github.com/cilium/kube-apate/)  | 9.5K | 18 | 0 | 4 |
| [grafana/cortex-tools](https://github.com/grafana/cortex-tools/)  | 9.5K | 270 | 179 | 133 |
| [kubernetes-sigs/cloud-provider-hua~](https://github.com/kubernetes-sigs/cloud-provider-huaweicloud/)  | 9.4K | 415 | 171 | 25 |
| [prometheus/statsd_exporter](https://github.com/prometheus/statsd_exporter/)  | 9.4K | 847 | 310 | 823 |
| [kubernetes-sigs/gateway-api](https://github.com/kubernetes-sigs/gateway-api/)  | 9.3K | 2.3K | 1.1K | 871 |
| [grafana/worldmap-panel](https://github.com/grafana/worldmap-panel/)  | 9.3K | 232 | 74 | 294 |
| [openebs/cstor-csi](https://github.com/openebs/cstor-csi/)  | 9.3K | 157 | 172 | 23 |
| [kubernetes-sigs/cri-tools](https://github.com/kubernetes-sigs/cri-tools/)  | 9.2K | 1.5K | 794 | 1.2K |
| [openebs/jiva-operator](https://github.com/openebs/jiva-operator/)  | 9.2K | 148 | 170 | 37 |
| [kubernetes/kompose](https://github.com/kubernetes/kompose/)  | 9.2K | 1.4K | 895 | 8.3K |
| [ClickHouse/clickhouse-jdbc-bridge](https://github.com/ClickHouse/clickhouse-jdbc-bridge/)  | 9.1K | 127 | 85 | 137 |
| [envoyproxy/go-control-plane](https://github.com/envoyproxy/go-control-plane/)  | 9.1K | 1.3K | 398 | 1.2K |
| [grpc/test-infra](https://github.com/grpc/test-infra/)  | 9.1K | 484 | 346 | 64 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-gcp/)  | 8.9K | 1.1K | 673 | 131 |
| [ClickHouse/libcxxabi](https://github.com/ClickHouse/libcxxabi/)  | 8.8K | 14 | 4 | 0 |
| [kubernetes-sigs/apiserver-network-~](https://github.com/kubernetes-sigs/apiserver-network-proxy/)  | 8.7K | 640 | 316 | 262 |
| [kubernetes/pod-security-admission](https://github.com/kubernetes/pod-security-admission/)  | 8.7K | 465 | 0 | 77 |
| [ClickHouse/graphouse](https://github.com/ClickHouse/graphouse/)  | 8.7K | 542 | 160 | 251 |
| [ClickHouse/clickhouse-connect](https://github.com/ClickHouse/clickhouse-connect/)  | 8.7K | 229 | 84 | 86 |
| [kubernetes-sigs/cli-experimental](https://github.com/kubernetes-sigs/cli-experimental/)  | 8.6K | 278 | 200 | 65 |
| [kubernetes/gengo](https://github.com/kubernetes/gengo/)  | 8.4K | 455 | 182 | 485 |
| [kubernetes-sigs/cloud-provider-bai~](https://github.com/kubernetes-sigs/cloud-provider-baiducloud/)  | 8.3K | 213 | 71 | 38 |
| [m3db/m3em](https://github.com/m3db/m3em/)  | 8.1K | 25 | 19 | 1 |
| [openebs/mayastor-extensions](https://github.com/openebs/mayastor-extensions/)  | 7.8K | 220 | 125 | 10 |
| [envoyproxy/ratelimit](https://github.com/envoyproxy/ratelimit/)  | 7.8K | 150 | 235 | 1.7K |
| [open-telemetry/opentelemetry-go-bu~](https://github.com/open-telemetry/opentelemetry-go-build-tools/)  | 7.7K | 309 | 242 | 21 |
| [openebs/device-localpv](https://github.com/openebs/device-localpv/)  | 7.7K | 55 | 49 | 15 |
| [kubernetes-sigs/krew](https://github.com/kubernetes-sigs/krew/)  | 7.6K | 478 | 479 | 5.3K |
| [prometheus/client_python](https://github.com/prometheus/client_python/)  | 7.6K | 512 | 389 | 3.2K |
| [grafana/grafana-polystat-panel](https://github.com/grafana/grafana-polystat-panel/)  | 7.5K | 199 | 124 | 74 |
| [kubernetes/kube-aggregator](https://github.com/kubernetes/kube-aggregator/)  | 7.4K | 2.8K | 11 | 223 |
| [grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client/)  | 7.4K | 402 | 135 | 71 |
| [kubernetes-sigs/aws-iam-authentica~](https://github.com/kubernetes-sigs/aws-iam-authenticator/)  | 7.3K | 515 | 313 | 1.9K |
| [kubernetes-sigs/aws-efs-csi-driver](https://github.com/kubernetes-sigs/aws-efs-csi-driver/)  | 7.3K | 689 | 585 | 568 |
| [openebs/upgrade](https://github.com/openebs/upgrade/)  | 7.2K | 122 | 135 | 10 |
| [grafana/postman-to-k6](https://github.com/grafana/postman-to-k6/)  | 7.2K | 607 | 58 | 268 |
| [kubernetes-sigs/instrumentation-to~](https://github.com/kubernetes-sigs/instrumentation-tools/)  | 7.1K | 92 | 6 | 25 |
| [docker/kitematic](https://github.com/docker/kitematic/)  | 7.1K | 2.3K | 534 | 12.2K |
| [grpc/grpc-web](https://github.com/grpc/grpc-web/)  | 7.7K | 888 | 622 | 7.3K |
| [kubernetes/dns](https://github.com/kubernetes/dns/)  | 7.3K | 710 | 318 | 802 |
| [grafana/opcua-datasource](https://github.com/grafana/opcua-datasource/)  | 7.3K | 370 | 44 | 45 |
| [prometheus/mysqld_exporter](https://github.com/prometheus/mysqld_exporter/)  | 6.9K | 605 | 412 | 1.6K |
| [openebs/libcstor](https://github.com/openebs/libcstor/)  | 6.9K | 77 | 89 | 14 |
| [open-telemetry/opentelemetry-demo](https://github.com/open-telemetry/opentelemetry-demo/)  | 6.8K | 420 | 549 | 690 |
| [cilium/hubble](https://github.com/cilium/hubble/)  | 6.8K | 796 | 757 | 2.5K |
| [grafana/attic](https://github.com/grafana/attic/)  | 6.8K | 426 | 0 | 1 |
| [envoyproxy/envoy-build-tools](https://github.com/envoyproxy/envoy-build-tools/)  | 6.8K | 320 | 182 | 36 |
| [kubernetes-sigs/reference-docs](https://github.com/kubernetes-sigs/reference-docs/)  | 6.8K | 492 | 233 | 65 |
| [grafana/kairosdb-datasource](https://github.com/grafana/kairosdb-datasource/)  | 6.7K | 110 | 43 | 31 |
| [cilium/image-tools](https://github.com/cilium/image-tools/)  | 6.7K | 239 | 192 | 11 |
| [helm/chartmuseum](https://github.com/helm/chartmuseum/)  | 6.6K | 613 | 333 | 3.1K |
| [etcd-io/dbtester](https://github.com/etcd-io/dbtester/)  | 6.5K | 1.2K | 288 | 250 |
| [kubernetes-sigs/apiserver-builder-~](https://github.com/kubernetes-sigs/apiserver-builder-alpha/)  | 6.5K | 1.7K | 434 | 717 |
| [envoyproxy/envoy-perf](https://github.com/envoyproxy/envoy-perf/)  | 6.5K | 142 | 160 | 107 |
| [gotd/botapi](https://github.com/gotd/botapi/)  | 6.4K | 653 | 313 | 18 |
| [kubernetes/csi-translation-lib](https://github.com/kubernetes/csi-translation-lib/)  | 6.3K | 658 | 0 | 10 |
| [open-telemetry/opamp-go](https://github.com/open-telemetry/opamp-go/)  | 6.3K | 92 | 110 | 60 |
| [kubernetes-sigs/kubetest2](https://github.com/kubernetes-sigs/kubetest2/)  | 6.3K | 535 | 177 | 259 |
| [kubernetes-sigs/sig-storage-local-~](https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner/)  | 6.1K | 623 | 213 | 866 |
| [kubernetes-sigs/slack-infra](https://github.com/kubernetes-sigs/slack-infra/)  | 6.1K | 108 | 44 | 85 |
| [cilium/docsearch-scraper-webhook](https://github.com/cilium/docsearch-scraper-webhook/)  | 5.9K | 38 | 32 | 3 |
| [siderolabs/kres](https://github.com/siderolabs/kres/)  | 5.9K | 152 | 199 | 22 |
| [ClickHouse/bzip2](https://github.com/ClickHouse/bzip2/)  | 5.9K | 141 | 0 | 1 |
| [prometheus/blackbox_exporter](https://github.com/prometheus/blackbox_exporter/)  | 5.9K | 455 | 436 | 3.4K |
| [grafana/devtools](https://github.com/grafana/devtools/)  | 5.9K | 139 | 22 | 9 |
| [docker/go](https://github.com/docker/go/)  | 5.8K | 24 | 8 | 17 |
| [kata-containers/govmm](https://github.com/kata-containers/govmm/)  | 5.8K | 402 | 151 | 304 |
| [docker/volumes-backup-extension](https://github.com/docker/volumes-backup-extension/)  | 5.8K | 228 | 76 | 47 |
| [istio/api](https://github.com/istio/api/)  | 5.8K | 1.6K | 2.4K | 404 |
| [kubernetes/website](https://github.com/kubernetes/website/)  | 5.7K | 39.7K | 29.1K | 3.6K |
| [docker/metadata-action](https://github.com/docker/metadata-action/)  | 5.7K | 318 | 157 | 563 |
| [grafana/azure-monitor-datasource](https://github.com/grafana/azure-monitor-datasource/)  | 5.6K | 170 | 18 | 91 |
| [prometheus/compliance](https://github.com/prometheus/compliance/)  | 5.6K | 147 | 79 | 105 |
| [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/)  | 5.5K | 989 | 602 | 4.5K |
| [kubernetes/component-helpers](https://github.com/kubernetes/component-helpers/)  | 5.5K | 329 | 0 | 11 |
| [m3db/m3storage](https://github.com/m3db/m3storage/)  | 5.5K | 38 | 15 | 3 |
| [openebs/api](https://github.com/openebs/api/)  | 5.3K | 88 | 99 | 7 |
| [kubernetes-sigs/cluster-addons](https://github.com/kubernetes-sigs/cluster-addons/)  | 5.3K | 261 | 89 | 147 |
| [prometheus/codemirror-promql](https://github.com/prometheus/codemirror-promql/)  | 5.3K | 516 | 151 | 35 |
| [VictoriaMetrics/metricsql](https://github.com/VictoriaMetrics/metricsql/)  | 5.2K | 102 | 6 | 136 |
| [kubernetes-sigs/aws-fsx-csi-driver](https://github.com/kubernetes-sigs/aws-fsx-csi-driver/)  | 5.2K | 409 | 204 | 103 |
| [grafana/kubernetes-app](https://github.com/grafana/kubernetes-app/)  | 5.2K | 199 | 23 | 397 |
| [helm/helm-classic](https://github.com/helm/helm-classic/)  | 5.2K | 574 | 274 | 578 |
| [etcd-io/zetcd](https://github.com/etcd-io/zetcd/)  | 5.1K | 135 | 58 | 1.6K |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-digitalocean/)  | 5.8K | 399 | 383 | 92 |
| [m3db/m3ctl](https://github.com/m3db/m3ctl/)  | 4.9K | 82 | 64 | 14 |
| [grafana/gel-app](https://github.com/grafana/gel-app/)  | 4.9K | 119 | 47 | 2 |
| [prometheus/snmp_exporter](https://github.com/prometheus/snmp_exporter/)  | 4.7K | 553 | 416 | 1.1K |
| [helm/monocular](https://github.com/helm/monocular/)  | 4.7K | 399 | 362 | 1.4K |
| [kata-containers/tests](https://github.com/kata-containers/tests/)  | 4.7K | 4.3K | 2.7K | 135 |
| [gotd/contrib](https://github.com/gotd/contrib/)  | 4.7K | 568 | 317 | 10 |
| [kubernetes-sigs/image-builder](https://github.com/kubernetes-sigs/image-builder/)  | 4.6K | 2.5K | 815 | 248 |
| [grafana/grafana-sensu-app](https://github.com/grafana/grafana-sensu-app/)  | 4.6K | 16 | 45 | 8 |
| [istio/old_auth_repo](https://github.com/istio/old_auth_repo/)  | 4.5K | 166 | 230 | 73 |
| [openebs/velero-plugin](https://github.com/openebs/velero-plugin/)  | 4.5K | 119 | 137 | 43 |


