package arrays

import (
	"errors"
	"fmt"
	"sort"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSortStringsFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sortStrings",
			Description: "Sorts an array of strings in ascending order.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.sortStrings(["c", "b", "a"]) // returns ["a", "b", "c"]`, packageName),
				fmt.Sprintf(`%s.sortStrings([3, "hello", 1]) // returns error("array contains non-strings")`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array of strings to sort.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "result",
				Description: "The sorted array of strings.",
			},
			{
				Type:        datatype.DataTypeError,
				Name:        "err",
				Description: "An error if the array contains non-strings.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()

			for _, element := range arr {
				_, err := element.AsString()

				if err != nil {
					return datavalue.Tuple(
						datavalue.Null(),
						datavalue.Error(errors.New("array contains non-strings")),
					)
				}
			}

			sortedArr := make([]datavalue.Value, len(arr))
			copy(sortedArr, arr)

			sort.Slice(sortedArr, func(a, b int) bool {
				strA, _ := sortedArr[a].AsString()
				strB, _ := sortedArr[b].AsString()

				return strA < strB
			})

			return datavalue.Tuple(datavalue.Array(sortedArr...), datavalue.Null())
		},
	)
}
