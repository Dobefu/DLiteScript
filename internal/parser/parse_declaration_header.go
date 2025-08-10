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

	if !nextToken.IsDataType() {
		return "", "", errorutil.NewErrorAt(
			errorutil.ErrorMsgInvalidDataType,
			p.tokenIdx,
			nextToken.Atom,
		)
	}

	varType := nextToken.Atom

	return varName, varType, nil
}
