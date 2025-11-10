package io

import (
	"errors"
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getExistsFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "exists",
			Description: "Returns whether the given file exists or not.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.exists("crazy.txt"") // returns a bool based on whether "crazy.txt" exists or not`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "path",
				Description: "The path for the file that should be checked whether it exists or not.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeBool,
				Name:        "state",
				Description: "The bool for whether the file exists or not",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the state of the file cannot be determined.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			path, _ := args[0].AsString()

			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				return datavalue.Bool(false)
			}
			return datavalue.Bool(true)
		},
	)
}
