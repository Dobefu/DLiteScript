package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getTrimLeftFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "trimLeft",
			Description: "Removes whitespace from the beginning of a string.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.trimLeft(" test ") // returns "test "`, packageName),
				fmt.Sprintf(`%s.trimLeft(" test") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trimLeft("test ") // returns "test "`, packageName),
				fmt.Sprintf(`%s.trimLeft("test") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trimLeft("   ") // returns ""`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to trim the left side of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The left trimmed string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()

			return datavalue.String(strings.TrimLeft(str, " "))
		},
	)
}
