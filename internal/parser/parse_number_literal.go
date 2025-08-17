package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseNumberLiteral(
	currentToken *token.Token,
) (ast.ExprNode, error) {
	return &ast.NumberLiteral{
		Value:    currentToken.Atom,
		StartPos: currentToken.StartPos,
		EndPos:   currentToken.EndPos,
	}, nil
}
