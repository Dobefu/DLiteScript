#!/usr/bin/env bash

set -e
cd "$(dirname "$0")/.."

mkdir -p output

PLATFORMS=(
    "darwin-x64:darwin:amd64"
    "darwin-arm64:darwin:arm64"
    "linux-x64:linux:amd64"
    "linux-arm:linux:arm"
    "linux-arm64:linux:arm64"
    "win32-x64.exe:windows:amd64"
    "win32-arm.exe:windows:arm"
    "win32-arm64.exe:windows:arm64"
)

# Build the DLiteScript binary for all platforms.
for platform in "${PLATFORMS[@]}"; do
    IFS=':' read -r platform_key goos goarch <<< "${platform}"
    OUTPUT="output/dlitescript-${platform_key}"

    echo "Compiling ${platform_key}"
    GOOS="${goos}" GOARCH="${goarch}" go build -buildvcs -ldflags="-s -w" -o "${OUTPUT}" .
done
