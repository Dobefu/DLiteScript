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

	if maxFunc.FunctionType != function.FunctionTypeVariadic {
		t.Fatalf("expected variadic function, got %T", maxFunc.FunctionType)
	}

	if maxFunc.Parameters[0].Type != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", maxFunc.Parameters[0].Type)
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

	result, err = maxFunc.Handler(
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
