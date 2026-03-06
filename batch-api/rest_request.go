package batchapi

import (
	"encoding/base64"
	"errors"
	"strings"

	u "net/url"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
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
	newInternal.Model
}

// NewRestRequest creates a new rest request
func NewRestRequest() *RestRequestModel {
	return &RestRequestModel{
		newInternal.NewBaseModel(),
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
	if internal.IsNil(rE) {
		return nil
	}

	serializers := []func(serialization.SerializationWriter) error{
		func(sw serialization.SerializationWriter) error {
			body, err := rE.GetBody()
			if err != nil {
				return err
			}

			encodedBody := base64.StdEncoding.EncodeToString(body)

			return sw.WriteStringValue(bodyKey, &encodedBody)
		},
		func(sw serialization.SerializationWriter) error {
			excludeResponseHeaders, err := rE.GetExcludeResponseHeaders()
			if err != nil {
				return err
			}

			return sw.WriteBoolValue(excludeResponseHeadersKey, excludeResponseHeaders)
		},
		func(sw serialization.SerializationWriter) error {
			headers, err := rE.GetHeaders()
			if err != nil {
				return err
			}

			// Create a new slice of serialization.Parsable
			parsableHeaders := make([]serialization.Parsable, len(headers))
			for i, header := range headers {
				parsableHeaders[i] = header
			}

			return sw.WriteCollectionOfObjectValues(headersKey, parsableHeaders)
		},
		func(sw serialization.SerializationWriter) error {
			id, err := rE.GetID()
			if err != nil {
				return err
			}

			// ensure request has an id BEFORE serializing
			if internal.IsNil(id) || *id == "" {
				idString := uuid.NewString()
				id = &idString
			}

			return sw.WriteStringValue(idKey, id)
		},
		func(sw serialization.SerializationWriter) error {
			method, err := rE.GetMethod()
			if err != nil {
				return err
			}
			if internal.IsNil(method) {
				return errors.New("method can't be nil")
			}

			strMethod := (*method).String()

			return sw.WriteStringValue(methodKey, &strMethod)
		},
		func(sw serialization.SerializationWriter) error {
			url, err := rE.GetURL()
			if err != nil {
				return err
			}

			return sw.WriteStringValue(urlKey, url)
		},
	}

	for _, serializer := range serializers {
		if err := serializer(writer); err != nil {
			return err
		}
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *RestRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(rE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		bodyKey: internalSerialization.DeserializeMutatedStringFunc(rE.SetBody, func(s *string) ([]byte, error) {
			if s == nil {
				return nil, nil
			}
			return base64.StdEncoding.DecodeString(*s)
		}),
		excludeResponseHeadersKey: internalSerialization.DeserializeBoolFunc(rE.SetExcludeResponseHeaders),
		headersKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc(rE.SetHeaders, CreateRestRequestHeaderFromDiscriminatorValue),
		idKey: internalSerialization.DeserializeStringFunc(rE.SetID),
		methodKey: internalSerialization.DeserializeMutatedStringFunc(rE.SetMethod, func(s *string) (*abstractions.HttpMethod, error) {
			if s == nil {
				return nil, nil
			}
			m, err := parseHttpMethod(*s)
			return &m, err
		}),
		urlKey: internalSerialization.DeserializeStringFunc(rE.SetURL),
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
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []byte](backingStore, bodyKey)
}

// SetBodyFromParsable serializes the provided parsable and sets the output to the request's body.
func (rE *RestRequestModel) SetBodyFromParsable(contentType string, parsable serialization.Parsable) error {
	if internal.IsNil(rE) {
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

	if internal.IsNil(headers) {
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
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, bodyKey, body)
}

// GetExcludeResponseHeaders returns if the request will exclude response headers.
func (rE *RestRequestModel) GetExcludeResponseHeaders() (*bool, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, excludeResponseHeadersKey)
}

// SetExcludeResponseHeaders set if to include or exclude response headers.
func (rE *RestRequestModel) SetExcludeResponseHeaders(excludeResponseHeaders *bool) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, excludeResponseHeadersKey, excludeResponseHeaders)
}

// GetHeaders returns the headers of the request.
func (rE *RestRequestModel) GetHeaders() ([]RestRequestHeader, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []RestRequestHeader](backingStore, headersKey)
}

// SetHeaders sets the headers for the request.
func (rE *RestRequestModel) SetHeaders(headers []RestRequestHeader) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, headersKey, headers)
}

// GetID returns the id of the request.
func (rE *RestRequestModel) GetID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, idKey)
}

// SetID sets the id of the request.
func (rE *RestRequestModel) SetID(id *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, idKey, id)
}

// GetMethod returns the method of the request
func (rE *RestRequestModel) GetMethod() (*abstractions.HttpMethod, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *abstractions.HttpMethod](backingStore, methodKey)
}

// SetMethod sets the method of the request
func (rE *RestRequestModel) SetMethod(method *abstractions.HttpMethod) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, methodKey, method)
}

// GetURL returns the relative URL of the request.
func (rE *RestRequestModel) GetURL() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, urlKey)
}

// SetURL sets the URL of the request (if not relative, will be converted).
func (rE *RestRequestModel) SetURL(url *string) error {
	if internal.IsNil(rE) {
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

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, urlKey, &relativeURL)
}
