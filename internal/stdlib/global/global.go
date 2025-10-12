// Package global provides the global functions for the standard library.
package global

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = ""

// GetGlobalFunctions returns the global functions for the standard library.
func GetGlobalFunctions() map[string]function.Info {
	return map[string]function.Info{
		"printf":  getPrintfFunction(),
		"sprintf": getSprintfFunction(),
		"dump":    getDumpFunction(),
		"exit":    getExitFunction(),
	}
}
