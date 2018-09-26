.PHONY: all dep build clean test coverage coverhtml lint help

all: build

test-fizzbuzz:
	cd fizzbuzz && go test --cover

test-jump:
	cd jump_game && go test --cover

default:
	go test -v ./...

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test --cover -v ./...

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

coverhtml: ## Generate global code coverage report in HTML
	sh ./tools/coverage.sh html;
