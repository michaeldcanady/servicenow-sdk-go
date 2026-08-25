package internalhttp

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWithClient(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				client := &http.Client{}
				config := &serviceNowRequestAdapterConfig{}

				opt := WithClient(client)
				err := opt(config)
				require.NoError(t, err)
				assert.Equal(t, &serviceNowRequestAdapterConfig{
					client: client,
				}, config)
			},
		},
		{
			name: "nil client",
			test: func(t *testing.T) {
				config := &serviceNowRequestAdapterConfig{}

				opt := WithClient(nil)
				err := opt(config)
				assert.Equal(t, errors.New("client is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithParseNodeFactory(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				factory := serialization.DefaultParseNodeFactoryInstance
				config := &serviceNowRequestAdapterConfig{}

				opt := WithParseNodeFactory(factory)
				err := opt(config)
				require.NoError(t, err)
				assert.Equal(t, &serviceNowRequestAdapterConfig{
					parseNodeFactory: factory,
				}, config)
			},
		},
		{
			name: "nil factory",
			test: func(t *testing.T) {
				config := &serviceNowRequestAdapterConfig{}

				opt := WithParseNodeFactory(nil)
				err := opt(config)
				assert.Equal(t, errors.New("factory is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithSerializationFactory(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				factory := serialization.DefaultSerializationWriterFactoryInstance
				config := &serviceNowRequestAdapterConfig{}

				opt := WithSerializationFactory(factory)
				err := opt(config)
				require.NoError(t, err)
				assert.Equal(t, &serviceNowRequestAdapterConfig{
					serializationWriterFactory: factory,
				}, config)
			},
		},
		{
			name: "nil factory",
			test: func(t *testing.T) {
				config := &serviceNowRequestAdapterConfig{}

				opt := WithSerializationFactory(nil)
				err := opt(config)
				assert.Equal(t, errors.New("factory is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServiceNowRequestAdapterDefaultOptions(t *testing.T) {
	tests := []struct {
		name   string
		config *serviceNowRequestAdapterConfig
		verify func(*testing.T, *serviceNowRequestAdapterConfig)
	}{
		{
			name:   "empty config gets all defaults",
			config: &serviceNowRequestAdapterConfig{},
			verify: func(t *testing.T, config *serviceNowRequestAdapterConfig) {
				assert.NotNil(t, config.client)
				assert.Equal(t, serialization.DefaultSerializationWriterFactoryInstance, config.serializationWriterFactory)
				assert.Equal(t, serialization.DefaultParseNodeFactoryInstance, config.parseNodeFactory)
			},
		},
		{
			name: "existing client without middleware is left untouched",
			config: &serviceNowRequestAdapterConfig{
				client: &http.Client{Timeout: 42},
			},
			verify: func(t *testing.T, config *serviceNowRequestAdapterConfig) {
				assert.Equal(t, &http.Client{Timeout: 42}, config.client)
			},
		},
		{
			name: "existing client with middleware is rebuilt via GetDefaultClient",
			config: &serviceNowRequestAdapterConfig{
				client:     &http.Client{Timeout: 42},
				middleware: []nethttplibrary.Middleware{nethttplibrary.NewHeadersInspectionHandler()},
			},
			verify: func(t *testing.T, config *serviceNowRequestAdapterConfig) {
				assert.NotNil(t, config.client)
				assert.NotEqual(t, &http.Client{Timeout: 42}, config.client)
			},
		},
		{
			name: "existing factories are preserved",
			config: &serviceNowRequestAdapterConfig{
				serializationWriterFactory: serialization.DefaultSerializationWriterFactoryInstance,
				parseNodeFactory:           serialization.DefaultParseNodeFactoryInstance,
			},
			verify: func(t *testing.T, config *serviceNowRequestAdapterConfig) {
				assert.Equal(t, serialization.DefaultSerializationWriterFactoryInstance, config.serializationWriterFactory)
				assert.Equal(t, serialization.DefaultParseNodeFactoryInstance, config.parseNodeFactory)
			},
		},
		{
			name: "logger builds a default client that logs requests",
			config: func() *serviceNowRequestAdapterConfig {
				config := &serviceNowRequestAdapterConfig{}
				require.NoError(t, WithLogger(&captureLogger{})(config))

				return config
			}(),
			verify: func(t *testing.T, config *serviceNowRequestAdapterConfig) {
				require.NotNil(t, config.client)

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()

				req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL+"/api/now/v1/table/incident", nil)
				require.NoError(t, err)

				resp, err := config.client.Do(req)
				require.NoError(t, err)
				require.NoError(t, resp.Body.Close())

				output := config.logger.(*captureLogger).output()
				assert.Contains(t, output, "DEBUG ")
				assert.Contains(t, output, "GET ")
				assert.Contains(t, output, "/api/now/v1/table/incident")
				assert.Contains(t, output, "INFO ")
				assert.Contains(t, output, "-> 200")
			},
		},
		{
			name: "existing client is left untouched even with a logger",
			config: func() *serviceNowRequestAdapterConfig {
				config := &serviceNowRequestAdapterConfig{}
				require.NoError(t, WithClient(&http.Client{Timeout: 42})(config))
				require.NoError(t, WithLogger(&captureLogger{})(config))

				return config
			}(),
			verify: func(t *testing.T, config *serviceNowRequestAdapterConfig) {
				assert.Equal(t, &http.Client{Timeout: 42}, config.client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := serviceNowRequestAdapterDefaultOptions()
			err := opt(tt.config)
			require.NoError(t, err)
			tt.verify(t, tt.config)
		})
	}
}

func TestWithLogger(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				logger := &internal.NoOpLogger{}
				config := &serviceNowRequestAdapterConfig{}

				opt := WithLogger(logger)
				err := opt(config)
				require.NoError(t, err)
				assert.Equal(t, &serviceNowRequestAdapterConfig{logger: logger}, config)
			},
		},
		{
			name: "nil logger",
			test: func(t *testing.T) {
				config := &serviceNowRequestAdapterConfig{}

				opt := WithLogger(nil)
				err := opt(config)
				assert.Equal(t, errors.New("logger is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithServiceNowClientOptions(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				config := &serviceNowRequestAdapterConfig{}

				option := WithServiceNowClientOptions()
				err := option(config)
				require.NoError(t, err)
				assert.IsType(t, &http.Client{}, config.client)
				assert.NotNil(t, config.client)
			},
		},
		{
			name: "option error",
			test: func(t *testing.T) {
				strct := newMockServiceNowClientOption()
				strct.On("ServiceNowClientOption", mock.IsType(&serviceNowClientConfig{})).Return(errors.New("opt error"))
				opt := strct.ServiceNowClientOption
				config := &serviceNowRequestAdapterConfig{}

				option := WithServiceNowClientOptions(opt)
				err := option(config)
				assert.Equal(t, errors.New("opt error"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
