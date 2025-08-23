package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluateComparisonBinaryExpr(t *testing.T) {
	t.Parallel()

	evaluator := NewEvaluator(io.Discard)

	tests := []struct {
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   datavalue.Value
	}{
		{
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Number(5),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      ">",
					TokenType: token.TokenTypeGreaterThan,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(true),
		},
		{
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Number(5),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      ">=",
					TokenType: token.TokenTypeGreaterThanOrEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(true),
		},
		{
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Number(5),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "<",
					TokenType: token.TokenTypeLessThan,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(true),
		},
		{
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Number(5),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "<=",
					TokenType: token.TokenTypeLessThanOrEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(true),
		},
	}

	for _, test := range tests {
		rawResult, err := evaluator.evaluateComparisonBinaryExpr(
			test.inputLeft,
			test.inputRight,
			test.inputNode,
		)

		if err != nil {
			t.Fatalf("error evaluating %s: %s", test.inputNode.Expr(), err)
		}

		if rawResult.Value.DataType() != test.expected.DataType() {
			t.Fatalf(
				"expected %s, got %s",
				test.expected.DataType().AsString(),
				rawResult.Value.DataType().AsString(),
			)
		}
	}
}

func TestEvaluateComparisonBinaryExprErr(t *testing.T) {
	t.Parallel()

	evaluator := NewEvaluator(io.Discard)

	tests := []struct {
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   datavalue.Value
	}{
		{
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.String("5"),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      ">",
					TokenType: token.TokenTypeGreaterThan,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(true),
		},
		{
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Number(5),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "==",
					TokenType: token.TokenTypeEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(true),
		},
	}

	for _, test := range tests {
		_, err := evaluator.evaluateComparisonBinaryExpr(
			test.inputLeft,
			test.inputRight,
			test.inputNode,
		)

		if err == nil {
			t.Fatalf("expected error evaluating \"%s\", got nil", test.inputNode.Expr())
		}
	}
}
