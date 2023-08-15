package list

import "errors"

type VisitFn func(int)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.count += 1
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head = newNode
	}
	l.count += 1
}

func (l *LinkedList) Size() int {
	return l.count
}

func (l *LinkedList) ForEach(fn VisitFn) {
	p := l.head
	for p != nil {
		fn(p.Value)
		p = p.Next
	}
}

func (l *LinkedList) Get(n int) int {
	p := l.head
	for i := 0; p != nil; i++ {
		if i == n {
			return p.Value
		}
		p = p.Next
	}
	return -1
}

func (l *LinkedList) InsertAt(index int, value int) error {
	if index < 0 || index > l.Size() {
		return errors.New("index out of bounds")
	}

	newNode := &Node{Value: value, Next: nil}

	if index == 0 {
		newNode.Next = l.head
		l.head = newNode
		if newNode.Next == nil {
			l.tail = newNode
		}
	} else {
		p := l.head
		for i := 0; p != nil; i++ {
			if i == (index - 1) {
				newNode.Next = p.Next
				p.Next = newNode
				if newNode.Next == nil {
					l.tail = newNode
				}
			}
			p = p.Next
		}
	}

	l.count += 1
	return nil
}

func (l *LinkedList) DeleteAt(index int) error {
	if l.count == 0 || index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		l.head = l.head.Next
		l.count -= 1
	} else {
		p := l.head
		for i := 0; p != nil; i++ {
			if i == (index - 1) {
				p.Next = p.Next.Next
				l.count -= 1
			}
			p = p.Next
		}
	}
	return nil
}

func (l *LinkedList) Length() int {
	return recLength(l.head)
}

func recLength(p *Node) int {
	if p == nil {
		return 0
	}
	return recLength(p.Next) + 1
}

func (l *LinkedList) IsPresent(value int) bool {
	p := l.head
	for p != nil {
		if p.Value == value {
			return true
		}
		p = p.Next
	}
	return false
}

func (l *LinkedList) IsPresentRec(value int) bool {
	return isPresentHelper(l.head, value)
}

func isPresentHelper(p *Node, value int) bool {
	if p == nil {
		return false
	}
	return (p.Value == value) || isPresentHelper(p.Next, value)
}
