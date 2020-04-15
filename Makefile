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