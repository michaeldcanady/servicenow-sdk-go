package internal

// Dictionary interface definition
type Dictionary[K, V any] interface {
	Get(K) (V, error)
	Add(K, V) error
	Update(K, V) error
	Contains(K) bool
	Remove(K) error
	Pop(K) (V, error)
}
