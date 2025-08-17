package lsp

import "github.com/Dobefu/DLiteScript/internal/ast"

func getAstNodeLabel(node ast.ExprNode) string {
	switch node.(type) {
	case *ast.FunctionCall:
		return "Function Call"

	default:
		return ""
	}
}
