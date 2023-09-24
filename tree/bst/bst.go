package bst

type VisitFn func(*Node)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewBSTNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

type BST struct {
	root *Node
}

func (b *BST) Insert(value int) {
	newNode := NewBSTNode(value)
	if b.root == nil {
		b.root = newNode
	} else {
		insertRecursive(b.root, newNode)
	}
}

func insertRecursive(root *Node, newNode *Node) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			insertRecursive(root.Left, newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
		} else {
			insertRecursive(root.Right, newNode)
		}
	}
}

func (b *BST) InOrder(fn VisitFn) {
	inorderRecursive(b.root, fn)
}

func (b *BST) PreOrder(fn VisitFn) {
	preOrderRecursive(b.root, fn)
}

func (b *BST) PostOrder(fn VisitFn) {
	postOrderRecursive(b.root, fn)
}

func postOrderRecursive(node *Node, fn VisitFn) {
	if node != nil {
		postOrderRecursive(node.Left, fn)
		postOrderRecursive(node.Right, fn)
		fn(node)
	}
}

func preOrderRecursive(node *Node, fn VisitFn) {
	if node != nil {
		fn(node)
		preOrderRecursive(node.Left, fn)
		preOrderRecursive(node.Right, fn)
	}
}

func inorderRecursive(node *Node, fn VisitFn) {
	if node != nil {
		inorderRecursive(node.Left, fn)
		fn(node)
		inorderRecursive(node.Right, fn)
	}
}

func NewBST() *BST {
	return &BST{root: nil}
}
