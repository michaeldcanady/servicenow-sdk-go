package mocking

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/new/testutils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

type MockParsable struct {
	serialization.Parsable
}

func (m *MockParsable) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

func (m *MockParsable) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func TestMockRequestAdapter(t *testing.T) {
	logger := testutils.NewLogger()
	tests := []struct {
		name     string
		response interface{}
		err      error
	}{
		{
			name:     "Successful Response",
			response: &MockParsable{},
			err:      nil,
		},
		{
			name:     "Error Response",
			response: nil,
			err:      assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &MockRequestAdapter{
				Response: tt.response,
				Error:    tt.err,
			}

			requestInfo := abstractions.NewRequestInformation()
			resp, err := adapter.Send(context.Background(), requestInfo, nil, nil)

			if tt.err != nil {
				if !assert.Error(t, err) {
					logger.LogFailure(t.Name(), err, tt)
				}
				assert.Nil(t, resp)
			} else {
				if !assert.NoError(t, err) {
					logger.LogFailure(t.Name(), err, tt)
				}
				assert.Equal(t, tt.response, resp)
			}
			assert.Equal(t, requestInfo, adapter.LastRequest)
		})
	}
}
