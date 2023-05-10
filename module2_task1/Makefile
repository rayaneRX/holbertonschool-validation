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
	pkill awesome-api || true
	rm -rf awesome-api awesome-api.log

test: ## Test the application
	curl http://localhost:9999
	curl http://localhost:9999/health

help: ## Print out the list of targets and their usage
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'

