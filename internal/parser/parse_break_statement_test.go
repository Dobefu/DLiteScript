package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseBreakStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected ast.ExprNode
	}{
		{
			name: "break",
			input: []*token.Token{
				token.NewToken("break", token.TokenTypeBreak, 0, 0),
			},
			expected: &ast.BreakStatement{
				Count: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "break 2",
			input: []*token.Token{
				token.NewToken("break", token.TokenTypeBreak, 0, 0),
				token.NewToken("2", token.TokenTypeNumber, 0, 0),
			},
			expected: &ast.BreakStatement{
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
				_, err := parser.GetNextToken()

				if err != nil {
					t.Fatalf("could not get token: %s", err.Error())
				}
			}

			expr, err := parser.parseBreakStatement()

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected.Expr() {
				t.Fatalf("expected %s, got %s", test.expected.Expr(), expr.Expr())
			}
		})
	}
}

func TestParseBreakStatementErr(t *testing.T) {
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
				token.NewToken("break", token.TokenTypeBreak, 0, 0),
				token.NewToken("bogus", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 11",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgInvalidNumber, "bogus"),
			),
		},
		{
			name: "less than 1",
			input: []*token.Token{
				token.NewToken("break", token.TokenTypeBreak, 0, 0),
				token.NewToken("-1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgBreakCountLessThanOne,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)

			if len(test.input) > 1 {
				_, err := parser.GetNextToken()

				if err != nil {
					t.Fatalf("could not get token: %s", err.Error())
				}
			}

			_, err := parser.parseBreakStatement()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
