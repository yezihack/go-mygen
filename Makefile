#!/usr/bin/env bash

dev:clean fmt bindata build run

fmt:
	gofmt -l -w ./

build:
	go build -v -o main/output/go-mygen ./main/

run:
	main/output/go-mygen

vendor:
	govendor add +e
	govendor remove +u

bindata:
	go-bindata -debug -pkg gomygen -o ./bindata.go tpl/...
clean:
	rm -rf main/output/go-mygen
	rm -rf main/output/markdown.md
	rm -rf main/output/db_entity/
	rm -rf main/output/db_models/

con:
	main/output/go-mygen -h localhost -P 3308 -u root -p 123456 -d kindled
