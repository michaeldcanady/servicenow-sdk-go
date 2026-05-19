package appointmentbookingapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// AvailabilityRequestBuilder provides operations to manage availability.
type AvailabilityRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAvailabilityRequestBuilder instantiates a new AvailabilityRequestBuilder.
func NewAvailabilityRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AvailabilityRequestBuilder {
	return &AvailabilityRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/availability", pathParameters),
	}
}

// Post sends a POST request to get availability.
func (rB *AvailabilityRequestBuilder) Post(ctx context.Context, body AvailabilityRequest, config *AvailabilityRequestBuilderPostRequestConfiguration) (AvailabilityResponse, error) {
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAvailabilityResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(AvailabilityResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *AvailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body AvailabilityRequest, config *AvailabilityRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// CalendarRequestBuilder provides operations to manage calendar.
type CalendarRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCalendarRequestBuilder instantiates a new CalendarRequestBuilder.
func NewCalendarRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CalendarRequestBuilder {
	return &CalendarRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/calendar{?catalog_id,location,opened_for}", pathParameters),
	}
}

// Get sends a GET request to retrieve calendar.
func (rB *CalendarRequestBuilder) Get(ctx context.Context, config *CalendarRequestBuilderGetRequestConfiguration) (CalendarResponse, error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCalendarResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(CalendarResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CalendarRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CalendarRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !internal.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}

// ConfigurationRequestBuilder provides operations to manage configuration.
type ConfigurationRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewConfigurationRequestBuilder instantiates a new ConfigurationRequestBuilder.
func NewConfigurationRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ConfigurationRequestBuilder {
	return &ConfigurationRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/configuration{?catalog_id}", pathParameters),
	}
}

// Get sends a GET request to retrieve configuration.
func (rB *ConfigurationRequestBuilder) Get(ctx context.Context, config *ConfigurationRequestBuilderGetRequestConfiguration) (ConfigurationResponse, error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateConfigurationResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(ConfigurationResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *ConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, config *ConfigurationRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !internal.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}

// ExecuteRuleConditionsRequestBuilder provides operations to manage execute rule conditions.
type ExecuteRuleConditionsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewExecuteRuleConditionsRequestBuilder instantiates a new ExecuteRuleConditionsRequestBuilder.
func NewExecuteRuleConditionsRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExecuteRuleConditionsRequestBuilder {
	return &ExecuteRuleConditionsRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/execute_rule_conditions", pathParameters),
	}
}

// Post sends a POST request to execute rule conditions.
func (rB *ExecuteRuleConditionsRequestBuilder) Post(ctx context.Context, body ExecuteRuleConditionsRequest, config *ExecuteRuleConditionsRequestBuilderPostRequestConfiguration) (ExecuteRuleConditionsResponse, error) {
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExecuteRuleConditionsResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(ExecuteRuleConditionsResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *ExecuteRuleConditionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExecuteRuleConditionsRequest, config *ExecuteRuleConditionsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// UserWindowRequestBuilder provides operations to manage user window.
type UserWindowRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewUserWindowRequestBuilder instantiates a new UserWindowRequestBuilder.
func NewUserWindowRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserWindowRequestBuilder {
	return &UserWindowRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/userwindow", pathParameters),
	}
}

// Post sends a POST request to get user window.
func (rB *UserWindowRequestBuilder) Post(ctx context.Context, body any, config *UserWindowRequestBuilderPostRequestConfiguration) (any, error) {
	// Generic implementation since details are missing
	return nil, nil
}
