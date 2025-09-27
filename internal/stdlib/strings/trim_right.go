package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getTrimRightFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "trimRight",
			Description: "Removes whitespace from the end of a string.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.trimRight(" test ") // returns " test"`, packageName),
				fmt.Sprintf(`%s.trimRight(" test") // returns " test"`, packageName),
				fmt.Sprintf(`%s.trimRight("test ") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trimRight("test") // returns "test"`, packageName),
				fmt.Sprintf(`%s.trimRight("   ") // returns ""`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to trim the right side of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The right trimmed string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()

			return datavalue.String(strings.TrimRight(str, " "))
		},
	)
}
