package actsubapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	messageKey    = "message"
	streamKey     = "stream"
	userKey       = "user"
	activitiesKey = "activities"
	statusKey     = "status"
)

// ActivitySubscription represents an activity subscriptions.
type ActivitySubscription struct {
	core.BaseModel
}

// NewActivitySubscription creates a new instance of ActivitySubscriptionModel.
func NewActivitySubscription() *ActivitySubscription {
	return &ActivitySubscription{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateActivitySubscriptionModelFromDiscriminatorValue creates a new instance of ActivitySubscriptionModel.
func CreateActivitySubscriptionModelFromDiscriminatorValue(serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitySubscription(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ActivitySubscription) Serialize(_ serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ActivitySubscription) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		messageKey:    internalSerialization.DeserializeStringFunc(m.SetMessage),
		streamKey:     internalSerialization.DeserializeStringFunc(m.SetStream),
		userKey:       internalSerialization.DeserializeStringFunc(m.SetUser),
		activitiesKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*Activity](CreateActivityFromDiscriminatorValue, m.SetActivities),
		statusKey:     internalSerialization.DeserializeInt64Func(m.SetStatus),
	}
}

// GetMessage returns the message value.
func (m *ActivitySubscription) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscription, *string](m, messageKey)
}

// SetMessage sets the message value.
func (m *ActivitySubscription) SetMessage(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, value)
}

// GetStream returns the stream value.
func (m *ActivitySubscription) GetStream() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscription, *string](m, streamKey)
}

// SetStream sets the stream value.
func (m *ActivitySubscription) SetStream(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, streamKey, value)
}

// GetUser returns the user value.
func (m *ActivitySubscription) GetUser() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscription, *string](m, userKey)
}

// SetUser sets the user value.
func (m *ActivitySubscription) SetUser(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userKey, value)
}

// GetActivities returns the activities value.
func (m *ActivitySubscription) GetActivities() ([]*Activity, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscription, []*Activity](m, activitiesKey)
}

// SetActivities sets the activities value.
func (m *ActivitySubscription) SetActivities(value []*Activity) error {
	return store.DefaultBackedModelMutatorFunc(m, activitiesKey, value)
}

// GetStatus returns the status value.
func (m *ActivitySubscription) GetStatus() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscription, *int64](m, statusKey)
}

// SetStatus sets the status value.
func (m *ActivitySubscription) SetStatus(value *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, value)
}
