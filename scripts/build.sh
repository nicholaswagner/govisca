#!/bin/bash

GOPATH=$(PWD)
echo "GOPATH=$GOPATH"

build=$(git rev-parse --short HEAD)
echo "Build SHA: $build"

echo "building ..."
go build -ldflags "-X main.Build=$build" src/github.com/nicholaswagner/main.go