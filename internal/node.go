package internal

type Node[T any] interface {
	GetNext() Node[T]
	SetNext(Node[T])
	GetValue() T
	SetValue(val T)
}

type nodeImpl[T any] struct {
	next  Node[T]
	value T
}

func NewNode[T any](elem T) Node[T] {
	return &nodeImpl[T]{
		value: elem,
		next:  nil,
	}
}

func (n *nodeImpl[T]) GetNext() Node[T] {
	return n.next
}

func (n *nodeImpl[T]) GetValue() T {
	return n.value
}

func (n *nodeImpl[T]) SetValue(val T) {
	n.value = val
}

func (n *nodeImpl[T]) SetNext(next Node[T]) {
	n.next = next
}
