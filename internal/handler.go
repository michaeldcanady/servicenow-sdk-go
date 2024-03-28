package internal

type Handler[T any] interface {
	Handle(T) error
	SetNext(Handler[T])
	Next() Handler[T]
}
