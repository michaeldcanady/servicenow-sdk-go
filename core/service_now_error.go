package core

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	errorKey = "error"
)

// ServiceNowError represents a Service-Now API error
type ServiceNowError struct {
	BackedModel
}

// BadRequestError represents a 400 Bad Request error
type BadRequestError struct {
	ServiceNowError
}

// UnauthorizedError represents a 401 Unauthorized error
type UnauthorizedError struct {
	ServiceNowError
}

// ForbiddenError represents a 403 Forbidden error
type ForbiddenError struct {
	ServiceNowError
}

// NotFoundError represents a 404 Not Found error
type NotFoundError struct {
	ServiceNowError
}

// TooManyRequestsError represents a 429 Too Many Requests error
type TooManyRequestsError struct {
	ServiceNowError
}

// ServerError represents a 5XX Server error
type ServerError struct {
	ServiceNowError
}

// NewServiceNowError instantiates a new Service-Now error
func NewServiceNowError() *ServiceNowError {
	return &ServiceNowError{
		NewBaseModel(),
	}
}

// CreateServiceNowErrorFromDiscriminatorValue is a parsable factory for creating a ServiceNowError
func CreateServiceNowErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceNowError(), nil
}

// CreateBadRequestErrorFromDiscriminatorValue creates a BadRequestError
func CreateBadRequestErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &BadRequestError{*NewServiceNowError()}, nil
}

// CreateUnauthorizedErrorFromDiscriminatorValue creates a UnauthorizedError
func CreateUnauthorizedErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &UnauthorizedError{*NewServiceNowError()}, nil
}

// CreateForbiddenErrorFromDiscriminatorValue creates a ForbiddenError
func CreateForbiddenErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &ForbiddenError{*NewServiceNowError()}, nil
}

// CreateNotFoundErrorFromDiscriminatorValue creates a NotFoundError
func CreateNotFoundErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &NotFoundError{*NewServiceNowError()}, nil
}

// CreateTooManyRequestsErrorFromDiscriminatorValue creates a TooManyRequestsError
func CreateTooManyRequestsErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &TooManyRequestsError{*NewServiceNowError()}, nil
}

// CreateServerErrorFromDiscriminatorValue creates a ServerError
func CreateServerErrorFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return &ServerError{*NewServiceNowError()}, nil
}

// DefaultErrorMapping returns the standard error mappings for Service-Now APIs.
func DefaultErrorMapping() abstractions.ErrorMappings {
	return abstractions.ErrorMappings{
		"400": CreateBadRequestErrorFromDiscriminatorValue,
		"401": CreateUnauthorizedErrorFromDiscriminatorValue,
		"403": CreateForbiddenErrorFromDiscriminatorValue,
		"404": CreateNotFoundErrorFromDiscriminatorValue,
		"429": CreateTooManyRequestsErrorFromDiscriminatorValue,
		"5XX": CreateServerErrorFromDiscriminatorValue,
		"XXX": CreateServiceNowErrorFromDiscriminatorValue,
	}
}

// Serialize writes the objects properties to the current writer.
func (exc *ServiceNowError) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(exc) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[MainErrorable](errorKey)(exc.GetError),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (exc *ServiceNowError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		errorKey: internalSerialization.DeserializeObjectValueFunc[MainErrorable](CreateMainErrorFromDiscriminatorValue)(exc.setError),
	}
}

// GetError returns the main error
func (exc *ServiceNowError) GetError() (MainErrorable, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceNowError, MainErrorable](exc, errorKey)
}

// setError sets the main error
func (exc *ServiceNowError) setError(mainError MainErrorable) error {
	return store.DefaultBackedModelMutatorFunc(exc, errorKey, mainError)
}

func (exc *ServiceNowError) Error() string {
	mainErr, _ := exc.GetError()
	msg, _ := mainErr.GetMessage()
	if msg != nil {
		return *msg
	}
	details, _ := mainErr.GetDetail()
	return *details
}
