package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseForStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeLBrace {
		return p.parseInfiniteLoop(startPos)
	}

	if nextToken.TokenType == token.TokenTypeVar {
		return p.parseVariableDeclarationLoop(startPos)
	}

	return p.parseLoop(startPos)
}

func (p *Parser) parseInfiniteLoop(startPos int) (ast.ExprNode, error) {
	loopBody, err := p.parseLoopBody()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        nil,
		Body:             loopBody,
		StartPos:         startPos,
		EndPos:           loopBody.EndPosition(),
		DeclaredVariable: "",
		RangeVariable:    "",
		RangeFrom:        nil,
		RangeTo:          nil,
		IsRange:          false,
	}, nil
}

func (p *Parser) parseLoop(startPos int) (ast.ExprNode, error) {
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
		StartPos:         startPos,
		EndPos:           loopBody.EndPosition(),
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

		blockStartPos := p.GetCurrentCharPos()

		return &ast.BlockStatement{
			Statements: []ast.ExprNode{},
			StartPos:   blockStartPos,
			EndPos:     blockStartPos,
		}, nil
	}

	return blockStatement, nil
}

func (p *Parser) parseVariableDeclarationLoop(
	startPos int,
) (ast.ExprNode, error) {
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
		return p.parseExplicitRangeLoop(startPos, varName)
	}

	if nextToken.TokenType == token.TokenTypeTo {
		return p.parseImplicitRangeLoop(startPos, varName)
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

	identifierStartPos := p.GetCurrentCharPos()
	condition := &ast.BinaryExpr{
		Left: &ast.Identifier{
			Value:    varName,
			StartPos: identifierStartPos,
			EndPos:   identifierStartPos + len(varName),
		},
		Operator: *operatorToken,
		Right:    rightSide,
		StartPos: identifierStartPos,
		EndPos:   rightSide.EndPosition(),
	}

	loopBody, err := p.parseLoopBody()

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		Condition:        condition,
		Body:             loopBody,
		StartPos:         startPos,
		EndPos:           loopBody.EndPosition(),
		DeclaredVariable: varName,
		RangeVariable:    "",
		RangeFrom:        nil,
		RangeTo:          nil,
		IsRange:          false,
	}, nil
}

func (p *Parser) parseExplicitRangeLoop(
	startPos int,
	varName string,
) (ast.ExprNode, error) {
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
		StartPos:         startPos,
		EndPos:           loopBody.EndPosition(),
		DeclaredVariable: varName,
		RangeVariable:    "",
		RangeFrom:        fromExpr,
		RangeTo:          toExpr,
		IsRange:          true,
	}, nil
}

func (p *Parser) parseImplicitRangeLoop(
	startPos int,
	varName string,
) (ast.ExprNode, error) {
	loopBody, toExpr, err := p.parseLoopBodyAndToExpr()

	if err != nil {
		return nil, err
	}

	zeroLiteralStartPos := p.GetCurrentCharPos()

	return &ast.ForStatement{
		Condition:        nil,
		Body:             loopBody,
		StartPos:         startPos,
		EndPos:           loopBody.EndPosition(),
		DeclaredVariable: varName,
		RangeVariable:    "",
		RangeFrom: &ast.NumberLiteral{
			Value:    "0",
			StartPos: zeroLiteralStartPos,
			EndPos:   zeroLiteralStartPos + 1,
		},
		RangeTo: toExpr,
		IsRange: true,
	}, nil
}

func (p *Parser) parseLoopBodyAndToExpr() (
	*ast.BlockStatement,
	ast.ExprNode,
	error,
) {
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
