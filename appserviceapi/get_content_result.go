package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetContentResult represents the result details retrieved from the getContent endpoint.
// The spec defines no schema for this response (it is a CI/relationship content graph).
type GetContentResult struct {
	core.BaseModel
}

// NewGetContentResult creates a new instance of GetContentResult.
func NewGetContentResult() *GetContentResult {
	return &GetContentResult{BaseModel: *core.NewBaseModel()}
}

// CreateGetContentResultFromDiscriminatorValue creates a new GetContentResult from a ParseNode.
func CreateGetContentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewGetContentResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *GetContentResult) Serialize(_ serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *GetContentResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}
