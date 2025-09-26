package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getLengthFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "length",
			Description: "Returns the length of an array.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.length([]) // returns 0", packageName),
				fmt.Sprintf("%s.length([1, 2, 3]) // returns 3", packageName),
				fmt.Sprintf("%s.length([1, 2, 3, 4, 5, 6]) // returns 6", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to get the length of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "length",
				Description: "The length of the array.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			array, _ := args[0].AsArray()

			return datavalue.Number(float64(len(array)))
		},
	)
}
