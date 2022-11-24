# Makefile.Common defines the target that each module
# must use. It must be included by each sub-module.
include ./Makefile.Common

GO_FILES := $(shell git ls-files ./**/*.go)

# Silent mode to reduce output.
MAKE_FLAGS ?= -s
MAKECMD ?= make $(MAKE_FLAGS)

COVERAGE_REPORT ?= bin/cover.xml
COVERAGE_OUT = $(COVERAGE_REPORT:.xml=.out)

# Eve is a multi-modular repository. For Go to find
# all modules within the project, we look for go.mod files.
# The set of go modules includes the found directories
# and the root module.

ALL_MODULES := $(shell find . -type f -name "go.mod" -exec dirname {} \; | sort | egrep  '^./' )
GO_MODULES = $(ALL_MODULES) $(PWD)

# List all modules within the repository for inspection.
all-modules:
	@echo $(ALL_MODULES) | tr ' ' '\n' | sort

# Define a delegation target for each module.
.PHONY: $(GO_MODULES)
$(GO_MODULES):
	$(MAKECMD) -C $@ $(TARGET)

# Trigger each module's delegation target.
.PHONY: for-all-target
for-all-target: $(GO_MODULES)

.PHONY: build
build: bin/buckle

bin/buckle: $(GO_FILES) go.mod go.sum
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $@

.PHONY: fmt
fmt:
	$(MAKECMD) for-all-target TARGET="module-fmt"

# make test also produces a code coverage per each module
# and merges all reports into a single one.
.PHONY: test
test: 
	$(MAKECMD) for-all-target TARGET="module-test"
	$(MAKECMD) merge-coverage
	gocov convert $(COVERAGE_OUT) | gocov-xml > $(COVERAGE_REPORT)

# merge-coverage looks for each module's coverage report
# and merges them into a single report named after the
# globally set variable (e.g. Jenkins worker).
.PHONY: merge-coverage
merge-coverage:
	gocovmerge $$(find . -name $(MODULE_COVERAGE_OUT)) > $(COVERAGE_OUT)

.PHONY: dev-tools
dev-tools:
	$(GO) install mvdan.cc/gofumpt@latest

# Tag creates and pushes a tag to the remote.
# It takes two arguments TAG and MSG to be specified
# by command line.

tag:
	git tag -s -a $$TAG -m "$$MSG"
	git push origin refs/tags/$$TAG