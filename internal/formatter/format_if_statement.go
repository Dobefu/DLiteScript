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

	if node.ElseBlock != nil {
		if len(node.ThenBlock.Statements) == 0 {
			result.WriteString("{}")
		} else {
			result.WriteString("{\n")

			for _, statement := range node.ThenBlock.Statements {
				if statement != nil {
					f.formatNode(statement, result, depth+1)
				}
			}

			f.addWhitespace(result, depth)
			result.WriteString("}")
		}

		result.WriteString(" else ")
		f.formatBlockStatement(node.ElseBlock, result, depth, false)
	} else {
		f.formatBlockStatement(node.ThenBlock, result, depth, false)
	}
}
