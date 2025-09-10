package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseIfStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	thenBlock, err := p.parseThenBlock(token.TokenTypeRBrace)

	if err != nil {
		return nil, err
	}

	endPos := thenBlock.EndPosition()
	var elseBlock *ast.BlockStatement

	if !p.isEOF {
		nextToken, err = p.PeekNextToken()

		if err != nil {
			return nil, err
		}

		if nextToken.TokenType == token.TokenTypeElse {
			elseBlock, err = p.handleElseBlock()

			if err != nil {
				return nil, err
			}

			endPos = elseBlock.EndPosition()
		}
	}

	return &ast.IfStatement{
		Condition: expr,
		ThenBlock: thenBlock,
		ElseBlock: elseBlock,
		StartPos:  startPos,
		EndPos:    endPos,
	}, nil
}

func (p *Parser) handleElseBlock() (*ast.BlockStatement, error) {
	_, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeIf {
		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}

		nestedExpr, err := p.parseIfStatement()

		if err != nil {
			return nil, err
		}

		return &ast.BlockStatement{
			Statements: []ast.ExprNode{nestedExpr},
			StartPos:   nestedExpr.StartPosition(),
			EndPos:     nestedExpr.EndPosition(),
		}, nil
	}

	elseBlock, err := p.parseElseBlock(token.TokenTypeRBrace)

	if err != nil {
		return nil, err
	}

	return elseBlock, nil
}

func (p *Parser) parseThenBlock(endToken token.Type) (*ast.BlockStatement, error) {
	_, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	thenBlock, err := p.parseBlock(&endToken)

	if err != nil {
		return nil, err
	}

	if thenBlock == nil {
		return &ast.BlockStatement{
			Statements: []ast.ExprNode{},
			StartPos:   p.GetCurrentCharPos(),
			EndPos:     p.GetCurrentCharPos(),
		}, nil
	}

	blockStatement, isBlockStatement := thenBlock.(*ast.BlockStatement)

	if !isBlockStatement {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgBlockStatementExpected,
			p.tokenIdx,
			thenBlock.Expr(),
		)
	}

	return blockStatement, nil
}

func (p *Parser) parseElseBlock(endToken token.Type) (*ast.BlockStatement, error) {
	_, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	elseBlock, err := p.parseBlock(&endToken)

	if err != nil {
		return nil, err
	}

	if elseBlock == nil {
		return &ast.BlockStatement{
			Statements: []ast.ExprNode{},
			StartPos:   p.GetCurrentCharPos(),
			EndPos:     p.GetCurrentCharPos(),
		}, nil
	}

	blockStatement, isBlockStatement := elseBlock.(*ast.BlockStatement)

	if !isBlockStatement {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgBlockStatementExpected,
			p.tokenIdx,
			elseBlock.Expr(),
		)
	}

	return blockStatement, nil
}
