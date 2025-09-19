package evaluator

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestEvaluateReturnStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		statement *ast.ReturnStatement
		expected  datavalue.Value
	}{
		{
			name: "no values",
			statement: &ast.ReturnStatement{
				Values:    []ast.ExprNode{},
				NumValues: 0,
				StartPos:  0,
				EndPos:    0,
			},
			expected: datavalue.Null(),
		},
		{
			name: "single number",
			statement: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value:    "1",
						StartPos: 0,
						EndPos:   1,
					},
				},
				NumValues: 1,
				StartPos:  0,
				EndPos:    1,
			},
			expected: datavalue.Number(1),
		},
		{
			name: "multiple values",
			statement: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.StringLiteral{
						Value:    "test",
						StartPos: 0,
						EndPos:   4,
					},
					&ast.NumberLiteral{
						Value:    "1",
						StartPos: 5,
						EndPos:   6,
					},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    6,
			},
			expected: datavalue.Tuple(datavalue.String("test"), datavalue.Number(1)),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(nil)
			result, err := evaluator.evaluateReturnStatement(test.statement)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result.Value.DataType != test.expected.DataType {
				t.Errorf(
					"expected %s, got %s",
					test.expected.DataType.AsString(),
					result.Value.DataType.AsString(),
				)
			}

			if result.Value.ToString() != test.expected.ToString() {
				t.Errorf(
					"expected %s, got %s",
					test.expected.ToString(),
					result.Value.ToString(),
				)
			}
		})
	}
}

func TestEvaluateReturnStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		statement *ast.ReturnStatement
		expected  string
	}{
		{
			name: "invalid value",
			statement: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value:    "test",
						StartPos: 0,
						EndPos:   1,
					},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    6,
			},
			expected: "strconv.ParseFloat: parsing \"test\": invalid syntax",
		},
		{
			name: "invalid value in tuple",
			statement: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.StringLiteral{
						Value:    "test",
						StartPos: 0,
						EndPos:   1,
					},
					&ast.NumberLiteral{
						Value:    "test",
						StartPos: 0,
						EndPos:   1,
					},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    1,
			},
			expected: "strconv.ParseFloat: parsing \"test\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(nil)
			_, err := evaluator.evaluateReturnStatement(test.statement)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected %s, got %s", test.expected, err.Error())
			}
		})
	}
}
