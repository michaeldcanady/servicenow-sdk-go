package core

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Model is a serializable, backed model: the union of serialization.Parsable and BackedModel.
type Model interface {
	serialization.Parsable
	BackedModel
}
