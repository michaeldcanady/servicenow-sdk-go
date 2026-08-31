// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CreateRequestConfiguration represents the configuration for a Create request.
type CreateRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
