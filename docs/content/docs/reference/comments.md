+++
title = 'Comments'
linkTitle = 'Comments'
description = 'Comment syntax in DLiteScript.'
weight = 0
draft = false
+++

Comments allow you to add notes and documentation to your code.
They are ignored by the interpreter and have no effect on program execution.

## Single-Line Comments

DLiteScript supports single-line comments using `//`.
Everything after `//` on that line is treated as a comment.

### Syntax

```go
// This is a comment.
```

### Examples

#### Full-Line Comments

```go
// This is a full-line comment
printf("Hello, world!\n")
```

#### Inline Comments

Comments can appear at the end of a line of code:

```go
var x number = 42 // This is an inline comment
printf("Value: %g\n", x) // Print the value
```

#### Multiple Comments

You can have multiple comment lines:

```go
// This is the first comment.
// This is the second comment.
// This is the third comment.
var y number = 10
```

#### Commented Code

Comments are often used to temporarily disable code:

```go
// printf("This line won't execute\n")
printf("This line will execute\n")
```

## Comments in Strings

The `//` sequence inside a string literal is not treated as a comment:

```go
printf("// This is not a comment\n") // This IS a comment
```

Output: `// This is not a comment`
