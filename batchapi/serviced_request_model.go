package batchapi

import (
	"encoding/base64"
	"math"
	"reflect"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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

// ServicedRequestModel represents Service-Now Batch API response's serviced request.
type ServicedRequestModel struct {
	core.BackedModel
}

// NewServicedRequest instantiates a new ServicedRequest.
func NewServicedRequest() *ServicedRequestModel {
	return &ServicedRequestModel{
		core.NewBaseModel(),
	}
}

// CreateServicedRequestFromDiscriminatorValue is a parsable factory for creating a BatchResponse.
func CreateServicedRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServicedRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (sR *ServicedRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(sR) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeMutatedStringFunc(bodyKey, func(body []byte) (*string, error) {
			if body == nil {
				return nil, nil
			}
			encodedBody := base64.StdEncoding.EncodeToString(body)
			return &encodedBody, nil
		}, sR.GetBody),
		internalSerialization.SerializeStringFunc(errorMessageKey, sR.GetErrorMessage),
		internalSerialization.SerializeISODurationFunc(executionTimeKey, sR.GetExecutionTime),
		internalSerialization.SerializeCollectionOfObjectValuesFunc[RestRequestHeader](headersKey, sR.GetHeaders),
		internalSerialization.SerializeStringFunc(idKey, sR.GetID),
		internalSerialization.SerializeStringFunc(redirectURLKey, sR.GetRedirectURL),
		internalSerialization.SerializeInt64Func(statusCodeKey, sR.GetStatusCode),
		internalSerialization.SerializeStringFunc(statusTextKey, sR.GetStatusText),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (sR *ServicedRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if conversion.IsNil(sR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		bodyKey: internalSerialization.DeserializeMutatedStringFunc(func(s *string) ([]byte, error) {
			if s == nil {
				return nil, nil
			}
			return base64.StdEncoding.DecodeString(*s)
		}, sR.setBody),
		errorMessageKey: internalSerialization.DeserializeStringFunc(sR.setErrorMessage),
		executionTimeKey: func(pn serialization.ParseNode) error {
			// ServiceNow returns execution_time as a number (milliseconds)
			// GetISODurationValue fails if it's a number in the JSON tree.
			floatTime, err := pn.GetFloat64Value()
			if err == nil && floatTime != nil {
				seconds := int(*floatTime / 1000)
				milliseconds := int(math.Mod(*floatTime, 1000))
				isoDuration := serialization.NewDuration(0, 0, 0, 0, 0, seconds, milliseconds)
				return sR.setExecutionTime(isoDuration)
			}

			duration, err := pn.GetISODurationValue()
			if err != nil {
				return err
			}

			return sR.setExecutionTime(duration)
		},
		headersKey:     internalSerialization.DeserializeCollectionOfObjectValuesFunc[RestRequestHeader](CreateRestRequestHeaderFromDiscriminatorValue, sR.setHeaders),
		idKey:          internalSerialization.DeserializeStringFunc(sR.setID),
		redirectURLKey: internalSerialization.DeserializeStringFunc(sR.setRedirectURL),
		statusCodeKey: func(pn serialization.ParseNode) error {
			// ServiceNow sometimes returns status_code as a number that gets parsed as float64
			// GetInt64Value fails if it's a float64 in the JSON tree.
			floatCode, err := pn.GetFloat64Value()
			if err != nil {
				return err
			}

			if floatCode != nil {
				statusCode := int64(*floatCode)
				return sR.setStatusCode(&statusCode)
			}

			return nil
		},
		statusTextKey: internalSerialization.DeserializeStringFunc(sR.setStatusText),
	}
}

// GetBodyAsParsable returns the body serialized.
func (sR *ServicedRequestModel) GetBodyAsParsable(constructor serialization.ParsableFactory) (serialization.Parsable, error) {
	if conversion.IsNil(sR) {
		return nil, nil
	}

	if err := throwErrors(sR, reflect.TypeOf(constructor).String()); err != nil {
		return nil, err
	}

	headers, err := sR.GetHeaders()
	if err != nil {
		return nil, err
	}

	body, err := sR.GetBody()
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}

	contentType := getHTTPHeader(headers, internalhttp.HTTPHeaderContentType, "")

	return serializeContent[serialization.Parsable](contentType, body, constructor)
}

// GetBody returns the raw body for the batch item.
func (sR *ServicedRequestModel) GetBody() ([]byte, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, []byte](sR, bodyKey)
}

// setBody sets the raw body for the batch item.
func (sR *ServicedRequestModel) setBody(body []byte) error {
	return store.DefaultBackedModelMutatorFunc(sR, bodyKey, body)
}

// GetErrorMessage returns, if present, the error messages.
func (sR *ServicedRequestModel) GetErrorMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, *string](sR, errorMessageKey)
}

// setErrorMessage sets the error messages.
func (sR *ServicedRequestModel) setErrorMessage(errorMessage *string) error {
	return store.DefaultBackedModelMutatorFunc(sR, errorMessageKey, errorMessage)
}

// GetExecutionTime returns time it took to execute the batch item request.
func (sR *ServicedRequestModel) GetExecutionTime() (*serialization.ISODuration, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, *serialization.ISODuration](sR, executionTimeKey)
}

// setExecutionTime sets the time it took to execute the batch item request.
func (sR *ServicedRequestModel) setExecutionTime(executionTime *serialization.ISODuration) error {
	return store.DefaultBackedModelMutatorFunc(sR, executionTimeKey, executionTime)
}

// GetHeaders returns headers for the batch item.
func (sR *ServicedRequestModel) GetHeaders() ([]RestRequestHeader, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, []RestRequestHeader](sR, headersKey)
}

// setHeaders sets headers for the batch item.
func (sR *ServicedRequestModel) setHeaders(headers []RestRequestHeader) error {
	return store.DefaultBackedModelMutatorFunc(sR, headersKey, headers)
}

// GetID returns ID of the batch item that matches the `rest_requests.id` parameter in the request.
func (sR *ServicedRequestModel) GetID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, *string](sR, idKey)
}

// setID sets the id of the batch item.
func (sR *ServicedRequestModel) setID(id *string) error {
	return store.DefaultBackedModelMutatorFunc(sR, idKey, id)
}

// GetRedirectURL, if present, returns redirect url for batch item.
func (sR *ServicedRequestModel) GetRedirectURL() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, *string](sR, redirectURLKey)
}

// setRedirectURL sets redirect url for batch item.
func (sR *ServicedRequestModel) setRedirectURL(redirectURL *string) error {
	return store.DefaultBackedModelMutatorFunc(sR, redirectURLKey, redirectURL)
}

// GetStatusCode returns status code for batch item.
func (sR *ServicedRequestModel) GetStatusCode() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, *int64](sR, statusCodeKey)
}

// setStatusCode sets status code for batch item.
func (sR *ServicedRequestModel) setStatusCode(statusCode *int64) error {
	return store.DefaultBackedModelMutatorFunc(sR, statusCodeKey, statusCode)
}

// GetStatusText returns status text for batch item.
func (sR *ServicedRequestModel) GetStatusText() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServicedRequestModel, *string](sR, statusTextKey)
}

// setStatusText sets status text for batch item.
func (sR *ServicedRequestModel) setStatusText(statusText *string) error {
	return store.DefaultBackedModelMutatorFunc(sR, statusTextKey, statusText)
}
