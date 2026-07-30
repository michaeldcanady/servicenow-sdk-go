package appserviceapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Nil-builder guard coverage for request builders in this package that did
// not otherwise have a dedicated test exercising the guard.

func TestRegisterServiceRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *RegisterServiceRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestRegisterServiceRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *RegisterServiceRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestCreateRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *CreateRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestCreateRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *CreateRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestFindServiceRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *FindServiceRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestFindServiceRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *FindServiceRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestPopulateServiceRequestBuilder_Put_NilBuilder(t *testing.T) {
	var builder *PopulateServiceRequestBuilder

	resp, err := builder.Put(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestPopulateServiceRequestBuilder_ToPutRequestInformation_NilBuilder(t *testing.T) {
	var builder *PopulateServiceRequestBuilder

	requestInfo, err := builder.ToPutRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestServiceDetailsRequestBuilder_Put_NilBuilder(t *testing.T) {
	var builder *ServiceDetailsRequestBuilder

	resp, err := builder.Put(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestServiceDetailsRequestBuilder_ToPutRequestInformation_NilBuilder(t *testing.T) {
	var builder *ServiceDetailsRequestBuilder

	requestInfo, err := builder.ToPutRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestRegisterServiceRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewRegisterServiceRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestCreateRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewCreateRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestFindServiceRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewFindServiceRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestPopulateServiceRequestBuilder_Put_NilRequestAdapter(t *testing.T) {
	builder := NewPopulateServiceRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Put(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestServiceDetailsRequestBuilder_Put_NilRequestAdapter(t *testing.T) {
	builder := NewServiceDetailsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Put(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

// ---------------------------------------------------------------------------
// Child-builder nil-receiver guards (AppServiceRequestBuilder, CsdmRequestBuilder,
// CsdmAppServiceItemRequestBuilder)
// ---------------------------------------------------------------------------

func TestAppServiceRequestBuilder_Create_NilReceiver(t *testing.T) {
	var builder *AppServiceRequestBuilder
	assert.Nil(t, builder.Create())
}

func TestAppServiceRequestBuilder_Csdm_NilReceiver(t *testing.T) {
	var builder *AppServiceRequestBuilder
	assert.Nil(t, builder.Csdm())
}

func TestCsdmRequestBuilder_FindService_NilReceiver(t *testing.T) {
	var builder *CsdmRequestBuilder
	assert.Nil(t, builder.FindService())
}

func TestCsdmRequestBuilder_RegisterService_NilReceiver(t *testing.T) {
	var builder *CsdmRequestBuilder
	assert.Nil(t, builder.RegisterService())
}

func TestCsdmRequestBuilder_ByID_NilReceiver(t *testing.T) {
	var builder *CsdmRequestBuilder
	assert.Nil(t, builder.ByID("service123"))
}

func TestCsdmAppServiceItemRequestBuilder_PopulateService_NilReceiver(t *testing.T) {
	var builder *CsdmAppServiceItemRequestBuilder
	assert.Nil(t, builder.PopulateService())
}

func TestCsdmAppServiceItemRequestBuilder_ServiceDetails_NilReceiver(t *testing.T) {
	var builder *CsdmAppServiceItemRequestBuilder
	assert.Nil(t, builder.ServiceDetails())
}
