package lsp

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

func parseDocumentToAst(text string) (ast.ExprNode, error) {
	tokenizer := tokenizer.NewTokenizer(text)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		return nil, fmt.Errorf("failed to tokenize file: %s", err.Error())
	}

	parser := parser.NewParser(tokens)
	ast, err := parser.Parse()

	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %s", err.Error())
	}

	return ast, nil
}
