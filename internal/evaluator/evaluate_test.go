package evaluator

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected float64
	}{
		{
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{Value: "5", Pos: 0},
				Right: &ast.BinaryExpr{
					Left: &ast.NumberLiteral{Value: "5", Pos: 2},
					Right: &ast.PrefixExpr{
						Operator: token.Token{
							Atom:      "-",
							TokenType: token.TokenTypeOperationSub,
						},
						Operand: &ast.FunctionCall{
							FunctionName: "abs",
							Arguments: []ast.ExprNode{
								&ast.Identifier{Value: "PI", Pos: 4},
							},
							Pos: 0,
						},
						Pos: 0,
					},
					Operator: token.Token{
						Atom:      "+",
						TokenType: token.TokenTypeOperationAdd,
					},
					Pos: 0,
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				Pos: 0,
			},
			expected: 5 + math.Abs(-5+math.Pi),
		},
	}

	for _, test := range tests {
		rawResult, err := NewEvaluator().Evaluate(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input.Expr(), err.Error())
		}

		result, err := rawResult.AsNumber()

		if err != nil {
			t.Fatalf("expected number, got type error: %s", err.Error())
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func BenchmarkEvaluate(b *testing.B) {
	for b.Loop() {
		_, _ = NewEvaluator().Evaluate(
			&ast.BinaryExpr{
				Left: &ast.NumberLiteral{Value: "1", Pos: 0},
				Right: &ast.BinaryExpr{
					Left: &ast.PrefixExpr{
						Operator: *token.NewToken("-", token.TokenTypeOperationSub),
						Operand:  &ast.NumberLiteral{Value: "2", Pos: 2},
						Pos:      0,
					},
					Right: &ast.BinaryExpr{
						Left:     &ast.NumberLiteral{Value: "3", Pos: 4},
						Right:    &ast.NumberLiteral{Value: "4", Pos: 6},
						Operator: *token.NewToken("/", token.TokenTypeOperationDiv),
						Pos:      0,
					},
					Operator: *token.NewToken("*", token.TokenTypeOperationMul),
					Pos:      0,
				},
				Operator: *token.NewToken("+", token.TokenTypeOperationAdd),
				Pos:      0,
			},
		)
	}
}
