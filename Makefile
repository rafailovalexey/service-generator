# Tidy

tidy:
	@echo "Tidy..."
	@go mod tidy

# Download

download:
	@echo "Downloading dependencies..."
	@go mod download
