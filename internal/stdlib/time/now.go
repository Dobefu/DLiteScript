package time

import (
	"fmt"
	"time"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getNowFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "now",
			Description: "Gets the current Unix timestamp.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf("%s.now() // returns the current timestamp", packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "timestamp",
				Description: "The current Unix timestamp.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, _ []datavalue.Value) datavalue.Value {
			return datavalue.Number(float64(time.Now().Unix()))
		},
	)
}
