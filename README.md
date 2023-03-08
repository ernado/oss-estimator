# oss-estimator

Set of tools to get info about GitHub project that is needed to estimate the cost of running it.

## Tools

| Tool                                           | Description                                                                 |
|------------------------------------------------|-----------------------------------------------------------------------------|
| [estimator-dl](./cmd/estimator-dl/main.go)     | Download from [gharchive.org](https://gharchive.org)                        |
| [estimator-sloc](./cmd/estimator-sloc/main.go) | SLOC count using [scc](https://github.com/boyter/scc/) (should be in $PATH) |


## Example

```console
$ go run ./cmd/estimator-sloc
+-------------------+
| go-faster/jx      |
+---------+---------+
| SLOC    | 9667    |
| Commits | 1408    |
| PR      | 65      |
| HEAD    | 8da2e3c |
+---------+---------+
```