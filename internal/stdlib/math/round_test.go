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

	round := functions["round"]

	if round.FunctionType != function.FunctionTypeFixed {
		t.Fatalf("expected fixed function, got %v", round.FunctionType)
	}

	if round.ArgKinds[0] != datatype.DataTypeNumber {
		t.Fatalf("expected number argument, got %v", round.ArgKinds[0])
	}

	result, err := round.Handler(nil, []datavalue.Value{datavalue.Number(1.5)})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Num != 2 {
		t.Fatalf("expected 2, got %v", result.Num)
	}
}
