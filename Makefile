default:
	go run ./cmd/estimator-list/ && go run ./cmd/estimator-aggregate/ && go generate
pull:
	go run ./cmd/estimator-list/ --pull && go run ./cmd/estimator-aggregate/ && go generate
force:
	go run ./cmd/estimator-list/ -f && go run ./cmd/estimator-aggregate/ && go generate
gen:
	go run ./cmd/estimator-aggregate/ && go generate