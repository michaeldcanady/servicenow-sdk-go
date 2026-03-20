package servicenowsdkgo

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
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
				authProvider := mocking.NewMockAuthenticationProvider()
				config, err := buildServiceClientConfig(WithInstance("test"), WithAuthenticationProvider(authProvider))
				assert.Nil(t, err)
				assert.NotNil(t, config)
			},
		},
		{
			name: "option error",
			test: func(t *testing.T) {
				strct := newMockServiceNowClientOption()
				strct.On("ServiceNowServiceClientOption", mock.IsType(&ServiceNowServiceClientConfig{})).Return(errors.New("option error"))
				option := strct.ServiceNowServiceClientOption

				config, err := buildServiceClientConfig(option)
				assert.Equal(t, errors.New("option error"), err)
				assert.Nil(t, config)
			},
		},
		{
			name: "missing auth and adapter",
			test: func(t *testing.T) {
				config, err := buildServiceClientConfig(WithInstance("test"))
				assert.Equal(t, errors.New("must provide either an AuthenticationProvider or a RequestAdapter"), err)
				assert.Nil(t, config)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
