package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatBlockStatement(
	node *ast.BlockStatement,
	result *strings.Builder,
	depth int,
	shouldIndent bool,
) {
	if shouldIndent {
		f.addWhitespace(result, depth)
	}

	if len(node.Statements) == 0 {
		result.WriteString("{}\n")

		return
	}

	result.WriteString("{\n")

	for _, statement := range node.Statements {
		if statement != nil {
			f.formatNode(statement, result, depth+1)
		}
	}

	f.addWhitespace(result, depth)
	result.WriteString("}\n")
}
