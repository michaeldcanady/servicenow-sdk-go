package model

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	errorKey = "error"
)

// ServicenowError represents a Service-Now API error
type ServicenowError struct {
	Model
}

// NewServicenowError instantiates a new Service-Now error
func NewServicenowError() *ServicenowError {
	return &ServicenowError{
		NewBaseModel(),
	}
}

// CreateServiceNowErrorFromDiscriminatorValue is a parsable factory for creating a ServicenowError
func CreateServiceNowErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServicenowError(), nil
}

// Serialize writes the objects properties to the current writer.
func (exc *ServicenowError) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(exc) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[MainErrorable](errorKey)(exc.GetError),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (exc *ServicenowError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		errorKey: internalSerialization.DeserializeObjectValueFunc[MainErrorable](CreateMainErrorFromDiscriminatorValue)(exc.setError),
	}
}

// GetError returns the main error
func (exc *ServicenowError) GetError() (MainErrorable, error) {
	if utils.IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, MainErrorable](backingStore, errorKey)
}

// setError sets the main error
func (exc *ServicenowError) setError(mainError MainErrorable) error {
	if utils.IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return kiota.DefaultBackedModelMutatorFunc(backingStore, errorKey, mainError)
}

func (exc *ServicenowError) Error() string {
	mainErr, _ := exc.GetError()
	msg, _ := mainErr.GetMessage()
	if msg != nil {
		return *msg
	}
	details, _ := mainErr.GetDetail()
	return *details
}
