package internal

import "github.com/microsoft/kiota-abstractions-go/serialization"

type Deserializer interface {
	Deserialize(contentType string, content []byte, parsableFactory serialization.ParsableFactory) (serialization.Parsable, error)
}
