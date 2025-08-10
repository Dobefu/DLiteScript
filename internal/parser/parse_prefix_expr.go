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
	case token.TokenTypeNumber:
		return p.parseNumberLiteral(currentToken)

	case token.TokenTypeOperationAdd, token.TokenTypeOperationSub:
		return p.parseUnaryOperator(currentToken, recursionDepth)

	case token.TokenTypeLParen:
		return p.parseParenthesizedExpr(recursionDepth)

	case token.TokenTypeIdentifier:
		return p.parseFunctionCallOrIdentifier(currentToken, recursionDepth)

	case token.TokenTypeString:
		return p.parseStringLiteral(currentToken)

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

	return &ast.PrefixExpr{
		Operator: *operatorToken,
		Operand:  operand,
		Pos:      p.tokenIdx - 1,
	}, nil
}

func (p *Parser) parseParenthesizedExpr(
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
		return p.parseFunctionCall(functionCallOrIdentifierToken.Atom, p.tokenIdx-1, recursionDepth+1)
	}

	return p.parseIdentifier(functionCallOrIdentifierToken)
}

func (p *Parser) parseStringLiteral(token *token.Token) (ast.ExprNode, error) {
	return &ast.StringLiteral{
		Value: token.Atom,
		Pos:   p.tokenIdx - 1,
	}, nil
}
