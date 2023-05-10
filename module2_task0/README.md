## Prerequisites
Same requirements as the previous task:

A Valid Go-Hugo website is provided
There are no Git Submodules
The theme ananke is installed
No directory dist/ committed
Makefile present
## Lifecycle
build:  Generate the website from the markdown and configuration files in the directory dist/.
clean:  Cleanup the content of the directory dist/
post:  Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME.
test: Test the application
help:  prints out the list of targets and their usage
