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
	# docker rmi -f $(shell docker images -q)
	# docker volume rm $(shell docker volume ls -q)
	# docker network rm $(shell docker network ls -q)

.PHONY: proto proto-lint proto-breaking proto-format
proto:
	buf generate
	buf lint
	buf format -w

proto-lint:
	buf lint

proto-breaking:
	buf breaking --against '.git#branch=main' --path ./api/proto

proto-format:
	buf format -w

.PHONY: gqlgen
gqlgen:
	go run github.com/99designs/gqlgen generate

.PHONY: monitoring
monitoring:
	docker-compose -f docker-compose.yml -f docker-compose.monitoring.yml up -d

.PHONY: openapi
openapi:
	go generate ./api/rest/...

.PHONY: generate
generate: proto gqlgen openapi
