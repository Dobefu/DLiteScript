package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (p *Parser) parseNullLiteral() (ast.ExprNode, error) {
	endPos := p.GetCurrentCharPos()
	startPos := endPos - 4

	return &ast.NullLiteral{
		StartPos: startPos,
		EndPos:   endPos,
	}, nil
}
