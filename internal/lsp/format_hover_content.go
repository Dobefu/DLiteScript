package lsp

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func formatHoverContent(node ast.ExprNode) string {
	var output strings.Builder

	nodeType := getAstNodeLabel(node)

	// If there's no node type, don't display anything.
	if nodeType == "" {
		return ""
	}

	output.WriteString("**" + nodeType + "**\n\n")

	output.WriteString("```dlitescript\n")
	output.WriteString(node.Expr())
	output.WriteString("```")

	return output.String()
}
