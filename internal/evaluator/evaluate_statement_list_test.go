package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateStatementList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.StatementList
		expected string
	}{
		{
			name:     "empty",
			input:    &ast.StatementList{Statements: []ast.ExprNode{}, StartPos: 0, EndPos: 0},
			expected: "",
		},
		{
			name: "single statement",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{
						Value:    "1",
						StartPos: 0,
						EndPos:   1,
					},
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
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
		expected string
	}{
		{
			name: "undefined function",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.FunctionCall{
						FunctionName: "bogus",
						Arguments: []ast.ExprNode{
							&ast.NumberLiteral{
								Value:    "1",
								StartPos: 0,
								EndPos:   1,
							},
						},
						StartPos: 0,
						EndPos:   0,
					},
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.evaluateStatementList(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}
