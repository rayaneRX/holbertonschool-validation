## Prerequisites
Same requirements as the previous task:

A Valid Go-Hugo website is provided
There are no Git Submodules
The theme ananke is installed
No directory dist/ committed
Makefile present
## Lifecycle
build:  Generate the website from the markdown and configuration files in the directory dist/
run:  Run the application in the background by executing the binary awesome-api
stop:  Stop the application
clean:  Stop the application, delete the binary awesome-api, and the log file awesome-api.log
test:  Test the application
help:  Print out the list of targets and their usage