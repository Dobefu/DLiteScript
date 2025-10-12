+++
title = 'Functions'
linkTitle = 'Functions'
description = 'Function declarations and calls in DLiteScript.'
weight = 0
draft = false
+++

Functions are reusable blocks of code that can accept parameters and return values.

## Function Declaration

Functions are declared using the `func` keyword followed by a name, parameters, return types, and body.

### Basic Syntax

```go
func name() {
  // function body
}
```

### Example

```go
func greet() {
  printf("Hello, world!\n")
}

greet() // call the function
```

## Parameters

Functions can accept zero or more parameters, each with a name and type.

### Single Parameter

```go
func greet(name string) {
  printf("Hello, %s!\n", name)
}

greet("John")
```

### Multiple Parameters

```go
func add(a number, b number) {
  printf("Sum: %g\n", a + b)
}

add(5, 3)
```

## Return Values

Functions can return values using the `return` statement.

### Single Return Value

```go
func double(x number) number {
  return x * 2
}

var result number = double(5) // result is 10
```

### Multiple Return Values

Functions can return multiple values.

```go
func getCoordinates() number, string {
  return 42, "hello"
}

var x number
var label string
x, label = getCoordinates()
```

### Parenthesized Return Types

Multiple return types can optionally be wrapped in parentheses for clarity.

```go
func getCoordinates() (number, string) {
  return 42, "hello"
}
```

Both styles are equivalent.

## Return Statement

The `return` statement exits the function and optionally returns values.

### Returning Values

```go
func getMax(a number, b number) number {
  if a > b {
    return a
  }

  return b
}
```

### Early Return

You can use `return` to exit a function early:

```go
func checkValue(x number) {
  if x < 0 {
    printf("Negative value\n")

    return
  }

  printf("Positive value: %g\n", x)
}
```

### Returning Null

Functions can return `null` as a value:

```go
func maybeGetValue(flag bool) number {
  if flag {
    return 42
  }

  return null
}
```

## Function Calls

Functions are called by their name followed by arguments in parentheses.

### Basic Call

```go
func sayHello() {
  printf("Hello!\n")
}

sayHello()
```

### With Arguments

```go
func multiply(a number, b number) number {
  return a * b
}

var result number = multiply(4, 5) // 20
```

### Using Spread Operator

When a function returns multiple values, use the spread operator to pass them as arguments to another function:

```go
func getTwo() number, number {
  return 10, 20
}

func addTwo(a number, b number) number {
  return a + b
}

var sum number = addTwo(...getTwo()) // sum is 30
```

You can also use it to pass array elements as arguments:

```go
func printThree(a number, b number, c number) {
  printf("%g, %g, %g\n", a, b, c)
}

var numbers []number = [1, 2, 3]
printThree(...numbers) // prints "1, 2, 3"
```

## Function Examples

### Simple Function

```go
func test() number {
  return 1
}

printf("test(): %g\n", test())
```

### Function with Parameter

```go
func testWithNumber(count number) number {
  return count
}

printf("testWithNumber(): %g\n", testWithNumber(2))
```

### Multiple Return Values

```go
func testMultiple() number, string {
  return 1, "hello"
}

printf("testMultiple(): %g, %s\n", ...testMultiple())
```

### Complex Function

```go
func testComplex(count number, name string) (number, string) {
  return count, null
}

printf("testComplex(): %g, %s\n", ...testComplex(1, "hello"))
```

## Scope

Functions have their own scope. Variables declared inside a function are local to that function.

```go
var global string = "global"

func myFunction() {
  var local string = "local"
  printf("%s\n", global) // can access global
  printf("%s\n", local)  // can access local
}

myFunction()
printf("%s\n", global) // works
printf("%s\n", local)  // Error: local not defined
```

## Recursion

Functions can call themselves recursively.

```go
func factorial(n number) number {
  if n <= 1 {
    return 1
  }

  return n * factorial(n - 1)
}

printf("5! = %g\n", factorial(5)) // 120
```
