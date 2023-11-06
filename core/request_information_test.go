package core

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestInformation(t *testing.T) {
	requestInfo := NewRequestInformation()

	expectedHeaders := make(http.Header)
	expectedOptions := make(map[string]RequestOption)
	expectedUri := NewUrlInformation()

	assert.Equal(t, expectedHeaders, requestInfo.Headers)
	assert.Equal(t, expectedOptions, requestInfo.options)
	assert.Equal(t, expectedUri, requestInfo.uri)
}

func TestNewRequestInformation_SetStreamContenr(t *testing.T) {

	content := []byte("Testing Test")

	requestInfo := NewRequestInformation()
	requestInfo.SetStreamContent(content)

	assert.Equal(t, content, requestInfo.Content)
	assert.Equal(t, binaryContentType, requestInfo.Headers.Get(contentTypeHeader))
}

func TestNewRequestInformation_AddQueryParameters(t *testing.T) {

	source := struct {
		Var1 string `query:"var_1"`
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

func TestNewRequestInformation_SetUri(t *testing.T) {

	url, err := url.Parse("https://www.example.com")
	if err != nil {
		t.Error(err)
	}

	expected := map[string]string{"request-raw-url": "https://www.example.com"}

	requestInfo := NewRequestInformation()
	requestInfo.SetUri(url)

	assert.Equal(t, expected, requestInfo.uri.PathParameters)
}
