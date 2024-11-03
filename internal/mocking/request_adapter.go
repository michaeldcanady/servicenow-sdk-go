package mocking

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

var _ abstractions.RequestAdapter = (*RequestAdapter)(nil)

type RequestAdapter struct {
	mock.Mock
}

// Send executes the HTTP request specified by the given RequestInformation and returns the deserialized response model.
func (m *RequestAdapter) Send(context context.Context, requestInfo *abstractions.RequestInformation, constructor serialization.ParsableFactory, errorMappings abstractions.ErrorMappings) (serialization.Parsable, error) {
	args := m.Called(context, requestInfo, constructor, errorMappings)
	return args.Get(0).(serialization.Parsable), args.Error(1)
}

// SendEnum executes the HTTP request specified by the given RequestInformation and returns the deserialized response model.
func (m *RequestAdapter) SendEnum(context context.Context, requestInfo *abstractions.RequestInformation, parser serialization.EnumFactory, errorMappings abstractions.ErrorMappings) (any, error) {
	args := m.Called(context, requestInfo, parser, errorMappings)
	return args.Get(0), args.Error(1)
}

// SendCollection executes the HTTP request specified by the given RequestInformation and returns the deserialized response model collection.
func (m *RequestAdapter) SendCollection(context context.Context, requestInfo *abstractions.RequestInformation, constructor serialization.ParsableFactory, errorMappings abstractions.ErrorMappings) ([]serialization.Parsable, error) {
	args := m.Called(context, requestInfo, constructor, errorMappings)
	return args.Get(0).([]serialization.Parsable), args.Error(1)
}

// SendEnumCollection executes the HTTP request specified by the given RequestInformation and returns the deserialized response model collection.
func (m *RequestAdapter) SendEnumCollection(context context.Context, requestInfo *abstractions.RequestInformation, parser serialization.EnumFactory, errorMappings abstractions.ErrorMappings) ([]any, error) {
	args := m.Called(context, requestInfo, parser, errorMappings)
	return args.Get(0).([]any), args.Error(1)
}

// SendPrimitive executes the HTTP request specified by the given RequestInformation and returns the deserialized primitive response model.
func (m *RequestAdapter) SendPrimitive(context context.Context, requestInfo *abstractions.RequestInformation, typeName string, errorMappings abstractions.ErrorMappings) (any, error) {
	args := m.Called(context, requestInfo, typeName, errorMappings)
	return args.Get(0), args.Error(1)
}

// SendPrimitiveCollection executes the HTTP request specified by the given RequestInformation and returns the deserialized primitive response model collection.
func (m *RequestAdapter) SendPrimitiveCollection(context context.Context, requestInfo *abstractions.RequestInformation, typeName string, errorMappings abstractions.ErrorMappings) ([]any, error) {
	args := m.Called(context, requestInfo, typeName, errorMappings)
	return args.Get(0).([]any), args.Error(1)
}

// SendNoContent executes the HTTP request specified by the given RequestInformation with no return content.
func (m *RequestAdapter) SendNoContent(context context.Context, requestInfo *abstractions.RequestInformation, errorMappings abstractions.ErrorMappings) error {
	args := m.Called(context, requestInfo, errorMappings)
	return args.Error(1)
}

// GetSerializationWriterFactory returns the serialization writer factory currently in use for the request adapter service.
func (m *RequestAdapter) GetSerializationWriterFactory() serialization.SerializationWriterFactory {
	args := m.Called()
	return args.Get(0).(serialization.SerializationWriterFactory)
}

// EnableBackingStore enables the backing store proxies for the SerializationWriters and ParseNodes in use.
func (m *RequestAdapter) EnableBackingStore(factory store.BackingStoreFactory) {
	_ = m.Called(factory)
}

// SetBaseUrl sets the base url for every request.
func (m *RequestAdapter) SetBaseUrl(baseUrl string) {
	_ = m.Called(baseUrl)
}

// GetBaseUrl gets the base url for every request.
func (m *RequestAdapter) GetBaseUrl() string {
	args := m.Called()
	return args.String(0)
}

// ConvertToNativeRequest converts the given RequestInformation into a native HTTP request.
func (m *RequestAdapter) ConvertToNativeRequest(context context.Context, requestInfo *abstractions.RequestInformation) (any, error) {
	args := m.Called(context, requestInfo)
	return args.Get(0), args.Error(1)
}
