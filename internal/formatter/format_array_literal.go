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

	for i, value := range node.Values {
		if value == nil {
			continue
		}

		values.WriteString(value.Expr())

		if i < len(node.Values)-1 {
			values.WriteString(", ")
		}
	}

	fmt.Fprintf(result, "[%s]", values.String())
	result.WriteString("\n")
}
