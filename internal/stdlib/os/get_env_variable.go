package os

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getGetEnvVariableFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "getEnvVariable",
			Description: "Gets the value of an environment variable.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.getEnvVariable("PATH") // returns something like "/usr/bin:/bin"`, packageName),
				fmt.Sprintf(`%s.getEnvVariable("HOME") // returns something like "/home/user"`, packageName),
				fmt.Sprintf(`%s.getEnvVariable("BOGUS") // returns ""`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "name",
				Description: "The name of the environment variable to get.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeString,
				Name:        "value",
				Description: "The value of the environment variable, or empty string if not found.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			name, _ := args[0].AsString()

			value := os.Getenv(name)

			return datavalue.String(value)
		},
	)
}
