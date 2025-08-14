package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (p *Parser) parseNullLiteral() (ast.ExprNode, error) {
	return &ast.NullLiteral{
		Pos: p.tokenIdx - 1,
	}, nil
}
