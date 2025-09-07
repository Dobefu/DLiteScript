package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetCeilFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["ceil"]; !ok {
		t.Fatalf("expected ceil function, got %v", functions)
	}

	ceilFunc := functions["ceil"]

	if ceilFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", ceilFunc.FunctionType)
	}

	if ceilFunc.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", ceilFunc.ArgKinds[0])
	}

	result, err := ceilFunc.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != 2 {
		t.Fatalf("expected 2, got %v", result.Num)
	}
}
