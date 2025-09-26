package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseContinueStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected ast.ExprNode
	}{
		{
			name: "continue",
			input: []*token.Token{
				token.NewToken("continue", token.TokenTypeContinue, 0, 0),
			},
			expected: &ast.ContinueStatement{
				Count: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "continue 2",
			input: []*token.Token{
				token.NewToken("continue", token.TokenTypeContinue, 0, 0),
				token.NewToken("2", token.TokenTypeNumber, 0, 0),
			},
			expected: &ast.ContinueStatement{
				Count: 2,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)

			if len(test.input) > 1 {
				_, _ = parser.GetNextToken()
			}

			expr, err := parser.parseContinueStatement()

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected.Expr() {
				t.Fatalf("expected %s, got %s", test.expected.Expr(), expr.Expr())
			}
		})
	}
}

func TestParseContinueStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no input",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid number",
			input: []*token.Token{
				token.NewToken("continue", token.TokenTypeContinue, 0, 0),
				token.NewToken("bogus", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 14",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgInvalidNumber, "bogus"),
			),
		},
		{
			name: "less than 1",
			input: []*token.Token{
				token.NewToken("continue", token.TokenTypeContinue, 0, 0),
				token.NewToken("-1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 11",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgContinueCountLessThanOne,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)

			if len(test.input) > 1 {
				_, _ = parser.GetNextToken()
			}

			_, err := parser.parseContinueStatement()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
