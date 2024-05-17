# Go compiler
GO := go

# Output binary name
BINARY := bin/rss-aggregator

# Output folder
OUTPUT := bin

# Source files
SOURCES := $(wildcard *.go)

# Build target
build:
	@$(GO) build -o $(BINARY) $(SOURCES)

# Run target
run:	build
	@./$(BINARY)

# Clean target
clean:
	rm -rf $(OUTPUT)