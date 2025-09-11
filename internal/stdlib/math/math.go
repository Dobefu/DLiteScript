// Package math provides maths functions for the standard library.
package math

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "math"

// GetMathFunctions returns the math functions for the standard library.
func GetMathFunctions() map[string]function.Info {
	return map[string]function.Info{
		"abs":   getAbsFunction(),
		"sin":   getSinFunction(),
		"cos":   getCosFunction(),
		"tan":   getTanFunction(),
		"sqrt":  getSqrtFunction(),
		"round": getRoundFunction(),
		"floor": getFloorFunction(),
		"ceil":  getCeilFunction(),
		"min":   getMinFunction(),
		"max":   getMaxFunction(),
	}
}
