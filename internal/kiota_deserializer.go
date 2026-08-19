package internal

import "github.com/microsoft/kiota-abstractions-go/serialization"

type KiotaDeserializer struct{}

func NewKiotaDeserializer() *KiotaDeserializer {
	return &KiotaDeserializer{}
}

func (kD *KiotaDeserializer) Deserialize(contentType string, content []byte, parsableFactory serialization.ParsableFactory) (serialization.Parsable, error) {
	return serialization.Deserialize(contentType, content, parsableFactory)
}
