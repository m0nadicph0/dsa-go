package house_robber_III

import "github.com/m0nadicph0/dsa-go/util"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// RobIII returns the maximum amount of money that can robed
// without alerting the police
func RobIII(root *Node) int {
	if root == nil {
		return 0
	}

	// Rob current
	robCurrent := root.Value
	if root.Left != nil {
		robCurrent += RobIII(root.Left.Left) + RobIII(root.Left.Right)
	}

	if root.Right != nil {
		robCurrent += RobIII(root.Right.Left) + RobIII(root.Right.Right)
	}

	// Skip current
	skipCurrent := RobIII(root.Left) + RobIII(root.Right)

	return util.Max(robCurrent, skipCurrent)
}
