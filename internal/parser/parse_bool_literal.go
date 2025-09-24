package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseBoolLiteral(token *token.Token) (ast.ExprNode, error) {
	return &ast.BoolLiteral{
		Value: token.Atom,
		Range: ast.Range{
			Start: ast.Position{
				Offset: token.StartPos,
				Line:   p.line,
				Column: p.column,
			},
			End: ast.Position{
				Offset: token.EndPos,
				Line:   p.line,
				Column: p.column + (token.EndPos - token.StartPos),
			},
		},
	}, nil
}
