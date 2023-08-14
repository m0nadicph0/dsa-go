package list

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
