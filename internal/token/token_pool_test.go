package token

import (
	"testing"
)

func TestPool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		atom      string
		tokenType Type
	}{
		{"+", TokenTypeOperationAdd},
		{"-", TokenTypeOperationSub},
		{"*", TokenTypeOperationMul},
		{"/", TokenTypeOperationDiv},
		{"%", TokenTypeOperationMod},
		{"**", TokenTypeOperationPow},
		{"(", TokenTypeLParen},
		{")", TokenTypeRParen},
		{",", TokenTypeComma},

		{"abs", TokenTypeIdentifier},
	}

	pool := NewPool()

	if pool.GetPoolSize() != len(pool.pool) {
		t.Errorf("expected pool size to be '%d', got '%d'", len(pool.pool), pool.GetPoolSize())
	}

	for _, test := range tests {
		token := pool.GetToken(test.atom, test.tokenType)

		if token.Atom != test.atom {
			t.Errorf("expected token atom to be '%s', got '%s'", test.atom, token.Atom)
		}

		if token.TokenType != test.tokenType {
			t.Errorf("expected token type to be '%d', got '%d'", test.tokenType, token.TokenType)
		}
	}
}
