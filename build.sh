#!/bin/bash

set -e

targets=(
    "linux/amd64/"
    "linux/arm64/"
    "linux/386/"
    "darwin/amd64/"
    "darwin/arm64/"
    "windows/amd64/.exe"
    "windows/386/.exe"
)

for target in "${targets[@]}"; do
    mkdir -p "build"

    IFS='/' read -r os arch ext <<< "${target}"
    to="build/webfetch-${os}-${arch}${ext}"
    from="cmd/webfetch/main.go"
    
    CGO_ENABLED=0 GOOS="${os}" GOARCH="${arch}" go build -o "${to}" "${from}"

    echo "Build for ${os}/${arch} finished."
done