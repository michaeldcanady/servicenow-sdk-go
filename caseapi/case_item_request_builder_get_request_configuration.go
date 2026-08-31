// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package caseapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// CaseItemRequestBuilderGetRequestConfiguration represents configuration for GET /case/{id}.
type CaseItemRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CaseItemRequestBuilderGetQueryParameters]
