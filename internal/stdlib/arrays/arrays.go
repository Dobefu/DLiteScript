// Package arrays provides the array functions for the standard library.
package arrays

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "arrays"

// GetArrayFunctions returns the array functions for the standard library.
func GetArrayFunctions() map[string]function.Info {
	return map[string]function.Info{
		"add":     getAddFunction(),
		"length":  getLengthFunction(),
		"reverse": getReverseFunction(),
		"join":    getJoinFunction(),
	}
}
