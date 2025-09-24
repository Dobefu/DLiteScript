// Package formatter provides a formatter for DLiteScript code.
package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

// Formatter represents a formatter for DLiteScript code.
type Formatter struct {
	indentSize    int
	indentChar    string
	maxLineLength int
}

// New creates a new formatter for DLiteScript code.
func New() *Formatter {
	return &Formatter{
		indentSize:    2,
		indentChar:    " ",
		maxLineLength: 80,
	}
}

// Format formats the DLiteScript code.
func (f *Formatter) Format(ast ast.ExprNode) string {
	result := &strings.Builder{}
	f.formatNode(ast, result, 0)

	return result.String()
}
