package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatBinaryExpr(
	node *ast.BinaryExpr,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)
	result.WriteString(node.Left.Expr())
	result.WriteString(" ")
	result.WriteString(node.Operator.Atom)
	result.WriteString(" ")
	result.WriteString(node.Right.Expr())
	result.WriteString("\n")
}
