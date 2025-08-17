package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseIdentifier(
	identifierToken *token.Token,
) (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()

	return &ast.Identifier{
		Value:    identifierToken.Atom,
		StartPos: startPos,
		EndPos:   startPos + len(identifierToken.Atom),
	}, nil
}
