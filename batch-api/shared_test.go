package batchapi

import (
	"errors"
	"testing"

	internalErrors "github.com/michaeldcanady/servicenow-sdk-go/internal/errors"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestThrowErrors(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(m *MockServicedRequest)
		expectedErr bool
	}{
		{
			name: "No Error",
			setup: func(m *MockServicedRequest) {
				m.On("GetStatusCode").Return(utils.ToPointer(int64(200)), nil)
			},
			expectedErr: false,
		},
		{
			name: "Status Error",
			setup: func(m *MockServicedRequest) {
				m.On("GetStatusCode").Return((*int64)(nil), errors.New("status error"))
			},
			expectedErr: true,
		},
		{
			name: "Mapped Error",
			setup: func(m *MockServicedRequest) {
				m.On("GetStatusCode").Return(utils.ToPointer(int64(400)), nil)
				m.On("GetErrorMessage").Return(utils.ToPointer(`{"error":"bad"}`), nil)
				m.On("GetHeaders").Return([]RestRequestHeader{}, nil)
			},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewMockServicedRequest()
			test.setup(m)

			err := throwErrors(m, "test")
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
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
				const contentType = "application/json-shared-test-successful"

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

				parsable, err := serializeContent[*mocking.MockParsable](contentType, []byte{}, factory)
				assert.Equal(t, errors.New("contentType is required"), err)
				assert.Nil(t, parsable)
			},
		},
		{
			name: "Bad content",
			test: func(t *testing.T) {
				const contentType = "application/json-shared-test-bad-content"

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
				const contentType = "application/json-shared-test-differing-type"

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return(mocking.NewMockParsable(), nil)

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*internalErrors.MainError](contentType, []byte{}, factory)
				assert.Equal(t, errors.New("result is not *internal.MainError"), err)
				assert.Nil(t, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
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
				header.On("GetName").Return(utils.ToPointer(internalHttp.HTTPHeaderContentType.String()), nil)
				header.On("GetValue").Return(utils.ToPointer("application/json"), nil)
				headers := []RestRequestHeader{header}
				defaultValue := ""
				value := getHTTPHeader(headers, internalHttp.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, "application/json", value)
			},
		},
		{
			name: "Name Error",
			test: func(t *testing.T) {
				header := mocking.NewMockBatchHeader()
				header.On("GetName").Return((*string)(nil), errors.New("no name"))
				header.On("GetValue").Return(utils.ToPointer("application/json"), nil)
				headers := []RestRequestHeader{header}
				defaultValue := ""
				value := getHTTPHeader(headers, internalHttp.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, defaultValue, value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
