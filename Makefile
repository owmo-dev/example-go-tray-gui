ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ROOT=$(shell pwd)

run:
	go run main.go

build: build-darwin

build-darwin:
	source .env && cd build; sh build-darwin.sh