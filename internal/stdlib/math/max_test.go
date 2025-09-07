package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetMaxFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["max"]; !ok {
		t.Fatalf("expected max function, got %v", functions)
	}

	maxFunc := functions["max"]

	if maxFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", maxFunc.FunctionType)
	}

	if maxFunc.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", maxFunc.ArgKinds[0])
	}

	result, err := maxFunc.Handler(
		nil,
		[]datavalue.Value{datavalue.Number(1.5), datavalue.Number(2.5)},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != math.Max(1.5, 2.5) {
		t.Fatalf("expected %f, got %v", math.Max(1.5, 2.5), result.Num)
	}
}
