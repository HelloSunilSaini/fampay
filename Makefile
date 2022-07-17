#!/bin/bash
build:
	go build -o fampay

test:
	go test ./...

docker_build:
	docker build .

run:
	docker compose up --build
