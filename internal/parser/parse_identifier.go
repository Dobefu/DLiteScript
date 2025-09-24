package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseIdentifier(
	identifierToken *token.Token,
) (ast.ExprNode, error) {
	return &ast.Identifier{
		Value: identifierToken.Atom,
		Range: ast.Range{
			Start: ast.Position{
				Offset: identifierToken.StartPos,
				Line:   p.line,
				Column: p.column,
			},
			End: ast.Position{
				Offset: identifierToken.EndPos,
				Line:   p.line,
				Column: p.column + (identifierToken.EndPos - identifierToken.StartPos),
			},
		},
	}, nil
}
