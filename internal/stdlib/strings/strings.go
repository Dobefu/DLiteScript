// Package strings provides the string functions for the standard library.
package strings

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "strings"

// GetStringsFunctions returns the string functions for the standard library.
func GetStringsFunctions() map[string]function.Info {
	return map[string]function.Info{
		"length":     getLengthFunction(),
		"find":       getFindFunction(),
		"has":        getHasFunction(),
		"replace":    getReplaceFunction(),
		"replaceAll": getReplaceAllFunction(),
		"split":      getSplitFunction(),
		"trim":       getTrimFunction(),
		"trimLeft":   getTrimLeftFunction(),
		"trimRight":  getTrimRightFunction(),
		"toUpper":    getToUpperFunction(),
		"toLower":    getToLowerFunction(),
		"substring":  getSubstringFunction(),
	}
}
