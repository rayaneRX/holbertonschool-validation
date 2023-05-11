lint: ## run the GolangCI-Lint tool
	golangci-lint run

build: ## Generate the website from the markdown and configuration files in the directory dist/
	make lint
	go build

run: ## Run the application in the background by executing the binary awesome-api
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application
	pkill awesome-api

clean: ## Stop the application, delete the binary awesome-api, and the log file awesome-api.log
	make stop || true
	rm -rf awesome-api awesome-api.log coverage-units.out coverage-integrations.out

test: ## Test the application
	make unit-tests
	make integration-tests

help: ## Print out the list of targets and their usage
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'

unit-tests: ## execute the Golang unit tests
	go test -v -short -coverprofile=coverage-units.out
	go tool cover -func=coverage-units.out

integration-tests: ## execute the Golang unit tests
	go test -v -coverprofile=coverage-integrations.out