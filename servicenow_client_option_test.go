package servicenowsdkgo

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
)

func TestWithURL(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				input := "https://exampleurl.com"
				option := withURL(input)
				config := &serviceNowServiceClientConfig{}
				err := option(config)
				assert.Nil(t, err)
				assert.Equal(t, input, config.rawURI)
			},
		},
		{
			name: "empty uri",
			test: func(t *testing.T) {
				input := " "
				option := withURL(input)
				config := &serviceNowServiceClientConfig{}
				err := option(config)
				assert.Equal(t, errors.New("url is empty"), err)
				assert.Equal(t, "", config.rawURI)
			},
		},
		{
			name: "invalid uri",
			test: func(t *testing.T) {
				input := "https://example url.com"
				option := withURL(input)
				config := &serviceNowServiceClientConfig{}
				err := option(config)
				assert.Equal(t, errors.New("parse \"https://example url.com\": invalid character \" \" in host name"), err)
				assert.Equal(t, "", config.rawURI)
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				input := "https://exampleurl.com"
				option := withURL(input)
				config := (*serviceNowServiceClientConfig)(nil)
				err := option(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithMiddleware(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				config := &serviceNowServiceClientConfig{}
				middleware := mocking.NewMockMiddleware()

				option := withMiddleware(middleware)

				err := option(config)

				assert.Nil(t, err)
				assert.Equal(t, middleware, config.middleware[0])
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				middleware := mocking.NewMockMiddleware()
				config := (*serviceNowServiceClientConfig)(nil)

				option := withMiddleware(middleware)
				err := option(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
		{
			name: "no middleware",
			test: func(t *testing.T) {
				config := &serviceNowServiceClientConfig{}

				option := withMiddleware()

				err := option(config)

				assert.Equal(t, errors.New("middleware is empty"), err)
				assert.Equal(t, ([]nethttplibrary.Middleware)(nil), config.middleware)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithInstance(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				instance := "test"
				config := &serviceNowServiceClientConfig{}

				option := withInstance(instance)
				err := option(config)

				assert.Nil(t, err)
				assert.Equal(t, instance, config.instance)
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				instance := "https://exampleurl.com"
				config := (*serviceNowServiceClientConfig)(nil)

				option := withURL(instance)

				err := option(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
		{
			name: "empty instance",
			test: func(t *testing.T) {
				instance := " "
				config := &serviceNowServiceClientConfig{}

				option := withInstance(instance)
				err := option(config)

				assert.Equal(t, errors.New("instance is empty"), err)
				assert.Equal(t, "", config.instance)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
