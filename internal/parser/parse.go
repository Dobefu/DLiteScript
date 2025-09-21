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
	if len(p.tokens) == 0 {
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

	return &ast.StatementList{
		Statements: statements,
		StartPos:   statements[0].StartPosition(),
		EndPos:     statements[len(statements)-1].EndPosition(),
	}, nil
}

func (p *Parser) parseStatements(endToken *token.Type) ([]ast.ExprNode, error) {
	statements := []ast.ExprNode{}

	for !p.isEOF {
		comments := p.handleOptionalNewlines()
		statements = append(statements, comments...)

		if endToken != nil {
			nextToken, _ := p.PeekNextToken()

			if nextToken.TokenType == *endToken {
				return statements, nil
			}
		}

		statement, err := p.parseStatement()

		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)

		comments, err = p.handleStatementEnd(endToken)

		if err != nil {
			return nil, err
		}

		statements = append(statements, comments...)
	}

	return statements, nil
}

func (p *Parser) handleStatementEnd(
	endToken *token.Type,
) ([]ast.ExprNode, error) {
	comments := []ast.ExprNode{}
	newlineCount := 0

	for !p.isEOF {
		nextToken, _ := p.PeekNextToken()

		if endToken != nil && nextToken.TokenType == *endToken {
			break
		}

		if nextToken.TokenType == token.TokenTypeNewline {
			_, _ = p.GetNextToken()
			newlineCount++

			continue
		}

		if newlineCount == 0 {
			return comments, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgUnexpectedToken,
				p.tokenIdx,
				nextToken.Atom,
			)
		}

		break
	}

	if newlineCount > 1 {
		comments = append(comments, &ast.NewlineLiteral{
			StartPos: 0,
			EndPos:   0,
		})
	}

	return comments, nil
}

func (p *Parser) parseStatement() (ast.ExprNode, error) {
	nextToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	switch nextToken.TokenType {
	case token.TokenTypeComment:
		return &ast.CommentLiteral{
			Value:    nextToken.Atom,
			StartPos: nextToken.StartPos,
			EndPos:   nextToken.EndPos,
		}, nil

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

	case token.TokenTypeImport:
		return p.parseImportStatement(nextToken)

	case token.TokenTypeLBrace:
		var endToken token.Type = token.TokenTypeRBrace

		return p.parseBlock(&endToken)

	default:
		return p.parseExpr(nextToken, nil, 0, 0)
	}
}

func (p *Parser) handleOptionalNewlines() []ast.ExprNode {
	comments := []ast.ExprNode{}
	newlineCount := 0

	for !p.isEOF {
		peek, _ := p.PeekNextToken()

		if peek.TokenType == token.TokenTypeNewline {
			_, _ = p.GetNextToken()
			newlineCount++

			continue
		}

		break
	}

	if newlineCount > 1 {
		comments = append(comments, &ast.NewlineLiteral{
			StartPos: 0,
			EndPos:   0,
		})
	}

	return comments
}
