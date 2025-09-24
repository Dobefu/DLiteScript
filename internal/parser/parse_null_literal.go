package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (p *Parser) parseNullLiteral() (ast.ExprNode, error) {
	endPos := p.GetCurrentPosition()
	startPos := ast.Position{
		Offset: endPos.Offset - 4,
		Line:   endPos.Line,
		Column: endPos.Column - 4,
	}

	return &ast.NullLiteral{
		Range: ast.Range{
			Start: startPos,
			End:   endPos,
		},
	}, nil
}
