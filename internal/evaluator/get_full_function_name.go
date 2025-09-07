package evaluator

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func getFullFunctionName(fc *ast.FunctionCall) string {
	if fc.Namespace == "" {
		return fc.FunctionName
	}

	return fmt.Sprintf("%s.%s", fc.Namespace, fc.FunctionName)
}
