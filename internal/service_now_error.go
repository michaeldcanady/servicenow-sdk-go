package internal

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	errorKey = "error"
)

// ServicenowError represents a Service-Now API error
type ServicenowError struct {
	BackedModel
}

// BadRequestError represents a 400 Bad Request error
type BadRequestError struct {
	ServicenowError
}

// UnauthorizedError represents a 401 Unauthorized error
type UnauthorizedError struct {
	ServicenowError
}

// ForbiddenError represents a 403 Forbidden error
type ForbiddenError struct {
	ServicenowError
}

// NotFoundError represents a 404 Not Found error
type NotFoundError struct {
	ServicenowError
}

// TooManyRequestsError represents a 429 Too Many Requests error
type TooManyRequestsError struct {
	ServicenowError
}

// ServerError represents a 5XX Server error
type ServerError struct {
	ServicenowError
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

// CreateBadRequestErrorFromDiscriminatorValue creates a BadRequestError
func CreateBadRequestErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &BadRequestError{*NewServicenowError()}, nil
}

// CreateUnauthorizedErrorFromDiscriminatorValue creates a UnauthorizedError
func CreateUnauthorizedErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &UnauthorizedError{*NewServicenowError()}, nil
}

// CreateForbiddenErrorFromDiscriminatorValue creates a ForbiddenError
func CreateForbiddenErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &ForbiddenError{*NewServicenowError()}, nil
}

// CreateNotFoundErrorFromDiscriminatorValue creates a NotFoundError
func CreateNotFoundErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &NotFoundError{*NewServicenowError()}, nil
}

// CreateTooManyRequestsErrorFromDiscriminatorValue creates a TooManyRequestsError
func CreateTooManyRequestsErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &TooManyRequestsError{*NewServicenowError()}, nil
}

// CreateServerErrorFromDiscriminatorValue creates a ServerError
func CreateServerErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &ServerError{*NewServicenowError()}, nil
}

// Serialize writes the objects properties to the current writer.
func (exc *ServicenowError) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(exc) {
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
	if conversion.IsNil(exc) {
		return nil, nil
	}

	backingStore := exc.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, MainErrorable](backingStore, errorKey)
}

// setError sets the main error
func (exc *ServicenowError) setError(mainError MainErrorable) error {
	if conversion.IsNil(exc) {
		return nil
	}

	backingStore := exc.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, errorKey, mainError)
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
