.PHONY: build run stop clean

build:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose down

clean:
	docker-compose down
	docker-compose rm -f
	docker rmi -f $(shell docker images -q)
	docker volume rm $(shell docker volume ls -q)
	docker network rm $(shell docker network ls -q)
