package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO: add tests
func TestThrowErrors(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSerializeContent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				const contentType = "application/json"

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return(mocking.NewMockParsable(), nil)

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*mocking.MockParsable](contentType, []byte{}, factory)
				assert.Nil(t, err)
				assert.Equal(t, mocking.NewMockParsable(), parsable)

				parseNode.AssertExpectations(t)
				parseNodeFactory.AssertExpectations(t)
			},
		},
		{
			name: "No content type",
			test: func(t *testing.T) {
				const contentType = ""

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return(mocking.NewMockParsable(), nil)

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*mocking.MockParsable](contentType, []byte{}, factory)
				assert.Equal(t, errors.New("contentType is required"), err)
				assert.Nil(t, parsable)
			},
		},
		{
			name: "Bad content",
			test: func(t *testing.T) {
				const contentType = "application/random"

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return((*mocking.MockParsable)(nil), errors.New("bad content"))

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*mocking.MockParsable](contentType, []byte{}, factory)
				assert.Equal(t, errors.New("bad content"), err)
				assert.Nil(t, parsable)
			},
		},
		{
			name: "differing type",
			test: func(t *testing.T) {
				const contentType = "application/random1"

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return(mocking.NewMockParsable(), nil)

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*newInternal.MainError](contentType, []byte{}, factory)
				assert.Equal(t, errors.New("result is not *internal.MainError"), err)
				assert.Nil(t, parsable)
			},
		},
		//TODO: add test for unregistered content type
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func toPointer[T any](value T) *T {
	return &value
}

// TODO: add tests
func TestGetHTTPHeader(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				header := mocking.NewMockBatchHeader()
				header.On("GetName").Return(toPointer(newInternal.HTTPHeaderContentType.String()), nil)
				header.On("GetValue").Return(toPointer("application/json"), nil)
				headers := []BatchHeaderable{header}
				defaultValue := ""
				value := getHTTPHeader(headers, newInternal.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, "application/json", value)
			},
		},
		{
			name: "Name Error",
			test: func(t *testing.T) {
				header := mocking.NewMockBatchHeader()
				header.On("GetName").Return((*string)(nil), errors.New("no name"))
				header.On("GetValue").Return(toPointer("application/json"), nil)
				headers := []BatchHeaderable{header}
				defaultValue := ""
				value := getHTTPHeader(headers, newInternal.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, defaultValue, value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
