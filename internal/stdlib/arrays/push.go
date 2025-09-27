package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getPushFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "push",
			Description: "Pushes an arbitrary number of values to an array.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.push([1, 2, 3], 4, 5, 6) // returns [1, 2, 3, 4, 5, 6]", packageName),
				fmt.Sprintf("%s.push([1, 2, 3], [4, 5, 6]) // returns [1, 2, 3, 4, 5, 6]", packageName),
			},
		},
		packageName,
		function.FunctionTypeMixedVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to push the values to.",
			},
			{
				Type:        datatype.DataTypeAny,
				Name:        "...values",
				Description: "The values to push to the array.",
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
