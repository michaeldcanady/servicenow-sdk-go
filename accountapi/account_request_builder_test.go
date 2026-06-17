package accountapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountRequestBuilder_Builders(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	tests := []struct {
		name     string
		builder  any
		expected map[string]string
	}{
		{
			name:     "ByID",
			builder:  builder.ByID("test-id"),
			expected: map[string]string{"account_id": "test-id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.builder)
			if tt.expected != nil {
				rb := tt.builder.(internal.RequestBuilder)
				for k, v := range tt.expected {
					assert.Equal(t, v, rb.GetPathParameters()[k])
				}
			}
		})
	}
}

func TestAccountRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockResponse := &AccountCollectionResponseMock{}
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil)

	res, err := builder.Get(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

type AccountCollectionResponseMock struct {
	internal.BaseServiceNowCollectionResponse[*AccountModel]
}

func (m *AccountCollectionResponseMock) Serialize(writer serialization.SerializationWriter) error {
	return nil
}
func (m *AccountCollectionResponseMock) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func TestAccountItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "account_id": "test-id"}, adapter)

	mockResponse := &AccountItemResponseMock{}
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil)

	res, err := builder.Get(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

type AccountItemResponseMock struct {
	internal.BaseServiceNowItemResponse[*AccountModel]
}

func (m *AccountItemResponseMock) Serialize(writer serialization.SerializationWriter) error {
	return nil
}
func (m *AccountItemResponseMock) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func TestAccountItemRequestBuilder_GetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "account_id": "test-id"}, adapter)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful Request Information Generation",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, requestInfo)
				assert.Equal(t, accountItemURLTemplate, requestInfo.UrlTemplate)
				assert.Equal(t, "test-id", requestInfo.PathParameters["account_id"])
			}
		})
	}
}
