+++
title = 'Syntax'
linkTitle = 'Syntax'
description = 'Lexical structure and syntax of DLiteScript.'
weight = -1
draft = false
+++

## Identifiers

Identifiers are names used for variables, constants, functions, and others.
They must start with a letter or underscore,
followed by any combination of letters, numbers, or underscores.

### Examples:

- `myVariable`
- `_privateVar`
- `counter123`

## Keywords

The following words are reserved and cannot be used as identifiers:

### General keywords

| Keyword    | Description            |
| ---------- | ---------------------- |
| `var`      | Variable declaration   |
| `const`    | Constant declaration   |
| `func`     | Function declaration   |
| `if`       | Conditional statement  |
| `else`     | Alternative branch     |
| `for`      | Loop statement         |
| `from`     | Loop range start       |
| `to`       | Loop range end         |
| `break`    | Exit loop              |
| `continue` | Skip to next iteration |
| `return`   | Return from function   |
| `import`   | Import module          |
| `as`       | Import alias           |
| `null`     | Null value             |
| `true`     | Boolean true           |
| `false`    | Boolean false          |

### Type Keywords

| Type     | Description  |
| -------- | ------------ |
| `number` | Numeric type |
| `string` | String type  |
| `bool`   | Boolean type |
| `any`    | Any type     |
| `error`  | Error type   |

## Literals

### Number Literals

Numeric values support integers and floating-point numbers.

#### Examples:

- `42`
- `3.14`
- `0.001`
- `-17`

### String Literals

Strings are enclosed in double quotes and support escape sequences.

#### Examples:

- `"Hello, world!"`
- `"Line 1\nLine 2"`
- `"Tab\tseparated"`

### Boolean Literals

Boolean values can be either `true` or `false`.

#### Examples:

- `true`
- `false`

### Null Literal

The `null` keyword represents the absence of a value.

### Array Literals

Arrays are declared with square brackets.

#### Examples:

- `[]`
- `[1, 2, 3]`
- `["hello", "world"]`

## Comments

Single-line comments start with `//`.

- `// This is a comment`
- `printf("Hello\n") // Inline comment`

## Operators

### Arithmetic Operators

| Operator | Description    |
| -------- | -------------- |
| `+`      | Addition       |
| `-`      | Subtraction    |
| `*`      | Multiplication |
| `/`      | Division       |
| `%`      | Modulo         |
| `**`     | Exponentiation |

### Assignment Operators

| Operator | Description             |
| -------- | ----------------------- |
| `=`      | Assignment              |
| `+=`     | Add and assign          |
| `-=`     | Subtract and assign     |
| `*=`     | Multiply and assign     |
| `/=`     | Divide and assign       |
| `%=`     | Modulo and assign       |
| `**=`    | Exponentiate and assign |

### Comparison Operators

| Operator | Description              |
| -------- | ------------------------ |
| `==`     | Equal to                 |
| `!=`     | Not equal to             |
| `>`      | Greater than             |
| `>=`     | Greater than or equal to |
| `<`      | Less than                |
| `<=`     | Less than or equal to    |

### Logical Operators

| Operator | Description |
| -------- | ----------- |
| `&&`     | Logical AND |
| `\|\|`   | Logical OR  |
| `!`      | Logical NOT |

### Special Operators

| Operator | Description     |
| -------- | --------------- |
| `...`    | Spread operator |
| `[]`     | Index operator  |

## Statements

### Expression Statements

Any expression can be a statement:

```go
printf("Hello\n")
x + 1
```

### Block Statements

Statements can be grouped and scoped with braces:

```go
{
  var x number = 1
  printf("%g\n", x)
}
```

## Code Structure

DLiteScript programs are sequences of statements. Statements are typically separated by newlines, though multiple statements can appear on the same line. Braces `{}` create new scopes for variables and constants.

```go
var x number = 10
printf("%g\n", x)

{
  var y number = 20
  printf("%g\n", y)
}
```
