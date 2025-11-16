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
			Description: "Writes data to the given file. Returns an error if the file does not exist.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.writeFile("data.txt", "This is a new line.") // writes "This is a new line." to "data.txt"`, packageName),
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
				Description: "An error, if the file cannot be written to or does not exist.",
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

			resultTuple, _ := result.AsTuple()
			fileExists, err := resultTuple[0].AsBool()

			if err != nil {
				return datavalue.Error(err)
			}

			if !fileExists {
				return datavalue.Error(fmt.Errorf("file %v does not exist", path))
			}

			err = os.WriteFile(path, []byte(content), 0600)

			if err != nil {
				return datavalue.Error(err)
			}

			return datavalue.Error(nil)
		},
	)
}
