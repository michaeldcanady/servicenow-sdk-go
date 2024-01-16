package internal

type Test[T any] struct {
	Title    string
	Input    interface{}
	Expected T
	Error    error
}
