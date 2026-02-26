package tableapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTableItemRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		err       error
	}{
		{
			name: "Successful",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(newInternal.NewBaseServiceNowItemResponse[*TableRecord](CreateTableRecordFromDiscriminatorValue), nil)
			},
			err: nil,
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
			err: errors.New("response is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := NewTableItemRequestBuilder3[*TableRecord](
				"https://example.com/api/now/v1/table/test/sysid",
				mockAdapter,
				CreateTableRecordFromDiscriminatorValue,
			)

			resp, err := builder.Get(context.Background(), nil)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockAdapter.AssertExpectations(t)
		})
	}
}

func TestTableItemRequestBuilder_Delete(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		err       error
	}{
		{
			name: "Successful",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			err: nil,
		},
		{
			name: "adapter returns error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("network error"))
			},
			err: errors.New("network error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := NewTableItemRequestBuilder3[*TableRecord](
				"https://example.com/api/now/v1/table/test/sysid",
				mockAdapter,
				CreateTableRecordFromDiscriminatorValue,
			)

			err := builder.Delete(context.Background(), nil)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
			} else {
				require.NoError(t, err)
			}

			mockAdapter.AssertExpectations(t)
		})
	}
}

func TestTableItemRequestBuilder_Put(t *testing.T) {
	tests := []struct {
		name      string
		body      *TableRecord
		setupMock func(m *mocking.MockRequestAdapter)
		err       error
	}{
		{
			name: "Successful",
			body: NewTableRecord(),
			setupMock: func(m *mocking.MockRequestAdapter) {
				serializationWriter := mocking.NewMockSerializationWriter()
				serializationWriter.On("WriteObjectValue", "", mock.AnythingOfType("*tableapi.TableRecord"), mock.Anything).Return(nil)
				serializationWriter.On("Close").Return(nil)
				serializationWriter.On("GetSerializedContent").Return([]byte("{}"), nil)

				serializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				serializationWriterFactory.On("GetSerializationWriter", "application/json").Return(serializationWriter, nil)

				m.On("GetSerializationWriterFactory").Return(serializationWriterFactory)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(newInternal.NewBaseServiceNowItemResponse[*TableRecord](CreateTableRecordFromDiscriminatorValue), nil)
			},
			err: nil,
		},
		{
			name: "nil body",
			body: nil,
			setupMock: func(m *mocking.MockRequestAdapter) {
			},
			err: errors.New("body is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := NewTableItemRequestBuilder3[*TableRecord](
				"https://example.com/api/now/v1/table/test/sysid",
				mockAdapter,
				CreateTableRecordFromDiscriminatorValue,
			)

			resp, err := builder.Put(context.Background(), tt.body, nil)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockAdapter.AssertExpectations(t)
		})
	}
}

func TestTableItemRequestBuilder_Patch(t *testing.T) {
	tests := []struct {
		name      string
		body      *TableRecord
		setupMock func(m *mocking.MockRequestAdapter)
		err       error
	}{
		{
			name: "Successful",
			body: NewTableRecord(),
			setupMock: func(m *mocking.MockRequestAdapter) {
				serializationWriter := mocking.NewMockSerializationWriter()
				serializationWriter.On("WriteObjectValue", "", mock.AnythingOfType("*tableapi.TableRecord"), mock.Anything).Return(nil)
				serializationWriter.On("Close").Return(nil)
				serializationWriter.On("GetSerializedContent").Return([]byte("{}"), nil)

				serializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				serializationWriterFactory.On("GetSerializationWriter", "application/json").Return(serializationWriter, nil)

				m.On("GetSerializationWriterFactory").Return(serializationWriterFactory)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(newInternal.NewBaseServiceNowItemResponse[*TableRecord](CreateTableRecordFromDiscriminatorValue), nil)
			},
			err: nil,
		},
		{
			name: "nil body",
			body: nil,
			setupMock: func(m *mocking.MockRequestAdapter) {
			},
			err: errors.New("body is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := NewTableItemRequestBuilder3[*TableRecord](
				"https://example.com/api/now/v1/table/test/sysid",
				mockAdapter,
				CreateTableRecordFromDiscriminatorValue,
			)

			resp, err := builder.Patch(context.Background(), tt.body, nil)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockAdapter.AssertExpectations(t)
		})
	}
}

func TestTableItemRequestBuilder_ToRequestInformation(t *testing.T) {
	mockAdapter := new(mocking.MockRequestAdapter)
	builder := NewTableItemRequestBuilder3[*TableRecord](
		"https://example.com/api/now/v1/table/test/sysid",
		mockAdapter,
		CreateTableRecordFromDiscriminatorValue,
	)

	t.Run("ToGetRequestInformation", func(t *testing.T) {
		config := &TableItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &TableItemRequestBuilderGetQueryParameters{
				DisplayValue: DisplayValue2All,
			},
		}
		requestInfo, err := builder.ToGetRequestInformation(context.Background(), config)
		require.NoError(t, err)
		assert.Equal(t, abstractions.GET, requestInfo.Method)
	})

	t.Run("ToDeleteRequestInformation", func(t *testing.T) {
		requestInfo, err := builder.ToDeleteRequestInformation(context.Background(), nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.DELETE, requestInfo.Method)
	})

	t.Run("ToPutRequestInformation", func(t *testing.T) {
		serializationWriter := mocking.NewMockSerializationWriter()
		serializationWriter.On("WriteObjectValue", "", mock.AnythingOfType("*tableapi.TableRecord"), mock.Anything).Return(nil)
		serializationWriter.On("Close").Return(nil)
		serializationWriter.On("GetSerializedContent").Return([]byte("{}"), nil)

		serializationWriterFactory := mocking.NewMockSerializationWriterFactory()
		serializationWriterFactory.On("GetSerializationWriter", "application/json").Return(serializationWriter, nil)

		mockAdapter.On("GetSerializationWriterFactory").Return(serializationWriterFactory)

		requestInfo, err := builder.ToPutRequestInformation(context.Background(), NewTableRecord(), nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.PUT, requestInfo.Method)
	})

	t.Run("ToPatchRequestInformation", func(t *testing.T) {
		serializationWriter := mocking.NewMockSerializationWriter()
		serializationWriter.On("WriteObjectValue", "", mock.AnythingOfType("*tableapi.TableRecord"), mock.Anything).Return(nil)
		serializationWriter.On("Close").Return(nil)
		serializationWriter.On("GetSerializedContent").Return([]byte("{}"), nil)

		serializationWriterFactory := mocking.NewMockSerializationWriterFactory()
		serializationWriterFactory.On("GetSerializationWriter", "application/json").Return(serializationWriter, nil)

		mockAdapter.On("GetSerializationWriterFactory").Return(serializationWriterFactory)

		requestInfo, err := builder.ToPatchRequestInformation(context.Background(), NewTableRecord(), nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.PATCH, requestInfo.Method)
	})
}
