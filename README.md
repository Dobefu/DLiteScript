# DLiteScript

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_DLiteScript&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Dobefu_DLiteScript)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dobefu/DLiteScript)](https://goreportcard.com/report/github.com/Dobefu/DLiteScript)

> [!WARNING]
> This repository is still a work-in-progress. It is nowhere near production-ready.

## Usage

- Run the application with a file, e.g.:

  ```bash
  go run main.go examples/00_simple/main.dl
  ```

## Supported constants

- `PI` - π
- `TAU` - τ (2π)
- `E` - Euler's number
- `PHI` - φ Golden ratio
- `LN2` - Natural logarithm of 2
- `LN10` - Natural logarithm of 10

## Supported functions

- `abs(x)` - Absolute value of `x`
- `sin(x)` - Sine value of `x`
- `cos(x)` - Cosine value of `x`
- `tan(x)` - Tangent value of `x`
- `sqrt(x)` - Square root
- `round(x)` - Round `x` to the nearest integer value
- `floor(x)` - Round `x` down to the nearest integer value
- `ceil(x)` - Round `x` up to the nearest integer value
- `min(x, y)` - Get the smallest of the values provided
- `max(x, y)` - Get the largest of the values provided
