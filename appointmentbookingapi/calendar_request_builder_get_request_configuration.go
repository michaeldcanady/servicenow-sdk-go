// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// CalendarRequestBuilderGetRequestConfiguration represents the configuration for GET /calendar.
type CalendarRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CalendarRequestBuilderGetQueryParameters]
