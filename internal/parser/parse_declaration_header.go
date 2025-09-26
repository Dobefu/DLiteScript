package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
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
			ast.Range{
				Start: ast.Position{
					Offset: p.tokenIdx,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: p.tokenIdx,
					Line:   p.line,
					Column: p.column,
				},
			},
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
				ast.Range{
					Start: ast.Position{
						Offset: nextToken.StartPos,
						Line:   p.line,
						Column: p.column,
					},
					End: ast.Position{
						Offset: nextToken.EndPos,
						Line:   p.line,
						Column: p.column,
					},
				},
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
				ast.Range{
					Start: ast.Position{
						Offset: elementTypeToken.StartPos,
						Line:   p.line,
						Column: p.column,
					},
					End: ast.Position{
						Offset: elementTypeToken.EndPos,
						Line:   p.line,
						Column: p.column,
					},
				},
				elementTypeToken.Atom,
			)
		}

		return "[]" + elementTypeToken.Atom, nil
	}

	if !typeToken.IsDataType() {
		return "", errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedToken,
			ast.Range{
				Start: ast.Position{
					Offset: typeToken.StartPos,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: typeToken.EndPos,
					Line:   p.line,
					Column: p.column,
				},
			},
			typeToken.Atom,
		)
	}

	return typeToken.Atom, nil
}
