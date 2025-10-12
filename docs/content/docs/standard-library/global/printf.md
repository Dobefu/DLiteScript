+++
title = 'printf'
linkTitle = 'printf'
description = 'Print formatted strings to standard output. Output text with format specifiers for numbers, strings, and booleans to the console. Part of the global namespace.'
weight = 0
draft = false
+++

Prints a formatted string to standard output.

## Examples

```go
printf("Hello, %s!\n", "world") // Hello, world!
printf("Number: %g\n", 42)      // Number: 42
printf("Boolean: %t\n", true)   // Boolean: true
```

## Format Specifiers

- `%s` - String
- `%g` - Number
- `%t` - Boolean
