package core

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
)

// RequestBuilder represents a builder for constructing HTTP request information.
type RequestBuilder struct {
	// PathParameters is a map of path parameters used in the URL template.
	PathParameters map[string]string

	// Client is an instance of the HTTP client used to send requests.
	Client Client

	// UrlTemplate is the URL template for constructing the request URL.
	UrlTemplate string //nolint:stylecheck
}

// NewRequestBuilder creates a new instance of the RequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created RequestBuilder.
func NewRequestBuilder(client Client, urlTemplate string, pathParameters map[string]string) *RequestBuilder {
	return &RequestBuilder{
		Client:         client,
		UrlTemplate:    urlTemplate,
		PathParameters: pathParameters,
	}
}

func (rB *RequestBuilder) ToPutRequestInformation(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation(PUT, config)
}

func (rB *RequestBuilder) ToPostRequestInformation(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation(POST, config)
}

func (rB *RequestBuilder) ToDeleteRequestInformation(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation(DELETE, config)
}

func (rB *RequestBuilder) ToGetRequestInformation(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation(GET, config)
}

func (rB *RequestBuilder) prepareData(rawData interface{}) ([]byte, error) {
	var data []byte
	var err error

	if rawData == nil {
		return data, nil
	}

	if reflect.TypeOf(rawData) == reflect.TypeOf([]byte{}) {
		return rawData.([]byte), nil
	}

	if reflect.TypeOf(rawData) == reflect.TypeOf(map[string]string{}) {
		data, err = json.Marshal(rawData)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %s", err)
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("unsupported type: %T", rawData)
	}

	return data, nil
}

func (rB *RequestBuilder) ToRequestInformation(method HttpMethod, config *RequestConfiguration) (*RequestInformation, error) {
	requestInfo := NewRequestInformation()

	if config != nil {
		if !isNil(config.QueryParameters) {
			err := requestInfo.AddQueryParameters(config.QueryParameters)
			if err != nil {
				return nil, err
			}
		}
		if !isNil(config.Data) {
			data, err := rB.prepareData(config.Data)
			if err != nil {
				return nil, err
			}
			if len(data) != 0 {
				mime := mimetype.Detect(data)

				requestInfo.Content = data
				requestInfo.Headers.Add("Content-Type", mime.String())
			}
		}
	}

	requestInfo.Method = method
	requestInfo.uri.PathParameters = rB.PathParameters
	requestInfo.uri.UrlTemplate = rB.UrlTemplate

	return requestInfo, nil
}

func (rB *RequestBuilder) SendGet(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToGetRequestInformation(config)
	if err != nil {
		return err
	}

	response, err := rB.Client.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

func (rB *RequestBuilder) SendPost(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToPostRequestInformation(config)
	if err != nil {
		return err
	}

	response, err := rB.Client.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

func (rB *RequestBuilder) SendDelete(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToDeleteRequestInformation(config)
	if err != nil {
		return err
	}

	_, err = rB.Client.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	return err
}

func (rB *RequestBuilder) SendPut(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToPutRequestInformation(config)
	if err != nil {
		return err
	}

	response, err := rB.Client.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}
