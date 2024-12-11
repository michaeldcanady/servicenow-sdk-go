package authentication

// Cache interface defines methods for a cache system.
type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}
