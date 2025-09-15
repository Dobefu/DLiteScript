package cmd

import (
	"sync"
)

var cmdMutex sync.Mutex
