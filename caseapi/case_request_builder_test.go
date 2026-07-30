package caseapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errNetwork is a stand-in for a transport-level error returned by the adapter.
var errNetwork = errors.New("network error")

// ---------------------------------------------------------------------------
// CaseRequestBuilder
// ---------------------------------------------------------------------------

func TestCaseRequestBuilder_ByID(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseRequestBuilder
	}{
		{
			name:    "valid builder returns child with id path parameter",
			builder: NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, &mocking.MockRequestAdapter{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child := tt.builder.ByID("test-id")
			require.NotNil(t, child)
			assert.Equal(t, "test-id", child.GetPathParameters()["id"])
			assert.Equal(t, caseItemURLTemplate, child.GetURLTemplate())
		})
	}

	t.Run("nil receiver returns nil", func(t *testing.T) {
		var builder *CaseRequestBuilder
		assert.Nil(t, builder.ByID("test-id"))
	})
}

func TestCaseRequestBuilder_FieldValues(t *testing.T) {
	t.Run("valid builder returns child with field_name path parameter", func(t *testing.T) {
		builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, &mocking.MockRequestAdapter{})
		child := builder.FieldValues("state")
		require.NotNil(t, child)
		assert.Equal(t, "state", child.GetPathParameters()["field_name"])
		assert.Equal(t, fieldValuesURLTemplate, child.GetURLTemplate())
	})

	t.Run("nil receiver returns nil", func(t *testing.T) {
		var builder *CaseRequestBuilder
		assert.Nil(t, builder.FieldValues("state"))
	})
}

func TestCaseRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseRequestBuilder
		wantErr error
	}{
		{
			name:    "happy path - builds request information",
			builder: NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, &mocking.MockRequestAdapter{}),
			wantErr: nil,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			builder: nil,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := tt.builder.ToGetRequestInformation(context.Background(), nil)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, requestInfo)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, requestInfo)
			assert.Equal(t, caseURLTemplate, requestInfo.UrlTemplate)
		})
	}
}

func TestCaseRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(m *mocking.MockRequestAdapter)
		nilRecv    bool
		nilAdapter bool
		wantErr    error
		wantNil    bool
	}{
		{
			name: "happy path - returns collection response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			nilRecv: true,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
		{
			name:       "nil request adapter returns ErrNilRequestAdapter",
			nilAdapter: true,
			wantErr:    snerrors.ErrNilRequestAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nilRecv {
				var builder *CaseRequestBuilder
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			if tt.nilAdapter {
				builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestCaseRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseRequestBuilder
		wantErr error
	}{
		{
			name: "happy path - builds request information",
			builder: func() *CaseRequestBuilder {
				adapter := &mocking.MockRequestAdapter{}
				adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				return NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			}(),
			wantErr: nil,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			builder: nil,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := tt.builder.ToPostRequestInformation(context.Background(), NewCaseResult(), nil)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, requestInfo)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, requestInfo)
		})
	}

	t.Run("nil adapter surfaces serialization error", func(t *testing.T) {
		builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)
		requestInfo, err := builder.ToPostRequestInformation(context.Background(), NewCaseResult(), nil)
		require.Error(t, err)
		assert.Nil(t, requestInfo)
	})
}

func TestCaseRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(m *mocking.MockRequestAdapter)
		nilRecv    bool
		nilAdapter bool
		wantErr    error
		wantNil    bool
	}{
		{
			name: "happy path - returns item response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			nilRecv: true,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
		{
			name:       "nil request adapter returns ErrNilRequestAdapter",
			nilAdapter: true,
			wantErr:    snerrors.ErrNilRequestAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nilRecv {
				var builder *CaseRequestBuilder
				resp, err := builder.Post(context.Background(), NewCaseResult(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			if tt.nilAdapter {
				builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)
				resp, err := builder.Post(context.Background(), NewCaseResult(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Post(context.Background(), NewCaseResult(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

// ---------------------------------------------------------------------------
// CaseItemRequestBuilder
// ---------------------------------------------------------------------------

func TestCaseItemRequestBuilder_Activities(t *testing.T) {
	t.Run("valid builder returns child", func(t *testing.T) {
		builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, &mocking.MockRequestAdapter{})
		child := builder.Activities()
		require.NotNil(t, child)
		assert.Equal(t, caseActivitiesURLTemplate, child.GetURLTemplate())
	})

	t.Run("nil receiver returns nil", func(t *testing.T) {
		var builder *CaseItemRequestBuilder
		assert.Nil(t, builder.Activities())
	})
}

func TestCaseItemRequestBuilder_FieldValues(t *testing.T) {
	t.Run("valid builder returns child with field_name path parameter", func(t *testing.T) {
		builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, &mocking.MockRequestAdapter{})
		child := builder.FieldValues("state")
		require.NotNil(t, child)
		assert.Equal(t, "state", child.GetPathParameters()["field_name"])
		assert.Equal(t, itemFieldValuesURLTemplate, child.GetURLTemplate())
	})

	t.Run("nil receiver returns nil", func(t *testing.T) {
		var builder *CaseItemRequestBuilder
		assert.Nil(t, builder.FieldValues("state"))
	})
}

func TestCaseItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseItemRequestBuilder
		wantErr error
	}{
		{
			name:    "happy path - builds request information",
			builder: NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, &mocking.MockRequestAdapter{}),
			wantErr: nil,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			builder: nil,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := tt.builder.ToGetRequestInformation(context.Background(), nil)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, requestInfo)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, requestInfo)
			assert.Equal(t, caseItemURLTemplate, requestInfo.UrlTemplate)
			assert.Equal(t, "test-id", requestInfo.PathParameters["id"])
		})
	}
}

func TestCaseItemRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(m *mocking.MockRequestAdapter)
		nilRecv    bool
		nilAdapter bool
		wantErr    error
		wantNil    bool
	}{
		{
			name: "happy path - returns item response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			nilRecv: true,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
		{
			name:       "nil request adapter returns ErrNilRequestAdapter",
			nilAdapter: true,
			wantErr:    snerrors.ErrNilRequestAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nilRecv {
				var builder *CaseItemRequestBuilder
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			if tt.nilAdapter {
				builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, nil)
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestCaseItemRequestBuilder_ToPutRequestInformation(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseItemRequestBuilder
		wantErr error
	}{
		{
			name: "happy path - builds request information",
			builder: func() *CaseItemRequestBuilder {
				adapter := &mocking.MockRequestAdapter{}
				adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				return NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)
			}(),
			wantErr: nil,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			builder: nil,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := tt.builder.ToPutRequestInformation(context.Background(), NewCaseResult(), nil)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, requestInfo)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, requestInfo)
		})
	}

	t.Run("nil adapter surfaces serialization error", func(t *testing.T) {
		builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, nil)
		requestInfo, err := builder.ToPutRequestInformation(context.Background(), NewCaseResult(), nil)
		require.Error(t, err)
		assert.Nil(t, requestInfo)
	})
}

func TestCaseItemRequestBuilder_Put(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(m *mocking.MockRequestAdapter)
		nilRecv    bool
		nilAdapter bool
		wantErr    error
		wantNil    bool
	}{
		{
			name: "happy path - returns item response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			nilRecv: true,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
		{
			name:       "nil request adapter returns ErrNilRequestAdapter",
			nilAdapter: true,
			wantErr:    snerrors.ErrNilRequestAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nilRecv {
				var builder *CaseItemRequestBuilder
				resp, err := builder.Put(context.Background(), NewCaseResult(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			if tt.nilAdapter {
				builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, nil)
				resp, err := builder.Put(context.Background(), NewCaseResult(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

			resp, err := builder.Put(context.Background(), NewCaseResult(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

// ---------------------------------------------------------------------------
// CaseActivitiesRequestBuilder
// ---------------------------------------------------------------------------

func TestCaseActivitiesRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseActivitiesRequestBuilder
		wantErr error
	}{
		{
			name:    "happy path - builds request information",
			builder: NewCaseActivitiesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, &mocking.MockRequestAdapter{}),
			wantErr: nil,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			builder: nil,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := tt.builder.ToGetRequestInformation(context.Background(), nil)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, requestInfo)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, requestInfo)
			assert.Equal(t, caseActivitiesURLTemplate, requestInfo.UrlTemplate)
		})
	}
}

func TestCaseActivitiesRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(m *mocking.MockRequestAdapter)
		nilRecv    bool
		nilAdapter bool
		wantErr    error
		wantNil    bool
	}{
		{
			name: "happy path - returns item response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*ActivitiesResultModel](CreateActivitiesResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			nilRecv: true,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
		{
			name:       "nil request adapter returns ErrNilRequestAdapter",
			nilAdapter: true,
			wantErr:    snerrors.ErrNilRequestAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nilRecv {
				var builder *CaseActivitiesRequestBuilder
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			if tt.nilAdapter {
				builder := NewCaseActivitiesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, nil)
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCaseActivitiesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

// ---------------------------------------------------------------------------
// CaseFieldValuesRequestBuilder
// ---------------------------------------------------------------------------

func TestNewItemFieldValuesRequestBuilderInternal(t *testing.T) {
	builder := NewItemFieldValuesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id", "field_name": "state"}, &mocking.MockRequestAdapter{})
	require.NotNil(t, builder)
	assert.Equal(t, itemFieldValuesURLTemplate, builder.GetURLTemplate())
}

func TestCaseFieldValuesRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name    string
		builder *CaseFieldValuesRequestBuilder
		wantErr error
	}{
		{
			name:    "happy path - builds request information",
			builder: NewCaseFieldValuesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "field_name": "state"}, &mocking.MockRequestAdapter{}),
			wantErr: nil,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			builder: nil,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := tt.builder.ToGetRequestInformation(context.Background(), nil)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, requestInfo)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, requestInfo)
			assert.Equal(t, fieldValuesURLTemplate, requestInfo.UrlTemplate)
		})
	}
}

func TestCaseFieldValuesRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(m *mocking.MockRequestAdapter)
		nilRecv    bool
		nilAdapter bool
		wantErr    error
		wantNil    bool
	}{
		{
			name: "happy path - returns item response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*FieldValuesResultModel](CreateFieldValuesResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
		{
			name:    "nil receiver returns ErrNilRequestBuilder",
			nilRecv: true,
			wantErr: snerrors.ErrNilRequestBuilder,
		},
		{
			name:       "nil request adapter returns ErrNilRequestAdapter",
			nilAdapter: true,
			wantErr:    snerrors.ErrNilRequestAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nilRecv {
				var builder *CaseFieldValuesRequestBuilder
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			if tt.nilAdapter {
				builder := NewCaseFieldValuesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "field_name": "state"}, nil)
				resp, err := builder.Get(context.Background(), nil)
				assert.Nil(t, resp)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCaseFieldValuesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "field_name": "state"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

// ---------------------------------------------------------------------------
// CustomerServiceRequestBuilder
// ---------------------------------------------------------------------------

func TestCustomerServiceRequestBuilder_Case(t *testing.T) {
	t.Run("valid builder returns child", func(t *testing.T) {
		builder := NewCustomerServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, &mocking.MockRequestAdapter{})
		child := builder.Case()
		require.NotNil(t, child)
		assert.Equal(t, caseURLTemplate, child.GetURLTemplate())
	})

	t.Run("nil receiver returns nil", func(t *testing.T) {
		var builder *CustomerServiceRequestBuilder
		assert.Nil(t, builder.Case())
	})
}
