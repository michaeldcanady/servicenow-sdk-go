package internal

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type Model interface {
	serialization.Parsable
	BackedModel
}
