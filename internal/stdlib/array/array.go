// Package array provides the array functions for the standard library.
package array

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "array"

// GetArrayFunctions returns the array functions for the standard library.
func GetArrayFunctions() map[string]function.Info {
	return map[string]function.Info{
		"add": getAddFunction(),
	}
}
