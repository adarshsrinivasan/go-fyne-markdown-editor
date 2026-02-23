# Makefile for go-fyne-markdown-editor
# Cross-platform Markdown editor built with Fyne

.PHONY: all build run clean deps test fmt lint vet package help

# Project metadata
APP_NAME       := MarkDown
APP_VERSION    := 1.0.0
APP_BUILD      := 1
APP_ID         := com.example.$(APP_NAME)
MODULE         := github.com/adarshsrinivasan/go-fyne-markdown-editor

# Build flags
LDFLAGS        := -ldflags "-s -w"
RELEASE_LDFLAGS := -ldflags "-s -w -linkmode=external"

# Directories
BUILD_DIR      := .
DIST_DIR       := dist

# Default target
all: build

## build: Compile the application binary
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) $(LDFLAGS) .
	@echo "Built: $(BUILD_DIR)/$(APP_NAME)"

## run: Build and run the application
run: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)

## clean: Remove build artifacts and generated files
clean:
	@echo "Cleaning..."
	@rm -f $(APP_NAME) $(APP_NAME).exe
	@rm -rf $(APP_NAME).app
	@rm -rf $(DIST_DIR)
	@rm -rf *.iconset
	@echo "Clean complete."

## deps: Download and verify Go module dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod verify
	@go mod tidy
	@echo "Dependencies ready."

## test: Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

## test-cover: Run tests with coverage report
test-cover:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## fmt: Format Go source code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete."

## lint: Run static analysis (requires golangci-lint)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

## vet: Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

## package: Create distributable package using fyne CLI (may hang on some macOS setups)
package:
	@echo "Packaging with fyne CLI..."
	@fyne package --app-version $(APP_VERSION) --name $(APP_NAME) --release --icon Icon.png
	@echo "Package created: $(APP_NAME).app"

## help: Show this help message
help:
	@echo "go-fyne-markdown-editor - Makefile targets"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'
