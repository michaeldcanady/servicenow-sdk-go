package actsubapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivitySubscriptionModel represents a generic model for Activity Subscriptions.
type ActivitySubscriptionModel struct {
	newInternal.BaseModel
}

// NewActivitySubscriptionModel creates a new instance of ActivitySubscriptionModel.
func NewActivitySubscriptionModel() *ActivitySubscriptionModel {
	return &ActivitySubscriptionModel{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

// CreateActivitySubscriptionModelFromDiscriminatorValue creates a new instance of ActivitySubscriptionModel.
func CreateActivitySubscriptionModelFromDiscriminatorValue(serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitySubscriptionModel(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ActivitySubscriptionModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ActivitySubscriptionModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}
