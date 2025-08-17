package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseConstantDeclaration() (ast.ExprNode, error) {
	// The "const" keyword has already been consumed,
	// so we should get the start position from the previous token.
	startPos := p.tokens[p.tokenIdx-1].StartPos
	varName, varType, err := p.parseDeclarationHeader()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType != token.TokenTypeAssign {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgConstantDeclarationWithNoValue,
			p.tokenIdx,
			varName,
		)
	}

	nextToken, err = p.GetNextToken()

	if err != nil {
		return nil, err
	}

	value, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	return &ast.ConstantDeclaration{
		Name:     varName,
		Type:     varType,
		Value:    value,
		StartPos: startPos,
		EndPos:   value.EndPosition(),
	}, nil
}
