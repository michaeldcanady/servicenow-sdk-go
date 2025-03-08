package mocking

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockRequestAdapter struct {
	mock.Mock
}

func NewMockRequestAdapter() *MockRequestAdapter {
	return &MockRequestAdapter{
		mock.Mock{},
	}
}

// Send executes the HTTP request specified by the given RequestInformation and returns the deserialized response model.
func (rA *MockRequestAdapter) Send(context context.Context, requestInfo *abstractions.RequestInformation, constructor serialization.ParsableFactory, errorMappings abstractions.ErrorMappings) (serialization.Parsable, error) {
	args := rA.Called(context, requestInfo, constructor, errorMappings)
	return args.Get(0).(serialization.Parsable), args.Error(1)
}

// SendEnum executes the HTTP request specified by the given RequestInformation and returns the deserialized response model.
func (rA *MockRequestAdapter) SendEnum(context context.Context, requestInfo *abstractions.RequestInformation, parser serialization.EnumFactory, errorMappings abstractions.ErrorMappings) (any, error) {
	args := rA.Called(context, requestInfo, parser, errorMappings)
	return args.Get(0), args.Error(1)
}

// SendCollection executes the HTTP request specified by the given RequestInformation and returns the deserialized response model collection.
func (rA *MockRequestAdapter) SendCollection(context context.Context, requestInfo *abstractions.RequestInformation, constructor serialization.ParsableFactory, errorMappings abstractions.ErrorMappings) ([]serialization.Parsable, error) {
	args := rA.Called(context, requestInfo, constructor, errorMappings)
	return args.Get(0).([]serialization.Parsable), args.Error(1)
}

// SendEnumCollection executes the HTTP request specified by the given RequestInformation and returns the deserialized response model collection.
func (rA *MockRequestAdapter) SendEnumCollection(context context.Context, requestInfo *abstractions.RequestInformation, parser serialization.EnumFactory, errorMappings abstractions.ErrorMappings) ([]any, error) {
	args := rA.Called(context, requestInfo, parser, errorMappings)
	return args.Get(0).([]any), args.Error(1)
}

// SendPrimitive executes the HTTP request specified by the given RequestInformation and returns the deserialized primitive response model.
func (rA *MockRequestAdapter) SendPrimitive(context context.Context, requestInfo *abstractions.RequestInformation, typeName string, errorMappings abstractions.ErrorMappings) (any, error) {
	args := rA.Called(context, requestInfo, typeName, errorMappings)
	return args.Get(0), args.Error(1)
}

// SendPrimitiveCollection executes the HTTP request specified by the given RequestInformation and returns the deserialized primitive response model collection.
func (rA *MockRequestAdapter) SendPrimitiveCollection(context context.Context, requestInfo *abstractions.RequestInformation, typeName string, errorMappings abstractions.ErrorMappings) ([]any, error) {
	args := rA.Called(context, requestInfo, typeName, errorMappings)
	return args.Get(0).([]any), args.Error(1)
}

// SendNoContent executes the HTTP request specified by the given RequestInformation with no return content.
func (rA *MockRequestAdapter) SendNoContent(context context.Context, requestInfo *abstractions.RequestInformation, errorMappings abstractions.ErrorMappings) error {
	args := rA.Called(context, requestInfo, errorMappings)
	return args.Error(1)
}

// GetSerializationWriterFactory returns the serialization writer factory currently in use for the request adapter service.
func (rA *MockRequestAdapter) GetSerializationWriterFactory() serialization.SerializationWriterFactory {
	args := rA.Called()
	return args.Get(0).(serialization.SerializationWriterFactory)
}

// EnableBackingStore enables the backing store proxies for the SerializationWriters and ParseNodes in use.
func (rA *MockRequestAdapter) EnableBackingStore(factory store.BackingStoreFactory) {
	_ = rA.Called(factory)
}

// SetBaseUrl sets the base url for every request.
func (rA *MockRequestAdapter) SetBaseUrl(baseUrl string) { //nolint:stylecheck
	_ = rA.Called(baseUrl)
}

// GetBaseUrl gets the base url for every request.
func (rA *MockRequestAdapter) GetBaseUrl() string { //nolint:stylecheck
	args := rA.Called()
	return args.String(0)
}

// ConvertToNativeRequest converts the given RequestInformation into a native HTTP request.
func (rA *MockRequestAdapter) ConvertToNativeRequest(context context.Context, requestInfo *abstractions.RequestInformation) (any, error) {
	args := rA.Called(context, requestInfo)
	return args.Get(0), args.Error(1)
}
