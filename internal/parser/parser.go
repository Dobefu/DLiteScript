// Package parser defines the actual parser for DLiteScript.
package parser

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

// Parser defines the parser itself.
type Parser struct {
	tokens   []*token.Token
	tokenIdx int
	tokenLen int
	charIdx  int
	isEOF    bool
}

// NewParser creates a new instance of the Parser struct.
func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:   tokens,
		tokenIdx: 0,
		tokenLen: len(tokens),
		charIdx:  0,
		isEOF:    len(tokens) == 0,
	}
}
