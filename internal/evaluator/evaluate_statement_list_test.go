package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateStatementList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   *ast.StatementList
		outFile io.Writer
	}{
		{
			name: "empty",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			outFile: io.Discard,
		},
		{
			name: "single statement",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			outFile: io.Discard,
		},
		{
			name: "output to writer",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.FunctionCall{
						Namespace:    "",
						FunctionName: "printf",
						Arguments: []ast.ExprNode{
							&ast.StringLiteral{
								Value: "test\n",
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 5, Line: 0, Column: 0},
								},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 5, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			outFile: &discardWriter{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(test.outFile)
			_, err := ev.evaluateStatementList(test.input)

			if err != nil {
				t.Errorf("error evaluating '%s': %s", test.input.Expr(), err.Error())
			}
		})
	}
}

func TestEvaluateStatementListErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.StatementList
		outFile  io.Writer
		expected string
	}{
		{
			name: "undefined function",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.FunctionCall{
						Namespace:    "",
						FunctionName: "bogus",
						Arguments: []ast.ExprNode{
							&ast.NumberLiteral{
								Value: "1",
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 1, Line: 0, Column: 0},
								},
							},
						},
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
			outFile: io.Discard,
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
			),
		},
		{
			name: "write error",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.FunctionCall{
						Namespace:    "",
						FunctionName: "printf",
						Arguments: []ast.ExprNode{
							&ast.StringLiteral{
								Value: "test\n",
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 5, Line: 0, Column: 0},
								},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 5, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			outFile:  &errWriter{},
			expected: "write error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(test.outFile)
			_, err := ev.evaluateStatementList(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
