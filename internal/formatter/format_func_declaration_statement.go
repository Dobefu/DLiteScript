package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatFuncDeclarationStatement(
	node *ast.FuncDeclarationStatement,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)
	result.WriteString(node.Expr())
	result.WriteString("\n")
}
