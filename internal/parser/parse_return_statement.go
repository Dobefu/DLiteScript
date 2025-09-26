package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseReturnStatement() (ast.ExprNode, error) {
	startPos := p.GetCurrentPosition()
	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, err
	}

	if p.isReturnValueTerminator(nextToken) {
		return &ast.ReturnStatement{
			Values:    []ast.ExprNode{},
			NumValues: 0,
			Range: ast.Range{
				Start: startPos,
				End:   p.GetCurrentPosition(),
			},
		}, nil
	}

	returnValues, err := p.parseReturnValues()

	if err != nil {
		return nil, err
	}

	return &ast.ReturnStatement{
		Values:    returnValues,
		NumValues: len(returnValues),
		Range: ast.Range{
			Start: startPos,
			End:   p.GetCurrentPosition(),
		},
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
					ast.Range{
						Start: ast.Position{
							Offset: p.tokenIdx + 1,
							Line:   p.line,
							Column: p.column,
						},
						End: ast.Position{
							Offset: p.tokenIdx + 1,
							Line:   p.line,
							Column: p.column,
						},
					},
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
