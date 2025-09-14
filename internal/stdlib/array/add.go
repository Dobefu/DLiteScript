package array

import (
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getAddFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "add",
			Description: "Adds an arbitrary number of values to an array.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				"add([1, 2, 3], 4, 5, 6) // returns [1, 2, 3, 4, 5, 6]",
				"add([1, 2, 3], [4, 5, 6]) // returns [1, 2, 3, 4, 5, 6]",
			},
		},
		packageName,
		function.FunctionTypeMixedVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "array",
				Description: "The array to add the values to.",
			},
			{
				Type:        datatype.DataTypeAny,
				Name:        "...values",
				Description: "The values to add to the array.",
			},
		},
		[]function.ArgInfo{},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			array, _ := args[0].AsArray()

			for _, arg := range args[1:] {
				nestedArray, err := arg.AsArray()

				if err == nil {
					array = append(array, nestedArray...)

					continue
				}

				array = append(array, arg)
			}

			return datavalue.Array(array...)
		},
	)
}
