package batchapi

import (
	"encoding/base64"
	"errors"
	u "net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	excludeResponseHeadersKey = "exclude_response_headers"
	methodKey                 = "method"
	urlKey                    = "url"
)

// RestRequest represents Service-Now Batch API request's request
type RestRequest interface {
	GetBody() ([]byte, error)
	SetBodyFromParsable(string, serialization.Parsable) error
	SetBody([]byte) error
	GetExcludeResponseHeaders() (*bool, error)
	SetExcludeResponseHeaders(*bool) error
	GetHeaders() ([]RestRequestHeader, error)
	SetHeaders([]RestRequestHeader) error
	GetID() (*string, error)
	SetID(*string) error
	GetMethod() (*abstractions.HttpMethod, error)
	SetMethod(*abstractions.HttpMethod) error
	GetURL() (*string, error)
	SetURL(*string) error
	serialization.Parsable
	kiotaStore.BackedModel
}

// RestRequestModel implementation of RestRequestable
type RestRequestModel struct {
	model.Model
}

// NewRestRequest creates a new rest request
func NewRestRequest() *RestRequestModel {
	return &RestRequestModel{
		model.NewBaseModel(),
	}
}

// CreateRestRequestFromDiscriminatorValue is a parsable factory for creating a BatchRequest
func CreateRestRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRestRequest(), nil
}

// CreateRestRequestFromRequestInformation
func CreateRestRequestFromRequestInformation(requestInfo *abstractions.RequestInformation, excludeResponseHeaders bool) (RestRequest, error) {
	request := NewRestRequest()
	if err := request.SetBody(requestInfo.Content); err != nil {
		return nil, err
	}
	headers, err := createRestRequestHeaderFromHeaders(requestInfo.Headers)
	if err != nil {
		return nil, err
	}
	if err := request.SetHeaders(headers); err != nil {
		return nil, err
	}
	if err := request.SetExcludeResponseHeaders(&excludeResponseHeaders); err != nil {
		return nil, err
	}
	if err := request.SetMethod(&requestInfo.Method); err != nil {
		return nil, err
	}
	uri, err := requestInfo.GetUri()
	if err != nil {
		return nil, err
	}
	uriString := uri.String()
	if err := request.SetURL(&uriString); err != nil {
		return nil, err
	}
	newID := uuid.NewString()
	if err := request.SetID(&newID); err != nil {
		return nil, err
	}

	return request, nil
}

// Serialize writes the objects properties to the current writer.
func (rE *RestRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(rE) {
		return nil
	}

	return kiota.Serialize(writer,
		kiota.SerializeMutatedStringFunc(bodyKey, func(body []byte) (*string, error) {
			encodedBody := base64.StdEncoding.EncodeToString(body)
			return &encodedBody, nil
		})(rE.GetBody),
		kiota.SerializeBoolFunc(excludeResponseHeadersKey)(rE.GetExcludeResponseHeaders),
		kiota.SerializeCollectionOfObjectValuesFunc[RestRequestHeader](headersKey)(rE.GetHeaders),
		kiota.SerializeStringFunc(idKey)(func() (*string, error) {
			id, err := rE.GetID()
			if err != nil {
				return nil, err
			}

			// ensure request has an id BEFORE serializing
			if utils.IsNil(id) || *id == "" {
				idString := uuid.NewString()
				id = &idString
			}

			return id, nil
		}),
		kiota.SerializeMutatedStringFunc(methodKey, func(method *abstractions.HttpMethod) (*string, error) {
			if utils.IsNil(method) {
				return nil, errors.New("method can't be nil")
			}
			strMethod := (*method).String()
			return &strMethod, nil
		})(rE.GetMethod),
		kiota.SerializeStringFunc(urlKey)(rE.GetURL),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *RestRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if utils.IsNil(rE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		bodyKey: kiota.DeserializeMutatedStringFunc(func(s *string) ([]byte, error) {
			if s == nil {
				return nil, nil
			}
			return base64.StdEncoding.DecodeString(*s)
		})(rE.SetBody),
		excludeResponseHeadersKey: kiota.DeserializeBoolFunc()(rE.SetExcludeResponseHeaders),
		headersKey:                kiota.DeserializeCollectionOfObjectValuesFunc[RestRequestHeader](CreateRestRequestHeaderFromDiscriminatorValue)(rE.SetHeaders),
		idKey:                     kiota.DeserializeStringFunc(rE.SetID),
		methodKey: kiota.DeserializeMutatedStringFunc(func(s *string) (*abstractions.HttpMethod, error) {
			if s == nil {
				return nil, nil
			}
			m, err := parseHttpMethod(*s)
			return &m, err
		})(rE.SetMethod),
		urlKey: internalSerialization.DeserializeStringFunc()(rE.SetURL),
	}
}

func parseHttpMethod(method string) (abstractions.HttpMethod, error) {
	switch strings.ToUpper(method) {
	case "GET":
		return abstractions.GET, nil
	case "POST":
		return abstractions.POST, nil
	case "PATCH":
		return abstractions.PATCH, nil
	case "DELETE":
		return abstractions.DELETE, nil
	case "OPTIONS":
		return abstractions.OPTIONS, nil
	case "CONNECT":
		return abstractions.CONNECT, nil
	case "PUT":
		return abstractions.PUT, nil
	case "TRACE":
		return abstractions.TRACE, nil
	case "HEAD":
		return abstractions.HEAD, nil
	default:
		return 0, errors.New("invalid HTTP method")
	}
}

// GetBody returns the requests body in bytes.
func (rE *RestRequestModel) GetBody() ([]byte, error) {
	if utils.IsNil(rE) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []byte](rE.GetBackingStore(), bodyKey)
}

// SetBodyFromParsable serializes the provided parsable and sets the output to the request's body.
func (rE *RestRequestModel) SetBodyFromParsable(contentType string, parsable serialization.Parsable) error {
	if utils.IsNil(rE) {
		return nil
	}

	registry := serialization.DefaultSerializationWriterFactoryInstance

	writer, err := registry.GetSerializationWriter(contentType)
	if err != nil {
		return err
	}

	headers, err := rE.GetHeaders()
	if err != nil {
		return err
	}

	if utils.IsNil(headers) {
		headers = make([]RestRequestHeader, 0)
	}

	if err := writer.WriteObjectValue("", parsable); err != nil {
		return err
	}

	content, err := writer.GetSerializedContent()
	if err != nil {
		return err
	}

	batchHeader := NewRestRequestHeader()
	// TODO: add to RequestHeader
	name := "Content-Type"
	if err := batchHeader.SetName(&name); err != nil {
		return err
	}
	if err := batchHeader.SetValue(&contentType); err != nil {
		return err
	}

	headers = append(headers, batchHeader)
	if err := rE.SetHeaders(headers); err != nil {
		return err
	}

	return rE.SetBody(content)
}

// SetBody sets the requests body to the provided content.
func (rE *RestRequestModel) SetBody(body []byte) error {
	if utils.IsNil(rE) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), bodyKey, body)
}

// GetExcludeResponseHeaders returns if the request will exclude response headers.
func (rE *RestRequestModel) GetExcludeResponseHeaders() (*bool, error) {
	if utils.IsNil(rE) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](rE.GetBackingStore(), excludeResponseHeadersKey)
}

// SetExcludeResponseHeaders set if to include or exclude response headers.
func (rE *RestRequestModel) SetExcludeResponseHeaders(excludeResponseHeaders *bool) error {
	if utils.IsNil(rE) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), excludeResponseHeadersKey, excludeResponseHeaders)
}

// GetHeaders returns the headers of the request.
func (rE *RestRequestModel) GetHeaders() ([]RestRequestHeader, error) {
	if utils.IsNil(rE) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []RestRequestHeader](rE.GetBackingStore(), headersKey)
}

// SetHeaders sets the headers for the request.
func (rE *RestRequestModel) SetHeaders(headers []RestRequestHeader) error {
	if utils.IsNil(rE) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), headersKey, headers)
}

// GetID returns the id of the request.
func (rE *RestRequestModel) GetID() (*string, error) {
	if utils.IsNil(rE) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](rE.GetBackingStore(), idKey)
}

// SetID sets the id of the request.
func (rE *RestRequestModel) SetID(id *string) error {
	if utils.IsNil(rE) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), idKey, id)
}

// GetMethod returns the method of the request
func (rE *RestRequestModel) GetMethod() (*abstractions.HttpMethod, error) {
	if utils.IsNil(rE) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *abstractions.HttpMethod](rE.GetBackingStore(), methodKey)
}

// SetMethod sets the method of the request
func (rE *RestRequestModel) SetMethod(method *abstractions.HttpMethod) error {
	if utils.IsNil(rE) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), methodKey, method)
}

// GetURL returns the relative URL of the request.
func (rE *RestRequestModel) GetURL() (*string, error) {
	if utils.IsNil(rE) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](rE.GetBackingStore(), urlKey)
}

// SetURL sets the URL of the request (if not relative, will be converted).
func (rE *RestRequestModel) SetURL(url *string) error {
	if utils.IsNil(rE) {
		return nil
	}

	parsedURL, err := u.Parse(*url)
	if err != nil {
		return err
	}

	// Ensure the URL is relative
	if parsedURL.IsAbs() || parsedURL.Host != "" {
		parsedURL.Scheme = ""
		parsedURL.Host = ""
	}

	relativeURL := parsedURL.String()

	// Ensure the path begins with "/api"
	if !strings.HasPrefix(relativeURL, "/api") {
		return errors.New("invalid URL: path doesn't begin with \"/api\"")
	}

	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), urlKey, &relativeURL)
}
