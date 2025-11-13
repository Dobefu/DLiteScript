// Package stdlib provides the standard library for DLiteScript.
package stdlib

import (
	"github.com/Dobefu/DLiteScript/internal/function"
	"github.com/Dobefu/DLiteScript/internal/stdlib/global"

	stdlibarrays "github.com/Dobefu/DLiteScript/internal/stdlib/arrays"
	stdliberrors "github.com/Dobefu/DLiteScript/internal/stdlib/errors"
	stdlibio "github.com/Dobefu/DLiteScript/internal/stdlib/io"
	stdlibmath "github.com/Dobefu/DLiteScript/internal/stdlib/math"
	stdlibos "github.com/Dobefu/DLiteScript/internal/stdlib/os"
	stdlibstrings "github.com/Dobefu/DLiteScript/internal/stdlib/strings"
	stdlibtime "github.com/Dobefu/DLiteScript/internal/stdlib/time"
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
	functionRegistry["strings"] = function.PackageInfo{
		Functions: stdlibstrings.GetStringsFunctions(),
	}
	functionRegistry["time"] = function.PackageInfo{
		Functions: stdlibtime.GetTimeFunctions(),
	}
	functionRegistry["os"] = function.PackageInfo{
		Functions: stdlibos.GetOSFunctions(),
	}
	functionRegistry["io"] = function.PackageInfo{
		Functions: stdlibio.GetIOFunctions(),
	}
}

// GetFunctionRegistry gets the function registry.
func GetFunctionRegistry() map[string]function.PackageInfo {
	return functionRegistry
}
