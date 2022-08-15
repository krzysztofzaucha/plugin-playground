export SHELL:=/bin/bash
export BASE_NAME:=$(shell basename ${PWD})
export IMAGE_BASE_NAME:=kz/${BASE_NAME}
export NETWORK:=${BASE_NAME}-network

default: help

help: ## Prints help for targets with comments
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-50s\033[0m %s\n", $$1, $$2}'
	@echo ""

###############
# Development #
###############

analysis: ## Run linter
	@docker run --rm \
		-v ${PWD}:${PWD} \
		-w ${PWD} \
		golangci/golangci-lint:v1.47.2 \
		golangci-lint run -v --config .golangci.yml

format-install: ## Install go formatting dependencies on the localhost
	@go install golang.org/x/tools/cmd/goimports@latest \
	 	&& go install mvdan.cc/gofumpt@latest

format: format-install ## Format the code
	@go mod tidy \
		&& goimports -w . \
		&& gofumpt -l -w .

#########
# Build #
#########

build-base: ## Build base image (without any plugins)
	@docker image build \
		--target=final-base \
		-t ${IMAGE_BASE_NAME}-base:latest .

build-plugin-command: ## Build command plugin image
	@docker image build \
		--target=final-plugin-command \
		-t ${IMAGE_BASE_NAME}-plugin-command:latest .

build-plugin-web-server: ## Build web-server plugin image
	@docker image build \
		--target=final-plugin-web-server \
		-t ${IMAGE_BASE_NAME}-plugin-web-server:latest .

#######
# Run #
#######

run-command: ## Run command plugin example
	@docker run --rm \
		${IMAGE_BASE_NAME}-plugin-command:latest

run-web-server: ## Run web-server plugin example
	@docker run --rm \
		${IMAGE_BASE_NAME}-plugin-web-server:latest

###############
# Danger Zone #
###############

reset: ## Cleanup
	@docker stop $(shell docker ps -aq) || true
	@docker system prune || true
	@docker volume rm $(shell docker volume ls -q) || true
	@docker rmi -f ${IMAGE_BASE_NAME}-base:latest || true # remove docker base image
	@docker rmi -f ${IMAGE_BASE_NAME}-plugin-command:latest || true # remove docker command image
	@docker rmi -f ${IMAGE_BASE_NAME}-plugin-web-server:latest || true # remove docker web-server image
