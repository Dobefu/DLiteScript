package parser

import (
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

	default:
		return nil, errorutil.NewErrorAt(
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

	startPos := p.GetCurrentCharPos()

	return &ast.PrefixExpr{
		Operator: *operatorToken,
		Operand:  operand,
		StartPos: startPos,
		EndPos:   operand.EndPosition(),
	}, nil
}

func (p *Parser) parseParenthesizedExpr(
	recursionDepth int,
) (ast.ExprNode, error) {
	err := p.handleOptionalNewlines()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(nextToken, nil, 0, recursionDepth+1)

	if err != nil {
		return nil, err
	}

	err = p.handleOptionalNewlines()

	if err != nil {
		return nil, err
	}

	rparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, errorutil.NewErrorAt(errorutil.ErrorMsgParenNotClosedAtEOF, p.tokenIdx)
	}

	if rparenToken.TokenType != token.TokenTypeRParen {
		return nil, errorutil.NewErrorAt(
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

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeLParen {
		return p.parseFunctionCall(
			functionCallOrIdentifierToken.Atom,
			recursionDepth+1,
		)
	}

	return p.parseIdentifier(functionCallOrIdentifierToken)
}
