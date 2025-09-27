package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getLastIndexOfFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "lastIndexOf",
			Description: "Gets the index of the last occurrence of a substring in a string.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.lastIndexOf("Hello World Hello", "Hello") // returns 12`, packageName),
				fmt.Sprintf(`%s.lastIndexOf("Hello World Hello", "World") // returns 6`, packageName),
				fmt.Sprintf(`%s.lastIndexOf("Hello World Hello", "xyz") // returns -1`, packageName),
				fmt.Sprintf(`%s.lastIndexOf("Hello World Hello", "") // returns 17`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to search in.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "substr",
				Description: "The substring to search for.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The index of the last occurrence, or -1 if not found.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			substr, _ := args[1].AsString()

			index := strings.LastIndex(str, substr)

			return datavalue.Number(float64(index))
		},
	)
}
