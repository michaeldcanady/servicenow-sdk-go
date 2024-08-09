package internal

type LinkedList[T any] interface {
	GetHead() Node[T]
	AddNode(Node[T])
	GetTail() Node[T]
}

type linkedListImpl[T any] struct {
	head Node[T]
	tail Node[T]
}

func NewLinkedList[T any]() LinkedList[T] {
	return &linkedListImpl[T]{
		head: nil,
		tail: nil,
	}
}

func (l *linkedListImpl[T]) GetHead() Node[T] {
	return l.head
}
func (l *linkedListImpl[T]) AddNode(node Node[T]) {
	if l == nil {
		l = &linkedListImpl[T]{}
	}

	if l.head == nil {
		l.head = node
	}

	if l.tail != nil {
		l.tail.SetNext(node)
	}
	l.tail = node
}
func (l *linkedListImpl[T]) GetTail() Node[T] {
	return l.tail
}
