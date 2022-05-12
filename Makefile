.DEFAULT: help

.PHONY: help
help:
	@echo 'Makefile for "holgerdocs" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  build-holgerdocs       - Build the hoglersync application.'
	@echo '  run-holgerdocs         - Run the holgerdocs subcommand on the test path.'
	@echo '  act                    - Run all GitHub actions workflows locally. (Requires Act installed)'

.PHONY: run-holgerdocs
run-holgerdocs:
	go run cmd/*.go 

.PHONY: build-holgerdocs
build-holgerdocs:
	go build -o holgerdocs ./...

.PHONY: act
act:
	act --workflows ".github/workflows/"