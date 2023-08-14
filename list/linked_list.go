package list

type VisitFn func(int)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head  *Node
	count int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (l *LinkedList) Append(value int) {
	if l.head == nil {
		l.head = &Node{
			Value: value,
			Next:  nil,
		}
	} else {
		p := l.head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = &Node{
			Value: value,
			Next:  nil,
		}
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
