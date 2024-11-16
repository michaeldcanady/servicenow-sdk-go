package internal

import "github.com/microsoft/kiota-abstractions-go/serialization"

// SerializationWriterAdapter adapter struct to handle conversions between "actual" type and Service-Now variation
type SerializationWriterAdapter struct {
	writer serialization.SerializationWriter //nolint:unused
}
