// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package tableapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableRequestBuilderGetRequestConfiguration represents the configuration for a Table collection GET request.
type TableRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[TableRequestBuilderGetQueryParameters]
