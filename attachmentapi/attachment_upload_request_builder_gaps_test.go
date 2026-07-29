package attachmentapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// failingMultipartBody is a MultipartBody whose GetPartValue fails for one named part. The real
// multipartBody only errors on an empty part name, which the SDK never passes, so a stub is the
// only way to reach the error branches guarding each required part.
type failingMultipartBody struct {
	failOn string
	parts  map[string]any
}

func newFailingMultipartBody(failOn string) *failingMultipartBody {
	return &failingMultipartBody{
		failOn: failOn,
		parts: map[string]any{
			"Content-Type": "text/plain",
			"table_name":   "incident",
			"table_sys_id": "d71f7935c0a80167",
			"uploadFile":   "contents",
		},
	}
}

func (b *failingMultipartBody) GetPartValue(name string) (any, error) {
	if name == b.failOn {
		return nil, errAdapter
	}

	return b.parts[name], nil
}

func (b *failingMultipartBody) AddOrReplacePart(name string, _ string, content any) error {
	b.parts[name] = content

	return nil
}

func (b *failingMultipartBody) RemovePart(name string) error {
	delete(b.parts, name)

	return nil
}

func (b *failingMultipartBody) SetRequestAdapter(abstractions.RequestAdapter) {}

func (b *failingMultipartBody) GetRequestAdapter() abstractions.RequestAdapter { return nil }

func (b *failingMultipartBody) GetBoundary() string { return "boundary" }

func (b *failingMultipartBody) Serialize(serialization.SerializationWriter) error { return nil }

func (b *failingMultipartBody) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

// newUploadTestBody builds a fully populated real multipart body, so Post gets past validation.
func newUploadTestBody(t *testing.T) abstractions.MultipartBody {
	t.Helper()

	value := "test"
	body := abstractions.NewMultipartBody()
	require.NoError(t, body.AddOrReplacePart("Content-Type", "text/plain", &value))
	require.NoError(t, body.AddOrReplacePart("table_name", "text/plain", &value))
	require.NoError(t, body.AddOrReplacePart("table_sys_id", "text/plain", &value))
	require.NoError(t, body.AddOrReplacePart("uploadFile", "text/plain", &value))

	return body
}

// newSerializingUploadAdapter returns an adapter that can serialize a multipart body, so Post
// reaches its Send call.
func newSerializingUploadAdapter(sendResult any, sendErr error) *mocking.MockRequestAdapter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	writer.On("GetSerializedContent").Return([]byte("content"), nil)
	writer.On("Close").Return(nil)

	factory := mocking.NewMockSerializationWriterFactory()
	factory.On("GetSerializationWriter", mock.Anything).Return(writer, nil)

	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(factory)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(sendResult, sendErr)

	return adapter
}

func TestAttachmentUploadRequestBuilder_Post_Gaps(t *testing.T) {
	ctx := context.Background()

	t.Run("nil body is rejected", func(t *testing.T) {
		builder := NewAttachmentUploadRequestBuilderInternal(
			map[string]string{"baseurl": "https://example.com"},
			mocking.NewMockRequestAdapter(),
		)

		file, err := builder.Post(ctx, nil, nil)

		require.ErrorIs(t, err, snerrors.ErrNilBody)
		assert.Nil(t, file)
	})

	t.Run("a part read failure propagates", func(t *testing.T) {
		for _, part := range []string{"Content-Type", "table_name", "table_sys_id", "uploadFile"} {
			t.Run(part, func(t *testing.T) {
				builder := NewAttachmentUploadRequestBuilderInternal(
					map[string]string{"baseurl": "https://example.com"},
					mocking.NewMockRequestAdapter(),
				)

				file, err := builder.Post(ctx, newFailingMultipartBody(part), nil)

				require.ErrorIs(t, err, errAdapter)
				assert.Nil(t, file)
			})
		}
	})

	t.Run("a nil response returns no file and no error", func(t *testing.T) {
		builder := NewAttachmentUploadRequestBuilderInternal(
			map[string]string{"baseurl": "https://example.com"},
			newSerializingUploadAdapter(nil, nil),
		)

		file, err := builder.Post(ctx, newUploadTestBody(t), nil)

		require.NoError(t, err)
		assert.Nil(t, file)
	})

	t.Run("a response of the wrong type is rejected", func(t *testing.T) {
		builder := NewAttachmentUploadRequestBuilderInternal(
			map[string]string{"baseurl": "https://example.com"},
			newSerializingUploadAdapter(mocking.NewMockParsable(), nil),
		)

		file, err := builder.Post(ctx, newUploadTestBody(t), nil)

		require.ErrorContains(t, err, "resp is not Fileable")
		assert.Nil(t, file)
	})

	t.Run("a body serialization failure propagates", func(t *testing.T) {
		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(errAdapter)
		writer.On("Close").Return(nil)

		factory := mocking.NewMockSerializationWriterFactory()
		factory.On("GetSerializationWriter", mock.Anything).Return(writer, nil)

		adapter := mocking.NewMockRequestAdapter()
		adapter.On("GetSerializationWriterFactory").Return(factory)

		builder := NewAttachmentUploadRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

		file, err := builder.Post(ctx, newUploadTestBody(t), nil)

		require.ErrorIs(t, err, errAdapter)
		assert.Nil(t, file)
		adapter.AssertNotCalled(t, "Send")
	})
}

func TestAttachmentUploadRequestBuilder_ToPostRequestInformation_NilInner(t *testing.T) {
	requestInfo, err := (&AttachmentUploadRequestBuilder{}).ToPostRequestInformation(
		context.Background(),
		abstractions.NewMultipartBody(),
		nil,
	)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}
