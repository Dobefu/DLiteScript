package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseForStatement() (ast.ExprNode, error) {
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeLBrace {
		return p.parseInfiniteLoop()
	}

	return p.parseLoop()
}

func (p *Parser) parseInfiniteLoop() (ast.ExprNode, error) {
	loopBody, err := p.parseLoopBody()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        nil,
		Body:             loopBody,
		Pos:              p.tokenIdx,
		DeclaredVariable: "",
		RangeVariable:    "",
		RangeFrom:        nil,
		RangeTo:          nil,
		IsRange:          false,
	}, nil
}

func (p *Parser) parseLoop() (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	condition, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	loopBody, err := p.parseLoopBody()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        condition,
		Body:             loopBody,
		Pos:              p.tokenIdx,
		DeclaredVariable: "",
		RangeVariable:    "",
		RangeFrom:        nil,
		RangeTo:          nil,
		IsRange:          false,
	}, nil
}

func (p *Parser) parseLoopBody() (*ast.BlockStatement, error) {
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType != token.TokenTypeLBrace {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgBlockStatementExpected,
			p.tokenIdx,
			nextToken.TokenType,
		)
	}

	_, err = p.GetNextToken()

	if err != nil {
		return nil, err
	}

	var endToken token.Type = token.TokenTypeRBrace
	blockNode, err := p.parseBlock(&endToken)

	if err != nil {
		return nil, err
	}

	blockStatement, isBlockStatement := blockNode.(*ast.BlockStatement)

	if !isBlockStatement {
		return nil, errorutil.NewError(
			errorutil.ErrorMsgBlockStatementExpected,
			blockNode.Expr(),
		)
	}

	return blockStatement, nil
}
