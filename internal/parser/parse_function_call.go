package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (p *Parser) parseFunctionCall(
	functionName string,
	functionPos int,
	recursionDepth int,
) (ast.ExprNode, error) {
	lparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if lparenToken.TokenType != token.TokenTypeLParen {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgExpectedOpenParen,
			p.tokenIdx,
			lparenToken.Atom,
		)
	}

	nextToken, err := p.PeekNextToken()

	if err != nil {
		return nil, errorutil.NewErrorAt(errorutil.ErrorMsgParenNotClosedAtEOF, p.tokenIdx)
	}

	var args []ast.ExprNode

	if nextToken.TokenType != token.TokenTypeRParen {
		args, err = p.parseFunctionCallArguments(recursionDepth + 1)

		if err != nil {
			return nil, err
		}
	}

	rparenToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if rparenToken.TokenType != token.TokenTypeRParen {
		return nil, errorutil.NewErrorAt(errorutil.ErrorMsgParenNotClosedAtEOF, p.tokenIdx)
	}

	return &ast.FunctionCall{
		FunctionName: functionName,
		Arguments:    args,
		Pos:          functionPos,
	}, nil
}

func (p *Parser) parseFunctionCallArguments(
	recursionDepth int,
) ([]ast.ExprNode, error) {
	// Pre-allocate the size of the slice, to reduce allocs.
	args := make([]ast.ExprNode, 0, 4)

	for {
		argToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		arg, err := p.parseExpr(argToken, nil, 0, recursionDepth+1)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		nextToken, err := p.PeekNextToken()

		if err != nil {
			return nil, errorutil.NewErrorAt(errorutil.ErrorMsgParenNotClosedAtEOF, p.tokenIdx)
		}

		if nextToken.TokenType == token.TokenTypeRParen {
			break
		}

		if nextToken.TokenType != token.TokenTypeComma {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgUnexpectedToken,
				p.tokenIdx,
				nextToken.Atom,
			)
		}

		_, err = p.GetNextToken()

		if err != nil {
			return nil, err
		}
	}

	return args, nil
}
