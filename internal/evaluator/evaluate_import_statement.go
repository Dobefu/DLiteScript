package evaluator

import (
	"fmt"
	"maps"
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

	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	if ext != "" {
		filename = filename[:len(filename)-len(ext)]
	}

	namespace := filename

	if node.Alias != "" {
		namespace = node.Alias
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

	e.namespaceFunctions[namespace] = make(map[string]*ast.FuncDeclarationStatement)

	importEvaluator := NewEvaluator(e.outFile)
	importEvaluator.SetCurrentFilePath(resolvedPath)

	_, err = importEvaluator.Evaluate(importedAST)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to evaluate imported file '%s': %s",
			path, err.Error(),
		)
	}

	if namespace == "_" {
		maps.Copy(e.userFunctions, importEvaluator.userFunctions)
		maps.Copy(e.outerScope, importEvaluator.outerScope)
	} else {
		maps.Copy(e.namespaceFunctions[namespace], importEvaluator.userFunctions)

		for varName, scopedValue := range importEvaluator.outerScope {
			e.outerScope[fmt.Sprintf("%s.%s", namespace, varName)] = scopedValue
		}
	}

	return controlflow.NewRegularResult(datavalue.Null()), nil
}
