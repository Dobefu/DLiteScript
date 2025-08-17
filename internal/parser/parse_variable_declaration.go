package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseVariableDeclaration() (*ast.VariableDeclaration, error) {
	// The "var" keyword has already been consumed.
	startPos := p.GetCurrentCharPos() - 3
	varName, varType, err := p.parseDeclarationHeader()

	if err != nil {
		return nil, err
	}

	var value ast.ExprNode
	endPos := p.GetCurrentCharPos()

	if !p.isEOF {
		nextToken, err := p.PeekNextToken()

		if err == nil && nextToken.TokenType == token.TokenTypeAssign {
			_, err = p.GetNextToken()

			if err != nil {
				return nil, err
			}

			nextToken, err = p.GetNextToken()

			if err != nil {
				return nil, err
			}

			value, err = p.parseExpr(nextToken, nil, 0, 0)

			if err != nil {
				return nil, err
			}

			endPos = value.EndPosition()
		}
	}

	return &ast.VariableDeclaration{
		Name:     varName,
		Type:     varType,
		Value:    value,
		StartPos: startPos,
		EndPos:   endPos,
	}, nil
}
