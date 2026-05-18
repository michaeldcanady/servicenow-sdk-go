package accountapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Account represents an account object in ServiceNow.
type Account struct {
	newInternal.BaseModel
}

// NewAccount creates a new instance of the Account model.
func NewAccount() *Account {
	return &Account{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

// CreateAccountFromDiscriminatorValue is a factory for creating an Account model.
func CreateAccountFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAccount(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *Account) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *Account) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}
