// Package parser defines the actual parser for DLiteScript.
package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// Parser defines the parser itself.
type Parser struct {
	tokens   []*token.Token
	tokenIdx int
	tokenLen int
	charIdx  int
	line     int
	column   int
	isEOF    bool
}

// NewParser creates a new instance of the Parser struct.
func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:   tokens,
		tokenIdx: 0,
		tokenLen: len(tokens),
		charIdx:  0,
		line:     0,
		column:   0,
		isEOF:    len(tokens) == 0,
	}
}

// GetCurrentPosition gets the current position.
func (p *Parser) GetCurrentPosition() ast.Position {
	return ast.Position{
		Offset: p.charIdx,
		Line:   p.line,
		Column: p.column,
	}
}

// AdvancePosition advances the position.
func (p *Parser) AdvancePosition(token *token.Token) {
	for _, char := range token.Atom {
		if char == '\n' {
			p.line++
			p.column = 0

			continue
		}

		p.column++
	}

	p.charIdx = token.EndPos
}
