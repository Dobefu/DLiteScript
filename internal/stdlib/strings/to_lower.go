package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getToLowerFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "toLower",
			Description: "Converts a string to lowercase.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.toLower("TEST") // returns "test"`, packageName),
				fmt.Sprintf(`%s.toLower("TEST 123") // returns "test 123"`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to convert to lowercase.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The lowercase version of the string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()

			return datavalue.String(strings.ToLower(str))
		},
	)
}
