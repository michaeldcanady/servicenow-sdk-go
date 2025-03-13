package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
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
func (exc *ServicenowError) Serialize(_ serialization.SerializationWriter) error {
	return errors.New("unsupported")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (exc *ServicenowError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		errorKey: func(pn serialization.ParseNode) error {
			parsable, err := pn.GetObjectValue(CreateMainErrorFromDiscriminatorValue)
			if err != nil {
				return err
			}
			mainError, ok := parsable.(MainErrorable)
			if !ok {
				return errors.New("parsable is not MainErrorable")
			}
			return exc.setError(mainError)
		},
	}
}

// GetError returns the main error
func (exc *ServicenowError) GetError() (MainErrorable, error) {
	if IsNil(exc) {
		return nil, nil
	}

	rawMainErr, err := exc.GetBackingStore().Get(errorKey)
	if err != nil {
		return nil, err
	}

	mainErr, ok := rawMainErr.(MainErrorable)
	if !ok {
		return nil, errors.New("rawMainErr is not MainErrorable")
	}

	return mainErr, nil
}

// setError sets the main error
func (exc *ServicenowError) setError(mainError MainErrorable) error {
	if IsNil(exc) {
		return nil
	}
	backingStore := exc.GetBackingStore()
	if IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}
	return backingStore.Set(errorKey, mainError)
}
