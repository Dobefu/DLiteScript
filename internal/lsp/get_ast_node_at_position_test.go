package lsp

import (
	"encoding/json"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestGetAstNodeAtPosition(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     ast.ExprNode
		charIndex int
		expected  ast.ExprNode
	}{
		{
			name: "null literal node",
			input: &ast.NullLiteral{
				StartPos: 0,
				EndPos:   4,
			},
			charIndex: 5,
			expected:  nil,
		},
		{
			name: "string literal node",
			input: &ast.StringLiteral{
				Value:    "test",
				StartPos: 0,
				EndPos:   4,
			},
			charIndex: 0,
			expected: &ast.StringLiteral{
				Value:    "test",
				StartPos: 0,
				EndPos:   4,
			},
		},
		{
			name: "binary expression",
			input: &ast.BinaryExpr{
				Left:  &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "5", StartPos: 4, EndPos: 5},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  2,
					EndPos:    3,
				},
				StartPos: 0,
				EndPos:   5,
			},
			charIndex: 0,
			expected:  &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			node := getAstNodeAtPosition(test.input, test.charIndex)

			nodeJSON, err := json.Marshal(node)

			if err != nil {
				t.Errorf("error marshalling node: %v", err)
			}

			expectedJSON, err := json.Marshal(test.expected)

			if err != nil {
				t.Errorf("error marshalling expected node: %v", err)
			}

			if string(nodeJSON) != string(expectedJSON) {
				t.Errorf("expected \"%v\", got \"%v\"", string(expectedJSON), string(nodeJSON))
			}
		})
	}
}
