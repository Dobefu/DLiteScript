package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseBoolLiteral(token *token.Token) (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()

	return &ast.BoolLiteral{
		Value:    token.Atom,
		StartPos: startPos,
		EndPos:   startPos + len(token.Atom),
	}, nil
}
