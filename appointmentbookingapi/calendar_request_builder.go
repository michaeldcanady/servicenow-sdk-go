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

const appointmentCalendarBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/calendar{?catalog_id,location,opened_for}"

// CalendarRequestBuilder provides operations to manage calendar.
type CalendarRequestBuilder struct {
	core.RequestBuilder
}

// NewCalendarRequestBuilderInternal instantiates a new [CalendarRequestBuilder].
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CalendarRequestBuilder {
	return &CalendarRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentCalendarBookingURLTemplate, pathParameters),
	}
}

// NewCalendarRequestBuilder instantiates a new [CalendarRequestBuilder] with the provided base URL
// and request adapter.
func NewCalendarRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *CalendarRequestBuilder {
	return NewCalendarRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get sends a GET request to retrieve calendar.
func (rB *CalendarRequestBuilder) Get(ctx context.Context, config *CalendarRequestBuilderGetRequestConfiguration) (CalendarItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCalendarItemResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	typedResp, ok := res.(CalendarItemResponse)
	if !ok {
		// TODO: standardize error
		return nil, errors.New("unexpected type")
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CalendarRequestBuilder) ToGetRequestInformation(_ context.Context, config *CalendarRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
