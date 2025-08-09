// Package token defines a Token struct.
package token

// Token defines a single expression token.
type Token struct {
	Atom      string
	TokenType Type
}

// NewToken creates a new Token.
func NewToken(atom string, tokenType Type) *Token {
	return &Token{
		Atom:      atom,
		TokenType: tokenType,
	}
}
