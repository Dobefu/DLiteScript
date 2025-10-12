+++
title = 'sprintf'
linkTitle = 'sprintf'
description = 'Return a formatted string. Format strings with placeholders for numbers, strings, and booleans for variable assignment. Part of the global namespace.'
weight = 0
draft = false
+++

Returns a formatted string without printing it.

## Examples

```go
var message string = sprintf("Sum: %g", 10 + 5)
printf("%s\n", message) // Sum: 15
```

## Format Specifiers

- `%s` - String
- `%g` - Number
- `%t` - Boolean

