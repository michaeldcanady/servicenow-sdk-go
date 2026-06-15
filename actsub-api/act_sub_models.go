package actsubapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
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
	internal.BaseModel
}

// NewActivitySubscriptionModel creates a new instance of ActivitySubscriptionModel.
func NewActivitySubscriptionModel() *ActivitySubscriptionModel {
	return &ActivitySubscriptionModel{
		BaseModel: *internal.NewBaseModel(),
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
		messageKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetMessage(val)
			}
			return nil
		},
		streamKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetStream(val)
			}
			return nil
		},
		userKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetUser(val)
			}
			return nil
		},
		activitiesKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetCollectionOfObjectValues(CreateActivityFromDiscriminatorValue)
			if err != nil {
				return err
			}
			if val != nil {
				res := make([]*Activity, len(val))
				for i, v := range val {
					if v != nil {
						res[i] = v.(*Activity)
					}
				}
				return m.SetActivities(res)
			}
			return nil
		},
		statusKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetInt64Value()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetStatus(val)
			}
			return nil
		},
	}
}

func (m *ActivitySubscriptionModel) GetMessage() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, messageKey)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (m *ActivitySubscriptionModel) SetMessage(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, messageKey, value)
}

func (m *ActivitySubscriptionModel) GetStream() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, streamKey)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (m *ActivitySubscriptionModel) SetStream(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, streamKey, value)
}

func (m *ActivitySubscriptionModel) GetUser() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, userKey)
}

func (m *ActivitySubscriptionModel) SetUser(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, userKey, value)
}

func (m *ActivitySubscriptionModel) GetActivities() ([]*Activity, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*Activity](backingStore, activitiesKey)
}

func (m *ActivitySubscriptionModel) SetActivities(value []*Activity) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, activitiesKey, value)
}

func (m *ActivitySubscriptionModel) GetStatus() (*int64, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, statusKey)
}

func (m *ActivitySubscriptionModel) SetStatus(value *int64) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, statusKey, value)
}
