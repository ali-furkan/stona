default: build

watch:
	@air

start:
	@go run alifurkan.co/storage

build:
	@go build -o main .

install:
	@go mod download
	@echo [Storage] - Dependencies downloaded

.PHONY: build start watch install