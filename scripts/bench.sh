#!/usr/bin/env bash

# Exit on error.
set -e

# Navigate to the root of the project.
cd "$(dirname "$0")/.."

# Run the tests.
go test -bench=. ./... -benchmem -run notest
