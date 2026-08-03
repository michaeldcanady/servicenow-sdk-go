package appointmentbookingapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AvailabilityRequestBuilder provides operations to manage availability.
type AvailabilityRequestBuilder struct {
	core.RequestBuilder
}

// NewAvailabilityRequestBuilderInternal instantiates a new AvailabilityRequestBuilder.
func NewAvailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AvailabilityRequestBuilder {
	return &AvailabilityRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/availability", pathParameters),
	}
}

// NewAvailabilityRequestBuilder instantiates a new [AvailabilityRequestBuilder] with the provided base URL
// and request adapter.
func NewAvailabilityRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AvailabilityRequestBuilder {
	return NewAvailabilityRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to get availability.
func (rB *AvailabilityRequestBuilder) Post(ctx context.Context, body AvailabilityRequest, config *AvailabilityRequestBuilderPostRequestConfiguration) (AvailabilityResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAvailabilityResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(AvailabilityResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *AvailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body AvailabilityRequest, config *AvailabilityRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// CalendarRequestBuilder provides operations to manage calendar.
type CalendarRequestBuilder struct {
	core.RequestBuilder
}

// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CalendarRequestBuilder {
	return &CalendarRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/calendar{?catalog_id,location,opened_for}", pathParameters),
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
func (rB *CalendarRequestBuilder) Get(ctx context.Context, config *CalendarRequestBuilderGetRequestConfiguration) (*CalendarResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCalendarResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*CalendarResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CalendarRequestBuilder) ToGetRequestInformation(_ context.Context, config *CalendarRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// ConfigurationRequestBuilder provides operations to manage configuration.
type ConfigurationRequestBuilder struct {
	core.RequestBuilder
}

// NewConfigurationRequestBuilderInternal instantiates a new ConfigurationRequestBuilder.
func NewConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ConfigurationRequestBuilder {
	return &ConfigurationRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/configuration{?catalog_id}", pathParameters),
	}
}

// NewConfigurationRequestBuilder instantiates a new [ConfigurationRequestBuilder] with the provided base URL
// and request adapter.
func NewConfigurationRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ConfigurationRequestBuilder {
	return NewConfigurationRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get sends a GET request to retrieve configuration.
func (rB *ConfigurationRequestBuilder) Get(ctx context.Context, config *ConfigurationRequestBuilderGetRequestConfiguration) (ConfigurationResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateConfigurationResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(ConfigurationResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *ConfigurationRequestBuilder) ToGetRequestInformation(_ context.Context, config *ConfigurationRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// ExecuteRuleConditionsRequestBuilder provides operations to manage execute rule conditions.
type ExecuteRuleConditionsRequestBuilder struct {
	core.RequestBuilder
}

// NewExecuteRuleConditionsRequestBuilderInternal instantiates a new ExecuteRuleConditionsRequestBuilder.
func NewExecuteRuleConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExecuteRuleConditionsRequestBuilder {
	return &ExecuteRuleConditionsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/execute_rule_conditions", pathParameters),
	}
}

// NewExecuteRuleConditionsRequestBuilder instantiates a new [ExecuteRuleConditionsRequestBuilder] with the provided base URL
// and request adapter.
func NewExecuteRuleConditionsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ExecuteRuleConditionsRequestBuilder {
	return NewExecuteRuleConditionsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to execute rule conditions.
func (rB *ExecuteRuleConditionsRequestBuilder) Post(ctx context.Context, body ExecuteRuleConditionsRequest, config *ExecuteRuleConditionsRequestBuilderPostRequestConfiguration) (ExecuteRuleConditionsResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExecuteRuleConditionsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(ExecuteRuleConditionsResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *ExecuteRuleConditionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExecuteRuleConditionsRequest, config *ExecuteRuleConditionsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// UserWindowRequestBuilder provides operations to manage user window.
type UserWindowRequestBuilder struct {
	core.RequestBuilder
}

// NewUserWindowRequestBuilderInternal instantiates a new UserWindowRequestBuilder.
func NewUserWindowRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserWindowRequestBuilder {
	return &UserWindowRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/userwindow", pathParameters),
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

// Post sends a POST request to get user window.
func (rB *UserWindowRequestBuilder) Post(ctx context.Context, body any, config *UserWindowRequestBuilderPostRequestConfiguration) (AvailabilityResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAvailabilityResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(AvailabilityResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *UserWindowRequestBuilder) ToPostRequestInformation(ctx context.Context, body any, config *UserWindowRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body.(serialization.Parsable))
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}
