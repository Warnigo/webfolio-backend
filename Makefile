GO = go
MAIN_GO = cmd/main.go
BINARY_NAME = main

.PHONY: all
all: build

# Build target
.PHONY: build
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_GO)

# Run target
.PHONY: run
run:
	$(GO) run $(MAIN_GO)

# Clean target
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
