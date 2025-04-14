package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}
				requestBuilder := NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters)

				assert.Equal(t, &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}, requestBuilder)
			},
		},
		{
			name: "nil pathParameter",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				requestBuilder := NewBaseRequestBuilder(requestAdapter, urlTemplate, nil)

				assert.Equal(t, &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: map[string]string{},
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}, requestBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestBuilder_GetPathParameters(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}

				assert.Equal(t, map[string]string{}, requestBuilder.GetPathParameters())
			},
		},
		{
			name: "nil requestBuilder",
			test: func(t *testing.T) {
				requestBuilder := (*BaseRequestBuilder)(nil)

				assert.Nil(t, requestBuilder.GetPathParameters())
			},
		},
		{
			name: "nil pathParameter",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: nil,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}

				assert.Nil(t, requestBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestBuilder_SetPathParameters(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: nil,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}

				err := requestBuilder.SetPathParameters(pathParameters)

				assert.Nil(t, err)
				assert.Equal(t, pathParameters, requestBuilder.PathParameters)
			},
		},
		{
			name: "nil requestBuilder",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}

				requestBuilder := (*BaseRequestBuilder)(nil)

				err := requestBuilder.SetPathParameters(pathParameters)

				assert.Nil(t, err)
			},
		},
		{
			name: "nil pathParameters",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}

				err := requestBuilder.SetPathParameters(nil)

				assert.Equal(t, errors.New("pathParameters is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestBuilder_GetRequestAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}

				assert.Equal(t, requestAdapter, requestBuilder.GetRequestAdapter())
			},
		},
		{
			name: "nil requestBuilder",
			test: func(t *testing.T) {
				requestBuilder := (*BaseRequestBuilder)(nil)

				assert.Nil(t, requestBuilder.GetRequestAdapter())
			},
		},
		{
			name: "nil requestAdapter",
			test: func(t *testing.T) {
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: nil,
				}}

				assert.Nil(t, requestBuilder.GetRequestAdapter())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestBuilder_SetRequestAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: nil,
				}}

				err := requestBuilder.SetRequestAdapter(requestAdapter)

				assert.Nil(t, err)
				assert.Equal(t, requestAdapter, requestBuilder.RequestAdapter)
			},
		},
		{
			name: "nil requestBuilder",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()

				requestBuilder := (*BaseRequestBuilder)(nil)

				err := requestBuilder.SetRequestAdapter(requestAdapter)

				assert.Nil(t, err)
			},
		},
		{
			name: "nil requestAdapter",
			test: func(t *testing.T) {
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: nil,
				}}

				err := requestBuilder.SetRequestAdapter(nil)

				assert.Equal(t, errors.New("requestAdapter is nil"), err)
				assert.Nil(t, requestBuilder.RequestAdapter)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestBuilder_GetURLTemplate(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    urlTemplate,
					RequestAdapter: requestAdapter,
				}}

				template := requestBuilder.GetURLTemplate()

				assert.Equal(t, urlTemplate, template)
			},
		},
		{
			name: "nil requestBuilder",
			test: func(t *testing.T) {
				requestBuilder := (*BaseRequestBuilder)(nil)

				template := requestBuilder.GetURLTemplate()

				assert.Equal(t, "", template)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestBuilder_SetURLTemplate(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				urlTemplate := "template"
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    "",
					RequestAdapter: requestAdapter,
				}}

				err := requestBuilder.SetURLTemplate(urlTemplate)

				assert.Nil(t, err)
				assert.Equal(t, urlTemplate, requestBuilder.UrlTemplate)
			},
		},
		{
			name: "nil requestBuilder",
			test: func(t *testing.T) {
				urlTemplate := "template"

				requestBuilder := (*BaseRequestBuilder)(nil)

				err := requestBuilder.SetURLTemplate(urlTemplate)

				assert.Nil(t, err)
			},
		},
		{
			name: "empty template",
			test: func(t *testing.T) {
				requestAdapter := mocking.NewMockRequestAdapter()
				pathParameters := map[string]string{}

				requestBuilder := &BaseRequestBuilder{abstractions.BaseRequestBuilder{
					PathParameters: pathParameters,
					UrlTemplate:    "",
					RequestAdapter: requestAdapter,
				}}

				err := requestBuilder.SetURLTemplate("")

				assert.Equal(t, errors.New("urlTemplate is empty"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
