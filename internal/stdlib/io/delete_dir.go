package io

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getDeleteDirFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "deleteDir",
			Description: "Deletes the given directory.",
			Since:       "v0.2.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.deleteDir("bad_memes") // deletes the given directory called "bad_memes" or something`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "path",
				Description: "The path for the folder that should be deleted.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "error",
				Description: "An error, if the folder cannot be deleted.",
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

			if result.Bool {
				err := os.RemoveAll(path)
				if err != nil {
					return datavalue.Error(err)
				}
				return datavalue.Error(nil)
			}
			return datavalue.Error(fmt.Errorf("Folder %v does not exist", path))
		},
	)
}
