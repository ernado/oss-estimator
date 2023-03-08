# oss-estimator

Set of tools to get info about GitHub project that is needed to estimate the cost of running it.

## Tools

| Tool                                           | Description                                                                 |
|------------------------------------------------|-----------------------------------------------------------------------------|
| [estimator-dl](./cmd/estimator-dl/main.go)     | Download from [gharchive.org](https://gharchive.org)                        |
| [estimator-sloc](./cmd/estimator-sloc/main.go) | SLOC count using [scc](https://github.com/boyter/scc/) (should be in $PATH) |
| [estimator-list](./cmd/estimator-list/main.go) | Concurrently fetch popular oss repos and stat for them                      |


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