// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// GetContentRequestConfiguration represents the configuration for a getContent request.
type GetContentRequestConfiguration = abstractions.RequestConfiguration[GetContentRequestBuilderGetQueryParameters]
