+++
title = 'os.getEnvVariable'
linkTitle = 'getEnvVariable'
description = 'Retrieve the value of an environment variable. Read system environment variables for configuration and runtime settings. Part of the os namespace.'
weight = 0
draft = false
+++

Retrieves the value of an environment variable.

## Examples

```go
var path string = os.getEnvVariable("PATH")
printf("PATH: %s\n", path)
```

