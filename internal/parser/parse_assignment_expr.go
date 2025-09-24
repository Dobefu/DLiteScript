package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (p *Parser) parseAssignmentExpr(
	leftExpr ast.ExprNode,
	minPrecedence int,
	recursionDepth int,
) (ast.ExprNode, error) {
	identifier, isIdentifier := leftExpr.(*ast.Identifier)
	indexExpr, isIndexExpr := leftExpr.(*ast.IndexExpr)

	if !isIdentifier && !isIndexExpr {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedToken,
			p.tokenIdx,
			leftExpr.Expr(),
		)
	}

	_, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	rightToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	rightExpr, err := p.parseExpr(rightToken, nil, minPrecedence, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	if isIdentifier {
		return &ast.AssignmentStatement{
			Left:  identifier,
			Right: rightExpr,
			Range: ast.Range{
				Start: leftExpr.GetRange().Start,
				End:   rightExpr.GetRange().End,
			},
		}, nil
	}

	return &ast.IndexAssignmentStatement{
		Array: indexExpr.Array,
		Index: indexExpr.Index,
		Right: rightExpr,
		Range: ast.Range{
			Start: leftExpr.GetRange().Start,
			End:   rightExpr.GetRange().End,
		},
	}, nil
}
