package evaluator

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestEvaluateSpreadExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		statement *ast.SpreadExpr
		expected  datavalue.Value
	}{
		{
			name: "single number",
			statement: &ast.SpreadExpr{
				Expression: &ast.NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			expected: datavalue.Number(1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(nil)
			result, err := evaluator.evaluateSpreadExpr(test.statement)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result.Value.DataType() != test.expected.DataType() {
				t.Errorf("expected %s, got %s", test.expected.DataType().AsString(), result.Value.DataType().AsString())
			}

			if result.Value.ToString() != test.expected.ToString() {
				t.Errorf("expected %s, got %s", test.expected.ToString(), result.Value.ToString())
			}
		})
	}
}
