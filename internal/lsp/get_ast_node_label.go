package lsp

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func getAstNodeLabel(node ast.ExprNode, isDebugMode bool) string {
	switch node.(type) {
	case *ast.FunctionCall:
		return "Function Call"

	case *ast.Identifier:
		return "Identifier"

	default:
		if isDebugMode {
			return fmt.Sprintf("Unknown Node: %T", node)
		}

		return ""
	}
}
