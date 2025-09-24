package lsp

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func getAstNodeAtPosition(
	node ast.ExprNode,
	charIndex int,
) ast.ExprNode {
	var bestNode ast.ExprNode

	node.Walk(func(currentNode ast.ExprNode) bool {
		if charIndex < currentNode.GetRange().Start.Offset || charIndex > currentNode.GetRange().End.Offset {
			return true
		}

		nodeSize := currentNode.GetRange().End.Offset - currentNode.GetRange().Start.Offset

		if bestNode == nil {
			bestNode = currentNode

			return true
		}

		bestNodeSize := bestNode.GetRange().End.Offset - bestNode.GetRange().Start.Offset

		if nodeSize <= bestNodeSize &&
			(currentNode.GetRange().Start.Offset >= bestNode.GetRange().Start.Offset &&
				currentNode.GetRange().End.Offset <= bestNode.GetRange().End.Offset) {
			bestNode = currentNode
		}

		return true
	})

	return bestNode
}
