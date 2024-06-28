SHELL := /bin/bash
MAKEFILE_PATH      := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_DIR        := $(shell dirname $(MAKEFILE_PATH))
SCRIPTS_DIR        := $(PROJECT_DIR)/scripts
UNLESS_PATH=$(PROJECT_DIR)/internal/unless
ASSERT_PATH=$(PROJECT_DIR)/pkg/declare/assert
EXPECT_PATH=$(PROJECT_DIR)/pkg/declare/expect


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
		python3 $(SCRIPTS_DIR)/declare.py assert $(UNLESS_PATH) $(ASSERT_PATH); \
		python3 $(SCRIPTS_DIR)/declare.py expect $(UNLESS_PATH) $(EXPECT_PATH); \
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
