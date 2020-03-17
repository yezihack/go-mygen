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

linux:
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=linux
	go build -a -o output/linux/go-mygen .
window:
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=windows
	go build -a -o output/window/go-mygen.exe .
mac:
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=darwin
	go build -a -o output/mac/go-mygen .
release:clear window linux mac tar
	echo "compiled fanish"

clear:
	rm -rf output/*

tar: 
	tar -czf output/go-mygen.window-amd64.tar.gz output/window/go-mygen.exe
	tar -czf output/go-mygen.linux-amd64.tar.gz output/linux/go-mygen
	tar -czf output/go-mygen.darwin-amd64.tar.gz output/mac/go-mygen
