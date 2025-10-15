#!/usr/bin/env bash

set -e
cd "$(dirname "$0")/.."

cd assets/playground
GOOS=js GOARCH=wasm go build

# Hugo takes care of this if the file doesn't exist,
# but we want to update the file in case it changes.
cp DLiteScript-playground-wasm ../../public/playground/DLiteScript-playground-wasm

echo "WASM binary built successfully"