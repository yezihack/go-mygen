#!/usr/bin/env bash
fmt:    
	gofmt -l -w ./

dev:fmt
	go build -v -o output/colorlog ./

test:
	go test -v

bench:
	go test -bench=. -benchmem
