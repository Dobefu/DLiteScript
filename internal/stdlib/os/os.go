// Package os provides os functions for the standard library.
package os

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "os"

// GetOSFunctions returns the os functions for the standard library.
func GetOSFunctions() map[string]function.Info {
	return map[string]function.Info{
		"getEnvVariable": getGetEnvVariableFunction(),
		"setEnvVariable": getSetEnvVariableFunction(),
	}
}
