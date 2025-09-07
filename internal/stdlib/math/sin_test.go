package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetSinFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["sin"]; !ok {
		t.Fatalf("expected sin function, got %v", functions)
	}

	sin := functions["sin"]

	if sin.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", sin.FunctionType)
	}

	if sin.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", sin.ArgKinds[0])
	}

	result, err := sin.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != math.Sin(1.5) {
		t.Fatalf("expected %f, got %v", math.Sin(1.5), result.Num)
	}
}
