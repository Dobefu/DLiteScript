package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// GetNextToken gets the next token and advances the current token index.
func (p *Parser) GetNextToken() (*token.Token, error) {
	if p.isEOF {
		var startPos, endPos ast.Position

		if p.tokenIdx < len(p.tokens) {
			startPos = ast.Position{
				Offset: p.tokens[p.tokenIdx].StartPos,
				Line:   p.line,
				Column: p.column,
			}
			endPos = ast.Position{
				Offset: p.tokens[p.tokenIdx].EndPos,
				Line:   p.line,
				Column: p.column,
			}
		} else {
			if len(p.tokens) > 0 {
				lastToken := p.tokens[len(p.tokens)-1]
				startPos = ast.Position{
					Offset: lastToken.EndPos,
					Line:   p.line,
					Column: p.column,
				}
			} else {
				startPos = ast.Position{
					Offset: 0,
					Line:   0,
					Column: 0,
				}
			}
			endPos = startPos
		}

		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedEOF,
			ast.Range{
				Start: startPos,
				End:   endPos,
			},
		)
	}

	next := p.tokens[p.tokenIdx]
	p.tokenIdx++
	p.AdvancePosition(next)

	if p.tokenIdx >= p.tokenLen {
		p.isEOF = true
	}

	return next, nil
}
