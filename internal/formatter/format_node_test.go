package formatter

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

type unknownNode struct{}

func (n *unknownNode) EndPosition() int                  { return 0 }
func (n *unknownNode) StartPosition() int                { return 0 }
func (n *unknownNode) Expr() string                      { return "unknown" }
func (n *unknownNode) Walk(fn func(_ ast.ExprNode) bool) { fn(n) }

func TestFormatNodeUnknownASTNode(t *testing.T) {
	t.Parallel()

	formatter := New()
	result := formatter.Format(&unknownNode{})

	if result != "unknown\n" {
		t.Errorf("expected \"unknown\\n\", got \"%s\"", result)
	}
}
