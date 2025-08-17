package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (p *Parser) parseNullLiteral() (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()

	return &ast.NullLiteral{
		StartPos: startPos,
		EndPos:   startPos + 4,
	}, nil
}
