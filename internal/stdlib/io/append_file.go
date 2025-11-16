package io

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getAppendFileFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "appendFile",
			Description: "Appends data to a given file without replacing the original content.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.appendFile("data.txt", "This is a new line.") // Appends "This is a new line." to "data.txt"`, packageName),
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
				Description: "The data that should be appended to the file.",
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

			resultTuple, _ := result.AsTuple()
			fileExists, err := resultTuple[0].AsBool()

			if err != nil {
				return datavalue.Error(err)
			}

			if !fileExists {
				return datavalue.Error(fmt.Errorf("file %v does not exist", path))
			}

			oldContent, err := os.ReadFile(path) // #nosec G304

			if err != nil {
				return datavalue.Error(err)
			}

			content = string(oldContent) + content
			err = os.WriteFile(path, []byte(content), 0600)

			if err != nil {
				return datavalue.Error(err)
			}

			return datavalue.Error(nil)
		},
	)
}
