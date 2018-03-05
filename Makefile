SHELL = /bin/bash
VERSION = 1.0.0
NAME = adventure-plan.com
ENV = dev
PORT = 8080

.PHONY: build_and_serve
build_and_serve: build serve

.PHONY: build
build: build_server

.PHONY: build_server
build_server:
	cd cmd/api && go build .

.PHONY: build_linux_server
build_linux_server:
	cd cmd/api && GOOS=linux go build .

.PHONY: build_app
build_app:
	cd ui && ionic build

.PHONY: serve
serve:
	./cmd/api/api

.PHONY: docker_build
docker_build: build_linux_server build_app
	docker build -t philbrookes/adventureplan:${VERSION} .

.PHONY: docker_serve
docker_serve: docker_build
	docker run -e "ENV=${ENV}" -p ${PORT}:${PORT} philbrookes/adventureplan:${VERSION}

.PHONY: docker_dev_serve
docker_dev_serve: docker_build
	docker run -e "ENV=${ENV}" -p ${PORT}:${PORT} philbrookes/adventureplan:${VERSION}

.PHONY: docker_dev_server_reload
docker_dev_server_reload: build_linux_server
	docker build -t philbrookes/adventureplan:${VERSION} .
	docker run -e "ENV=${ENV}" -p ${PORT}:${PORT} philbrookes/adventureplan:${VERSION}

.PHONY: docker_dev_ui_reload
docker_dev_ui_reload: build_app
	docker build -t philbrookes/adventureplan:${VERSION} .
	docker run -e "ENV=${ENV}" -p ${PORT}:${PORT} philbrookes/adventureplan:${VERSION}
