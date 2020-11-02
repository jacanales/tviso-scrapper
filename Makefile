.ONESHELL:

SHELL := /bin/bash
.SHELLFLAGS := -ec

help:
	@echo 'Usage: make [target] ...'
	@echo
	@echo 'targets:'
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s : | sort)"

.DEFAULT_GOAL := help

%:
	@echo 'Invalid target. type `make help` to get a list of available targets'
	@exit 1

up:
	docker-compose up -d

down:
	docker-compose down

download-wiremock: ## Download WireMock extensions to run tests
	@echo "downloading wiremock extensions ..."
	etc/bin/wiremock_extension.sh
	@echo "... done"

generate:
	go generate ./...

check-style: ## Run golangci check
	golangci-lint run ./...

coverage:
	go test -v -covermode=count -coverprofile=coverage.out ./...

go-install-golangci-lint: ## Install golangci-lint
	GOFLAGS="" go get \
    		github.com/golangci/golangci-lint/cmd/golangci-lint@'v1.32.0'

go-install-mockgen: ## Install golangci-lint
	GOFLAGS="" go get \
    		github.com/golang/mock/mockgen@'v1.4.4'

go-install-dev-vendors: go-install-golangci-lint ## Install modules for needed for development
	GOFLAGS="" go get \
		github.com/cespare/reflex \
		gotest.tools/gotestsum@'v0.6.0' \
		github.com/golang/mock/mockgen@'v1.4.4'
