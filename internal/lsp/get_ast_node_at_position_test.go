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
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 4, Line: 0, Column: 0},
				},
			},
			charIndex: 5,
			expected:  nil,
		},
		{
			name: "string literal node",
			input: &ast.StringLiteral{
				Value: "test",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 4, Line: 0, Column: 0},
				},
			},
			charIndex: 0,
			expected: &ast.StringLiteral{
				Value: "test",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 4, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "binary expression",
			input: &ast.BinaryExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  2,
					EndPos:    3,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			charIndex: 0,
			expected: &ast.Identifier{
				Value: "x",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
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
