// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportsResponse represents a response containing a collection of export results.
type ExportsResponse interface {
	core.ServiceNowCollectionResponse[*ExportResult]
}

// CreateExportsResponseFromDiscriminatorValue instantiates a new ExportsResponse.
func CreateExportsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ExportResult](CreateExportResultFromDiscriminatorValue), nil
}
