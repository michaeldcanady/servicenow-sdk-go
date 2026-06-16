package mocking

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// MockRequestAdapter is a mock implementation of the abstractions.RequestAdapter interface.
type MockRequestAdapter struct {
	LastRequest *abstractions.RequestInformation
	Response    interface{}
	Error       error
}

func (m *MockRequestAdapter) Send(ctx context.Context, requestInfo *abstractions.RequestInformation, constructor serialization.ParsableFactory, errorMapping abstractions.ErrorMappings) (serialization.Parsable, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response.(serialization.Parsable), nil
}

func (m *MockRequestAdapter) SendCollection(ctx context.Context, requestInfo *abstractions.RequestInformation, constructor serialization.ParsableFactory, errorMapping abstractions.ErrorMappings) ([]serialization.Parsable, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response.([]serialization.Parsable), nil
}

func (m *MockRequestAdapter) SendPrimitive(ctx context.Context, requestInfo *abstractions.RequestInformation, typeName string, errorMapping abstractions.ErrorMappings) (interface{}, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response, nil
}

func (m *MockRequestAdapter) SendPrimitiveCollection(ctx context.Context, requestInfo *abstractions.RequestInformation, typeName string, errorMapping abstractions.ErrorMappings) ([]interface{}, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response.([]interface{}), nil
}

func (m *MockRequestAdapter) SendNoResponseContent(ctx context.Context, requestInfo *abstractions.RequestInformation, errorMapping abstractions.ErrorMappings) error {
	m.LastRequest = requestInfo
	return m.Error
}

func (m *MockRequestAdapter) SendNoContent(ctx context.Context, requestInfo *abstractions.RequestInformation, errorMapping abstractions.ErrorMappings) error {
	m.LastRequest = requestInfo
	return m.Error
}

func (m *MockRequestAdapter) SendEnum(ctx context.Context, requestInfo *abstractions.RequestInformation, parser serialization.EnumFactory, errorMapping abstractions.ErrorMappings) (any, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response, nil
}

func (m *MockRequestAdapter) SendEnumCollection(ctx context.Context, requestInfo *abstractions.RequestInformation, parser serialization.EnumFactory, errorMapping abstractions.ErrorMappings) ([]any, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response.([]any), nil
}

func (m *MockRequestAdapter) ConvertToNativeRequest(ctx context.Context, requestInfo *abstractions.RequestInformation) (any, error) {
	m.LastRequest = requestInfo
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response, nil
}

func (m *MockRequestAdapter) GetSerializationWriterFactory() serialization.SerializationWriterFactory {
	return nil
}

func (m *MockRequestAdapter) GetParseNodeFactory() serialization.ParseNodeFactory {
	return nil
}

func (m *MockRequestAdapter) GetSerializationWriterFactoryCollection() []serialization.SerializationWriterFactory {
	return nil
}

func (m *MockRequestAdapter) GetParseNodeFactoryCollection() []serialization.ParseNodeFactory {
	return nil
}

func (m *MockRequestAdapter) SetBaseUrl(baseUrl string) {
}

func (m *MockRequestAdapter) GetBaseUrl() string {
	return ""
}

func (m *MockRequestAdapter) EnableBackingStore(factory store.BackingStoreFactory) {
}
