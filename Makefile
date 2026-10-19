.DEFAULT_GOAL := help

.PHONY: help build test vet install

help: ## list targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-8s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## build ./zlay
	go build -o zlay ./cmd/zlay

test: ## go test ./...
	go test ./...

vet: ## go vet ./...
	go vet ./...

install: build ## install zlay to ~/.local/bin
	install -m 0755 zlay $(HOME)/.local/bin/zlay
