// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// ExportsRequestBuilderGetRequestConfiguration represents the GET request configuration for the Exports resource.
type ExportsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ExportsRequestBuilderGetQueryParameters]
