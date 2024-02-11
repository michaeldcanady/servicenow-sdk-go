package internal

type PageResult[T any] interface {
	Result() []*T
	NextLink() string
	PreviousLink() string
	FirstLink() string
	LastLink() string
}
