package batchapi

import (
	"encoding/base64"
	"errors"
	"reflect"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	bodyKey          = "body"
	errorMessageKey  = "error_message"
	executionTimeKey = "execution_time"
	headersKey       = "headers"
	idKey            = "id"
	redirectURLKey   = "redirect_url"
	statusCodeKey    = "status_code"
	statusTextKey    = "status_text"
)

// ServicedRequestable represents Service-Now Batch API response's serviced request.
type ServicedRequestable interface {
	GetBody(serialization.ParsableFactory) (serialization.Parsable, error)
	GetRawBody() ([]byte, error)
	setRawBody([]byte) error
	GetErrorMessage() (*string, error)
	setErrorMessage(*string) error
	GetExecutionTime() (*serialization.ISODuration, error)
	setExecutionTime(*serialization.ISODuration) error
	GetHeaders() ([]BatchHeaderable, error)
	setHeaders([]BatchHeaderable) error
	GetID() (*string, error)
	setID(*string) error
	GetRedirectURL() (*string, error)
	setRedirectURL(*string) error
	GetStatusCode() (*int64, error)
	setStatusCode(*int64) error
	GetStatusText() (*string, error)
	setStatusText(*string) error
	serialization.Parsable
	store.BackedModel
}

// ServicedRequest represents Service-Now Batch API response's serviced request.
type ServicedRequest struct {
	newInternal.Model
}

// NewServicedRequest instantiates a new ServicedRequest.
func NewServicedRequest() *ServicedRequest {
	return &ServicedRequest{
		newInternal.NewBaseModel(),
	}
}

// CreateServicedRequestFromDiscriminatorValue is a parsable factory for creating a BatchResponseable.
func CreateServicedRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServicedRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (sR *ServicedRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(sR) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (sR *ServicedRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(sR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		bodyKey: func(pn serialization.ParseNode) error {
			encodedBody, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			body, err := base64.StdEncoding.DecodeString(*encodedBody)
			if err != nil {
				return err
			}

			return sR.setRawBody(body)
		},
		errorMessageKey: func(pn serialization.ParseNode) error {
			errorMessage, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return sR.setErrorMessage(errorMessage)
		},
		executionTimeKey: func(pn serialization.ParseNode) error {
			duration, err := pn.GetISODurationValue()
			if err != nil {
				return err
			}

			return sR.setExecutionTime(duration)
		},
		headersKey: func(pn serialization.ParseNode) error {
			headers, err := pn.GetCollectionOfObjectValues(CreateBatchHeader2FromDiscriminatorValue)
			if err != nil {
				return err
			}

			batchHeaders := make([]BatchHeaderable, len(headers))
			for index, header := range headers {
				batchHeader, ok := header.(BatchHeaderable)
				if !ok {
					return errors.New("header is not BatchHeaderable")
				}
				batchHeaders[index] = batchHeader
			}

			return sR.setHeaders(batchHeaders)
		},
		idKey: func(pn serialization.ParseNode) error {
			id, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return sR.setID(id)
		},
		redirectURLKey: func(pn serialization.ParseNode) error {
			redirectURL, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return sR.setRedirectURL(redirectURL)
		},
		statusCodeKey: func(pn serialization.ParseNode) error {
			statusCode, err := pn.GetInt64Value()
			if err != nil {
				return err
			}

			return sR.setStatusCode(statusCode)
		},
		statusTextKey: func(pn serialization.ParseNode) error {
			statusText, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return sR.setStatusText(statusText)
		},
	}
}

// GetBody returns the body serialized.
func (sR *ServicedRequest) GetBody(constructor serialization.ParsableFactory) (serialization.Parsable, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	if err := throwErrors(sR, reflect.TypeOf(constructor).Elem().Name()); err != nil {
		return nil, err
	}

	headers, err := sR.GetHeaders()
	if err != nil {
		return nil, err
	}

	body, err := sR.GetRawBody()
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}

	contentType := getHTTPHeader(headers, newInternal.HTTPHeaderContentType, "")

	return serializeContent[serialization.Parsable](contentType, body, constructor)
}

// GetRawBody returns the raw body for the batch item.
func (sR *ServicedRequest) GetRawBody() ([]byte, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	body, err := sR.GetBackingStore().Get(bodyKey)
	if err != nil {
		return nil, err
	}

	typedBody, ok := body.([]byte)
	if !ok {
		return nil, errors.New("body is not []byte")
	}

	return typedBody, nil
}

// setRawBody sets the raw body for the batch item.
func (sR *ServicedRequest) setRawBody(body []byte) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(bodyKey, body)
}

// GetErrorMessage returns, if present, the error messages.
func (sR *ServicedRequest) GetErrorMessage() (*string, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	message, err := sR.GetBackingStore().Get(errorMessageKey)
	if err != nil {
		return nil, err
	}

	stringMessage, ok := message.(*string)
	if !ok {
		return nil, errors.New("message is not *string")
	}

	return stringMessage, nil
}

// setErrorMessage sets the error messages.
func (sR *ServicedRequest) setErrorMessage(errorMessage *string) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(errorMessageKey, errorMessage)
}

// GetExecutionTime returns time it took to execute the batch item request.
func (sR *ServicedRequest) GetExecutionTime() (*serialization.ISODuration, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	executionTime, err := sR.GetBackingStore().Get(executionTimeKey)
	if err != nil {
		return nil, err
	}

	typedExecutionTime, ok := executionTime.(*serialization.ISODuration)
	if !ok {
		return nil, errors.New("executionTime is not *serialization.ISODuration")
	}

	return typedExecutionTime, nil
}

// setExecutionTime sets the time it took to execute the batch item request.
func (sR *ServicedRequest) setExecutionTime(executionTime *serialization.ISODuration) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(executionTimeKey, executionTime)
}

// GetHeaders returns headers for the batch item.
func (sR *ServicedRequest) GetHeaders() ([]BatchHeaderable, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	headers, err := sR.GetBackingStore().Get(headersKey)
	if err != nil {
		return nil, err
	}

	typedHeaders, ok := headers.([]BatchHeaderable)
	if !ok {
		return nil, errors.New("headers is not []BatchHeaderable")
	}

	return typedHeaders, nil
}

// setHeaders sets headers for the batch item.
func (sR *ServicedRequest) setHeaders(headers []BatchHeaderable) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(headersKey, headers)
}

// GetID returns ID of the batch item that matches the `rest_requests.id` parameter in the request.
func (sR *ServicedRequest) GetID() (*string, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	id, err := sR.GetBackingStore().Get(idKey)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*string)
	if !ok {
		return nil, errors.New("id is not *string")
	}

	return typedID, nil
}

// setID sets the id of the batch item.
func (sR *ServicedRequest) setID(id *string) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(idKey, id)
}

// GetRedirectURL, if present, returns redirect url for batch item.
func (sR *ServicedRequest) GetRedirectURL() (*string, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	redirectURL, err := sR.GetBackingStore().Get(redirectURLKey)
	if err != nil {
		return nil, err
	}

	typedRedirectURL, ok := redirectURL.(*string)
	if !ok {
		return nil, errors.New("redirectURL is not *string")
	}

	return typedRedirectURL, nil
}

// setRedirectURL sets redirect url for batch item.
func (sR *ServicedRequest) setRedirectURL(redirectURL *string) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(redirectURLKey, redirectURL)
}

// GetStatusCode returns status code for batch item.
func (sR *ServicedRequest) GetStatusCode() (*int64, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	statusCode, err := sR.GetBackingStore().Get(statusCodeKey)
	if err != nil {
		return nil, err
	}

	typedStatusCode, ok := statusCode.(*int64)
	if !ok {
		return nil, errors.New("statusCode is not *int64")
	}

	return typedStatusCode, nil
}

// setStatusCode sets status code for batch item.
func (sR *ServicedRequest) setStatusCode(statusCode *int64) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(statusCodeKey, statusCode)
}

// GetStatusText returns status text for batch item.
func (sR *ServicedRequest) GetStatusText() (*string, error) {
	if internal.IsNil(sR) {
		return nil, nil
	}

	statusText, err := sR.GetBackingStore().Get(statusTextKey)
	if err != nil {
		return nil, err
	}

	typedStatusText, ok := statusText.(*string)
	if !ok {
		return nil, errors.New("statusCode is not *string")
	}

	return typedStatusText, nil
}

// setStatusText sets status text for batch item.
func (sR *ServicedRequest) setStatusText(statusText *string) error {
	if internal.IsNil(sR) {
		return nil
	}

	return sR.GetBackingStore().Set(statusTextKey, statusText)
}
