package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseVariableDeclaration() (*ast.VariableDeclaration, error) {
	varName, varType, err := p.parseDeclarationHeader()

	if err != nil {
		return nil, err
	}

	var value ast.ExprNode

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
		}
	}

	return &ast.VariableDeclaration{
		Name:     varName,
		Type:     varType,
		Value:    value,
		StartPos: p.tokenIdx - 1,
		EndPos:   p.tokenIdx,
	}, nil
}
