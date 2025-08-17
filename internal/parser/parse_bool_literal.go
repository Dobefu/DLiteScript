package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseBoolLiteral(token *token.Token) (ast.ExprNode, error) {
	return &ast.BoolLiteral{
		Value:    token.Atom,
		StartPos: token.StartPos,
		EndPos:   token.EndPos,
	}, nil
}
