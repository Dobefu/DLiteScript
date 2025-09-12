package math

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getTanFunction() function.Info {
	return function.MakeFunction(
		"tan",
		"Returns the tangent value of a number.",
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "num",
				Description: "The number to find the tangent value of.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The tangent value of the provided number.",
			},
		},
		true,
		"v0.1.0",
		function.DeprecationInfo{IsDeprecated: false, Description: "", Version: ""},
		[]string{
			fmt.Sprintf("%s.tan(1.5) // returns 14.101419947171719", packageName),
			fmt.Sprintf("%s.tan(1) // returns 1.557407724654902", packageName),
			fmt.Sprintf("%s.tan(0) // returns 0", packageName),
			fmt.Sprintf("%s.tan(-1.5) // returns -14.101419947171719", packageName),
			fmt.Sprintf("%s.tan(-1) // returns -1.557407724654902", packageName),
			fmt.Sprintf("%s.tan(-0) // returns 0", packageName),
		},
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Tan(arg0))
		},
	)
}
