.DEFAULT_GOAL := help
.PHONY: 

da:=docker exec -d app

init:
	make build
	make start

init-dev:
	make build

build:
	docker-compose up -d --build
	docker ps -a

start:
	$(da) go run /usr/src/app/main.go

stop:
	docker stop app db
	docker rm app db

restart:
	make stop
	make build
	make start

restart-dev:
	make stop
	make build

