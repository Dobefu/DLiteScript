package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getFloorFunction() function.Info {
	return function.MakeFunction(
		"floor",
		"Returns the input number, rounded down to the nearest whole number.",
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to process.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The input number, rounded down to the nearest whole number.",
			},
		},
		true,
		"v0.1.0",
		function.DeprecationInfo{IsDeprecated: false, Description: "", Version: ""},
		[]string{
			fmt.Sprintf("%s.floor(1.5) // returns 1", packageName),
			fmt.Sprintf("%s.floor(1.2) // returns 1", packageName),
			fmt.Sprintf("%s.floor(1) // returns 1", packageName),
			fmt.Sprintf("%s.floor(-1.5) // returns -2", packageName),
		},
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Floor(arg0))
		},
	)
}
