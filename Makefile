SHELL := /bin/bash
MAKEFILE_PATH      := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_DIR        := $(shell dirname $(MAKEFILE_PATH))
SCRIPTS_DIR        := $(PROJECT_DIR)/scripts


help:
	@echo "Please use 'make <TARGET>' where <TARGET> is one of:"
	@echo
	@echo "    help            Show this dialogue."
	@echo "    tidy            Tidy Go module and format Go files."
	@echo


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
