package formatter

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatIndexAssignmentStatement(
	node *ast.IndexAssignmentStatement,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	totalLineLength := node.Range.End.Offset - node.Range.Start.Offset
	arrayLiteral, isArrayLiteral := node.Right.(*ast.ArrayLiteral)

	if isArrayLiteral && totalLineLength > f.maxLineLength {
		fmt.Fprintf(result, "%s[%s] = [\n", node.Array.Expr(), node.Index.Expr())

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

	result.WriteString(node.Expr())
	result.WriteString("\n")
}
