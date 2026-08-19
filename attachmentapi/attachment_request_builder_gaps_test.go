package attachmentapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const gapsTestRawURL = "https://example.service-now.com/api/now/attachment"

// errAdapter is a stand-in for a transport-level failure from the request adapter.
var errAdapter = errors.New("adapter failure")

func TestNewAttachmentRequestBuilder(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}

	builder := NewAttachmentRequestBuilder(gapsTestRawURL, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, gapsTestRawURL, builder.GetPathParameters()[internal.RawURLKey])
	assert.Equal(t, adapter, builder.GetRequestAdapter())
}

// TestAttachmentRequestBuilder_Head covers the request lifecycle past the nil guard. Head is
// unusual here: it goes through SendNoContent and hands back the inspected response headers
// rather than a deserialized body.
func TestAttachmentRequestBuilder_Head(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path returns inspected response headers",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errAdapter)
			},
			wantErr: errAdapter,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			test.setupMock(adapter)
			builder := NewAttachmentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			headers, err := builder.Head(context.Background(), nil)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, headers)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, headers)
			}

			adapter.AssertExpectations(t)
		})
	}

	t.Run("registers header inspection", func(t *testing.T) {
		adapter := &mocking.MockRequestAdapter{}
		adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		builder := NewAttachmentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

		config := &AttachmentRequestBuilderGetRequestConfiguration{}
		_, err := builder.Head(context.Background(), config)

		require.NoError(t, err)
		assert.NotEmpty(t, config.Options, "Head must register a header-inspection option")
	})
}

func TestAttachmentRequestBuilder_ToHeadRequestInformation(t *testing.T) {
	t.Run("builds a HEAD request", func(t *testing.T) {
		builder := NewAttachmentRequestBuilderInternal(
			map[string]string{"baseurl": "https://example.com"},
			&mocking.MockRequestAdapter{},
		)

		requestInfo, err := builder.ToHeadRequestInformation(context.Background(), nil)

		require.NoError(t, err)
		require.NotNil(t, requestInfo)
		assert.Equal(t, abstractions.HEAD, requestInfo.Method)
		assert.Equal(t, attachmentURLTemplate, requestInfo.UrlTemplate)
	})

	t.Run("nil builder", func(t *testing.T) {
		var builder *AttachmentRequestBuilder

		requestInfo, err := builder.ToHeadRequestInformation(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, requestInfo)
	})

	t.Run("nil inner request builder", func(t *testing.T) {
		requestInfo, err := (&AttachmentRequestBuilder{}).ToHeadRequestInformation(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, requestInfo)
	})
}

// TestAttachmentRequestBuilder_NavigationNilReceivers covers the nil-receiver guard on each
// navigation method, which returns a nil child rather than panicking.
func TestAttachmentRequestBuilder_NavigationNilReceivers(t *testing.T) {
	var builder *AttachmentRequestBuilder

	assert.Nil(t, builder.ByID("sys_id"))
	assert.Nil(t, builder.File())
	assert.Nil(t, builder.Upload())
}

// TestAttachmentRequestBuilder_NavigationClones guards the path-parameter clone: a child must
// not be able to write back into its parent.
func TestAttachmentRequestBuilder_NavigationClones(t *testing.T) {
	parent := NewAttachmentRequestBuilderInternal(
		map[string]string{"baseurl": "https://example.com"},
		&mocking.MockRequestAdapter{},
	)

	item := parent.ByID("sys_id_123")

	require.NotNil(t, item)
	assert.Equal(t, "sys_id_123", item.GetPathParameters()[sysIDKey])
	assert.NotContains(t, parent.GetPathParameters(), sysIDKey)
}

func TestCreateFileFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateFileFromDiscriminatorValue(nil)

	require.NoError(t, err)
	assert.NotNil(t, parsable)
	assert.IsType(t, &File{}, parsable)
}

// TestFile_CreatedByAlias covers GetCreatedBy/SetCreatedBy, which are aliases over the
// sys_created_by accessors.
func TestFile_CreatedByAlias(t *testing.T) {
	file := NewFile()

	require.NoError(t, file.SetCreatedBy(internal.ToPointer("admin")))

	createdBy, err := file.GetCreatedBy()
	require.NoError(t, err)
	require.NotNil(t, createdBy)
	assert.Equal(t, "admin", *createdBy)

	// The alias must read and write the same backing-store entry as the sys_created_by pair.
	sysCreatedBy, err := file.GetSysCreatedBy()
	require.NoError(t, err)
	require.NotNil(t, sysCreatedBy)
	assert.Equal(t, "admin", *sysCreatedBy)
}

// TestAttachmentRequestBuilder_NilInnerBuilder covers a gap between the guards: Get and Head
// only check the outer receiver, while ToGetRequestInformation and ToHeadRequestInformation
// also reject a nil embedded RequestBuilder. A builder with a nil inner therefore gets past
// the verb's own guard and fails when building the request information.
func TestAttachmentRequestBuilder_NilInnerBuilder(t *testing.T) {
	ctx := context.Background()

	t.Run("Get", func(t *testing.T) {
		response, err := (&AttachmentRequestBuilder{}).Get(ctx, nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, response)
	})

	t.Run("Head", func(t *testing.T) {
		headers, err := (&AttachmentRequestBuilder{}).Head(ctx, nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, headers)
	})

	t.Run("ToGetRequestInformation", func(t *testing.T) {
		requestInfo, err := (&AttachmentRequestBuilder{}).ToGetRequestInformation(ctx, nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, requestInfo)
	})
}

// TestAttachmentRequestBuilder_GetWrongResponseType covers Get's failed type assertion.
func TestAttachmentRequestBuilder_GetWrongResponseType(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)

	builder := NewAttachmentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	response, err := builder.Get(context.Background(), nil)

	require.ErrorContains(t, err, "res is not *AttachmentCollectionResponse")
	assert.Nil(t, response)
}
