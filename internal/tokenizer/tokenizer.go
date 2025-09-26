// Package tokenizer provides a struct that handles tokenisation of a string.
package tokenizer

import (
	"unicode/utf8"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

// Tokenizer defines the tokenizer itself.
type Tokenizer struct {
	exp     string
	expLen  int
	expIdx  int
	byteIdx int
	line    int
	col     int
	isEOF   bool
}

// NewTokenizer creates a new instance of the Tokenizer struct.
func NewTokenizer(exp string) *Tokenizer {
	return &Tokenizer{
		exp:     exp,
		expLen:  utf8.RuneCountInString(exp),
		expIdx:  0,
		byteIdx: 0,
		line:    0,
		col:     0,
		isEOF:   len(exp) == 0,
	}
}

// GetCurrentPosition gets the current position.
func (t *Tokenizer) GetCurrentPosition() ast.Position {
	return ast.Position{
		Offset: t.byteIdx,
		Line:   t.line,
		Column: t.col,
	}
}
