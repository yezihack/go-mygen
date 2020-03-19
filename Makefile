#!/usr/bin/env bash

dev:clean fmt bindata build run

all: dev con

fmt:
	gofmt -l -w ./

run:
	output/go-mygen

bindata:
	go-bindata -pkg main -o ./bindata.go assets/tpl

bench:
	go test -test.bench=".*"  -benchmem

win: bindata window
	output/window/go-mygen.exe help

linux:fmt
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=linux
	go build -a -o output/linux/go-mygen .
window:fmt
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=windows
	go build -a -o output/window/go-mygen.exe .
mac:fmt
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=darwin
	go build -a -o output/mac/go-mygen .
release:clean window linux mac tar
	echo "compiled fanish"

clean:
	rm -rf output/*

tar: 
	tar -czf output/go-mygen.window-amd64.tar.gz output/window/go-mygen.exe
	tar -czf output/go-mygen.linux-amd64.tar.gz output/linux/go-mygen
	tar -czf output/go-mygen.darwin-amd64.tar.gz output/mac/go-mygen
