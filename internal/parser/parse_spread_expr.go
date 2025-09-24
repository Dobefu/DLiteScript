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
		Range: ast.Range{
			Start: ast.Position{
				Offset: currentToken.StartPos,
				Line:   p.line,
				Column: p.column,
			},
			End: ast.Position{
				Offset: nextToken.EndPos,
				Line:   p.line,
				Column: p.column + (nextToken.EndPos - currentToken.StartPos),
			},
		},
	}, nil
}
