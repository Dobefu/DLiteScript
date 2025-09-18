// Package stdlib provides the standard library for DLiteScript.
package stdlib

import (
	"github.com/Dobefu/DLiteScript/internal/function"
	"github.com/Dobefu/DLiteScript/internal/stdlib/global"

	stdlibarrays "github.com/Dobefu/DLiteScript/internal/stdlib/arrays"
	stdliberrors "github.com/Dobefu/DLiteScript/internal/stdlib/errors"
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
	functionRegistry["arrays"] = function.PackageInfo{
		Functions: stdlibarrays.GetArrayFunctions(),
	}
	functionRegistry["errors"] = function.PackageInfo{
		Functions: stdliberrors.GetErrorFunctions(),
	}
}

// GetFunctionRegistry gets the function registry.
func GetFunctionRegistry() map[string]function.PackageInfo {
	return functionRegistry
}
