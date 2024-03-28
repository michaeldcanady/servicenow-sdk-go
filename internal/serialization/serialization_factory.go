package serialization

type SerializationFactory struct {
	serailizers map[string]Serializer
}

func (s *SerializationFactory) CreateSerializer(contentType string) Serializer {

}
