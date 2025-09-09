package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseFunctionDeclaration() (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType != token.TokenTypeIdentifier {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnexpectedToken,
			nextToken.StartPos,
			nextToken.Atom,
		)
	}

	funcName := nextToken.Atom
	startPos := nextToken.StartPos
	args, err := p.getArgs()

	if err != nil {
		return nil, err
	}

	returnTypes, err := p.getReturnTypes()

	if err != nil {
		return nil, err
	}

	var endToken token.Type = token.TokenTypeRBrace
	leftBrace, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if leftBrace.TokenType != token.TokenTypeLBrace {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnexpectedToken,
			leftBrace.StartPos,
			leftBrace.Atom,
		)
	}

	body, err := p.parseBlock(&endToken)

	if err != nil {
		return nil, err
	}

	return &ast.FuncDeclarationStatement{
		Name:            funcName,
		Args:            args,
		Body:            body,
		ReturnValues:    returnTypes,
		NumReturnValues: len(returnTypes),
		StartPos:        startPos,
		EndPos:          p.GetCurrentCharPos(),
	}, nil
}

func (p *Parser) getArgs() ([]ast.FuncParameter, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType != token.TokenTypeLParen {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnexpectedToken,
			nextToken.StartPos,
			nextToken.Atom,
		)
	}

	return p.parseArguments()
}

func (p *Parser) parseArguments() ([]ast.FuncParameter, error) {
	args := make([]ast.FuncParameter, 0)

	for !p.isEOF {
		nextToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		if nextToken.TokenType == token.TokenTypeRParen {
			break
		}

		if nextToken.TokenType == token.TokenTypeComma {
			continue
		}

		arg, err := p.parseFunctionArgument(nextToken)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	return args, nil
}

func (p *Parser) parseFunctionArgument(nameToken *token.Token) (ast.FuncParameter, error) {
	if nameToken.TokenType != token.TokenTypeIdentifier {
		return ast.FuncParameter{}, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnexpectedToken,
			nameToken.StartPos,
			nameToken.Atom,
		)
	}

	typeToken, err := p.GetNextToken()

	if err != nil {
		return ast.FuncParameter{}, err
	}

	dataType, err := p.parseDataType(typeToken)

	if err != nil {
		return ast.FuncParameter{}, err
	}

	return ast.FuncParameter{
		Name: nameToken.Atom,
		Type: dataType,
	}, nil
}

func (p *Parser) getReturnTypes() ([]string, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType == token.TokenTypeLParen {
		returnTypes, err := p.parseReturnTypes(token.TokenTypeRParen)

		if err != nil {
			return nil, err
		}

		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}

		return returnTypes, nil
	}

	dataType, err := p.parseDataType(nextToken)

	if err != nil {
		return nil, err
	}

	returnTypes := make([]string, 0)
	returnTypes = append(returnTypes, dataType)

	return p.parseReturnTypes(token.TokenTypeLBrace, returnTypes...)
}

func (p *Parser) parseReturnTypes(
	endToken token.Type,
	initialTypes ...string,
) ([]string, error) {
	returnTypes := make([]string, 0, len(initialTypes)+5)
	returnTypes = append(returnTypes, initialTypes...)

	for !p.isEOF {
		if p.tokenIdx >= p.tokenLen {
			break
		}

		peekToken := p.tokens[p.tokenIdx]

		if peekToken.TokenType == endToken {
			break
		}

		if peekToken.TokenType == token.TokenTypeComma {
			_, err := p.GetNextToken()

			if err != nil {
				return nil, err
			}

			continue
		}

		nextToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		if nextToken.TokenType != token.TokenTypeIdentifier && !nextToken.IsDataType() {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgUnexpectedToken,
				nextToken.StartPos,
				nextToken.Atom,
			)
		}

		returnTypes = append(returnTypes, nextToken.Atom)
	}

	return returnTypes, nil
}
