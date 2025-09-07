package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetAbsFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["abs"]; !ok {
		t.Fatalf("expected abs function, got %v", functions)
	}

	abs := functions["abs"]

	if abs.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", abs.FunctionType)
	}

	if abs.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", abs.ArgKinds[0])
	}

	result, err := abs.Handler(nil, []datavalue.Value{datavalue.Number(-1)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != 1 {
		t.Fatalf("expected 1, got %v", result.Num)
	}
}
