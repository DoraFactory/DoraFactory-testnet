#!/usr/bin/make -f

build: go.sum
	CGO_ENABLED=1 go build -mod=readonly  -o build/dogd ./cmd/dogd

install: go.sum
	CGO_ENABLED=1 go install -mod=readonly  ./cmd/dogd