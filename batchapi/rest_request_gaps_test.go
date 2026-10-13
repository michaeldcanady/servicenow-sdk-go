// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package batchapi

import (
	"encoding/base64"
	"net/url"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateRestRequestHeaderFromHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  func() *abstractions.RequestHeaders
		expected map[string]string
	}{
		{
			name:     "nil headers yield an empty collection",
			headers:  func() *abstractions.RequestHeaders { return nil },
			expected: map[string]string{},
		},
		{
			name:     "empty headers yield an empty collection",
			headers:  abstractions.NewRequestHeaders,
			expected: map[string]string{},
		},
		{
			// Kiota's RequestHeaders normalizes keys to lower case, so the batch sub-request
			// headers come out lower-cased too.
			name: "single-valued headers are carried across",
			headers: func() *abstractions.RequestHeaders {
				headers := abstractions.NewRequestHeaders()
				headers.Add("Accept", "application/json")

				return headers
			},
			expected: map[string]string{"accept": "application/json"},
		},
		{
			name: "multiple keys each produce a header",
			headers: func() *abstractions.RequestHeaders {
				headers := abstractions.NewRequestHeaders()
				headers.Add("Accept", "application/json")
				headers.Add("Content-Type", "application/json")

				return headers
			},
			expected: map[string]string{
				"accept":       "application/json",
				"content-type": "application/json",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			batchHeaders, err := createRestRequestHeaderFromHeaders(test.headers())

			require.NoError(t, err)
			require.NotNil(t, batchHeaders)
			require.Len(t, batchHeaders, len(test.expected))

			actual := make(map[string]string, len(batchHeaders))
			for _, header := range batchHeaders {
				name, err := header.GetName()
				require.NoError(t, err)
				require.NotNil(t, name)

				value, err := header.GetValue()
				require.NoError(t, err)
				require.NotNil(t, value)

				actual[*name] = *value
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

// TestCreateRestRequestHeaderFromHeaders_MultiValued covers the flattening of a multi-valued
// header into the single string a batch sub-request header carries. Kiota stores header values
// in a set, so the joined order is not deterministic and the assertion allows either order.
func TestCreateRestRequestHeaderFromHeaders_MultiValued(t *testing.T) {
	headers := abstractions.NewRequestHeaders()
	headers.Add("Accept", "application/json", "text/plain")

	batchHeaders, err := createRestRequestHeaderFromHeaders(headers)

	require.NoError(t, err)
	require.Len(t, batchHeaders, 1)

	value, err := batchHeaders[0].GetValue()
	require.NoError(t, err)
	require.NotNil(t, value)
	assert.ElementsMatch(t, []string{"application/json", "text/plain"}, strings.Split(*value, ", "))
}

// TestRestRequest_SetURL_Gaps covers SetURL's two rejections and its relativization of an
// absolute URL.
func TestRestRequest_SetURL_Gaps(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		expected   string
		wantErrMsg string
	}{
		{
			name:     "absolute URL is made relative",
			url:      "https://example.com/api/now/table/incident",
			expected: "/api/now/table/incident",
		},
		{
			name:     "already-relative URL is kept",
			url:      "/api/now/table/incident",
			expected: "/api/now/table/incident",
		},
		{
			name:       "path outside /api is rejected",
			url:        "/some/other/path",
			wantErrMsg: `path doesn't begin with "/api"`,
		},
		{
			name:       "unparsable URL is rejected",
			url:        "https://example.com/api\n",
			wantErrMsg: "invalid control character",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model := NewRestRequest()

			err := model.SetURL(internal.ToPointer(test.url))

			if test.wantErrMsg != "" {
				require.ErrorContains(t, err, test.wantErrMsg)

				return
			}

			require.NoError(t, err)

			actual, err := model.GetURL()
			require.NoError(t, err)
			require.NotNil(t, actual)
			assert.Equal(t, test.expected, *actual)
		})
	}
}

// TestCreateRestRequestFromRequestInformation_Gaps covers the parts of the converter the
// existing happy-path test misses: headers actually being carried over, the body being
// preserved, an ID being minted, and the failure when the URI cannot be resolved.
func TestCreateRestRequestFromRequestInformation_Gaps(t *testing.T) {
	t.Run("carries body, method, headers and mints an ID", func(t *testing.T) {
		requestInfo := abstractions.NewRequestInformation()
		requestInfo.Method = abstractions.POST
		requestInfo.SetUri(url.URL{Scheme: "https", Host: "example.com", Path: "/api/now/table/incident"})
		requestInfo.Headers.Add("Accept", "application/json")
		requestInfo.Content = []byte(`{"short_description":"test"}`)

		request, err := CreateRestRequestFromRequestInformation(requestInfo, true)

		require.NoError(t, err)
		require.NotNil(t, request)

		body, err := request.GetBody()
		require.NoError(t, err)
		assert.JSONEq(t, `{"short_description":"test"}`, string(body))

		method, err := request.GetMethod()
		require.NoError(t, err)
		require.NotNil(t, method)
		assert.Equal(t, abstractions.POST, *method)

		// SetURL strips the scheme and host: a batch sub-request carries a relative URL,
		// since the batch envelope itself already targets the instance.
		requestURL, err := request.GetURL()
		require.NoError(t, err)
		require.NotNil(t, requestURL)
		assert.Equal(t, "/api/now/table/incident", *requestURL)

		exclude, err := request.GetExcludeResponseHeaders()
		require.NoError(t, err)
		require.NotNil(t, exclude)
		assert.True(t, *exclude)

		headers, err := request.GetHeaders()
		require.NoError(t, err)
		require.Len(t, headers, 1)

		id, err := request.GetID()
		require.NoError(t, err)
		require.NotNil(t, id)
		assert.NotEmpty(t, *id, "a sub-request must carry an ID so the response can be correlated")
	})

	t.Run("unresolvable URI fails", func(t *testing.T) {
		// A RequestInformation with neither a URI nor a resolvable template cannot produce a URL.
		request, err := CreateRestRequestFromRequestInformation(abstractions.NewRequestInformation(), false)

		require.Error(t, err)
		assert.Nil(t, request)
	})
}

// TestRestRequest_SerializeGaps covers the mutating serializers, which the existing Serialize
// test does not reach: base64 body encoding, the nil-method rejection, and the ID being minted
// during serialization when the model has none.
func TestRestRequest_SerializeGaps(t *testing.T) {
	t.Run("body is base64 encoded", func(t *testing.T) {
		model := NewRestRequest()
		require.NoError(t, model.SetBody([]byte("hello")))
		require.NoError(t, model.SetMethod(internal.ToPointer(abstractions.GET)))

		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

		require.NoError(t, model.Serialize(writer))

		encoded := base64.StdEncoding.EncodeToString([]byte("hello"))
		writer.AssertCalled(t, "WriteStringValue", bodyKey, internal.ToPointer(encoded))
	})

	t.Run("nil method is rejected", func(t *testing.T) {
		model := NewRestRequest()
		require.NoError(t, model.SetBody([]byte("hello")))

		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

		err := model.Serialize(writer)

		require.ErrorContains(t, err, "method can't be nil")
	})

	t.Run("an empty ID is replaced with a generated one", func(t *testing.T) {
		model := NewRestRequest()
		require.NoError(t, model.SetID(internal.ToPointer("")))
		require.NoError(t, model.SetMethod(internal.ToPointer(abstractions.GET)))

		var writtenID *string
		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteStringValue", idKey, mock.Anything).Run(func(args mock.Arguments) {
			writtenID, _ = args.Get(1).(*string)
		}).Return(nil)
		writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

		require.NoError(t, model.Serialize(writer))

		require.NotNil(t, writtenID)
		assert.NotEmpty(t, *writtenID, "serialization must mint an ID when the model has none")
	})
}

// TestRestRequest_FieldDeserializers exercises each deserializer rather than merely asserting
// the map is non-nil, which is all the existing test does.
func TestRestRequest_FieldDeserializers(t *testing.T) {
	t.Run("body is base64 decoded", func(t *testing.T) {
		model := NewRestRequest()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return(internal.ToPointer(base64.StdEncoding.EncodeToString([]byte("hello"))), nil)

		require.NoError(t, model.GetFieldDeserializers()[bodyKey](node))

		body, err := model.GetBody()
		require.NoError(t, err)
		assert.Equal(t, []byte("hello"), body)
	})

	t.Run("a nil body deserializes to no body", func(t *testing.T) {
		model := NewRestRequest()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return((*string)(nil), nil)

		require.NoError(t, model.GetFieldDeserializers()[bodyKey](node))

		body, err := model.GetBody()
		require.NoError(t, err)
		assert.Nil(t, body)
	})

	t.Run("a body that is not base64 is rejected", func(t *testing.T) {
		model := NewRestRequest()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return(internal.ToPointer("not-base64!!!"), nil)

		require.Error(t, model.GetFieldDeserializers()[bodyKey](node))
	})

	t.Run("method is parsed", func(t *testing.T) {
		model := NewRestRequest()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return(internal.ToPointer("POST"), nil)

		require.NoError(t, model.GetFieldDeserializers()[methodKey](node))

		method, err := model.GetMethod()
		require.NoError(t, err)
		require.NotNil(t, method)
		assert.Equal(t, abstractions.POST, *method)
	})

	t.Run("a nil method deserializes to no method", func(t *testing.T) {
		model := NewRestRequest()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return((*string)(nil), nil)

		require.NoError(t, model.GetFieldDeserializers()[methodKey](node))
	})

	t.Run("an unknown method is rejected", func(t *testing.T) {
		model := NewRestRequest()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return(internal.ToPointer("BREW"), nil)

		require.Error(t, model.GetFieldDeserializers()[methodKey](node))
	})
}
