package parser

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// parseImportStatement parses an import statement.
func (p *Parser) parseImportStatement(
	nextToken *token.Token,
) (ast.ExprNode, error) {
	if nextToken == nil {
		return nil, errorutil.NewErrorAt(
			errorutil.StageParse,
			errorutil.ErrorMsgUnexpectedEOF,
			p.tokenIdx,
		)
	}

	importToken := nextToken
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

	importStmt := &ast.ImportStatement{
		Path: &ast.StringLiteral{
			Value:    pathToken.Atom,
			StartPos: pathToken.StartPos,
			EndPos:   pathToken.EndPos,
		},
		Namespace: pathToken.Atom,
		Alias:     "",
		StartPos:  startPos,
		EndPos:    pathToken.EndPos,
	}

	nextToken, err = p.PeekNextToken()

	if err != nil {
		return importStmt, nil
	}

	if nextToken.TokenType == token.TokenTypeAs {
		_, _ = p.GetNextToken()
		aliasToken, err := p.GetNextToken()

		if err != nil {
			return nil, err
		}

		if aliasToken.TokenType != token.TokenTypeIdentifier {
			return nil, errorutil.NewErrorAt(
				errorutil.StageParse,
				errorutil.ErrorMsgUnexpectedToken,
				p.tokenIdx-1,
				aliasToken.Atom,
			)
		}

		importStmt.Alias = aliasToken.Atom
		importStmt.EndPos = aliasToken.EndPos
	}

	return importStmt, nil
}
