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
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.FunctionCall{
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 4, EndPos: 5},
				},
				StartPos: 3,
				EndPos:   6,
			},
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ",", TokenType: token.TokenTypeComma},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.FunctionCall{
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 4, EndPos: 5},
					&ast.NumberLiteral{Value: "1", StartPos: 7, EndPos: 8},
				},
				StartPos: 3,
				EndPos:   8,
			},
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)

		// Skip the function name token.
		startCharPos := parser.GetCurrentCharPos()
		_, err := parser.GetNextToken()

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)
		}

		expr, err := parser.parseFunctionCall("abs", startCharPos)

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
			input:    []*token.Token{},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgExpectedOpenParen, "abs"),
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ",", TokenType: token.TokenTypeComma},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "abs"),
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		_, err := parser.parseFunctionCall("abs", 0)

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
			{Atom: "(", TokenType: token.TokenTypeLParen},
			{Atom: "1", TokenType: token.TokenTypeNumber},
			{Atom: ",", TokenType: token.TokenTypeComma},
			{Atom: "2", TokenType: token.TokenTypeNumber},
			{Atom: ")", TokenType: token.TokenTypeRParen},
		}).parseFunctionCall("min", 0)
	}
}
