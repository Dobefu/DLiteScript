package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSliceFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "slice",
			Description: "Extracts a portion of an array.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 1, 4) // returns [2, 3, 4]`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 0, 3) // returns [1, 2, 3]`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 2, 2) // returns []`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 0, 0) // returns []`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], -1, 0) // returns []`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 0, -1) // returns [1, 2, 3, 4]`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 6, 1) // returns []`, packageName),
				fmt.Sprintf(`%s.slice([1, 2, 3, 4, 5], 0, 6) // returns [1, 2, 3, 4, 5]`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to slice.",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "start",
				Description: "The starting index (inclusive).",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "end",
				Description: "The ending index (exclusive).",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "result",
				Description: "The sliced portion of the array.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()
			startFloat, _ := args[1].AsNumber()
			endFloat, _ := args[2].AsNumber()

			start := int(startFloat)
			end := int(endFloat)
			arrLen := len(arr)

			if start < 0 {
				start = max(arrLen+start, 0)
			}

			if end < 0 {
				end = max(arrLen+end, 0)
			}

			if start > arrLen {
				start = arrLen
			}

			if end > arrLen {
				end = arrLen
			}

			if start > end {
				start = end
			}

			slice := arr[start:end]

			return datavalue.Array(slice...)
		},
	)
}
