#!/usr/bin/env bash

dev:clean fmt bindata build run

all: dev con

fmt:
	gofmt -l -w ./

build:
	go build -v -o go-mygen .

run:
	./go-mygen

vendor:
	govendor add +e
	govendor remove +u

bindata:
	go-bindata -pkg main -o ./bindata.go assets/tpl

clean:
	rm -rf go-mygen
	rm -rf output/markdown*
	rm -rf output/db_entity/
	rm -rf output/db_models/

con:
	./go-mygen -h localhost -P 3306 -u root -d mysql
	#./go-mygen -h 192.168.31.142 -P 3306 -u root -d sgfoot -p root

bench:
	go test -test.bench=".*"  -benchmem
