help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Setup some tools
	go get -u github.com/golang/lint/golint

lint: ## Run the Golint
	go vet
	golint -min_confidence 0.6 -set_exit_status
