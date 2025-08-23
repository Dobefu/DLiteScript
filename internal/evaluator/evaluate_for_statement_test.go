package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateForStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected datavalue.Value
	}{
		{
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{Count: 1, StartPos: 0, EndPos: 0},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "",
				RangeVariable:    "",
				RangeFrom:        nil,
				RangeTo:          nil,
				IsRange:          false,
			},
			expected: datavalue.Null(),
		},
		{
			input: &ast.ForStatement{
				Condition: &ast.BoolLiteral{
					Value:    "true",
					StartPos: 0,
					EndPos:   0,
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{Count: 1, StartPos: 0, EndPos: 0},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "",
				RangeVariable:    "",
				RangeFrom:        nil,
				RangeTo:          nil,
				IsRange:          false,
			},
			expected: datavalue.Null(),
		},
		{
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ContinueStatement{
							Count:    1,
							StartPos: 0,
							EndPos:   0,
						},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "i",
				RangeVariable:    "",
				RangeFrom: &ast.NumberLiteral{
					Value:    "0",
					StartPos: 0,
					EndPos:   0,
				},
				RangeTo: &ast.NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   0,
				},
				IsRange: true,
			},
			expected: datavalue.Null(),
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator(io.Discard).Evaluate(test.input)

		if err != nil {
			t.Fatalf("error evaluating for statement: \"%s\"", err.Error())
		}

		if result.Value.DataType().AsString() != test.expected.DataType().AsString() {
			t.Fatalf(
				"expected \"%v\", got \"%v\" at position %d",
				test.expected.DataType().AsString(),
				result.Value.DataType().AsString(),
				test.input.StartPosition(),
			)
		}
	}
}

func TestEvaluateForStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected string
	}{
		{
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{Count: 1, StartPos: 0, EndPos: 0},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "i",
				RangeVariable:    "",
				RangeFrom:        ast.ExprNode(nil),
				RangeTo:          nil,
				IsRange:          true,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
		},
		{
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name: "i",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value:    "0",
							StartPos: 0,
							EndPos:   0,
						},
						StartPos: 0,
						EndPos:   0,
					},
					&ast.ForStatement{
						Condition: nil,
						Body: &ast.BlockStatement{
							Statements: []ast.ExprNode{
								&ast.BreakStatement{Count: 1, StartPos: 0, EndPos: 0},
							},
							StartPos: 0,
							EndPos:   0,
						},
						StartPos:         0,
						EndPos:           0,
						DeclaredVariable: "i",
						RangeVariable:    "",
						RangeFrom: &ast.NumberLiteral{
							Value:    "0",
							StartPos: 0,
							EndPos:   0,
						},
						RangeTo: nil,
						IsRange: true,
					},
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator(io.Discard).Evaluate(test.input)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != test.expected {
			t.Fatalf(
				"expected \"%v\", got \"%v\" at position %d",
				test.expected,
				err.Error(),
				test.input.StartPosition(),
			)
		}
	}
}
