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
		f.formatSimpleForCondition(node, result)
		result.WriteString(" ")
		f.formatBlockStatement(node.Body, result, depth, false)

		return
	}

	if node.RangeTo != nil {
		f.formatRangeForCondition(node, result)
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

func (f *Formatter) formatSimpleForCondition(
	node *ast.ForStatement,
	result *strings.Builder,
) {
	if node.Condition == nil {
		return
	}

	result.WriteString(" ")

	if node.DeclaredVariable != "" {
		f.formatForConditionWithVariable(node, result)

		return
	}

	f.formatForConditionWithoutVariable(node, result)
}

func (f *Formatter) formatForConditionWithVariable(
	node *ast.ForStatement,
	result *strings.Builder,
) {
	binaryExpr, isBinaryExpr := node.Condition.(*ast.BinaryExpr)

	if isBinaryExpr {
		fmt.Fprintf(
			result,
			"var %s %s %s",
			node.DeclaredVariable,
			binaryExpr.Operator.Atom,
			binaryExpr.Right.Expr(),
		)

		return
	}

	fmt.Fprintf(result, "var %s %s", node.DeclaredVariable, node.Condition.Expr())
}

func (f *Formatter) formatForConditionWithoutVariable(
	node *ast.ForStatement,
	result *strings.Builder,
) {
	binaryExpr, isBinaryExpr := node.Condition.(*ast.BinaryExpr)

	if isBinaryExpr {
		fmt.Fprintf(
			result,
			"%s %s %s",
			binaryExpr.Left.Expr(),
			binaryExpr.Operator.Atom,
			binaryExpr.Right.Expr(),
		)

		return
	}

	result.WriteString(node.Condition.Expr())
}

func (f *Formatter) formatRangeForCondition(
	node *ast.ForStatement,
	result *strings.Builder,
) {
	if node.HasExplicitFrom {
		f.formatForExplicitRange(node, result)

		return
	}

	f.formatForImplicitRange(node, result)
}

func (f *Formatter) formatForExplicitRange(
	node *ast.ForStatement,
	result *strings.Builder,
) {
	if node.DeclaredVariable != "" {
		fmt.Fprintf(
			result,
			" var %s from %s to %s",
			node.DeclaredVariable,
			node.RangeFrom.Expr(),
			node.RangeTo.Expr(),
		)

		return
	}

	fmt.Fprintf(
		result,
		" from %s to %s",
		node.RangeFrom.Expr(),
		node.RangeTo.Expr(),
	)
}

func (f *Formatter) formatForImplicitRange(
	node *ast.ForStatement,
	result *strings.Builder,
) {
	if node.DeclaredVariable != "" {
		fmt.Fprintf(
			result,
			" var %s to %s",
			node.DeclaredVariable,
			node.RangeTo.Expr(),
		)

		return
	}

	fmt.Fprintf(result, " to %s", node.RangeTo.Expr())
}
