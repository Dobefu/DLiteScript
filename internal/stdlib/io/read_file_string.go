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
			Description: "Reads the content of a file and returns it as a string.",
			Since:       "v0.1.2",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.readFileString("names.txt") // returns the raw content of the given file as a string"`, packageName),
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
				Description: "The content of the given file, or empty string if not found.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the content of the given file cannot be read or returned as a string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			path, _ := args[0].AsString()

			value, err := os.ReadFile(path)
			if err != nil {
				return datavalue.Error(err)
			}

			return datavalue.String(string(value))
		},
	)
}
