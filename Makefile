# Variables
BUILD_DIR = dist
MAIN_FILE = main.go
APP_NAME ?= $(shell grep APP_NAME .env | cut -d '=' -f2)
VERSION ?= $(shell grep VERSION .env | cut -d '=' -f2)

# Targets
.PHONY: build clean run

build:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin
