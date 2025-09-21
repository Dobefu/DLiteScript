package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// parseImportStatement parses an import statement.
func (p *Parser) parseImportStatement() (ast.ExprNode, error) {
	importToken := p.tokens[p.tokenIdx-1]
	startPos := importToken.StartPos

	pathToken, err := p.GetNextToken()

	if err != nil {
		return nil, err
	}

	if pathToken.TokenType != token.TokenTypeString {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedToken,
			p.tokenIdx-1,
			pathToken.Atom,
		)
	}

	return &ast.ImportStatement{
		Path: &ast.StringLiteral{
			Value:    pathToken.Atom,
			StartPos: pathToken.StartPos,
			EndPos:   pathToken.EndPos,
		},
		Namespace: pathToken.Atom,
		StartPos:  startPos,
		EndPos:    pathToken.EndPos,
	}, nil
}
