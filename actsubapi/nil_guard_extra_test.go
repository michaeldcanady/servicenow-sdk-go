package actsubapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests cover nil-builder guard paths for request builders that do not
// otherwise have a dedicated _test.go file.

func TestSubscriberItemRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *SubscriberItemRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestSubscriberItemRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *SubscriberItemRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestFollowingItemRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *FollowingItemRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestFollowingItemRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *FollowingItemRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestSubObjectsRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *SubObjectsRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestSubObjectsRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *SubObjectsRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestContextsRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *ContextsRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestContextsRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *ContextsRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestFacetsInstanceRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *FacetsInstanceRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestFacetsInstanceRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewFacetsInstanceRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestFacetsInstanceRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *FacetsInstanceRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}
