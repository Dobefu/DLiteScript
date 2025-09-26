package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseArrayLiteral(
	recursionDepth int,
) (ast.ExprNode, error) {
	startPos := p.GetCurrentPosition()
	var values []ast.ExprNode

	p.handleOptionalNewlines()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeRBracket {
		_, _ = p.GetNextToken()

		return &ast.ArrayLiteral{
			Values: values,
			Range: ast.Range{
				Start: startPos,
				End:   p.GetCurrentPosition(),
			},
		}, nil
	}

	values, err = p.parseArrayLiteralValues(recursionDepth)

	if err != nil {
		return nil, err
	}

	return &ast.ArrayLiteral{
		Values: values,
		Range: ast.Range{
			Start: startPos,
			End:   p.GetCurrentPosition(),
		},
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
			ast.Range{
				Start: ast.Position{
					Offset: nextToken.StartPos,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: nextToken.EndPos,
					Line:   p.line,
					Column: p.column,
				},
			},
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
