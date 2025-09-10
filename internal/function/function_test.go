package function

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestMakeFunction(t *testing.T) {
	t.Parallel()

	function := MakeFunction(
		FunctionTypeFixed,
		[]datatype.DataType{datatype.DataTypeNumber},
		func(_ EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			num, err := args[0].AsNumber()

			if err != nil {
				return datavalue.Number(0)
			}

			return datavalue.Number(num)
		},
	)

	result, err := function.Handler(nil, []datavalue.Value{datavalue.Number(1)})

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result.DataType() != datatype.DataTypeNumber {
		t.Errorf("expected DataTypeNumber, got %v", result.DataType())
	}
}
