package arrays

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getJoinFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "join",
			Description: "Joins an array of strings into a single string using a delimiter.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.join(["a", "b", "c"], ",") // returns "a,b,c"`, packageName),
				fmt.Sprintf(`%s.join(["a", "b", "c"], " ") // returns "a b c"`, packageName),
				fmt.Sprintf(`%s.join(["a"], ",") // returns "a"`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "arr",
				Description: "The array of strings to join.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "delimiter",
				Description: "The delimiter to join the strings with.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The joined string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arr, _ := args[0].AsArray()
			delimiter, _ := args[1].AsString()

			strValues := make([]string, len(arr))

			for i, val := range arr {
				str, _ := val.AsString()
				strValues[i] = str
			}

			return datavalue.String(strings.Join(strValues, delimiter))
		},
	)
}
