ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ROOT=$(shell pwd)

run:
	go run main.go

build: build-darwin build-windows

build-darwin:
	source .env && cd build; sh build-darwin.sh

build-windows:
	source .env && cd build; sh build-windows.sh