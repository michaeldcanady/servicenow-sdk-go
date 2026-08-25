package servicenowsdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
				require.NoError(t, err)
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

type wiringCaptureLogger struct {
	mu       sync.Mutex
	messages []string
}

func (l *wiringCaptureLogger) Log(message string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, fmt.Sprintf(message, args...))
}

func (l *wiringCaptureLogger) output() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	return strings.Join(l.messages, "\n")
}

func TestGetRequestAdapter_LoggerWiring(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":[]}`))
	}))
	defer server.Close()

	logger := &wiringCaptureLogger{}
	authProvider := mocking.NewMockAuthenticationProvider()
	authProvider.On("AuthenticateRequest", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	client, err := NewServiceNowServiceClient(
		WithURL(server.URL),
		WithAuthenticationProvider(authProvider),
		WithLogger(logger),
	)
	require.NoError(t, err)

	_, _ = client.Now().Table("incident").Get(context.Background(), nil)

	output := logger.output()
	assert.Contains(t, output, "DEBUG ")
	assert.Contains(t, output, "GET http://")
	assert.Contains(t, output, "/api/now/v1/table/incident")
	assert.Contains(t, output, "INFO ")
	assert.Contains(t, output, "-> 200")
}
