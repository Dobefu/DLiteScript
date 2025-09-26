// Package time provides time functions for the standard library.
package time

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "time"

// GetTimeFunctions returns the time functions for the standard library.
func GetTimeFunctions() map[string]function.Info {
	return map[string]function.Info{
		"now": getNowFunction(),
	}
}
