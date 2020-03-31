#!/usr/bin/env bash
.PHONY : release fmt bindata bench window linux mac clean tar

VERSION := 3.3.0

release:clean bindata window linux mac tar
	echo "compiled fanish"

fmt:
	gofmt -l -w ./

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

clean:
	rm -rf output/*

tar: 
	tar -czf output/go-mygen$(VERSION).window-amd64.tar.gz output/window/go-mygen.exe
	tar -czf output/go-mygen$(VERSION).linux-amd64.tar.gz output/linux/go-mygen
	tar -czf output/go-mygen$(VERSION).darwin-amd64.tar.gz output/mac/go-mygen
