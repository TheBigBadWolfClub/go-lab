package pkg

type LinkedList[T any] struct {
	Len  int
	Head *linkedListNode[T]
}

func (l *LinkedList[T]) Add(v T) {
	newNode := NewSingleListNode(v)
	if l.Head == nil {
		l.Head = newNode
		l.Len++
		return
	}

	tail := l.Tail()
	tail.next = newNode
	l.Len++
}

func (l *LinkedList[T]) AddNode(v *linkedListNode[T]) {

	if l.Head == nil {
		l.Head = v
		l.Len++
		return
	}

	tail := l.Tail().next
	tail.next = v
}

func (l *LinkedList[T]) Get(index int) (T, bool) {
	var zero T
	if l.Head == nil || l.Len <= index {
		return zero, false
	}

	cur := l.Head
	for i := 1; i <= index; i++ {
		cur = cur.next
	}

	return cur.value, true
}

func (l *LinkedList[T]) Delete(index int) error {
	if l.Head == nil || l.Len <= index {
		return IndexNotFound
	}

	if index == 0 {
		l.Head = l.Head.next
		l.Len--
		return nil
	}

	cur := l.Head
	for i := 1; i < index; i++ {
		cur = cur.next
	}

	if cur.next == nil {
		return nil
	}

	cur.next = cur.next.next
	l.Len--
	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Len == 0
}

func (l *LinkedList[T]) Size() int {
	return l.Len
}

func (l *LinkedList[T]) Tail() *linkedListNode[T] {
	if l.Head == nil || l.Len == 0 {
		return nil
	}

	cur := l.Head
	for i := 1; i < l.Len; i++ {
		cur = cur.next
	}
	return cur
}

func (l *LinkedList[T]) DeleteHead() (ok bool) {
	if l.Head == nil {
		return true
	}

	l.Head = l.Head.next
	l.Len--
	return true
}

func (l *LinkedList[T]) DeleteIndex(index int) (ok bool) {
	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	if cur.next == nil {
		return true
	}
	cur.next = cur.next.next
	return true
}

func (l *LinkedList[T]) Reverse() {
	var cur, next, prev *linkedListNode[T]
	cur = l.Head
	for cur != nil {
		// swap
		next = cur.next
		cur.next = prev

		// prepare next
		prev = cur
		cur = next // b
	}

	l.Head = prev
}

type linkedListNode[T any] struct {
	next  *linkedListNode[T]
	value T
}

func NewSingleListNode[T any](v T) *linkedListNode[T] {
	return &linkedListNode[T]{
		next:  nil,
		value: v,
	}
}
