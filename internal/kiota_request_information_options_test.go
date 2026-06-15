package internal

import (
	"errors"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestWithMethod(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				config := &KiotaRequestInformation{
					&abstractions.RequestInformation{},
				}

				opt := WithMethod(abstractions.GET)
				err := opt(config)
				assert.Nil(t, err)
				assert.Equal(t, abstractions.GET, config.Method)
			},
		},
		{
			name: "Nil config",
			test: func(t *testing.T) {
				config := (*KiotaRequestInformation)(nil)

				opt := WithMethod(abstractions.GET)
				err := opt(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithURLTemplate(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				config := &KiotaRequestInformation{
					&abstractions.RequestInformation{},
				}

				opt := WithURLTemplate("template")
				err := opt(config)
				assert.Nil(t, err)
				assert.Equal(t, "template", config.UrlTemplate)
			},
		},
		{
			name: "Empty template",
			test: func(t *testing.T) {
				config := &KiotaRequestInformation{
					&abstractions.RequestInformation{},
				}

				opt := WithURLTemplate(" ")
				err := opt(config)
				assert.Equal(t, errors.New("template is empty"), err)
			},
		},
		{
			name: "Nil config",
			test: func(t *testing.T) {
				config := (*KiotaRequestInformation)(nil)

				opt := WithURLTemplate("template")
				err := opt(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithPathParameters(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				config := &KiotaRequestInformation{
					&abstractions.RequestInformation{},
				}

				opt := WithPathParameters(map[string]string{"baseurl": "url"})
				err := opt(config)
				assert.Nil(t, err)
				assert.Equal(t, map[string]string{"baseurl": "url"}, config.PathParameters)
			},
		},
		{
			name: "Nil parameters",
			test: func(t *testing.T) {
				config := &KiotaRequestInformation{
					&abstractions.RequestInformation{},
				}

				opt := WithPathParameters(nil)
				err := opt(config)
				assert.Equal(t, errors.New("pathParameters is nil"), err)
			},
		},
		{
			name: "Empty parameters",
			test: func(t *testing.T) {
				config := &KiotaRequestInformation{
					&abstractions.RequestInformation{},
				}

				opt := WithPathParameters(map[string]string{})
				err := opt(config)
				assert.Equal(t, errors.New("pathParameters is empty"), err)
			},
		},
		{
			name: "Nil config",
			test: func(t *testing.T) {
				config := (*KiotaRequestInformation)(nil)

				opt := WithPathParameters(map[string]string{})
				err := opt(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
