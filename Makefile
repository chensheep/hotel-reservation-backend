BINARY_NAME=api

create_test_mongodb:
	docker run --name mongodb -d -p 27017:27017 mongo:latest

start_test_mongodb:
	docker start mongodb

stop_test_mongodb:
	docker stop mongodb

build:
	go build -o bin/$(BINARY_NAME)

run: build
	./bin/$(BINARY_NAME)

test:
	go test -v ./...

.PHONY: build run test 