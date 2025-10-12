+++
title = 'Variables and Constants'
linkTitle = 'Variables'
description = 'DLiteScript variable and constant declarations including syntax, scoping rules, type annotations, shadowing, and naming conventions with examples.'
weight = 0
draft = false
+++

DLiteScript supports both mutable variables and immutable constants,
both of which require explicit type declarations.

## Variable Declaration

Variables are declared using the `var` keyword followed by a name, type, and optional initial value.

### Syntax

```go
var name type = value
var name type // initialized with zero value
```

### Examples

```go
var count number = 42
var message string = "Hello"
var isActive bool = true

var uninitialized number // defaults to 0
var emptyString string // defaults to ""
var defaultBool bool // defaults to false
```

### Reassignment

Variables can be reassigned after declaration:

```go
var x number = 10
x = 20
x = x + 5
x += 5
```

## Constant Declaration

Constants are declared using the `const` keyword.
Once initialized, their values cannot be changed.

### Syntax

```go
const name type = value
```

Constants must be initialized at declaration.
You cannot declare a constant without a value.

### Examples

```go
const PI number = 3.14159
const NAME string = "DLiteScript"
const MAX_SIZE number = 100
```

### Immutability

Attempting to reassign a constant results in an error:

```go
const x number = 10
x = 20 // Error: cannot reassign constant
```

## Scoping

Both variables and constants follow block scoping rules.
A variable or constant declared within a block is only accessible within that block and any nested blocks.

### Global Scope

Declarations at the top level are accessible throughout the entire file:

```go
var globalVar string = "accessible everywhere in the file"

{
  printf("%s\n", globalVar) // works
}
```

### Block Scope

Declarations inside blocks are only accessible within that block:

```go
{
  var blockVar number = 10
  printf("%g\n", blockVar) // works
}

printf("%g\n", blockVar) // Error: blockVar not defined
```

### Nested Scopes

Inner blocks can access variables from outer blocks, but not vice versa:

```go
var outer string = "outer"

{
  var inner string = "inner"

  printf("%s\n", outer) // works
  printf("%s\n", inner) // works
}

printf("%s\n", outer) // works
printf("%s\n", inner) // Error: inner not defined
```

### Shadowing

Variables and constants in inner scopes can shadow (hide) names from outer scopes:

```go
var x number = 10

{
  var x number = 20 // different variable, shadows outer x
  printf("%g\n", x) // prints 20
}

printf("%g\n", x) // prints 10
```

## Type Annotations

All variable and constant declarations must include explicit type annotations. DLiteScript does not perform type inference.

```go
var count number = 42 // type required
var name = "Alice" // Error: missing type annotation
```

## Naming Conventions

Variable and constant names:

- Must start with a letter or underscore
- Can contain letters, numbers, and underscores
- Are case-sensitive
- Cannot be reserved keywords

```go
var myVariable number = 1 // valid
var _private string = "ok" // valid
var count123 number = 5 // valid

var 1count number = 5 // Error: cannot start with number
var for number = 1 // Error: 'for' is a keyword
```
