package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluatePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.PrefixExpr
		expected float64
	}{
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
				Operand: &ast.NumberLiteral{Value: "5", Pos: 1},
				Pos:     0,
			},
			expected: -5,
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Operand: &ast.NumberLiteral{Value: "5", Pos: 1},
				Pos:     0,
			},
			expected: 5,
		},
	}

	for _, test := range tests {
		rawResult, err := NewEvaluator(io.Discard).evaluatePrefixExpr(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input.Expr(), err.Error())
		}

		result, err := rawResult.Value.AsNumber()

		if err != nil {
			t.Fatalf("expected number, got type error: %s", err.Error())
		}

		if result != test.expected {
			t.Errorf("expected '%f', got '%f'", test.expected, result)
		}
	}
}

func TestEvaluatePrefixExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.PrefixExpr
		expected string
	}{
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
				},
				Operand: nil,
				Pos:     0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Operand: &ast.StringLiteral{Value: "test", Pos: 1},
				Pos:     0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "string"),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator(io.Discard).evaluatePrefixExpr(test.input)

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
	}
}
