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

	if nextToken.TokenType == token.TokenTypeVar {
		return p.parseVariableDeclarationLoop()
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
		StartPos:         p.tokenIdx,
		EndPos:           p.tokenIdx,
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
		StartPos:         p.tokenIdx,
		EndPos:           p.tokenIdx,
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
		if blockNode != nil {
			return nil, errorutil.NewError(
				errorutil.ErrorMsgBlockStatementExpected,
				blockNode.Expr(),
			)
		}

		return &ast.BlockStatement{
			Statements: []ast.ExprNode{blockNode},
			StartPos:   p.tokenIdx,
			EndPos:     p.tokenIdx,
		}, nil
	}

	return blockStatement, nil
}

func (p *Parser) parseVariableDeclarationLoop() (ast.ExprNode, error) {
	_, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType != token.TokenTypeIdentifier {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnexpectedIdentifier,
			p.tokenIdx,
			nextToken.TokenType,
		)
	}

	varName := nextToken.Atom

	nextToken, err = p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeFrom {
		return p.parseExplicitRangeLoop(varName)
	}

	if nextToken.TokenType == token.TokenTypeTo {
		return p.parseImplicitRangeLoop(varName)
	}

	operatorToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	nextToken, err = p.GetNextToken()

	if err != nil {
		return nil, err
	}

	rightSide, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	condition := &ast.BinaryExpr{
		Left:     &ast.Identifier{Value: varName, StartPos: p.tokenIdx, EndPos: p.tokenIdx + 1},
		Operator: *operatorToken,
		Right:    rightSide,
		StartPos: p.tokenIdx,
		EndPos:   p.tokenIdx,
	}

	loopBody, err := p.parseLoopBody()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        condition,
		Body:             loopBody,
		StartPos:         p.tokenIdx,
		EndPos:           p.tokenIdx,
		DeclaredVariable: varName,
		RangeVariable:    "",
		RangeFrom:        nil,
		RangeTo:          nil,
		IsRange:          false,
	}, nil
}

func (p *Parser) parseExplicitRangeLoop(varName string) (ast.ExprNode, error) {
	_, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	fromExpr, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	toToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if toToken.TokenType != token.TokenTypeTo {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnexpectedToken,
			p.tokenIdx,
			toToken.TokenType,
		)
	}

	loopBody, toExpr, err := p.parseLoopBodyAndToExpr()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        nil,
		Body:             loopBody,
		StartPos:         p.tokenIdx,
		EndPos:           p.tokenIdx,
		DeclaredVariable: varName,
		RangeVariable:    "",
		RangeFrom:        fromExpr,
		RangeTo:          toExpr,
		IsRange:          true,
	}, nil
}

func (p *Parser) parseImplicitRangeLoop(varName string) (ast.ExprNode, error) {
	loopBody, toExpr, err := p.parseLoopBodyAndToExpr()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        nil,
		Body:             loopBody,
		StartPos:         p.tokenIdx,
		EndPos:           p.tokenIdx,
		DeclaredVariable: varName,
		RangeVariable:    "",
		RangeFrom:        &ast.NumberLiteral{Value: "0", StartPos: p.tokenIdx, EndPos: p.tokenIdx + 1},
		RangeTo:          toExpr,
		IsRange:          true,
	}, nil
}

func (p *Parser) parseLoopBodyAndToExpr() (*ast.BlockStatement, ast.ExprNode, error) {
	_, err := p.GetNextToken()

	if err != nil {
		return nil, nil, err
	}

	valueToken, err := p.GetNextToken()

	if err != nil {
		return nil, nil, err
	}

	toExpr, err := p.parseExpr(valueToken, nil, 0, 0)

	if err != nil {
		return nil, nil, err
	}

	loopBody, err := p.parseLoopBody()

	if err != nil {
		return nil, nil, err
	}

	return loopBody, toExpr, nil
}
