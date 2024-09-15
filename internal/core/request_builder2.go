package core

import (
	"context"
	"encoding/json"
	"fmt"
)

type RequestBuilder2 interface {
	Sendable
	GetPathParameters() map[string]string
	GetClient() ClientSendable
	GetURLTemplate() string
}

type Sendable interface {
	Send(ctx context.Context, method HttpMethod, config RequestConfiguration) (interface{}, error)
}

type requestBuilder2 struct {
	pathParameters map[string]string
	client         ClientSendable
	urlTemplate    string
}

func NewRequestBuilder2(client ClientSendable, urlTemplate string, pathParameters map[string]string) RequestBuilder2 {
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

func (rB *requestBuilder2) GetClient() ClientSendable {
	return rB.client
}

func (rB *requestBuilder2) GetURLTemplate() string {
	return rB.urlTemplate
}

func (rB *requestBuilder2) toRequestInformation(method HttpMethod, config RequestConfiguration) (RequestInformation, error) {
	requestInfo := NewRequestInformation(
		WithMethod(method),
		WithPathParams(rB.pathParameters),
		WithURITemplate(rB.urlTemplate),
	)

	if config == nil {
		return requestInfo, nil
	}

	if config, ok := config.(SupportsQueryParams[any]); ok && config.GetQueryParams() != nil {
		if err := requestInfo.AddQueryParameters(config.GetQueryParams()); err != nil {
			return nil, fmt.Errorf("failed to add query parameters: %w", err)
		}
	}

	if config, ok := config.(SupportsData[any]); ok && config.GetData() != nil {
		data, err := json.Marshal(config.GetData())
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %w", err)
		}
		if len(data) > 0 {
			requestInfo.SetContent(data, jsonContentType)
		}
	}

	return requestInfo, nil
}

func (rB *requestBuilder2) Send(ctx context.Context, method HttpMethod, config RequestConfiguration) (interface{}, error) {
	requestInfo, err := rB.toRequestInformation(method, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create request information: %w", err)
	}

	response, err := rB.client.SendWithContext(ctx, requestInfo, config.GetErrorMapping())
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if method == MethodDelete || response == nil {
		return nil, nil
	}

	rawResp := config.GetResponse()

	if err = ParseResponse(response, &rawResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return rawResp, nil
}

//func (rB *requestBuilder2) Send(ctx context.Context, method HttpMethod, opts ...RequestConfigurationOption) (interface{}, error) {
//	config := ApplyOptions(opts...)
//
//	requestInfo, err := rB.toRequestInformation(method, config)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create request information: %w", err)
//	}
//
//	response, err := rB.client.SendWithContext(ctx, requestInfo, config.ErrorMapping)
//	if err != nil {
//		return nil, fmt.Errorf("request failed: %w", err)
//	}
//
//	if method == DELETE || response == nil {
//		return nil, nil
//	}
//
//	if err = ParseResponse(response, &config.Response); err != nil {
//		return nil, fmt.Errorf("failed to parse response: %w", err)
//	}
//
//	return config.Response, nil
//}
