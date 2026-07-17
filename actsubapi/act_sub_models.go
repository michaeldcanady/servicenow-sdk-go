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

// ActivitySubscriptionModel represents a generic model for Activity Subscriptions.
type ActivitySubscriptionModel struct {
	core.BaseModel
}

// NewActivitySubscriptionModel creates a new instance of ActivitySubscriptionModel.
func NewActivitySubscriptionModel() *ActivitySubscriptionModel {
	return &ActivitySubscriptionModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateActivitySubscriptionModelFromDiscriminatorValue creates a new instance of ActivitySubscriptionModel.
func CreateActivitySubscriptionModelFromDiscriminatorValue(serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitySubscriptionModel(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ActivitySubscriptionModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ActivitySubscriptionModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		messageKey:    internalSerialization.DeserializeStringFunc(m.SetMessage),
		streamKey:     internalSerialization.DeserializeStringFunc(m.SetStream),
		userKey:       internalSerialization.DeserializeStringFunc(m.SetUser),
		activitiesKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*Activity](CreateActivityFromDiscriminatorValue, m.SetActivities),
		statusKey:     internalSerialization.DeserializeInt64Func(m.SetStatus),
	}
}

func (m *ActivitySubscriptionModel) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscriptionModel, *string](m, messageKey)
}

func (m *ActivitySubscriptionModel) SetMessage(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, value)
}

func (m *ActivitySubscriptionModel) GetStream() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscriptionModel, *string](m, streamKey)
}

func (m *ActivitySubscriptionModel) SetStream(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, streamKey, value)
}

func (m *ActivitySubscriptionModel) GetUser() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscriptionModel, *string](m, userKey)
}

func (m *ActivitySubscriptionModel) SetUser(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userKey, value)
}

func (m *ActivitySubscriptionModel) GetActivities() ([]*Activity, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscriptionModel, []*Activity](m, activitiesKey)
}

func (m *ActivitySubscriptionModel) SetActivities(value []*Activity) error {
	return store.DefaultBackedModelMutatorFunc(m, activitiesKey, value)
}

func (m *ActivitySubscriptionModel) GetStatus() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitySubscriptionModel, *int64](m, statusKey)
}

func (m *ActivitySubscriptionModel) SetStatus(value *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, value)
}
