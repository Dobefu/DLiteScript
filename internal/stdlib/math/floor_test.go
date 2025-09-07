package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetFloorFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["floor"]; !ok {
		t.Fatalf("expected floor function, got %v", functions)
	}

	floor := functions["floor"]

	if floor.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", floor.FunctionType)
	}

	if floor.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", floor.ArgKinds[0])
	}

	result, err := floor.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != 1 {
		t.Fatalf("expected 1, got %v", result.Num)
	}
}
