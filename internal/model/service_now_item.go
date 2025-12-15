package model

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ServiceNowItem interface {
	store.BackedModel
	serialization.Parsable
	GetSysID() (*string, error)
}
