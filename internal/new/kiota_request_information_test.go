package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "no options",
			test: func(t *testing.T) {
				requestInformation, err := NewRequestInformation()

				assert.Nil(t, err)
				assert.Equal(t, &KiotaRequestInformation{abstractions.NewRequestInformation()}, requestInformation)
			},
		},
		{
			name: "construction error",
			test: func(t *testing.T) {
				strct := mocking.NewMockOption[*KiotaRequestInformation]()
				strct.On("Option", mock.AnythingOfType("*internal.KiotaRequestInformation")).Return(errors.New("error"))
				opt := strct.Option

				requestInformation, err := NewRequestInformation(opt)

				assert.Equal(t, errors.New("error"), err)
				assert.Nil(t, requestInformation)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestNewRequestInformationWithMethodAndURLTemplateAndPathParameters(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestKiotaRequestInformation_AddQueryParameters(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {

				params := struct {
					ExampleValue string `url:"example_value,omitempty"`
				}{
					ExampleValue: "test",
				}

				info := &KiotaRequestInformation{
					RequestInformation: abstractions.NewRequestInformation(),
				}

				info.AddQueryParameters(params)

				assert.Equal(t, map[string]any{"example_value": []interface{}{"test"}}, info.QueryParametersAny)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
