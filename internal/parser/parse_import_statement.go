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
			ast.Range{
				Start: ast.Position{
					Offset: p.tokenIdx - 1,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: p.tokenIdx - 1,
					Line:   p.line,
					Column: p.column,
				},
			},
			pathToken.Atom,
		)
	}

	importStmt := &ast.ImportStatement{
		Path: &ast.StringLiteral{
			Value: pathToken.Atom,
			Range: ast.Range{
				Start: ast.Position{
					Offset: pathToken.StartPos,
					Line:   p.line,
					Column: p.column,
				},
				End: ast.Position{
					Offset: pathToken.EndPos,
					Line:   p.line,
					Column: p.column + (pathToken.EndPos - pathToken.StartPos),
				},
			},
		},
		Namespace: pathToken.Atom,
		Alias:     "",
		Range: ast.Range{
			Start: ast.Position{
				Offset: startPos,
				Line:   p.line,
				Column: p.column,
			},
			End: ast.Position{
				Offset: pathToken.EndPos,
				Line:   p.line,
				Column: p.column + (pathToken.EndPos - pathToken.StartPos),
			},
		},
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
				ast.Range{
					Start: ast.Position{
						Offset: p.tokenIdx - 1,
						Line:   p.line,
						Column: p.column,
					},
					End: ast.Position{
						Offset: p.tokenIdx - 1,
						Line:   p.line,
						Column: p.column,
					},
				},
				aliasToken.Atom,
			)
		}

		importStmt.Alias = aliasToken.Atom
		importStmt.Range.End = ast.Position{
			Offset: aliasToken.EndPos,
			Line:   p.line,
			Column: p.column + (aliasToken.EndPos - aliasToken.StartPos),
		}
	}

	return importStmt, nil
}
