#!/usr/bin/env bash

test:clean fmt build

dev:clean fmt bindata build run

all:clean fmt bindata build run

fmt:
	gofmt -l -w ./

build:
	go build -v -o output/gm2m ./

run:
	output/gm2m
vendor:
	govendor add +e
	govendor remove +u

bindata:
	go-bindata -debug -pkg gm2m -o ./bindata.go tpl/...
clean:
	rm -rf output/gm2m
	rm -rf output/markdown.md
	rm -rf output/db_entity/
	rm -rf output/db_models/
