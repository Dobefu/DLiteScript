package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateAssignmentStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected any
	}{
		{
			name: "assignment statement",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name: "x",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value: "0",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.AssignmentStatement{
						Left: &ast.Identifier{
							Value: "x",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
						Right: &ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 4, Line: 0, Column: 0},
								End:   ast.Position{Offset: 7, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 7, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 7, Line: 0, Column: 0},
				},
			},
			expected: 1.0,
		},
		{
			name: "assignment statement in block scope",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name: "x",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value: "0",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.BlockStatement{
						Statements: []ast.ExprNode{
							&ast.VariableDeclaration{
								Name: "x",
								Type: "number",
								Value: &ast.NumberLiteral{
									Value: "5",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 1, Line: 0, Column: 0},
								},
							},
							&ast.AssignmentStatement{
								Left: &ast.Identifier{
									Value: "x",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
								Right: &ast.NumberLiteral{
									Value: "42",
									Range: ast.Range{
										Start: ast.Position{Offset: 4, Line: 0, Column: 0},
										End:   ast.Position{Offset: 6, Line: 0, Column: 0},
									},
								},
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 6, Line: 0, Column: 0},
								},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 6, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 6, Line: 0, Column: 0},
				},
			},
			expected: 42.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			eval := NewEvaluator(io.Discard)

			result, err := eval.Evaluate(test.input)

			if err != nil {
				t.Fatalf("failed to evaluate assignment statement: %s", err.Error())
			}

			if result.Value.DataType != datatype.DataTypeNumber {
				t.Fatalf("expected number result, got %v", result.Value.DataType)
			}

			num, err := result.Value.AsNumber()

			if err != nil {
				t.Fatalf("expected number result: %s", err.Error())
			}

			if num != test.expected {
				t.Errorf("expected %f, got %f", test.expected, num)
			}
		})
	}
}

func TestEvaluateAssignmentStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "assignment to undefined variable",
			input: &ast.AssignmentStatement{
				Left: &ast.Identifier{
					Value: "undefined_var",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 13, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 16, Line: 0, Column: 0},
						End:   ast.Position{Offset: 17, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 17, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "undefined_var"),
			),
		},
		{
			name: "assignment to constant",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.ConstantDeclaration{
						Name: "const_var",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value: "5",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.AssignmentStatement{
						Left: &ast.Identifier{
							Value: "const_var",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 9, Line: 0, Column: 0},
							},
						},
						Right: &ast.NumberLiteral{
							Value: "10",
							Range: ast.Range{
								Start: ast.Position{Offset: 12, Line: 0, Column: 0},
								End:   ast.Position{Offset: 14, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 14, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 14, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgReassignmentToConstant, "const_var"),
			),
		},
		{
			name: "assignment to constant in block scope",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.ConstantDeclaration{
						Name: "block_const",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value: "5",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.AssignmentStatement{
						Left: &ast.Identifier{
							Value: "block_const",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 11, Line: 0, Column: 0},
							},
						},
						Right: &ast.NumberLiteral{
							Value: "10",
							Range: ast.Range{
								Start: ast.Position{Offset: 14, Line: 0, Column: 0},
								End:   ast.Position{Offset: 16, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 16, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 16, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgReassignmentToConstant, "block_const"),
			),
		},
		{
			name: "assignment with right side evaluation error",
			input: &ast.AssignmentStatement{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "undefined_func",
					Arguments:    []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 18, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 18, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "undefined_func"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			eval := NewEvaluator(io.Discard)

			_, err := eval.Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected error \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
