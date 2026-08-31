// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// FindServiceRequestConfiguration represents the configuration for a find_service request.
type FindServiceRequestConfiguration = abstractions.RequestConfiguration[FindServiceQueryParameters]
