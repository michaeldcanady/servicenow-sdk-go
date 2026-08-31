// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPopulationMethodModel_GettersSetters(t *testing.T) {
	model := NewPopulationMethod()

	err := model.SetType(internal.ToPointer("discovery"))
	require.NoError(t, err)

	got, err := model.GetType()
	require.NoError(t, err)
	assert.Equal(t, internal.ToPointer("discovery"), got)
}

func TestPopulationMethodModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *PopulationMethodModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewPopulationMethod(),
		},
		{
			name: "happy path - writes type",
			model: func() *PopulationMethodModel {
				m := NewPopulationMethod()
				_ = m.SetType(internal.ToPointer("discovery"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", typeKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *PopulationMethodModel {
				m := NewPopulationMethod()
				_ = m.SetType(internal.ToPointer("discovery"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", typeKey, mock.Anything).Return(errWrite)
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

func TestPopulationMethodModel_GetFieldDeserializers(t *testing.T) {
	model := NewPopulationMethod()
	deserializers := model.GetFieldDeserializers()
	assert.NotNil(t, deserializers[typeKey])
	assert.Len(t, deserializers, 1)
}

func TestCreatePopulationMethodFromDiscriminatorValue(t *testing.T) {
	t.Run("GetChildNode error propagates", func(t *testing.T) {
		parseNode := mocking.NewMockParseNode()
		parseNode.On("GetChildNode", typeKey).Return(mocking.NewMockParseNode(), errWrite)

		result, err := CreatePopulationMethodFromDiscriminatorValue(parseNode)
		require.ErrorIs(t, err, errWrite)
		assert.Nil(t, result)
	})

	t.Run("GetStringValue error propagates", func(t *testing.T) {
		childNode := mocking.NewMockParseNode()
		childNode.On("GetStringValue").Return((*string)(nil), errWrite)

		parseNode := mocking.NewMockParseNode()
		parseNode.On("GetChildNode", typeKey).Return(childNode, nil)

		result, err := CreatePopulationMethodFromDiscriminatorValue(parseNode)
		require.ErrorIs(t, err, errWrite)
		assert.Nil(t, result)
	})

	t.Run("nil type value returns error", func(t *testing.T) {
		childNode := mocking.NewMockParseNode()
		childNode.On("GetStringValue").Return((*string)(nil), nil)

		parseNode := mocking.NewMockParseNode()
		parseNode.On("GetChildNode", typeKey).Return(childNode, nil)

		result, err := CreatePopulationMethodFromDiscriminatorValue(parseNode)
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("unknown type returns error", func(t *testing.T) {
		childNode := mocking.NewMockParseNode()
		childNode.On("GetStringValue").Return(internal.ToPointer("bogus"), nil)

		parseNode := mocking.NewMockParseNode()
		parseNode.On("GetChildNode", typeKey).Return(childNode, nil)

		result, err := CreatePopulationMethodFromDiscriminatorValue(parseNode)
		require.Error(t, err)
		assert.Nil(t, result)
	})

	dispatchTests := []struct {
		typeValue    string
		assertResult func(t *testing.T, v serialization.Parsable)
	}{
		{"cmdb_group_based", func(t *testing.T, v serialization.Parsable) {
			assert.IsType(t, &CmdbGroupBasedPopulationMethodModel{}, v)
		}},
		{"discovery", func(t *testing.T, v serialization.Parsable) { assert.IsType(t, &DiscoveryPopulationMethodModel{}, v) }},
		{"tag_list", func(t *testing.T, v serialization.Parsable) { assert.IsType(t, &TagListPopulationMethodModel{}, v) }},
	}

	for _, tt := range dispatchTests {
		t.Run(tt.typeValue, func(t *testing.T) {
			childNode := mocking.NewMockParseNode()
			childNode.On("GetStringValue").Return(internal.ToPointer(tt.typeValue), nil)

			parseNode := mocking.NewMockParseNode()
			parseNode.On("GetChildNode", typeKey).Return(childNode, nil)

			result, err := CreatePopulationMethodFromDiscriminatorValue(parseNode)
			require.NoError(t, err)
			tt.assertResult(t, result)
		})
	}
}
