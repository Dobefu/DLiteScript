package lsp

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func formatHoverContent(node ast.ExprNode, isDebugMode bool) string {
	var output strings.Builder

	nodeType := getAstNodeLabel(node, isDebugMode)

	if isDebugMode {
		output.WriteString(fmt.Sprintf("**🔴 Debug Mode** %s\n\n", nodeType))
	}

	// If there's no node type, don't display anything.
	if nodeType == "" {
		return ""
	}

	output.WriteString("```dlitescript\n")
	output.WriteString(node.Expr())
	output.WriteString("\n```")

	return output.String()
}
