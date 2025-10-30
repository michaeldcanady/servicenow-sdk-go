package internal

import (
	"errors"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuildServiceNowRequestAdapterConfig(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				config, err := buildServiceNowRequestAdapterConfig()
				assert.Nil(t, err)
				// can't test config.Client since it can/does contain functions as properties
				assert.Equal(t, serialization.DefaultSerializationWriterFactoryInstance, config.serializationWriterFactory)
				assert.Equal(t, serialization.DefaultParseNodeFactoryInstance, config.parseNodeFactory)
			},
		},
		{
			name: "option error",
			test: func(t *testing.T) {
				strct := newMockServiceNowRequestAdapterOption()
				strct.On("ServiceNowClientOption", mock.IsType(&serviceNowRequestAdapterConfig{})).Return(errors.New("opt error"))
				opt := strct.ServiceNowClientOption
				config, err := buildServiceNowRequestAdapterConfig(opt)
				assert.Nil(t, config)
				assert.Equal(t, errors.New("opt error"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
