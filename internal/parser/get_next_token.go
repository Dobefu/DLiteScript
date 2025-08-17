package parser

import (
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// GetNextToken gets the next token and advances the current token index.
func (p *Parser) GetNextToken() (*token.Token, error) {
	if p.isEOF {
		return nil, errorutil.NewErrorAt(errorutil.ErrorMsgUnexpectedEOF, p.tokenIdx)
	}

	next := p.tokens[p.tokenIdx]
	p.tokenIdx++

	p.charIdx = next.EndPos

	if p.tokenIdx >= p.tokenLen {
		p.isEOF = true
	}

	return next, nil
}
