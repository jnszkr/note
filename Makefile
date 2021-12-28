
.PHONY: all
all: lint test

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: test
test:
	go test ./... -cover