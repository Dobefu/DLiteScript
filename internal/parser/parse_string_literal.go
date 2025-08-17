package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseStringLiteral(token *token.Token) (ast.ExprNode, error) {
	return &ast.StringLiteral{
		Value:    token.Atom,
		StartPos: token.StartPos,
		EndPos:   token.EndPos,
	}, nil
}
