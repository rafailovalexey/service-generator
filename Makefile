# Download

download:
	go mod download

# Build

build:
	go build -o build/main main.go

.PHONY: download, build
