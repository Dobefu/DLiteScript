package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatFunctionCall(
	node *ast.FunctionCall,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	functionName := node.FunctionName

	if node.Namespace != "" {
		functionName = node.Namespace + "." + node.FunctionName
	}

	result.WriteString(functionName)
	result.WriteString("(")

	if len(node.Arguments) > 0 {
		for i, arg := range node.Arguments {
			if arg == nil {
				continue
			}

			if i > 0 {
				result.WriteString(", ")
			}

			result.WriteString(arg.Expr())
		}
	}

	result.WriteString(")\n")
}
