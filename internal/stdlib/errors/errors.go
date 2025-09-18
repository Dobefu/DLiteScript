// Package errors provides the errors functions for the standard library.
package errors

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "errors"

// GetErrorFunctions returns the error functions for the standard library.
func GetErrorFunctions() map[string]function.Info {
	return map[string]function.Info{
		"new": getNewFunction(),
	}
}
