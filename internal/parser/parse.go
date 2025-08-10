package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() (ast.ExprNode, error) {
	if len(p.tokens) <= 0 {
		return nil, nil
	}

	statements := []ast.ExprNode{}

	for !p.isEOF {
		statement, err := p.parseStatement()

		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)

		err = p.handleStatementEnd()

		if err != nil {
			return nil, err
		}
	}

	if len(statements) == 0 {
		return nil, nil
	}

	if len(statements) == 1 {
		return statements[0], nil
	}

	return &ast.StatementList{
		Statements: statements,
		Pos:        statements[0].Position(),
	}, nil
}

func (p *Parser) handleStatementEnd() error {
	hasNewlines := false

	for !p.isEOF {
		nextToken, err := p.PeekNextToken()

		if err != nil {
			return err
		}

		if nextToken.TokenType == token.TokenTypeNewline {
			_, err = p.GetNextToken()

			if err != nil {
				return err
			}

			hasNewlines = true

			continue
		}

		if !hasNewlines {
			return errorutil.NewErrorAt(
				errorutil.ErrorMsgUnexpectedToken,
				p.tokenIdx,
				nextToken.Atom,
			)
		}

		break
	}

	return nil
}

func (p *Parser) parseStatement() (ast.ExprNode, error) {
	token, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	return p.parseExpr(token, nil, 0, 0)
}

func (p *Parser) handleOptionalNewlines() error {
	for !p.isEOF {
		peek, err := p.PeekNextToken()

		if err != nil {
			return err
		}

		if peek.TokenType == token.TokenTypeNewline {
			if _, err := p.GetNextToken(); err != nil {
				return err
			}

			continue
		}

		break
	}

	return nil
}
