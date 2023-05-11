build: check  ## publish the site
	hugo

clean:  ## empties the dist directory
	rm -rf dist/*

post:  ## Add a new page to your site with a specific title
	hugo new posts/$(POST_NAME).md
	sed -i "s/title: .*/title: \"$(POST_TITLE)\"/" content/posts/$(POST_NAME).md

check: ## Check lint of markdonws source files and analysis of the links
	markdownlint README.md content/posts/*.md
	markdown-link-check README.md content/posts/*.md

validate: ## validate the html file in dist/
	-python3 w3c_validator.py ./dist/index.html
	
help: ## Print out the list of targets and their usage
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'