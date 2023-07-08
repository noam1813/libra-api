.DEFAULT_GOAL := help
.PHONY: 

da:=docker exec -d app

init:
	make build
	make start

build:
	docker-compose up -d --build
	docker ps -a

start:
	$(da) go run /usr/src/app/main.go

stop:
	docker stop libra-api

restart:
	make stop
	make build
	make start

