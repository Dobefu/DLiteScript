package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func TestGetCosFunction(t *testing.T) {
	t.Parallel()

	functions := GetMathFunctions()

	if _, ok := functions["cos"]; !ok {
		t.Fatalf("expected cos function, got %v", functions)
	}

	cosFunc := functions["cos"]

	if cosFunc.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", cosFunc.FunctionType)
	}

	if cosFunc.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", cosFunc.ArgKinds[0])
	}

	result, err := cosFunc.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != math.Cos(1.5) {
		t.Fatalf("expected %f, got %v", math.Cos(1.5), result.Num)
	}
}
