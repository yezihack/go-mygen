#!/usr/bin/env bash

dev:clean fmt bindata build run

fmt:
	gofmt -l -w ./

build:
	go build -v -o gomygen/output/gomygen ./gomygen/

run:
	gomygen/output/gomygen

vendor:
	govendor add +e
	govendor remove +u

bindata:
	go-bindata -debug -pkg gomygen -o ./bindata.go tpl/...
clean:
	rm -rf gomygen/output/gomygen
	rm -rf gomygen/output/markdown.md
	rm -rf gomygen/output/db_entity/
	rm -rf gomygen/output/db_models/

con:
	gomygen/output/gomygen -h localhost -P 3308 -u root -p 123456 -d kindled
