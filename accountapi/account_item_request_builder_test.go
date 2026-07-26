package accountapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountItemRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *AccountItemRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestAccountItemRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *AccountItemRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestAccountItemRequestBuilder_Get_Success(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "account_id": "test-id"}, adapter)

	require.NotNil(t, builder)
}
