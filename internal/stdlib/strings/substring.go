package strings

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSubstringFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "substring",
			Description: "Extracts a substring from a string using start position and length.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.substring("Hello World", 0, 5) // returns ("Hello", null)`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", 6, 5) // returns ("World", null)`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", -5, 3) // returns ("Wor", null)`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", 0, -6) // returns ("Hello", null)`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", -5, -1) // returns ("Worl", null)`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", 20, 5) // returns ("", error("start index out of bounds: 20 >= 11"))`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", 0, -20) // returns ("", error("negative length results in empty string: length -20"))`, packageName),
				fmt.Sprintf(`%s.substring("Hello World", 0, 20) // returns ("Hello World", error("length exceeds string bounds: requested 20, available 11"))`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "str",
				Description: "The string to extract from.",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "start",
				Description: "The starting position (0-based index).",
			},
			{
				Type:        datatype.DataTypeNumber,
				Name:        "length",
				Description: "The number of characters to extract.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "result",
				Description: "The extracted substring.",
			},
			{
				Type:        datatype.DataTypeError,
				Name:        "err",
				Description: "An error if the indices are invalid.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			str, _ := args[0].AsString()
			startFloat, _ := args[1].AsNumber()
			lengthFloat, _ := args[2].AsNumber()

			start := int(startFloat)
			length := int(lengthFloat)
			strLen := len(str)

			if start < 0 {
				start = max(strLen+start, 0)
			}

			if start >= strLen {
				return datavalue.Tuple(
					datavalue.String(""),
					datavalue.Error(
						fmt.Errorf(
							"start index out of bounds: %d >= %d",
							start,
							strLen,
						),
					),
				)
			}

			if length < 0 {
				end := strLen + length

				if end <= start {
					return datavalue.Tuple(
						datavalue.String(""),
						datavalue.Error(
							fmt.Errorf(
								"negative length results in empty string: length %d",
								length,
							),
						),
					)
				}

				return datavalue.Tuple(
					datavalue.String(str[start:end]),
					datavalue.Null(),
				)
			}

			end := start + length

			if end > strLen {
				end = strLen

				return datavalue.Tuple(
					datavalue.String(str[start:end]),
					datavalue.Error(
						fmt.Errorf(
							"length exceeds string bounds: requested %d, available %d",
							length,
							end-start,
						),
					),
				)
			}

			return datavalue.Tuple(datavalue.String(str[start:end]), datavalue.Null())
		},
	)
}
