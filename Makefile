#!/usr/bin/env bash
.PHONY : release fmt bindata bench window linux mac clean tar

VERSION := 3.3.0

release:clean bindata window linux mac tar
	echo "compiled fanish"

dev:fmt bindata
	rm -rf output/dev/*
#	go build -a -o output/dev/go-mygen .
	go build -o output/dev/go-mygen .

fmt:
	gofmt -l -w ./

bindata:
	go-bindata -pkg main -o ./bindata.go assets/tpl

bench:
	go test -test.bench=".*"  -benchmem

win: bindata window
	output/window/go-mygen.exe help

linux:fmt
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o output/linux/go-mygen .

window:fmt
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o output/window/go-mygen.exe .

mac:fmt
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o output/mac/go-mygen .

clean:
	rm -rf output/*

tar: 
	tar -zcvf output/go-mygen$(VERSION).window-amd64.tar.gz output/window/go-mygen.exe
	tar -zcvf output/go-mygen$(VERSION).linux-amd64.tar.gz output/linux/go-mygen
	tar -zcvf output/go-mygen$(VERSION).darwin-amd64.tar.gz output/mac/go-mygen

