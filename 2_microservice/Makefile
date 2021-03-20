export PATH := $(shell pwd)/bin:$(PATH)

.PHONY: run, build

build:
	@echo Building Movie Service...
	@echo Exporting go build env variable...
	@export CGO_ENABLED=0 GOOS=linux GO111MODULE=on
	@echo Build started...
	@/usr/local/go/bin/go build -o ./build/movie-service ./cmd/run_server/main.go
	@echo Build done
