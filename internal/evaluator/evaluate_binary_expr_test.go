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

func TestEvaluateBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.BinaryExpr
		expected datavalue.Value
	}{
		{
			name: "addition",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Number(10),
		},
		{
			name: "subtraction",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Number(0),
		},
		{
			name: "multiplication",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Number(25),
		},
		{
			name: "division",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Number(1),
		},
		{
			name: "modulo",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Number(0),
		},
		{
			name: "power",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "^",
					TokenType: token.TokenTypeOperationPow,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Number(3125),
		},
		{
			name: "equality",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "==",
					TokenType: token.TokenTypeEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Bool(true),
		},
		{
			name: "greater than or equal",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "6", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      ">=",
					TokenType: token.TokenTypeGreaterThanOrEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Bool(false),
		},
		{
			name: "logical and",
			input: &ast.BinaryExpr{
				Left:  &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				Right: &ast.BoolLiteral{Value: "true", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "&&",
					TokenType: token.TokenTypeLogicalAnd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: datavalue.Bool(true),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rawResult, err := NewEvaluator(io.Discard).evaluateBinaryExpr(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err.Error())
			}

			if rawResult.Value.DataType() != test.expected.DataType() {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.DataType().AsString(),
					rawResult.Value.DataType().AsString(),
				)
			}
		})
	}
}

func TestEvaluateBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.BinaryExpr
		expected string
	}{
		{
			name: "left operand is nil",
			input: &ast.BinaryExpr{
				Left:  nil,
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
		},
		{
			name: "right operand is nil",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: nil,
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
		},
		{
			name: "division by zero",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "0", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: errorutil.ErrorMsgDivByZero,
		},
		{
			name: "modulo by zero",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "0", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: errorutil.ErrorMsgModByZero,
		},
		{
			name: "unknown operator",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "0", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, ","),
		},
		{
			name: "undefined identifier",
			input: &ast.BinaryExpr{
				Left:  &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "0", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "x"),
		},
		{
			name: "undefined identifier",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "0", StartPos: 2, EndPos: 3},
				Right: &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 1,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "x"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateBinaryExpr(test.input)

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
