.PHONY: fmt
fmt:
	$(GO) mod tidy
	gofumpt -w .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race ./...
