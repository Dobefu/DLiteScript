package formatter

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatPrefixExpr(
	node *ast.PrefixExpr,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	fmt.Fprintf(result, "%s%s", node.Operator.Atom, node.Operand.Expr())
	result.WriteString("\n")
}
