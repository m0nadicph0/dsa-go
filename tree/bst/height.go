package bst

import "github.com/m0nadicph0/dsa-go/util"

func CalculateHeight(node *Node) int {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return 0
	}
	leftHeight := CalculateHeight(node.Left)
	rightHeight := CalculateHeight(node.Right)
	return 1 + util.Max(leftHeight, rightHeight)
}
