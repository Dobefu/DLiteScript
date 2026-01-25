package evaluator

import (
	"fmt"
	"io"
	"strings"
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
		expected error
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
			expected: errCouldNotEvaluateForStatement,
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
			expected: errCouldNotEvaluateForStatement,
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
			expected: errCouldNotEvaluateForStatement,
		},
		{
			name: "declareLoopVariable error with invalid range from",
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
				RangeFrom: &ast.Identifier{
					Value: "bogus",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				RangeTo: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				IsRange:         true,
				HasExplicitFrom: false,
			},
			expected: fmt.Errorf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"), // nolint: err113
		},
		{
			name: "executeForIteration error with invalid body",
			input: &ast.ForStatement{
				Condition: nil,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.Identifier{
							Value: "bogus",
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
			expected: fmt.Errorf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"), // nolint: err113
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !strings.Contains(err.Error(), test.expected.Error()) {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.Error(),
					err.Error(),
				)
			}
		})
	}
}

func TestDeclareLoopVariable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		node     *ast.ForStatement
		expected error
	}{
		{
			name: "declareLoopVariable with outerScope assignment",
			node: &ast.ForStatement{
				DeclaredVariable: "outerVar",
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
				Condition:       nil,
				Body:            nil,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(io.Discard)
			err := evaluator.declareLoopVariable(test.node)

			if err != nil && test.expected == nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if err == nil && test.expected != nil {
				t.Fatalf("expected error, got nil")
			}

			_, hasScopedValue := evaluator.outerScope[test.node.DeclaredVariable]

			if !hasScopedValue {
				t.Fatalf(
					"variable \"%s\" not found in outerScope",
					test.node.DeclaredVariable,
				)
			}
		})
	}
}
