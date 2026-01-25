package lsp

import (
	"errors"
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

var errFailedToTokenizeFile = errors.New("failed to tokenize file")
var errFailedToParseFile = errors.New("failed to parse file")

func parseDocumentToAst(text string) (ast.ExprNode, error) {
	tokenizer := tokenizer.NewTokenizer(text)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errFailedToTokenizeFile, err)
	}

	parser := parser.NewParser(tokens)
	ast, err := parser.Parse()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errFailedToParseFile, err)
	}

	return ast, nil
}
