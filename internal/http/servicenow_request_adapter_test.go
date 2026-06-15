package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServiceNowRequestAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				authProvider := mocking.NewMockAuthenticationProvider()

				_, err := NewServiceNowRequestAdapter(authProvider)
				assert.Nil(t, err)
			},
		},
		{
			name: "nil AuthenticationProvider",
			test: func(t *testing.T) {
				_, err := NewServiceNowRequestAdapter(nil)
				assert.Equal(t, errors.New("authenticationProvider cannot be nil"), err)
			},
		},
		{
			name: "successful",
			test: func(t *testing.T) {
				authProvider := mocking.NewMockAuthenticationProvider()
				strct := mocking.NewMockOption[*serviceNowRequestAdapterConfig]()
				strct.On("Option", mock.IsType(&serviceNowRequestAdapterConfig{})).Return(errors.New("opt error"))
				opt := strct.Option

				_, err := NewServiceNowRequestAdapter(authProvider, opt)
				assert.Equal(t, errors.New("opt error"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
