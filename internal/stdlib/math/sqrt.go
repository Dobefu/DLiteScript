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
			Description: " the square root of a number.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.sqrt(4) // returns (2, null)", packageName),
				fmt.Sprintf("%s.sqrt(16) // returns (4, null)", packageName),
				fmt.Sprintf("%s.sqrt(0) // returns (0, null)", packageName),
				fmt.Sprintf("%s.sqrt(-1) // returns (null, err)", packageName),
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
			{
				Type:        datatype.DataTypeError,
				Name:        "err",
				Description: "An error if the square root cannot be calculated.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			num, _ := args[0].AsNumber()

			sqrt := math.Sqrt(num)

			if math.IsNaN(sqrt) {
				return datavalue.Tuple(
					datavalue.Null(),
					datavalue.Error(
						fmt.Errorf(
							"cannot calculate square root of negative number: %f",
							num,
						),
					),
				)
			}

			return datavalue.Tuple(
				datavalue.Number(sqrt),
				datavalue.Null(),
			)
		},
	)
}
