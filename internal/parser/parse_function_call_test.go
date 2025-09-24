package parser

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected *ast.FunctionCall
	}{
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 5, EndPos: 6},
			},
			expected: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 4, Line: 0, Column: 0},
							End:   ast.Position{Offset: 5, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 6, Line: 0, Column: 0},
				},
			},
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 5, EndPos: 6},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 6, EndPos: 7},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 7, EndPos: 8},
			},
			expected: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 4, Line: 0, Column: 0},
							End:   ast.Position{Offset: 5, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 6, Line: 0, Column: 0},
							End:   ast.Position{Offset: 7, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 8, Line: 0, Column: 0},
				},
			},
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		_, err := parser.GetNextToken()

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)
		}

		expr, err := parser.parseFunctionCall("", "abs", 0)

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)
		}

		if expr.GetRange().Start.Offset != test.expected.Range.Start.Offset {
			t.Errorf(
				"expected StartPos to be \"%d\", got \"%d\"",
				test.expected.Range.Start.Offset,
				expr.GetRange().Start.Offset,
			)
		}

		if expr.GetRange().End.Offset != test.expected.Range.End.Offset {
			t.Errorf(
				"expected EndPos to be \"%d\", got \"%d\"",
				test.expected.Range.End.Offset,
				expr.GetRange().End.Offset,
			)
		}
	}
}

func TestParseFunctionCallErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 4, EndPos: 5},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 5, EndPos: 6},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 5, EndPos: 6},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier, StartPos: 5, EndPos: 8},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "abs"),
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		_, err := parser.GetNextToken()

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)
		}

		_, err = parser.parseFunctionCall("", "abs", 0)

		if err == nil {
			t.Fatalf("expected error, got none for input \"%v\"", test.input)
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				errors.Unwrap(err).Error(),
			)
		}
	}
}

func BenchmarkParseFunctionCall(b *testing.B) {
	for b.Loop() {
		_, _ = NewParser([]*token.Token{
			{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 0, EndPos: 1},
			{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 1, EndPos: 2},
			{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 2, EndPos: 3},
			{Atom: "2", TokenType: token.TokenTypeNumber, StartPos: 3, EndPos: 4},
			{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 4, EndPos: 5},
		}).parseFunctionCall("", "min", 0)
	}
}
