#!/usr/bin/env bash

set -e
cd "$(dirname "$0")/.."

go test -bench=. ./... -benchmem -run notest
