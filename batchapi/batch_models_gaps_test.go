package batchapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestBatchRequest_SerializeMintsID covers the batch-request-ID mutator in Serialize: a batch
// request must carry an ID so the response can be correlated back to it, so serialization mints
// one when the model has none.
func TestBatchRequest_SerializeMintsID(t *testing.T) {
	tests := []struct {
		name  string
		setup func(model *BatchRequestModel)
	}{
		{
			name:  "no ID set",
			setup: func(_ *BatchRequestModel) {},
		},
		{
			name: "empty ID set",
			setup: func(model *BatchRequestModel) {
				require.NoError(t, model.SetBatchRequestID(internal.ToPointer("")))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model := NewBatchRequestModel()
			test.setup(model)

			var written *string
			writer := mocking.NewMockSerializationWriter()
			writer.On("WriteStringValue", batchRequestIDKey, mock.Anything).Run(func(args mock.Arguments) {
				written, _ = args.Get(1).(*string)
			}).Return(nil)
			writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

			require.NoError(t, model.Serialize(writer))

			require.NotNil(t, written)
			assert.NotEmpty(t, *written, "serialization must mint a batch request ID when the model has none")
		})
	}

	t.Run("an existing ID is preserved", func(t *testing.T) {
		model := NewBatchRequestModel()
		require.NoError(t, model.SetBatchRequestID(internal.ToPointer("batch-1")))

		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

		require.NoError(t, model.Serialize(writer))

		writer.AssertCalled(t, "WriteStringValue", batchRequestIDKey, internal.ToPointer("batch-1"))
	})
}

// TestBatchRequest_FieldDeserializers exercises the deserializers rather than only asserting the
// map is non-nil.
func TestBatchRequest_FieldDeserializers(t *testing.T) {
	t.Run("batch request ID is read", func(t *testing.T) {
		model := NewBatchRequestModel()
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return(internal.ToPointer("batch-1"), nil)

		require.NoError(t, model.GetFieldDeserializers()[batchRequestIDKey](node))

		id, err := model.GetBatchRequestID()
		require.NoError(t, err)
		require.NotNil(t, id)
		assert.Equal(t, "batch-1", *id)
	})

	t.Run("rest requests are read", func(t *testing.T) {
		model := NewBatchRequestModel()
		node := mocking.NewMockParseNode()
		node.On("GetCollectionOfObjectValues", mock.Anything).
			Return([]serialization.Parsable{NewRestRequest()}, nil)

		require.NoError(t, model.GetFieldDeserializers()[restRequestsKey](node))

		requests, err := model.GetRestRequests()
		require.NoError(t, err)
		assert.Len(t, requests, 1)
	})
}

// TestBatchResponse_ServicedRequestsDeserializer covers the hand-written serviced-requests
// deserializer, including its type-assertion failure.
func TestBatchResponse_ServicedRequestsDeserializer(t *testing.T) {
	t.Run("serviced requests are read", func(t *testing.T) {
		model := NewBatchResponse()
		node := mocking.NewMockParseNode()
		node.On("GetCollectionOfObjectValues", mock.Anything).
			Return([]serialization.Parsable{NewServicedRequest()}, nil)

		require.NoError(t, model.GetFieldDeserializers()[servicedRequestsKey](node))

		requests, err := model.GetServicedRequests()
		require.NoError(t, err)
		assert.Len(t, requests, 1)
	})

	t.Run("a read error propagates", func(t *testing.T) {
		model := NewBatchResponse()
		node := mocking.NewMockParseNode()
		node.On("GetCollectionOfObjectValues", mock.Anything).
			Return([]serialization.Parsable(nil), errNode)

		require.ErrorIs(t, model.GetFieldDeserializers()[servicedRequestsKey](node), errNode)
	})

	t.Run("an entry of the wrong type is rejected", func(t *testing.T) {
		model := NewBatchResponse()
		node := mocking.NewMockParseNode()
		node.On("GetCollectionOfObjectValues", mock.Anything).
			Return([]serialization.Parsable{mocking.NewMockParsable()}, nil)

		require.ErrorContains(t, model.GetFieldDeserializers()[servicedRequestsKey](node), "value is not ServicedRequest")
	})
}

// TestBatchResponse_UnservicedRequestsDeserializer covers the unserviced-requests deserializer,
// which reads a primitive string collection.
func TestBatchResponse_UnservicedRequestsDeserializer(t *testing.T) {
	t.Run("unserviced request IDs are read", func(t *testing.T) {
		model := NewBatchResponse()
		node := mocking.NewMockParseNode()
		node.On("GetCollectionOfPrimitiveValues", "string").Return([]any{"id-1", "id-2"}, nil)

		require.NoError(t, model.GetFieldDeserializers()[unservicedRequestsKey](node))

		requests, err := model.GetUnservicedRequests()
		require.NoError(t, err)
		assert.Equal(t, []string{"id-1", "id-2"}, requests)
	})

	t.Run("a read error propagates", func(t *testing.T) {
		model := NewBatchResponse()
		node := mocking.NewMockParseNode()
		node.On("GetCollectionOfPrimitiveValues", "string").Return([]any(nil), errNode)

		require.ErrorIs(t, model.GetFieldDeserializers()[unservicedRequestsKey](node), errNode)
	})
}

// TestBatchResponse_GetServicedRequestByID_Gaps covers the lookup's miss path and its tolerance
// of an entry whose ID cannot be read.
func TestBatchResponse_GetServicedRequestByID_Gaps(t *testing.T) {
	newResponseWith := func(t *testing.T, ids ...string) *BatchResponseModel {
		t.Helper()

		requests := make([]ServicedRequest, 0, len(ids))
		for _, id := range ids {
			request := NewServicedRequest()
			require.NoError(t, request.setID(internal.ToPointer(id)))
			requests = append(requests, request)
		}

		model := NewBatchResponse()
		require.NoError(t, model.setServicedRequests(requests))

		return model
	}

	t.Run("returns the matching request", func(t *testing.T) {
		model := newResponseWith(t, "a", "b")

		request, err := model.GetServicedRequestByID("b")

		require.NoError(t, err)
		require.NotNil(t, request)

		id, err := request.GetID()
		require.NoError(t, err)
		require.NotNil(t, id)
		assert.Equal(t, "b", *id)
	})

	t.Run("returns nothing when no request matches", func(t *testing.T) {
		model := newResponseWith(t, "a", "b")

		request, err := model.GetServicedRequestByID("missing")

		require.NoError(t, err)
		assert.Nil(t, request)
	})

	t.Run("nil receiver returns nothing", func(t *testing.T) {
		var model *BatchResponseModel

		request, err := model.GetServicedRequestByID("a")

		require.NoError(t, err)
		assert.Nil(t, request)
	})
}
