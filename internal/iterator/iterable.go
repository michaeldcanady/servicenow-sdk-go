package iterator

type Iterable[T any] interface {
	Next() (T, error)
}
