package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseArrayLiteral(
	recursionDepth int,
) (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()
	var values []ast.ExprNode

	err := p.handleOptionalNewlines()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeRBracket {
		_, _ = p.GetNextToken()

		return &ast.ArrayLiteral{
			Values:   values,
			StartPos: startPos,
			EndPos:   p.GetCurrentCharPos(),
		}, nil
	}

	values, err = p.parseArrayLiteralValues(recursionDepth)

	if err != nil {
		return nil, err
	}

	return &ast.ArrayLiteral{
		Values:   values,
		StartPos: startPos,
		EndPos:   p.GetCurrentCharPos(),
	}, nil
}

func (p *Parser) parseArrayLiteralValues(
	recursionDepth int,
) ([]ast.ExprNode, error) {
	var values []ast.ExprNode

	for {
		expr, err := p.parseArrayLiteralExpr(recursionDepth)

		if err != nil {
			return nil, err
		}

		values = append(values, expr)
		nextToken, err := p.PeekNextToken()

		if err != nil {
			return nil, err
		}

		if nextToken.TokenType == token.TokenTypeRBracket {
			_, _ = p.GetNextToken()

			break
		}

		if nextToken.TokenType == token.TokenTypeComma {
			_, _ = p.GetNextToken()

			continue
		}

		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedToken,
			nextToken.StartPos,
			nextToken.Atom,
		)
	}

	return values, nil
}

func (p *Parser) parseArrayLiteralExpr(
	recursionDepth int,
) (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(nextToken, nil, 0, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	return expr, nil
}
