// Package io provides io functions for the standard library.
package io

import (
	"github.com/Dobefu/DLiteScript/internal/function"
)

const packageName = "io"

// GetIOFunctions returns the io functions for the standard library.
func GetIOFunctions() map[string]function.Info {
	return map[string]function.Info{
		"readFile":   readFileFunction(),
		"writeFile":  writeFileFunction(),
		"exists":     existsFunction(),
		"deleteFile": deleteFileFunction(),
		"deleteDir":  deleteDirFunction(),
	}
}
