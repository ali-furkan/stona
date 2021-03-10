GO ?= go

default: build

watch:
	@air

start:
	@$(GO) run github.com/ali-furkqn/stona

build:
	@$(GO) build -o stona .

# I may will seperate it like core and cli 
release:
	@echo [Stona Release] - Building for darwin
	@GOOS=darwin GOARCH=amd64 $(GO) build -o bin/darwin/stona-amd64
	@echo [Stona Release] - Building for linux
	@GOOS=linux GOARCH=amd64 $(GO) build -o bin/linux/stona-amd64
	@GOOS=linux GOARCH=arm $(GO) build -o bin/linux/stona-arm
	@GOOS=linux GOARCH=arm64 $(GO) build -o bin/linux/stona-arm64
	@echo [Stona Release] - Building for windows
	@GOOS=windows GOARCH=amd64 $(GO) build -o bin/windows/stona-amd64.exe
	@GOOS=windows GOARCH=arm $(GO) build -o bin/windows/stona-arm.exe
	@GOOS=windows GOARCH=arm64 $(GO) build -o bin/windows/stona-arm64.exe

# NOTE: Add test

install:
	@echo [Stona Install] - Downloading dependencies...
	@go mod download
	@echo [Stona Install] - Dependencies downloaded

.PHONY: build start watch release install