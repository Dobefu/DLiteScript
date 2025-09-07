# DLiteScript

[![Go Version](https://img.shields.io/github/go-mod/go-version/Dobefu/DLiteScript)](https://golang.org/)
[![License](https://img.shields.io/github/license/Dobefu/DLiteScript)](https://golang.org/)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_DLiteScript&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Dobefu_DLiteScript)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_DLiteScript&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Dobefu_DLiteScript)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dobefu/DLiteScript)](https://goreportcard.com/report/github.com/Dobefu/DLiteScript)

> [!WARNING]
> This repository is still a work-in-progress. It is nowhere near production-ready.

A delightfully simple scripting language.

## Table of Contents

- [Features](#features)
- [Quick Start](#quick-start)
- [Examples](#examples)
- [Installation](#installation)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Static typing**: Explicit types with types like `string`, `number`, and `bool`
- **No magic**: All functionality is visible and explicit
- **LSP support**: Language Server Protocol for IDE integration

## Quick Start

### Hello World

```dlitescript
printf("Hello, World!\n")
```

Run it:

```bash
dlitescript hello.dl
```

### Variables and Types

```dlitescript
var name string = "DLiteScript"
var version number = 0.1
var isSupported bool = true

printf("Welcome to %s v%g\n", name, version)
printf("Supported: %t\n", isActive)
```

### Control Flow

```dlitescript
var test number = 85

if test > 80 {
  printf("The number is over 80\n")
} else {
  printf("The number is 80 or lower\n")
}
```

### Loops

```dlitescript
for var i from 0 to 5 {
  printf("Count: %g\n", i)
}

for var i to 3 {
  if i == 2 {
    continue
  }
  printf("Iteration %g\n", i)
}
```

### Functions

```dlitescript
func greet(name string) string {
  printf("Hello, %s!\n", name)
}

greet("Developer")
```

## Examples

Check out the `examples/` directory for examples covering:

- Basic syntax and operations
- Variables and type declarations
- Control flow and loops
- Functions with multiple return values
- And much more!

## Installation

```bash
git clone https://github.com/Dobefu/DLiteScript.git
cd DLiteScript
go mod download
go build -o dlitescript ./cmd
```

## Development

```bash
# Run tests
go test ./...

# Build binary
go build -o dlitescript .
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

MIT License - see [LICENSE](LICENSE) for details.
