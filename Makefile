.DEFAULT_GOAL := help
.PHONY: 

da:=docker exec -d app

init:
	make build
	make start

init-dev:
	make build

build:
	make network/create
	docker-compose up -d --build
	docker ps -a

start:
	$(da) go run /usr/src/app/main.go

stop:
	docker stop app db
	docker rm app db
	make network/delete

restart:
	make stop
	make build
	make start

restart-dev:
	make stop
	make build

network/create:
	docker network create app-net
	docker network ls

network/delete:
	docker network rm app-net
	docker network ls
