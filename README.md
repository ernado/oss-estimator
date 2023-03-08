# oss-estimator

Set of tools to get info about GitHub project that is needed to estimate the cost of running it.

## Tools

| Tool                                           | Description                                                                 |
|------------------------------------------------|-----------------------------------------------------------------------------|
| [estimator-dl](./cmd/estimator-dl/main.go)     | Download from [gharchive.org](https://gharchive.org)                        |
| [estimator-get](./cmd/estimator-get/main.go)   | Pull request and commit count                                               |
| [estimator-sloc](./cmd/estimator-sloc/main.go) | SLOC count using [scc](https://github.com/boyter/scc/) (should be in $PATH) |


## Example

```console
$ time estimator-sloc --org cilium --repo cilium
head 427461643b5c1493fb7b7662cb4501add3e446d3 refs/heads/master
Go 295919
ReStructuredText 32295
JSON 15093
C Header 13910
C 8972
Shell 7791
Go Template 2321
BASH 1818
Makefile 1617
Plain Text 1238
Protocol Buffers 984
Python 709
Jenkins Buildfile 545
Smarty Template 485
Dockerfile 457
License 344
SVG 190
gitignore 135
JavaScript 109
Systemd 97
Docker ignore 76
CSS 71
Lua 71
sed 63
Terraform 52
AWK 24
XML 14
Ruby 9
TOML 4
HTML 3

real    0m9,992s
user    0m1,679s
sys     0m0,419s
```