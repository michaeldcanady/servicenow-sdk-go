package appointmentbookingapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appointmentAppointmentBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/appointment"

// AppointmentRequestBuilder represents the appointment request builder.
type AppointmentRequestBuilder struct {
	core.RequestBuilder
}

// NewAppointmentRequestBuilderInternal creates a new instance of AppointmentRequestBuilder.
func NewAppointmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppointmentRequestBuilder {
	return &AppointmentRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentAppointmentBookingURLTemplate, pathParameters),
	}
}

// NewAppointmentRequestBuilder instantiates a new [AppointmentRequestBuilder] with the provided base URL
// and request adapter.
func NewAppointmentRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AppointmentRequestBuilder {
	return NewAppointmentRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to book or reschedule an appointment.
func (rB *AppointmentRequestBuilder) Post(ctx context.Context, body AppointmentRequest, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (AppointmentResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAppointmentResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(AppointmentResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*AppointmentResponse)(nil))
	}

	return typedResp, nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *AppointmentRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppointmentRequest, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body); err != nil {
		return nil, err
	}

	return requestInfo, nil
}
