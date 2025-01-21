# Variables
BUILD_DIR = dist
MAIN_FILE = main.go
ASSISTANT_APP_NAME = foo
VERSION = 0.1.0

# Targets
.PHONY: build clean run

build:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(ASSISTANT_APP_NAME) $(MAIN_FILE)
	sudo rm /usr/local/bin/$(ASSISTANT_APP_NAME)
	sudo cp $(BUILD_DIR)/$(ASSISTANT_APP_NAME) /usr/local/bin
