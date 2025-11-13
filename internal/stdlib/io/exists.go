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
			Description: "Returns whether the given file/dir exists or not.",
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
				Description: "The path for the file/dir that should be checked whether it exists or not.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeBool,
				Name:        "state",
				Description: "The bool for whether the file/dir exists or not",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the state of the file/dir cannot be determined.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			path, _ := args[0].AsString()

			_, err := os.Stat(path)

			if err == nil {
				return datavalue.Tuple(datavalue.Bool(true), datavalue.Null())
			}

			if errors.Is(err, os.ErrNotExist) {
				return datavalue.Tuple(datavalue.Bool(false), datavalue.Null())
			}

			return datavalue.Tuple(datavalue.Bool(false), datavalue.Error(err))
		},
	)
}
