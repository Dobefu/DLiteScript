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
					&ast.NumberLiteral{Value: "1", StartPos: 4, EndPos: 5},
				},
				StartPos: 0,
				EndPos:   6,
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
					&ast.NumberLiteral{Value: "1", StartPos: 4, EndPos: 5},
					&ast.NumberLiteral{Value: "1", StartPos: 6, EndPos: 7},
				},
				StartPos: 0,
				EndPos:   8,
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

		if expr.StartPosition() != test.expected.StartPosition() {
			t.Errorf(
				"expected StartPos to be \"%d\", got \"%d\"",
				test.expected.StartPosition(),
				expr.StartPosition(),
			)
		}

		if expr.EndPosition() != test.expected.EndPosition() {
			t.Errorf(
				"expected EndPos to be \"%d\", got \"%d\"",
				test.expected.EndPosition(),
				expr.EndPosition(),
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
