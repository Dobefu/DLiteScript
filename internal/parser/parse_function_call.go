package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseFunctionCall(
	namespace string,
	functionName string,
	recursionDepth int,
) (ast.ExprNode, error) {
	// The function name has already been consumed,
	// so we should get the start position from the previous token.
	startCharPos := p.tokens[p.tokenIdx-1].StartPos
	lparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if lparenToken.TokenType != token.TokenTypeLParen {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgExpectedOpenParen,
			p.tokenIdx,
			lparenToken.Atom,
		)
	}

	err = p.handleOptionalNewlines()

	if err != nil {
		return nil, err
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgParenNotClosedAtEOF,
			p.tokenIdx,
		)
	}

	var args []ast.ExprNode

	if nextToken.TokenType != token.TokenTypeRParen {
		args, err = p.parseFunctionCallArguments(recursionDepth + 1)

		if err != nil {
			return nil, err
		}
	}

	nextToken, err = p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if nextToken.TokenType != token.TokenTypeRParen {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgParenNotClosedAtEOF,
			p.tokenIdx,
		)
	}

	return &ast.FunctionCall{
		Namespace:    namespace,
		FunctionName: functionName,
		Arguments:    args,
		StartPos:     startCharPos,
		EndPos:       p.GetCurrentCharPos(),
	}, nil
}

func (p *Parser) parseFunctionCallArguments(
	recursionDepth int,
) ([]ast.ExprNode, error) {
	// Pre-allocate the size of the slice, to reduce allocs.
	args := make([]ast.ExprNode, 0, 4)

	for {
		arg, err := p.parseArgument(recursionDepth)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		hasArgumentEnded, err := p.isEndOfArguments()

		if err != nil {
			return nil, err
		}

		if hasArgumentEnded {
			break
		}

		if err := p.consumeComma(); err != nil {
			return nil, err
		}

		isTrailingComma, err := p.isTrailingComma()

		if err != nil {
			return nil, err
		}

		if isTrailingComma {
			break
		}
	}

	return args, nil
}

func (p *Parser) isEndOfArguments() (bool, error) {
	if err := p.handleOptionalNewlines(); err != nil {
		return false, err
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return false, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgParenNotClosedAtEOF,
			p.tokenIdx,
		)
	}

	return nextToken.TokenType == token.TokenTypeRParen, nil
}

func (p *Parser) consumeComma() error {
	if err := p.handleOptionalNewlines(); err != nil {
		return err
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgParenNotClosedAtEOF,
			p.tokenIdx,
		)
	}

	if nextToken.TokenType != token.TokenTypeComma {
		return errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgUnexpectedToken,
			p.tokenIdx,
			nextToken.Atom,
		)
	}

	_, err = p.GetNextToken()

	return err
}

func (p *Parser) isTrailingComma() (bool, error) {
	if err := p.handleOptionalNewlines(); err != nil {
		return false, err
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return false, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgUnexpectedEOF,
			p.tokenIdx,
		)
	}

	return nextToken.TokenType == token.TokenTypeRParen, nil
}

func (p *Parser) parseArgument(recursionDepth int) (ast.ExprNode, error) {
	argToken, err := p.GetNextToken()

	if err != nil {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParsing,
			errorutil.ErrorMsgUnexpectedEOF,
			p.tokenIdx,
		)
	}

	return p.parseExpr(argToken, nil, 0, recursionDepth+1)
}
