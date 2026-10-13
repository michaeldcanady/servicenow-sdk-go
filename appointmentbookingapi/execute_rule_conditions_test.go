// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// ExecuteRuleConditionsRequestModel
// ---------------------------------------------------------------------------

func TestCreateExecuteRuleConditionsRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExecuteRuleConditionsRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestExecuteRuleConditionsRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewExecuteRuleConditionsRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{catalogIDKey, otherInputsKey, taskIDKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

// ---------------------------------------------------------------------------
// ExecuteRuleConditionsResponse
// ---------------------------------------------------------------------------

func TestCreateExecuteRuleConditionsResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExecuteRuleConditionsResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// ExecuteRuleConditionsResult
// ---------------------------------------------------------------------------

func TestCreateExecuteRuleConditionsResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExecuteRuleConditionsResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestExecuteRuleConditionsResult_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ExecuteRuleConditionsResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewExecuteRuleConditionsResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ExecuteRuleConditionsResult {
				m := NewExecuteRuleConditionsResult()
				_ = m.SetDedicatedCapacity(internal.ToPointer(true))
				_ = m.SetFutureMaxBookableDays(internal.ToPointer("30"))
				_ = m.SetRuleID(internal.ToPointer("rule-id"))
				_ = m.SetRuleName(internal.ToPointer("rule-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", dedicatedCapacityKey, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ExecuteRuleConditionsResult {
				m := NewExecuteRuleConditionsResult()
				_ = m.SetDedicatedCapacity(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", dedicatedCapacityKey, mock.Anything).Return(errWrite)
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

func TestExecuteRuleConditionsResult_GetFieldDeserializers(t *testing.T) {
	model := NewExecuteRuleConditionsResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{dedicatedCapacityKey, futureMaxBookableDaysKey, ruleIDKey, ruleNameKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}

func TestExecuteRuleConditionsRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ExecuteRuleConditionsRequestModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			// Unlike the primitive serializers, SerializeAnyFunc has no nil-skip, so an
			// otherwise-empty model still emits the otherInputs key with a nil value.
			name:  "empty model still writes otherInputs",
			model: NewExecuteRuleConditionsRequest(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", otherInputsKey, nil).Return(nil)
			},
		},
		{
			name: "happy path - writes all fields",
			model: func() *ExecuteRuleConditionsRequestModel {
				m := NewExecuteRuleConditionsRequest()
				_ = m.SetCatalogID(internal.ToPointer("catalog-id"))
				_ = m.SetOtherInputs(map[string]any{"key": "value"})
				_ = m.SetTaskID(internal.ToPointer("task-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteAnyValue", otherInputsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ExecuteRuleConditionsRequestModel {
				m := NewExecuteRuleConditionsRequest()
				_ = m.SetCatalogID(internal.ToPointer("catalog-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", catalogIDKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if test.setupMock != nil {
				test.setupMock(writer)
			}

			err := test.model.Serialize(writer)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
