package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getReplaceAllFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "replaceAll",
			Description: "Replaces all occurrences of a substring in a string with a new substring.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.replaceAll("test some test string", "test", "new") // returns "new some new string"`, packageName),
				fmt.Sprintf(`%s.replaceAll("test some test string", "bogus", "new") // returns "test some test string"`, packageName),
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
				Description: "The new string with the substring replaced all occurrences of the substring.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			substr, _ := args[1].AsString()
			newSubstr, _ := args[2].AsString()

			return datavalue.String(strings.ReplaceAll(str, substr, newSubstr))
		},
	)
}
