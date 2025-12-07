package parser

import (
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseContinueStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentPosition()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeNumber &&
		!strings.Contains(nextToken.Atom, ".") {
		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}

		continueCount, err := strconv.Atoi(nextToken.Atom)

		if err != nil {
			return nil, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgInvalidNumber,
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
				nextToken.Atom,
			)
		}

		if continueCount < 1 {
			return nil, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgContinueCountLessThanOne,
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

		endPos := p.GetCurrentPosition()

		return &ast.ContinueStatement{
			Count: continueCount,
			Range: ast.Range{
				Start: startPos,
				End:   endPos,
			},
		}, nil
	}

	endPos := p.GetCurrentPosition()

	return &ast.ContinueStatement{
		Count: 1,
		Range: ast.Range{
			Start: startPos,
			End:   endPos,
		},
	}, nil
}
