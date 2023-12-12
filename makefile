# Makefile

.DEFAULT_GOAL := build

ARCH ?= amd64

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -o alert-plugin main.go email.go util.go

run:
	go run main.go

docker-build:
	docker build -t alert-plugin:v1 --platform=linux/$(ARCH) .

docker-run:
	docker run -p 8080:8080 --platform=linux/$(ARCH) alert-plugin:v1

clean:
	rm -f alert-plugin

.PHONY: build run docker-build docker-run clean
