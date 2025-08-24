package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateIfStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected datavalue.Value
	}{
		{
			name: "false condition",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value:    "false",
					StartPos: 0,
					EndPos:   0,
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   0,
						},
					},
					StartPos: 0,
					EndPos:   0,
				},
				ElseBlock: nil,
				StartPos:  0,
				EndPos:    0,
			},
			expected: datavalue.Null(),
		},
		{
			name: "true condition",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value:    "true",
					StartPos: 0,
					EndPos:   0,
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   0,
						},
					},
					StartPos: 0,
					EndPos:   0,
				},
				ElseBlock: nil,
				StartPos:  0,
				EndPos:    0,
			},
			expected: datavalue.Number(1),
		},
		{
			name: "false condition, else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value:    "false",
					StartPos: 0,
					EndPos:   0,
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   0,
					EndPos:     0,
				},
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value:    "2",
							StartPos: 0,
							EndPos:   0,
						},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: datavalue.Number(2),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := NewEvaluator(io.Discard).Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating if statement: \"%s\"", err.Error())
			}

			if result.Value.DataType().AsString() != test.expected.DataType().AsString() {
				t.Fatalf(
					"expected \"%v\", got \"%v\" at position %d",
					test.expected.DataType().AsString(),
					result.Value.DataType().AsString(),
					test.input.StartPosition(),
				)
			}
		})
	}
}

func TestEvaluateIfStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "condition is number",
			input: &ast.IfStatement{
				Condition: &ast.NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   0,
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   0,
						},
					},
					StartPos: 0,
					EndPos:   0,
				},
				ElseBlock: nil,
				StartPos:  0,
				EndPos:    0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "number"),
		},
		{
			name: "condition is undefined identifier",
			input: &ast.IfStatement{
				Condition: &ast.Identifier{
					Value:    "undefined_var",
					StartPos: 0,
					EndPos:   0,
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   0,
						},
					},
					StartPos: 0,
					EndPos:   0,
				},
				ElseBlock: nil,
				StartPos:  0,
				EndPos:    0,
			},
			expected: "undefined identifier: 'undefined_var' at position 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%v\", got \"%v\"", test.expected, err.Error())
			}
		})
	}
}
