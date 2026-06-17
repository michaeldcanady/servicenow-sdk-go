package mocking

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/testutils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockRequestAdapter(t *testing.T) {
	logger := testutils.NewLogger()
	tests := []struct {
		name     string
		response interface{}
		err      error
	}{
		{
			name:     "Successful Response",
			response: NewMockParsable(),
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
			adapter := NewMockRequestAdapter()
			adapter.On("Send", context.Background(), mock.Anything, mock.Anything, mock.Anything).Return(tt.response, tt.err)

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
			adapter.AssertExpectations(t)
		})
	}
}
