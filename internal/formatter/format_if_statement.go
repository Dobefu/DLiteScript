package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatIfStatement(
	node *ast.IfStatement,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)
	result.WriteString("if ")
	result.WriteString(node.Condition.Expr())
	result.WriteString(" ")
	f.formatBlockStatement(node.ThenBlock, result, depth, false)

	if node.ElseBlock != nil {
		f.addWhitespace(result, depth)
		result.WriteString("else ")
		f.formatBlockStatement(node.ElseBlock, result, depth, false)
	}
}
