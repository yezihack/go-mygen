#!/usr/bin/env bash

test:fmt clean build

dev:fmt clean build run fmt-output

all:fmt build run fmt-output

fmt:
	goimports -l -w ./

fmt-output:
	goimports -l -w ./output/

build:
	go build -o output/m2m ./

run:
	output/m2m

clean:
	rm -rf output/m2m
	rm -rf output/markdown.md
