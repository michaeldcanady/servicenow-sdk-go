package batchapi

import (
	"errors"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// deserializerRegistrations makes each registered content type unique. Kiota's
// RegisterDefaultDeserializer keeps the FIRST factory registered for a content type and
// silently ignores later ones, and the registry is a process-wide singleton. Reusing a fixed
// content type therefore breaks these tests on a repeat run (go test -count=2), because the
// stale factory from the previous run stays registered and the fresh mock is never called.
var deserializerRegistrations atomic.Uint64

// uniqueContentType returns a content type no previous registration has claimed.
func uniqueContentType(name string) string {
	return fmt.Sprintf("application/json-%s-%d", name, deserializerRegistrations.Add(1))
}

func TestThrowErrors(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(m *MockServicedRequest)
		expectedErr bool
	}{
		{
			name: "No Error",
			setup: func(m *MockServicedRequest) {
				m.On("GetStatusCode").Return(internal.ToPointer(int64(200)), nil)
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
				m.On("GetStatusCode").Return(internal.ToPointer(int64(400)), nil)
				m.On("GetErrorMessage").Return(internal.ToPointer(`{"error":"bad"}`), nil)
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
				contentType := uniqueContentType("shared-test-successful")

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return(mocking.NewMockParsable(), nil)

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*mocking.MockParsable](contentType, []byte{}, factory)
				require.NoError(t, err)
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
				contentType := uniqueContentType("shared-test-bad-content")

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
				contentType := uniqueContentType("shared-test-differing-type")

				strct := mocking.NewMockParsableFactory()
				factory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetObjectValue", mock.AnythingOfType("serialization.ParsableFactory")).Return(mocking.NewMockParsable(), nil)

				parseNodeFactory := mocking.NewMockParseNodeFactory()
				parseNodeFactory.On("GetValidContentType").Return(contentType, nil)
				parseNodeFactory.On("GetRootParseNode", contentType, []byte{}).Return(parseNode, nil)

				abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return parseNodeFactory })

				parsable, err := serializeContent[*core.MainError](contentType, []byte{}, factory)
				assert.Equal(t, errors.New("result is not *core.MainError"), err)
				assert.Nil(t, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestGetHTTPHeader(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				header := mocking.NewMockBatchHeader()
				header.On("GetName").Return(internal.ToPointer(internalhttp.HTTPHeaderContentType.String()), nil)
				header.On("GetValue").Return(internal.ToPointer("application/json"), nil)
				headers := []RestRequestHeader{header}
				defaultValue := ""
				value := getHTTPHeader(headers, internalhttp.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, "application/json", value)
			},
		},
		{
			name: "Not Found",
			test: func(t *testing.T) {
				headers := []RestRequestHeader{}
				defaultValue := "default"
				value := getHTTPHeader(headers, internalhttp.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, defaultValue, value)
			},
		},
		{
			name: "Name Error",
			test: func(t *testing.T) {
				header := mocking.NewMockBatchHeader()
				header.On("GetName").Return((*string)(nil), errors.New("no name"))
				header.On("GetValue").Return(internal.ToPointer("application/json"), nil)
				headers := []RestRequestHeader{header}
				defaultValue := ""
				value := getHTTPHeader(headers, internalhttp.HTTPHeaderContentType, defaultValue)
				assert.Equal(t, defaultValue, value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
