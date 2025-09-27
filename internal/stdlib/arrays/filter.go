package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getFilterFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "filter",
			Description: "Filters out falsy values from an array.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.filter([1, 0, 2, false, 3, "", 4]) // returns ([0, false, ""], [1, 2, 3, 4])`, packageName),
				fmt.Sprintf(`%s.filter([true, false, "hello", ""]) // returns ([false, ""], [true, "hello"])`, packageName),
				fmt.Sprintf(`%s.filter([1, 2, 3]) // returns ([], [1, 2, 3])`, packageName),
				fmt.Sprintf(`%s.filter([]) // returns ([], [])`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to filter.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "filteredOut",
				Description: "Array of elements that were filtered out.",
			},
			{
				Type:        datatype.DataTypeArray,
				Name:        "remaining",
				Description: "Array of remaining elements.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()

			filteredOut := make([]datavalue.Value, 0)
			remaining := make([]datavalue.Value, 0)

			for _, element := range arr {
				if !element.IsTruthy() {
					filteredOut = append(filteredOut, element)

					continue
				}

				remaining = append(remaining, element)
			}

			return datavalue.Tuple(
				datavalue.Array(filteredOut...),
				datavalue.Array(remaining...),
			)
		},
	)
}
