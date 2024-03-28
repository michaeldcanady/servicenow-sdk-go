package serialization

// Unmarshaller interface that all unmarshallers must implement
type Unmarshaller interface {
	Unmarshal(data []byte, v interface{}) error
}
