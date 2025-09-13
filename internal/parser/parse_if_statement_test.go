package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseIfStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "true condition",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
			},
			expected: "if true { (1) }",
		},
		{
			name: "true condition with else",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
				token.NewToken("else", token.TokenTypeElse, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("2", token.TokenTypeNumber, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
			},
			expected: "if true { (1) } else { (2) }",
		},
		{
			name: "true condition with empty else",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
				token.NewToken("else", token.TokenTypeElse, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
			},
			expected: "if true { (1) } else { () }",
		},
		{
			name: "else if",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
				token.NewToken("else", token.TokenTypeElse, 0, 0),
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("false", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
				token.NewToken("else", token.TokenTypeElse, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
			},
			expected: "if true { () } else { (if false { () } else { () }) }",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			expr, err := p.Parse()

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, expr.Expr())
			}
		})
	}
}

func TestParseIfStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "no next token after if",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token after if",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "{"),
			),
		},
		{
			name: "no block statement after if",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token in then block",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("=", token.TokenTypeAssign, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 4",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "="),
			),
		},
		{
			name: "no next token after else",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
				token.NewToken("else", token.TokenTypeElse, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token after else",
			input: []*token.Token{
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("}", token.TokenTypeRBrace, 0, 0),
				token.NewToken("else", token.TokenTypeElse, 0, 0),
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 6",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.Parse()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestHandleElseBlockErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "nested if statement parsing error",
			input: []*token.Token{
				token.NewToken("{", token.TokenTypeLBrace, 0, 0),
				token.NewToken("if", token.TokenTypeIf, 0, 0),
				token.NewToken("true", token.TokenTypeBool, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 3",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.handleElseBlock()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseThenBlockErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseThenBlock(token.TokenTypeRBrace)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
