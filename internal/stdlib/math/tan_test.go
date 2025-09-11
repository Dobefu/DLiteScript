package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetTanFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["tan"]; !ok {
		t.Fatalf("expected tan function, got %v", functions)
	}

	tan := functions["tan"]

	if tan.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", tan.FunctionType)
	}

	if tan.Parameters[0].Type != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", tan.Parameters[0].Type)
	}

	result, err := tan.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != math.Tan(1.5) {
		t.Fatalf("expected %f, got %v", math.Tan(1.5), result.Num)
	}
}
