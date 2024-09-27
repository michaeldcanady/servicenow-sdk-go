package auth

type Cache[T any] interface {
	Store(string, T) error
	Retrieve(string) (T, error)
}
