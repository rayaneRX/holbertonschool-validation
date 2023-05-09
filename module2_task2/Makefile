lint: ## run the GolangCI-Lint tool
	golangci-lint run

build: ## Generate the website from the markdown and configuration files in the directory dist/
	go build

run: ## Run the application in the background by executing the binary awesome-api
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application
	kill "$(pgrep awesome-api)"

clean: ## Stop the application, delete the binary awesome-api, and the log file awesome-api.log
	pkill awesome-api
	rm -rf awesome-api awesome-api.log

test: ## Test the application
	curl http://localhost:9999
	curl http://localhost:9999

help: ## Print out the list of targets and their usage
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'

unit-tests: ## Run Golang unit tests with code coverage
	go test -v -short -coverprofile=coverage-units.out
	