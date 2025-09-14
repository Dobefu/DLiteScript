package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSqrtFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sqrt",
			Description: "Returns the square root of a number.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.sqrt(4) // returns 2", packageName),
				fmt.Sprintf("%s.sqrt(16) // returns 4", packageName),
				fmt.Sprintf("%s.sqrt(0) // returns 0", packageName),
				fmt.Sprintf("%s.sqrt(-1) // returns null", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the square root of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The square root of the provided number.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			num, _ := args[0].AsNumber()

			sqrt := math.Sqrt(num)

			if math.IsNaN(sqrt) {
				return datavalue.Null()
			}

			return datavalue.Number(sqrt)
		},
	)
}
