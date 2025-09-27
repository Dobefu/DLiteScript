package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSpliceFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "splice",
			Description: "Removes elements from an array and inserts new elements at the same position.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.splice([1, 2, 3, 4, 5], 1, 2) // returns ([2, 3], [1, 4, 5])`, packageName),
				fmt.Sprintf(`%s.splice([1, 2, 3, 4, 5], 1, 2, 6, 7) // returns ([2, 3], [1, 6, 7, 4, 5])`, packageName),
				fmt.Sprintf(`%s.splice([1, 2, 3], 0, 0, 0) // returns ([], [0, 1, 2, 3])`, packageName),
				fmt.Sprintf(`%s.splice([1, 2, 3, 4, 5], -1, 1) // returns ([5], [1, 2, 3, 4])`, packageName),
				fmt.Sprintf(`%s.splice([1, 2, 3, 4, 5], 6, 1) // returns ([], [1, 2, 3, 4, 5])`, packageName),
				fmt.Sprintf(`%s.splice([1, 2, 3, 4, 5], 0, -1) // returns ([1, 2, 3, 4, 5], [])`, packageName),
				fmt.Sprintf(`%s.splice([1, 2, 3, 4, 5], 0, 6) // returns ([1, 2, 3, 4, 5], [])`, packageName),
			},
		},
		packageName,
		function.FunctionTypeMixedVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to splice.",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "start",
				Description: "The index at which to start removing elements.",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "deleteCount",
				Description: "The number of elements to remove.",
			},
			{
				Type:        datatype.DataTypeAny,
				Name:        "...items",
				Description: "Optional elements to insert at the start index.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "removed",
				Description: "Array of removed elements.",
			},
			{
				Type:        datatype.DataTypeArray,
				Name:        "result",
				Description: "The modified array.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()
			startFloat, _ := args[1].AsNumber()
			deleteCountFloat, _ := args[2].AsNumber()

			start := int(startFloat)
			deleteCount := int(deleteCountFloat)
			arrLen := len(arr)

			if start < 0 {
				start = max(arrLen+start, 0)
			}

			if start > arrLen {
				start = arrLen
			}

			if deleteCount < 0 {
				deleteCount = 0
			}

			if deleteCount > arrLen-start {
				deleteCount = arrLen - start
			}

			removed := arr[start : start+deleteCount]

			newArr := make([]datavalue.Value, 0, arrLen-deleteCount+len(args)-3)
			newArr = append(newArr, arr[:start]...)
			newArr = append(newArr, args[3:]...)
			newArr = append(newArr, arr[start+deleteCount:]...)

			return datavalue.Tuple(datavalue.Array(removed...), datavalue.Array(newArr...))
		},
	)
}
