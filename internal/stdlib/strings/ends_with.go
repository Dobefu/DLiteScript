package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getEndsWithFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "endsWith",
			Description: "Checks if a string ends with a given substring.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.endsWith("Hello World", "World") // returns true`, packageName),
				fmt.Sprintf(`%s.endsWith("Hello World", "Hello") // returns false`, packageName),
				fmt.Sprintf(`%s.endsWith("Hello World", "") // returns true`, packageName),
				fmt.Sprintf(`%s.endsWith("", "World") // returns false`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to check.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "substr",
				Description: "The substring to check for.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeBool,
				Name:        "result",
				Description: "True if the string ends with the substring, false otherwise.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			suffix, _ := args[1].AsString()

			return datavalue.Bool(strings.HasSuffix(str, suffix))
		},
	)
}
