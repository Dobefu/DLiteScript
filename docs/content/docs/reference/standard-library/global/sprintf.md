+++
title = 'sprintf'
linkTitle = 'sprintf'
description = 'Returns a formatted string without printing it.'
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

