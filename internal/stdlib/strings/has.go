package strings

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getHasFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "has",
			Description: "Checks if a substring is present in a string.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				`has("some test string", "") // returns true`,
				`has("some test string", "test") // returns true`,
				`has("some test string", "bogus") // returns false`,
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
				Type:        datatype.DataTypeBool,
				Name:        "has",
				Description: "Whether or not the substring is present in the string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			substr, _ := args[1].AsString()

			return datavalue.Bool(strings.Contains(str, substr))
		},
	)
}
