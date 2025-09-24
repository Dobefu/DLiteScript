package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseVariableDeclaration() (*ast.VariableDeclaration, error) {
	// The "var" keyword has already been consumed,
	// so we should get the start position from the previous token.
	prevToken := p.tokens[p.tokenIdx-1]

	startPos := ast.Position{
		Offset: prevToken.StartPos,
		Line:   p.line,
		Column: p.column - (prevToken.EndPos - prevToken.StartPos),
	}
	varName, varType, err := p.parseDeclarationHeader()

	if err != nil {
		return nil, err
	}

	var value ast.ExprNode
	endPos := p.GetCurrentPosition()

	if !p.isEOF {
		nextToken, err := p.PeekNextToken()

		if err == nil && nextToken.TokenType == token.TokenTypeAssign {
			_, _ = p.GetNextToken()

			nextToken, err = p.GetNextToken()

			if err != nil {
				return nil, err
			}

			value, err = p.parseExpr(nextToken, nil, 0, 0)

			if err != nil {
				return nil, err
			}

			endPos = value.GetRange().End
		}
	}

	return &ast.VariableDeclaration{
		Name:  varName,
		Type:  varType,
		Value: value,
		Range: ast.Range{
			Start: startPos,
			End:   endPos,
		},
	}, nil
}
