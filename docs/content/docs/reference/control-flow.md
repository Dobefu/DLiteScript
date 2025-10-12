+++
title = 'Control Flow'
linkTitle = 'Control Flow'
description = 'Control flow statements and loops in DLiteScript.'
weight = 0
draft = false
+++

Control flow statements allow you to control the execution path of your program based on one or more conditions.

## If Statements

The `if` statement executes code conditionally based on a boolean expression.

### Basic If

```go
if condition {
  // code to execute if condition is true
}
```

### Example

```go
var coverage number = 85

if coverage > 80 {
  printf("That's some pretty good coverage\n")
}
```

### If-Else

The `else` clause provides an alternative code path when the condition is false.

```go
if condition {
  // code if true
} else {
  // code if false
}
```

### Example

```go
var coverage number = 40

if coverage >= 80 {
  printf("Coverage is at or above 80%\n")
} else {
  printf("Coverage is below 80%\n")
}
```

### Else-If Chains

Multiple conditions can be checked using `else if`.

```go
if condition1 {
  // code if condition1 is true
} else if condition2 {
  // code if condition2 is true
} else {
  // code if all conditions are false
}
```

### Example

```go
var score number = 75

if score >= 90 {
  printf("Grade: A\n")
} else if score >= 80 {
  printf("Grade: B\n")
} else if score >= 70 {
  printf("Grade: C\n")
} else {
  printf("Grade: F\n")
}
```

### Optional Parentheses

Parentheses around conditions are optional but can be used for clarity:

```go
var test bool = true

if test {
  printf("without parentheses\n")
}

if (test) {
  printf("with parentheses\n")
}
```

## For Loops

The `for` statement provides various ways to create loops.

### Infinite Loop

A `for` loop without any condition runs indefinitely (use `break` to exit).

```go
for {
  printf("This loops forever\n")
  break // exit the loop
}
```

### Condition Loop

A `for` loop with a condition runs while the condition is true.

```go
var count number = 0

for count < 5 {
  printf("Count: %g\n", count)
  count += 1
}
```

### Range Loop (to)

Loop from 0 up to (but not including) a value.

```go
for var i to 3 {
  printf("Iteration %g\n", i)
}
// Prints: 0, 1, 2
```

### Range Loop (from-to)

Loop from a starting value up to (but not including) an ending value.

```go
for var i from 5 to 8 {
  printf("Iteration %g\n", i)
}
// Prints: 5, 6, 7
```

### Comparison Loop

Loop while a variable meets a comparison condition.

```go
for var i < 3 {
  printf("Iteration %g\n", i)
}
// Prints: 0, 1, 2
```

## Loop Control

### Break Statement

The `break` statement exits the current loop immediately.

```go
for var i to 10 {
  if i == 5 {
    break
  }
  printf("%g\n", i)
}
// Prints: 0, 1, 2, 3, 4
```

### Break with Depth

You can break out of multiple nested loops by specifying a depth.

```go
for var i to 3 {
  for var j to 3 {
    if i == 1 && j == 1 {
      break 2 // breaks out of both loops
    }
    printf("i=%g, j=%g\n", i, j)
  }
}
```

### Continue Statement

The `continue` statement skips the rest of the current iteration and moves to the next one.

```go
for var i to 5 {
  if i == 2 {
    continue
  }
  printf("%g\n", i)
}
// Prints: 0, 1, 3, 4 (skips 2)
```

## Nested Loops

Loops can be nested inside other loops.

```go
for var i to 2 {
  for var j to 2 {
    printf("i=%g, j=%g\n", i, j)
  }
}
// Output:
// i=0, j=0
// i=0, j=1
// i=1, j=0
// i=1, j=1
```

## Practical Examples

### Conditional Logic with Multiple Conditions

```go
var test1 bool = false
var test2 bool = false
var test3 bool = true

if test1 {
  printf("test1 is true\n")
} else if test2 {
  printf("test2 is true\n")
} else if test3 {
  printf("test3 is true\n")
} else {
  printf("none of the conditions are true\n")
}
// Output: test3 is true
```

### Loop with Conditional Break

```go
var isComplete bool = false
var counter number = 0

for !isComplete {
  printf("Iteration %g\n", counter)
  counter += 1

  if counter >= 3 {
    isComplete = true
  }
}
```
