package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() (ast.ExprNode, error) {
	return p.parseBlock(nil)
}

func (p *Parser) parseBlock(endToken *token.Type) (ast.ExprNode, error) {
	if len(p.tokens) <= 0 {
		return nil, nil
	}

	statements, err := p.parseStatements(endToken)

	if err != nil {
		return nil, err
	}

	if endToken != nil {
		_, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}
	}

	if len(statements) == 0 {
		return nil, nil
	}

	if endToken != nil {
		return &ast.BlockStatement{
			Statements: statements,
			StartPos:   statements[0].StartPosition(),
			EndPos:     statements[len(statements)-1].EndPosition(),
		}, nil
	}

	if len(statements) == 1 {
		return statements[0], nil
	}

	if len(statements) == 0 {
		return nil, nil
	}

	return &ast.StatementList{
		Statements: statements,
		StartPos:   statements[0].StartPosition(),
		EndPos:     statements[len(statements)-1].EndPosition(),
	}, nil
}

func (p *Parser) parseStatements(endToken *token.Type) ([]ast.ExprNode, error) {
	statements := []ast.ExprNode{}

	for !p.isEOF {
		err := p.handleOptionalNewlines()

		if err != nil {
			return nil, err
		}

		if endToken != nil {
			nextToken, err := p.PeekNextToken()

			if err != nil {
				return nil, err
			}

			if nextToken.TokenType == *endToken {
				return statements, nil
			}
		}

		statement, err := p.parseStatement()

		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)

		err = p.handleStatementEnd(endToken)

		if err != nil {
			return nil, err
		}
	}

	return statements, nil
}

func (p *Parser) handleStatementEnd(endToken *token.Type) error {
	hasNewlines := false

	for !p.isEOF {
		nextToken, err := p.PeekNextToken()

		if err != nil {
			return err
		}

		if endToken != nil && nextToken.TokenType == *endToken {
			break
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
				errorutil.StageParsing,
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
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	switch nextToken.TokenType {
	case token.TokenTypeVar:
		return p.parseVariableDeclaration()

	case token.TokenTypeConst:
		return p.parseConstantDeclaration()

	case token.TokenTypeIf:
		return p.parseIfStatement()

	case token.TokenTypeFor:
		return p.parseForStatement()

	case token.TokenTypeBreak:
		return p.parseBreakStatement()

	case token.TokenTypeContinue:
		return p.parseContinueStatement()

	case token.TokenTypeFunc:
		return p.parseFunctionDeclaration()

	case token.TokenTypeReturn:
		return p.parseReturnStatement()

	case token.TokenTypeLBrace:
		var endToken token.Type = token.TokenTypeRBrace

		return p.parseBlock(&endToken)

	default:
		return p.parseExpr(nextToken, nil, 0, 0)
	}
}

func (p *Parser) handleOptionalNewlines() error {
	for !p.isEOF {
		peek, err := p.PeekNextToken()

		if err != nil {
			return err
		}

		if peek.TokenType == token.TokenTypeNewline {
			_, err := p.GetNextToken()

			if err != nil {
				return err
			}

			continue
		}

		break
	}

	return nil
}
