package evaluator

import (
	"io"
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	funcDeclaration := &ast.FuncDeclarationStatement{
		Name: "someFunction",
		Args: []ast.FuncParameter{},
		Body: &ast.BlockStatement{
			Statements: []ast.ExprNode{},
			StartPos:   0,
			EndPos:     1,
		},
		ReturnValues:    []string{},
		NumReturnValues: 0,
		StartPos:        0,
		EndPos:          1,
	}

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected *controlflow.EvaluationResult
	}{
		{
			name: "binary expression",
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
							Namespace:    "math",
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
			expected: controlflow.NewRegularResult(
				datavalue.Number(5 + math.Abs(-5+math.Pi)),
			),
		},
		{
			name: "index assignment statement",
			input: &ast.IndexAssignmentStatement{
				Array:    &ast.Identifier{Value: "someArray", StartPos: 0, EndPos: 1},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(
				datavalue.Array(datavalue.Number(1)),
			),
		},
		{
			name:  "function declaration statement",
			input: funcDeclaration,
			expected: controlflow.NewRegularResult(
				datavalue.Function(funcDeclaration),
			),
		},
		{
			name: "spread expression",
			input: &ast.SpreadExpr{
				Expression: &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos:   0,
				EndPos:     1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(1)),
		},
		{
			name: "index expression",
			input: &ast.IndexExpr{
				Array:    &ast.Identifier{Value: "someArray", StartPos: 0, EndPos: 1},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(0)),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			ev.outerScope["someArray"] = &Variable{
				Value: datavalue.Array(datavalue.Number(0)),
				Type:  "array",
			}

			result, err := ev.Evaluate(test.input)

			if err != nil {
				t.Errorf("error evaluating \"%s\": %s", test.input.Expr(), err.Error())
			}

			if result.Value.DataType != test.expected.Value.DataType {
				t.Errorf(
					"expected \"%T\", got \"%T\"",
					test.expected.Value.DataType,
					result.Value.DataType,
				)
			}

			if !result.Value.Equals(test.expected.Value) {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected.Value.ToString(),
					result.Value.ToString(),
				)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "function call",
			input: &ast.FunctionCall{
				Namespace:    "",
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(io.Discard)

			_, err := evaluator.Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating \"%s\": %s", test.input.Expr(), err.Error())
			}

			if evaluator.Output() != test.expected {
				t.Errorf("expected \"%v\", got \"%v\"", test.expected, evaluator.Output())
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
