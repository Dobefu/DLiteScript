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
		name     string
		input    ast.ExprNode
		expected datavalue.Value
	}{
		{
			name: "infinite loop",
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
			name: "loop with range variable",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   0,
					EndPos:     0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "i",
				RangeVariable:    "i",
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
		{
			name: "loop with condition",
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
			name: "loop with continue statement",
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
		{
			name: "loop with break statement count > 1",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{Count: 3, StartPos: 0, EndPos: 0},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "i",
				RangeVariable:    "i",
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
		{
			name: "loop with continue statement count > 1",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ContinueStatement{Count: 2, StartPos: 0, EndPos: 0},
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos:         0,
				EndPos:           0,
				DeclaredVariable: "i",
				RangeVariable:    "i",
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
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "invalid range from",
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
			expected: fmt.Sprintf(
				"%s: %s: %s",
				ErrMsgCouldNotEvaluateForStatement,
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
			),
		},
		{
			name: "invalid range to",
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
			expected: fmt.Sprintf(
				"%s: %s: %s",
				ErrMsgCouldNotEvaluateForStatement,
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
			),
		},
		{
			name: "invalid condition",
			input: &ast.ForStatement{
				Condition: &ast.StringLiteral{
					Value:    "not_a_bool",
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
			expected: fmt.Sprintf(
				"%s: %s: %s",
				ErrMsgCouldNotEvaluateForStatement,
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "string"),
			),
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
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
