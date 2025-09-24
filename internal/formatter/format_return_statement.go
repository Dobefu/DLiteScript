package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatReturnStatement(
	node *ast.ReturnStatement,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	if len(node.Values) == 0 {
		result.WriteString("return\n")

		return
	}

	if len(node.Values) == 1 {
		totalLineLength := node.Range.End.Offset - node.Range.Start.Offset
		arrayLiteral, isArrayLiteral := node.Values[0].(*ast.ArrayLiteral)

		if isArrayLiteral && totalLineLength > f.maxLineLength {
			result.WriteString("return [\n")

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
	}

	result.WriteString(node.Expr())
	result.WriteString("\n")
}
