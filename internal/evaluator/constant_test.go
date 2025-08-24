package evaluator

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestConstant(t *testing.T) {
	t.Parallel()

	constant := Constant{
		Value: datavalue.Number(1),
		Type:  "number",
	}

	if !constant.GetValue().Equals(datavalue.Number(1)) {
		t.Errorf("expected 1, got %v", constant.GetValue())
	}

	if constant.GetType() != "number" {
		t.Errorf("expected number, got %s", constant.GetType())
	}

	if !constant.IsConstant() {
		t.Errorf("expected constant, got %v", constant)
	}
}
