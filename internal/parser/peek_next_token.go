package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// PeekNextToken gets the next token without advancing the current token index.
func (p *Parser) PeekNextToken() (*token.Token, error) {
	if p.isEOF {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedEOF,
			ast.Range{
				Start: ast.Position{
					Offset: p.tokenIdx,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: p.tokenIdx,
					Line:   p.line,
					Column: p.column,
				},
			},
		)
	}

	return p.tokens[p.tokenIdx], nil
}
