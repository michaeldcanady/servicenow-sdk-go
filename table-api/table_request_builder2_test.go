package tableapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTableRequestBuilder2_Get(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		err       error
	}{
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
		{
			name: "adapter returns wrong type",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(mocking.NewMockParsable(), nil)
			},
			err: errors.New("resp is not *internal.ServiceNowCollectionResponse[*github.com/michaeldcanady/servicenow-sdk-go/table-api.TableRecord]"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := NewTableRequestBuilder2[*TableRecord](
				"https://example.com/api/now/v1/table/test",
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

func TestTableRequestBuilder2_Post(t *testing.T) {
	tests := []struct {
		name      string
		body      *TableRecord
		setupMock func(m *mocking.MockRequestAdapter)
		expectErr bool
		err       error
	}{
		{
			name:      "nil body returns error",
			body:      nil,
			setupMock: func(m *mocking.MockRequestAdapter) {},
			expectErr: true,
			err:       errors.New("body is nil"),
		},
		{
			name: "adapter returns error",
			body: &TableRecord{},
			setupMock: func(m *mocking.MockRequestAdapter) {
				serializationWriter := mocking.NewMockSerializationWriter()
				serializationWriter.On("WriteObjectValue", "", mock.AnythingOfType("*tableapi.TableRecord"), mock.Anything).Return(nil)
				serializationWriter.On("Close").Return(nil)
				serializationWriter.On("GetSerializedContent").Return([]byte(""), nil)

				serializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				serializationWriterFactory.On("GetSerializationWriter", "application/json").Return(serializationWriter, nil)

				m.On("GetSerializationWriterFactory").Return(serializationWriterFactory)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("network error"))
			},
			expectErr: true,
			err:       errors.New("network error"),
		},
		{
			name: "adapter returns wrong type",
			body: &TableRecord{},
			setupMock: func(m *mocking.MockRequestAdapter) {
				serializationWriter := mocking.NewMockSerializationWriter()
				serializationWriter.On("WriteObjectValue", "", mock.AnythingOfType("*tableapi.TableRecord"), mock.Anything).Return(nil)
				serializationWriter.On("Close").Return(nil)
				serializationWriter.On("GetSerializedContent").Return([]byte(""), nil)

				serializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				serializationWriterFactory.On("GetSerializationWriter", "application/json").Return(serializationWriter, nil)

				m.On("GetSerializationWriterFactory").Return(serializationWriterFactory)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(mocking.NewMockParsable(), nil)
			},
			expectErr: true,
			err:       errors.New("resp is not *internal.ServiceNowItemResponse[*github.com/michaeldcanady/servicenow-sdk-go/table-api.TableRecord]"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdapter := new(mocking.MockRequestAdapter)
			tt.setupMock(mockAdapter)

			builder := NewTableRequestBuilder2[*TableRecord](
				"https://example.com/api/now/v1/table/test",
				mockAdapter,
				CreateTableRecordFromDiscriminatorValue,
			)

			resp, err := builder.Post(context.Background(), tt.body, nil)
			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockAdapter.AssertExpectations(t)
		})
	}
}
