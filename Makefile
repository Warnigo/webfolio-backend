GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
BINARY_NAME=main

build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go

clean:
	$(GOCLEAN)

test:
	$(GOTEST) ./...

fmt:
	$(GOFMT) -s -w .

all: fmt build test

run: build
	./$(BINARY_NAME)

reset: clean all
