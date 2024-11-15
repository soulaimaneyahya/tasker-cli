# Makefile

# Variables
SRC_DIR := ./src
DIST_DIR := ./dist
TODO_DIR := ./todo/*.txt
MAIN_BINARY := $(DIST_DIR)/main

# Default target
all: fmt build

# Format target
fmt:
	go fmt $(SRC_DIR)/*.go

# Golint
golint:
	golint $(SRC_DIR)/

# Build target
build: golint
	mkdir -p $(DIST_DIR)
	go build -o $(MAIN_BINARY) $(SRC_DIR)

# Run target
run: build
	$(MAIN_BINARY)

# Clean target
clean:
	rm -f $(MAIN_BINARY)
	rm -f $(TODO_DIR)
