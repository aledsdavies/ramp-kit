# Makefile for a Go project using Air for live reloading and Templ for templating

# Variables
PROJECT_NAME := my-go-project
GO_FILES := $(shell find . -type f -name '*.go')
AIR := $(shell command -v air 2> /dev/null)
GO := $(shell command -v go 2> /dev/null)
TEMPL := $(shell command -v templ 2> /dev/null)

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all              Build the project"
	@echo "  build            Build the project"
	@echo "  dev              Run the project with live reloading (Air)"
	@echo "  clean            Clean the project"
	@echo "  test             Run tests"
	@echo "  deps             Install dependencies"
	@echo "  help             Display this help message"

# Build target
.PHONY: build
build: check-deps clean
	$(TEMPL) generate
	$(GO) build -o ./.tmp/main ./cmd/

# Run with Air for live reloading
.PHONY: dev
dev: check-deps clean
	make -j3 dev/templ dev/server dev/sync_assets

dev/server:
	$(AIR) \
	--build.pre_cmd "bun run scripts/bundle-css.mjs" \
	--build.cmd "go build -o .tmp/main ./cmd/ && templ generate --notify-proxy --proxyport=8090" \
	--build.bin ".tmp/main" \
	--build.delay "100" \
	--build.exclude_dir "frontend/node_modules,public/css,public/js" \
	--build.include_ext "go,css" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

dev/templ:
	$(TEMPL) generate --watch --proxyport=8090 --proxy="http://localhost:8080" --open-browser=false -v

# Clean target
.PHONY: clean
clean:
	$(GO) clean
	find ./views -name '*_templ.go' -type f -delete
	rm -rf .tmp/

# Run tests
.PHONY: test
test: check-deps
	$(GO) test ./...

# Install dependencies
.PHONY: deps
deps: check-deps
	$(GO) mod tidy

# Check for dependencies
.PHONY: check-deps
check-deps:
ifndef AIR
	$(error "Air is not installed. Please install it from https://github.com/cosmtrek/air")
endif
ifndef GO
	$(error "Go is not installed. Please install it from https://golang.org/")
endif
ifndef TEMPL
	$(error "Templ is not installed. Please install it from https://github.com/benbjohnson/templ")
endif

# Default target to display help message
.PHONY: all
all: help

