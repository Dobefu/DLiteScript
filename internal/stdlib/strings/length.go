package strings

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getLengthFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "length",
			Description: "Gets the length of a string.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.length("") // returns 0`, packageName),
				fmt.Sprintf(`%s.length("test") // returns 4`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to get the length of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "length",
				Description: "The length of the string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()

			return datavalue.Number(float64(len(str)))
		},
	)
}
