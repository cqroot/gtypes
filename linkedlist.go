package gtypes

type node[T comparable] struct {
	data T
	prev *node[T]
	next *node[T]
}

type LinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Add(value T) {
	newNode := node[T]{
		data: value,
		prev: nil,
		next: nil,
	}

	if l.size == 0 {
		newNode.prev = &newNode
		newNode.next = &newNode
		l.head = &newNode
		l.tail = &newNode
		l.size++
		return
	}

	newNode.prev = l.tail
	newNode.next = l.head
	l.tail.next = &newNode
	l.head.prev = &newNode
	l.tail = &newNode
	l.size++
}

func (l *LinkedList[T]) Remove(index int) {
	if index < 0 || index >= l.size {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}

	n := l.head
	for range index {
		n = n.next
	}
	n.prev.next = n.next
	n.next.prev = n.prev

	if index == 0 {
		l.head = n.next
	}
	if index == l.size {
		l.tail = n.prev
	}

	l.size--
}

func (l LinkedList[T]) IndexOf(value T) int {
	n := l.head
	for i := range l.size {
		if n.data == value {
			return i
		}
		n = n.next
	}
	return -1
}

func (l LinkedList[T]) ForEach(f func(i int, data T) error) error {
	if l.size == 0 {
		return nil
	}

	n := l.head
	for i := 0; i < l.size; i++ {
		err := f(i, n.data)
		if err != nil {
			return err
		}
		n = n.next
	}
	return nil
}
