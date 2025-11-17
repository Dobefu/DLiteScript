<p align="center">
    <img
        align="center"
        alt="Logo"
        aria-hidden="true"
        src="/media/logo.svg"
        width="64"
    />
</p>

<h1 align="center">DLiteScript</h1>

<p align="center">A delightfully simple scripting language.</p>

<p align="center">
    <a href="https://go.dev/">
        <img
            alt="Go Version"
            src="https://img.shields.io/github/go-mod/go-version/Dobefu/DLiteScript"
        /></a>
    <a href="https://go.dev/">
        <img
            alt="License"
            src="https://img.shields.io/github/license/Dobefu/DLiteScript"
        /></a>
    <a href="https://sonarcloud.io/summary/new_code?id=Dobefu_DLiteScript">
        <img
            alt="Quality Gate Status"
            src="https://sonarcloud.io/api/project_badges/measure?project=Dobefu_DLiteScript&metric=alert_status"
        /></a>
    <a href="https://sonarcloud.io/summary/overall?id=Dobefu_DLiteScript">
        <img
            alt="Coverage"
            src="https://sonarcloud.io/api/project_badges/measure?project=Dobefu_DLiteScript&metric=coverage"
        /></a>
    <a href="https://goreportcard.com/report/github.com/Dobefu/DLiteScript">
        <img
            alt="Go Report Card"
            src="https://goreportcard.com/badge/github.com/Dobefu/DLiteScript"
        /></a>
    <a href="https://discord.gg/KBAykpBgXR">
        <img
            alt="Discord"
            src="https://dcbadge.limes.pink/api/server/KBAykpBgXR?style=flat"
        /></a>
    </p>

> [!WARNING]
> This repository is still a work-in-progress. It is nowhere near production-ready.

## Table of Contents

- [Features](#features)
- [Quick Start](#quick-start)
- [Examples](#examples)
- [Installation](#installation)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- Static typing
- No magic
- LSP support
- Built-in formatter
- Built-in REPL
- Full support for [VSCode](https://github.com/Dobefu/vscode-dlitescript)
- Full support for [Neovim](https://github.com/Dobefu/nvim-dlitescript)
- Extensive [Documentation](https://dlitescript.com) with an online [playground](https://dlitescript.com/playground/)

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
go build -buildvcs
```

## Development

```bash
# Run tests
go test ./...

# Build binary
go build -buildvcs
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

MIT License - see [LICENSE](LICENSE) for details.
