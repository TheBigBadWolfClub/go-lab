package pkg

type SingleLinkList[T any] struct {
	next  *SingleLinkList[T]
	valid bool
	value T
}

func (l *SingleLinkList[T]) Add(v T) {
	if !l.valid {
		l.next = nil
		l.valid = true
		l.value = v
		return
	}

	tail := l.tail()
	tail.next = &SingleLinkList[T]{
		next:  nil,
		valid: true,
		value: v,
	}
}

func (l *SingleLinkList[T]) Get(index int) (T, bool) {
	var zero T
	if !l.valid {
		return zero, false
	}

	cur := l
	for i := 0; ; i++ {
		if i == index {
			return cur.value, cur.valid
		}

		cur = cur.next
		if cur == nil {
			return zero, false
		}
	}
}

func (l *SingleLinkList[T]) Delete(index int) error {
	if !l.valid {
		return IndexNotFound
	}

	if l.deleteHead(index) {
		return nil
	}

	cur := l
	for i := 0; ; i++ {
		if cur == nil || !cur.valid {
			return IndexNotFound
		}

		// is next to be deleted
		if i+1 == index && cur.next.valid {
			cur.next = cur.next.next
			return nil
		}

		cur = cur.next
	}
}

func (l *SingleLinkList[T]) IsEmpty() bool {
	return !l.valid
}

func (l *SingleLinkList[T]) Size() int {

	var count int
	cur := l
	for {
		if cur == nil || !cur.valid {
			break
		}
		count++
		cur = cur.next
	}

	return count
}

func (l *SingleLinkList[T]) tail() *SingleLinkList[T] {
	if !l.valid {
		return nil
	}

	cur := l
	for {
		if cur.next == nil {
			return cur
		}
		cur = cur.next
	}
}

func (l *SingleLinkList[T]) deleteHead(index int) (ok bool) {
	if index != 0 {
		return false
	}

	if l.next != nil {
		l.valid = l.next.valid
		l.value = l.next.value
		l.next = l.next.next // need to be last
		return true

	}

	var zero T
	l.value = zero
	l.valid = false
	return true
}
