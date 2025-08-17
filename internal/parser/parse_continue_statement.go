package parser

import (
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseContinueStatement() (ast.ExprNode, error) {
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeNumber &&
		!strings.Contains(nextToken.Atom, ".") {
		if _, err := p.GetNextToken(); err != nil {
			return nil, err
		}

		continueCount, err := strconv.Atoi(nextToken.Atom)

		if err != nil {
			return nil, err
		}

		if continueCount < 1 {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgContinueCountLessThanOne,
				p.tokenIdx,
			)
		}

		return &ast.ContinueStatement{
			Count:    continueCount,
			StartPos: p.tokenIdx,
			EndPos:   p.tokenIdx,
		}, nil
	}

	return &ast.ContinueStatement{
		Count:    1,
		StartPos: p.tokenIdx,
		EndPos:   p.tokenIdx,
	}, nil
}
