# Define the Go binary
GO_BIN = go
PROTOC_BIN = protoc
PROTOC_GEN_GO_BIN = protoc-gen-go
PROTOC_GEN_GRPC_BIN = protoc-gen-go-grpc

# Define the protocol buffer source directory
PROTO_DIR = proto

# Define the output directory for generated files
OUT_DIR = .

# Define the protocol buffer file pattern
PROTO_FILES = $(PROTO_DIR)/**/*.proto

# Define the Go test and build flags
GO_TEST_FLAGS = -v
GO_BUILD_FLAGS =

# Targets
.PHONY: all run test generate clean build help

all: run

run:
	@echo "Running the application..."
	$(GO_BIN) run ./cmd/dev.go

test:
	@echo "Running tests..."
	$(GO_BIN) test $(GO_TEST_FLAGS) ./...

generate:
	@echo "Generating protocol buffers..."
	$(PROTOC_BIN) --go_out=$(OUT_DIR) --go_opt=paths=source_relative --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(PROTO_FILES)

clean:
	@echo "Cleaning up..."
	find . -name '*.pb.go' -type f -delete

# Optional target to build the binaries
build:
	@echo "Building the application..."
	$(GO_BIN) build $(GO_BUILD_FLAGS) ./cmd/...

# Help target
help:
	@echo "Developer Environment commands:"
	@echo "  all        - Run the application"
	@echo "  run        - Run the application"
	@echo "  test       - Run tests"
	@echo "  generate   - Generate protocol buffer code"
	@echo "  clean      - Clean up generated files"
	@echo "  build      - Build the application binaries"
	@echo "  help       - Display this help message"
