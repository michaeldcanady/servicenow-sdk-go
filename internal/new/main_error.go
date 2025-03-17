package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	detailKey  = "Detail"
	messageKey = "Message"
	statusKey  = "Status"
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
	if IsNil(backingStore) {
		return nil, errors.New("backingStore is nil")
	}

	detail, err := backingStore.Get(detailKey)
	if err != nil {
		return nil, err
	}

	detailStr, ok := detail.(*string)
	if !ok {
		return nil, errors.New("detail is not *string")
	}

	return detailStr, nil
}

// setDetail sets the error details.
func (exc *MainError) setDetail(detail *string) error {
	if IsNil(exc) {
		return nil
	}
	backingStore := exc.GetBackingStore()
	if IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}
	return backingStore.Set(detailKey, detail)
}

// GetMessage gets the error message.
func (exc *MainError) GetMessage() (*string, error) {
	if IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	if IsNil(backingStore) {
		return nil, errors.New("backingStore is nil")
	}

	message, err := backingStore.Get(messageKey)
	if err != nil {
		return nil, err
	}

	detailStr, ok := message.(*string)
	if !ok {
		return nil, errors.New("message is not *string")
	}

	return detailStr, nil
}

// setMessage sets the error message.
func (exc *MainError) setMessage(message *string) error {
	if IsNil(exc) {
		return nil
	}
	backingStore := exc.GetBackingStore()
	if IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}
	return backingStore.Set(messageKey, message)
}

// GetStatus returns the status.
func (exc *MainError) GetStatus() (*string, error) {
	if IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	if IsNil(backingStore) {
		return nil, errors.New("backingStore is nil")
	}

	status, err := backingStore.Get(statusKey)
	if err != nil {
		return nil, err
	}

	detailStr, ok := status.(*string)
	if !ok {
		return nil, errors.New("status is not *string")
	}

	return detailStr, nil
}

// setStatus sets the status.
func (exc *MainError) setStatus(status *string) error {
	if IsNil(exc) {
		return nil
	}
	backingStore := exc.GetBackingStore()
	if IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}
	return backingStore.Set(statusKey, status)
}
