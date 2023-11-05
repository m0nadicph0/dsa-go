package fast_and_slow

func IsPalindrome(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	reversed := reverse(slow)
	for head != nil && reversed != nil {
		if head.value != reversed.value {
			return false
		}
		head = head.next
		reversed = reversed.next
	}
	return true
}

func reverse(head *Node) *Node {
	var prev *Node
	var next *Node
	current := head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}
