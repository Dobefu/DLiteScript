package math

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getTanFunction() function.Info {
	return function.MakeFunction(
		function.FunctionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Tan(arg0))
		},
	)
}
