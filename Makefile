# Variables
BUILD_DIR = dist
MAIN_FILE = main.go
APP_NAME = foo
VERSION = 0.1.0

# Targets
.PHONY: build clean run

build:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	sudo rm /usr/local/bin/$(APP_NAME)
	sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin
