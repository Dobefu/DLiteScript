#!/usr/bin/env bash

# Exit on error.
set -e

# Navigate to the root of the project.
cd "$(dirname "$0")/.."

# Run the tests.
go test "./..." -coverprofile="coverage.out" -covermode=count -parallel="$(nproc)"

# Display the coverage statistics.
go tool cover -func coverage.out

# Generate and display an HTML report if the --coverage flag is provided.
if [ "$1" == "--coverage" ]; then
  go tool cover -html=coverage.out
fi
