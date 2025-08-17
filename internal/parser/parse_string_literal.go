package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseStringLiteral(token *token.Token) (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()

	return &ast.StringLiteral{
		Value:    token.Atom,
		StartPos: startPos,
		EndPos:   startPos + len(token.Atom),
	}, nil
}
