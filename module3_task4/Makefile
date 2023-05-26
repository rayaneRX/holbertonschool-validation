COVERAGE=coverage
lint: ## Does the linting
	markdownlint DEPLOY.md README.md || true


build:  ## publish the site
	hugo
	go build

run: ## Run the application in background, and write logs into a file named awesome-api.log
	./awesome-api >./awesome-api.log 2>&1 &

stop: ##  Stop the application
	pkill awesome-api

clean:  ## empties the dist directory
	rm -rf dist/
	pkill awesome-api || true
	rm -rf awesome-api awesome-api.log coverage-integrations.out coverage-units.out

post:  ## Add a new page to your site with a specific title
	hugo new posts/$(POST_NAME).md
	sed -i "s/title: .*/title: \"$(POST_TITLE)\"/" content/posts/$(POST_NAME).md

test: ## Tests for browser responses
	make unit-tests
	make integration-tests
	make validate

unit-tests: ## Run unit tests
	go test -v -short -coverprofile=$(COVERAGE)-units.out

integration-tests: ## Run integration tests
	go test -v -coverprofile=$(COVERAGE)-integrations.out

check: ## Check lint of markdonws source files and analysis of the links
	markdownlint content/posts/welcome.md
	markdown-link-check content/posts/welcome.md

validate: ## validate the html file in dist/
	-python3 w3c_validator.py -f errors.txt dist/index.html
	
help: ## Prints out the list of targets and their usage
	@grep "##" Makefile | grep -v grep | sed "s/##//"

package: ## produces a file awesome-website.zip
	zip -r awesome-website.zip awesome-api dist/
