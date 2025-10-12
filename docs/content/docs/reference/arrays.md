+++
title = 'Arrays'
linkTitle = 'Arrays'
description = 'DLiteScript arrays including creation, indexing, modification, iteration, spread operator, and built-in array functions for common operations with examples.'
weight = 0
draft = false
+++

Arrays are ordered collections of values of the same type.

## Array Declaration

Arrays are declared using square brackets `[]` followed by the element type.

### Syntax

```go
var name []type = [elements]
```

### Examples

```go
var empty []number = []
var numbers []number = [1, 2, 3, 4, 5]
var names []string = ["Alice", "Bob", "Charlie"]
var flags []bool = [true, false, true]
```

## Array Types

Array type annotations use `[]` followed by the element type.

### Common Array Types

- `[]number` - Array of numbers
- `[]string` - Array of strings
- `[]bool` - Array of booleans

## Accessing Elements

Array elements are accessed using the index operator `[]` with a zero-based index.

### Reading Elements

```go
var fruits []string = ["apple", "banana", "cherry"]

printf("%s\n", fruits[0]) // "apple"
printf("%s\n", fruits[1]) // "banana"
printf("%s\n", fruits[2]) // "cherry"
```

### Index Out of Bounds

Accessing an index that doesn't exist results in an error:

```go
var numbers []number = [1, 2, 3]

printf("%g\n", numbers[10]) // Error: array index out of bounds
```

## Modifying Elements

You can modify array elements by assigning to a specific index.

```go
var numbers []number = [1, 2, 3]

numbers[1] = 10
printf("%g\n", numbers[1]) // 10

printf("%s\n", numbers) // [1, 10, 3]
```

## Array Operations

### Concatenation

The `+` operator can be used to concatenate arrays:

```go
var arr1 []number = [1, 2, 3]
var arr2 []number = [4, 5, 6]
var combined []number = arr1 + arr2

printf("%s\n", combined) // [1, 2, 3, 4, 5, 6]
```

### Reassignment

Arrays can be reassigned to new values:

```go
var array1 []number = []

printf("Empty array: %s\n", array1) // []

array1 = [4, 5, 6]
printf("Array with [4, 5, 6]: %s\n", array1) // [4, 5, 6]
```

## Using Arrays with Functions

### Passing Arrays as Arguments

Arrays can be passed to functions:

```go
func printArray(arr []number) {
  printf("%s\n", arr)
}

var numbers []number = [1, 2, 3]
printArray(numbers)
```

### Returning Arrays

Functions can return arrays:

```go
func getRange() []number {
  return [1, 2, 3, 4, 5]
}

var numbers []number = getRange()
```

### Spread Operator with Arrays

Use the spread operator (`...`) to expand array elements as individual function arguments:

```go
func sum(a number, b number, c number) number {
  return a + b + c
}

var numbers []number = [1, 2, 3]
var total number = sum(...numbers) // expands to sum(1, 2, 3)

printf("%g\n", total) // 6
```

## Iterating Over Arrays

Use a `for` loop to iterate over array indices:

```go
var fruits []string = ["apple", "banana", "cherry"]

for var i to 3 {
  printf("%s\n", fruits[i])
}
```

## Practical Examples

### Building an Array

```go
var numbers []number = [1, 2, 3]

printf("Initial: %s\n", numbers)

numbers = numbers + [4, 5]
printf("After adding: %s\n", numbers) // [1, 2, 3, 4, 5]
```

### Array Modification

```go
var array2 []number = [1, 2, 3, 4, 5, 6]

printf("array2[1]: %g\n", array2[1]) // 2

array2[1] = 10

printf("array2[1]: %g\n", array2[1]) // 10
printf("array2: %s\n", array2) // [1, 10, 3, 4, 5, 6]
```

## Type Safety

All elements in an array must be of the same type.
Type mismatches will result in errors.

```go
var numbers []number = [1, 2, 3]

// This would cause an error:
// numbers = ["a", "b", "c"] // Error: type mismatch
```
