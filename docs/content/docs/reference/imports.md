+++
title = 'Imports'
linkTitle = 'Imports'
description = 'Import statements and module system in DLiteScript.'
weight = 0
draft = false
+++

The import system allows you to organize code across multiple files and
reuse functions and variables from other modules.

## Import Statement

Import statements load code from other DLiteScript files.

### Syntax

```go
import "path/to/file.dl"
import "path/to/file.dl" as alias
```

### Basic Import

Import a file using a relative or absolute path.
The filename (without extension) becomes the namespace:

```go
import "./utils.dl"

var result number = utils.add(5, 3)
```

In this example:

- Functions and variables from `utils.dl` are accessed via the `utils` namespace
- `utils.add()` calls the `add` function from that file

## Namespaces

By default, imports create a namespace based on the filename.

### Accessing Imported Items

Use the namespace followed by a dot to access functions and variables:

```go
import "./utils.dl"

var sum number = utils.add(1, 2)
var product number = utils.multiply(3, 4)
var pi number = utils.PI
```

### File Structure

**utils.dl:**

```go
func add(a number, b number) number {
  return a + b
}

func multiply(a number, b number) number {
  return a * b
}

var PI number = 3.14159
```

**main.dl:**

```go
import "./utils.dl"

printf("5 + 3 = %g\n", utils.add(5, 3))
printf("PI = %g\n", utils.PI)
```

## Import Aliases

You can specify a custom namespace using the `as` keyword.

### Custom Namespace

```go
import "./utilities.dl" as util

var result number = util.add(10, 20)
```

### Global Import

Use `_` as the alias to import into the global scope:

```go
import "./test.dl" as _

var result number = test() // No namespace prefix needed
```

With `as _`, all functions and variables are imported directly into the global scope.

## Import Paths

### Relative Paths

Relative paths are resolved relative to the current file:

```go
import "./utils.dl"        // Same directory
import "../shared/math.dl" // Parent directory
import "./lib/helpers.dl"  // Subdirectory
```

### Absolute Paths

You can also use absolute paths:

```go
import "/home/user/project/utils.dl"
```

## Module Organization

### Practical Example

**Project structure:**

```
project/
  main.dl
  utils.dl
  test.dl
```

**utils.dl:**

```go
func add(a number, b number) number {
  return a + b
}

func multiply(a number, b number) number {
  return a * b
}

var PI number = 3.14159
```

**test.dl:**

```go
func test() number {
  return 1
}
```

**main.dl:**

```go
import "./utils.dl"
import "./test.dl" as _

var x number = 5
var y number = 3

var sum number = utils.add(x, y)
var product number = utils.multiply(x, y)

printf("x = %g, y = %g\n", x, y)
printf("x + y = %g\n", sum)
printf("x * y = %g\n", product)
printf("PI = %g\n", utils.PI)
printf("test = %g\n", test())
```

## Important Notes

### File Extension

By convention, DLiteScript files use the `.dl` extension:

```go
import "./utils.dl"
```

### Import Order

Imports are evaluated when the import statement is encountered.
It's common practice to place all imports at the top of the file.
