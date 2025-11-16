package io

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getReadFileStringFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "readFileString",
			Description: "Reads the content of a file and returns it as a string. Returns an error if the file does not exist.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.readFileString("data.txt") // returns the raw content of the given file as a string or an error if the file does not exist`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "path",
				Description: "The path for the file that should be read.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "value",
				Description: "The content of the given file as a string, or an error if the file does not exist.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the content of the given file cannot be read or the file does not exist.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			path, _ := args[0].AsString()

			//#nosec G304
			value, err := os.ReadFile(path)
			if err != nil {
				return datavalue.Error(err)
			}

			return datavalue.String(string(value))
		},
	)
}
