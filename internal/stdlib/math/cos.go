package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getCosFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "cos",
			Description: "Gets the cosine value of a number.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.cos(1.5) // returns 0.0707372016677029", packageName),
				fmt.Sprintf("%s.cos(1) // returns 0.5403023058681398", packageName),
				fmt.Sprintf("%s.cos(0) // returns 1", packageName),
				fmt.Sprintf("%s.cos(-1.5) // returns 0.0707372016677029", packageName),
				fmt.Sprintf("%s.cos(-1) // returns 0.5403023058681398", packageName),
				fmt.Sprintf("%s.cos(-0) // returns 1", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the cosine value of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The cosine value of the provided number.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Cos(arg0))
		},
	)
}
