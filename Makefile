FILE_PATH=${shell pwd}


all: help

## build			Build project in this dir .
.PHONY: build
build:
	@echo "make build";
	go build -mod=vendor -o $(FILE_PATH)/weChatRobot  $(FILE_PATH)/main.go;
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -mod=vendor -o $(FILE_PATH)/weChatRobot.exe  $(FILE_PATH)/main.go;

## help			print this help message and exit.
.PHONY: help
help: Makefile
	@echo ""
	@echo "Choose a command run in "$(FILE_PATH)":"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Valid target values are:"
	@echo ""
	@sed -n 's/^## //p' $<