package lsp

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/stdlib"
)

// AstNodeInfo represents information about an AST node.
type AstNodeInfo struct {
	Label       string
	Description string
}

func getAstNodeInfo(node ast.ExprNode, isDebugMode bool) *AstNodeInfo {
	switch n := node.(type) {
	case *ast.FunctionCall:
		return getFunctionCallInfo(n)

	case *ast.Identifier:
		return &AstNodeInfo{Label: "Identifier", Description: "An identifier"}

	default:
		if isDebugMode {
			return &AstNodeInfo{
				Label:       "Unknown Node",
				Description: fmt.Sprintf("Unknown Node: %T", node),
			}
		}

		return nil
	}
}

func getFunctionCallInfo(n *ast.FunctionCall) *AstNodeInfo {
	registry := stdlib.GetFunctionRegistry()
	pkg, hasPkg := registry[n.Namespace]
	nodeInfo := &AstNodeInfo{Label: "Function Call", Description: ""}

	if !hasPkg {
		return nodeInfo
	}

	function, hasFunction := pkg.Functions[n.FunctionName]

	if !hasFunction {
		return nodeInfo
	}

	var description strings.Builder

	description.WriteString("\n\n```dlitescript\n")
	description.WriteString(function.Expr())
	description.WriteString("\n```\n\n")
	description.WriteString(function.Description)
	description.WriteString("\n\n")

	if len(function.Parameters) > 0 {
		description.WriteString("**Parameters:**\n")

		for _, param := range function.Parameters {
			description.WriteString("```dlitescript\n")
			description.WriteString(param.Name)
			description.WriteString(" ")
			description.WriteString(param.Type.AsString())
			description.WriteString("\n```\n")
		}
	}

	if len(function.ReturnValues) > 0 {
		description.WriteString("\n**Return Values:**\n")

		for _, ret := range function.ReturnValues {
			description.WriteString("```dlitescript\n")
			description.WriteString(ret.Type.AsString())
			description.WriteString("\n```\n")
		}
	}

	nodeInfo.Description = description.String()

	return nodeInfo
}
