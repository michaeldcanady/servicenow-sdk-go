package batchapi

import (
	"encoding/base64"
	"errors"
	"strings"

	u "net/url"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	excludeResponseHeadersKey = "exclude_response_headers"
	methodKey                 = "method"
	urlKey                    = "url"
)

// RestRequestable represents Service-Now Batch API request's request
type RestRequestable interface {
	GetBody() ([]byte, error)
	SetBodyFromParsable(string, serialization.Parsable) error
	SetBody([]byte) error
	GetExcludeResponseHeaders() (*bool, error)
	SetExcludeResponseHeaders(*bool) error
	GetHeaders() ([]BatchHeaderable, error)
	SetHeaders([]BatchHeaderable) error
	GetID() (*string, error)
	SetID(*string) error
	GetMethod() (*abstractions.HttpMethod, error)
	SetMethod(*abstractions.HttpMethod) error
	GetURL() (*string, error)
	SetURL(*string) error
	serialization.Parsable
	store.BackedModel
}

// RestRequest implementation of RestRequestable
type RestRequest struct {
	newInternal.Model
}

// NewRestRequest creates a new rest request
func NewRestRequest() *RestRequest {
	return &RestRequest{
		newInternal.NewBaseModel(),
	}
}

// CreateRestRequestFromDiscriminatorValue is a parsable factory for creating a BatchRequestable
func CreateRestRequestFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewRestRequest(), nil
}

// headers support headers types
type headers interface {
	*abstractions.RequestHeaders
}

// createBatchableHeadersFromHeaders converts headers to BatchHeaderable
func createBatchableHeadersFromHeaders[h headers](headers h) ([]BatchHeaderable, error) {
	batchHeaders := make([]BatchHeaderable, 0)

	if requestHeaders, ok := interface{}(headers).(*abstractions.RequestHeaders); ok {
		for _, key := range requestHeaders.ListKeys() {
			batchHeader := NewBatchHeader()
			values := requestHeaders.Get(key)
			if err := batchHeader.SetName(&key); err != nil {
				return nil, err
			}
			valuesString := strings.Join(values, ", ")
			if err := batchHeader.SetValue(&valuesString); err != nil {
				return nil, err
			}
			batchHeaders = append(batchHeaders, batchHeader)
		}
		return batchHeaders, nil
	}

	return nil, nil
}

// CreateRestRequestFromRequestInformation
func CreateRestRequestFromRequestInformation(requestInfo *abstractions.RequestInformation, excludeResponseHeaders bool) (RestRequestable, error) {
	request := NewRestRequest()
	if err := request.SetBody(requestInfo.Content); err != nil {
		return nil, err
	}
	headers, err := createBatchableHeadersFromHeaders(requestInfo.Headers)
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
func (rE *RestRequest) Serialize(writer serialization.SerializationWriter) error {
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
func (rE *RestRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(rE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		bodyKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (bodyKey) not implemented")
		},
		excludeResponseHeadersKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (excludeResponseHeadersKey) not implemented")
		},
		headersKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (headersKey) not implemented")
		},
		methodKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (methodKey) not implemented")
		},
		urlKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (urlKey) not implemented")
		},
	}
}

// GetBody returns the requests body in bytes.
func (rE *RestRequest) GetBody() ([]byte, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	body, err := rE.GetBackingStore().Get(bodyKey)
	if err != nil {
		return nil, err
	}

	typedBody, ok := body.([]byte)
	if !ok {
		return nil, errors.New("body is not []byte")
	}

	return typedBody, nil
}

// SetBodyFromParsable serializes the provided parsable and sets the output to the request's body.
func (rE *RestRequest) SetBodyFromParsable(contentType string, parsable serialization.Parsable) error {
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
		headers = make([]BatchHeaderable, 0)
	}

	if err := writer.WriteObjectValue("", parsable); err != nil {
		return err
	}

	content, err := writer.GetSerializedContent()
	if err != nil {
		return err
	}

	batchHeader := NewBatchHeader()
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
func (rE *RestRequest) SetBody(body []byte) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(bodyKey, body)
}

// GetExcludeResponseHeaders returns if the request will exclude response headers.
func (rE *RestRequest) GetExcludeResponseHeaders() (*bool, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	excludeResponseHeaders, err := rE.GetBackingStore().Get(excludeResponseHeadersKey)
	if err != nil {
		return nil, err
	}

	typedExcludeResponseHeaders, ok := excludeResponseHeaders.(*bool)
	if !ok {
		return nil, errors.New("excludeResponseHeaders is not *bool")
	}

	return typedExcludeResponseHeaders, nil
}

// SetExcludeResponseHeaders set if to include or exclude response headers.
func (rE *RestRequest) SetExcludeResponseHeaders(excludeResponseHeaders *bool) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(excludeResponseHeadersKey, excludeResponseHeaders)
}

// GetHeaders returns the headers of the request.
func (rE *RestRequest) GetHeaders() ([]BatchHeaderable, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	headers, err := rE.GetBackingStore().Get(headersKey)
	if err != nil {
		return nil, err
	}

	typedheaders, ok := headers.([]BatchHeaderable)
	if !ok {
		return nil, errors.New("headers is not []BatchHeaderable")
	}

	return typedheaders, nil
}

// SetHeaders sets the headers for the request.
func (rE *RestRequest) SetHeaders(headers []BatchHeaderable) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(headersKey, headers)
}

// GetID returns the id of the request.
func (rE *RestRequest) GetID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	id, err := rE.GetBackingStore().Get(idKey)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*string)
	if !ok {
		return nil, errors.New("id is not *string")
	}

	return typedID, nil
}

// SetID sets the id of the request.
func (rE *RestRequest) SetID(id *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(idKey, id)
}

// GetMethod returns the method of the request
func (rE *RestRequest) GetMethod() (*abstractions.HttpMethod, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	method, err := rE.GetBackingStore().Get(methodKey)
	if err != nil {
		return nil, err
	}

	typedMethod, ok := method.(*abstractions.HttpMethod)
	if !ok {
		return nil, errors.New("method is not *abstractions.HttpMethod")
	}

	return typedMethod, nil
}

// SetMethod sets the method of the request
func (rE *RestRequest) SetMethod(method *abstractions.HttpMethod) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(methodKey, method)
}

// GetURL returns the relative URL of the request.
func (rE *RestRequest) GetURL() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	url, err := rE.GetBackingStore().Get(urlKey)
	if err != nil {
		return nil, err
	}

	typedURL, ok := url.(*string)
	if !ok {
		return nil, errors.New("url is not *string")
	}

	return typedURL, nil
}

// SetURL sets the URL of the request (if not relative, will be converted).
func (rE *RestRequest) SetURL(url *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	parsedURL, err := u.Parse(*url)
	if err != nil {
		return err
	}

	// Ensure the URL is relative
	if parsedURL.IsAbs() {
		parsedURL.Scheme = ""
		parsedURL.Host = ""
	}

	relativeURL := parsedURL.String()

	// Ensure the path begins with "/api"
	if !strings.HasPrefix(relativeURL, "/api") {
		return errors.New("invalid URL: path doesn't begin with \"/api\"")
	}

	return rE.GetBackingStore().Set(urlKey, &relativeURL)
}
