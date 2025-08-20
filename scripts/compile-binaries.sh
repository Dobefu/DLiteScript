#!/usr/bin/env bash

# Exit on error.
set -e

# Navigate to the root of the project.
cd "$(dirname "$0")/.."

# Create the output directory if it doesn't already exist.
mkdir -p output

# Set the target platforms.
PLATFORMS=(
    "darwin-x64:darwin:amd64"
    "darwin-arm64:darwin:arm64"
    "linux-x64:linux:amd64"
    "linux-arm:linux:arm"
    "linux-arm64:linux:arm64"
    "win32-x64:windows:amd64"
    "win32-arm:windows:arm"
    "win32-arm64:windows:arm64"
)

# Build the DLiteScript binary for all platforms.
for platform in "${PLATFORMS[@]}"; do
    IFS=':' read -r platform_key goos goarch <<< "$platform"
    OUTPUT="output/dlitescript-${platform_key}"

    echo "Compiling ${platform_key}"
    GOOS=$goos GOARCH=$goarch go build -ldflags="-s -w" -o "$OUTPUT" .
done
