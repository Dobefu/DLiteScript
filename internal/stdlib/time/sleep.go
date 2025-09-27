package time

import (
	"fmt"
	"time"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func getSleepFunction() function.Info {
	return function.MakeFunction(
		function.Documentation{
			Name:        "sleep",
			Description: "Pauses execution for the specified number of milliseconds.",
			Since:       "v0.1.1",
			DeprecationInfo: function.DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{
				fmt.Sprintf(`%s.sleep(1000) // pauses execution for 1000 milliseconds (1 second)`, packageName),
				fmt.Sprintf(`%s.sleep(500) // pauses execution for 500 milliseconds (0.5 seconds)`, packageName),
				fmt.Sprintf(`%s.sleep(0) // no pause`, packageName),
			},
		},
		packageName,
		function.FunctionTypeFixed,
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNumber,
				Name:        "milliseconds",
				Description: "The number of milliseconds to pause execution.",
			},
		},
		[]function.ArgInfo{
			{
				Type:        datatype.DataTypeNull,
				Name:        "result",
				Description: "Always returns null after the pause.",
			},
		},
		true,
		func(_ function.EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			milliseconds, _ := args[0].AsNumber()

			if milliseconds < 0 {
				milliseconds = 0
			}

			time.Sleep(time.Duration(milliseconds) * time.Millisecond)

			return datavalue.Null()
		},
	)
}
