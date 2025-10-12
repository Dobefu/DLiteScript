+++
title = 'dump'
linkTitle = 'dump'
description = 'Outputs the value and type of arguments for debugging.'
weight = 0
draft = false
+++

Outputs the value and type of one or more arguments for debugging.

## Examples

```go
dump("test")    // Outputs: string("test")
dump(42)        // Outputs: number(42)
dump(true)      // Outputs: bool(true)
dump([1, 2, 3]) // Outputs: array([1, 2, 3])
```
