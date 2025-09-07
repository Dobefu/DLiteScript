package global

import (
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getPrintfFunction() function.Info {
	return function.MakeFunction(
		function.FunctionTypeMixedVariadic,
		[]datatype.DataType{datatype.DataTypeString},
		func(e function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
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

				case datatype.DataTypeFunction:
					formatArgs[i-1] = "function"

				case datatype.DataTypeTuple:
					formatArgs[i-1] = args[i].ToString()

				default:
					formatArgs[i-1] = "unknown"
				}
			}

			e.AddToBuffer(format, formatArgs...)

			return datavalue.Null()
		},
	)
}
