package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatStatementList(
	node *ast.StatementList,
	result *strings.Builder,
	depth int,
) {
	for _, statement := range node.Statements {
		if statement == nil {
			continue
		}

		f.formatNode(statement, result, depth)
	}
}
