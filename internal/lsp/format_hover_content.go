package lsp

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func formatHoverContent(
	node ast.ExprNode,
	isDebugMode bool,
) string {
	var output strings.Builder
	nodeType := getAstNodeInfo(node, isDebugMode)

	if isDebugMode {
		output.WriteString("**🔴 Debug Mode** | ")
		output.WriteString(fmt.Sprintf("**%s**", nodeType.Label))
		output.WriteString("\n\n---\n\n")
	}

	// If there's no node type, don't display anything.
	if nodeType == nil {
		return ""
	}

	output.WriteString(nodeType.Description)

	return output.String()
}
