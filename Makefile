SHELL := /bin/bash
MAKEFILE_PATH      := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_DIR        := $(shell dirname $(MAKEFILE_PATH))
SCRIPTS_DIR        := $(PROJECT_DIR)/scripts
IS_PKG_PATH=$(PROJECT_DIR)/internal/is/is.go
ASSERT_PATH=$(PROJECT_DIR)/pkg/declare/assert/assert.go
EXPECT_PATH=$(PROJECT_DIR)/pkg/declare/expect/expect.go


help:
	@echo "Please use 'make <TARGET>' where <TARGET> is one of:"
	@echo
	@echo "    declare         Write the declare package files."
	@echo "    help            Show this dialogue."
	@echo "    test            Run tests locally."
	@echo "    tidy            Tidy Go module and format Go files."
	@echo

declare:
	@if [[ ! -z "$(shell git status -s)" ]]; then \
		echo "Error: Uncommitted changes detected! Commit changes and try again."; \
	else \
		echo "Generating declare package files..."; \
		python3 $(SCRIPTS_DIR)/declare.py assert $(PROJECT_DIR)/internal/is/is.go $(PROJECT_DIR)/pkg/declare/assert/assert.go; \
		python3 $(SCRIPTS_DIR)/declare.py expect $(PROJECT_DIR)/internal/is/is.go $(PROJECT_DIR)/pkg/declare/expect/expect.go; \
	fi
	@if [[ -z "$$(git status -s)" ]]; then \
		echo "No changes applied."; \
	else \
		echo "Uncommitted changes:"; \
		git status -bs; \
	fi;

test:
	@echo "Running tests..."
	@go test ./...

tidy:
	@if [[ ! -z "$(shell git status -s)" ]]; then \
		echo "Error: Uncommitted changes detected! Commit changes and try again."; \
	else \
		echo "Running go mod tidy..."; \
		go mod tidy; \
		echo "Running go fmt..."; \
		for gofile in $$(find . -name '*.go'); do \
			go fmt $$gofile; \
		done; \
	fi
	@if [[ -z "$$(git status -s)" ]]; then \
		echo "No changes applied."; \
	else \
		echo "Uncommitted changes:"; \
		git status -bs; \
	fi;
