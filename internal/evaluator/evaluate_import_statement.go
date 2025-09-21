package evaluator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

// evaluateImportStatement evaluates an import statement.
func (e *Evaluator) evaluateImportStatement(
	node *ast.ImportStatement,
) (*controlflow.EvaluationResult, error) {
	path := node.Path.Value
	resolvedPath := path

	if !filepath.IsAbs(path) && e.currentFilePath != "" {
		currentDir := filepath.Dir(e.currentFilePath)
		resolvedPath = filepath.Join(currentDir, path)
	}

	fileContent, err := os.ReadFile(filepath.Clean(resolvedPath))

	if err != nil {
		return nil, fmt.Errorf(
			"failed to read imported file '%s': %s",
			path,
			err.Error(),
		)
	}

	t := tokenizer.NewTokenizer(string(fileContent))
	tokens, err := t.Tokenize()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to tokenize imported file '%s': %s",
			path, err.Error(),
		)
	}

	p := parser.NewParser(tokens)
	importedAST, err := p.Parse()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse imported file '%s': %s",
			path, err.Error(),
		)
	}

	_, err = e.Evaluate(importedAST)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to evaluate imported file '%s': %s",
			path, err.Error(),
		)
	}

	return controlflow.NewRegularResult(datavalue.Null()), nil
}
