+++
title = 'strings.find'
linkTitle = 'find'
description = 'Finds the first occurrence of a substring.'
weight = 0
draft = false
+++

Finds the first occurrence of a substring (alias for `indexOf`).

## Examples

```go
printf("%g\n", strings.find("hello world", "world")) // 6
printf("%g\n", strings.find("hello world", "xyz"))   // -1
```

