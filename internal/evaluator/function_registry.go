package evaluator

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

type functionHandler func(
	e *Evaluator,
	args []datavalue.Value,
) (datavalue.Value, error)

type functionInfo struct {
	handler  functionHandler
	argCount int
	argKinds []datatype.DataType
}

func makeFunction(
	argKinds []datatype.DataType,
	impl func(e *Evaluator, args []datavalue.Value) datavalue.Value,
) functionInfo {
	handler := func(
		e *Evaluator,
		args []datavalue.Value,
	) (datavalue.Value, error) {
		for i, arg := range args {
			if arg.DataType() == argKinds[i] {
				continue
			}

			return datavalue.Null(), errorutil.NewError(
				errorutil.ErrorMsgTypeUnknownDataType,
				arg.DataType().AsString(),
			)
		}

		return impl(e, args), nil
	}

	return functionInfo{
		handler:  handler,
		argCount: len(argKinds),
		argKinds: argKinds,
	}
}

var functionRegistry = map[string]functionInfo{
	"println": makeFunction(
		[]datatype.DataType{datatype.DataTypeString},
		func(e *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsString()

			fmt.Fprintf(&e.buf, "%s", arg0)

			return datavalue.Null()
		},
	),

	"abs": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Abs(arg0))
		},
	),
	"sin": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Sin(arg0))
		},
	),
	"cos": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Cos(arg0))
		},
	),
	"tan": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Tan(arg0))
		},
	),
	"sqrt": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Sqrt(arg0))
		},
	),
	"round": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Round(arg0))
		},
	),
	"floor": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Floor(arg0))
		},
	),
	"ceil": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()

			return datavalue.Number(math.Ceil(arg0))
		},
	),
	"min": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber, datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()
			arg1, _ := args[1].AsNumber()

			return datavalue.Number(math.Min(arg0, arg1))
		},
	),
	"max": makeFunction(
		[]datatype.DataType{datatype.DataTypeNumber, datatype.DataTypeNumber},
		func(_ *Evaluator, args []datavalue.Value) datavalue.Value {
			arg0, _ := args[0].AsNumber()
			arg1, _ := args[1].AsNumber()

			return datavalue.Number(math.Max(arg0, arg1))
		},
	),
}
