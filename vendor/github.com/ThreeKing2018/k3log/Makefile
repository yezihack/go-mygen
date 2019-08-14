#!/usr/bin/env bash

fmt:
	goimports -l -w .

v:
	govendor add +e
	govendor remove +u

install: fmt clean
