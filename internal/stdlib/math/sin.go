package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSinFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sin",
			Description: "Returns the sine value of a number.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.sin(1.5) // returns 0.9974949866040544", packageName),
				fmt.Sprintf("%s.sin(1) // returns 0.8414709848078965", packageName),
				fmt.Sprintf("%s.sin(0) // returns 0", packageName),
				fmt.Sprintf("%s.sin(-1.5) // returns -0.9974949866040544", packageName),
				fmt.Sprintf("%s.sin(-1) // returns -0.8414709848078965", packageName),
				fmt.Sprintf("%s.sin(-0) // returns 0", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the sine value of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The sine value of the provided number.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Sin(arg0))
		},
	)
}
