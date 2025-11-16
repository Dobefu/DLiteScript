+++
title = 'arrays.filter'
linkTitle = 'filter'
description = 'Filters out falsy values from an array. Part of the arrays namespace.'
weight = 0
draft = false
+++

Filters out falsy values from an array.

## Examples

```go
arrays.filter([1, 0, 2, false, 3, "", 4]) // returns ([0, false, ""], [1, 2, 3, 4])
arrays.filter([true, false, "hello", ""]) // returns ([false, ""], [true, "hello"])
arrays.filter([1, 2, 3]) // returns ([], [1, 2, 3])
arrays.filter([]) // returns ([], [])
```
