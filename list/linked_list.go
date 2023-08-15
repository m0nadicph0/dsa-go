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
