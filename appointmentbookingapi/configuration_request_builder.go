// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appointmentConfigurationBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/configuration{?catalog_id}"

// ConfigurationRequestBuilder provides operations to manage configuration.
type ConfigurationRequestBuilder struct {
	core.RequestBuilder
}

// NewConfigurationRequestBuilderInternal instantiates a new [ConfigurationRequestBuilder].
func NewConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ConfigurationRequestBuilder {
	return &ConfigurationRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentConfigurationBookingURLTemplate, pathParameters),
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
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(ConfigurationResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*ConfigurationResponse)(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a [abstractions.RequestInformation] object for a GET request.
func (rB *ConfigurationRequestBuilder) ToGetRequestInformation(_ context.Context, config *ConfigurationRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
