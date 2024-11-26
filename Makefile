BUILD_DIR = ./bin
BINARY = $(BUILD_DIR)/gospur

# CLI Commands
init: build
	@$(BINARY) init

version: build	
	@$(BINARY) version

# Development Commands
run: build
	@$(BINARY)

build: 
	@go build -o $(BINARY) .

test:
	@go test -v ./...

test-race:
	@go test -v ./... --race

# Only for local testing
release:
	@goreleaser release --snapshot --clean