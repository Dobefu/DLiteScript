package formatter

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatArrayLiteral(
	node *ast.ArrayLiteral,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	if len(node.Values) == 0 {
		result.WriteString("[]\n")

		return
	}

	var values strings.Builder
	var validValues []ast.ExprNode

	for _, value := range node.Values {
		if value != nil {
			validValues = append(validValues, value)
		}
	}

	if len(validValues) == 0 {
		result.WriteString("[]\n")

		return
	}

	for i, value := range validValues {
		values.WriteString(value.Expr())

		if i < len(validValues)-1 {
			values.WriteString(", ")
		}
	}

	currentIndent := strings.Repeat(f.indentChar, f.indentSize*depth)
	content := values.String()
	totalLength := len(currentIndent) + len(content) + 2

	if totalLength > f.maxLineLength {
		result.WriteString("[\n")

		for i, value := range validValues {
			f.addWhitespace(result, depth+1)
			result.WriteString(value.Expr())

			if i < len(validValues)-1 {
				result.WriteString(",")
			}

			result.WriteString("\n")
		}

		f.addWhitespace(result, depth)
		result.WriteString("]\n")

		return
	}

	fmt.Fprintf(result, "[%s]", values.String())
	result.WriteString("\n")
}
