package parser

import (
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseContinueStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeNumber &&
		!strings.Contains(nextToken.Atom, ".") {
		_, _ = p.GetNextToken()
		continueCount, err := strconv.Atoi(nextToken.Atom)

		if err != nil {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgInvalidNumber,
				p.tokenIdx,
				nextToken.Atom,
			)
		}

		if continueCount < 1 {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgContinueCountLessThanOne,
				p.tokenIdx,
			)
		}

		endPos := p.GetCurrentCharPos()

		return &ast.ContinueStatement{
			Count:    continueCount,
			StartPos: startPos,
			EndPos:   endPos,
		}, nil
	}

	endPos := startPos + 8

	return &ast.ContinueStatement{
		Count:    1,
		StartPos: startPos,
		EndPos:   endPos,
	}, nil
}
