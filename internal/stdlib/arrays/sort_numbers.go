package arrays

import (
	"errors"
	"fmt"
	"sort"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSortNumbersFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sortNumbers",
			Description: "Sorts an array of numbers in ascending order.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.sortNumbers([3, 1, 4, 1, 5]) // returns [1, 1, 3, 4, 5]`, packageName),
				fmt.Sprintf(`%s.sortNumbers([1, 2, 10]) // returns [1, 2, 10]`, packageName),
				fmt.Sprintf(`%s.sortNumbers([3, "hello", 1]) // returns error("array contains non-numbers")`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array of numbers to sort.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "result",
				Description: "The sorted array of numbers.",
			},
			{
				Type:        datatype.DataTypeError,
				Name:        "err",
				Description: "An error if the array contains non-numbers.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()

			for _, element := range arr {
				_, err := element.AsNumber()

				if err != nil {
					return datavalue.Tuple(
						datavalue.Null(),
						datavalue.Error(errors.New("array contains non-numbers")),
					)
				}
			}

			sortedArr := make([]datavalue.Value, len(arr))
			copy(sortedArr, arr)

			sort.Slice(sortedArr, func(a, b int) bool {
				numA, _ := sortedArr[a].AsNumber()
				numB, _ := sortedArr[b].AsNumber()

				return numA < numB
			})

			return datavalue.Tuple(datavalue.Array(sortedArr...), datavalue.Null())
		},
	)
}
