#!/bin/bash
mkdir -p output
# build go binary
GOOS=$OS GOARCH=$ARCH go build -o output/inventory ./cmd/inventory
# build reactjs
cd web && npm run build && rm -rf ../output/static && cp -r build ../output/static