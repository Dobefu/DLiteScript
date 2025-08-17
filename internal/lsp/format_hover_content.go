package lsp

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func formatHoverContent(node ast.ExprNode, isDebugMode bool) string {
	var output strings.Builder

	if isDebugMode {
		output.WriteString("**🔴 Debug Mode 🔴**\n\n")
	}

	nodeType := getAstNodeLabel(node, isDebugMode)

	// If there's no node type, don't display anything.
	if nodeType == "" {
		return ""
	}

	output.WriteString("**" + nodeType + "**\n\n")

	output.WriteString("```dlitescript\n")
	output.WriteString(node.Expr())
	output.WriteString("\n```")

	return output.String()
}
