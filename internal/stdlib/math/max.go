package math

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getMaxFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "max",
			Description: "Gets the larger of two provided numbers.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.max(1, 2, 3) // returns (3, null)", packageName),
				fmt.Sprintf("%s.max(1.5, 2.5, 3.5) // returns (3.5, null)", packageName),
				fmt.Sprintf("%s.max(-1, -2, -3) // returns (-1, null)", packageName),
				fmt.Sprintf("%s.max(-1.5, -2.5, -3.5) // returns (-1.5, null)", packageName),
			},
		},
		packageName,
		function.FunctionTypeVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "...nums",
				Description: "The numbers to find the maximum of. At least two numbers are required.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The largest of the provided numbers.",
			},
			{
				Type:        datatype.DataTypeError,
				Name:        "err",
				Description: "An error if the maximum cannot be calculated.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			if len(args) < 2 {
				return datavalue.Tuple(
					datavalue.Null(),
					datavalue.Error(
						fmt.Errorf(
							"max requires at least 2 arguments, got %d",
							len(args),
						),
					),
				)
			}

			maxValue, _ := args[0].AsNumber()

			for _, arg := range args[1:] {
				num, _ := arg.AsNumber()

				if num > maxValue {
					maxValue = num
				}
			}

			return datavalue.Tuple(datavalue.Number(maxValue), datavalue.Null())
		},
	)
}
