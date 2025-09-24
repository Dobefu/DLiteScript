package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseVariableDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected *ast.VariableDeclaration
	}{
		{
			name: "simple variable declaration",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
			},
			expected: &ast.VariableDeclaration{
				Name:  "x",
				Type:  "number",
				Value: nil,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "variable declaration with initialisation",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: &ast.VariableDeclaration{
				Name: "x",
				Type: "number",
				Value: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
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

			p := NewParser(test.input)
			p.tokenIdx = 1
			declaration, err := p.parseVariableDeclaration()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if declaration.Expr() != test.expected.Expr() {
				t.Fatalf(
					"expected %s, got %s",
					test.expected.Expr(),
					declaration.Expr(),
				)
			}
		})
	}
}

func TestParseVariableDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "no tokens after var",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected EOF after assignment operator",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 4",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid expression after assignment",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			p.tokenIdx = 1

			if len(test.input) <= 1 {
				p.isEOF = true
			}

			_, err := p.parseVariableDeclaration()

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
