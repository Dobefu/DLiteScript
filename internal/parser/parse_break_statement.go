package parser

import (
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseBreakStatement() (ast.ExprNode, error) {
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeNumber &&
		!strings.Contains(nextToken.Atom, ".") {
		if _, err := p.GetNextToken(); err != nil {
			return nil, err
		}

		breakCount, err := strconv.Atoi(nextToken.Atom)

		if err != nil {
			return nil, err
		}

		if breakCount < 1 {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgBreakCountLessThanOne,
				p.tokenIdx,
			)
		}

		return &ast.BreakStatement{
			Count:    breakCount,
			StartPos: p.tokenIdx,
			EndPos:   p.tokenIdx,
		}, nil
	}

	return &ast.BreakStatement{
		Count:    1,
		StartPos: p.tokenIdx,
		EndPos:   p.tokenIdx,
	}, nil
}
