package pkg

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	valid bool
	value T
}

type LinkedList[T any] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) Size() int {
	return l.count
}

func (l *LinkedList[T]) AddHead(v T) {

	node := Node[T]{
		next:  l.head,
		prev:  nil,
		valid: true,
		value: v,
	}
	l.head = &node

	if node.next != nil {
		node.next.prev = &node
	}

	if l.tail == nil {
		l.tail = &node
	}

	l.count++
}

func (l *LinkedList[T]) AddTail(v T) {
	node := Node[T]{
		next:  nil,
		prev:  l.tail,
		valid: true,
		value: v,
	}
	l.tail = &node

	if node.prev != nil {
		node.prev.next = &node
	}

	if l.head == nil {
		l.head = &node
	}

	l.count++
}

func (l *LinkedList[T]) Add(v T, i int) error {
	node, err := l.linearSearch(i)
	if err != nil {
		return err
	}

	newNode := &Node[T]{
		next:  node,
		prev:  node.prev,
		valid: true,
		value: v,
	}
	newNode.prev.next = newNode
	node.prev = newNode

	l.count++
	return nil
}

func (l *LinkedList[T]) Delete(index int) error {
	if index == 0 {
		return l.DeleteHead()
	}
	if index == l.tailIndex() {
		return l.DeleteTail()
	}

	node, err := l.linearSearch(index)
	if err != nil {
		return err
	}

	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev

	l.count--
	return nil
}

func (l *LinkedList[T]) DeleteHead() error {
	if l.Size() == 0 {
		return IndexOutOfBounds
	}

	delNode := l.head
	l.head = delNode.next

	if delNode.next != nil {
		delNode.next.prev = nil
	}

	if delNode == l.tail {
		l.tail = l.head
	}

	l.count--
	return nil
}

func (l *LinkedList[T]) DeleteTail() error {
	if l.Size() == 0 {
		return IndexOutOfBounds
	}

	delNode := l.tail
	l.tail = delNode.prev

	if delNode.prev != nil {
		delNode.prev.next = nil
	}

	if delNode == l.head {
		l.head = l.tail
	}

	l.count--
	return nil
}

func (l *LinkedList[T]) linearSearch(index int) (*Node[T], error) {
	cur := l.head
	for i := 0; i < l.Size(); i++ {
		if i == index {
			return cur, nil
		}
		cur = cur.next
	}

	return nil, IndexOutOfBounds
}

func (l *LinkedList[T]) tailIndex() int {
	return l.Size() - 1
}
