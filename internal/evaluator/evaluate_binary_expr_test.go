package evaluator

import (
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
			name: "number addition",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Number(10),
		},
		{
			name: "number subtraction",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Number(0),
		},
		{
			name: "number multiplication",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Number(25),
		},
		{
			name: "number division",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Number(1),
		},
		{
			name: "number modulo",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Number(0),
		},
		{
			name: "number power",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "^",
					TokenType: token.TokenTypeOperationPow,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Number(3125),
		},
		{
			name: "number equality",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "==",
					TokenType: token.TokenTypeEqual,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Bool(true),
		},
		{
			name: "number greater than or equal",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "6",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      ">=",
					TokenType: token.TokenTypeGreaterThanOrEqual,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Bool(false),
		},
		{
			name: "boolean logical and",
			input: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "&&",
					TokenType: token.TokenTypeLogicalAnd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Bool(true),
		},
		{
			name: "array addition",
			input: &ast.BinaryExpr{
				Left: &ast.ArrayLiteral{
					Values: []ast.ExprNode{&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					}},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.ArrayLiteral{
					Values: []ast.ExprNode{&ast.NumberLiteral{
						Value: "2",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					}},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Array(
				datavalue.Number(1),
				datavalue.Number(2),
			),
		},
		{
			name: "string addition",
			input: &ast.BinaryExpr{
				Left: &ast.StringLiteral{
					Value: "test",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.StringLiteral{
					Value: "test",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.String("testtest"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rawResult, err := NewEvaluator(io.Discard).evaluateBinaryExpr(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err.Error())
			}

			if rawResult.Value.DataType != test.expected.DataType {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.DataType.AsString(),
					rawResult.Value.DataType.AsString(),
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
				Left: nil,
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
			),
		},
		{
			name: "right operand is nil",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: nil,
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "null", "number"),
			),
		},
		{
			name: "division by zero",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "/",
					TokenType: token.TokenTypeOperationDiv,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				errorutil.ErrorMsgDivByZero,
			),
		},
		{
			name: "modulo by zero",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "%",
					TokenType: token.TokenTypeOperationMod,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				errorutil.ErrorMsgModByZero,
			),
		},
		{
			name: "unknown operator",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, ","),
			),
		},
		{
			name: "undefined identifier",
			input: &ast.BinaryExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "x"),
			),
		},
		{
			name: "undefined identifier",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Right: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "x"),
			),
		},
		{
			name: "boolean addition",
			input: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgCannotConcat, "bool", "bool"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateBinaryExpr(test.input)

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

func TestEvaluateArithmeticBinaryExprNumberErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		left     datavalue.Value
		right    datavalue.Value
		node     *ast.BinaryExpr
		expected string
	}{
		{
			name:  "null values",
			left:  datavalue.Null(),
			right: datavalue.Null(),
			node: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"could not get binary expr value as number: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "null"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateArithmeticBinaryExprNumber(
				test.left,
				test.right,
				test.node,
			)

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

func TestEvaluateArithmeticBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		node     *ast.BinaryExpr
		expected string
	}{
		{
			name: "addition",
			node: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "",
					TokenType: token.Type(-1),
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, ""),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateArithmeticBinaryExpr(
				datavalue.Number(5),
				datavalue.Number(5),
				test.node,
			)

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

func TestEvaluateArithmeticBinaryExprArrayErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		left     datavalue.Value
		right    datavalue.Value
		node     *ast.BinaryExpr
		expected string
	}{
		{
			name:  "invalid left operand",
			left:  datavalue.Number(5),
			right: datavalue.Array(datavalue.Number(1)),
			node: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 2, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "array", "number"),
			),
		},
		{
			name:  "invalid right operand",
			left:  datavalue.Array(datavalue.Number(1)),
			right: datavalue.Number(5),
			node: &ast.BinaryExpr{
				Left: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 2, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "array", "number"),
			),
		},
		{
			name:  "invalid operator",
			left:  datavalue.Array(datavalue.Number(1)),
			right: datavalue.Array(datavalue.Number(1)),
			node: &ast.BinaryExpr{
				Left: &ast.ArrayLiteral{
					Values: []ast.ExprNode{&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					}},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "",
					TokenType: token.Type(-1),
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, ""),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateArithmeticBinaryExprArray(
				test.left,
				test.right,
				test.node,
			)

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

func TestEvaluateArithmeticBinaryExprStringErr(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		left     datavalue.Value
		right    datavalue.Value
		node     *ast.BinaryExpr
		expected string
	}{
		{
			name:  "invalid left operand",
			left:  datavalue.Number(5),
			right: datavalue.String("5"),
			node: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.StringLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "string", "number"),
		},
		{
			name:  "invalid operator",
			left:  datavalue.String("5"),
			right: datavalue.String("5"),
			node: &ast.BinaryExpr{
				Left: &ast.StringLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "",
					TokenType: token.Type(-1),
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, ""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateArithmeticBinaryExprString(
				test.left,
				test.right,
				test.node,
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}

}

func TestGetBinaryExprValueAsBoolErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		left     datavalue.Value
		right    datavalue.Value
		expected string
	}{
		{
			name:  "invalid left operand",
			left:  datavalue.String("5"),
			right: datavalue.Bool(true),
			expected: fmt.Sprintf(
				"could not get binary expr value as bool: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "string"),
			),
		},
		{
			name:  "invalid right operand",
			left:  datavalue.Bool(true),
			right: datavalue.String("5"),
			expected: fmt.Sprintf(
				"could not get binary expr value as bool: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "string"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, _, err := NewEvaluator(io.Discard).getBinaryExprValueAsBool(
				test.left,
				test.right,
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}

func TestGetBinaryExprValueAsNumberErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		left     datavalue.Value
		right    datavalue.Value
		expected string
	}{
		{
			name:  "invalid left operand",
			left:  datavalue.String("5"),
			right: datavalue.Number(5),
			expected: fmt.Sprintf(
				"could not get binary expr value as number: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "string"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, _, err := NewEvaluator(io.Discard).getBinaryExprValueAsNumber(
				test.left,
				test.right,
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}

func TestGetBinaryExprValueAsStringErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		left     datavalue.Value
		right    datavalue.Value
		expected string
	}{
		{
			name:  "invalid left operand",
			left:  datavalue.Number(5),
			right: datavalue.String("5"),
			expected: fmt.Sprintf(
				"could not get binary expr value as string: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "string", "number"),
			),
		},
		{
			name:  "invalid right operand",
			left:  datavalue.String("5"),
			right: datavalue.Number(5),
			expected: fmt.Sprintf(
				"could not get binary expr value as string: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "string", "number"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, _, err := NewEvaluator(io.Discard).getBinaryExprValueAsString(
				test.left,
				test.right,
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
