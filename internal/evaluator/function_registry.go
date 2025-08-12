package evaluator

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

type functionType int

const (
	functionTypeFixed functionType = iota
	functionTypeVariadic
	functionTypeMixedVariadic
)

type functionHandler func(
	e *Evaluator,
	args []datavalue.Value,
) (datavalue.Value, error)

type functionInfo struct {
	handler      functionHandler
	functionType functionType
	argKinds     []datatype.DataType
}

func makeFunction(
	functionType functionType,
	argKinds []datatype.DataType,
	impl func(e *Evaluator, args []datavalue.Value) datavalue.Value,
) functionInfo {
	handler := func(e *Evaluator, args []datavalue.Value) (datavalue.Value, error) {
		return impl(e, args), nil
	}

	return functionInfo{
		handler:      handler,
		functionType: functionType,
		argKinds:     argKinds,
	}
}

var functionRegistry = map[string]functionInfo{
	"printf": makeFunction(
		functionTypeMixedVariadic,
		[]datatype.DataType{datatype.DataTypeString},
		func(e *Evaluator, args []datavalue.Value) datavalue.Value {
			format, _ := args[0].AsString()
			formatArgs := make([]any, len(args)-1)

			for i := 1; i < len(args); i++ {
				switch args[i].DataType() {
				case datatype.DataTypeString:
					str, _ := args[i].AsString()
					formatArgs[i-1] = str

				case datatype.DataTypeNumber:
					num, _ := args[i].AsNumber()
					formatArgs[i-1] = num

				case datatype.DataTypeBool:
					num, _ := args[i].AsBool()
					formatArgs[i-1] = num

				case datatype.DataTypeNull:
					formatArgs[i-1] = "null"

				default:
					formatArgs[i-1] = "unknown"
				}
			}

			fmt.Fprintf(&e.buf, format, formatArgs...)

			return datavalue.Null()
		},
	),
	"abs": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Abs(arg0))
		},
	),
	"sin": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Sin(arg0))
		},
	),
	"cos": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Cos(arg0))
		},
	),
	"tan": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Tan(arg0))
		},
	),
	"sqrt": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Sqrt(arg0))
		},
	),
	"round": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Round(arg0))
		},
	),
	"floor": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Floor(arg0))
		},
	),
	"ceil": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Ceil(arg0))
		},
	),
	"min": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber, datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()
			arg1, _ := args[1].AsNumber()

			return datavalue.Number(math.Min(arg0, arg1))
		},
	),
	"max": makeFunction(
		functionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber, datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()
			arg1, _ := args[1].AsNumber()

			return datavalue.Number(math.Max(arg0, arg1))
		},
	),
}
