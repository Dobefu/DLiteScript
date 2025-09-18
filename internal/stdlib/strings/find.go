package strings

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getFindFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "find",
			Description: "Returns the index of the first occurrence of a substring in a string.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.find("some test string", "") // returns 0`, packageName),
				fmt.Sprintf(`%s.find("some test string", "test") // returns 5`, packageName),
				fmt.Sprintf(`%s.find("some test string", "bogus") // returns -1`, packageName),
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
				Name:        "index",
				Description: "The index of the first occurrence of the substring in the string.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			substr, _ := args[1].AsString()

			return datavalue.Number(float64(strings.Index(str, substr)))
		},
	)
}
