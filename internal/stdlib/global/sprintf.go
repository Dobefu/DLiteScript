package global

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSprintfFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sprintf",
			Description: "Returns a formatted string.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				"sprintf(\"test %s\", \"string\") // returns \"test string\"",
				"sprintf(\"test %d\", 1) // returns \"test 1\"",
				"sprintf(\"test %t\", true) // returns \"test true\"",
				"sprintf(\"test %s\", null) // returns \"test null\"",
			},
		},
		packageName,
		function.FunctionTypeMixedVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "format",
				Description: "The format string.",
			},
			{
				Type:        datatype.DataTypeAny,
				Name:        "...args",
				Description: "The arguments to format.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The formatted string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			format, _ := args[0].AsString()
			format = strings.ReplaceAll(format, "%d", "%f")

			if len(args) == 1 {
				return datavalue.String(format)
			}

			formatArgs := make([]any, len(args)-1)

			for i := 1; i < len(args); i++ {
				switch args[i].DataType {
				case
					datatype.DataTypeString:
					str, _ := args[i].AsString()
					formatArgs[i-1] = str

				case
					datatype.DataTypeNumber:
					num, _ := args[i].AsNumber()
					formatArgs[i-1] = num

				case
					datatype.DataTypeBool:
					num, _ := args[i].AsBool()
					formatArgs[i-1] = num

				case
					datatype.DataTypeNull:
					formatArgs[i-1] = "null"

				case
					datatype.DataTypeFunction:
					formatArgs[i-1] = "function"

				case
					datatype.DataTypeTuple,
					datatype.DataTypeArray,
					datatype.DataTypeError,
					datatype.DataTypeAny:
					formatArgs[i-1] = args[i].ToString()
				}
			}

			result := fmt.Sprintf(format, formatArgs...)

			return datavalue.String(result)
		},
	)
}
