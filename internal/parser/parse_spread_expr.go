package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseSpreadExpr(
	currentToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(
		nextToken,
		nil,
		p.getBindingPower(currentToken, true),
		recursionDepth,
	)

	if err != nil {
		return nil, err
	}

	return &ast.SpreadExpr{
		Expression: expr,
		StartPos:   currentToken.StartPos,
		EndPos:     nextToken.EndPos,
	}, nil
}
