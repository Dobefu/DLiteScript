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
			Description: "Appends data to the given file.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.appendFile("secrets.txt", "this language is awesome!") // appends "this language is awesome! to "secrets.txt" or something without replacing the original content.`, packageName),
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

			if result.Bool {
				oldContent, err := os.ReadFile(path)
				if err != nil {
					return datavalue.Error(err)
				}
				content = string(oldContent) + content
				err = os.WriteFile(path, []byte(content), 0644)
				if err != nil {
					return datavalue.Error(err)
				}
			}
			return datavalue.Error(fmt.Errorf("File %v does not exist", path))
		},
	)
}
