#!/bin/bash
# create output folder
mkdir -p output
rm -rf ./output/*

# build go binary into output folder
# GOOS=$OS GOARCH=$ARCH go build -o output/inventory ./cmd/inventory
GOOS=linux GOARCH=amd64 go build -o output/inventory-linux-amd64 ./cmd/inventory
GOOS=windows GOARCH=amd64 go build -o output/inventory-win-amd64.exe ./cmd/inventory
GOOS=windows GOARCH=386 go build -o output/inventory-win-386.exe ./cmd/inventory

# build reactjs into output folder
cd web && npm run build && rm -rf ../output/static && cp -r build ../output/static
