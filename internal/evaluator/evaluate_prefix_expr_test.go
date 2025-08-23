package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluatePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.PrefixExpr
		expected datavalue.Value
	}{
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &ast.NumberLiteral{Value: "5", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
			},
			expected: datavalue.Number(-5),
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &ast.NumberLiteral{Value: "5", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
			},
			expected: datavalue.Number(5),
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "!",
					TokenType: token.TokenTypeNot,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &ast.BoolLiteral{Value: "true", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
			},
			expected: datavalue.Bool(false),
		},
	}

	for _, test := range tests {
		rawResult, err := NewEvaluator(io.Discard).evaluatePrefixExpr(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input.Expr(), err.Error())
		}

		if rawResult.Value.DataType().AsString() != test.expected.DataType().AsString() {
			t.Errorf(
				"expected '%v', got '%v'",
				test.expected.DataType().AsString(),
				rawResult.Value.DataType().AsString(),
			)
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
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  nil,
				StartPos: 0,
				EndPos:   0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &ast.StringLiteral{Value: "test", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "string"),
		},
		{
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "!",
					TokenType: token.TokenTypeNot,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  nil,
				StartPos: 0,
				EndPos:   0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "null"),
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
