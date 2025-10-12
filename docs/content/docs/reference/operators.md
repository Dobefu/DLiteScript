+++
title = 'Operators'
linkTitle = 'Operators'
description = 'DLiteScript operators including arithmetic, comparison, logical, assignment, and special operators like spread and index operators with examples.'
weight = 0
draft = false
+++

DLiteScript provides a variety of operators for performing operations on values.

## Arithmetic Operators

Arithmetic operators perform mathematical calculations on numeric values.

| Operator | Description    | Example  | Result |
| -------- | -------------- | -------- | ------ |
| `+`      | Addition       | `5 + 3`  | `8`    |
| `-`      | Subtraction    | `5 - 3`  | `2`    |
| `*`      | Multiplication | `5 * 3`  | `15`   |
| `/`      | Division       | `10 / 2` | `5`    |
| `%`      | Modulo         | `10 % 3` | `1`    |
| `**`     | Exponentiation | `2 ** 3` | `8`    |

### Examples

```go
var sum number = 10 + 5        // 15
var difference number = 10 - 5 // 5
var product number = 10 * 5    // 50
var quotient number = 10 / 5   // 2
var remainder number = 10 % 3  // 1
var power number = 2 ** 3      // 8
```

### String Concatenation

The `+` operator also works for string concatenation:

```go
var greeting string = "Hello, " + "world!" // "Hello, world!"
```

### Array Concatenation

The `+` operator concatenates arrays as well:

```go
var arr1 []number = [1, 2, 3]
var arr2 []number = [4, 5, 6]
var combined []number = arr1 + arr2 // [1, 2, 3, 4, 5, 6]
```

## Assignment Operators

Assignment operators assign values to variables.

### Basic Assignment

| Operator | Description | Example |
| -------- | ----------- | ------- |
| `=`      | Assignment  | `x = 5` |

### Compound Assignment

Compound assignment operators perform an operation and assign the result in one step.

| Operator | Description             | Example   | Equivalent to |
| -------- | ----------------------- | --------- | ------------- |
| `+=`     | Add and assign          | `x += 5`  | `x = x + 5`   |
| `-=`     | Subtract and assign     | `x -= 5`  | `x = x - 5`   |
| `*=`     | Multiply and assign     | `x *= 5`  | `x = x * 5`   |
| `/=`     | Divide and assign       | `x /= 5`  | `x = x / 5`   |
| `%=`     | Modulo and assign       | `x %= 5`  | `x = x % 5`   |
| `**=`    | Exponentiate and assign | `x **= 2` | `x = x ** 2`  |

### Examples

```go
var x number = 10

x += 5  // x went from 10 to 15
x -= 3  // x went from 15 to 12
x *= 2  // x went from 12 to 24
x /= 4  // x went from 24 to 6
x %= 4  // x went from 6 to 2
x **= 3 // x went from 2 to 8
```

## Comparison Operators

Comparison operators compare two values and return a boolean result.

| Operator | Description              | Example  | Result |
| -------- | ------------------------ | -------- | ------ |
| `==`     | Equal to                 | `5 == 5` | `true` |
| `!=`     | Not equal to             | `5 != 3` | `true` |
| `>`      | Greater than             | `5 > 3`  | `true` |
| `>=`     | Greater than or equal to | `5 >= 5` | `true` |
| `<`      | Less than                | `3 < 5`  | `true` |
| `<=`     | Less than or equal to    | `3 <= 5` | `true` |

### Type Comparison

Comparisons between different types are allowed:

```go
printf("%t\n", 1 == 1)           // true
printf("%t\n", "test" == "test") // true
printf("%t\n", null == null)     // true
printf("%t\n", 1 != 2)           // true
printf("%t\n", 1 != "test")      // true (different types)
printf("%t\n", true != "test")   // true (different types)
```

## Logical Operators

Logical operators perform boolean logic operations.

| Operator | Description | Example           | Result |
| -------- | ----------- | ----------------- | ------ |
| `&&`     | Logical AND | `true && true`    | `true` |
| `\|\|`   | Logical OR  | `true \|\| false` | `true` |
| `!`      | Logical NOT | `!false`          | `true` |

### Examples

```go
var a bool = true
var b bool = false

printf("%t\n", a && a)  // true
printf("%t\n", a && b)  // false
printf("%t\n", a || b)  // true
printf("%t\n", b || b)  // false
printf("%t\n", !a)      // false
printf("%t\n", !b)      // true
```

### Short-Circuit Evaluation

Logical operators use short-circuit evaluation:

- `&&` returns `false` without evaluating the right operand if the left is `false`
- `||` returns `true` without evaluating the right operand if the left is `true`

## Special Operators

### Spread Operator (`...`)

The spread operator expands array elements, typically used in function calls to pass array elements as individual arguments.

```go
func printThree(a number, b number, c number) {
  printf("%g, %g, %g\n", a, b, c)
}

var numbers []number = [1, 2, 3]
printThree(...numbers) // expands to printThree(1, 2, 3)
```

### Index Operator (`[]`)

The index operator accesses elements in an array by their position (zero-based).

```go
var fruits []string = ["apple", "banana", "cherry"]

printf("%s\n", fruits[0]) // "apple"
printf("%s\n", fruits[1]) // "banana"
printf("%s\n", fruits[2]) // "cherry"
```

You can also use it to assign values:

```go
var numbers []number = [1, 2, 3]
numbers[1] = 10
printf("%g\n", numbers[1]) // 10
```

## Operator Precedence

Operators are evaluated in the following order (highest to lowest precedence):

1. `**` (Exponentiation)
2. `!` (Logical NOT)
3. `*`, `/`, `%` (Multiplication, Division, Modulo)
4. `+`, `-` (Addition, Subtraction)
5. `<`, `<=`, `>`, `>=` (Comparison)
6. `==`, `!=` (Equality)
7. `&&` (Logical AND)
8. `||` (Logical OR)

Use parentheses to explicitly control evaluation order:

```go
var result1 number = 2 + 3 * 4     // 14 (multiplication first)
var result2 number = (2 + 3) * 4   // 20 (parentheses first)
```
