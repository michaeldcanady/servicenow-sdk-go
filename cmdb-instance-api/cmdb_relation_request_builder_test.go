package cmdbinstanceapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCmdbRelationRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := newInternal.NewBaseServiceNowItemResponse[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("relation creation failed"))
			},
			expectedErr: errors.New("relation creation failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCmdbRelationRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)
			resp, err := builder.Post(context.Background(), nil, nil)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestCmdbRelationRequestBuilder_ByID(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCmdbRelationRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)

	itemBuilder := builder.ByID("456")

	assert.NotNil(t, itemBuilder)
	assert.Equal(t, "456", itemBuilder.GetPathParameters()["rel_sys_id"])
}
