# Tidy

tidy:
	@echo "Tidy..."
	@go mod tidy

# Download

download:
	@echo "Downloading dependencies..."
	@go mod download

# Run

run:
	@echo "Running..."
	@go run main.go

# Build

build:
	@echo "Building..."
	@go build -o build/main main.go

.PHONY: tidy, download, run, build
