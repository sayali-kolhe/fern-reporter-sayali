# Makefile for a Gin-based Golang project using gox for cross-compilation

# Project Name
BINARY_NAME=fern

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOPKG=$(GOBASE)

# Go build and run commands
.PHONY: all build run clean cross-compile docker-build docker-run

all: build

build:
	@echo "🚀 Building..."
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(BINARY_NAME) $(GOPKG)

run:
	@echo "Running..."
	@GOBIN=$(GOBIN) ./bin/$(BINARY_NAME)

clean:
	@echo "🧹 Cleaning..."
	@GOBIN=$(GOBIN) go clean
	@rm -f $(GOBIN)/$(BINARY_NAME)


#Generate Go code from proto files
proto-generate:
	rm -rf ./fernreporter_pb
	mkdir ./fernreporter_pb
	protoc --go_out=./fernreporter_pb  \
       --go-grpc_out=./fernreporter_pb  \
       ./proto/fern-reporter.proto

# Cross-compilation with gox
cross-compile: proto-generate
	@echo "🛠️ Cross compiling for Linux and Mac..."
	@gox -osarch="linux/amd64 darwin/amd64" -output "$(GOBIN)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}" $(GOPKG)

# Testing
test:
	@echo "🧪 Running Tests..."
	@go test $(TEST_FLAGS) -coverprofile=profile.cov ./...

docker-build: cross-compile
	@echo "🐳 Building Local Docker image..."
	@docker build -t fern-app . -f Dockerfile-local

docker-run-local: docker-build
	@echo "🏃 Running application in Docker..."
	@docker run -it -p 8080:8080 -p 50051:50051 fern-app

docker-build-multi-arch:
	@echo "Building multi arch docker image and pushing..."
	@docker buildx build --platform linux/amd64,linux/arm64,linux/arm64/v8 -t anoop2811/fern-reporter:latest --push .

cover:
	@echo "📊 Running Tests with Coverage..."
	@go test -coverprofile=profile.cov ./...