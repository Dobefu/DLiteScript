package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseImportStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "import statement",
			input: []*token.Token{
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: "import \"test\"",
		},
		{
			name: "import statement with alias",
			input: []*token.Token{
				{Atom: "test", TokenType: token.TokenTypeString},
				{Atom: "as", TokenType: token.TokenTypeAs},
				{Atom: "test", TokenType: token.TokenTypeIdentifier},
			},
			expected: "import \"test\" as test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			importStmt, err := p.parseImportStatement(&token.Token{
				Atom:      "import",
				TokenType: token.TokenTypeImport,
				StartPos:  0,
				EndPos:    6,
			})

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if importStmt.Expr() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, importStmt.Expr())
			}
		})
	}

	t.Run("import statement with EOF after path", func(t *testing.T) {
		t.Parallel()

		p := NewParser([]*token.Token{
			{Atom: "test", TokenType: token.TokenTypeString},
		})

		importStmt, err := p.parseImportStatement(&token.Token{
			Atom:      "import",
			TokenType: token.TokenTypeImport,
			StartPos:  0,
			EndPos:    6,
		})

		if err != nil {
			t.Fatalf("expected no error, got: %s", err.Error())
		}

		expected := "import \"test\""
		if importStmt.Expr() != expected {
			t.Fatalf("expected \"%s\", got \"%s\"", expected, importStmt.Expr())
		}
	})
}

func TestParseImportStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     []*token.Token
		nextToken *token.Token
		expected  string
	}{
		{
			name:      "unexpected EOF ",
			input:     []*token.Token{},
			nextToken: nil,
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name:  "unexpected EOF after import keyword",
			input: []*token.Token{},
			nextToken: &token.Token{
				Atom:      "import",
				TokenType: token.TokenTypeImport,
				StartPos:  0,
				EndPos:    6,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid path type",
			input: []*token.Token{
				{Atom: "test", TokenType: token.TokenTypeIdentifier},
			},
			nextToken: &token.Token{
				Atom:      "import",
				TokenType: token.TokenTypeImport,
				StartPos:  0,
				EndPos:    6,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "test"),
			),
		},
		{
			name: "unexpected EOF after alias",
			input: []*token.Token{
				{Atom: "test", TokenType: token.TokenTypeString},
				{Atom: "as", TokenType: token.TokenTypeAs},
			},
			nextToken: &token.Token{
				Atom:      "import",
				TokenType: token.TokenTypeImport,
				StartPos:  0,
				EndPos:    6,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid alias token type",
			input: []*token.Token{
				{Atom: "test", TokenType: token.TokenTypeString},
				{Atom: "as", TokenType: token.TokenTypeAs},
				{Atom: "123", TokenType: token.TokenTypeNumber},
			},
			nextToken: &token.Token{
				Atom:      "import",
				TokenType: token.TokenTypeImport,
				StartPos:  0,
				EndPos:    6,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "123"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseImportStatement(test.nextToken)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
