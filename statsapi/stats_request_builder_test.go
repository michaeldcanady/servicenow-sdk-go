package statsapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewStatsRequestBuilder(t *testing.T) {
	mockAdapter := new(mocking.MockRequestAdapter)
	builder := NewStatsRequestBuilder("https://example.com/api/now/stats/incident", mockAdapter)

	require.NotNil(t, builder)
	assert.Equal(t, "https://example.com/api/now/stats/incident", builder.GetPathParameters()[internal.RawURLKey])
}

func TestStatsRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name      string
		builder   *StatsRequestBuilder
		setupMock func(m *mocking.MockRequestAdapter)
		err       error
	}{
		{
			name:      "nil builder",
			builder:   nil,
			setupMock: func(m *mocking.MockRequestAdapter) {},
			err:       snerrors.ErrNilRequestBuilder,
		},
		{
			name:      "nil request adapter",
			builder:   NewStatsRequestBuilderInternal(map[string]string{}, nil),
			setupMock: func(m *mocking.MockRequestAdapter) {},
			err:       snerrors.ErrNilRequestAdapter,
		},
		{
			name: "adapter returns error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("network error"))
			},
			err: errors.New("network error"),
		},
		{
			name: "adapter returns nil response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			err: snerrors.ErrNilResponse,
		},
		{
			name: "adapter returns wrong type",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(mocking.NewMockParsable(), nil)
			},
			err: errors.New("resp is not *core.ServiceNowItemResponse[*github.com/michaeldcanady/servicenow-sdk-go/statsapi.StatsResultModel]"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := tt.builder
			if builder == nil && tt.name != "nil builder" {
				builder = NewStatsRequestBuilder("https://example.com/api/now/stats/incident", mockAdapter)
			}

			resp, err := builder.Get(context.Background(), nil)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
				assert.Nil(t, resp)
			} else if tt.name == "nil builder" {
				require.NoError(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockAdapter.AssertExpectations(t)
		})
	}
}

func TestStatsRequestBuilder_ToGetRequestInformation(t *testing.T) {
	builder := NewStatsRequestBuilder("https://example.com/api/now/stats/incident", new(mocking.MockRequestAdapter))

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.NoError(t, err)
	require.NotNil(t, requestInfo)
}

func TestStatsRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *StatsRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}
