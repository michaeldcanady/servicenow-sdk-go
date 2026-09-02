// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package accountapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// AccountRequestBuilderGetRequestConfiguration represents the configuration for a GET request.
type AccountRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[AccountRequestBuilderGetQueryParameters]
