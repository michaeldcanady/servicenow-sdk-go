// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package aggregationapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// StatsRequestBuilderGetRequestConfiguration represents the configuration for a Stats API GET request.
type StatsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[StatsRequestBuilderGetQueryParameters]
