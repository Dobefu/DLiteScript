package errors

import (
	"errors"
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getNewFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "new",
			Description: "Creates a new error with a given message.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.new(\"some error\") // creates a new error", packageName),
			},
		},
		packageName,
		function.FunctionTypeVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "message",
				Description: "The message to create the error with.",
			},
		},
		[]function.ArgInfo{},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			message, _ := args[0].AsString()

			return datavalue.Error(errors.New(message))
		},
	)
}
