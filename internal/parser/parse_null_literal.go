package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (p *Parser) parseNullLiteral() (ast.ExprNode, error) {
	return &ast.NullLiteral{
		StartPos: p.tokenIdx - 1,
		EndPos:   p.tokenIdx,
	}, nil
}
