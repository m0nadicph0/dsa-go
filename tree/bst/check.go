package bst

func IsBST(root *Node) bool {
	return isBSTHelper(root, nil, nil)
}

func isBSTHelper(node *Node, min *Node, max *Node) bool {
	// Base case: an empty tree is a BST.
	if node == nil {
		return true
	}

	// Check if the current node's value is within the valid range.
	if (min != nil && node.Value <= min.Value) || (max != nil && node.Value >= max.Value) {
		return false
	}

	// Recursively check the left and right subtrees with updated min and max.
	return isBSTHelper(node.Left, min, node) && isBSTHelper(node.Right, node, max)
}
