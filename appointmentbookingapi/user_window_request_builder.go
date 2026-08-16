package appointmentbookingapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appointmentUserWindowBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/userwindow"

// UserWindowRequestBuilder provides operations to manage the user window.
type UserWindowRequestBuilder struct {
	core.RequestBuilder
}

// NewUserWindowRequestBuilderInternal instantiates a new UserWindowRequestBuilder.
func NewUserWindowRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserWindowRequestBuilder {
	return &UserWindowRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentUserWindowBookingURLTemplate, pathParameters),
	}
}

// NewUserWindowRequestBuilder instantiates a new [UserWindowRequestBuilder] with the provided base URL
// and request adapter.
func NewUserWindowRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UserWindowRequestBuilder {
	return NewUserWindowRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to get the user window.
func (rB *UserWindowRequestBuilder) Post(ctx context.Context, body *UserWindowRequest, config *UserWindowRequestBuilderPostRequestConfiguration) (UserWindowResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUserWindowResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	typedResp, ok := res.(UserWindowResponse)
	if !ok {
		// TODO: standardize error
		return nil, errors.New("unexpected type")
	}

	return typedResp, nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *UserWindowRequestBuilder) ToPostRequestInformation(ctx context.Context, body *UserWindowRequest, config *UserWindowRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return requestInfo, nil
}
