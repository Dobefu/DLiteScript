package parser

import (
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseDeclarationHeader() (string, string, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return "", "", err
	}

	if nextToken.TokenType != token.TokenTypeIdentifier {
		return "", "", errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedIdentifier,
			p.tokenIdx,
			nextToken.Atom,
		)
	}

	varName := nextToken.Atom
	nextToken, err = p.GetNextToken()

	if err != nil {
		return "", "", err
	}

	varType, err := p.parseDataType(nextToken)

	if err != nil {
		return "", "", err
	}

	return varName, varType, nil
}

func (p *Parser) parseDataType(typeToken *token.Token) (string, error) {
	if typeToken.TokenType == token.TokenTypeLBracket {
		nextToken, err := p.GetNextToken()

		if err != nil {
			return "", err
		}

		if nextToken.TokenType != token.TokenTypeRBracket {
			return "", errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgUnexpectedToken,
				nextToken.StartPos,
				nextToken.Atom,
			)
		}

		elementTypeToken, err := p.GetNextToken()

		if err != nil {
			return "", err
		}

		if !elementTypeToken.IsDataType() {
			return "", errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgUnexpectedToken,
				elementTypeToken.StartPos,
				elementTypeToken.Atom,
			)
		}

		return "[]" + elementTypeToken.Atom, nil
	}

	if !typeToken.IsDataType() {
		return "", errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedToken,
			typeToken.StartPos,
			typeToken.Atom,
		)
	}

	return typeToken.Atom, nil
}
