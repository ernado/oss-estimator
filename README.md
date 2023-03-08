# oss-estimator

Set of tools to get info about GitHub project that is needed to estimate the cost of running it.

Tools:

| Tool                                   | Description                                                                 |
|----------------------------------------|-----------------------------------------------------------------------------|
| [estimator-dl](./cmd/estimator-dl)     | Download from [gharchive.org](https://gharchive.org)                        |
| [estimator-get](./cmd/estimator-get)   | Pull request and commit count                                               |
| [estimator-sloc](./cmd/estimator-sloc) | SLOC count using [scc](https://github.com/boyter/scc/) (should be in $PATH) |