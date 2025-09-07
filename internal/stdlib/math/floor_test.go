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

	floorFunc := functions["floor"]

	if floorFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", floorFunc.FunctionType)
	}

	if floorFunc.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", floorFunc.ArgKinds[0])
	}

	result, err := floorFunc.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != 1 {
		t.Fatalf("expected 1, got %v", result.Num)
	}
}
