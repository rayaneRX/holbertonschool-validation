build: ## Generate the website from the markdown and configuration files in the directory dist/.
	hugo

clean: ## Cleanup the content of the directory dist/
	rm -rf dist/*

post: ## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME.
	hugo new --kind post posts/$(POST_NAME).md
	sed -i 's/title: .*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md

help: ## prints out the list of targets and their usage
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'
