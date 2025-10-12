+++
title = 'strings.find'
linkTitle = 'find'
description = 'Find the first occurrence of a substring. Search for text patterns in strings and return the matching substring or empty string. Part of the strings namespace.'
weight = 0
draft = false
+++

Finds the first occurrence of a substring (alias for `indexOf`).

## Examples

```go
printf("%g\n", strings.find("hello world", "world")) // 6
printf("%g\n", strings.find("hello world", "xyz"))   // -1
```

