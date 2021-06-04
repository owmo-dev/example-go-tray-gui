ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ROOT=$(shell pwd)

run:
	go run main.go

build: build-darwin build-windows build-linux

build-darwin:
	source .env && cd build; sh build-darwin.sh

build-windows:
	source .env && cd build; sh build-windows.sh

build-linux: docker-build-linux docker-clean

docker-build-linux:
	docker build -t example-tray-gui -f build/linux/Dockerfile .
	docker run -v $(ROOT)/bin:/example-tray-gui/bin -t example-tray-gui bash -c 'export VERSION=${VERSION} && export NAME=${NAME} && export NAME_LOWER=${NAME_LOWER} && cd build; bash build-linux.sh'

docker-clean:
	docker rm $(shell docker ps --all -q)
	docker rmi $(shell docker images | grep example-tray-gui | tr -s ' ' | cut -d ' ' -f 3)
