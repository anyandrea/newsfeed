.PHONY: run dev binary setup glide start-mysql stop-mysql test
SHELL := /bin/bash

all: run

run: binary
	scripts/run.sh

dev: stop-mysql start-mysql
	scripts/dev.sh

binary:
	GOARCH=amd64 GOOS=linux go build -i -o newsfeed

setup:
	go get -v -u github.com/codegangsta/gin
	go get -v -u github.com/Masterminds/glide

glide:
	glide install --force

start-mysql:
	docker run --name newsfeeddb \
		-e MYSQL_ROOT_PASSWORD=blibb \
		-e MYSQL_DATABASE=newsfeed_db \
		-e MYSQL_USER=blubb \
		-e MYSQL_PASSWORD=blabb \
		-p "3306:3306" \
		-d mariadb:10
	# scripts/db_setup.sh

stop-mysql:
	docker kill newsfeeddb || true
	docker rm -f newsfeeddb || true

test:
	GOARCH=amd64 GOOS=linux go test $$(go list ./... | grep -v /vendor/)
