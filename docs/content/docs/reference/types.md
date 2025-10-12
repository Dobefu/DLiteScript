+++
title = 'Types'
linkTitle = 'Types'
description = 'The available data types in the DLiteScript programming language.'
weight = 0
draft = false
+++

DLiteScript is a statically typed language, meaning every variable must have a declared type.

## Basic Types

### Number

The `number` type represents numeric values, including both integers and floating-point numbers.

#### Examples:

```go
var count number = 42
var pi number = 3.14159
var temperature number = -17.5
```

### String

The `string` type represents text data. Strings are enclosed in double quotes.

#### Examples:

```go
var name string = "Alice"
var message string = "Hello, world!"
var empty string = ""
```

### Bool

The `bool` type represents boolean values: `true` or `false`.

#### Examples:

```go
var isActive bool = true
var hasError bool = false
var isLessThanTwo bool = (1 < 2)
```

### Null

The `null` value represents the absence of a value.
It's used in comparisons and as a return value,
but cannot be directly assigned to typed variables.

When variables are declared without initialization, they receive zero values instead:
- `number` defaults to `0`
- `string` defaults to `""`
- `bool` defaults to `false`
- Arrays default to `[]`

#### Examples:

```go
var count number // count is 0, not null

var result = (null == null) // true

func maybeReturn() number {
  return null
}
```

## Special Types

### Any

The `any` type accepts values of any type.
This provides flexibility when the specific type isn't known at compile time.

#### Examples:

```go
var unknown any = 42
unknown = "now a string"
unknown = true
```

### Error

The `error` type is used for error handling and represents error values.

#### Examples:

```go
var err error // err has a zero value of null
```

## Composite Types

### Arrays

Arrays are declared using square brackets `[]` followed by the element type.
All elements in an array must be of the same type.

#### Syntax:

- `[]number` - Array of numbers
- `[]string` - Array of strings
- `[]bool` - Array of booleans

#### Examples:

```go
var numbers []number = [1, 2, 3, 4, 5]
var names []string = ["Alice", "Bob", "Charlie"]
var flags []bool = [true, false, true]
var empty []number = []
```

## Type Conversion

DLiteScript does not perform implicit type conversion. Types must match exactly in assignments and operations.

```go
var x number = 42
var y string = "42"

printf("%t\n", x == y) // false
```

