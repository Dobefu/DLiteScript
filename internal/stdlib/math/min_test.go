package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetMinFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["min"]; !ok {
		t.Fatalf("expected min function, got %v", functions)
	}

	minFunc := functions["min"]

	if minFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", minFunc.FunctionType)
	}

	if minFunc.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", minFunc.ArgKinds[0])
	}

	result, err := minFunc.Handler(
		nil,
		[]datavalue.Value{datavalue.Number(1.5), datavalue.Number(2.5)},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != math.Min(1.5, 2.5) {
		t.Fatalf("expected %f, got %v", math.Min(1.5, 2.5), result.Num)
	}
}
