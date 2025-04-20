package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServiceNowErrorThrower(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				deserializer := mocking.NewMockDeserializer()
				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()

				thrower := NewServiceNowErrorThrower(errorRegistry, deserializer)

				assert.Equal(t, deserializer, thrower.deserializer)
				assert.Equal(t, errorRegistry, thrower.errorRegistry)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServiceNowErrorThrower_resolveErrorFactory(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful 5XX",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				errorMapping := abstractions.ErrorMappings{"5XX": factory}

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
				}

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Nil(t, err)
				assert.NotNil(t, factory)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "Successful 4XX",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(400)

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				errorMapping := abstractions.ErrorMappings{"4XX": factory}

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
				}

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Nil(t, err)
				assert.NotNil(t, factory)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "Successful exact",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(401)

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				errorMapping := abstractions.ErrorMappings{"401": factory}

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
				}

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Nil(t, err)
				assert.NotNil(t, factory)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "Successful any",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(403)

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				errorMapping := abstractions.ErrorMappings{"XXX": factory}

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
				}

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Nil(t, err)
				assert.NotNil(t, factory)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "No factory",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(403)

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				errorMapping := abstractions.ErrorMappings{}

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
				}

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Equal(t, errors.New("no error factory registered"), err)
				assert.Nil(t, factory)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "Registry error",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(403)

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return((abstractions.ErrorMappings)(nil), errors.New("registry error"))

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
				}

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Equal(t, errors.New("registry error"), err)
				assert.Nil(t, factory)
				errorRegistry.AssertExpectations(t)
			},
		},
		{
			name: "nil thrower",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)

				thrower := (*ServiceNowErrorThrower)(nil)

				factory, err := thrower.resolveErrorFactory(typeName, statusCode)

				assert.Equal(t, NewNilPointerError("et is nil"), err)
				assert.Nil(t, factory)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServiceNowErrorThrower_Throw(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)
				contentType := "application/json"
				content := []byte{}

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parsable := mocking.NewMockParsableError()

				errorMapping := abstractions.ErrorMappings{"5XX": factory}

				deserializer := mocking.NewMockDeserializer()
				deserializer.On("Deserialize", contentType, content, mock.AnythingOfType("serialization.ParsableFactory")).Return(parsable, nil)

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
					deserializer:  deserializer,
				}

				err := thrower.Throw(typeName, statusCode, contentType, content)

				assert.Equal(t, parsable, err)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
				parsable.AssertExpectations(t)
				deserializer.AssertExpectations(t)
			},
		},
		{
			name: "Deserializer error",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)
				contentType := "application/json"
				content := []byte{}

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				errorMapping := abstractions.ErrorMappings{"5XX": factory}

				deserializer := mocking.NewMockDeserializer()
				deserializer.On("Deserialize", contentType, content, mock.AnythingOfType("serialization.ParsableFactory")).Return((*mocking.MockParsable)(nil), errors.New("failed to deserialize"))

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
					deserializer:  deserializer,
				}

				err := thrower.Throw(typeName, statusCode, contentType, content)

				assert.Equal(t, errors.New("failed to deserialize"), err)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
				deserializer.AssertExpectations(t)
			},
		},
		{
			name: "Deserializer error",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)
				contentType := "application/json"
				content := []byte{}

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parsable := mocking.NewMockParsable()

				errorMapping := abstractions.ErrorMappings{"5XX": factory}

				deserializer := mocking.NewMockDeserializer()
				deserializer.On("Deserialize", contentType, content, mock.AnythingOfType("serialization.ParsableFactory")).Return(parsable, nil)

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
					deserializer:  deserializer,
				}

				err := thrower.Throw(typeName, statusCode, contentType, content)

				assert.Equal(t, fmt.Errorf("%T is not error", parsable), err)
				errorRegistry.AssertExpectations(t)
				strct.AssertExpectations(t)
				parsable.AssertExpectations(t)
				deserializer.AssertExpectations(t)
			},
		},
		{
			name: "Resolve error",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)
				contentType := "application/json"
				content := []byte{}

				parsable := mocking.NewMockParsableError()

				errorMapping := abstractions.ErrorMappings{}

				deserializer := mocking.NewMockDeserializer()

				errorRegistry := mocking.NewMockDictionary[string, abstractions.ErrorMappings]()
				errorRegistry.On("Get", typeName).Return(errorMapping, nil)

				thrower := &ServiceNowErrorThrower{
					errorRegistry: errorRegistry,
					deserializer:  deserializer,
				}

				err := thrower.Throw(typeName, statusCode, contentType, content)

				assert.Equal(t, errors.New("no error factory registered"), err)
				errorRegistry.AssertExpectations(t)
				parsable.AssertExpectations(t)
				deserializer.AssertExpectations(t)
			},
		},
		{
			name: "nil thrower",
			test: func(t *testing.T) {
				typeName := "type"
				statusCode := int64(500)
				contentType := "application/json"
				content := []byte{}

				thrower := (*ServiceNowErrorThrower)(nil)

				err := thrower.Throw(typeName, statusCode, contentType, content)

				assert.Equal(t, NewNilPointerError("et is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
