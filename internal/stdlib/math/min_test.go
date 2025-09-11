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

	if minFunc.FunctionType != function.FunctionTypeVariadic {
		t.Fatalf("expected variadic function, got %T", minFunc.FunctionType)
	}

	if minFunc.Parameters[0].Type != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", minFunc.Parameters[0].Type)
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

	result, err = minFunc.Handler(
		nil,
		[]datavalue.Value{},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.DataType() != datatype.DataTypeNull {
		t.Fatalf("expected null for no arguments, got %v", result.DataType())
	}
}
