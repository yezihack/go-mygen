#!/usr/bin/env bash

dev:clean fmt bindata build run

fmt:
	gofmt -l -w ./

build:
	go build -v -o gm2m/output/gm2m ./gm2m/

run:
	gm2m/output/gm2m

vendor:
	govendor add +e
	govendor remove +u

bindata:
	go-bindata -debug -pkg gm2m -o ./bindata.go tpl/...
clean:
	rm -rf gm2m/output/gm2m
	rm -rf gm2m/output/markdown.md
	rm -rf gm2m/output/db_entity/
	rm -rf gm2m/output/db_models/

con:
	gm2m/output/gm2m -h localhost -P 3308 -u root -p 123456 -d kindled
