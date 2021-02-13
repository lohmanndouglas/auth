.PHONY: vendor run

vendor:
	go mod tidy
	go mod vendor

run: vendor
	GOFLAGS="-mod=vendor" realize start --run