package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getIndexOfFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "indexOf",
			Description: "Returns the index of the first occurrence of a substring in a string.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.indexOf("Hello World", "Hello") // returns 0`, packageName),
				fmt.Sprintf(`%s.indexOf("Hello World", "World") // returns 6`, packageName),
				fmt.Sprintf(`%s.indexOf("Hello World", "xyz") // returns -1`, packageName),
				fmt.Sprintf(`%s.indexOf("Hello World", "") // returns 0`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to search in.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "substr",
				Description: "The substring to search for.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "result",
				Description: "The index of the first occurrence, or -1 if not found.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			substr, _ := args[1].AsString()

			index := strings.Index(str, substr)

			return datavalue.Number(float64(index))
		},
	)
}
