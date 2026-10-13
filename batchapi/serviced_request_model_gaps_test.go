// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package batchapi

import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errNode is a stand-in for a parse-node read failure.
var errNode = errors.New("parse node error")

// TestServicedRequestModel_ExecutionTimeDeserializer covers the execution_time deserializer,
// which has to cope with ServiceNow returning a millisecond number where an ISO 8601 duration
// is nominally expected.
func TestServicedRequestModel_ExecutionTimeDeserializer(t *testing.T) {
	t.Run("a millisecond number is converted to a duration", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return(internal.ToPointer(1500.0), nil)

		require.NoError(t, model.GetFieldDeserializers()[executionTimeKey](node))

		executionTime, err := model.GetExecutionTime()
		require.NoError(t, err)
		require.NotNil(t, executionTime)
		assert.Equal(t, 1, executionTime.GetSeconds())
		assert.Equal(t, 500, executionTime.GetMilliSeconds())
	})

	t.Run("a nil number falls through to the ISO duration reader", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return((*float64)(nil), nil)
		node.On("GetISODurationValue").Return(serialization.NewDuration(0, 0, 0, 0, 0, 2, 250), nil)

		require.NoError(t, model.GetFieldDeserializers()[executionTimeKey](node))

		executionTime, err := model.GetExecutionTime()
		require.NoError(t, err)
		require.NotNil(t, executionTime)
		assert.Equal(t, 2, executionTime.GetSeconds())
	})

	t.Run("a number read error falls through to the ISO duration reader", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return((*float64)(nil), errNode)
		node.On("GetISODurationValue").Return(serialization.NewDuration(0, 0, 0, 0, 0, 3, 0), nil)

		require.NoError(t, model.GetFieldDeserializers()[executionTimeKey](node))

		executionTime, err := model.GetExecutionTime()
		require.NoError(t, err)
		require.NotNil(t, executionTime)
		assert.Equal(t, 3, executionTime.GetSeconds())
	})

	t.Run("an ISO duration read error propagates", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return((*float64)(nil), nil)
		node.On("GetISODurationValue").Return((*serialization.ISODuration)(nil), errNode)

		require.ErrorIs(t, model.GetFieldDeserializers()[executionTimeKey](node), errNode)
	})
}

// TestServicedRequestModel_StatusCodeDeserializer covers the status_code deserializer, which
// reads a float because ServiceNow's JSON numbers land as float64 in the parse tree.
func TestServicedRequestModel_StatusCodeDeserializer(t *testing.T) {
	t.Run("a number is narrowed to an int64", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return(internal.ToPointer(200.0), nil)

		require.NoError(t, model.GetFieldDeserializers()[statusCodeKey](node))

		statusCode, err := model.GetStatusCode()
		require.NoError(t, err)
		require.NotNil(t, statusCode)
		assert.Equal(t, int64(200), *statusCode)
	})

	t.Run("a nil number leaves the status code unset", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return((*float64)(nil), nil)

		require.NoError(t, model.GetFieldDeserializers()[statusCodeKey](node))

		statusCode, err := model.GetStatusCode()
		require.NoError(t, err)
		assert.Nil(t, statusCode)
	})

	t.Run("a read error propagates", func(t *testing.T) {
		model := NewServicedRequest()
		node := mocking.NewMockParseNode()
		node.On("GetFloat64Value").Return((*float64)(nil), errNode)

		require.ErrorIs(t, model.GetFieldDeserializers()[statusCodeKey](node), errNode)
	})
}

// TestServicedRequestModel_BodyDeserializer covers the base64 body deserializer.
func TestServicedRequestModel_BodyDeserializer(t *testing.T) {
	tests := []struct {
		name     string
		raw      *string
		expected []byte
		wantErr  bool
	}{
		{
			name:     "base64 body is decoded",
			raw:      internal.ToPointer(base64.StdEncoding.EncodeToString([]byte(`{"result":1}`))),
			expected: []byte(`{"result":1}`),
		},
		{
			name:     "nil body decodes to no body",
			raw:      nil,
			expected: nil,
		},
		{
			name:    "a body that is not base64 is rejected",
			raw:     internal.ToPointer("not-base64!!!"),
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model := NewServicedRequest()
			node := mocking.NewMockParseNode()
			node.On("GetStringValue").Return(test.raw, nil)

			err := model.GetFieldDeserializers()[bodyKey](node)

			if test.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			body, err := model.GetBody()
			require.NoError(t, err)
			assert.Equal(t, test.expected, body)
		})
	}
}

// TestServicedRequestModel_SerializeGaps covers the body mutator in Serialize, which the
// existing Serialize test does not populate.
func TestServicedRequestModel_SerializeGaps(t *testing.T) {
	t.Run("body is base64 encoded", func(t *testing.T) {
		model := NewServicedRequest()
		require.NoError(t, model.setBody([]byte("hello")))

		writer := newPermissiveWriter()

		require.NoError(t, model.Serialize(writer))

		encoded := base64.StdEncoding.EncodeToString([]byte("hello"))
		writer.AssertCalled(t, "WriteStringValue", bodyKey, internal.ToPointer(encoded))
	})

	t.Run("an absent body writes nothing for the body key", func(t *testing.T) {
		model := NewServicedRequest()

		writer := newPermissiveWriter()

		require.NoError(t, model.Serialize(writer))

		writer.AssertNotCalled(t, "WriteStringValue", bodyKey, mock.Anything)
	})
}

// newPermissiveWriter returns a serialization writer that accepts any write, for tests that
// assert on which calls happened rather than on failures.
func newPermissiveWriter() *mocking.MockSerializationWriter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteInt64Value", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteISODurationValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

	return writer
}
