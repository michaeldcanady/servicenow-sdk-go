// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServiceDetailsRequestModel_GettersSetters(t *testing.T) {
	model := NewServiceDetailsRequest()
	details := NewBasicDetails()

	err := model.SetBasicDetails(details)
	require.NoError(t, err)

	got, err := model.GetBasicDetails()
	require.NoError(t, err)
	assert.Equal(t, details, got)

	relationship := NewServiceRelationship()

	err = model.SetRelationships(relationship)
	require.NoError(t, err)

	gotRelationship, err := model.GetRelationships()
	require.NoError(t, err)
	assert.Equal(t, relationship, gotRelationship)
}

func TestServiceDetailsRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ServiceDetailsRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewServiceDetailsRequest(),
		},
		{
			name: "happy path - writes nested basic details",
			model: func() *ServiceDetailsRequest {
				m := NewServiceDetailsRequest()
				_ = m.SetBasicDetails(NewBasicDetails())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", basicDetailsKey, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "nested object write error propagates",
			model: func() *ServiceDetailsRequest {
				m := NewServiceDetailsRequest()
				_ = m.SetBasicDetails(NewBasicDetails())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", basicDetailsKey, mock.Anything, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestServiceDetailsRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewServiceDetailsRequest()
	deserializers := model.GetFieldDeserializers()
	assert.NotNil(t, deserializers[basicDetailsKey])
	assert.Len(t, deserializers, 1)
}

func TestServiceDetailsRequestModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewServiceDetailsRequest()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetObjectValue", mock.Anything).Return(nil, errWrite)

	err := deserializers[basicDetailsKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreateServiceDetailsRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateServiceDetailsRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
