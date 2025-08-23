package evaluator

import (
	"io"
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
				Left: &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.BinaryExpr{
					Left: &ast.NumberLiteral{Value: "5", StartPos: 2, EndPos: 3},
					Right: &ast.PrefixExpr{
						Operator: token.Token{
							Atom:      "-",
							TokenType: token.TokenTypeOperationSub,
							StartPos:  0,
							EndPos:    0,
						},
						Operand: &ast.FunctionCall{
							FunctionName: "abs",
							Arguments: []ast.ExprNode{
								&ast.Identifier{Value: "PI", StartPos: 4, EndPos: 5},
							},
							StartPos: 0,
							EndPos:   0,
						},
						StartPos: 0,
						EndPos:   0,
					},
					Operator: token.Token{
						Atom:      "+",
						TokenType: token.TokenTypeOperationAdd,
						StartPos:  0,
						EndPos:    0,
					},
					StartPos: 0,
					EndPos:   0,
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: 5 + math.Abs(-5+math.Pi),
		},
	}

	for _, test := range tests {
		rawResult, err := NewEvaluator(io.Discard).Evaluate(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input.Expr(), err.Error())
		}

		result, err := rawResult.Value.AsNumber()

		if err != nil {
			t.Fatalf("expected number, got type error: %s", err.Error())
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestOutput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected string
	}{
		{
			input: &ast.FunctionCall{
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{Value: "test", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.input.Expr(), func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(io.Discard)

			_, err := evaluator.Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating '%s': %s", test.input.Expr(), err.Error())
			}

			if evaluator.Output() != test.expected {
				t.Errorf("expected '%v', got '%v'", test.expected, evaluator.Output())
			}
		})
	}
}

func BenchmarkEvaluate(b *testing.B) {
	for b.Loop() {
		_, _ = NewEvaluator(io.Discard).Evaluate(
			&ast.BinaryExpr{
				Left: &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right: &ast.BinaryExpr{
					Left: &ast.PrefixExpr{
						Operator: *token.NewToken("-", token.TokenTypeOperationSub, 0, 0),
						Operand:  &ast.NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
						StartPos: 0,
						EndPos:   0,
					},
					Right: &ast.BinaryExpr{
						Left:     &ast.NumberLiteral{Value: "3", StartPos: 4, EndPos: 5},
						Right:    &ast.NumberLiteral{Value: "4", StartPos: 6, EndPos: 7},
						Operator: *token.NewToken("/", token.TokenTypeOperationDiv, 0, 0),
						StartPos: 0,
						EndPos:   0,
					},
					Operator: *token.NewToken("*", token.TokenTypeOperationMul, 0, 0),
					StartPos: 0,
					EndPos:   0,
				},
				Operator: *token.NewToken("+", token.TokenTypeOperationAdd, 0, 0),
				StartPos: 0,
				EndPos:   0,
			},
		)
	}
}
