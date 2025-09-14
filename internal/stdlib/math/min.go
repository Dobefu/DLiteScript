package math

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getMinFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "min",
			Description: "Returns the smaller of two provided numbers.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.min(1, 2, 3) // returns 1", packageName),
				fmt.Sprintf("%s.min(1.5, 2.5, 3.5) // returns 1.5", packageName),
				fmt.Sprintf("%s.min(-1, -2, -3) // returns -3", packageName),
				fmt.Sprintf("%s.min(-1.5, -2.5, -3.5) // returns -3.5", packageName),
			},
		},
		packageName,
		function.FunctionTypeVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "...nums",
				Description: "The numbers to find the minimum of. At least two numbers are required.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The smallest of the provided numbers.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			if len(args) < 2 {
				return datavalue.Null()
			}

			minValue, _ := args[0].AsNumber()

			for _, arg := range args[1:] {
				num, _ := arg.AsNumber()

				if num < minValue {
					minValue = num
				}
			}

			return datavalue.Number(minValue)
		},
	)
}
