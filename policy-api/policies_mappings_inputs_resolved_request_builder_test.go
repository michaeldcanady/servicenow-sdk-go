package policyapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPoliciesMappingsInputsResolvedRequestBuilderInternal(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}
	requestAdapter := mocking.NewMockRequestAdapter()
	builder := NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(pathParameters, requestAdapter)
	assert.NotNil(t, builder)
}

func TestPoliciesMappingsInputsResolvedRequestBuilder_Get(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}

	tests := []struct {
		name      string
		mock      func(*mocking.MockRequestAdapter)
		builder   *PoliciesMappingsInputsResolvedRequestBuilder
		config    *PoliciesMappingsInputsResolvedRequestBuilderGetRequestConfiguration
		wantErr   bool
		errStderr string
	}{
		{
			name: "Valid",
			mock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(NewBasePoliciesMappingsInputCollectionResponse(), nil)
			},
			wantErr: false,
		},
		{
			name:    "NilRB",
			builder: nil,
			wantErr: false,
		},
		{
			name: "SendError",
			mock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send error"))
			},
			wantErr: true,
		},
		{
			name: "NilResponse",
			mock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "WrongResponseType",
			mock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mocking.NewMockParsable(), nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestAdapter := mocking.NewMockRequestAdapter()
			if tt.mock != nil {
				tt.mock(requestAdapter)
			}

			builder := tt.builder
			if builder == nil && tt.name != "NilRB" {
				builder = NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(pathParameters, requestAdapter)
			}

			res, err := builder.Get(context.Background(), tt.config)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.name == "NilRB" {
					assert.Nil(t, res)
				} else {
					assert.NotNil(t, res)
				}
			}
		})
	}
}

func TestPoliciesMappingsInputsResolvedRequestBuilder_ToGetRequestInformation(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}
	requestAdapter := mocking.NewMockRequestAdapter()
	builder := NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(pathParameters, requestAdapter)

	tests := []struct {
		name    string
		builder *PoliciesMappingsInputsResolvedRequestBuilder
		config  *PoliciesMappingsInputsResolvedRequestBuilderGetRequestConfiguration
	}{
		{"NilConfig", builder, nil},
		{"EmptyConfig", builder, &PoliciesMappingsInputsResolvedRequestBuilderGetRequestConfiguration{}},
		{"NilRB", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqInfo, err := tt.builder.ToGetRequestInformation(context.Background(), tt.config)
			assert.NoError(t, err)
			if tt.builder == nil {
				assert.Nil(t, reqInfo)
			} else {
				assert.NotNil(t, reqInfo)
			}
		})
	}
}

func NewBasePoliciesMappingsInputCollectionResponse() *newInternal.BaseServiceNowCollectionResponse[*PoliciesMappingsInput] {
	return newInternal.NewBaseServiceNowCollectionResponse[*PoliciesMappingsInput](CreatePoliciesMappingsInputFromDiscriminatorValue)
}
