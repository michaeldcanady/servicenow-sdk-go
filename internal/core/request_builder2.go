package core

import (
	"context"
	"encoding/json"
	"fmt"
)

type RequestBuilder2 interface {
	ToRequestInformation(method HttpMethod, config *RequestConfiguration) (RequestInformation, error)
	GetPathParameters() map[string]string
	GetClient() Client2
	GetURLTemplate() string
	Send(ctx context.Context, method HttpMethod, opts ...RequestConfigurationOption) (interface{}, error)
}

type requestBuilder2 struct {
	pathParameters map[string]string
	client         Client2
	urlTemplate    string
}

func NewRequestBuilder2(client Client2, urlTemplate string, pathParameters map[string]string) RequestBuilder2 {
	return &requestBuilder2{
		client:         client,
		urlTemplate:    urlTemplate,
		pathParameters: pathParameters,
	}
}

func (rB *requestBuilder2) GetPathParameters() map[string]string {
	copy := make(map[string]string)
	for key, value := range rB.pathParameters {
		copy[key] = value
	}
	return copy
}

func (rB *requestBuilder2) GetClient() Client2 {
	return rB.client
}

func (rB *requestBuilder2) GetURLTemplate() string {
	return rB.urlTemplate
}

func (rB *requestBuilder2) ToRequestInformation(method HttpMethod, config *RequestConfiguration) (RequestInformation, error) {
	requestInfo := NewRequestInformation(
		WithMethod(method),
		WithPathParams(rB.pathParameters),
		WithURITemplate(rB.urlTemplate),
	)

	if config == nil {
		return requestInfo, nil
	}

	if err := requestInfo.AddQueryParameters(config.QueryParameters); err != nil {
		return nil, fmt.Errorf("failed to add query parameters: %w", err)
	}

	if config.Data != nil {
		data, err := json.Marshal(config.Data)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %w", err)
		}
		if len(data) > 0 {
			requestInfo.SetContent(data, jsonContentType)
		}
	}

	return requestInfo, nil
}

func (rB *requestBuilder2) Send(ctx context.Context, method HttpMethod, opts ...RequestConfigurationOption) (interface{}, error) {
	config := ApplyOptions(opts...)

	requestInfo, err := rB.ToRequestInformation(method, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create request information: %w", err)
	}

	response, err := rB.client.SendWithContext(ctx, requestInfo, config.ErrorMapping.(ErrorMapping))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if method == DELETE || config.Response == nil {
		return nil, nil
	}

	if err = ParseResponse(response, &config.Response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return config.Response, nil
}
