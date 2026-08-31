// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateServiceResponse represents the response containing the created application service details.
type CreateServiceResponse interface {
	core.ServiceNowItemResponse[*CreateServiceResult]
}

// CreateCreateServiceResponseFromDiscriminatorValue creates a new CreateServiceResponse from a ParseNode.
func CreateCreateServiceResponseFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return core.ServiceNowItemResponseFromDiscriminatorValue[*CreateServiceResult](CreateCreateServiceResultFromDiscriminatorValue)(parseNode)
}
