package math

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getMinFunction() function.Info {
	return function.MakeFunction(
		function.FunctionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber, datatype.DataTypeNumber},
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()
			arg1, _ := args[1].AsNumber()

			return datavalue.Number(math.Min(arg0, arg1))
		},
	)
}
