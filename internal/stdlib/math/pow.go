package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getPowFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "pow",
			Description: "Gets a number raised to the power of another number.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.pow(2, 3) // returns 8", packageName),
				fmt.Sprintf("%s.pow(10, 2) // returns 100", packageName),
				fmt.Sprintf("%s.pow(4, 0.5) // returns 2", packageName),
				fmt.Sprintf("%s.pow(0, 0) // returns 1", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "base",
				Description: "The base number.",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "exponent",
				Description: "The exponent to raise the base number to.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The number raised to the power of the exponent.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			base, _ := args[0].AsNumber()
			exponent, _ := args[1].AsNumber()

			result := math.Pow(base, exponent)

			return datavalue.Number(result)
		},
	)
}
