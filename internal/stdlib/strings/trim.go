package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getTrimFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "trim",
			Description: "Removes whitespace from the beginning and end of a string.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.trim(" test ") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trim(" test") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trim("test ") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trim("test") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trim("   ") // returns ""`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to trim.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The trimmed string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()

			return datavalue.String(strings.TrimSpace(str))
		},
	)
}
