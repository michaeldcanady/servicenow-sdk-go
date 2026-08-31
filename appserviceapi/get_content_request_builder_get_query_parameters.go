// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

// GetContentRequestBuilderGetQueryParameters represents the query parameters for a getContent request.
type GetContentRequestBuilderGetQueryParameters struct {
	// Mode controls whether to return minimum or full details of each CI (shallow or full).
	Mode *string `uriparametername:"mode"`
}
