MAKEFLAGS += --silent

all: clean build

## help: Prints a list of available build targets.
help:
	echo "Usage: make <OPTIONS> ... <TARGETS>"
	echo ""
	echo "Available targets are:"
	echo ''
	sed -n 's/^##//p' ${PWD}/Makefile | column -t -s ':' | sed -e 's/^/ /'
	echo
	echo "Targets run by default are: `sed -n 's/^all: //p' ./Makefile | sed -e 's/ /, /g' | sed -e 's/\(.*\), /\1, and /'`"


## test: Executes any tests.
test:
	go test -race -timeout 30s ./...

## build: Builds a custom 'k6' with the local extension. 
build:
	xk6 build --with $(shell go list -m)=.

## clean: Removes any previously created artifacts/downloads.
clean:
	echo "Cleaning up..."
	rm -f ./k6

.PHONY: test clean help build