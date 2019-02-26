#!/usr/bin/env bash

test:fmt clean build

dev:fmt clean build run fmt-output

all:fmt build run fmt-output

fmt:
	goimports -l -w ./

fmt-output:
	goimports -l -w ./output/

build:
	go build -o output/gm2m ./

run:
	output/gm2m

clean:
	rm -rf output/gm2m
	rm -rf output/markdown.md
