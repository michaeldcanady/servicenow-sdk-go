package internal

type BaseHandler[T any] struct {
	next Handler[T]
}

func NewBaseHandler[T any]() *BaseHandler[T] {
	return &BaseHandler[T]{}
}

// SetNext method for BaseHandler
func (b *BaseHandler[T]) SetNext(handler Handler[T]) {
	b.next = handler
}

func (b *BaseHandler[T]) Next() Handler[T] {
	return b.next
}

func (b *BaseHandler[T]) Handle(input T) error {
	if !IsNil(b.Next()) {
		return b.Next().Handle(input)
	}
	return nil
}
