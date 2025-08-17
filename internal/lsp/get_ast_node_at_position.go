package lsp

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func getAstNodeAtPosition(
	node ast.ExprNode,
	charIndex int,
) ast.ExprNode {
	var bestNode ast.ExprNode
	closestDistance := -1

	node.Walk(func(currentNode ast.ExprNode) bool {
		if charIndex < currentNode.StartPosition() || charIndex >= currentNode.EndPosition() {
			return true
		}

		nodeSize := currentNode.EndPosition() - currentNode.StartPosition()
		distance := charIndex - currentNode.StartPosition()

		if bestNode == nil {
			bestNode = currentNode
			closestDistance = distance
		} else {
			bestNodeSize := bestNode.EndPosition() - bestNode.StartPosition()

			if nodeSize < bestNodeSize ||
				(nodeSize == bestNodeSize && distance < closestDistance) {
				bestNode = currentNode
				closestDistance = distance
			}
		}

		return true
	})

	return bestNode
}
