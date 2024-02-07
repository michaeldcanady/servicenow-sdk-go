package core

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestInformation(t *testing.T) {
	requestInfo := NewRequestInformation()

	expectedHeaders := make(http.Header)
	expectedOptions := make(map[string]RequestOption)
	expectedURI := NewURLInformation()

	assert.Equal(t, expectedHeaders, requestInfo.Headers)
	assert.Equal(t, expectedOptions, requestInfo.options)
	assert.Equal(t, expectedURI, requestInfo.uri)
}

func TestNewRequestInformationSetStreamContenr(t *testing.T) {
	content := []byte("Testing Test")

	requestInfo := NewRequestInformation()
	requestInfo.SetStreamContent(content)

	assert.Equal(t, content, requestInfo.Content)
	assert.Equal(t, binaryContentType, requestInfo.Headers.Get(contentTypeHeader))
}

func TestNewRequestInformationAddQueryParameters(t *testing.T) {
	source := struct {
		Var1 string `url:"var_1"`
	}{
		Var1: "Val1",
	}

	expected := map[string]string{"var_1": "Val1"}

	requestInfo := NewRequestInformation()
	err := requestInfo.AddQueryParameters(source)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, requestInfo.uri.QueryParameters)
}

func TestNewRequestInformationSetUri(t *testing.T) {
	url, err := url.Parse("https://www.example.com")
	if err != nil {
		t.Error(err)
	}

	expected := map[string]string{"request-raw-url": "https://www.example.com"}

	requestInfo := NewRequestInformation()
	requestInfo.SetUri(url)

	assert.Equal(t, expected, requestInfo.uri.PathParameters)
}

func TestNewRequestInformationGetContentReader(t *testing.T) {
	requestInfo := NewRequestInformation()
	reader := requestInfo.getContentReader()
	assert.IsType(t, &bytes.Reader{}, reader)
}

func TestNewRequestInformationUrl(t *testing.T) {
	url, err := url.Parse("https://www.example.com")
	if err != nil {
		t.Error(err)
	}

	requestInfo := NewRequestInformation()
	requestInfo.SetUri(url)
	uri, err := requestInfo.Url()
	assert.NoError(t, err)

	assert.Equal(t, "https://www.example.com", uri)
}

func TestNewRequestInformation_ToRequest(t *testing.T) {
	url, err := url.Parse("https://www.example.com")
	if err != nil {
		t.Error(err)
	}

	requestInfo := NewRequestInformation()
	requestInfo.SetUri(url)
	requestInfo.Method = GET

	request, err := requestInfo.ToRequest()
	assert.NoError(t, err)

	expected, err := http.NewRequest("GET", "https://www.example.com", http.NoBody) // Use nil directly here
	assert.NoError(t, err)

	assert.Equal(t, expected.Method, request.Method)
	assert.Equal(t, expected.URL, request.URL)
	assert.Equal(t, expected.Body, request.Body)
}

func TestNewRequestInformationToRequestWithContext(t *testing.T) {
	url, err := url.Parse("https://www.example.com")
	if err != nil {
		t.Error(err)
	}

	requestInfo := NewRequestInformation()
	requestInfo.SetUri(url)
	requestInfo.Method = GET
	request, err := requestInfo.ToRequestWithContext(context.TODO())
	assert.NoError(t, err)

	expected, err := http.NewRequestWithContext(context.TODO(), "GET", "https://www.example.com", bytes.NewReader([]byte(nil)))
	assert.NoError(t, err)

	assert.Equal(t, expected.Method, request.Method)
	assert.Equal(t, expected.URL, request.URL)
	assert.Equal(t, expected.Body, request.Body)
	assert.Equal(t, expected.Context(), request.Context())
}

type test[T any] struct {
	title string
	// setup to make needed modifications for a specific test
	setup func()
	// cleanup to undo changes do to reusable items
	cleanup     func()
	input       interface{}
	expected    T
	shouldErr   bool
	expectedErr error
}

func TestNewRequestInformationAddHeaders(t *testing.T) {
	tests := []test[http.Header]{
		{
			title: "Test Struct Headers",
			input: struct {
				Header1 string `header:"header-1"`
				Header2 string `header:"header-2"`
				Header3 string `header:"header-3"`
			}{
				Header1: "value1",
				Header2: "value2",
				Header3: "value3",
			},
			expected: http.Header{
				"Header-1": []string{"value1"},
				"Header-2": []string{"value2"},
				"Header-3": []string{"value3"},
			},
			shouldErr: false,
		},
		{
			title: "Test http.Header Headers",
			input: http.Header{
				"Header-1": []string{"value1"},
				"Header-2": []string{"value2"},
				"Header-3": []string{"value3"},
			},
			expected: http.Header{
				"Header-1": []string{"value1"},
				"Header-2": []string{"value2"},
				"Header-3": []string{"value3"},
			},
			shouldErr: false,
		},
		{
			title:       "Test string Headers",
			input:       "bad headers",
			expected:    http.Header{},
			shouldErr:   true,
			expectedErr: ErrInvalidHeaderType,
		},
	}

	for _, test := range tests {
		requestInfo := NewRequestInformation()
		err := requestInfo.AddHeaders(test.input)

		if !test.shouldErr {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.expected, requestInfo.Headers)
	}
}
