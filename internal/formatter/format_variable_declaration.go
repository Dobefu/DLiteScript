package formatter

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatVariableDeclaration(
	node *ast.VariableDeclaration,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	if node.Value == nil {
		fmt.Fprintf(result, "var %s %s\n", node.Name, node.Type)

		return
	}

	varDeclarationPrefix := fmt.Sprintf("var %s %s = ", node.Name, node.Type)
	totalLineLength := node.Range.End.Offset - node.Range.Start.Offset

	arrayLiteral, isArrayLiteral := node.Value.(*ast.ArrayLiteral)

	if isArrayLiteral && totalLineLength > f.maxLineLength {
		result.WriteString(varDeclarationPrefix)
		result.WriteString("[\n")

		for _, value := range arrayLiteral.Values {
			if value == nil {
				continue
			}

			f.addWhitespace(result, depth+1)
			result.WriteString(value.Expr())
			result.WriteString(",\n")
		}

		f.addWhitespace(result, depth)
		result.WriteString("]\n")

		return
	}

	fmt.Fprintf(result, "var %s %s = ", node.Name, node.Type)

	var valueBuilder strings.Builder
	f.formatNode(node.Value, &valueBuilder, 0)
	valueStr := strings.TrimSuffix(valueBuilder.String(), "\n")
	result.WriteString(valueStr)
	result.WriteString("\n")
}
