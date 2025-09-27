package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSplitFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "split",
			Description: "Splits a string into an array of strings using a delimiter.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.split("a,b,c", ",") // returns ["a", "b", "c"]`, packageName),
				fmt.Sprintf(`%s.split("a b c", " ") // returns ["a", "b", "c"]`, packageName),
				fmt.Sprintf(`%s.split("a b c", ",") // returns ["a b c"]`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to split.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "delimiter",
				Description: "The delimiter to split the string on.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeArray,
				Name:        "result",
				Description: "An array of strings, split on the delimiter.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			delimiter, _ := args[1].AsString()

			parts := strings.Split(str, delimiter)
			result := make([]datavalue.Value, len(parts))

			for i, part := range parts {
				result[i] = datavalue.String(part)
			}

			return datavalue.Array(result...)
		},
	)
}
