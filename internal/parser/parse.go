package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() (ast.ExprNode, error) {
	if len(p.tokens) <= 0 {
		return nil, errorutil.NewErrorAt(errorutil.ErrorMsgEmptyExpression, 0)
	}

	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	ast, err := p.parseExpr(nextToken, nil, 0, 0)

	if err != nil {
		return nil, err
	}

	err = p.checkTrailingTokens()

	if err != nil {
		return nil, err
	}

	return ast, nil
}

func (p *Parser) checkTrailingTokens() error {
	for !p.isEOF {
		peek, err := p.PeekNextToken()

		if err != nil {
			return err
		}

		if peek.TokenType != token.TokenTypeNewline {
			return errorutil.NewErrorAt(
				errorutil.ErrorMsgUnexpectedToken,
				p.tokenIdx,
				peek.Atom,
			)
		}

		if _, err := p.GetNextToken(); err != nil {
			return err
		}
	}

	return nil
}
