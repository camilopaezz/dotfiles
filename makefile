.PHONY: build run clean install

# Build the dotfiles manager
build:
	go build -o bin/dotfiles ./src

# Run the dotfiles manager
run:
	go run ./src

# Clean build artifacts
clean:
	rm -rf bin/

# Install Go dependencies
deps:
	go mod download
	go mod tidy