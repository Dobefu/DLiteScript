// Package token defines a Token struct.
package token

// Token defines a single expression token.
type Token struct {
	Atom      string
	TokenType Type
	StartPos  int
	EndPos    int
}

// NewToken creates a new Token.
func NewToken(atom string, tokenType Type, startPos int, endPos int) *Token {
	return &Token{
		Atom:      atom,
		TokenType: tokenType,
		StartPos:  startPos,
		EndPos:    endPos,
	}
}
