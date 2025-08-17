package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseIdentifier(
	identifierToken *token.Token,
) (ast.ExprNode, error) {
	return &ast.Identifier{
		Value:    identifierToken.Atom,
		StartPos: p.tokenIdx - 1,
		EndPos:   p.tokenIdx,
	}, nil
}
