package global

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getDumpFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "dump",
			Description: "Outputs a formatted representation of the provided value.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				`dump("test") // prints "string: "test"`,
				"dump([1, 2, 3]) // prints the formatted array",
				"dump(1, 2, 3) // prints the values",
			},
		},
		packageName,
		function.FunctionTypeVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeAny,
				Name:        "...values",
				Description: "The values to dump.",
			},
		},
		[]function.ArgInfo{},
		true,
		func(e function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			for _, arg := range args {
				dumpSingleValue(e, arg, 0)
			}

			return datavalue.Null()
		},
	)
}

func dumpSingleValue(
	e function.EvaluatorInterface,
	value datavalue.Value,
	indent int,
) {
	indentStr := strings.Repeat("  ", indent)

	switch value.DataType() {
	case datatype.DataTypeString:
		e.AddToBuffer(fmt.Sprintf("%s\"%s\"\n", indentStr, value.ToString()))

	case datatype.DataTypeNumber:
		e.AddToBuffer(fmt.Sprintf("%s%s\n", indentStr, value.ToString()))

	case datatype.DataTypeBool:
		e.AddToBuffer(fmt.Sprintf("%s%s\n", indentStr, value.ToString()))

	case datatype.DataTypeNull:
		e.AddToBuffer(fmt.Sprintf("%snull\n", indentStr))

	case datatype.DataTypeFunction:
		e.AddToBuffer(fmt.Sprintf("%sfunction\n", indentStr))

	case datatype.DataTypeArray:
		e.AddToBuffer(fmt.Sprintf("%sarray[%d]:\n", indentStr, len(value.Values)))

		for i, item := range value.Values {
			e.AddToBuffer(fmt.Sprintf("%s  [%d]: ", indentStr, i))

			dumpSingleValue(e, item, indent+1)
		}

	case datatype.DataTypeTuple:
		e.AddToBuffer(fmt.Sprintf("%stuple[%d]:\n", indentStr, len(value.Values)))

		for i, item := range value.Values {
			e.AddToBuffer(fmt.Sprintf("%s  (%d): ", indentStr, i))
			dumpSingleValue(e, item, indent+1)
		}

	case datatype.DataTypeAny:
		e.AddToBuffer(fmt.Sprintf("%s%s\n", indentStr, value.ToString()))
	}
}
