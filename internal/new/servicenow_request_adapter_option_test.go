package internal

import (
	"errors"
	"net/http"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
				assert.Nil(t, err)
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
				assert.Nil(t, err)
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
				assert.Nil(t, err)
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
				assert.Nil(t, err)
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
