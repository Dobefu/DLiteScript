package formatter

import (
	"fmt"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatFuncDeclarationStatement(
	node *ast.FuncDeclarationStatement,
	result *strings.Builder,
	depth int,
) {
	f.addWhitespace(result, depth)

	argStrings := make([]string, len(node.Args))

	for i, arg := range node.Args {
		argStrings[i] = fmt.Sprintf("%s %s", arg.Name, arg.Type)
	}

	result.WriteString("func ")
	result.WriteString(node.Name)
	result.WriteString("(")
	result.WriteString(strings.Join(argStrings, ", "))
	result.WriteString(")")

	if node.NumReturnValues > 0 {
		result.WriteString(" ")
		result.WriteString(strings.Join(node.ReturnValues, ", "))
	}

	blockStmt, hasBlockStmt := node.Body.(*ast.BlockStatement)

	if hasBlockStmt {
		result.WriteString(" ")

		if len(blockStmt.Statements) == 0 {
			result.WriteString("{}\n")
		} else {
			result.WriteString("{\n")

			for _, statement := range blockStmt.Statements {
				if statement == nil {
					continue
				}

				f.formatNode(statement, result, depth+1)
			}

			f.addWhitespace(result, depth)
			result.WriteString("}\n")
		}
	} else {
		result.WriteString(" {\n")
		f.formatNode(node.Body, result, depth+1)
		f.addWhitespace(result, depth)
		result.WriteString("}\n")
	}
}
