package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluateEqualityBinaryExpr(t *testing.T) {
	t.Parallel()

	function := &ast.FuncDeclarationStatement{
		Name: "printf",
		Args: []ast.FuncParameter{},
		Body: &ast.StringLiteral{
			Value:    "1",
			StartPos: 0,
			EndPos:   3,
		},
		ReturnValues: []string{
			"string",
		},
		NumReturnValues: 1,
		StartPos:        0,
		EndPos:          18,
	}

	tests := []struct {
		name       string
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   datavalue.Value
	}{
		{
			name:       "number",
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
		{
			name:       "string",
			inputLeft:  datavalue.String("5"),
			inputRight: datavalue.String("5"),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.StringLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.StringLiteral{Value: "5", StartPos: 2, EndPos: 3},
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
		{
			name:       "boolean",
			inputLeft:  datavalue.Bool(true),
			inputRight: datavalue.Bool(true),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				Right: &ast.BoolLiteral{Value: "true", StartPos: 2, EndPos: 3},
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
		{
			name:       "null",
			inputLeft:  datavalue.Null(),
			inputRight: datavalue.Null(),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NullLiteral{StartPos: 0, EndPos: 1},
				Right: &ast.NullLiteral{StartPos: 2, EndPos: 3},
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
		{
			name:       "function",
			inputLeft:  datavalue.Function(nil),
			inputRight: datavalue.Function(nil),
			inputNode: &ast.BinaryExpr{
				Left:  function,
				Right: function,
				Operator: token.Token{
					Atom:      "==",
					TokenType: token.TokenTypeEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(false),
		},
		{
			name:       "tuple",
			inputLeft:  datavalue.Tuple(datavalue.Number(1), datavalue.Number(2)),
			inputRight: datavalue.Tuple(datavalue.Number(1), datavalue.Number(2)),
			inputNode: &ast.BinaryExpr{
				Left:  nil,
				Right: nil,
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
		{
			name:       "different data types",
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.String("5"),
			inputNode: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right: &ast.StringLiteral{Value: "5", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "==",
					TokenType: token.TokenTypeEqual,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Bool(false),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := NewEvaluator(io.Discard).evaluateEqualityBinaryExpr(
				test.inputLeft,
				test.inputRight,
				test.inputNode,
			)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.inputNode.Expr(), err)
			}

			if result.Value.DataType().AsString() != test.expected.DataType().AsString() {
				t.Fatalf(
					"expected \"%v\", got \"%v\"",
					test.expected.DataType().AsString(),
					result.Value.DataType().AsString(),
				)
			}
		})
	}
}
