package mocking

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// MockResponseHandler is a mock implementation of the abstractions.ResponseHandler interface.
type MockResponseHandler struct {
	Response interface{}
	Error    error
}

func (m *MockResponseHandler) HandleResponse(response interface{}, errorMapping abstractions.ErrorMappings) (interface{}, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response, nil
}
