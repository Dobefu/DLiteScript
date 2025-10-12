+++
title = 'arrays.slice'
linkTitle = 'slice'
description = 'Extract a part of an array by start and end index. Return a new array with elements from the range without changing the original. Part of the arrays namespace.'
weight = 0
draft = false
+++

Extracts a portion of an array.

## Examples

```go
var arr []number = [1, 2, 3, 4, 5]
var sliced []number = arrays.slice(arr, 1, 3)
printf("%s\n", sliced) // [2, 3]
```

