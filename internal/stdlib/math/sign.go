package math

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSignFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sign",
			Description: "Gets the sign of a number: -1 for negative, 0 for zero, 1 for positive.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.sign(-5) // returns -1", packageName),
				fmt.Sprintf("%s.sign(0) // returns 0", packageName),
				fmt.Sprintf("%s.sign(0.1) // returns 1", packageName),
				fmt.Sprintf("%s.sign(-0) // returns 0", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the sign of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The sign of the provided number: -1, 0, or 1.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			num, _ := args[0].AsNumber()

			if num > 0 {
				return datavalue.Number(1)
			}

			if num < 0 {
				return datavalue.Number(-1)
			}

			return datavalue.Number(0)
		},
	)
}
