package tableapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const itemTestRawURL = "https://example.com/api/now/v1/table/test/sysid"

// newFailingBodyAdapter returns an adapter whose serialization writer fails on write. A
// body-write failure is the only way the ToPut/ToPatch request-information builders can fail
// once past their nil guard, so it is the only way to reach Put's and Patch's error branch.
func newFailingBodyAdapter() *mocking.MockRequestAdapter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(errTransport)
	writer.On("Close").Return(nil)

	factory := mocking.NewMockSerializationWriterFactory()
	factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

	adapter := new(mocking.MockRequestAdapter)
	adapter.On("GetSerializationWriterFactory").Return(factory)

	return adapter
}

// newWritingAdapter returns an adapter that serializes a body successfully, so tests can reach
// the Send call and exercise what the adapter returns.
func newWritingAdapter() *mocking.MockRequestAdapter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(nil)
	writer.On("Close").Return(nil)
	writer.On("GetSerializedContent").Return([]byte("{}"), nil)

	factory := mocking.NewMockSerializationWriterFactory()
	factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

	adapter := new(mocking.MockRequestAdapter)
	adapter.On("GetSerializationWriterFactory").Return(factory)

	return adapter
}

// bodyVerb names one of the two verbs that carry a request body, so the shared failure matrix
// below can be applied to both.
type bodyVerb struct {
	name string
	call func(builder *TableItemRequestBuilder[*TableRecord], body *TableRecord) (any, error)
}

func bodyVerbs() []bodyVerb {
	ctx := context.Background()

	return []bodyVerb{
		{
			name: "Put",
			call: func(builder *TableItemRequestBuilder[*TableRecord], body *TableRecord) (any, error) {
				return builder.Put(ctx, body, nil)
			},
		},
		{
			name: "Patch",
			call: func(builder *TableItemRequestBuilder[*TableRecord], body *TableRecord) (any, error) {
				return builder.Patch(ctx, body, nil)
			},
		},
	}
}

// TestTableItemRequestBuilder_BodyVerbFailures covers the branches Put and Patch share past
// their nil-body guard: a body that fails to serialize, a transport error, an empty response,
// and a response of the wrong type.
func TestTableItemRequestBuilder_BodyVerbFailures(t *testing.T) {
	tests := []struct {
		name    string
		adapter func() *mocking.MockRequestAdapter
		wantErr error
	}{
		{
			name:    "body serialization failure",
			adapter: newFailingBodyAdapter,
			wantErr: errTransport,
		},
		{
			name: "adapter error propagates",
			adapter: func() *mocking.MockRequestAdapter {
				adapter := newWritingAdapter()
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errTransport)

				return adapter
			},
			wantErr: errTransport,
		},
		{
			name: "nil response returns ErrNilResponse",
			adapter: func() *mocking.MockRequestAdapter {
				adapter := newWritingAdapter()
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

				return adapter
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, verb := range bodyVerbs() {
		for _, test := range tests {
			t.Run(verb.name+"/"+test.name, func(t *testing.T) {
				adapter := test.adapter()
				builder := NewDefaultTableItemRequestBuilder(itemTestRawURL, adapter)

				response, err := verb.call(builder, NewTableRecord())

				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, response)
			})
		}

		t.Run(verb.name+"/wrong response type", func(t *testing.T) {
			adapter := newWritingAdapter()
			adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(mocking.NewMockParsable(), nil)
			builder := NewDefaultTableItemRequestBuilder(itemTestRawURL, adapter)

			response, err := verb.call(builder, NewTableRecord())

			require.ErrorContains(t, err, "resp is not")
			assert.Nil(t, response)
		})
	}
}

// TestTableItemRequestBuilder_GetWrongResponseType covers Get's failed type assertion, the one
// Get branch its existing tests leave out.
func TestTableItemRequestBuilder_GetWrongResponseType(t *testing.T) {
	adapter := new(mocking.MockRequestAdapter)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)

	builder := NewDefaultTableItemRequestBuilder(itemTestRawURL, adapter)

	response, err := builder.Get(context.Background(), nil)

	require.ErrorContains(t, err, "resp is not")
	assert.Nil(t, response)
}

// TestTableItemRequestBuilder_ToRequestInformationNilGuards covers the nil-receiver guard on
// each of the four request-information builders.
func TestTableItemRequestBuilder_ToRequestInformationNilGuards(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		call func(builder *TableItemRequestBuilder[*TableRecord]) (any, error)
	}{
		{
			name: "ToGetRequestInformation",
			call: func(b *TableItemRequestBuilder[*TableRecord]) (any, error) {
				return b.ToGetRequestInformation(ctx, nil)
			},
		},
		{
			name: "ToDeleteRequestInformation",
			call: func(b *TableItemRequestBuilder[*TableRecord]) (any, error) {
				return b.ToDeleteRequestInformation(ctx, nil)
			},
		},
		{
			name: "ToPutRequestInformation",
			call: func(b *TableItemRequestBuilder[*TableRecord]) (any, error) {
				return b.ToPutRequestInformation(ctx, NewTableRecord(), nil)
			},
		},
		{
			name: "ToPatchRequestInformation",
			call: func(b *TableItemRequestBuilder[*TableRecord]) (any, error) {
				return b.ToPatchRequestInformation(ctx, NewTableRecord(), nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name+"/nil builder", func(t *testing.T) {
			var builder *TableItemRequestBuilder[*TableRecord]

			requestInfo, err := test.call(builder)

			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, requestInfo)
		})

		t.Run(test.name+"/nil inner request builder", func(t *testing.T) {
			builder := &TableItemRequestBuilder[*TableRecord]{}

			requestInfo, err := test.call(builder)

			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, requestInfo)
		})
	}
}

// TestTableItemRequestBuilder_ToBodyRequestInformationSerializationFailure covers the
// body-write error branch inside ToPutRequestInformation and ToPatchRequestInformation.
func TestTableItemRequestBuilder_ToBodyRequestInformationSerializationFailure(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		call func(builder *TableItemRequestBuilder[*TableRecord]) (any, error)
	}{
		{
			name: "ToPutRequestInformation",
			call: func(b *TableItemRequestBuilder[*TableRecord]) (any, error) {
				return b.ToPutRequestInformation(ctx, NewTableRecord(), nil)
			},
		},
		{
			name: "ToPatchRequestInformation",
			call: func(b *TableItemRequestBuilder[*TableRecord]) (any, error) {
				return b.ToPatchRequestInformation(ctx, NewTableRecord(), nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			builder := NewDefaultTableItemRequestBuilder(itemTestRawURL, newFailingBodyAdapter())

			requestInfo, err := test.call(builder)

			require.ErrorIs(t, err, errTransport)
			assert.Nil(t, requestInfo)
		})
	}
}

// TestNewDefaultTableItemRequestBuilders covers the two default-parsable convenience
// constructors.
func TestNewDefaultTableItemRequestBuilders(t *testing.T) {
	adapter := new(mocking.MockRequestAdapter)

	t.Run("from raw URL", func(t *testing.T) {
		builder := NewDefaultTableItemRequestBuilder(itemTestRawURL, adapter)

		require.NotNil(t, builder)
		assert.Equal(t, itemTestRawURL, builder.GetPathParameters()["request-raw-url"])
		assert.Equal(t, adapter, builder.GetRequestAdapter())
	})

	t.Run("from path parameters", func(t *testing.T) {
		pathParameters := map[string]string{"baseurl": "https://example.com", "table": "incident"}

		builder := NewDefaultTableItemRequestBuilderInternal(pathParameters, adapter)

		require.NotNil(t, builder)
		assert.Equal(t, pathParameters, builder.GetPathParameters())
		assert.Equal(t, adapter, builder.GetRequestAdapter())
	})
}
