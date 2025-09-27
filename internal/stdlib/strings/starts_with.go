package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getStartsWithFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "startsWith",
			Description: "Checks if a string starts with a given substring.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.startsWith("Hello World", "Hello") // returns true`, packageName),
				fmt.Sprintf(`%s.startsWith("Hello World", "World") // returns false`, packageName),
				fmt.Sprintf(`%s.startsWith("Hello World", "") // returns true`, packageName),
				fmt.Sprintf(`%s.startsWith("", "Hello") // returns false`, packageName),
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
				Name:        "prefix",
				Description: "The substring to check for.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeBool,
				Name:        "result",
				Description: "True if the string starts with the substring, false otherwise.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			prefix, _ := args[1].AsString()

			return datavalue.Bool(strings.HasPrefix(str, prefix))
		},
	)
}
