#!/bin/bash

VERSION=$(git rev-parse --short HEAD)
# CGO_ENABLED=1
# create output folder
mkdir -p output/$VERSION

rm -rf output/$VERSION/*

# build go binary into output folder
# GOOS=$OS GOARCH=$ARCH go build -o output/inventory ./cmd/inventory
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "output/$VERSION/inventory-linux-amd64-$VERSION" ./cmd/inventory
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "output/$VERSION/inventory-win-amd64-$VERSION.exe" ./cmd/inventory
CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o "output/$VERSION/inventory-win-386-$VERSION.exe" ./cmd/inventory

# build reactjs into output folder
cd web && npm run build && rm -rf "../output/$VERSION/static" && cp -r build "../output/$VERSION/static"
