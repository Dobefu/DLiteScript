package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getAbsFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "abs",
			Description: "Gets the absolute value of a number.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.abs(-1) // returns 1", packageName),
				fmt.Sprintf("%s.abs(1) // returns 1", packageName),
				fmt.Sprintf("%s.abs(0) // returns 0", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the absolute value of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The absolute value of the provided number.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Abs(arg0))
		},
	)
}
