package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetRoundFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["round"]; !ok {
		t.Fatalf("expected round function, got %v", functions)
	}

	roundFunc := functions["round"]

	if roundFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", roundFunc.FunctionType)
	}

	if roundFunc.Parameters[0].Type != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", roundFunc.Parameters[0].Type)
	}

	result, err := roundFunc.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != 2 {
		t.Fatalf("expected 2, got %v", result.Num)
	}
}
