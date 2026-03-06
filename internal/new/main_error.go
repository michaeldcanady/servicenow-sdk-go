package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	detailKey  = "detail"
	messageKey = "message"
	statusKey  = "status"
)

type MainErrorable interface {
	GetDetail() (*string, error)
	GetMessage() (*string, error)
	GetStatus() (*string, error)
	Model
	serialization.Parsable
}

// MainError represents internal error of Service-Now API error
type MainError struct {
	Model
}

// NewMainError instantiates a new MainError
func NewMainError() *MainError {
	return &MainError{
		NewBaseModel(),
	}
}

// CreateBatchRequest2FromDiscriminatorValue is a parsable factory for creating a MainError
func CreateMainErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewMainError(), nil
}

// Serialize writes the objects properties to the current writer.
func (exc *MainError) Serialize(_ serialization.SerializationWriter) error {
	return errors.New("unsupported")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (exc *MainError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		detailKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return exc.setDetail(val)
		},
		messageKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return exc.setMessage(val)
		},
		statusKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return exc.setStatus(val)
		},
	}
}

// GetDetail returns the error details.
func (exc *MainError) GetDetail() (*string, error) {
	if IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelAccessorFunc[store.BackingStore, *string](backingStore, detailKey)
}

// setDetail sets the error details.
func (exc *MainError) setDetail(detail *string) error {
	if IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, detailKey, detail)
}

// GetMessage gets the error message.
func (exc *MainError) GetMessage() (*string, error) {
	if IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelAccessorFunc[store.BackingStore, *string](backingStore, messageKey)
}

// setMessage sets the error message.
func (exc *MainError) setMessage(message *string) error {
	if IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, messageKey, message)
}

// GetStatus returns the status.
func (exc *MainError) GetStatus() (*string, error) {
	if IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelAccessorFunc[store.BackingStore, *string](backingStore, statusKey)
}

// setStatus sets the status.
func (exc *MainError) setStatus(status *string) error {
	if IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, statusKey, status)
}
