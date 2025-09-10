package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateVariableDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.VariableDeclaration
		expected string
	}{
		{
			name: "evaluation error",
			input: &ast.VariableDeclaration{
				Name: "x",
				Type: "int",
				Value: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "bogus",
					Arguments:    []ast.ExprNode{},
					StartPos:     4,
					EndPos:       18,
				},
				StartPos: 0,
				EndPos:   18,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "type mismatch",
			input: &ast.VariableDeclaration{
				Name:     "x",
				Type:     "int",
				Value:    &ast.StringLiteral{Value: "5", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeMismatch, "int", "string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
		})

		_, err := NewEvaluator(io.Discard).evaluateVariableDeclaration(test.input)

		if err == nil {
			t.Fatalf("expected error evaluating %s, got nil", test.input.Expr())
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Fatalf("expected \"%s\", got \"%s\"", test.expected, errors.Unwrap(err).Error())
		}
	}
}
