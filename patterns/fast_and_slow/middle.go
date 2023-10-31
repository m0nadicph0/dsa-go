package fast_and_slow

func GetMiddleNode(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
