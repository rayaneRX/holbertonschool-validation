build: ## Generate the website from the markdown and configuration files in the directory dist/.
	go build	

clean: ## Cleanup the content of the directory dist/
	rm -rf dist/*

run: ## Run the application in background by executing the binary awesome-api
	./awesome-api >./awesome-api.log 2>&1 &

help: ## prints out the list of targets and their usage
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'
