package io

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getCreateFileFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "createFile",
			Description: "Creates an empty file.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.createFile("awesome.txt") // creates an empty file called "awesome.txt" or something`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "path",
				Description: "The path for the file that should be created.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the file cannot be created.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			path, _ := args[0].AsString()

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

			if !result.Bool {
				err := os.WriteFile(path, []byte(""), 0600)
				if err != nil {
					return datavalue.Error(err)
				}

				return datavalue.Error(nil)
			}

			return datavalue.Error(fmt.Errorf("file %v already exists", path))
		},
	)
}
