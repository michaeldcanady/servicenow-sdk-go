package internal

import (
	"encoding/json"
	"fmt"
)

type RequestBuilder interface {
	ToRequestInformation(method HTTPMethod, config RequestConfiguration) (RequestInformation, error)
}

type requestBuilder struct {
}

func (rB *requestBuilder) ToRequestInformation(method HTTPMethod, config RequestConfiguration) (RequestInformation, error) {
	requestInfo := NewRequestInformation()
	requestInfo.Method = method
	requestInfo.uri.PathParameters = rB.PathParameters
	requestInfo.uri.UrlTemplate = rB.UrlTemplate

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %s", err)
		}
		requestInfo.Content = jsonData
	}

	if params != nil {
		err := requestInfo.AddQueryParameters(params)
		if err != nil {
			return nil, err
		}
	}
	return requestInfo, nil
}
