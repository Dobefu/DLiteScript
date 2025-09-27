package arrays

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getPopFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "pop",
			Description: "Removes and gets the last element from an array.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.pop([1, 2, 3, 4]) // returns (4, [1, 2, 3])`, packageName),
				fmt.Sprintf(`%s.pop(["hello", "world"]) // returns ("world", ["hello"])`, packageName),
				fmt.Sprintf(`%s.pop([]) // returns (null, [])`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array to remove the last element from.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeAny,
				Name:        "element",
				Description: "The last element that was removed, or null if array was empty.",
			},
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array with the last element removed.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()

			if len(arr) == 0 {
				return datavalue.Tuple(datavalue.Null(), datavalue.Array())
			}

			lastElement := arr[len(arr)-1]
			newArr := make([]datavalue.Value, len(arr)-1)
			copy(newArr, arr[:len(arr)-1])

			return datavalue.Tuple(lastElement, datavalue.Array(newArr...))
		},
	)
}
