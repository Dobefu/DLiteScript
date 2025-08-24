package token

import "testing"

func TestNewToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		atom      string
		tokenType Type
		startPos  int
		endPos    int
		expected  *Token
	}{
		{
			name:      "number",
			atom:      "123",
			tokenType: TokenTypeNumber,
			startPos:  0,
			endPos:    3,
			expected:  NewToken("123", TokenTypeNumber, 0, 3),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			token := NewToken(test.atom, test.tokenType, test.startPos, test.endPos)

			if token.Atom != test.expected.Atom {
				t.Fatalf("expected %s, got %s", test.expected.Atom, token.Atom)
			}

			if token.TokenType != test.expected.TokenType {
				t.Fatalf("expected %T, got %T", test.expected.TokenType, token.TokenType)
			}

			if token.StartPos != test.expected.StartPos {
				t.Fatalf("expected %d, got %d", test.expected.StartPos, token.StartPos)
			}

			if token.EndPos != test.expected.EndPos {
				t.Fatalf("expected %d, got %d", test.expected.EndPos, token.EndPos)
			}
		})
	}
}
