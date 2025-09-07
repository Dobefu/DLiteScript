package evaluator

import (
	"errors"
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected float64
	}{
		{
			name:     "PI",
			input:    &ast.Identifier{Value: "PI", StartPos: 0, EndPos: 1},
			expected: math.Pi,
		},
		{
			name: "variable declaration",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name: "test",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   0,
						},
						StartPos: 0,
						EndPos:   0,
					},
					&ast.Identifier{Value: "test", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: 1,
		},
		{
			name: "scoped identifier",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name: "test",
						Type: "number",
						Value: &ast.NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
						StartPos: 0,
						EndPos:   0,
					},
					&ast.Identifier{Value: "test", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rawResult, err := NewEvaluator(io.Discard).Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating \"%s\": %s", test.input.Expr(), err.Error())
			}

			result, err := rawResult.Value.AsNumber()

			if err != nil {
				t.Fatalf("expected number, got type error: %s", err.Error())
			}

			if result != test.expected {
				t.Errorf("expected %f, got %f", test.expected, result)
			}
		})
	}
}

func TestEvaluateIdentifierErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name:     "undefined identifier",
			input:    &ast.Identifier{Value: "bogus", StartPos: 0, EndPos: 1},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).Evaluate(test.input)

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
