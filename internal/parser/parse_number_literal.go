package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseNumberLiteral(
	currentToken *token.Token,
) (ast.ExprNode, error) {
	return &ast.NumberLiteral{
		Value: currentToken.Atom,
		Range: ast.Range{
			Start: ast.Position{
				Offset: currentToken.StartPos,
				Line:   p.line,
				Column: p.column,
			},
			End: ast.Position{
				Offset: currentToken.EndPos,
				Line:   p.line,
				Column: p.column + (currentToken.EndPos - currentToken.StartPos),
			},
		},
	}, nil
}
