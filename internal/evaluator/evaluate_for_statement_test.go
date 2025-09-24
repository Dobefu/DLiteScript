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
						&ast.BreakStatement{
							Count: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "",
				RangeVariable:    "",
				RangeFrom:        nil,
				RangeTo:          nil,
				IsRange:          false,
				HasExplicitFrom:  false,
			},
			expected: datavalue.Null(),
		},
		{
			name: "loop with range variable",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "i",
				RangeVariable:    "i",
				RangeFrom: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				RangeTo: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				IsRange:         true,
				HasExplicitFrom: false,
			},
			expected: datavalue.Null(),
		},
		{
			name: "loop with condition",
			input: &ast.ForStatement{
				Condition: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{
							Count: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "",
				RangeVariable:    "",
				RangeFrom:        nil,
				RangeTo:          nil,
				IsRange:          false,
				HasExplicitFrom:  false,
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
							Count: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "i",
				RangeVariable:    "",
				RangeFrom: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				RangeTo: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				IsRange:         true,
				HasExplicitFrom: false,
			},
			expected: datavalue.Null(),
		},
		{
			name: "loop with break statement count > 1",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{
							Count: 3,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "i",
				RangeVariable:    "i",
				RangeFrom: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				RangeTo: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				IsRange:         true,
				HasExplicitFrom: false,
			},
			expected: datavalue.Null(),
		},
		{
			name: "loop with continue statement count > 1",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ContinueStatement{
							Count: 2,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "i",
				RangeVariable:    "i",
				RangeFrom: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				RangeTo: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				IsRange:         true,
				HasExplicitFrom: false,
			},
			expected: datavalue.Null(),
		},
	}

	for _, test := range tests {
		result, err := NewEvaluator(io.Discard).Evaluate(test.input)

		if err != nil {
			t.Fatalf("error evaluating for statement: \"%s\"", err.Error())
		}

		if result.Value.DataType.AsString() != test.expected.DataType.AsString() {
			t.Fatalf(
				"expected \"%v\", got \"%v\" at position %d",
				test.expected.DataType.AsString(),
				result.Value.DataType.AsString(),
				test.input.GetRange().Start.Offset,
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
						&ast.BreakStatement{
							Count: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "i",
				RangeVariable:    "",
				RangeFrom:        ast.ExprNode(nil),
				RangeTo:          nil,
				IsRange:          true,
				HasExplicitFrom:  false,
			},
			expected: fmt.Sprintf(
				"%s: %s",
				fmt.Sprintf(
					ErrMsgCouldNotEvaluateForStatement,
					errorutil.StageEvaluate.String(),
				),
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
							Value: "0",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.ForStatement{
						Condition: nil,
						Body: &ast.BlockStatement{
							Statements: []ast.ExprNode{
								&ast.BreakStatement{
									Count: 1,
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 0, Line: 0, Column: 0},
									},
								},
							},
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
						DeclaredVariable: "i",
						RangeVariable:    "",
						RangeFrom: &ast.NumberLiteral{
							Value: "0",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
						RangeTo:         nil,
						IsRange:         true,
						HasExplicitFrom: false,
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s",
				fmt.Sprintf(
					ErrMsgCouldNotEvaluateForStatement,
					errorutil.StageEvaluate.String(),
				),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
			),
		},
		{
			name: "invalid condition",
			input: &ast.ForStatement{
				Condition: &ast.StringLiteral{
					Value: "not_a_bool",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.BreakStatement{
							Count: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				DeclaredVariable: "",
				RangeVariable:    "",
				RangeFrom:        nil,
				RangeTo:          nil,
				IsRange:          false,
				HasExplicitFrom:  false,
			},
			expected: fmt.Sprintf(
				"%s: %s",
				fmt.Sprintf(
					ErrMsgCouldNotEvaluateForStatement,
					errorutil.StageEvaluate.String(),
				),
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
