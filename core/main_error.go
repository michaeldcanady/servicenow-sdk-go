package core

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	internalStore "github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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
	BackedModel
	serialization.Parsable
}

// MainError represents internal error of Service-Now API error
type MainError struct {
	BackedModel
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
	if conversion.IsNil(exc) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(detailKey, exc.GetDetail),
		internalSerialization.SerializeStringFunc(messageKey, exc.GetMessage),
		internalSerialization.SerializeStringFunc(statusKey, exc.GetStatus),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (exc *MainError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		detailKey:  internalSerialization.DeserializeStringFunc(exc.SetDetail),
		messageKey: internalSerialization.DeserializeStringFunc(exc.SetMessage),
		statusKey:  internalSerialization.DeserializeStringFunc(exc.SetStatus),
	}
}

// GetDetail returns the error details.
func (exc *MainError) GetDetail() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*MainError, *string](exc, detailKey)
}

// SetDetail sets the error details.
func (exc *MainError) SetDetail(detail *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(exc, detailKey, detail)
}

// GetMessage gets the error message.
func (exc *MainError) GetMessage() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*MainError, *string](exc, messageKey)
}

// SetMessage sets the error message.
func (exc *MainError) SetMessage(message *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(exc, messageKey, message)
}

// GetStatus returns the status.
func (exc *MainError) GetStatus() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*MainError, *string](exc, statusKey)
}

// SetStatus sets the status.
func (exc *MainError) SetStatus(status *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(exc, statusKey, status)
}
