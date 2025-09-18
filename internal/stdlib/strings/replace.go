package strings

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getReplaceFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "replace",
			Description: "Replaces a substring in a string with a new substring.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				`replace("some test string", "test", "new") // returns "some new string"`,
				`replace("some test string", "bogus", "new") // returns "some test string"`,
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to replace the substring in.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "substr",
				Description: "The substring to replace.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "newSubstr",
				Description: "The new substring to replace the substring with.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "newStr",
				Description: "The new string with the substring replaced.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			substr, _ := args[1].AsString()
			newSubstr, _ := args[2].AsString()

			return datavalue.String(strings.Replace(str, substr, newSubstr, 1))
		},
	)
}
