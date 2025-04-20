package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

// TODO: add tests
func TestApplyOptions(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {

				type configType struct {
				}

				config := &configType{}
				option := mocking.NewMockOption[*configType]()
				option.On("Option", config).Return(nil)

				err := ApplyOptions(config, option.Option)

				assert.Nil(t, err)
				option.AssertExpectations(t)
			},
		},
		{
			name: "Option error",
			test: func(t *testing.T) {

				type configType struct {
				}

				config := &configType{}
				option := mocking.NewMockOption[*configType]()
				option.On("Option", config).Return(errors.New("failed to apply option"))

				err := ApplyOptions(config, option.Option)

				assert.Equal(t, errors.New("failed to apply option"), err)
				option.AssertExpectations(t)
			},
		},
		{
			name: "No options",
			test: func(t *testing.T) {

				type configType struct {
				}

				config := &configType{}
				option := mocking.NewMockOption[*configType]()

				err := ApplyOptions(config)
				assert.Nil(t, err)
				option.AssertExpectations(t)
			},
		},
		{
			name: "No options",
			test: func(t *testing.T) {

				type configType struct {
				}

				config := (*configType)(nil)
				option := mocking.NewMockOption[*configType]()

				err := ApplyOptions(config, option.Option)
				assert.Nil(t, err)
				option.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
