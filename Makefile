default: build

watch:
	@air

start:
	@go run stona

build:
	@go build -o main .

install:
	@go mod download
	@echo [Storage] - Dependencies downloaded

.PHONY: build start watch install