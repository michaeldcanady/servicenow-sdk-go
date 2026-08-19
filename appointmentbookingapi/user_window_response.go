package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserWindowResponse represents the response returned by the userwindow endpoint.
type UserWindowResponse = core.ServiceNowItemResponse[*UserWindowResult]

// CreateUserWindowResponseFromDiscriminatorValue is a factory for creating a UserWindowResponse.
func CreateUserWindowResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UserWindowResult](CreateUserWindowResultFromDiscriminatorValue), nil
}

// UserWindowResult represents the result details returned by the userwindow endpoint. The spec
// defines no schema for this response.
type UserWindowResult struct {
	core.BaseModel
}

// NewUserWindowResult creates a new instance of UserWindowResult.
func NewUserWindowResult() *UserWindowResult {
	return &UserWindowResult{BaseModel: *core.NewBaseModel()}
}

// CreateUserWindowResultFromDiscriminatorValue creates a new UserWindowResult from a ParseNode.
func CreateUserWindowResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserWindowResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *UserWindowResult) Serialize(_ serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *UserWindowResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}
