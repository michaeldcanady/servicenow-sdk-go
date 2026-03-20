package model

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
func (exc *MainError) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(exc) {
		return nil
	}

	return kiota.Serialize(writer,
		kiota.SerializeStringFunc(detailKey)(exc.GetDetail),
		kiota.SerializeStringFunc(messageKey)(exc.GetMessage),
		kiota.SerializeStringFunc(statusKey)(exc.GetStatus),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (exc *MainError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		detailKey:  kiota.DeserializeStringFunc(exc.SetDetail),
		messageKey: kiota.DeserializeStringFunc(exc.SetMessage),
		statusKey:  kiota.DeserializeStringFunc(exc.SetStatus),
	}
}

// GetDetail returns the error details.
func (exc *MainError) GetDetail() (*string, error) {
	if utils.IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelAccessorFunc[store.BackingStore, *string](backingStore, detailKey)
}

// SetDetail sets the error details.
func (exc *MainError) SetDetail(detail *string) error {
	if utils.IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, detailKey, detail)
}

// GetMessage gets the error message.
func (exc *MainError) GetMessage() (*string, error) {
	if utils.IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelAccessorFunc[store.BackingStore, *string](backingStore, messageKey)
}

// SetMessage sets the error message.
func (exc *MainError) SetMessage(message *string) error {
	if utils.IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, messageKey, message)
}

// GetStatus returns the status.
func (exc *MainError) GetStatus() (*string, error) {
	if utils.IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelAccessorFunc[store.BackingStore, *string](backingStore, statusKey)
}

// SetStatus sets the status.
func (exc *MainError) SetStatus(status *string) error {
	if utils.IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, statusKey, status)
}
