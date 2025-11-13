package io

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getWriteFileFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "writeFile",
			Description: "Writes data to the given file.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.writeFile("secrets.txt", "this language is awesome!") // writes "this language is awesome!" to "secrets.txt" or something`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "path",
				Description: "The path for the file that should be written to.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "string",
				Description: "The data that should be written to the file.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the file does not exist or cannot be written to.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			path, _ := args[0].AsString()
			content, _ := args[1].AsString()

			exists := getExistsFunction()
			result, err := exists.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(path),
				},
			)
			if err != nil {
				return datavalue.Error(err)
			}

			if result.Bool {
				err := os.WriteFile(path, []byte(content), 0600)
				if err != nil {
					return datavalue.Error(err)
				}
			}
			return datavalue.Error(fmt.Errorf("File %v does not exist", path))
		},
	)
}
