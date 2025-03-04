package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
)

func TestWithMiddleware(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				config := &serviceNowClientConfig{}
				middleware := mocking.NewMockMiddleware()

				option := WithMiddleware(middleware)
				err := option(config)
				assert.Nil(t, err)
				assert.Equal(t, &serviceNowClientConfig{
					middleware: []nethttplibrary.Middleware{middleware},
				}, config)
			},
		},
		{
			name: "missing middleware",
			test: func(t *testing.T) {
				config := &serviceNowClientConfig{}

				option := WithMiddleware()
				err := option(config)
				assert.Equal(t, errors.New("middleware is empty"), err)
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				middleware := mocking.NewMockMiddleware()
				option := WithMiddleware(middleware)
				err := option(nil)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
