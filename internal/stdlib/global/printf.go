package global

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getPrintfFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "printf",
			Description: "Prints a formatted string.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				"printf(\"test %s\", \"string\") // prints \"test string\"",
				"printf(\"test %d\", 1) // prints \"test 1\"",
				"printf(\"test %t\", true) // prints \"test true\"",
				"printf(\"test %s\", null) // prints \"test null\"",
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
		[]function.ArgInfo{},
		true,
		func(e function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			format, _ := args[0].AsString()
			format = strings.ReplaceAll(format, "%d", "%f")

			if len(args) == 1 {
				e.AddToBuffer(format)

				return datavalue.Null()
			}

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

				case datatype.DataTypeArray:
					formatArgs[i-1] = args[i].ToString()

				case datatype.DataTypeAny:
					formatArgs[i-1] = args[i].ToString()
				}
			}

			e.AddToBuffer(format, formatArgs...)

			return datavalue.Null()
		},
	)
}
