package parser

import (
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseBreakStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentPosition()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeNumber &&
		!strings.Contains(nextToken.Atom, ".") {
		_, _ = p.GetNextToken()

		breakCount, err := strconv.Atoi(nextToken.Atom)

		if err != nil {
			return nil, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgInvalidNumber,
				p.tokenIdx,
				nextToken.Atom,
			)
		}

		if breakCount < 1 {
			return nil, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgBreakCountLessThanOne,
				p.tokenIdx,
			)
		}

		endPos := p.GetCurrentPosition()

		return &ast.BreakStatement{
			Count: breakCount,
			Range: ast.Range{
				Start: startPos,
				End:   endPos,
			},
		}, nil
	}

	endPos := p.GetCurrentPosition()

	return &ast.BreakStatement{
		Count: 1,
		Range: ast.Range{
			Start: startPos,
			End:   endPos,
		},
	}, nil
}
