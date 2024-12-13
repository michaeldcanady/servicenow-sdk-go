package iterator

import "github.com/microsoft/kiota-abstractions-go/serialization"

type HasResult interface {
	GetResult() ([]serialization.Parsable, error)
}
