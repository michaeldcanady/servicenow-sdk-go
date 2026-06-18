package core

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type Model interface {
	serialization.Parsable
	BackedModel
}
