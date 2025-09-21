package formatter

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatForStatement(
	node *ast.ForStatement,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)
	result.WriteString("for")

	if !node.IsRange {
		if node.Condition != nil {
			result.WriteString(" ")
			result.WriteString(node.Condition.Expr())
		}

		result.WriteString(" ")
		f.formatBlockStatement(node.Body, result, depth, false)

		return
	}

	if node.RangeFrom != nil && node.RangeTo != nil {
		if node.DeclaredVariable != "" {
			fmt.Fprintf(
				result,
				" var %s from %s to %s",
				node.DeclaredVariable,
				node.RangeFrom.Expr(),
				node.RangeTo.Expr(),
			)
		} else {
			fmt.Fprintf(
				result,
				" from %s to %s",
				node.RangeFrom.Expr(),
				node.RangeTo.Expr(),
			)
		}

		result.WriteString(" ")
		f.formatBlockStatement(node.Body, result, depth, false)

		return
	}

	if node.RangeVariable != "" {
		fmt.Fprintf(result, " %s", node.RangeVariable)
	}

	result.WriteString(" ")
	f.formatNode(node.Body, result, depth)
}
