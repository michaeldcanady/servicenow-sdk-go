package appointmentbookingapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type AppointmentRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewAppointmentRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppointmentRequestBuilder {
	return &AppointmentRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/appointment", pathParameters),
	}
}

// Post sends a POST request to book or reschedule an appointment.
func (rB *AppointmentRequestBuilder) Post(ctx context.Context, body AppointmentRequest, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (AppointmentResponse, error) {
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAppointmentResponseFromDiscriminatorValue, nil)
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
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return requestInfo, nil
}
