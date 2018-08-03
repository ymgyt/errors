all: generate lint test

generate:
	go generate ./...

lint:
	gometalinter ./...

test:
	go test -v

.PHONY: generate lint test
