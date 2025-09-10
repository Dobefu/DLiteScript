package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseBinaryExpr(
	operatorToken *token.Token,
	leftExpr ast.ExprNode,
	rightToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	if leftExpr == nil {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgUnexpectedToken,
			p.tokenIdx,
			operatorToken.Atom,
		)
	}

	rightExpr, err := p.parseExpr(
		rightToken,
		nil,
		p.getBindingPower(operatorToken, false)+1,
		recursionDepth+1,
	)

	if err != nil {
		return nil, err
	}

	return &ast.BinaryExpr{
		Left:     leftExpr,
		Right:    rightExpr,
		Operator: *operatorToken,
		StartPos: leftExpr.StartPosition(),
		EndPos:   rightExpr.EndPosition(),
	}, nil
}
