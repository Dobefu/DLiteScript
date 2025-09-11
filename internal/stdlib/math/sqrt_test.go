package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetSqrtFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["sqrt"]; !ok {
		t.Fatalf("expected sqrt function, got %v", functions)
	}

	sqrtFunc := functions["sqrt"]

	if sqrtFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", sqrtFunc.FunctionType)
	}

	if sqrtFunc.Parameters[0].Type != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", sqrtFunc.Parameters[0].Type)
	}

	result, err := sqrtFunc.Handler(
		nil,
		[]datavalue.Value{datavalue.Number(1.5), datavalue.Number(2.5)},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != math.Sqrt(1.5) {
		t.Fatalf("expected %f, got %v", math.Sqrt(1.5), result.Num)
	}
}
