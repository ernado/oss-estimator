all: list aggregate generate
list:
	go run ./cmd/estimator-list
pull:
	go run ./cmd/estimator-list --pull
force:
	go run ./cmd/estimator-list -f
aggregate:
	go run ./cmd/estimator-aggregate
generate:
	go generate
fetch: force aggregate generate
update: pull aggregate generate

.PHONY: all list pull force aggregate generate fetch update