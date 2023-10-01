package twoptr

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewList() *LinkedList {
	return &LinkedList{head: nil}
}

func (ll *LinkedList) Insert(value int) {
	if ll.head == nil {
		ll.head = &Node{
			value: value,
			next:  nil,
		}
		return
	}
	p := ll.head
	ll.head = &Node{
		value: value,
		next:  p,
	}
}

func (ll *LinkedList) Print() {
	p := ll.head
	for p != nil {
		fmt.Println(p.value)
		p = p.next
	}
}

func (ll *LinkedList) RemoveNth(n int) {
	ll.head = removeNthLastNode(ll.head, n)
}

func (ll *LinkedList) ToSlice() []int {
	result := make([]int, 0)
	p := ll.head
	for p != nil {
		result = append(result, p.value)
		p = p.next
	}
	return result
}

func removeNthLastNode(head *Node, n int) *Node {

	left, right := head, head
	for i := 0; i < n && right != nil; i++ {
		right = right.next
	}

	if right == nil {
		return head.next
	}

	for right.next != nil {
		left = left.next
		right = right.next
	}
	left.next = left.next.next
	return head
}
