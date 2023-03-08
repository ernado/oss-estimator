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
git head: 427461643b5c1493fb7b7662cb4501add3e446d3 refs/heads/master
languages:
 go 295919
 restructuredtext 32295
 json 15093
 c header 13910
 c 8972
 shell 7791
 go template 2321
 bash 1818
 makefile 1617
 plain text 1238
 protocol buffers 984
 python 709
 jenkins buildfile 545
 smarty template 485
 dockerfile 457
 license 344
 svg 190
 gitignore 135
 javascript 109
 systemd 97
 docker ignore 76
 css 71
 lua 71
 sed 63
 terraform 52
 awk 24
 xml 14
 ruby 9
 toml 4
 html 3
Total: 309739 SLOC
Languages that are counted: [Go C Go Template BASH Python]

real    0m9,992s
user    0m1,679s
sys     0m0,419s
```