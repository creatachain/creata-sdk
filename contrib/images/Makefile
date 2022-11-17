all: creatad-env

creatad-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag creatasdk/creatad-env creatad-env

.PHONY: all creatad-env
