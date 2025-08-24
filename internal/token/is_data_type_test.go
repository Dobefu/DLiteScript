package token

import "testing"

func TestIsDataType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		token    Token
		expected bool
	}{
		{
			name: "number type",
			token: Token{
				Atom:      "number",
				TokenType: TokenTypeTypeNumber,
				StartPos:  0,
				EndPos:    6,
			},
			expected: true,
		},
		{
			name: "number",
			token: Token{
				Atom:      "123",
				TokenType: TokenTypeNumber,
				StartPos:  0,
				EndPos:    3,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := test.token.IsDataType()

			if result != test.expected {
				t.Fatalf("expected %t, got %t", test.expected, result)
			}
		})
	}
}
