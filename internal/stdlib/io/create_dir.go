package io

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getCreateDirFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "createDir",
			Description: "Creates an empty directory.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.createDir("memes") // creates an empty directory called "memes" or something`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "path",
				Description: "The path for the directory that should be created.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the directory cannot be created.",
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

			resultTuple, _ := result.AsTuple()
			dirExists, err := resultTuple[0].AsBool()

			if err != nil {
				return datavalue.Error(err)
			}

			if dirExists {
				return datavalue.Error(fmt.Errorf("directory %v already exists", path))
			}

			err = os.MkdirAll(path, 0700)

			if err != nil {
				return datavalue.Error(err)
			}

			return datavalue.Error(nil)
		},
	)
}
