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
		if charIndex < currentNode.StartPosition() || charIndex >= currentNode.EndPosition() {
			return true
		}

		nodeSize := currentNode.EndPosition() - currentNode.StartPosition()

		if bestNode == nil {
			bestNode = currentNode

			return true
		}

		bestNodeSize := bestNode.EndPosition() - bestNode.StartPosition()

		if nodeSize <= bestNodeSize &&
			(currentNode.StartPosition() >= bestNode.StartPosition() &&
				currentNode.EndPosition() <= bestNode.EndPosition()) {
			bestNode = currentNode
		}

		return true
	})

	return bestNode
}
