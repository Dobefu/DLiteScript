package global

import (
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getExitFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "exit",
			Description: "Exits the running script with a given exit code.",
			Since:       "v0.1.0",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				"exit(0) // exits the program with a success status",
				"exit(1) // exits the program with an error status",
			},
		},
		packageName,
		function.FunctionTypeVariadic,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "code",
				Description: "The exit code to return.",
			},
		},
		[]function.ArgInfo{},
		true,
		func(e function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			code, _ := args[0].AsNumber()
			e.Terminate(byte(code))

			return datavalue.Null()
		},
	)
}
