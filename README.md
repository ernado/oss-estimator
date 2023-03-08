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
| [kubernetes](https://github.com/kubernetes)  | 3434971 | 381074 | 190484 | 243089 |
| [tensorflow](https://github.com/tensorflow)  | 3351209 | 144463 | 22710 | 171882 |
| [ClickHouse](https://github.com/ClickHouse)  | 2341463 | 166223 | 34848 | 33869 |
| [kubernetes-sigs](https://github.com/kubernetes-sigs)  | 1997680 | 117395 | 71019 | 98275 |
| [envoyproxy](https://github.com/envoyproxy)  | 1871907 | 46722 | 24304 | 28102 |
| [open-telemetry](https://github.com/open-telemetry)  | 1780557 | 61864 | 61019 | 28606 |
| [grpc](https://github.com/grpc)  | 1683336 | 74131 | 38788 | 86819 |
| [grafana](https://github.com/grafana)  | 1679277 | 76610 | 50201 | 103185 |
| [docker](https://github.com/docker)  | 1266401 | 173748 | 30587 | 146425 |
| [openebs](https://github.com/openebs)  | 1165411 | 17867 | 10441 | 10172 |
| [golang](https://github.com/golang)  | 1061006 | 55722 | 2865 | 109154 |
| [cilium](https://github.com/cilium)  | 867372 | 47596 | 21845 | 23625 |
| [m3db](https://github.com/m3db)  | 736760 | 10433 | 5079 | 4646 |
| [istio](https://github.com/istio)  | 583580 | 47045 | 62373 | 38489 |
| [apache](https://github.com/apache)  | 420406 | 22275 | 521 | 5657 |
| [prometheus](https://github.com/prometheus)  | 377349 | 30866 | 19120 | 89530 |
| [facebook](https://github.com/facebook)  | 357548 | 15583 | 13253 | 203498 |
| [pixie-io](https://github.com/pixie-io)  | 352478 | 11768 | 549 | 4380 |
| [LINBIT](https://github.com/LINBIT)  | 310875 | 4396 | 11 | 649 |
| [vectordotdev](https://github.com/vectordotdev)  | 308968 | 14392 | 10198 | 14352 |
| [uber](https://github.com/uber)  | 253779 | 1572 | 248 | 5943 |
| [siderolabs](https://github.com/siderolabs)  | 232773 | 7105 | 9062 | 4779 |
| [kata-containers](https://github.com/kata-containers)  | 223850 | 22768 | 10297 | 7000 |
| [Netflix](https://github.com/Netflix)  | 208149 | 4152 | 2264 | 551 |
| [VictoriaMetrics](https://github.com/VictoriaMetrics)  | 205963 | 8126 | 2587 | 10986 |
| [etcd-io](https://github.com/etcd-io)  | 185291 | 26381 | 11277 | 52282 |
| [containers](https://github.com/containers)  | 174455 | 18210 | 10101 | 17085 |
| [VKCOM](https://github.com/VKCOM)  | 101459 | 9819 | 3089 | 877 |
| [gotd](https://github.com/gotd)  | 74334 | 6222 | 2440 | 1162 |
| [helm](https://github.com/helm)  | 71506 | 28920 | 26875 | 48778 |
| [vuejs](https://github.com/vuejs)  | 60290 | 3546 | 2389 | 202656 |
| [go-faster](https://github.com/go-faster)  | 50751 | 4707 | 670 | 151 |
| [vitalif](https://github.com/vitalif)  | 40876 | 1224 | 12 | 69 |
| [ogen-go](https://github.com/ogen-go)  | 38140 | 3872 | 1050 | 472 |
| [falcosecurity](https://github.com/falcosecurity)  | 13984 | 3557 | 1446 | 5681 |


| Project                                        | SLOC | Commits | PRs | Stars |
|------------------------------------------------|------|---------|-----|-------|
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow/)  | 3351209 | 144463 | 22710 | 171882 |
| [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/)  | 1488546 | 114442 | 73545 | 96462 |
| [ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse/)  | 1112577 | 109729 | 32220 | 27404 |
| [golang/go](https://github.com/golang/go/)  | 1061006 | 55722 | 2865 | 109154 |
| [grafana/grafana](https://github.com/grafana/grafana/)  | 889820 | 41339 | 33532 | 54290 |
| [envoyproxy/envoy](https://github.com/envoyproxy/envoy/)  | 807577 | 17306 | 17107 | 21571 |
| [openebs/openebs-docs](https://github.com/openebs/openebs-docs/)  | 640750 | 2366 | 832 | 37 |
| [grpc/grpc](https://github.com/grpc/grpc/)  | 514803 | 52698 | 21634 | 37233 |
| [grpc/grpc-ios](https://github.com/grpc/grpc-ios/)  | 512418 | 96 | 103 | 22 |
| [envoyproxy/envoy-wasm](https://github.com/envoyproxy/envoy-wasm/)  | 504503 | 8484 | 454 | 205 |
| [m3db/m3](https://github.com/m3db/m3/)  | 477343 | 4211 | 3564 | 4366 |
| [ClickHouse/grpc](https://github.com/ClickHouse/grpc/)  | 397291 | 47064 | 13 | 1 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector-contrib/)  | 389240 | 9606 | 16162 | 1524 |
| [facebook/react](https://github.com/facebook/react/)  | 357548 | 15583 | 13253 | 203498 |
| [pixie-io/pixie](https://github.com/pixie-io/pixie/)  | 352478 | 11768 | 549 | 4380 |
| [istio/istio](https://github.com/istio/istio/)  | 339081 | 19854 | 26467 | 32549 |
| [kubernetes/autoscaler](https://github.com/kubernetes/autoscaler/)  | 332826 | 6459 | 3739 | 6592 |
| [cilium/cilium](https://github.com/cilium/cilium/)  | 324822 | 22284 | 16999 | 14569 |
| [docker/docker-ce](https://github.com/docker/docker-ce/)  | 324419 | 54308 | 662 | 5586 |
| [LINBIT/linstor-server](https://github.com/LINBIT/linstor-server/)  | 310875 | 4396 | 11 | 649 |
| [apache/mesos](https://github.com/apache/mesos/)  | 305614 | 18184 | 450 | 5029 |
| [docker/labs](https://github.com/docker/labs/)  | 304415 | 718 | 398 | 11147 |
| [envoyproxy/envoy-website](https://github.com/envoyproxy/envoy-website/)  | 299034 | 428 | 252 | 33 |
| [kubernetes-sigs/security-profiles-~](https://github.com/kubernetes-sigs/security-profiles-operator/)  | 284401 | 1597 | 1334 | 465 |
| [vectordotdev/vector](https://github.com/vectordotdev/vector/)  | 274642 | 9352 | 9365 | 12907 |
| [kubernetes/test-infra](https://github.com/kubernetes/test-infra/)  | 270560 | 51941 | 24615 | 3571 |
| [docker/get-involved](https://github.com/docker/get-involved/)  | 264430 | 1635 | 36 | 24 |
| [grafana/loki](https://github.com/grafana/loki/)  | 241078 | 4717 | 5180 | 18489 |
| [grpc/grpc-java](https://github.com/grpc/grpc-java/)  | 235165 | 5735 | 6854 | 10402 |
| [uber/peloton](https://github.com/uber/peloton/)  | 216371 | 705 | 10 | 582 |
| [cilium/pwru](https://github.com/cilium/pwru/)  | 194155 | 181 | 125 | 1156 |
| [ClickHouse/clickhouse-website-cont~](https://github.com/ClickHouse/clickhouse-website-content/)  | 186631 | 1 | 2 | 2 |
| [cilium/busybox](https://github.com/cilium/busybox/)  | 185793 | 16646 | 0 | 0 |
| [containers/podman](https://github.com/containers/podman/)  | 174455 | 18210 | 10101 | 17085 |
| [kubernetes-sigs/cluster-api](https://github.com/kubernetes-sigs/cluster-api/)  | 173316 | 8851 | 5326 | 2775 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-instrumentation/)  | 167856 | 8835 | 5066 | 1214 |
| [kubernetes/kops](https://github.com/kubernetes/kops/)  | 167601 | 19350 | 10375 | 14758 |
| [Netflix/titus-control-plane](https://github.com/Netflix/titus-control-plane/)  | 157302 | 1663 | 1267 | 319 |
| [ClickHouse/ssl](https://github.com/ClickHouse/ssl/)  | 156067 | 16 | 2 | 2 |
| [prometheus/prometheus](https://github.com/prometheus/prometheus/)  | 155113 | 10934 | 6786 | 47081 |
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/)  | 144520 | 5903 | 1822 | 8122 |
| [grpc/grpc-go](https://github.com/grpc/grpc-go/)  | 143240 | 4462 | 3913 | 17694 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/)  | 141088 | 927 | 1797 | 224 |
| [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize/)  | 139479 | 6327 | 3002 | 9410 |
| [kubernetes/apiserver](https://github.com/kubernetes/apiserver/)  | 131328 | 6180 | 25 | 494 |
| [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers/)  | 130813 | 10136 | 3007 | 3088 |
| [siderolabs/talos](https://github.com/siderolabs/talos/)  | 130197 | 3899 | 5193 | 3677 |
| [etcd-io/etcd](https://github.com/etcd-io/etcd/)  | 125664 | 19331 | 9188 | 42804 |
| [apache/aurora](https://github.com/apache/aurora/)  | 114792 | 4091 | 71 | 628 |
| [open-telemetry/opentelemetry-sandb~](https://github.com/open-telemetry/opentelemetry-sandbox-web-js/)  | 113692 | 2822 | 60 | 10 |
| [kubernetes-sigs/vsphere-csi-driver](https://github.com/kubernetes-sigs/vsphere-csi-driver/)  | 106976 | 2269 | 1841 | 218 |
| [grafana/agent](https://github.com/grafana/agent/)  | 105905 | 1465 | 1770 | 980 |
| [openebs/maya](https://github.com/openebs/maya/)  | 100071 | 1799 | 1669 | 180 |
| [open-telemetry/opentelemetry-java](https://github.com/open-telemetry/opentelemetry-java/)  | 96376 | 3268 | 3620 | 1452 |
| [kubernetes/kubectl](https://github.com/kubernetes/kubectl/)  | 95965 | 2962 | 313 | 2215 |
| [kubernetes/minikube](https://github.com/kubernetes/minikube/)  | 90069 | 20969 | 7603 | 26015 |
| [ClickHouse/UnixODBC](https://github.com/ClickHouse/UnixODBC/)  | 88921 | 5 | 2 | 0 |
| [kubernetes-sigs/cloud-provider-azu~](https://github.com/kubernetes-sigs/cloud-provider-azure/)  | 84466 | 3425 | 3165 | 210 |
| [docker/cli](https://github.com/docker/cli/)  | 78919 | 8462 | 2687 | 3901 |
| [kubernetes/ingress-gce](https://github.com/kubernetes/ingress-gce/)  | 78027 | 4552 | 1564 | 1192 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-azure/)  | 77206 | 3614 | 2182 | 244 |
| [kubernetes-sigs/aws-load-balancer-~](https://github.com/kubernetes-sigs/aws-load-balancer-controller/)  | 76499 | 630 | 1197 | 3264 |
| [grafana/metrictank](https://github.com/grafana/metrictank/)  | 73770 | 6529 | 1242 | 628 |
| [ClickHouse/clickhouse-java](https://github.com/ClickHouse/clickhouse-java/)  | 73753 | 1723 | 634 | 1171 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)  | 73650 | 3831 | 2676 | 543 |
| [grafana/k6](https://github.com/grafana/k6/)  | 73040 | 5057 | 1451 | 19694 |
| [kubernetes-sigs/kui](https://github.com/kubernetes-sigs/kui/)  | 72373 | 4834 | 5272 | 2435 |
| [grpc/grpc-experiments](https://github.com/grpc/grpc-experiments/)  | 71879 | 633 | 232 | 1069 |
| [kubernetes/client-go](https://github.com/kubernetes/client-go/)  | 65629 | 3800 | 204 | 7399 |
| [kata-containers/runtime](https://github.com/kata-containers/runtime/)  | 63906 | 2751 | 1453 | 2114 |
| [grpc/grpc-dotnet](https://github.com/grpc/grpc-dotnet/)  | 63899 | 859 | 941 | 3545 |
| [VKCOM/statshouse](https://github.com/VKCOM/statshouse/)  | 61053 | 236 | 214 | 120 |
| [open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js/)  | 61017 | 1709 | 2005 | 1798 |
| [openebs/mayastor-control-plane](https://github.com/openebs/mayastor-control-plane/)  | 60917 | 1231 | 447 | 27 |
| [kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns/)  | 60511 | 3011 | 1858 | 6119 |
| [vuejs/vue](https://github.com/vuejs/vue/)  | 60290 | 3546 | 2389 | 202656 |
| [kubernetes/legacy-cloud-providers](https://github.com/kubernetes/legacy-cloud-providers/)  | 60213 | 1807 | 0 | 47 |
| [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/)  | 60137 | 2517 | 30 | 661 |
| [envoyproxy/envoy-mobile](https://github.com/envoyproxy/envoy-mobile/)  | 57633 | 1794 | 2208 | 550 |
| [open-telemetry/opentelemetry-dotnet](https://github.com/open-telemetry/opentelemetry-dotnet/)  | 56924 | 2277 | 2792 | 2276 |
| [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/)  | 56865 | 6900 | 4907 | 14521 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/)  | 56099 | 1812 | 1064 | 286 |
| [openebs/mayastor](https://github.com/openebs/mayastor/)  | 54169 | 1877 | 1148 | 457 |
| [open-telemetry/experimental-arrow-~](https://github.com/open-telemetry/experimental-arrow-collector/)  | 53440 | 4416 | 37 | 11 |
| [gotd/td](https://github.com/gotd/td/)  | 52603 | 3084 | 788 | 1056 |
| [open-telemetry/opentelemetry-js-co~](https://github.com/open-telemetry/opentelemetry-js-contrib/)  | 51352 | 1585 | 990 | 420 |
| [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/)  | 51245 | 3162 | 8 | 206 |
| [open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/)  | 51094 | 1943 | 2389 | 3670 |
| [Netflix/titus-executor](https://github.com/Netflix/titus-executor/)  | 50847 | 2489 | 997 | 232 |
| [m3db/m3metrics](https://github.com/m3db/m3metrics/)  | 50672 | 233 | 194 | 9 |
| [open-telemetry/opentelemetry-colle~](https://github.com/open-telemetry/opentelemetry-collector/)  | 50320 | 4416 | 5092 | 2759 |
| [kubernetes-sigs/multi-tenancy](https://github.com/kubernetes-sigs/multi-tenancy/)  | 48929 | 2300 | 1064 | 930 |
| [envoyproxy/pytooling](https://github.com/envoyproxy/pytooling/)  | 48908 | 615 | 619 | 6 |
| [grpc/grpc-swift](https://github.com/grpc/grpc-swift/)  | 48316 | 1564 | 981 | 1718 |
| [m3db/pilosa](https://github.com/m3db/pilosa/)  | 46556 | 4474 | 5 | 1 |
| [openebs/istgt](https://github.com/openebs/istgt/)  | 46447 | 241 | 358 | 20 |
| [grafana/azure-data-explorer-dataso~](https://github.com/grafana/azure-data-explorer-datasource/)  | 45793 | 790 | 309 | 38 |
| [open-telemetry/opentelemetry-pytho~](https://github.com/open-telemetry/opentelemetry-python-contrib/)  | 45783 | 1754 | 1021 | 413 |
| [kubernetes-sigs/cli-utils](https://github.com/kubernetes-sigs/cli-utils/)  | 45255 | 1086 | 543 | 113 |
| [kubernetes/dashboard](https://github.com/kubernetes/dashboard/)  | 44851 | 4440 | 4943 | 12308 |
| [istio/old_mixer_repo](https://github.com/istio/old_mixer_repo/)  | 42543 | 741 | 1091 | 68 |
| [grpc/grpc-node](https://github.com/grpc/grpc-node/)  | 42337 | 4216 | 1367 | 3853 |
| [open-telemetry/opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python/)  | 42221 | 1404 | 1728 | 1207 |
| [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)  | 42148 | 2221 | 1300 | 1877 |
| [open-telemetry/opentelemetry-cpp-c~](https://github.com/open-telemetry/opentelemetry-cpp-contrib/)  | 42051 | 157 | 180 | 78 |
| [helm/helm](https://github.com/helm/helm/)  | 41739 | 6809 | 4836 | 23880 |
| [open-telemetry/opentelemetry-ebpf](https://github.com/open-telemetry/opentelemetry-ebpf/)  | 41435 | 267 | 110 | 85 |
| [vitalif/vitastor](https://github.com/vitalif/vitastor/)  | 40876 | 1224 | 12 | 69 |
| [kubernetes-sigs/sig-windows-samples](https://github.com/kubernetes-sigs/sig-windows-samples/)  | 40825 | 52 | 3 | 5 |
| [VKCOM/VKUI](https://github.com/VKCOM/VKUI/)  | 40406 | 9583 | 2875 | 757 |
| [m3db/m3aggregator](https://github.com/m3db/m3aggregator/)  | 39089 | 177 | 142 | 13 |
| [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/)  | 38012 | 2978 | 1847 | 6199 |
| [uber/kraken](https://github.com/uber/kraken/)  | 37408 | 867 | 238 | 5361 |
| [kubernetes/cloud-provider-alibaba-~](https://github.com/kubernetes/cloud-provider-alibaba-cloud/)  | 37304 | 799 | 283 | 321 |
| [ogen-go/ogen](https://github.com/ogen-go/ogen/)  | 36658 | 3267 | 700 | 454 |
| [ClickHouse/libpq](https://github.com/ClickHouse/libpq/)  | 36352 | 35 | 7 | 1 |
| [openebs/cstor-operators](https://github.com/openebs/cstor-operators/)  | 36302 | 298 | 358 | 83 |
| [VictoriaMetrics/operator](https://github.com/VictoriaMetrics/operator/)  | 35457 | 662 | 313 | 278 |
| [ClickHouse/clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  | 34226 | 1262 | 459 | 2302 |
| [istio/get-istioctl](https://github.com/istio/get-istioctl/)  | 34103 | 13 | 0 | 6 |
| [kubernetes/cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp/)  | 33691 | 1459 | 420 | 78 |
| [open-telemetry/opentelemetry-cpp](https://github.com/open-telemetry/opentelemetry-cpp/)  | 33554 | 1005 | 1187 | 431 |
| [kubernetes/perf-tests](https://github.com/kubernetes/perf-tests/)  | 33389 | 3123 | 1888 | 764 |
| [open-telemetry/opentelemetry-dotne~](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/)  | 33382 | 732 | 850 | 230 |
| [prometheus/alertmanager](https://github.com/prometheus/alertmanager/)  | 31884 | 2812 | 1677 | 5544 |
| [open-telemetry/opentelemetry-php](https://github.com/open-telemetry/opentelemetry-php/)  | 30979 | 508 | 591 | 519 |
| [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/)  | 30917 | 1258 | 308 | 232 |
| [ClickHouse/libc-headers](https://github.com/ClickHouse/libc-headers/)  | 30907 | 18 | 4 | 0 |
| [open-telemetry/opentelemetry-rust](https://github.com/open-telemetry/opentelemetry-rust/)  | 30530 | 574 | 597 | 997 |
| [docker/machine](https://github.com/docker/machine/)  | 29940 | 3463 | 1957 | 6562 |
| [cilium/ebpf](https://github.com/cilium/ebpf/)  | 29933 | 1290 | 686 | 4160 |
| [m3db/m3cluster](https://github.com/m3db/m3cluster/)  | 29627 | 238 | 227 | 21 |
| [open-telemetry/opentelemetry-log-c~](https://github.com/open-telemetry/opentelemetry-log-collection/)  | 28053 | 267 | 373 | 92 |
| [envoyproxy/nighthawk](https://github.com/envoyproxy/nighthawk/)  | 27852 | 578 | 655 | 272 |
| [kubernetes-sigs/alibaba-cloud-csi-~](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/)  | 27638 | 1341 | 590 | 444 |
| [kubernetes/release](https://github.com/kubernetes/release/)  | 27346 | 5187 | 2392 | 436 |
| [docker/docker-py](https://github.com/docker/docker-py/)  | 27332 | 3299 | 1538 | 6131 |
| [open-telemetry/opentelemetry-go-co~](https://github.com/open-telemetry/opentelemetry-go-contrib/)  | 27256 | 1046 | 3210 | 684 |
| [cilium/hubble-ui](https://github.com/cilium/hubble-ui/)  | 26768 | 385 | 426 | 237 |
| [ClickHouse/libhdfs3](https://github.com/ClickHouse/libhdfs3/)  | 26723 | 64 | 34 | 23 |
| [ClickHouse/clickhouse-cpp](https://github.com/ClickHouse/clickhouse-cpp/)  | 26671 | 673 | 172 | 215 |
| [open-telemetry/opentelemetry-java-~](https://github.com/open-telemetry/opentelemetry-java-contrib/)  | 26343 | 562 | 644 | 80 |
| [open-telemetry/opentelemetry-swift](https://github.com/open-telemetry/opentelemetry-swift/)  | 25907 | 618 | 252 | 126 |
| [envoyproxy/java-control-plane](https://github.com/envoyproxy/java-control-plane/)  | 25478 | 235 | 206 | 246 |
| [grafana/grafana-plugin-sdk-go](https://github.com/grafana/grafana-plugin-sdk-go/)  | 25368 | 415 | 531 | 157 |
| [kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)  | 25340 | 642 | 349 | 733 |
| [kubernetes-sigs/kubefed](https://github.com/kubernetes-sigs/kubefed/)  | 25048 | 2745 | 978 | 2450 |
| [docker/compose-cli](https://github.com/docker/compose-cli/)  | 24959 | 3024 | 1214 | 910 |
| [grafana/terraform-provider-grafana](https://github.com/grafana/terraform-provider-grafana/)  | 24623 | 697 | 522 | 312 |
| [siderolabs/theila](https://github.com/siderolabs/theila/)  | 24315 | 100 | 114 | 45 |
| [kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/)  | 23693 | 2570 | 1234 | 4370 |
| [open-telemetry/opentelemetry-opera~](https://github.com/open-telemetry/opentelemetry-operator/)  | 23641 | 781 | 1019 | 675 |
| [openebs/elves](https://github.com/openebs/elves/)  | 23246 | 234 | 47 | 17 |
| [openebs/node-disk-manager](https://github.com/openebs/node-disk-manager/)  | 23219 | 451 | 533 | 170 |
| [istio/old_pilot_repo](https://github.com/istio/old_pilot_repo/)  | 22997 | 688 | 1241 | 138 |
| [istio/proxy](https://github.com/istio/proxy/)  | 22423 | 2482 | 4232 | 700 |
| [kubernetes-sigs/azuredisk-csi-driv~](https://github.com/kubernetes-sigs/azuredisk-csi-driver/)  | 21740 | 2724 | 1337 | 113 |
| [envoyproxy/gateway](https://github.com/envoyproxy/gateway/)  | 21624 | 539 | 632 | 886 |
| [istio/operator](https://github.com/istio/operator/)  | 21286 | 508 | 780 | 174 |
| [kubernetes-sigs/structured-merge-d~](https://github.com/kubernetes-sigs/structured-merge-diff/)  | 20928 | 495 | 208 | 74 |
| [prometheus/procfs](https://github.com/prometheus/procfs/)  | 20804 | 572 | 372 | 618 |
| [prometheus/client_golang](https://github.com/prometheus/client_golang/)  | 20747 | 1447 | 741 | 4436 |
| [docker/buildx](https://github.com/docker/buildx/)  | 20627 | 1420 | 722 | 2473 |
| [kubernetes/cloud-provider-vsphere](https://github.com/kubernetes/cloud-provider-vsphere/)  | 20499 | 1087 | 472 | 189 |
| [kubernetes/cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack/)  | 20280 | 2670 | 1368 | 507 |
| [open-telemetry/opentelemetry-ruby-~](https://github.com/open-telemetry/opentelemetry-ruby-contrib/)  | 19976 | 895 | 241 | 29 |
| [prometheus/node_exporter](https://github.com/prometheus/node_exporter/)  | 19276 | 1966 | 1410 | 8559 |
| [envoyproxy/data-plane-api](https://github.com/envoyproxy/data-plane-api/)  | 19231 | 2388 | 540 | 510 |
| [kubernetes-sigs/descheduler](https://github.com/kubernetes-sigs/descheduler/)  | 19077 | 1334 | 667 | 3212 |
| [cilium/cilium-olm](https://github.com/cilium/cilium-olm/)  | 19028 | 341 | 55 | 11 |
| [ClickHouse/libgsasl](https://github.com/ClickHouse/libgsasl/)  | 18930 | 21 | 8 | 0 |
| [grpc/grpc-dart](https://github.com/grpc/grpc-dart/)  | 18720 | 228 | 297 | 759 |
| [cilium/cilium-cli](https://github.com/cilium/cilium-cli/)  | 18381 | 1387 | 1135 | 205 |
| [kubernetes-sigs/gcp-filestore-csi-~](https://github.com/kubernetes-sigs/gcp-filestore-csi-driver/)  | 17882 | 593 | 403 | 68 |
| [istio/tools](https://github.com/istio/tools/)  | 17492 | 1793 | 2388 | 302 |
| [openebs/dynamic-nfs-provisioner](https://github.com/openebs/dynamic-nfs-provisioner/)  | 17333 | 112 | 109 | 103 |
| [docker/compose](https://github.com/docker/compose/)  | 17332 | 4010 | 3515 | 28796 |
| [docker/compose-on-kubernetes](https://github.com/docker/compose-on-kubernetes/)  | 17216 | 236 | 114 | 1425 |
| [kubernetes-sigs/etcdadm](https://github.com/kubernetes-sigs/etcdadm/)  | 16898 | 1219 | 248 | 652 |
| [kubernetes/api](https://github.com/kubernetes/api/)  | 16874 | 6711 | 10 | 514 |
| [ClickHouse/libcxx](https://github.com/ClickHouse/libcxx/)  | 16510 | 38 | 14 | 0 |
| [open-telemetry/opentelemetry-ruby](https://github.com/open-telemetry/opentelemetry-ruby/)  | 16445 | 763 | 919 | 375 |
| [ClickHouse/clickhouse-odbc](https://github.com/ClickHouse/clickhouse-odbc/)  | 16287 | 1110 | 172 | 222 |
| [kubernetes-sigs/aws-ebs-csi-driver](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/)  | 16204 | 1654 | 882 | 720 |
| [ClickHouse/boost](https://github.com/ClickHouse/boost/)  | 16195 | 83 | 28 | 1 |
| [ClickHouse/antlr4-runtime](https://github.com/ClickHouse/antlr4-runtime/)  | 16181 | 306 | 0 | 1 |
| [kubernetes-sigs/boskos](https://github.com/kubernetes-sigs/boskos/)  | 15965 | 923 | 129 | 93 |
| [openebs/dynamic-localpv-provisioner](https://github.com/openebs/dynamic-localpv-provisioner/)  | 15882 | 154 | 123 | 88 |
| [kubernetes-sigs/gcp-compute-persis~](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver/)  | 15729 | 1373 | 824 | 133 |
| [prometheus/client_java](https://github.com/prometheus/client_java/)  | 15646 | 578 | 490 | 1929 |
| [grafana/carbon-relay-ng](https://github.com/grafana/carbon-relay-ng/)  | 15553 | 1118 | 258 | 454 |
| [go-faster/yamlx](https://github.com/go-faster/yamlx/)  | 15442 | 540 | 38 | 5 |
| [kubernetes-sigs/azurefile-csi-driv~](https://github.com/kubernetes-sigs/azurefile-csi-driver/)  | 15239 | 2443 | 886 | 113 |
| [ClickHouse/boost-extra](https://github.com/ClickHouse/boost-extra/)  | 15221 | 34 | 0 | 0 |
| [kubernetes-sigs/promo-tools](https://github.com/kubernetes-sigs/promo-tools/)  | 15135 | 1617 | 602 | 125 |
| [etcd-io/raft](https://github.com/etcd-io/raft/)  | 15020 | 1195 | 19 | 139 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-openstack/)  | 15014 | 1438 | 977 | 220 |
| [ClickHouse/ch-go](https://github.com/ClickHouse/ch-go/)  | 14953 | 1096 | 230 | 230 |
| [kubernetes/cloud-provider-aws](https://github.com/kubernetes/cloud-provider-aws/)  | 14921 | 975 | 401 | 294 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-ibmcloud/)  | 14588 | 727 | 840 | 55 |
| [openebs/jiva](https://github.com/openebs/jiva/)  | 14573 | 804 | 337 | 133 |
| [kubernetes-sigs/kind](https://github.com/kubernetes-sigs/kind/)  | 14569 | 3800 | 1596 | 11234 |
| [docker/libcompose](https://github.com/docker/libcompose/)  | 14354 | 673 | 361 | 587 |
| [docker/go-docker](https://github.com/docker/go-docker/)  | 14335 | 7 | 4 | 179 |
| [kubernetes/cloud-provider](https://github.com/kubernetes/cloud-provider/)  | 14187 | 1049 | 0 | 182 |
| [VictoriaMetrics/grafana-datasource](https://github.com/VictoriaMetrics/grafana-datasource/)  | 14133 | 189 | 29 | 23 |
| [falcosecurity/falco](https://github.com/falcosecurity/falco/)  | 13984 | 3557 | 1446 | 5681 |
| [kubernetes/component-base](https://github.com/kubernetes/component-base/)  | 13743 | 1264 | 0 | 84 |
| [kubernetes-sigs/node-feature-disco~](https://github.com/kubernetes-sigs/node-feature-discovery/)  | 13630 | 1550 | 832 | 541 |
| [docker/awesome-compose](https://github.com/docker/awesome-compose/)  | 13508 | 267 | 255 | 21867 |
| [prometheus/promlens](https://github.com/prometheus/promlens/)  | 13504 | 78 | 81 | 307 |
| [grafana/har-to-k6](https://github.com/grafana/har-to-k6/)  | 13480 | 627 | 84 | 77 |
| [istio/bots](https://github.com/istio/bots/)  | 13471 | 638 | 652 | 10 |
| [m3db/m3coordinator](https://github.com/m3db/m3coordinator/)  | 13259 | 68 | 54 | 4 |
| [m3db/m3ninx](https://github.com/m3db/m3ninx/)  | 13230 | 67 | 77 | 3 |
| [envoyproxy/playground](https://github.com/envoyproxy/playground/)  | 13086 | 192 | 168 | 7 |
| [go-faster/hx](https://github.com/go-faster/hx/)  | 13037 | 1473 | 12 | 4 |
| [prometheus/common](https://github.com/prometheus/common/)  | 12964 | 497 | 362 | 224 |
| [open-telemetry/opentelemetry-php-c~](https://github.com/open-telemetry/opentelemetry-php-contrib/)  | 12762 | 320 | 112 | 22 |
| [m3db/m3x](https://github.com/m3db/m3x/)  | 12696 | 208 | 206 | 17 |
| [etcd-io/bbolt](https://github.com/etcd-io/bbolt/)  | 12588 | 1164 | 220 | 6237 |
| [siderolabs/sidero](https://github.com/siderolabs/sidero/)  | 12377 | 366 | 839 | 257 |
| [docker/app](https://github.com/docker/app/)  | 12333 | 1544 | 650 | 1584 |
| [grafana/ksonnet](https://github.com/grafana/ksonnet/)  | 12180 | 434 | 0 | 0 |
| [m3db/m3db-operator](https://github.com/m3db/m3db-operator/)  | 11985 | 230 | 247 | 134 |
| [openebs/openebsctl](https://github.com/openebs/openebsctl/)  | 11906 | 112 | 105 | 27 |
| [kata-containers/agent](https://github.com/kata-containers/agent/)  | 11891 | 833 | 515 | 237 |
| [openebs/website](https://github.com/openebs/website/)  | 11856 | 1152 | 342 | 11 |
| [grafana/jslib.k6.io](https://github.com/grafana/jslib.k6.io/)  | 11563 | 189 | 71 | 32 |
| [open-telemetry/opentelemetry-erlang](https://github.com/open-telemetry/opentelemetry-erlang/)  | 11535 | 1144 | 355 | 253 |
| [prometheus/pushgateway](https://github.com/prometheus/pushgateway/)  | 11509 | 669 | 267 | 2514 |
| [kubernetes-sigs/blob-csi-driver](https://github.com/kubernetes-sigs/blob-csi-driver/)  | 11381 | 1663 | 645 | 97 |
| [openebs/zfs-localpv](https://github.com/openebs/zfs-localpv/)  | 11373 | 275 | 320 | 261 |
| [docker/getting-started](https://github.com/docker/getting-started/)  | 11356 | 190 | 212 | 2541 |
| [ClickHouse/clickhouse-presentations](https://github.com/ClickHouse/clickhouse-presentations/)  | 11213 | 515 | 39 | 834 |
| [m3db/m3msg](https://github.com/m3db/m3msg/)  | 11194 | 62 | 55 | 15 |
| [docker/engine-api](https://github.com/docker/engine-api/)  | 11164 | 9154 | 327 | 267 |
| [kubernetes-sigs/kubebuilder-declar~](https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/)  | 11087 | 591 | 262 | 188 |
| [cilium/proxy](https://github.com/cilium/proxy/)  | 10984 | 567 | 116 | 82 |
| [kubernetes-sigs/controller-tools](https://github.com/kubernetes-sigs/controller-tools/)  | 10903 | 773 | 455 | 549 |
| [envoyproxy/xds-relay](https://github.com/envoyproxy/xds-relay/)  | 10813 | 127 | 162 | 128 |
| [openebs/lvm-localpv](https://github.com/openebs/lvm-localpv/)  | 10735 | 182 | 144 | 141 |
| [open-telemetry/assign-reviewers-ac~](https://github.com/open-telemetry/assign-reviewers-action/)  | 10717 | 5 | 6 | 6 |
| [kubernetes/code-generator](https://github.com/kubernetes/code-generator/)  | 10645 | 1073 | 17 | 1394 |
| [grafana/jmeter-to-k6](https://github.com/grafana/jmeter-to-k6/)  | 10629 | 347 | 24 | 57 |
| [kubernetes/utils](https://github.com/kubernetes/utils/)  | 10611 | 893 | 233 | 256 |
| [kubernetes/node-problem-detector](https://github.com/kubernetes/node-problem-detector/)  | 10541 | 753 | 434 | 2258 |
| [grpc/grpc-kotlin](https://github.com/grpc/grpc-kotlin/)  | 10497 | 230 | 184 | 977 |
| [kubernetes/cli-runtime](https://github.com/kubernetes/cli-runtime/)  | 10483 | 913 | 5 | 240 |
| [grafana/tanka](https://github.com/grafana/tanka/)  | 10369 | 449 | 469 | 1896 |
| [istio/pkg](https://github.com/istio/pkg/)  | 10199 | 859 | 763 | 47 |
| [kubernetes-sigs/kube-batch](https://github.com/kubernetes-sigs/kube-batch/)  | 10197 | 1345 | 691 | 1035 |
| [istio/istio.io](https://github.com/istio/istio.io/)  | 10154 | 8352 | 10950 | 681 |
| [istio/old_mixerclient_repo](https://github.com/istio/old_mixerclient_repo/)  | 10004 | 228 | 409 | 15 |
| [kubernetes/kubeadm](https://github.com/kubernetes/kubeadm/)  | 9995 | 1023 | 520 | 3362 |
| [etcd-io/jetcd](https://github.com/etcd-io/jetcd/)  | 9809 | 1185 | 742 | 970 |
| [kubernetes-sigs/secrets-store-csi-~](https://github.com/kubernetes-sigs/secrets-store-csi-driver/)  | 9699 | 1195 | 747 | 939 |
| [istio/ztunnel](https://github.com/istio/ztunnel/)  | 9696 | 275 | 301 | 90 |
| [go-faster/jx](https://github.com/go-faster/jx/)  | 9667 | 1408 | 65 | 80 |
| [cilium/kube-apate](https://github.com/cilium/kube-apate/)  | 9590 | 18 | 0 | 4 |
| [grafana/cortex-tools](https://github.com/grafana/cortex-tools/)  | 9567 | 270 | 179 | 133 |
| [kubernetes-sigs/cloud-provider-hua~](https://github.com/kubernetes-sigs/cloud-provider-huaweicloud/)  | 9428 | 415 | 171 | 25 |
| [prometheus/statsd_exporter](https://github.com/prometheus/statsd_exporter/)  | 9406 | 847 | 310 | 823 |
| [kubernetes-sigs/gateway-api](https://github.com/kubernetes-sigs/gateway-api/)  | 9380 | 2300 | 1115 | 871 |
| [grafana/worldmap-panel](https://github.com/grafana/worldmap-panel/)  | 9361 | 232 | 74 | 294 |
| [openebs/cstor-csi](https://github.com/openebs/cstor-csi/)  | 9349 | 157 | 172 | 23 |
| [kubernetes-sigs/cri-tools](https://github.com/kubernetes-sigs/cri-tools/)  | 9232 | 1532 | 794 | 1217 |
| [openebs/jiva-operator](https://github.com/openebs/jiva-operator/)  | 9218 | 148 | 170 | 37 |
| [kubernetes/kompose](https://github.com/kubernetes/kompose/)  | 9201 | 1463 | 895 | 8309 |
| [ClickHouse/clickhouse-jdbc-bridge](https://github.com/ClickHouse/clickhouse-jdbc-bridge/)  | 9158 | 127 | 85 | 137 |
| [envoyproxy/go-control-plane](https://github.com/envoyproxy/go-control-plane/)  | 9128 | 1365 | 398 | 1287 |
| [grpc/test-infra](https://github.com/grpc/test-infra/)  | 9113 | 484 | 346 | 64 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-gcp/)  | 8963 | 1112 | 673 | 131 |
| [ClickHouse/libcxxabi](https://github.com/ClickHouse/libcxxabi/)  | 8810 | 14 | 4 | 0 |
| [kubernetes-sigs/apiserver-network-~](https://github.com/kubernetes-sigs/apiserver-network-proxy/)  | 8794 | 640 | 316 | 262 |
| [kubernetes/pod-security-admission](https://github.com/kubernetes/pod-security-admission/)  | 8784 | 465 | 0 | 77 |
| [ClickHouse/graphouse](https://github.com/ClickHouse/graphouse/)  | 8737 | 542 | 160 | 251 |
| [ClickHouse/clickhouse-connect](https://github.com/ClickHouse/clickhouse-connect/)  | 8723 | 229 | 84 | 86 |
| [kubernetes-sigs/cli-experimental](https://github.com/kubernetes-sigs/cli-experimental/)  | 8601 | 278 | 200 | 65 |
| [kubernetes/gengo](https://github.com/kubernetes/gengo/)  | 8416 | 455 | 182 | 485 |
| [kubernetes-sigs/cloud-provider-bai~](https://github.com/kubernetes-sigs/cloud-provider-baiducloud/)  | 8390 | 213 | 71 | 38 |
| [m3db/m3em](https://github.com/m3db/m3em/)  | 8125 | 25 | 19 | 1 |
| [openebs/mayastor-extensions](https://github.com/openebs/mayastor-extensions/)  | 7899 | 220 | 125 | 10 |
| [envoyproxy/ratelimit](https://github.com/envoyproxy/ratelimit/)  | 7866 | 150 | 235 | 1798 |
| [open-telemetry/opentelemetry-go-bu~](https://github.com/open-telemetry/opentelemetry-go-build-tools/)  | 7775 | 309 | 242 | 21 |
| [openebs/device-localpv](https://github.com/openebs/device-localpv/)  | 7745 | 55 | 49 | 15 |
| [kubernetes-sigs/krew](https://github.com/kubernetes-sigs/krew/)  | 7644 | 478 | 479 | 5396 |
| [prometheus/client_python](https://github.com/prometheus/client_python/)  | 7601 | 512 | 389 | 3254 |
| [grafana/grafana-polystat-panel](https://github.com/grafana/grafana-polystat-panel/)  | 7580 | 199 | 124 | 74 |
| [kubernetes/kube-aggregator](https://github.com/kubernetes/kube-aggregator/)  | 7488 | 2080 | 11 | 223 |
| [grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client/)  | 7470 | 402 | 135 | 71 |
| [kubernetes-sigs/aws-iam-authentica~](https://github.com/kubernetes-sigs/aws-iam-authenticator/)  | 7355 | 515 | 313 | 1968 |
| [kubernetes-sigs/aws-efs-csi-driver](https://github.com/kubernetes-sigs/aws-efs-csi-driver/)  | 7318 | 689 | 585 | 568 |
| [openebs/upgrade](https://github.com/openebs/upgrade/)  | 7223 | 122 | 135 | 10 |
| [grafana/postman-to-k6](https://github.com/grafana/postman-to-k6/)  | 7221 | 607 | 58 | 268 |
| [kubernetes-sigs/instrumentation-to~](https://github.com/kubernetes-sigs/instrumentation-tools/)  | 7192 | 92 | 6 | 25 |
| [docker/kitematic](https://github.com/docker/kitematic/)  | 7101 | 2335 | 534 | 12269 |
| [grpc/grpc-web](https://github.com/grpc/grpc-web/)  | 7078 | 888 | 622 | 7380 |
| [kubernetes/dns](https://github.com/kubernetes/dns/)  | 7033 | 710 | 318 | 802 |
| [grafana/opcua-datasource](https://github.com/grafana/opcua-datasource/)  | 7031 | 370 | 44 | 45 |
| [prometheus/mysqld_exporter](https://github.com/prometheus/mysqld_exporter/)  | 6994 | 605 | 412 | 1699 |
| [openebs/libcstor](https://github.com/openebs/libcstor/)  | 6959 | 77 | 89 | 14 |
| [open-telemetry/opentelemetry-demo](https://github.com/open-telemetry/opentelemetry-demo/)  | 6884 | 420 | 549 | 690 |
| [cilium/hubble](https://github.com/cilium/hubble/)  | 6843 | 796 | 757 | 2540 |
| [grafana/attic](https://github.com/grafana/attic/)  | 6818 | 426 | 0 | 1 |
| [envoyproxy/envoy-build-tools](https://github.com/envoyproxy/envoy-build-tools/)  | 6816 | 320 | 182 | 36 |
| [kubernetes-sigs/reference-docs](https://github.com/kubernetes-sigs/reference-docs/)  | 6815 | 492 | 233 | 65 |
| [grafana/kairosdb-datasource](https://github.com/grafana/kairosdb-datasource/)  | 6755 | 110 | 43 | 31 |
| [cilium/image-tools](https://github.com/cilium/image-tools/)  | 6736 | 239 | 192 | 11 |
| [helm/chartmuseum](https://github.com/helm/chartmuseum/)  | 6654 | 613 | 333 | 3119 |
| [etcd-io/dbtester](https://github.com/etcd-io/dbtester/)  | 6595 | 1226 | 288 | 250 |
| [kubernetes-sigs/apiserver-builder-~](https://github.com/kubernetes-sigs/apiserver-builder-alpha/)  | 6580 | 1070 | 434 | 717 |
| [envoyproxy/envoy-perf](https://github.com/envoyproxy/envoy-perf/)  | 6501 | 142 | 160 | 107 |
| [gotd/botapi](https://github.com/gotd/botapi/)  | 6402 | 653 | 313 | 18 |
| [kubernetes/csi-translation-lib](https://github.com/kubernetes/csi-translation-lib/)  | 6360 | 658 | 0 | 10 |
| [open-telemetry/opamp-go](https://github.com/open-telemetry/opamp-go/)  | 6328 | 92 | 110 | 60 |
| [kubernetes-sigs/kubetest2](https://github.com/kubernetes-sigs/kubetest2/)  | 6307 | 535 | 177 | 259 |
| [kubernetes-sigs/slack-infra](https://github.com/kubernetes-sigs/slack-infra/)  | 6124 | 108 | 44 | 85 |
| [kubernetes-sigs/sig-storage-local-~](https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner/)  | 6124 | 623 | 213 | 866 |
| [cilium/docsearch-scraper-webhook](https://github.com/cilium/docsearch-scraper-webhook/)  | 5993 | 38 | 32 | 3 |
| [siderolabs/kres](https://github.com/siderolabs/kres/)  | 5987 | 152 | 199 | 22 |
| [ClickHouse/bzip2](https://github.com/ClickHouse/bzip2/)  | 5964 | 141 | 0 | 1 |
| [prometheus/blackbox_exporter](https://github.com/prometheus/blackbox_exporter/)  | 5942 | 455 | 436 | 3484 |
| [grafana/devtools](https://github.com/grafana/devtools/)  | 5940 | 139 | 22 | 9 |
| [docker/go](https://github.com/docker/go/)  | 5884 | 24 | 8 | 17 |
| [kata-containers/govmm](https://github.com/kata-containers/govmm/)  | 5875 | 402 | 151 | 304 |
| [docker/volumes-backup-extension](https://github.com/docker/volumes-backup-extension/)  | 5873 | 228 | 76 | 47 |
| [istio/api](https://github.com/istio/api/)  | 5811 | 1685 | 2454 | 404 |
| [kubernetes/website](https://github.com/kubernetes/website/)  | 5767 | 39075 | 29119 | 3683 |
| [docker/metadata-action](https://github.com/docker/metadata-action/)  | 5744 | 318 | 157 | 563 |
| [grafana/azure-monitor-datasource](https://github.com/grafana/azure-monitor-datasource/)  | 5611 | 170 | 18 | 91 |
| [prometheus/compliance](https://github.com/prometheus/compliance/)  | 5606 | 147 | 79 | 105 |
| [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/)  | 5559 | 989 | 602 | 4546 |
| [kubernetes/component-helpers](https://github.com/kubernetes/component-helpers/)  | 5547 | 329 | 0 | 11 |
| [m3db/m3storage](https://github.com/m3db/m3storage/)  | 5511 | 38 | 15 | 3 |
| [openebs/api](https://github.com/openebs/api/)  | 5374 | 88 | 99 | 7 |
| [kubernetes-sigs/cluster-addons](https://github.com/kubernetes-sigs/cluster-addons/)  | 5371 | 261 | 89 | 147 |
| [prometheus/codemirror-promql](https://github.com/prometheus/codemirror-promql/)  | 5312 | 516 | 151 | 35 |
| [VictoriaMetrics/metricsql](https://github.com/VictoriaMetrics/metricsql/)  | 5287 | 102 | 6 | 136 |
| [kubernetes-sigs/aws-fsx-csi-driver](https://github.com/kubernetes-sigs/aws-fsx-csi-driver/)  | 5247 | 409 | 204 | 103 |
| [grafana/kubernetes-app](https://github.com/grafana/kubernetes-app/)  | 5231 | 199 | 23 | 397 |
| [helm/helm-classic](https://github.com/helm/helm-classic/)  | 5214 | 574 | 274 | 578 |
| [etcd-io/zetcd](https://github.com/etcd-io/zetcd/)  | 5152 | 135 | 58 | 1065 |
| [kubernetes-sigs/cluster-api-provid~](https://github.com/kubernetes-sigs/cluster-api-provider-digitalocean/)  | 5086 | 399 | 383 | 92 |
| [m3db/m3ctl](https://github.com/m3db/m3ctl/)  | 4984 | 82 | 64 | 14 |
| [grafana/gel-app](https://github.com/grafana/gel-app/)  | 4915 | 119 | 47 | 2 |
| [prometheus/snmp_exporter](https://github.com/prometheus/snmp_exporter/)  | 4779 | 553 | 416 | 1185 |
| [helm/monocular](https://github.com/helm/monocular/)  | 4741 | 399 | 362 | 1435 |
| [kata-containers/tests](https://github.com/kata-containers/tests/)  | 4723 | 4332 | 2783 | 135 |
| [gotd/contrib](https://github.com/gotd/contrib/)  | 4722 | 568 | 317 | 10 |
| [kubernetes-sigs/image-builder](https://github.com/kubernetes-sigs/image-builder/)  | 4634 | 2592 | 815 | 248 |
| [grafana/grafana-sensu-app](https://github.com/grafana/grafana-sensu-app/)  | 4610 | 16 | 45 | 8 |
| [istio/old_auth_repo](https://github.com/istio/old_auth_repo/)  | 4574 | 166 | 230 | 73 |
| [openebs/velero-plugin](https://github.com/openebs/velero-plugin/)  | 4537 | 119 | 137 | 43 |


