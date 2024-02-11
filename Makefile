# Download

download:
	@echo "Downloading dependencies..."
	@go mod download

# Build

build:
	@echo "Building..."
	@go build -o build/main main.go

# Tidy

tidy:
	@echo "Tidy..."
	@go mod tidy

.PHONY: download, build, tidy
