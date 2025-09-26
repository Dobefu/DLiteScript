package arrays

import (
	"fmt"
	"slices"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getReverseFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "reverse",
			Description: "Reverses the order of elements in an array.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.reverse([1, 2, 3]) // returns [3, 2, 1]", packageName),
				fmt.Sprintf("%s.reverse(['a', 'b', 'c']) // returns ['c', 'b', 'a']", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to reverse.",
			},
		},
		[]function.ArgInfo{},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			array, _ := args[0].AsArray()
			slices.Reverse(array)

			return datavalue.Array(array...)
		},
	)
}
