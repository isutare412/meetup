.DEFAULT_TARGET := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: tools
tools: ## Install tools for development.
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: docs
docs: ## Generate swagger docs.
	swag init \
		-d "${ROOTDIR}/cmd,${ROOTDIR}/pkg/controller/http,${ROOTDIR}/pkg/core/domain" \
		-o ${ROOTDIR}/api

.PHONY: run
run: ## Run server locally.
	@go run ./cmd/...

.PHONY: test
test: ## Test without infrastructures.
	@go test ./pkg/...

.PHONY: test-infra
test-infra: ## Test using infrastructures.
	@go test ./tests/...
