package os

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSetEnvVariableFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "setEnvVariable",
			Description: "Sets the value of an environment variable.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.setEnvVariable("VARIABLE", "value") // sets the VARIABLE environment variable`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "name",
				Description: "The name of the environment variable to set.",
			},
			{
				Type:        datatype.DataTypeString,
				Name:        "value",
				Description: "The value to set for the environment variable.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeError,
				Name:        "error",
				Description: "An error if the environment variable cannot be set.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			name, _ := args[0].AsString()
			value, _ := args[1].AsString()

			err := os.Setenv(name, value)

			return datavalue.Error(err)
		},
	)
}
