// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package caseapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// CaseFieldValuesRequestBuilderGetRequestConfiguration represents configuration for GET /case/field_values/{field_name}.
type CaseFieldValuesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CaseFieldValuesRequestBuilderGetQueryParameters]
