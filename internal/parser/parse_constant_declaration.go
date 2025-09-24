package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseConstantDeclaration() (ast.ExprNode, error) {
	// Get start position from current position
	startPos := p.GetCurrentPosition()
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
			errorutil.StageParse,
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
		Name:  varName,
		Type:  varType,
		Value: value,
		Range: ast.Range{
			Start: startPos,
			End:   p.GetCurrentPosition(),
		},
	}, nil
}
