.PHONY: lint py-lint

# Lint all Go Project Euler solutions
# Requires Go toolchain and golangci-lint on PATH
GO_MODULE_DIRS := $(dir $(wildcard Go/*/go.mod))

lint:
	@echo "Linting Go Project Euler modules with golangci-lint..."
	@for dir in $(GO_MODULE_DIRS); do \
		echo "==> $$dir"; \
		( cd "$$dir" && golangci-lint run ./... ) || exit $$?; \
	done

py-lint:
	@ruff check .
