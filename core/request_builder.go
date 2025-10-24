package core

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
)

func NewRequestBuilder2(client Client2, urlTemplate string, pathParameters map[string]string) *RequestBuilder {
	return &RequestBuilder{
		Client:         client,
		Client2:        client,
		UrlTemplate:    urlTemplate,
		PathParameters: pathParameters,
	}
}

func (rB *RequestBuilder) ToPutRequestInformation2(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(PUT, config)
}

func (rB *RequestBuilder) ToPostRequestInformation3(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(POST, config)
}

func (rB *RequestBuilder) ToDeleteRequestInformation2(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(DELETE, config)
}

func (rB *RequestBuilder) ToGetRequestInformation2(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(GET, config)
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

func (rB *RequestBuilder) ToRequestInformation3(method HttpMethod, config *RequestConfiguration) (*RequestInformation, error) {
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

func (rB *RequestBuilder) SendGet3(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToGetRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := rB.Client2.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

func (rB *RequestBuilder) SendPost4(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToPostRequestInformation3(config)
	if err != nil {
		return err
	}

	response, err := rB.Client2.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

func (rB *RequestBuilder) SendDelete3(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToDeleteRequestInformation2(config)
	if err != nil {
		return err
	}

	_, err = rB.Client2.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	return err
}

func (rB *RequestBuilder) SendPut3(ctx context.Context, config *RequestConfiguration) error {
	requestInfo, err := rB.ToPutRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := rB.Client2.SendWithContext(ctx, requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}
