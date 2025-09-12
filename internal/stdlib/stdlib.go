// Package stdlib provides the standard library for DLiteScript.
package stdlib

import (
	"github.com/Dobefu/DLiteScript/internal/function"
	"github.com/Dobefu/DLiteScript/internal/stdlib/global"

	stdlibarray "github.com/Dobefu/DLiteScript/internal/stdlib/array"
	stdlibmath "github.com/Dobefu/DLiteScript/internal/stdlib/math"
)

var functionRegistry = map[string]function.PackageInfo{
	"": {
		Functions: map[string]function.Info{},
	},
}

func init() {
	functionRegistry[""] = function.PackageInfo{
		Functions: global.GetGlobalFunctions(),
	}
	functionRegistry["math"] = function.PackageInfo{
		Functions: stdlibmath.GetMathFunctions(),
	}
	functionRegistry["array"] = function.PackageInfo{
		Functions: stdlibarray.GetArrayFunctions(),
	}
}

// GetFunctionRegistry gets the function registry.
func GetFunctionRegistry() map[string]function.PackageInfo {
	return functionRegistry
}
