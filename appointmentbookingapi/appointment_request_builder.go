package appointmentbookingapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type AppointmentRequestBuilder struct {
	core.RequestBuilder
}

func NewAppointmentRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppointmentRequestBuilder {
	return &AppointmentRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/appointment", pathParameters),
	}
}

// Post sends a POST request to book or reschedule an appointment.
func (rB *AppointmentRequestBuilder) Post(ctx context.Context, body AppointmentRequest, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (AppointmentResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAppointmentResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(AppointmentResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *AppointmentRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppointmentRequest, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
