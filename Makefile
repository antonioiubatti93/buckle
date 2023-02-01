MODULE_COVERAGE_OUT := module_coverage.out
TEST_REPORT ?= bin/junit.xml
TEST_FLAGS = -race -tags=integration -coverprofile=bin/$(MODULE_COVERAGE_OUT)
GOTESTSUM_FLAGS = --format testname --junitfile $(TEST_REPORT)

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
