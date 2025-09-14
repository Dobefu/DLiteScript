package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseReturnStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentCharPos()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if p.isReturnValueTerminator(nextToken) {
		return &ast.ReturnStatement{
			Values:    []ast.ExprNode{},
			NumValues: 0,
			StartPos:  startPos,
			EndPos:    p.GetCurrentCharPos(),
		}, nil
	}

	returnValues, err := p.parseReturnValues()

	if err != nil {
		return nil, err
	}

	return &ast.ReturnStatement{
		Values:    returnValues,
		NumValues: len(returnValues),
		StartPos:  startPos,
		EndPos:    p.GetCurrentCharPos(),
	}, nil
}

func (p *Parser) parseReturnValues() ([]ast.ExprNode, error) {
	var returnValues []ast.ExprNode

	for !p.isEOF {
		nextToken, _ := p.PeekNextToken()

		if p.isReturnValueTerminator(nextToken) {
			break
		}

		if nextToken.TokenType == token.TokenTypeComma {
			if len(returnValues) == 0 {
				return nil, errorutil.NewErrorAt(
					errorutil.StageParse,
					errorutil.ErrorMsgUnexpectedToken,
					p.tokenIdx+1,
					nextToken.Atom,
				)
			}

			_ = p.consumeReturnComma()

			continue
		}

		value, err := p.parseNextReturnValue()

		if err != nil {
			return nil, err
		}

		returnValues = append(returnValues, value)
	}

	return returnValues, nil
}

func (p *Parser) isReturnValueTerminator(t *token.Token) bool {
	return t.TokenType == token.TokenTypeNewline || t.TokenType == token.TokenTypeRBrace
}

func (p *Parser) consumeReturnComma() error {
	_, err := p.GetNextToken()

	return err
}

func (p *Parser) parseNextReturnValue() (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	return p.parseExpr(nextToken, nil, 0, 0)
}
