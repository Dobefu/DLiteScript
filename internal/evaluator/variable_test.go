package evaluator

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestVariable(t *testing.T) {
	t.Parallel()

	variable := Variable{
		Value: datavalue.Number(1),
		Type:  "number",
	}

	if variable.GetValue() != datavalue.Number(1) {
		t.Errorf("expected 1, got %v", variable.GetValue())
	}

	if variable.GetType() != "number" {
		t.Errorf("expected number, got %s", variable.GetType())
	}

	if variable.IsConstant() {
		t.Errorf("expected not constant, got %v", variable)
	}
}
