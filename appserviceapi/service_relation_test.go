// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServiceRelationModel_GettersSetters(t *testing.T) {
	model := NewServiceRelation()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Parent", func(v any) error { return model.setParent(v.(*string)) }, func() (any, error) { return model.GetParent() }, internal.ToPointer("parent123")},
		{"Child", func(v any) error { return model.setChild(v.(*string)) }, func() (any, error) { return model.GetChild() }, internal.ToPointer("child456")},
		{"Type", func(v any) error { return model.setType(v.(*string)) }, func() (any, error) { return model.GetType() }, internal.ToPointer("Depends on::Used by")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestServiceRelationModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ServiceRelation
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewServiceRelation(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ServiceRelation {
				m := NewServiceRelation()
				_ = m.setParent(internal.ToPointer("parent123"))
				_ = m.setChild(internal.ToPointer("child456"))
				_ = m.setType(internal.ToPointer("Depends on::Used by"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ServiceRelation {
				m := NewServiceRelation()
				_ = m.setParent(internal.ToPointer("parent123"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", parentKey, mock.Anything).Return(errWrite)
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

func TestServiceRelationModel_GetFieldDeserializers(t *testing.T) {
	model := NewServiceRelation()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{parentKey, childKey, typeKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestCreateServiceRelationFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateServiceRelationFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
