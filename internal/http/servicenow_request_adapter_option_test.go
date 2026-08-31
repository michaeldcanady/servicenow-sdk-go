// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package internalhttp

import (
	"errors"
	"net/http"
	"testing"

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
