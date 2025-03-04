package servicenowsdkgo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuildServiceClientConfig(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				strct := newMockServiceNowClientOption()
				strct.On("ServiceNowServiceClientOption", mock.IsType(&serviceNowServiceClientConfig{})).Return(nil)
				option := strct.ServiceNowServiceClientOption

				config, err := buildServiceClientConfig(option)
				assert.Nil(t, err)
				assert.NotNil(t, config)
			},
		},
		{
			name: "option error",
			test: func(t *testing.T) {
				strct := newMockServiceNowClientOption()
				strct.On("ServiceNowServiceClientOption", mock.IsType(&serviceNowServiceClientConfig{})).Return(errors.New("option error"))
				option := strct.ServiceNowServiceClientOption

				config, err := buildServiceClientConfig(option)
				assert.Equal(t, errors.New("option error"), err)
				assert.Nil(t, config)
			},
		},
		{
			name: "instance and raw uri",
			test: func(t *testing.T) {
				config, err := buildServiceClientConfig(withInstance("fdasfsda"), withURL("https://dafsfd.com"))
				assert.Equal(t, errors.New("rawURL and instance cannot be used together"), err)
				assert.Nil(t, config)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
