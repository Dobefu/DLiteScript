package parser

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parsePrefixExpr(
	currentToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	switch currentToken.TokenType {
	case
		token.TokenTypeNumber:
		return p.parseNumberLiteral(currentToken)

	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub,
		token.TokenTypeNot:
		return p.parseUnaryOperator(currentToken, recursionDepth)

	case
		token.TokenTypeLParen:
		return p.parseParenthesizedExpr(recursionDepth)

	case
		token.TokenTypeLBracket:
		return p.parseArrayLiteral(recursionDepth)

	case
		token.TokenTypeIdentifier:
		return p.parseFunctionCallOrIdentifier(currentToken, recursionDepth)

	case
		token.TokenTypeString:
		return p.parseStringLiteral(currentToken)

	case
		token.TokenTypeBool:
		return p.parseBoolLiteral(currentToken)

	case
		token.TokenTypeNull:
		return p.parseNullLiteral()

	case
		token.TokenTypeOperationSpread:
		return p.parseSpreadExpr(currentToken, recursionDepth)

	default:
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedToken,
			p.tokenIdx,
			currentToken.Atom,
		)
	}
}

func (p *Parser) parseUnaryOperator(
	operatorToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()
	if err != nil {
		return nil, err
	}

	operand, err := p.parseExpr(
		nextToken,
		nil,
		p.getBindingPower(operatorToken, true),
		recursionDepth+1,
	)

	if err != nil {
		return nil, err
	}

	startPos := p.GetCurrentPosition()

	return &ast.PrefixExpr{
		Operator: *operatorToken,
		Operand:  operand,
		Range: ast.Range{
			Start: startPos,
			End:   operand.GetRange().End,
		},
	}, nil
}

func (p *Parser) parseParenthesizedExpr(
	recursionDepth int,
) (ast.ExprNode, error) {
	p.handleOptionalNewlines()
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(nextToken, nil, 0, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	p.handleOptionalNewlines()
	rparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgParenNotClosedAtEOF,
			p.tokenIdx,
		)
	}

	if rparenToken.TokenType != token.TokenTypeRParen {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgExpectedCloseParen,
			p.tokenIdx,
			rparenToken.Atom,
		)
	}

	return expr, nil
}

func (p *Parser) parseFunctionCallOrIdentifier(
	functionCallOrIdentifierToken *token.Token,
	recursionDepth int,
) (ast.ExprNode, error) {
	if p.isEOF {
		return p.parseIdentifier(functionCallOrIdentifierToken)
	}

	nextToken, _ := p.PeekNextToken()

	if nextToken.TokenType == token.TokenTypeLParen {
		return p.parseFunctionCall(
			"",
			functionCallOrIdentifierToken.Atom,
			recursionDepth+1,
		)
	}

	if nextToken.TokenType == token.TokenTypeDot {
		namespace := functionCallOrIdentifierToken.Atom

		_, _ = p.GetNextToken()
		functionNameOrIdentifierToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		if functionNameOrIdentifierToken.TokenType != token.TokenTypeIdentifier {
			return nil, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgUnexpectedToken,
				p.tokenIdx,
				functionNameOrIdentifierToken.Atom,
			)
		}

		nextToken, err := p.PeekNextToken()

		if err != nil {
			return &ast.Identifier{
				Value: fmt.Sprintf("%s.%s", namespace, functionNameOrIdentifierToken.Atom),
				Range: ast.Range{
					Start: ast.Position{
						Offset: functionNameOrIdentifierToken.StartPos,
						Line:   p.line,
						Column: p.column,
					},
					End: ast.Position{
						Offset: functionNameOrIdentifierToken.EndPos,
						Line:   p.line,
						Column: p.column + (functionNameOrIdentifierToken.EndPos - functionNameOrIdentifierToken.StartPos),
					},
				},
			}, nil
		}

		if nextToken.TokenType == token.TokenTypeLParen {
			return p.parseFunctionCall(
				namespace,
				functionNameOrIdentifierToken.Atom,
				recursionDepth+1,
			)
		}

		return &ast.Identifier{
			Value: fmt.Sprintf("%s.%s", namespace, functionNameOrIdentifierToken.Atom),
			Range: ast.Range{
				Start: ast.Position{
					Offset: functionNameOrIdentifierToken.StartPos,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: functionNameOrIdentifierToken.EndPos,
					Line:   p.line,
					Column: p.column + (functionNameOrIdentifierToken.EndPos - functionNameOrIdentifierToken.StartPos),
				},
			},
		}, nil
	}

	return p.parseIdentifier(functionCallOrIdentifierToken)
}
