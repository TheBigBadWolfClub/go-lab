package pkg

type DoubleLinkedNode[T any] struct {
	next  *DoubleLinkedNode[T]
	prev  *DoubleLinkedNode[T]
	valid bool
	value T
}

type DoubleLinkedList[T any] struct {
	head  *DoubleLinkedNode[T]
	tail  *DoubleLinkedNode[T]
	count int
}

func (l *DoubleLinkedList[T]) Size() int {
	return l.count
}

func (l *DoubleLinkedList[T]) AddHead(v T) {

	node := DoubleLinkedNode[T]{
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

func (l *DoubleLinkedList[T]) AddTail(v T) {
	node := DoubleLinkedNode[T]{
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

func (l *DoubleLinkedList[T]) Add(v T, i int) error {
	node, err := l.linearSearch(i)
	if err != nil {
		return err
	}

	newNode := &DoubleLinkedNode[T]{
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

func (l *DoubleLinkedList[T]) Delete(index int) error {
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

func (l *DoubleLinkedList[T]) DeleteHead() error {
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

func (l *DoubleLinkedList[T]) DeleteTail() error {
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

func (l *DoubleLinkedList[T]) linearSearch(index int) (*DoubleLinkedNode[T], error) {
	cur := l.head
	for i := 0; i < l.Size(); i++ {
		if i == index {
			return cur, nil
		}
		cur = cur.next
	}

	return nil, IndexOutOfBounds
}

func (l *DoubleLinkedList[T]) tailIndex() int {
	return l.Size() - 1
}
