package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getLog10Function() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "log10",
			Description: "Returns the base 10 logarithm of a number.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.log10(1) // returns 0", packageName),
				fmt.Sprintf("%s.log10(10) // returns 1", packageName),
				fmt.Sprintf("%s.log10(100) // returns 2", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the base 10 logarithm of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The base 10 logarithm of the provided number.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			num, _ := args[0].AsNumber()

			return datavalue.Number(math.Log10(num))
		},
	)
}
