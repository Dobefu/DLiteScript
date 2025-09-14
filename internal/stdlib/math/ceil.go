package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getCeilFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "ceil",
			Description: "Returns the input number, rounded up to the nearest whole number.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.ceil(1.5) // returns 2", packageName),
				fmt.Sprintf("%s.ceil(1.2) // returns 2", packageName),
				fmt.Sprintf("%s.ceil(1) // returns 1", packageName),
				fmt.Sprintf("%s.ceil(-1.5) // returns -1", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to round up.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The input number, rounded up to the nearest whole number.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Ceil(arg0))
		},
	)
}
