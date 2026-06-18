package appointmentbookingapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type AppointmentRequestBuilder struct {
	internal.RequestBuilder
}

func NewAppointmentRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppointmentRequestBuilder {
	return &AppointmentRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/appointment", pathParameters),
	}
}

// Post sends a POST request to book or reschedule an appointment.
func (rB *AppointmentRequestBuilder) Post(ctx context.Context, body AppointmentRequest, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (AppointmentResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAppointmentResponseFromDiscriminatorValue, internal.DefaultErrorMapping())
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
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
