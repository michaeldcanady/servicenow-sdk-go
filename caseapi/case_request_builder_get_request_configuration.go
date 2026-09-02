// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package caseapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// CaseRequestBuilderGetRequestConfiguration represents configuration for GET /case.
type CaseRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CaseRequestBuilderGetQueryParameters]
