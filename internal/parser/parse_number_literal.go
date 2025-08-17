package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseNumberLiteral(currentToken *token.Token) (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()

	return &ast.NumberLiteral{
		Value:    currentToken.Atom,
		StartPos: startPos,
		EndPos:   startPos + len(currentToken.Atom),
	}, nil
}
