package model

import (
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type Model interface {
	serialization.Parsable
	internal.BackedModel
}
