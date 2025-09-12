package math

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getMaxFunction() function.Info {
	return function.MakeFunction(
		"max",
		"Returns the larger of two provided numbers.",
		packageName,
		function.FunctionTypeVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "...nums",
				Description: "The numbers to process. At least two numbers are required.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The largest of the provided numbers.",
			},
		},
		true,
		"v0.1.0",
		function.DeprecationInfo{IsDeprecated: false, Description: "", Version: ""},
		[]string{
			fmt.Sprintf("%s.max(1, 2, 3) // returns 3", packageName),
			fmt.Sprintf("%s.max(1.5, 2.5, 3.5) // returns 3.5", packageName),
			fmt.Sprintf("%s.max(-1, -2, -3) // returns -1", packageName),
			fmt.Sprintf("%s.max(-1.5, -2.5, -3.5) // returns -1.5", packageName),
		},
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			if len(args) < 2 {
				return datavalue.Null()
			}

			maxValue, _ := args[0].AsNumber()

			for _, arg := range args[1:] {
				num, _ := arg.AsNumber()

				if num > maxValue {
					maxValue = num
				}
			}

			return datavalue.Number(maxValue)
		},
	)
}
