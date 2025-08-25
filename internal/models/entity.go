//go:build preview.tableApiV2

package models

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	sysIDKey = "sys_id"
)

type Entity struct {
	newInternal.Model
}

// NewEntity creates a new instance of Entity
func NewEntity() *Entity {
	return &Entity{newInternal.NewBaseModel()}
}

// CreateEntityFromDiscriminatorValue is a parsable factory for creating an EntityModel
func CreateEntityFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewEntity(), nil
}

// Serialize writes the objects properties to the current writer.
func (rE *Entity) Serialize(_ serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("serialization not supported")
}

// GetFieldDeserializers implements serialization.Parsable.
func (rE *Entity) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	panic("unimplemented")
}

func (e *Entity) GetSysID() (*string, error) {
	store := e.GetBackingStore()
	raw, err := store.Get(sysIDKey)
	if err != nil {
		return nil, err
	}

}
