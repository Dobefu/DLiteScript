package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getContainsFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "contains",
			Description: "Checks if an array contains a given value.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.contains([1, 2, 3], 2) // returns true`, packageName),
				fmt.Sprintf(`%s.contains([1, 2, 3], 4) // returns false`, packageName),
				fmt.Sprintf(`%s.contains(["hello", "world"], "hello") // returns true`, packageName),
				fmt.Sprintf(`%s.contains(["hello", "world"], "hi") // returns false`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to search in.",
			},
			{
				Type:        datatype.DataTypeAny,
				Name:        "value",
				Description: "The value to search for.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeBool,
				Name:        "result",
				Description: "True if the array contains the value, false otherwise.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()
			searchValue := args[1]

			for _, element := range arr {
				if element.Equals(searchValue) {
					return datavalue.Bool(true)
				}
			}

			return datavalue.Bool(false)
		},
	)
}
