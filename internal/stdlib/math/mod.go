package math

import (
	"errors"
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getModFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "mod",
			Description: "Gets the remainder of a number divided by another number.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.mod(2, 3) // returns (2, null)", packageName),
				fmt.Sprintf("%s.mod(10, 2) // returns (0, null)", packageName),
				fmt.Sprintf("%s.mod(4, 0.5) // returns (0, null)", packageName),
				fmt.Sprintf("%s.mod(0, 0) // returns (0, err)", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "base",
				Description: "The base number.",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "divisor",
				Description: "The divisor to divide the base number by.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The remainder of the base number divided by the divisor.",
			},
			{
				Type:        datatype.DataTypeError,
				Name:        "err",
				Description: "An error if the remainder cannot be calculated.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			base, _ := args[0].AsNumber()
			divisor, _ := args[1].AsNumber()

			result := math.Mod(base, divisor)

			if math.IsNaN(result) {
				return datavalue.Tuple(
					datavalue.Null(),
					datavalue.Error(errors.New("cannot mod by zero")),
				)
			}

			return datavalue.Tuple(datavalue.Number(result), datavalue.Null())
		},
	)
}
