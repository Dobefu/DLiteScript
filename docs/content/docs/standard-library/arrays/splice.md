+++
title = 'arrays.splice'
linkTitle = 'splice'
description = 'Remove or replace elements in an array. Modify array by removing elements at index and inserting new values at that position. Part of the arrays namespace.'
weight = 0
draft = false
+++

Removes or replaces elements in an array.

## Examples

```go
var arr []number = [1, 2, 3, 4, 5]
arr = arrays.splice(arr, 2, 2) // Remove 2 elements starting at index 2
printf("%s\n", arr) // [1, 2, 5]
```

