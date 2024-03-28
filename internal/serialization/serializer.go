package serialization

// Serializer interface that all serializers must implement
type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
}
